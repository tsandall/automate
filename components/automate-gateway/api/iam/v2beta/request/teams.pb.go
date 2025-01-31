// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2beta/request/teams.proto

package request

import (
	fmt "fmt"
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

type ListTeamsReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTeamsReq) Reset()         { *m = ListTeamsReq{} }
func (m *ListTeamsReq) String() string { return proto.CompactTextString(m) }
func (*ListTeamsReq) ProtoMessage()    {}
func (*ListTeamsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e5a482ab4c91df5, []int{0}
}

func (m *ListTeamsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTeamsReq.Unmarshal(m, b)
}
func (m *ListTeamsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTeamsReq.Marshal(b, m, deterministic)
}
func (m *ListTeamsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTeamsReq.Merge(m, src)
}
func (m *ListTeamsReq) XXX_Size() int {
	return xxx_messageInfo_ListTeamsReq.Size(m)
}
func (m *ListTeamsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTeamsReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListTeamsReq proto.InternalMessageInfo

type GetTeamReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTeamReq) Reset()         { *m = GetTeamReq{} }
func (m *GetTeamReq) String() string { return proto.CompactTextString(m) }
func (*GetTeamReq) ProtoMessage()    {}
func (*GetTeamReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e5a482ab4c91df5, []int{1}
}

func (m *GetTeamReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTeamReq.Unmarshal(m, b)
}
func (m *GetTeamReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTeamReq.Marshal(b, m, deterministic)
}
func (m *GetTeamReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTeamReq.Merge(m, src)
}
func (m *GetTeamReq) XXX_Size() int {
	return xxx_messageInfo_GetTeamReq.Size(m)
}
func (m *GetTeamReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTeamReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetTeamReq proto.InternalMessageInfo

func (m *GetTeamReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CreateTeamReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Projects             []string `protobuf:"bytes,3,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTeamReq) Reset()         { *m = CreateTeamReq{} }
func (m *CreateTeamReq) String() string { return proto.CompactTextString(m) }
func (*CreateTeamReq) ProtoMessage()    {}
func (*CreateTeamReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e5a482ab4c91df5, []int{2}
}

func (m *CreateTeamReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTeamReq.Unmarshal(m, b)
}
func (m *CreateTeamReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTeamReq.Marshal(b, m, deterministic)
}
func (m *CreateTeamReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTeamReq.Merge(m, src)
}
func (m *CreateTeamReq) XXX_Size() int {
	return xxx_messageInfo_CreateTeamReq.Size(m)
}
func (m *CreateTeamReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTeamReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTeamReq proto.InternalMessageInfo

func (m *CreateTeamReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateTeamReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateTeamReq) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type UpdateTeamReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Projects             []string `protobuf:"bytes,3,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTeamReq) Reset()         { *m = UpdateTeamReq{} }
func (m *UpdateTeamReq) String() string { return proto.CompactTextString(m) }
func (*UpdateTeamReq) ProtoMessage()    {}
func (*UpdateTeamReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e5a482ab4c91df5, []int{3}
}

func (m *UpdateTeamReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTeamReq.Unmarshal(m, b)
}
func (m *UpdateTeamReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTeamReq.Marshal(b, m, deterministic)
}
func (m *UpdateTeamReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTeamReq.Merge(m, src)
}
func (m *UpdateTeamReq) XXX_Size() int {
	return xxx_messageInfo_UpdateTeamReq.Size(m)
}
func (m *UpdateTeamReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTeamReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTeamReq proto.InternalMessageInfo

func (m *UpdateTeamReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateTeamReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateTeamReq) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type DeleteTeamReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTeamReq) Reset()         { *m = DeleteTeamReq{} }
func (m *DeleteTeamReq) String() string { return proto.CompactTextString(m) }
func (*DeleteTeamReq) ProtoMessage()    {}
func (*DeleteTeamReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e5a482ab4c91df5, []int{4}
}

func (m *DeleteTeamReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTeamReq.Unmarshal(m, b)
}
func (m *DeleteTeamReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTeamReq.Marshal(b, m, deterministic)
}
func (m *DeleteTeamReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTeamReq.Merge(m, src)
}
func (m *DeleteTeamReq) XXX_Size() int {
	return xxx_messageInfo_DeleteTeamReq.Size(m)
}
func (m *DeleteTeamReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTeamReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTeamReq proto.InternalMessageInfo

func (m *DeleteTeamReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type AddTeamMembersReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserIds              []string `protobuf:"bytes,2,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddTeamMembersReq) Reset()         { *m = AddTeamMembersReq{} }
func (m *AddTeamMembersReq) String() string { return proto.CompactTextString(m) }
func (*AddTeamMembersReq) ProtoMessage()    {}
func (*AddTeamMembersReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e5a482ab4c91df5, []int{5}
}

func (m *AddTeamMembersReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddTeamMembersReq.Unmarshal(m, b)
}
func (m *AddTeamMembersReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddTeamMembersReq.Marshal(b, m, deterministic)
}
func (m *AddTeamMembersReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddTeamMembersReq.Merge(m, src)
}
func (m *AddTeamMembersReq) XXX_Size() int {
	return xxx_messageInfo_AddTeamMembersReq.Size(m)
}
func (m *AddTeamMembersReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddTeamMembersReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddTeamMembersReq proto.InternalMessageInfo

func (m *AddTeamMembersReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AddTeamMembersReq) GetUserIds() []string {
	if m != nil {
		return m.UserIds
	}
	return nil
}

type GetTeamMembershipReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTeamMembershipReq) Reset()         { *m = GetTeamMembershipReq{} }
func (m *GetTeamMembershipReq) String() string { return proto.CompactTextString(m) }
func (*GetTeamMembershipReq) ProtoMessage()    {}
func (*GetTeamMembershipReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e5a482ab4c91df5, []int{6}
}

func (m *GetTeamMembershipReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTeamMembershipReq.Unmarshal(m, b)
}
func (m *GetTeamMembershipReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTeamMembershipReq.Marshal(b, m, deterministic)
}
func (m *GetTeamMembershipReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTeamMembershipReq.Merge(m, src)
}
func (m *GetTeamMembershipReq) XXX_Size() int {
	return xxx_messageInfo_GetTeamMembershipReq.Size(m)
}
func (m *GetTeamMembershipReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTeamMembershipReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetTeamMembershipReq proto.InternalMessageInfo

func (m *GetTeamMembershipReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type RemoveTeamMembersReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserIds              []string `protobuf:"bytes,2,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveTeamMembersReq) Reset()         { *m = RemoveTeamMembersReq{} }
func (m *RemoveTeamMembersReq) String() string { return proto.CompactTextString(m) }
func (*RemoveTeamMembersReq) ProtoMessage()    {}
func (*RemoveTeamMembersReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e5a482ab4c91df5, []int{7}
}

func (m *RemoveTeamMembersReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveTeamMembersReq.Unmarshal(m, b)
}
func (m *RemoveTeamMembersReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveTeamMembersReq.Marshal(b, m, deterministic)
}
func (m *RemoveTeamMembersReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveTeamMembersReq.Merge(m, src)
}
func (m *RemoveTeamMembersReq) XXX_Size() int {
	return xxx_messageInfo_RemoveTeamMembersReq.Size(m)
}
func (m *RemoveTeamMembersReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveTeamMembersReq.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveTeamMembersReq proto.InternalMessageInfo

func (m *RemoveTeamMembersReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RemoveTeamMembersReq) GetUserIds() []string {
	if m != nil {
		return m.UserIds
	}
	return nil
}

type GetTeamsForMemberReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTeamsForMemberReq) Reset()         { *m = GetTeamsForMemberReq{} }
func (m *GetTeamsForMemberReq) String() string { return proto.CompactTextString(m) }
func (*GetTeamsForMemberReq) ProtoMessage()    {}
func (*GetTeamsForMemberReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e5a482ab4c91df5, []int{8}
}

func (m *GetTeamsForMemberReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTeamsForMemberReq.Unmarshal(m, b)
}
func (m *GetTeamsForMemberReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTeamsForMemberReq.Marshal(b, m, deterministic)
}
func (m *GetTeamsForMemberReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTeamsForMemberReq.Merge(m, src)
}
func (m *GetTeamsForMemberReq) XXX_Size() int {
	return xxx_messageInfo_GetTeamsForMemberReq.Size(m)
}
func (m *GetTeamsForMemberReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTeamsForMemberReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetTeamsForMemberReq proto.InternalMessageInfo

func (m *GetTeamsForMemberReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ApplyV2DataMigrationsReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplyV2DataMigrationsReq) Reset()         { *m = ApplyV2DataMigrationsReq{} }
func (m *ApplyV2DataMigrationsReq) String() string { return proto.CompactTextString(m) }
func (*ApplyV2DataMigrationsReq) ProtoMessage()    {}
func (*ApplyV2DataMigrationsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e5a482ab4c91df5, []int{9}
}

func (m *ApplyV2DataMigrationsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyV2DataMigrationsReq.Unmarshal(m, b)
}
func (m *ApplyV2DataMigrationsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyV2DataMigrationsReq.Marshal(b, m, deterministic)
}
func (m *ApplyV2DataMigrationsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyV2DataMigrationsReq.Merge(m, src)
}
func (m *ApplyV2DataMigrationsReq) XXX_Size() int {
	return xxx_messageInfo_ApplyV2DataMigrationsReq.Size(m)
}
func (m *ApplyV2DataMigrationsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyV2DataMigrationsReq.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyV2DataMigrationsReq proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ListTeamsReq)(nil), "chef.automate.api.iam.v2beta.ListTeamsReq")
	proto.RegisterType((*GetTeamReq)(nil), "chef.automate.api.iam.v2beta.GetTeamReq")
	proto.RegisterType((*CreateTeamReq)(nil), "chef.automate.api.iam.v2beta.CreateTeamReq")
	proto.RegisterType((*UpdateTeamReq)(nil), "chef.automate.api.iam.v2beta.UpdateTeamReq")
	proto.RegisterType((*DeleteTeamReq)(nil), "chef.automate.api.iam.v2beta.DeleteTeamReq")
	proto.RegisterType((*AddTeamMembersReq)(nil), "chef.automate.api.iam.v2beta.AddTeamMembersReq")
	proto.RegisterType((*GetTeamMembershipReq)(nil), "chef.automate.api.iam.v2beta.GetTeamMembershipReq")
	proto.RegisterType((*RemoveTeamMembersReq)(nil), "chef.automate.api.iam.v2beta.RemoveTeamMembersReq")
	proto.RegisterType((*GetTeamsForMemberReq)(nil), "chef.automate.api.iam.v2beta.GetTeamsForMemberReq")
	proto.RegisterType((*ApplyV2DataMigrationsReq)(nil), "chef.automate.api.iam.v2beta.ApplyV2DataMigrationsReq")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/iam/v2beta/request/teams.proto", fileDescriptor_7e5a482ab4c91df5)
}

var fileDescriptor_7e5a482ab4c91df5 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xc1, 0x6b, 0xf2, 0x40,
	0x10, 0xc5, 0x31, 0x7e, 0x7c, 0xd5, 0xa1, 0x0a, 0x0d, 0x1e, 0x52, 0x11, 0x2a, 0x39, 0x14, 0x2f,
	0xcd, 0x82, 0xbd, 0x0b, 0xb6, 0xd2, 0x52, 0x5a, 0x29, 0x84, 0xb6, 0x87, 0x5e, 0xca, 0x24, 0x99,
	0xea, 0x16, 0x37, 0xbb, 0xee, 0x4e, 0x2c, 0xfe, 0xf7, 0x25, 0xd1, 0x7a, 0xd2, 0x43, 0xc1, 0xdb,
	0xbe, 0x9d, 0x79, 0xef, 0x07, 0xc3, 0x83, 0x51, 0xaa, 0x95, 0xd1, 0x39, 0xe5, 0xec, 0x04, 0x16,
	0xac, 0x15, 0x32, 0x5d, 0xcd, 0x90, 0xe9, 0x1b, 0xd7, 0x02, 0x8d, 0x14, 0x12, 0x95, 0x58, 0x0d,
	0x13, 0x62, 0x14, 0x96, 0x96, 0x05, 0x39, 0x16, 0x4c, 0xa8, 0x5c, 0x64, 0xac, 0x66, 0xed, 0xf7,
	0xd2, 0x39, 0x7d, 0x46, 0xbf, 0xce, 0x08, 0x8d, 0x8c, 0x24, 0xaa, 0x68, 0xe3, 0x08, 0xdb, 0x70,
	0xfa, 0x24, 0x1d, 0xbf, 0x94, 0x86, 0x98, 0x96, 0x61, 0x0f, 0xe0, 0x9e, 0x2a, 0x19, 0xd3, 0xd2,
	0x6f, 0x83, 0x27, 0xb3, 0xa0, 0xd6, 0xaf, 0x0d, 0x9a, 0xb1, 0x27, 0xb3, 0xf0, 0x19, 0x5a, 0xb7,
	0x96, 0x90, 0xe9, 0xc0, 0x82, 0xef, 0xc3, 0xbf, 0x1c, 0x15, 0x05, 0x5e, 0xf5, 0x53, 0xbd, 0xfd,
	0x2e, 0x34, 0x8c, 0xd5, 0x5f, 0x94, 0xb2, 0x0b, 0xea, 0xfd, 0xfa, 0xa0, 0x19, 0xef, 0x74, 0x19,
	0xf8, 0x6a, 0xb2, 0x23, 0x06, 0x5e, 0x40, 0x6b, 0x42, 0x0b, 0x3a, 0x18, 0x18, 0x8e, 0xe0, 0x6c,
	0x9c, 0x65, 0xe5, 0x74, 0x4a, 0x2a, 0x21, 0xeb, 0xf6, 0x51, 0xcf, 0xa1, 0x51, 0x38, 0xb2, 0x1f,
	0x32, 0x73, 0x81, 0x57, 0x11, 0x4e, 0x4a, 0xfd, 0x90, 0xb9, 0xf0, 0x12, 0x3a, 0xdb, 0x03, 0x6d,
	0xfd, 0x73, 0x69, 0xf6, 0x71, 0xc6, 0xd0, 0x89, 0x49, 0xe9, 0x15, 0x1d, 0x03, 0xe5, 0xee, 0xb4,
	0xdd, 0x64, 0xec, 0x43, 0x75, 0x21, 0x18, 0x1b, 0xb3, 0x58, 0xbf, 0x0d, 0x27, 0xc8, 0x38, 0x95,
	0x33, 0x8b, 0x2c, 0x75, 0x5e, 0xe2, 0x6e, 0xa6, 0xef, 0x8f, 0x33, 0xc9, 0xf3, 0x22, 0x89, 0x52,
	0xad, 0x44, 0x59, 0x85, 0x5d, 0x89, 0xc4, 0xdf, 0x8b, 0x95, 0xfc, 0xaf, 0x3a, 0x75, 0xfd, 0x13,
	0x00, 0x00, 0xff, 0xff, 0x5b, 0x7d, 0x94, 0x5a, 0x95, 0x02, 0x00, 0x00,
}
