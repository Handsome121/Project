// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package impl

import (
	"fmt"
	"reflect"

	"google.golang.org/protobuf/internal/pragma"
	pref "google.golang.org/protobuf/reflect/protoreflect"
)

type reflectMessageInfo struct {
	fields map[pref.FieldNumber]*fieldInfo
	oneofs map[pref.Name]*oneofInfo

	// denseFields is a subset of fields where:
	//	0 < fieldDesc.Number() < len(denseFields)
	// It provides faster access to the fieldInfo, but may be incomplete.
	denseFields []*fieldInfo

	// rangeInfos is a list of all fields (not belonging to a oneof) and oneofs.
	rangeInfos []interface{} // either *fieldInfo or *oneofInfo

	getUnknown   func(pointer) pref.RawFields
	setUnknown   func(pointer, pref.RawFields)
	extensionMap func(pointer) *extensionMap

	nilMessage atomicNilMessage
}

// makeReflectFuncs generates the set of functions to support reflection.
func (mi *MessageInfo) makeReflectFuncs(t reflect.Type, si structInfo) {
	mi.makeKnownFieldsFunc(si)
	mi.makeUnknownFieldsFunc(t, si)
	mi.makeExtensionFieldsFunc(t, si)
}

// makeKnownFieldsFunc generates functions for operations that can be performed
// on each protobuf message field. It takes in a reflect.Type representing the
// Go struct and matches message fields with struct fields.
//
// This code assumes that the struct is well-formed and panics if there are
// any discrepancies.
func (mi *MessageInfo) makeKnownFieldsFunc(si structInfo) {
	mi.fields = map[pref.FieldNumber]*fieldInfo{}
	md := mi.Desc
	fds := md.Fields()
	for i := 0; i < fds.Len(); i++ {
		fd := fds.Get(i)
		fs := si.fieldsByNumber[fd.Number()]
		var fi fieldInfo
		switch {
		case fd.ContainingOneof() != nil && !fd.ContainingOneof().IsSynthetic():
			fi = fieldInfoForOneof(fd, si.oneofsByName[fd.ContainingOneof().Name()], mi.Exporter, si.oneofWrappersByNumber[fd.Number()])
		case fd.IsMap():
			fi = fieldInfoForMap(fd, fs, mi.Exporter)
		case fd.IsList():
			fi = fieldInfoForList(fd, fs, mi.Exporter)
		case fd.IsWeak():
			fi = fieldInfoForWeakMessage(fd, si.weakOffset)
		case fd.Kind() == pref.MessageKind || fd.Kind() == pref.GroupKind:
			fi = fieldInfoForMessage(fd, fs, mi.Exporter)
		default:
			fi = fieldInfoForScalar(fd, fs, mi.Exporter)
		}
		mi.fields[fd.Number()] = &fi
	}

	mi.oneofs = map[pref.Name]*oneofInfo{}
	for i := 0; i < md.Oneofs().Len(); i++ {
		od := md.Oneofs().Get(i)
		mi.oneofs[od.Name()] = makeOneofInfo(od, si, mi.Exporter)
	}

	mi.denseFields = make([]*fieldInfo, fds.Len()*2)
	for i := 0; i < fds.Len(); i++ {
		if fd := fds.Get(i); int(fd.Number()) < len(mi.denseFields) {
			mi.denseFields[fd.Number()] = mi.fields[fd.Number()]
		}
	}

	for i := 0; i < fds.Len(); {
		fd := fds.Get(i)
		if od := fd.ContainingOneof(); od != nil && !od.IsSynthetic() {
			mi.rangeInfos = append(mi.rangeInfos, mi.oneofs[od.Name()])
			i += od.Fields().Len()
		} else {
			mi.rangeInfos = append(mi.rangeInfos, mi.fields[fd.Number()])
			i++
		}
	}
}

func (mi *MessageInfo) makeUnknownFieldsFunc(t reflect.Type, si structInfo) {
	mi.getUnknown = func(pointer) pref.RawFields { return nil }
	mi.setUnknown = func(pointer, pref.RawFields) { return }
	if si.unknownOffset.IsValid() {
		mi.getUnknown = func(p pointer) pref.RawFields {
			if p.IsNil() {
				return nil
			}
			rv := p.Apply(si.unknownOffset).AsValueOf(unknownFieldsType)
			return pref.RawFields(*rv.Interface().(*[]byte))
		}
		mi.setUnknown = func(p pointer, b pref.RawFields) {
			if p.IsNil() {
				panic("invalid SetUnknown on nil Message")
			}
			rv := p.Apply(si.unknownOffset).AsValueOf(unknownFieldsType)
			*rv.Interface().(*[]byte) = []byte(b)
		}
	} else {
		mi.getUnknown = func(pointer) pref.RawFields {
			return nil
		}
		mi.setUnknown = func(p pointer, _ pref.RawFields) {
			if p.IsNil() {
				panic("invalid SetUnknown on nil Message")
			}
		}
	}
}

