// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2beta/response/users.proto

package response

import (
	fmt "fmt"
	common "github.com/chef/automate/components/automate-gateway/api/iam/v2beta/common"
	proto "github.com/golang/protobuf/proto"
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

type CreateUserResp struct {
	User                 *common.User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateUserResp) Reset()         { *m = CreateUserResp{} }
func (m *CreateUserResp) String() string { return proto.CompactTextString(m) }
func (*CreateUserResp) ProtoMessage()    {}
func (*CreateUserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_525b6514ab436d23, []int{0}
}

func (m *CreateUserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResp.Unmarshal(m, b)
}
func (m *CreateUserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResp.Marshal(b, m, deterministic)
}
func (m *CreateUserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResp.Merge(m, src)
}
func (m *CreateUserResp) XXX_Size() int {
	return xxx_messageInfo_CreateUserResp.Size(m)
}
func (m *CreateUserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResp proto.InternalMessageInfo

func (m *CreateUserResp) GetUser() *common.User {
	if m != nil {
		return m.User
	}
	return nil
}

type ListUsersResp struct {
	Users                []*common.User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListUsersResp) Reset()         { *m = ListUsersResp{} }
func (m *ListUsersResp) String() string { return proto.CompactTextString(m) }
func (*ListUsersResp) ProtoMessage()    {}
func (*ListUsersResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_525b6514ab436d23, []int{1}
}

func (m *ListUsersResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersResp.Unmarshal(m, b)
}
func (m *ListUsersResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersResp.Marshal(b, m, deterministic)
}
func (m *ListUsersResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersResp.Merge(m, src)
}
func (m *ListUsersResp) XXX_Size() int {
	return xxx_messageInfo_ListUsersResp.Size(m)
}
func (m *ListUsersResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersResp.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersResp proto.InternalMessageInfo

func (m *ListUsersResp) GetUsers() []*common.User {
	if m != nil {
		return m.Users
	}
	return nil
}

type DeleteUserResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserResp) Reset()         { *m = DeleteUserResp{} }
func (m *DeleteUserResp) String() string { return proto.CompactTextString(m) }
func (*DeleteUserResp) ProtoMessage()    {}
func (*DeleteUserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_525b6514ab436d23, []int{2}
}

func (m *DeleteUserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserResp.Unmarshal(m, b)
}
func (m *DeleteUserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserResp.Marshal(b, m, deterministic)
}
func (m *DeleteUserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserResp.Merge(m, src)
}
func (m *DeleteUserResp) XXX_Size() int {
	return xxx_messageInfo_DeleteUserResp.Size(m)
}
func (m *DeleteUserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserResp.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserResp proto.InternalMessageInfo

type GetUserResp struct {
	User                 *common.User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetUserResp) Reset()         { *m = GetUserResp{} }
func (m *GetUserResp) String() string { return proto.CompactTextString(m) }
func (*GetUserResp) ProtoMessage()    {}
func (*GetUserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_525b6514ab436d23, []int{3}
}

func (m *GetUserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResp.Unmarshal(m, b)
}
func (m *GetUserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResp.Marshal(b, m, deterministic)
}
func (m *GetUserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResp.Merge(m, src)
}
func (m *GetUserResp) XXX_Size() int {
	return xxx_messageInfo_GetUserResp.Size(m)
}
func (m *GetUserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResp proto.InternalMessageInfo

func (m *GetUserResp) GetUser() *common.User {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateUserResp struct {
	User                 *common.User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateUserResp) Reset()         { *m = UpdateUserResp{} }
func (m *UpdateUserResp) String() string { return proto.CompactTextString(m) }
func (*UpdateUserResp) ProtoMessage()    {}
func (*UpdateUserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_525b6514ab436d23, []int{4}
}

func (m *UpdateUserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserResp.Unmarshal(m, b)
}
func (m *UpdateUserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserResp.Marshal(b, m, deterministic)
}
func (m *UpdateUserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserResp.Merge(m, src)
}
func (m *UpdateUserResp) XXX_Size() int {
	return xxx_messageInfo_UpdateUserResp.Size(m)
}
func (m *UpdateUserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserResp proto.InternalMessageInfo

func (m *UpdateUserResp) GetUser() *common.User {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateSelfResp struct {
	User                 *common.User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateSelfResp) Reset()         { *m = UpdateSelfResp{} }
func (m *UpdateSelfResp) String() string { return proto.CompactTextString(m) }
func (*UpdateSelfResp) ProtoMessage()    {}
func (*UpdateSelfResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_525b6514ab436d23, []int{5}
}

func (m *UpdateSelfResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateSelfResp.Unmarshal(m, b)
}
func (m *UpdateSelfResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateSelfResp.Marshal(b, m, deterministic)
}
func (m *UpdateSelfResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateSelfResp.Merge(m, src)
}
func (m *UpdateSelfResp) XXX_Size() int {
	return xxx_messageInfo_UpdateSelfResp.Size(m)
}
func (m *UpdateSelfResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateSelfResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateSelfResp proto.InternalMessageInfo

func (m *UpdateSelfResp) GetUser() *common.User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateUserResp)(nil), "chef.automate.api.iam.v2beta.CreateUserResp")
	proto.RegisterType((*ListUsersResp)(nil), "chef.automate.api.iam.v2beta.ListUsersResp")
	proto.RegisterType((*DeleteUserResp)(nil), "chef.automate.api.iam.v2beta.DeleteUserResp")
	proto.RegisterType((*GetUserResp)(nil), "chef.automate.api.iam.v2beta.GetUserResp")
	proto.RegisterType((*UpdateUserResp)(nil), "chef.automate.api.iam.v2beta.UpdateUserResp")
	proto.RegisterType((*UpdateSelfResp)(nil), "chef.automate.api.iam.v2beta.UpdateSelfResp")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/iam/v2beta/response/users.proto", fileDescriptor_525b6514ab436d23)
}

var fileDescriptor_525b6514ab436d23 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x3d, 0x4b, 0x03, 0x41,
	0x10, 0x86, 0x39, 0xfc, 0x28, 0x36, 0x78, 0x48, 0xaa, 0x20, 0x16, 0x61, 0x2b, 0x1b, 0x77, 0x20,
	0x82, 0xd8, 0x88, 0xe0, 0x07, 0x46, 0x08, 0x16, 0x4a, 0x1a, 0xbb, 0xb9, 0x73, 0x92, 0x2c, 0x64,
	0x6f, 0x96, 0x9d, 0x39, 0xc5, 0x7f, 0x2f, 0xb7, 0xc1, 0x68, 0x25, 0x6a, 0x6c, 0x97, 0x79, 0x9e,
	0x79, 0x97, 0x77, 0xcc, 0x45, 0xcd, 0x21, 0x72, 0x43, 0x8d, 0x0a, 0x60, 0xab, 0x1c, 0x50, 0xe9,
	0x78, 0x8e, 0x4a, 0xaf, 0xf8, 0x06, 0x18, 0x3d, 0x78, 0x0c, 0xf0, 0x32, 0xaa, 0x48, 0x11, 0x12,
	0x49, 0xe4, 0x46, 0x08, 0x5a, 0xa1, 0x24, 0x2e, 0x26, 0x56, 0xee, 0x1f, 0xd6, 0x0b, 0x9a, 0xb9,
	0x0f, 0xd4, 0x61, 0xf4, 0xce, 0x63, 0x70, 0x2b, 0xe4, 0xe0, 0xfc, 0x17, 0xfa, 0x9a, 0x43, 0xe0,
	0xe6, 0xab, 0xdc, 0x8e, 0x4d, 0x79, 0x95, 0x08, 0x95, 0xa6, 0x42, 0xe9, 0x81, 0x24, 0xf6, 0x4f,
	0xcd, 0x76, 0x37, 0x30, 0x28, 0x86, 0xc5, 0x51, 0x6f, 0x64, 0xdd, 0x77, 0xdb, 0x5d, 0xa6, 0xf2,
	0xbc, 0xbd, 0x33, 0x7b, 0x13, 0x2f, 0xda, 0xbd, 0x48, 0x16, 0x9d, 0x99, 0x9d, 0xbc, 0x69, 0x50,
	0x0c, 0xb7, 0x7e, 0x68, 0x5a, 0x01, 0x76, 0xdf, 0x94, 0xd7, 0xb4, 0xa4, 0xcf, 0x50, 0xf6, 0xc6,
	0xf4, 0x6e, 0x49, 0x37, 0xce, 0x38, 0x36, 0xe5, 0x34, 0x3e, 0xff, 0xc7, 0x6f, 0xd7, 0xa6, 0x47,
	0x5a, 0xce, 0x36, 0x31, 0x5d, 0xde, 0x3f, 0x4d, 0xe6, 0x5e, 0x17, 0x6d, 0xe5, 0x6a, 0x0e, 0xd0,
	0x51, 0xeb, 0x1e, 0xe1, 0x0f, 0xa7, 0x53, 0xed, 0xe6, 0x62, 0x4f, 0xde, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xc2, 0x5b, 0x87, 0x29, 0x78, 0x02, 0x00, 0x00,
}
