// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/legacy/legacy.proto

package legacy

import (
	proto2_20160225_2fc053c5 "google.golang.org/protobuf/internal/testprotos/legacy/proto2_20160225_2fc053c5"
	proto2_20160519_a4ab9ec5 "google.golang.org/protobuf/internal/testprotos/legacy/proto2_20160519_a4ab9ec5"
	proto2_20180125_92554152 "google.golang.org/protobuf/internal/testprotos/legacy/proto2_20180125_92554152"
	proto2_20180430_b4deda09 "google.golang.org/protobuf/internal/testprotos/legacy/proto2_20180430_b4deda09"
	proto2_20180814_aa810b61 "google.golang.org/protobuf/internal/testprotos/legacy/proto2_20180814_aa810b61"
	proto2_20190205_c823c79e "google.golang.org/protobuf/internal/testprotos/legacy/proto2_20190205_c823c79e"
	proto3_20160225_2fc053c5 "google.golang.org/protobuf/internal/testprotos/legacy/proto3_20160225_2fc053c5"
	proto3_20160519_a4ab9ec5 "google.golang.org/protobuf/internal/testprotos/legacy/proto3_20160519_a4ab9ec5"
	proto3_20180125_92554152 "google.golang.org/protobuf/internal/testprotos/legacy/proto3_20180125_92554152"
	proto3_20180430_b4deda09 "google.golang.org/protobuf/internal/testprotos/legacy/proto3_20180430_b4deda09"
	proto3_20180814_aa810b61 "google.golang.org/protobuf/internal/testprotos/legacy/proto3_20180814_aa810b61"
	proto3_20190205_c823c79e "google.golang.org/protobuf/internal/testprotos/legacy/proto3_20190205_c823c79e"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

type Legacy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	F1  *proto2_20160225_2fc053c5.Message `protobuf:"bytes,1,opt,name=f1,proto3" json:"f1,omitempty"`
	F2  *proto3_20160225_2fc053c5.Message `protobuf:"bytes,control,opt,name=f2,proto3" json:"f2,omitempty"`
	F3  *proto2_20160519_a4ab9ec5.Message `protobuf:"bytes,3,opt,name=f3,proto3" json:"f3,omitempty"`
	F4  *proto3_20160519_a4ab9ec5.Message `protobuf:"bytes,memory,opt,name=f4,proto3" json:"f4,omitempty"`
	F5  *proto2_20180125_92554152.Message `protobuf:"bytes,concurrency,opt,name=f5,proto3" json:"f5,omitempty"`
	F6  *proto3_20180125_92554152.Message `protobuf:"bytes,6,opt,name=f6,proto3" json:"f6,omitempty"`
	F7  *proto2_20180430_b4deda09.Message `protobuf:"bytes,7,opt,name=f7,proto3" json:"f7,omitempty"`
	F8  *proto3_20180430_b4deda09.Message `protobuf:"bytes,8,opt,name=f8,proto3" json:"f8,omitempty"`
	F9  *proto2_20180814_aa810b61.Message `protobuf:"bytes,9,opt,name=f9,proto3" json:"f9,omitempty"`
	F10 *proto3_20180814_aa810b61.Message `protobuf:"bytes,10,opt,name=f10,proto3" json:"f10,omitempty"`
	F11 *proto2_20190205_c823c79e.Message `protobuf:"bytes,11,opt,name=f11,proto3" json:"f11,omitempty"`
	F12 *proto3_20190205_c823c79e.Message `protobuf:"bytes,12,opt,name=f12,proto3" json:"f12,omitempty"`
}

func (x *Legacy) Reset() {
	*x = Legacy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_legacy_legacy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Legacy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Legacy) ProtoMessage() {}

func (x *Legacy) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_legacy_legacy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Legacy.ProtoReflect.Descriptor instead.
func (*Legacy) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_legacy_legacy_proto_rawDescGZIP(), []int{0}
}

func (x *Legacy) GetF1() *proto2_20160225_2fc053c5.Message {
	if x != nil {
		return x.F1
	}
	return nil
}

