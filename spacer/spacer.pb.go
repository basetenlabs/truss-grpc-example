// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: spacer.proto

package yapper

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type YapRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *YapRequest) Reset() {
	*x = YapRequest{}
	mi := &file_spacer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *YapRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YapRequest) ProtoMessage() {}

func (x *YapRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spacer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YapRequest.ProtoReflect.Descriptor instead.
func (*YapRequest) Descriptor() ([]byte, []int) {
	return file_spacer_proto_rawDescGZIP(), []int{0}
}

func (x *YapRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type YapReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Reply         string                 `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *YapReply) Reset() {
	*x = YapReply{}
	mi := &file_spacer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *YapReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YapReply) ProtoMessage() {}

func (x *YapReply) ProtoReflect() protoreflect.Message {
	mi := &file_spacer_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YapReply.ProtoReflect.Descriptor instead.
func (*YapReply) Descriptor() ([]byte, []int) {
	return file_spacer_proto_rawDescGZIP(), []int{1}
}

func (x *YapReply) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

var File_spacer_proto protoreflect.FileDescriptor

const file_spacer_proto_rawDesc = "" +
	"\n" +
	"\fspacer.proto\x12\x06yapper\"&\n" +
	"\n" +
	"YapRequest\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\" \n" +
	"\bYapReply\x12\x14\n" +
	"\x05reply\x18\x01 \x01(\tR\x05reply2>\n" +
	"\rYapperService\x12-\n" +
	"\x03Yap\x12\x12.yapper.YapRequest\x1a\x10.yapper.YapReply0\x01B\rZ\vasdf/yapperb\x06proto3"

var (
	file_spacer_proto_rawDescOnce sync.Once
	file_spacer_proto_rawDescData []byte
)

func file_spacer_proto_rawDescGZIP() []byte {
	file_spacer_proto_rawDescOnce.Do(func() {
		file_spacer_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_spacer_proto_rawDesc), len(file_spacer_proto_rawDesc)))
	})
	return file_spacer_proto_rawDescData
}

var file_spacer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_spacer_proto_goTypes = []any{
	(*YapRequest)(nil), // 0: yapper.YapRequest
	(*YapReply)(nil),   // 1: yapper.YapReply
}
var file_spacer_proto_depIdxs = []int32{
	0, // 0: yapper.YapperService.Yap:input_type -> yapper.YapRequest
	1, // 1: yapper.YapperService.Yap:output_type -> yapper.YapReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_spacer_proto_init() }
func file_spacer_proto_init() {
	if File_spacer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spacer_proto_rawDesc), len(file_spacer_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spacer_proto_goTypes,
		DependencyIndexes: file_spacer_proto_depIdxs,
		MessageInfos:      file_spacer_proto_msgTypes,
	}.Build()
	File_spacer_proto = out.File
	file_spacer_proto_goTypes = nil
	file_spacer_proto_depIdxs = nil
}
