// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

func init() { proto.RegisterFile("server.proto", fileDescriptor_ad098daeda4239f7) }

var fileDescriptor_ad098daeda4239f7 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x0f, 0x6a, 0x8b, 0x63, 0x8b, 0x90, 0x0a, 0x42, 0xc0, 0x8a, 0xe8, 0xb9, 0x07, 0x3d,
	0x78, 0x13, 0x2c, 0x15, 0x15, 0x4a, 0x0f, 0x0b, 0xad, 0x20, 0x5e, 0x52, 0x77, 0x64, 0x83, 0x31,
	0x89, 0x49, 0xb6, 0xd0, 0xf7, 0xf4, 0x81, 0x64, 0xb3, 0x09, 0xda, 0xdd, 0x28, 0x3d, 0xce, 0x9f,
	0xff, 0xfb, 0x98, 0x59, 0x16, 0x7a, 0x16, 0xcd, 0x0a, 0xcd, 0x48, 0x1b, 0xe5, 0x14, 0xe9, 0x4b,
	0xe5, 0xf8, 0x1b, 0x7f, 0x65, 0x8e, 0x2b, 0x69, 0xe9, 0x60, 0x63, 0xac, 0x3b, 0xf4, 0xc0, 0x94,
	0x02, 0xe3, 0xd0, 0x2b, 0x90, 0x09, 0x57, 0xd4, 0xd3, 0xe5, 0xd7, 0x2e, 0xf4, 0x67, 0xbf, 0x11,
	0x72, 0x0d, 0x1d, 0x1f, 0xac, 0xc9, 0xd1, 0x68, 0x53, 0x76, 0xb7, 0x42, 0xe9, 0xe8, 0x71, 0x23,
	0xcd, 0xd0, 0x6a, 0x25, 0x2d, 0x92, 0x1b, 0xe8, 0xde, 0xe6, 0x79, 0x56, 0x0a, 0x24, 0x83, 0x66,
	0xa7, 0x14, 0x48, 0x87, 0x89, 0xb0, 0x02, 0x22, 0x3f, 0x03, 0x98, 0xa0, 0x40, 0x87, 0x5e, 0x71,
	0x92, 0x68, 0x3f, 0xe6, 0x28, 0xab, 0x0c, 0x0d, 0x3d, 0x4b, 0x3c, 0x07, 0x3a, 0xfa, 0x26, 0x00,
	0x73, 0x9d, 0xb3, 0xe0, 0x4b, 0xae, 0x94, 0xb2, 0x04, 0x26, 0x5a, 0x1e, 0xa0, 0x7b, 0x8f, 0x6e,
	0x9b, 0x95, 0x52, 0xf7, 0x55, 0x68, 0x34, 0x8d, 0x61, 0x7f, 0xca, 0xad, 0x57, 0xd9, 0xf6, 0xb7,
	0xfd, 0xd0, 0x6e, 0x4d, 0x4f, 0x13, 0x0a, 0xcf, 0x44, 0xc7, 0x0b, 0x1c, 0x2e, 0x98, 0xe0, 0xd5,
	0x86, 0x4f, 0xb8, 0x2c, 0x94, 0x7a, 0x27, 0xe7, 0x0d, 0x66, 0x9e, 0x4d, 0x43, 0x85, 0x2b, 0x99,
	0xe1, 0x67, 0x89, 0xd6, 0xd1, 0x8b, 0xff, 0x4b, 0x3f, 0xb7, 0x2e, 0xd0, 0x58, 0xae, 0x64, 0xeb,
	0xd6, 0x90, 0x47, 0xdf, 0xf0, 0xaf, 0xe7, 0xda, 0x34, 0xde, 0x7b, 0xde, 0x61, 0x9a, 0x2f, 0x3b,
	0xfe, 0x27, 0xbb, 0xfa, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xe5, 0xe0, 0xcf, 0x4c, 0xb3, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NotificationsClient is the client API for Notifications service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotificationsClient interface {
	// Publish a notification
	Notify(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Response, error)
	// Manage notification alerting rules
	AddRule(ctx context.Context, in *Rule, opts ...grpc.CallOption) (*RuleAddResponse, error)
	DeleteRule(ctx context.Context, in *RuleIdentifier, opts ...grpc.CallOption) (*RuleDeleteResponse, error)
	UpdateRule(ctx context.Context, in *Rule, opts ...grpc.CallOption) (*RuleUpdateResponse, error)
	GetRule(ctx context.Context, in *RuleIdentifier, opts ...grpc.CallOption) (*RuleGetResponse, error)
	ListRules(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RuleListResponse, error)
	ValidateWebhook(ctx context.Context, in *URLValidationRequest, opts ...grpc.CallOption) (*URLValidationResponse, error)
	// Health checks and metadata
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
}