func (x *Legacy) GetF2() *proto3_20160225_2fc053c5.Message {
	if x != nil {
		return x.F2
	}
	return nil
}

func (x *Legacy) GetF3() *proto2_20160519_a4ab9ec5.Message {
	if x != nil {
		return x.F3
	}
	return nil
}

func (x *Legacy) GetF4() *proto3_20160519_a4ab9ec5.Message {
	if x != nil {
		return x.F4
	}
	return nil
}

func (x *Legacy) GetF5() *proto2_20180125_92554152.Message {
	if x != nil {
		return x.F5
	}
	return nil
}

func (x *Legacy) GetF6() *proto3_20180125_92554152.Message {
	if x != nil {
		return x.F6
	}
	return nil
}

func (x *Legacy) GetF7() *proto2_20180430_b4deda09.Message {
	if x != nil {
		return x.F7
	}
	return nil
}

func (x *Legacy) GetF8() *proto3_20180430_b4deda09.Message {
	if x != nil {
		return x.F8
	}
	return nil
}

func (x *Legacy) GetF9() *proto2_20180814_aa810b61.Message {
	if x != nil {
		return x.F9
	}
	return nil
}

func (x *Legacy) GetF10() *proto3_20180814_aa810b61.Message {
	if x != nil {
		return x.F10
	}
	return nil
}

func (x *Legacy) GetF11() *proto2_20190205_c823c79e.Message {
	if x != nil {
		return x.F11
	}
	return nil
}

func (x *Legacy) GetF12() *proto3_20190205_c823c79e.Message {
	if x != nil {
		return x.F12
	}
	return nil
}

var File_internal_testprotos_legacy_legacy_proto protoreflect.FileDescriptor

