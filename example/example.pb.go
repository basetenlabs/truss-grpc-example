// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: example.proto

package example

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

type ExampleRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExampleRequest) Reset() {
	*x = ExampleRequest{}
	mi := &file_example_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExampleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleRequest) ProtoMessage() {}

func (x *ExampleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleRequest.ProtoReflect.Descriptor instead.
func (*ExampleRequest) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{0}
}

func (x *ExampleRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ExampleReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Reply         string                 `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExampleReply) Reset() {
	*x = ExampleReply{}
	mi := &file_example_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExampleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleReply) ProtoMessage() {}

func (x *ExampleReply) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleReply.ProtoReflect.Descriptor instead.
func (*ExampleReply) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{1}
}

func (x *ExampleReply) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

var File_example_proto protoreflect.FileDescriptor

const file_example_proto_rawDesc = "" +
	"\n" +
	"\rexample.proto\x12\aexample\"*\n" +
	"\x0eExampleRequest\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"$\n" +
	"\fExampleReply\x12\x14\n" +
	"\x05reply\x18\x01 \x01(\tR\x05reply2M\n" +
	"\x0eExampleService\x12;\n" +
	"\aExample\x12\x17.example.ExampleRequest\x1a\x15.example.ExampleReply0\x01B\vZ\t./exampleb\x06proto3"

var (
	file_example_proto_rawDescOnce sync.Once
	file_example_proto_rawDescData []byte
)

func file_example_proto_rawDescGZIP() []byte {
	file_example_proto_rawDescOnce.Do(func() {
		file_example_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_example_proto_rawDesc), len(file_example_proto_rawDesc)))
	})
	return file_example_proto_rawDescData
}

var file_example_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_example_proto_goTypes = []any{
	(*ExampleRequest)(nil), // 0: example.ExampleRequest
	(*ExampleReply)(nil),   // 1: example.ExampleReply
}
var file_example_proto_depIdxs = []int32{
	0, // 0: example.ExampleService.Example:input_type -> example.ExampleRequest
	1, // 1: example.ExampleService.Example:output_type -> example.ExampleReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_example_proto_init() }
func file_example_proto_init() {
	if File_example_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_example_proto_rawDesc), len(file_example_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_proto_goTypes,
		DependencyIndexes: file_example_proto_depIdxs,
		MessageInfos:      file_example_proto_msgTypes,
	}.Build()
	File_example_proto = out.File
	file_example_proto_goTypes = nil
	file_example_proto_depIdxs = nil
}
