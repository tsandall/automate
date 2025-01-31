// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2beta/users.proto

package v2beta

import (
	context "context"
	fmt "fmt"
	request "github.com/chef/automate/components/automate-gateway/api/iam/v2beta/request"
	response "github.com/chef/automate/components/automate-gateway/api/iam/v2beta/response"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/api"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/iam"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("components/automate-gateway/api/iam/v2beta/users.proto", fileDescriptor_90353f4df809fe2f)
}

var fileDescriptor_90353f4df809fe2f = []byte{
	// 536 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x65, 0x4a, 0x43, 0x73, 0xfc, 0x68, 0xfb, 0x28, 0xc5, 0xb2, 0x3a, 0x79, 0x82, 0x94,
	0xd8, 0x52, 0x91, 0x18, 0x32, 0x00, 0xa2, 0x20, 0x16, 0x26, 0x50, 0x19, 0xd8, 0x2e, 0xce, 0x8b,
	0x73, 0x92, 0xed, 0xbb, 0xfa, 0xce, 0xa0, 0x0a, 0xb1, 0x64, 0xcc, 0xca, 0xdf, 0xe2, 0x8d, 0x8d,
	0x95, 0x09, 0x46, 0xc4, 0xc6, 0xcc, 0xdf, 0x80, 0xce, 0x97, 0xc6, 0x8e, 0x52, 0x47, 0x97, 0xd1,
	0xbe, 0xcf, 0x7b, 0xf7, 0xfd, 0xd8, 0xef, 0x8e, 0x3c, 0x89, 0x78, 0x2a, 0x78, 0x86, 0x99, 0x92,
	0x21, 0x2d, 0x14, 0x4f, 0xa9, 0xc2, 0x7e, 0x4c, 0x15, 0x7e, 0xa2, 0x17, 0x21, 0x15, 0x2c, 0x64,
	0x34, 0x0d, 0x3f, 0x9e, 0x0c, 0x51, 0xd1, 0xb0, 0x90, 0x98, 0xcb, 0x40, 0xe4, 0x5c, 0x71, 0x38,
	0x8a, 0x26, 0x38, 0x0e, 0x2e, 0x2b, 0x02, 0x2a, 0x58, 0xc0, 0x68, 0x1a, 0x18, 0xd2, 0x3b, 0x8a,
	0x39, 0x8f, 0x13, 0xac, 0x1a, 0xd0, 0x2c, 0xe3, 0x8a, 0x2a, 0xc6, 0xb3, 0x79, 0xad, 0xf7, 0x74,
	0x83, 0x3d, 0x73, 0x3c, 0x2f, 0x50, 0xaa, 0xe6, 0xde, 0xde, 0xb3, 0x8d, 0xea, 0xa5, 0xe0, 0x99,
	0xc4, 0xa5, 0x06, 0xcf, 0xaf, 0x6c, 0x90, 0x8b, 0x28, 0xac, 0xd6, 0xa3, 0x7e, 0x8c, 0x59, 0x5f,
	0xf0, 0x84, 0x45, 0x17, 0x2d, 0x0a, 0x9b, 0x74, 0xd0, 0x69, 0x56, 0x3a, 0x9c, 0xfc, 0xeb, 0x92,
	0xed, 0x33, 0x9d, 0x09, 0x7e, 0x38, 0x84, 0x9c, 0xe6, 0x48, 0x15, 0xea, 0x67, 0x38, 0x0e, 0xd6,
	0x7d, 0xda, 0xa0, 0x26, 0xdf, 0xe2, 0xb9, 0xf7, 0xc8, 0x1e, 0x96, 0xc2, 0x8f, 0xa6, 0xa5, 0x7b,
	0x8b, 0x10, 0x5a, 0xa8, 0xc9, 0xa0, 0xfa, 0x1e, 0xd3, 0xd2, 0xdd, 0x81, 0x4e, 0x54, 0x51, 0xb3,
	0xd2, 0xbd, 0x49, 0xba, 0x8c, 0xa6, 0x66, 0x69, 0x56, 0xba, 0x00, 0x7b, 0x8b, 0xc7, 0x81, 0x81,
	0xa6, 0xbf, 0xfe, 0x7e, 0xbd, 0x76, 0xd7, 0xdf, 0x5f, 0x19, 0x8b, 0x6a, 0x61, 0x6b, 0xe0, 0xf4,
	0xe0, 0x9b, 0x43, 0xba, 0x6f, 0x98, 0x54, 0x46, 0xae, 0xb7, 0x3e, 0xe0, 0x02, 0xd4, 0x32, 0xc7,
	0xd6, 0xac, 0x14, 0xfe, 0xfb, 0x2b, 0x5c, 0x3a, 0x70, 0x3d, 0x47, 0x3a, 0x5a, 0x35, 0xd9, 0x83,
	0x3b, 0xb5, 0x49, 0xc2, 0xa4, 0x32, 0x1e, 0xb0, 0xea, 0x01, 0xdf, 0x1d, 0x72, 0xe3, 0x35, 0x56,
	0x1b, 0xc1, 0x83, 0xf5, 0x81, 0xe6, 0x98, 0x8e, 0xfe, 0xd0, 0x92, 0x9c, 0xff, 0x84, 0x7d, 0xb2,
	0x5b, 0x07, 0x1f, 0x7c, 0x66, 0xa3, 0x2f, 0xd3, 0xd2, 0xdd, 0x86, 0xad, 0x18, 0x95, 0x4e, 0x4b,
	0x1a, 0x69, 0xf5, 0xe2, 0xac, 0x74, 0x77, 0xe1, 0x76, 0xfd, 0x2e, 0x46, 0x23, 0xe0, 0xc2, 0xe1,
	0x8a, 0x40, 0xa8, 0x0b, 0xe0, 0xa7, 0x43, 0xc8, 0x4b, 0x4c, 0xd0, 0x6e, 0xa6, 0x6a, 0xd2, 0x62,
	0xa6, 0x9a, 0xb0, 0x14, 0x3e, 0x6b, 0xd3, 0xd9, 0x81, 0xce, 0xa8, 0x42, 0x5b, 0x8c, 0x96, 0xa6,
	0xcb, 0x90, 0x46, 0xaa, 0xd7, 0x26, 0xf5, 0xdb, 0x21, 0xe4, 0x4c, 0x8c, 0x2c, 0x0f, 0x4a, 0x4d,
	0x5a, 0x48, 0x35, 0x61, 0x29, 0xfc, 0x7c, 0x8d, 0x54, 0x51, 0xa1, 0x36, 0x52, 0x86, 0x34, 0x52,
	0x5e, 0x8b, 0x54, 0x7d, 0x6e, 0xfe, 0x2c, 0xec, 0xde, 0x61, 0x32, 0xb6, 0xb3, 0xd3, 0xa4, 0xb5,
	0x9d, 0x81, 0xa5, 0xf0, 0x95, 0x39, 0x3a, 0xad, 0x62, 0x07, 0x04, 0x16, 0x12, 0xba, 0xe8, 0x52,
	0xee, 0x10, 0x0e, 0x96, 0xdf, 0x37, 0x04, 0xef, 0x7b, 0xf7, 0x9a, 0x82, 0x12, 0x93, 0xf1, 0xb2,
	0xdf, 0x8b, 0x57, 0x1f, 0x4e, 0x63, 0xa6, 0x26, 0xc5, 0x30, 0x88, 0x78, 0x1a, 0xea, 0xbc, 0x8b,
	0x9b, 0x33, 0xb4, 0xbf, 0xd0, 0x87, 0x9d, 0xea, 0xfa, 0x7c, 0xfc, 0x3f, 0x00, 0x00, 0xff, 0xff,
	0xad, 0x1d, 0xa0, 0x40, 0xb9, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersClient interface {
	CreateUser(ctx context.Context, in *request.CreateUserReq, opts ...grpc.CallOption) (*response.CreateUserResp, error)
	ListUsers(ctx context.Context, in *request.ListUsersReq, opts ...grpc.CallOption) (*response.ListUsersResp, error)
	GetUser(ctx context.Context, in *request.GetUserReq, opts ...grpc.CallOption) (*response.GetUserResp, error)
	DeleteUser(ctx context.Context, in *request.DeleteUserReq, opts ...grpc.CallOption) (*response.DeleteUserResp, error)
	UpdateUser(ctx context.Context, in *request.UpdateUserReq, opts ...grpc.CallOption) (*response.UpdateUserResp, error)
	UpdateSelf(ctx context.Context, in *request.UpdateSelfReq, opts ...grpc.CallOption) (*response.UpdateSelfResp, error)
}

type usersClient struct {
	cc *grpc.ClientConn
}

func NewUsersClient(cc *grpc.ClientConn) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) CreateUser(ctx context.Context, in *request.CreateUserReq, opts ...grpc.CallOption) (*response.CreateUserResp, error) {
	out := new(response.CreateUserResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2beta.Users/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) ListUsers(ctx context.Context, in *request.ListUsersReq, opts ...grpc.CallOption) (*response.ListUsersResp, error) {
	out := new(response.ListUsersResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2beta.Users/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetUser(ctx context.Context, in *request.GetUserReq, opts ...grpc.CallOption) (*response.GetUserResp, error) {
	out := new(response.GetUserResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2beta.Users/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) DeleteUser(ctx context.Context, in *request.DeleteUserReq, opts ...grpc.CallOption) (*response.DeleteUserResp, error) {
	out := new(response.DeleteUserResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2beta.Users/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) UpdateUser(ctx context.Context, in *request.UpdateUserReq, opts ...grpc.CallOption) (*response.UpdateUserResp, error) {
	out := new(response.UpdateUserResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2beta.Users/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) UpdateSelf(ctx context.Context, in *request.UpdateSelfReq, opts ...grpc.CallOption) (*response.UpdateSelfResp, error) {
	out := new(response.UpdateSelfResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2beta.Users/UpdateSelf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
type UsersServer interface {
	CreateUser(context.Context, *request.CreateUserReq) (*response.CreateUserResp, error)
	ListUsers(context.Context, *request.ListUsersReq) (*response.ListUsersResp, error)
	GetUser(context.Context, *request.GetUserReq) (*response.GetUserResp, error)
	DeleteUser(context.Context, *request.DeleteUserReq) (*response.DeleteUserResp, error)
	UpdateUser(context.Context, *request.UpdateUserReq) (*response.UpdateUserResp, error)
	UpdateSelf(context.Context, *request.UpdateSelfReq) (*response.UpdateSelfResp, error)
}

// UnimplementedUsersServer can be embedded to have forward compatible implementations.
type UnimplementedUsersServer struct {
}

func (*UnimplementedUsersServer) CreateUser(ctx context.Context, req *request.CreateUserReq) (*response.CreateUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUsersServer) ListUsers(ctx context.Context, req *request.ListUsersReq) (*response.ListUsersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (*UnimplementedUsersServer) GetUser(ctx context.Context, req *request.GetUserReq) (*response.GetUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedUsersServer) DeleteUser(ctx context.Context, req *request.DeleteUserReq) (*response.DeleteUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (*UnimplementedUsersServer) UpdateUser(ctx context.Context, req *request.UpdateUserReq) (*response.UpdateUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (*UnimplementedUsersServer) UpdateSelf(ctx context.Context, req *request.UpdateSelfReq) (*response.UpdateSelfResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSelf not implemented")
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.CreateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2beta.Users/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).CreateUser(ctx, req.(*request.CreateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ListUsersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2beta.Users/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).ListUsers(ctx, req.(*request.ListUsersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2beta.Users/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUser(ctx, req.(*request.GetUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.DeleteUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2beta.Users/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).DeleteUser(ctx, req.(*request.DeleteUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.UpdateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2beta.Users/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).UpdateUser(ctx, req.(*request.UpdateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_UpdateSelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.UpdateSelfReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).UpdateSelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2beta.Users/UpdateSelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).UpdateSelf(ctx, req.(*request.UpdateSelfReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.iam.v2beta.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _Users_CreateUser_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _Users_ListUsers_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Users_GetUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Users_DeleteUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Users_UpdateUser_Handler,
		},
		{
			MethodName: "UpdateSelf",
			Handler:    _Users_UpdateSelf_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "components/automate-gateway/api/iam/v2beta/users.proto",
}