type notificationsClient struct {
	cc *grpc.ClientConn
}

func NewNotificationsClient(cc *grpc.ClientConn) NotificationsClient {
	return &notificationsClient{cc}
}

func (c *notificationsClient) Notify(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/notifications.Notifications/Notify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) AddRule(ctx context.Context, in *Rule, opts ...grpc.CallOption) (*RuleAddResponse, error) {
	out := new(RuleAddResponse)
	err := c.cc.Invoke(ctx, "/notifications.Notifications/AddRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) DeleteRule(ctx context.Context, in *RuleIdentifier, opts ...grpc.CallOption) (*RuleDeleteResponse, error) {
	out := new(RuleDeleteResponse)
	err := c.cc.Invoke(ctx, "/notifications.Notifications/DeleteRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) UpdateRule(ctx context.Context, in *Rule, opts ...grpc.CallOption) (*RuleUpdateResponse, error) {
	out := new(RuleUpdateResponse)
	err := c.cc.Invoke(ctx, "/notifications.Notifications/UpdateRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) GetRule(ctx context.Context, in *RuleIdentifier, opts ...grpc.CallOption) (*RuleGetResponse, error) {
	out := new(RuleGetResponse)
	err := c.cc.Invoke(ctx, "/notifications.Notifications/GetRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) ListRules(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RuleListResponse, error) {
	out := new(RuleListResponse)
	err := c.cc.Invoke(ctx, "/notifications.Notifications/ListRules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) ValidateWebhook(ctx context.Context, in *URLValidationRequest, opts ...grpc.CallOption) (*URLValidationResponse, error) {
	out := new(URLValidationResponse)
	err := c.cc.Invoke(ctx, "/notifications.Notifications/ValidateWebhook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/notifications.Notifications/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationsServer is the server API for Notifications service.
type NotificationsServer interface {
	// Publish a notification
	Notify(context.Context, *Event) (*Response, error)
	// Manage notification alerting rules
	AddRule(context.Context, *Rule) (*RuleAddResponse, error)
	DeleteRule(context.Context, *RuleIdentifier) (*RuleDeleteResponse, error)
	UpdateRule(context.Context, *Rule) (*RuleUpdateResponse, error)
	GetRule(context.Context, *RuleIdentifier) (*RuleGetResponse, error)
	ListRules(context.Context, *Empty) (*RuleListResponse, error)
	ValidateWebhook(context.Context, *URLValidationRequest) (*URLValidationResponse, error)
	// Health checks and metadata
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
}

func RegisterNotificationsServer(s *grpc.Server, srv NotificationsServer) {
	s.RegisterService(&_Notifications_serviceDesc, srv)
}

func _Notifications_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.Notifications/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).Notify(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_AddRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Rule)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).AddRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.Notifications/AddRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).AddRule(ctx, req.(*Rule))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_DeleteRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RuleIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).DeleteRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.Notifications/DeleteRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).DeleteRule(ctx, req.(*RuleIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_UpdateRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Rule)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).UpdateRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.Notifications/UpdateRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).UpdateRule(ctx, req.(*Rule))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_GetRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RuleIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).GetRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.Notifications/GetRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).GetRule(ctx, req.(*RuleIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_ListRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).ListRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.Notifications/ListRules",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).ListRules(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_ValidateWebhook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(URLValidationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).ValidateWebhook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.Notifications/ValidateWebhook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).ValidateWebhook(ctx, req.(*URLValidationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.Notifications/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Notifications_serviceDesc = grpc.ServiceDesc{
	ServiceName: "notifications.Notifications",
	HandlerType: (*NotificationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Notify",
			Handler:    _Notifications_Notify_Handler,
		},
		{
			MethodName: "AddRule",
			Handler:    _Notifications_AddRule_Handler,
		},
		{
			MethodName: "DeleteRule",
			Handler:    _Notifications_DeleteRule_Handler,
		},
		{
			MethodName: "UpdateRule",
			Handler:    _Notifications_UpdateRule_Handler,
		},
		{
			MethodName: "GetRule",
			Handler:    _Notifications_GetRule_Handler,
		},
		{
			MethodName: "ListRules",
			Handler:    _Notifications_ListRules_Handler,
		},
		{
			MethodName: "ValidateWebhook",
			Handler:    _Notifications_ValidateWebhook_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _Notifications_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}