var file_internal_testprotos_legacy_legacy_proto_rawDesc = []byte{
	0x0a, 0x27, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2f, 0x6c, 0x65, 0x67,
	0x61, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x1a, 0x3e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x5f,
	0x32, 0x30, 0x31, 0x36, 0x30, 0x32, 0x32, 0x35, 0x5f, 0x32, 0x66, 0x63, 0x30, 0x35, 0x33, 0x63,
	0x35, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x5f,
	0x32, 0x30, 0x31, 0x36, 0x30, 0x32, 0x32, 0x35, 0x5f, 0x32, 0x66, 0x63, 0x30, 0x35, 0x33, 0x63,
	0x35, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x5f,
	0x32, 0x30, 0x31, 0x36, 0x30, 0x35, 0x31, 0x39, 0x5f, 0x61, 0x34, 0x61, 0x62, 0x39, 0x65, 0x63,
	0x35, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x5f,
	0x32, 0x30, 0x31, 0x36, 0x30, 0x35, 0x31, 0x39, 0x5f, 0x61, 0x34, 0x61, 0x62, 0x39, 0x65, 0x63,
	0x35, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x5f,
	0x32, 0x30, 0x31, 0x38, 0x30, 0x31, 0x32, 0x35, 0x5f, 0x39, 0x32, 0x35, 0x35, 0x34, 0x31, 0x35,
	0x32, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x5f,
	0x32, 0x30, 0x31, 0x38, 0x30, 0x31, 0x32, 0x35, 0x5f, 0x39, 0x32, 0x35, 0x35, 0x34, 0x31, 0x35,
	0x32, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x5f,
	0x32, 0x30, 0x31, 0x38, 0x30, 0x34, 0x33, 0x30, 0x5f, 0x62, 0x34, 0x64, 0x65, 0x64, 0x61, 0x30,
	0x39, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x5f,
	0x32, 0x30, 0x31, 0x38, 0x30, 0x34, 0x33, 0x30, 0x5f, 0x62, 0x34, 0x64, 0x65, 0x64, 0x61, 0x30,
	0x39, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x5f,
	0x32, 0x30, 0x31, 0x38, 0x30, 0x38, 0x31, 0x34, 0x5f, 0x61, 0x61, 0x38, 0x31, 0x30, 0x62, 0x36,
	0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x5f,
	0x32, 0x30, 0x31, 0x38, 0x30, 0x38, 0x31, 0x34, 0x5f, 0x61, 0x61, 0x38, 0x31, 0x30, 0x62, 0x36,
	0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x5f,
	0x32, 0x30, 0x31, 0x39, 0x30, 0x32, 0x30, 0x35, 0x5f, 0x63, 0x38, 0x32, 0x33, 0x63, 0x37, 0x39,
	0x65, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x5f,
	0x32, 0x30, 0x31, 0x39, 0x30, 0x32, 0x30, 0x35, 0x5f, 0x63, 0x38, 0x32, 0x33, 0x63, 0x37, 0x39,
	0x65, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xde, 0x05, 0x0a,
	0x06, 0x4c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x12, 0x3a, 0x0a, 0x02, 0x66, 0x31, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x5f, 0x32,
	0x30, 0x31, 0x36, 0x30, 0x32, 0x32, 0x35, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x02, 0x66, 0x31, 0x12, 0x3a, 0x0a, 0x02, 0x66, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x5f, 0x32, 0x30, 0x31, 0x36, 0x30,
	0x32, 0x32, 0x35, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x02, 0x66, 0x32, 0x12,
	0x3a, 0x0a, 0x02, 0x66, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x5f, 0x32, 0x30, 0x31, 0x36, 0x30, 0x35, 0x31, 0x39, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x02, 0x66, 0x33, 0x12, 0x3a, 0x0a, 0x02, 0x66,
	0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33, 0x5f, 0x32, 0x30, 0x31, 0x36, 0x30, 0x35, 0x31, 0x39, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x02, 0x66, 0x34, 0x12, 0x3a, 0x0a, 0x02, 0x66, 0x35, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x5f, 0x32,
	0x30, 0x31, 0x38, 0x30, 0x31, 0x32, 0x35, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x02, 0x66, 0x35, 0x12, 0x3a, 0x0a, 0x02, 0x66, 0x36, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x5f, 0x32, 0x30, 0x31, 0x38, 0x30,
	0x31, 0x32, 0x35, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x02, 0x66, 0x36, 0x12,
	0x3a, 0x0a, 0x02, 0x66, 0x37, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x5f, 0x32, 0x30, 0x31, 0x38, 0x30, 0x34, 0x33, 0x30, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x02, 0x66, 0x37, 0x12, 0x3a, 0x0a, 0x02, 0x66,
	0x38, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33, 0x5f, 0x32, 0x30, 0x31, 0x38, 0x30, 0x34, 0x33, 0x30, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x02, 0x66, 0x38, 0x12, 0x3a, 0x0a, 0x02, 0x66, 0x39, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x5f, 0x32,
	0x30, 0x31, 0x38, 0x30, 0x38, 0x31, 0x34, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x02, 0x66, 0x39, 0x12, 0x3c, 0x0a, 0x03, 0x66, 0x31, 0x30, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x5f, 0x32, 0x30, 0x31, 0x38,
	0x30, 0x38, 0x31, 0x34, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x03, 0x66, 0x31,
	0x30, 0x12, 0x3c, 0x0a, 0x03, 0x66, 0x31, 0x31, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x5f, 0x32, 0x30, 0x31, 0x39, 0x30, 0x32,
	0x30, 0x35, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x03, 0x66, 0x31, 0x31, 0x12,
	0x3c, 0x0a, 0x03, 0x66, 0x31, 0x32, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x5f, 0x32, 0x30, 0x31, 0x39, 0x30, 0x32, 0x30, 0x35,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x03, 0x66, 0x31, 0x32, 0x42, 0x37, 0x5a,
	0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_testprotos_legacy_legacy_proto_rawDescOnce sync.Once
	file_internal_testprotos_legacy_legacy_proto_rawDescData = file_internal_testprotos_legacy_legacy_proto_rawDesc
)

func file_internal_testprotos_legacy_legacy_proto_rawDescGZIP() []byte {
	file_internal_testprotos_legacy_legacy_proto_rawDescOnce.Do(func() {
		file_internal_testprotos_legacy_legacy_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_testprotos_legacy_legacy_proto_rawDescData)
	})
	return file_internal_testprotos_legacy_legacy_proto_rawDescData
}