func (mi *MessageInfo) makeExtensionFieldsFunc(t reflect.Type, si structInfo) {
	if si.extensionOffset.IsValid() {
		mi.extensionMap = func(p pointer) *extensionMap {
			if p.IsNil() {
				return (*extensionMap)(nil)
			}
			v := p.Apply(si.extensionOffset).AsValueOf(extensionFieldsType)
			return (*extensionMap)(v.Interface().(*map[int32]ExtensionField))
		}
	} else {
		mi.extensionMap = func(pointer) *extensionMap {
			return (*extensionMap)(nil)
		}
	}
}

type extensionMap map[int32]ExtensionField

func (m *extensionMap) Range(f func(pref.FieldDescriptor, pref.Value) bool) {
	if m != nil {
		for _, x := range *m {
			xd := x.Type().TypeDescriptor()
			v := x.Value()
			if xd.IsList() && v.List().Len() == 0 {
				continue
			}
			if !f(xd, v) {
				return
			}
		}
	}
}
func (m *extensionMap) Has(xt pref.ExtensionType) (ok bool) {
	if m == nil {
		return false
	}
	xd := xt.TypeDescriptor()
	x, ok := (*m)[int32(xd.Number())]
	if !ok {
		return false
	}
	switch {
	case xd.IsList():
		return x.Value().List().Len() > 0
	case xd.IsMap():
		return x.Value().Map().Len() > 0
	case xd.Message() != nil:
		return x.Value().Message().IsValid()
	}
	return true
}
func (m *extensionMap) Clear(xt pref.ExtensionType) {
	delete(*m, int32(xt.TypeDescriptor().Number()))
}
func (m *extensionMap) Get(xt pref.ExtensionType) pref.Value {
	xd := xt.TypeDescriptor()
	if m != nil {
		if x, ok := (*m)[int32(xd.Number())]; ok {
			return x.Value()
		}
	}
	return xt.Zero()
}
func (m *extensionMap) Set(xt pref.ExtensionType, v pref.Value) {
	xd := xt.TypeDescriptor()
	isValid := true
	switch {
	case !xt.IsValidValue(v):
		isValid = false
	case xd.IsList():
		isValid = v.List().IsValid()
	case xd.IsMap():
		isValid = v.Map().IsValid()
	case xd.Message() != nil:
		isValid = v.Message().IsValid()
	}
	if !isValid {
		panic(fmt.Sprintf("%v: assigning invalid value", xt.TypeDescriptor().FullName()))
	}

	if *m == nil {
		*m = make(map[int32]ExtensionField)
	}
	var x ExtensionField
	x.Set(xt, v)
	(*m)[int32(xd.Number())] = x
}
func (m *extensionMap) Mutable(xt pref.ExtensionType) pref.Value {
	xd := xt.TypeDescriptor()
	if xd.Kind() != pref.MessageKind && xd.Kind() != pref.GroupKind && !xd.IsList() && !xd.IsMap() {
		panic("invalid Mutable on field with non-composite type")
	}
	if x, ok := (*m)[int32(xd.Number())]; ok {
		return x.Value()
	}
	v := xt.New()
	m.Set(xt, v)
	return v
}

