// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: rpc_handshake.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HandshakeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKey string `protobuf:"bytes,1,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
}

func (x *HandshakeRequest) Reset() {
	*x = HandshakeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_handshake_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandshakeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandshakeRequest) ProtoMessage() {}

func (x *HandshakeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_handshake_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandshakeRequest.ProtoReflect.Descriptor instead.
func (*HandshakeRequest) Descriptor() ([]byte, []int) {
	return file_rpc_handshake_proto_rawDescGZIP(), []int{0}
}

func (x *HandshakeRequest) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

type HandshakeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthKey   string `protobuf:"bytes,1,opt,name=authKey,proto3" json:"authKey,omitempty"`
	PublicKey string `protobuf:"bytes,2,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
}

func (x *HandshakeResponse) Reset() {
	*x = HandshakeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_handshake_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandshakeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandshakeResponse) ProtoMessage() {}

func (x *HandshakeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_handshake_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandshakeResponse.ProtoReflect.Descriptor instead.
func (*HandshakeResponse) Descriptor() ([]byte, []int) {
	return file_rpc_handshake_proto_rawDescGZIP(), []int{1}
}

func (x *HandshakeResponse) GetAuthKey() string {
	if x != nil {
		return x.AuthKey
	}
	return ""
}

func (x *HandshakeResponse) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

var File_rpc_handshake_proto protoreflect.FileDescriptor

var file_rpc_handshake_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x70, 0x63, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x30, 0x0a, 0x10, 0x48, 0x61, 0x6e,
	0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x4b, 0x0a, 0x11, 0x48,
	0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x42, 0x0c, 0x5a, 0x0a, 0x6d, 0x73, 0x5f, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_handshake_proto_rawDescOnce sync.Once
	file_rpc_handshake_proto_rawDescData = file_rpc_handshake_proto_rawDesc
)

func file_rpc_handshake_proto_rawDescGZIP() []byte {
	file_rpc_handshake_proto_rawDescOnce.Do(func() {
		file_rpc_handshake_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_handshake_proto_rawDescData)
	})
	return file_rpc_handshake_proto_rawDescData
}

var file_rpc_handshake_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_handshake_proto_goTypes = []interface{}{
	(*HandshakeRequest)(nil),  // 0: pb.HandshakeRequest
	(*HandshakeResponse)(nil), // 1: pb.HandshakeResponse
}
var file_rpc_handshake_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_handshake_proto_init() }
func file_rpc_handshake_proto_init() {
	if File_rpc_handshake_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_handshake_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandshakeRequest); i {
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
		file_rpc_handshake_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandshakeResponse); i {
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
			RawDescriptor: file_rpc_handshake_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_handshake_proto_goTypes,
		DependencyIndexes: file_rpc_handshake_proto_depIdxs,
		MessageInfos:      file_rpc_handshake_proto_msgTypes,
	}.Build()
	File_rpc_handshake_proto = out.File
	file_rpc_handshake_proto_rawDesc = nil
	file_rpc_handshake_proto_goTypes = nil
	file_rpc_handshake_proto_depIdxs = nil
}