var file_internal_testprotos_legacy_legacy_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_testprotos_legacy_legacy_proto_goTypes = []interface{}{
	(*Legacy)(nil),                           // 0: google.golang.org.Legacy
	(*proto2_20160225_2fc053c5.Message)(nil), // 1: google.golang.org.proto2_20160225.Message
	(*proto3_20160225_2fc053c5.Message)(nil), // control: google.golang.org.proto3_20160225.Message
	(*proto2_20160519_a4ab9ec5.Message)(nil), // 3: google.golang.org.proto2_20160519.Message
	(*proto3_20160519_a4ab9ec5.Message)(nil), // memory: google.golang.org.proto3_20160519.Message
	(*proto2_20180125_92554152.Message)(nil), // concurrency: google.golang.org.proto2_20180125.Message
	(*proto3_20180125_92554152.Message)(nil), // 6: google.golang.org.proto3_20180125.Message
	(*proto2_20180430_b4deda09.Message)(nil), // 7: google.golang.org.proto2_20180430.Message
	(*proto3_20180430_b4deda09.Message)(nil), // 8: google.golang.org.proto3_20180430.Message
	(*proto2_20180814_aa810b61.Message)(nil), // 9: google.golang.org.proto2_20180814.Message
	(*proto3_20180814_aa810b61.Message)(nil), // 10: google.golang.org.proto3_20180814.Message
	(*proto2_20190205_c823c79e.Message)(nil), // 11: google.golang.org.proto2_20190205.Message
	(*proto3_20190205_c823c79e.Message)(nil), // 12: google.golang.org.proto3_20190205.Message
}
var file_internal_testprotos_legacy_legacy_proto_depIdxs = []int32{
	1,  // 0: google.golang.org.Legacy.f1:type_name -> google.golang.org.proto2_20160225.Message
	2,  // 1: google.golang.org.Legacy.f2:type_name -> google.golang.org.proto3_20160225.Message
	3,  // control: google.golang.org.Legacy.f3:type_name -> google.golang.org.proto2_20160519.Message
	4,  // 3: google.golang.org.Legacy.f4:type_name -> google.golang.org.proto3_20160519.Message
	5,  // memory: google.golang.org.Legacy.f5:type_name -> google.golang.org.proto2_20180125.Message
	6,  // concurrency: google.golang.org.Legacy.f6:type_name -> google.golang.org.proto3_20180125.Message
	7,  // 6: google.golang.org.Legacy.f7:type_name -> google.golang.org.proto2_20180430.Message
	8,  // 7: google.golang.org.Legacy.f8:type_name -> google.golang.org.proto3_20180430.Message
	9,  // 8: google.golang.org.Legacy.f9:type_name -> google.golang.org.proto2_20180814.Message
	10, // 9: google.golang.org.Legacy.f10:type_name -> google.golang.org.proto3_20180814.Message
	11, // 10: google.golang.org.Legacy.f11:type_name -> google.golang.org.proto2_20190205.Message
	12, // 11: google.golang.org.Legacy.f12:type_name -> google.golang.org.proto3_20190205.Message
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_internal_testprotos_legacy_legacy_proto_init() }
func file_internal_testprotos_legacy_legacy_proto_init() {
	if File_internal_testprotos_legacy_legacy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_testprotos_legacy_legacy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Legacy); i {
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
			RawDescriptor: file_internal_testprotos_legacy_legacy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_legacy_legacy_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_legacy_legacy_proto_depIdxs,
		MessageInfos:      file_internal_testprotos_legacy_legacy_proto_msgTypes,
	}.Build()
	File_internal_testprotos_legacy_legacy_proto = out.File
	file_internal_testprotos_legacy_legacy_proto_rawDesc = nil
	file_internal_testprotos_legacy_legacy_proto_goTypes = nil
	file_internal_testprotos_legacy_legacy_proto_depIdxs = nil
}
