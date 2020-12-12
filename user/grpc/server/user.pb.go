// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: server/user.proto

package server

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_server_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exits  bool `protobuf:"varint,1,opt,name=exits,proto3" json:"exits,omitempty"`
	Active bool `protobuf:"varint,2,opt,name=active,proto3" json:"active,omitempty"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_server_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserResponse) GetExits() bool {
	if x != nil {
		return x.Exits
	}
	return false
}

func (x *UserResponse) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

type RelationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromUsername string `protobuf:"bytes,1,opt,name=fromUsername,proto3" json:"fromUsername,omitempty"`
	ToUsername   string `protobuf:"bytes,2,opt,name=toUsername,proto3" json:"toUsername,omitempty"`
}

func (x *RelationRequest) Reset() {
	*x = RelationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationRequest) ProtoMessage() {}

func (x *RelationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationRequest.ProtoReflect.Descriptor instead.
func (*RelationRequest) Descriptor() ([]byte, []int) {
	return file_server_user_proto_rawDescGZIP(), []int{2}
}

func (x *RelationRequest) GetFromUsername() string {
	if x != nil {
		return x.FromUsername
	}
	return ""
}

func (x *RelationRequest) GetToUsername() string {
	if x != nil {
		return x.ToUsername
	}
	return ""
}

type RelationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exits bool `protobuf:"varint,1,opt,name=exits,proto3" json:"exits,omitempty"`
}

func (x *RelationResponse) Reset() {
	*x = RelationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationResponse) ProtoMessage() {}

func (x *RelationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationResponse.ProtoReflect.Descriptor instead.
func (*RelationResponse) Descriptor() ([]byte, []int) {
	return file_server_user_proto_rawDescGZIP(), []int{3}
}

func (x *RelationResponse) GetExits() bool {
	if x != nil {
		return x.Exits
	}
	return false
}

type AvatarName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *AvatarName) Reset() {
	*x = AvatarName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarName) ProtoMessage() {}

func (x *AvatarName) ProtoReflect() protoreflect.Message {
	mi := &file_server_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarName.ProtoReflect.Descriptor instead.
func (*AvatarName) Descriptor() ([]byte, []int) {
	return file_server_user_proto_rawDescGZIP(), []int{4}
}

func (x *AvatarName) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AvatarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sucess bool `protobuf:"varint,1,opt,name=sucess,proto3" json:"sucess,omitempty"`
}

func (x *AvatarResponse) Reset() {
	*x = AvatarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarResponse) ProtoMessage() {}

func (x *AvatarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarResponse.ProtoReflect.Descriptor instead.
func (*AvatarResponse) Descriptor() ([]byte, []int) {
	return file_server_user_proto_rawDescGZIP(), []int{5}
}

func (x *AvatarResponse) GetSucess() bool {
	if x != nil {
		return x.Sucess
	}
	return false
}

var File_server_user_proto protoreflect.FileDescriptor

var file_server_user_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x1d, 0x0a, 0x0b, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x3c, 0x0a, 0x0c, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78,
	0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x78, 0x69, 0x74, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x55, 0x0a, 0x0f, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x66,
	0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x28, 0x0a, 0x10, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x65, 0x78, 0x69, 0x74, 0x73, 0x22, 0x20, 0x0a, 0x0a, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x28, 0x0a, 0x0e, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x75, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73,
	0x75, 0x63, 0x65, 0x73, 0x73, 0x32, 0xc5, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a,
	0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3a, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x12, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_server_user_proto_rawDescOnce sync.Once
	file_server_user_proto_rawDescData = file_server_user_proto_rawDesc
)

func file_server_user_proto_rawDescGZIP() []byte {
	file_server_user_proto_rawDescOnce.Do(func() {
		file_server_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_user_proto_rawDescData)
	})
	return file_server_user_proto_rawDescData
}

var file_server_user_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_server_user_proto_goTypes = []interface{}{
	(*UserRequest)(nil),      // 0: server.UserRequest
	(*UserResponse)(nil),     // 1: server.UserResponse
	(*RelationRequest)(nil),  // 2: server.RelationRequest
	(*RelationResponse)(nil), // 3: server.RelationResponse
	(*AvatarName)(nil),       // 4: server.AvatarName
	(*AvatarResponse)(nil),   // 5: server.AvatarResponse
}
var file_server_user_proto_depIdxs = []int32{
	0, // 0: server.UserService.CheckUser:input_type -> server.UserRequest
	2, // 1: server.UserService.CheckRelation:input_type -> server.RelationRequest
	4, // 2: server.UserService.ChangeAvatar:input_type -> server.AvatarName
	1, // 3: server.UserService.CheckUser:output_type -> server.UserResponse
	3, // 4: server.UserService.CheckRelation:output_type -> server.RelationResponse
	5, // 5: server.UserService.ChangeAvatar:output_type -> server.AvatarResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_server_user_proto_init() }
func file_server_user_proto_init() {
	if File_server_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRequest); i {
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
		file_server_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResponse); i {
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
		file_server_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationRequest); i {
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
		file_server_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationResponse); i {
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
		file_server_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarName); i {
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
		file_server_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarResponse); i {
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
			RawDescriptor: file_server_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_user_proto_goTypes,
		DependencyIndexes: file_server_user_proto_depIdxs,
		MessageInfos:      file_server_user_proto_msgTypes,
	}.Build()
	File_server_user_proto = out.File
	file_server_user_proto_rawDesc = nil
	file_server_user_proto_goTypes = nil
	file_server_user_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	CheckUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	CheckRelation(ctx context.Context, in *RelationRequest, opts ...grpc.CallOption) (*RelationResponse, error)
	ChangeAvatar(ctx context.Context, in *AvatarName, opts ...grpc.CallOption) (*AvatarResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CheckUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/server.UserService/CheckUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CheckRelation(ctx context.Context, in *RelationRequest, opts ...grpc.CallOption) (*RelationResponse, error) {
	out := new(RelationResponse)
	err := c.cc.Invoke(ctx, "/server.UserService/CheckRelation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangeAvatar(ctx context.Context, in *AvatarName, opts ...grpc.CallOption) (*AvatarResponse, error) {
	out := new(AvatarResponse)
	err := c.cc.Invoke(ctx, "/server.UserService/ChangeAvatar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	CheckUser(context.Context, *UserRequest) (*UserResponse, error)
	CheckRelation(context.Context, *RelationRequest) (*RelationResponse, error)
	ChangeAvatar(context.Context, *AvatarName) (*AvatarResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) CheckUser(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUser not implemented")
}
func (*UnimplementedUserServiceServer) CheckRelation(context.Context, *RelationRequest) (*RelationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckRelation not implemented")
}
func (*UnimplementedUserServiceServer) ChangeAvatar(context.Context, *AvatarName) (*AvatarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeAvatar not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_CheckUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CheckUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.UserService/CheckUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CheckUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CheckRelation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CheckRelation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.UserService/CheckRelation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CheckRelation(ctx, req.(*RelationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangeAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AvatarName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangeAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.UserService/ChangeAvatar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangeAvatar(ctx, req.(*AvatarName))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckUser",
			Handler:    _UserService_CheckUser_Handler,
		},
		{
			MethodName: "CheckRelation",
			Handler:    _UserService_CheckRelation_Handler,
		},
		{
			MethodName: "ChangeAvatar",
			Handler:    _UserService_ChangeAvatar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/user.proto",
}
