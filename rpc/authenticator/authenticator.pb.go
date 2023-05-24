// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.9
// source: authenticator/authenticator.proto

package authenticator

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetContextExecCredentialsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContextName string `protobuf:"bytes,1,opt,name=context_name,json=contextName,proto3" json:"context_name,omitempty"`
}

func (x *GetContextExecCredentialsRequest) Reset() {
	*x = GetContextExecCredentialsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authenticator_authenticator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetContextExecCredentialsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContextExecCredentialsRequest) ProtoMessage() {}

func (x *GetContextExecCredentialsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authenticator_authenticator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContextExecCredentialsRequest.ProtoReflect.Descriptor instead.
func (*GetContextExecCredentialsRequest) Descriptor() ([]byte, []int) {
	return file_authenticator_authenticator_proto_rawDescGZIP(), []int{0}
}

func (x *GetContextExecCredentialsRequest) GetContextName() string {
	if x != nil {
		return x.ContextName
	}
	return ""
}

type GetContextExecCredentialsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawCredentials []byte `protobuf:"bytes,1,opt,name=raw_credentials,json=rawCredentials,proto3" json:"raw_credentials,omitempty"`
}

func (x *GetContextExecCredentialsResponse) Reset() {
	*x = GetContextExecCredentialsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authenticator_authenticator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetContextExecCredentialsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContextExecCredentialsResponse) ProtoMessage() {}

func (x *GetContextExecCredentialsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authenticator_authenticator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContextExecCredentialsResponse.ProtoReflect.Descriptor instead.
func (*GetContextExecCredentialsResponse) Descriptor() ([]byte, []int) {
	return file_authenticator_authenticator_proto_rawDescGZIP(), []int{1}
}

func (x *GetContextExecCredentialsResponse) GetRawCredentials() []byte {
	if x != nil {
		return x.RawCredentials
	}
	return nil
}

var File_authenticator_authenticator_proto protoreflect.FileDescriptor

var file_authenticator_authenticator_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63,
	0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x22,
	0x45, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x45, 0x78, 0x65,
	0x63, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4c, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x45, 0x78, 0x65, 0x63, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x72,
	0x61, 0x77, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x72, 0x61, 0x77, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x32, 0xaa, 0x01, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x98, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x45, 0x78, 0x65, 0x63, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x12, 0x3c, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65,
	0x6e, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x45, 0x78, 0x65, 0x63,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63,
	0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x45, 0x78, 0x65, 0x63, 0x43, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x69, 0x6f, 0x2f, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f,
	0x76, 0x32, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_authenticator_authenticator_proto_rawDescOnce sync.Once
	file_authenticator_authenticator_proto_rawDescData = file_authenticator_authenticator_proto_rawDesc
)

func file_authenticator_authenticator_proto_rawDescGZIP() []byte {
	file_authenticator_authenticator_proto_rawDescOnce.Do(func() {
		file_authenticator_authenticator_proto_rawDescData = protoimpl.X.CompressGZIP(file_authenticator_authenticator_proto_rawDescData)
	})
	return file_authenticator_authenticator_proto_rawDescData
}

var (
	file_authenticator_authenticator_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
	file_authenticator_authenticator_proto_goTypes  = []interface{}{
		(*GetContextExecCredentialsRequest)(nil),  // 0: telepresence.authenticator.GetContextExecCredentialsRequest
		(*GetContextExecCredentialsResponse)(nil), // 1: telepresence.authenticator.GetContextExecCredentialsResponse
	}
)
var file_authenticator_authenticator_proto_depIdxs = []int32{
	0, // 0: telepresence.authenticator.Authenticator.GetContextExecCredentials:input_type -> telepresence.authenticator.GetContextExecCredentialsRequest
	1, // 1: telepresence.authenticator.Authenticator.GetContextExecCredentials:output_type -> telepresence.authenticator.GetContextExecCredentialsResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_authenticator_authenticator_proto_init() }
func file_authenticator_authenticator_proto_init() {
	if File_authenticator_authenticator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authenticator_authenticator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetContextExecCredentialsRequest); i {
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
		file_authenticator_authenticator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetContextExecCredentialsResponse); i {
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
			RawDescriptor: file_authenticator_authenticator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authenticator_authenticator_proto_goTypes,
		DependencyIndexes: file_authenticator_authenticator_proto_depIdxs,
		MessageInfos:      file_authenticator_authenticator_proto_msgTypes,
	}.Build()
	File_authenticator_authenticator_proto = out.File
	file_authenticator_authenticator_proto_rawDesc = nil
	file_authenticator_authenticator_proto_goTypes = nil
	file_authenticator_authenticator_proto_depIdxs = nil
}