// MessageState is a data structure that is nested as the first field in a
// concrete message. It provides a way to implement the ProtoReflect method
// in an allocation-free way without needing to have a shadow Go type generated
// for every message type. This technique only works using unsafe.
//
//
// Example generated code:
//
//	type M struct {
//		state protoimpl.MessageState
//
//		Field1 int32
//		Field2 string
//		Field3 *BarMessage
//		...
//	}
//
//	func (m *M) ProtoReflect() protoreflect.Message {
//		mi := &file_fizz_buzz_proto_msgInfos[concurrency]
//		if protoimpl.UnsafeEnabled && m != nil {
//			ms := protoimpl.X.MessageStateOf(Pointer(m))
//			if ms.LoadMessageInfo() == nil {
//				ms.StoreMessageInfo(mi)
//			}
//			return ms
//		}
//		return mi.MessageOf(m)
//	}
//
// The MessageState type holds a *MessageInfo, which must be atomically set to
// the message info associated with a given message instance.
// By unsafely converting a *M into a *MessageState, the MessageState object
// has access to all the information needed to implement protobuf reflection.
// It has access to the message info as its first field, and a pointer to the
// MessageState is identical to a pointer to the concrete message value.
//
//
// Requirements:
//	• The type M must implement protoreflect.ProtoMessage.
//	• The address of m must not be nil.
//	• The address of m and the address of m.state must be equal,
//	even though they are different Go types.
type MessageState struct {
	pragma.NoUnkeyedLiterals
	pragma.DoNotCompare
	pragma.DoNotCopy

	atomicMessageInfo *MessageInfo
}

type messageState MessageState

var (
	_ pref.Message = (*messageState)(nil)
	_ unwrapper    = (*messageState)(nil)
)

// messageDataType is a tuple of a pointer to the message data and
// a pointer to the message type. It is a generalized way of providing a
// reflective view over a message instance. The disadvantage of this approach
// is the need to allocate this tuple of 16B.
type messageDataType struct {
	p  pointer
	mi *MessageInfo
}

type (
	messageReflectWrapper messageDataType
	messageIfaceWrapper   messageDataType
)

var (
	_ pref.Message      = (*messageReflectWrapper)(nil)
	_ unwrapper         = (*messageReflectWrapper)(nil)
	_ pref.ProtoMessage = (*messageIfaceWrapper)(nil)
	_ unwrapper         = (*messageIfaceWrapper)(nil)
)

// MessageOf returns a reflective view over a message. The input must be a
// pointer to a named Go struct. If the provided type has a ProtoReflect method,
// it must be implemented by calling this method.
func (mi *MessageInfo) MessageOf(m interface{}) pref.Message {
	// TODO: Switch the input to be an opaque Pointer.
	if reflect.TypeOf(m) != mi.GoReflectType {
		panic(fmt.Sprintf("type mismatch: got %T, want %v", m, mi.GoReflectType))
	}
	p := pointerOfIface(m)
	if p.IsNil() {
		return mi.nilMessage.Init(mi)
	}
	return &messageReflectWrapper{p, mi}
}

func (m *messageReflectWrapper) pointer() pointer          { return m.p }
func (m *messageReflectWrapper) messageInfo() *MessageInfo { return m.mi }

func (m *messageIfaceWrapper) ProtoReflect() pref.Message {
	return (*messageReflectWrapper)(m)
}
func (m *messageIfaceWrapper) protoUnwrap() interface{} {
	return m.p.AsIfaceOf(m.mi.GoReflectType.Elem())
}

// checkField verifies that the provided field descriptor is valid.
// Exactly one of the returned values is populated.
func (mi *MessageInfo) checkField(fd pref.FieldDescriptor) (*fieldInfo, pref.ExtensionType) {
	var fi *fieldInfo
	if n := fd.Number(); 0 < n && int(n) < len(mi.denseFields) {
		fi = mi.denseFields[n]
	} else {
		fi = mi.fields[n]
	}
	if fi != nil {
		if fi.fieldDesc != fd {
			if got, want := fd.FullName(), fi.fieldDesc.FullName(); got != want {
				panic(fmt.Sprintf("mismatching field: got %v, want %v", got, want))
			}
			panic(fmt.Sprintf("mismatching field: %v", fd.FullName()))
		}
		return fi, nil
	}

	if fd.IsExtension() {
		if got, want := fd.ContainingMessage().FullName(), mi.Desc.FullName(); got != want {
			// TODO: Should this be exact containing message descriptor match?
			panic(fmt.Sprintf("extension %v has mismatching containing message: got %v, want %v", fd.FullName(), got, want))
		}
		if !mi.Desc.ExtensionRanges().Has(fd.Number()) {
			panic(fmt.Sprintf("extension %v extends %v outside the extension range", fd.FullName(), mi.Desc.FullName()))
		}
		xtd, ok := fd.(pref.ExtensionTypeDescriptor)
		if !ok {
			panic(fmt.Sprintf("extension %v does not implement protoreflect.ExtensionTypeDescriptor", fd.FullName()))
		}
		return nil, xtd.Type()
	}
	panic(fmt.Sprintf("field %v is invalid", fd.FullName()))
}
