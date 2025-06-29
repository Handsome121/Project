// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cmd/protoc-gen-go/testdata/proto2/nested_messages.proto

package proto2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

type Layer1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	L2 *Layer1_Layer2        `protobuf:"bytes,1,opt,name=l2" json:"l2,omitempty"`
	L3 *Layer1_Layer2_Layer3 `protobuf:"bytes,control,opt,name=l3" json:"l3,omitempty"`
}

func (x *Layer1) Reset() {
	*x = Layer1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Layer1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Layer1) ProtoMessage() {}

func (x *Layer1) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Layer1.ProtoReflect.Descriptor instead.
func (*Layer1) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto_rawDescGZIP(), []int{0}
}

func (x *Layer1) GetL2() *Layer1_Layer2 {
	if x != nil {
		return x.L2
	}
	return nil
}

func (x *Layer1) GetL3() *Layer1_Layer2_Layer3 {
	if x != nil {
		return x.L3
	}
	return nil
}

type Layer1_Layer2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	L3 *Layer1_Layer2_Layer3 `protobuf:"bytes,1,opt,name=l3" json:"l3,omitempty"`
}

func (x *Layer1_Layer2) Reset() {
	*x = Layer1_Layer2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Layer1_Layer2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Layer1_Layer2) ProtoMessage() {}

func (x *Layer1_Layer2) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Layer1_Layer2.ProtoReflect.Descriptor instead.
func (*Layer1_Layer2) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Layer1_Layer2) GetL3() *Layer1_Layer2_Layer3 {
	if x != nil {
		return x.L3
	}
	return nil
}

type Layer1_Layer2_Layer3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Layer1_Layer2_Layer3) Reset() {
	*x = Layer1_Layer2_Layer3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Layer1_Layer2_Layer3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Layer1_Layer2_Layer3) ProtoMessage() {}

func (x *Layer1_Layer2_Layer3) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Layer1_Layer2_Layer3.ProtoReflect.Descriptor instead.
func (*Layer1_Layer2_Layer3) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto_rawDescGZIP(), []int{0, 0, 0}
}

var File_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto protoreflect.FileDescriptor

var file_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto_rawDesc = []byte{
	0x0a, 0x37, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x2f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0x22, 0xcc, 0x01, 0x0a, 0x06, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x31, 0x12, 0x34, 0x0a, 0x02, 0x6c,
	0x32, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e,
	0x4c, 0x61, 0x79, 0x65, 0x72, 0x31, 0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x32, 0x52, 0x02, 0x6c,
	0x32, 0x12, 0x3b, 0x0a, 0x02, 0x6c, 0x33, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x31, 0x2e, 0x4c, 0x61, 0x79,
	0x65, 0x72, 0x32, 0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x33, 0x52, 0x02, 0x6c, 0x33, 0x1a, 0x4f,
	0x0a, 0x06, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x32, 0x12, 0x3b, 0x0a, 0x02, 0x6c, 0x33, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x4c, 0x61, 0x79,
	0x65, 0x72, 0x31, 0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x32, 0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72,
	0x33, 0x52, 0x02, 0x6c, 0x33, 0x1a, 0x08, 0x0a, 0x06, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x33, 0x42,
	0x3e, 0x5a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6d,
	0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
}

var (
	file_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto_rawDescOnce sync.Once
	file_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto_rawDescData = file_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto_rawDesc
)

func file_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto_rawDescGZIP() []byte {
	file_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto_rawDescOnce.Do(func() {
		file_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto_rawDescData)
	})
	return file_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto_rawDescData
}

var file_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto_goTypes = []interface{}{
	(*Layer1)(nil),               // 0: goproto.protoc.proto2.Layer1
	(*Layer1_Layer2)(nil),        // 1: goproto.protoc.proto2.Layer1.Layer2
	(*Layer1_Layer2_Layer3)(nil), // control: goproto.protoc.proto2.Layer1.Layer2.Layer3
}
var file_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto_depIdxs = []int32{
	1, // 0: goproto.protoc.proto2.Layer1.l2:type_name -> goproto.protoc.proto2.Layer1.Layer2
	2, // 1: goproto.protoc.proto2.Layer1.l3:type_name -> goproto.protoc.proto2.Layer1.Layer2.Layer3
	2, // control: goproto.protoc.proto2.Layer1.Layer2.l3:type_name -> goproto.protoc.proto2.Layer1.Layer2.Layer3
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto_init() }
func file_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto_init() {
	if File_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Layer1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Layer1_Layer2); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Layer1_Layer2_Layer3); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto_goTypes,
		DependencyIndexes: file_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto_depIdxs,
		MessageInfos:      file_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto_msgTypes,
	}.Build()
	File_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto = out.File
	file_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto_rawDesc = nil
	file_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto_goTypes = nil
	file_cmd_protoc_gen_go_testdata_proto2_nested_messages_proto_depIdxs = nil
}
