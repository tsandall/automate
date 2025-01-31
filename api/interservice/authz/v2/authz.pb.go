// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/authz/v2/authz.proto

package v2

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

// TODO (sr): consider dropping this validation:
// a) it takes time
// b) bad input will with certainty lead to an "unauthorized" response
type IsAuthorizedReq struct {
	Subjects             []string `protobuf:"bytes,1,rep,name=subjects,proto3" json:"subjects,omitempty" toml:"subjects,omitempty" mapstructure:"subjects,omitempty"`
	Resource             string   `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty" toml:"resource,omitempty" mapstructure:"resource,omitempty"`
	Action               string   `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty" toml:"action,omitempty" mapstructure:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *IsAuthorizedReq) Reset()         { *m = IsAuthorizedReq{} }
func (m *IsAuthorizedReq) String() string { return proto.CompactTextString(m) }
func (*IsAuthorizedReq) ProtoMessage()    {}
func (*IsAuthorizedReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_744c9d667935614c, []int{0}
}

func (m *IsAuthorizedReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsAuthorizedReq.Unmarshal(m, b)
}
func (m *IsAuthorizedReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsAuthorizedReq.Marshal(b, m, deterministic)
}
func (m *IsAuthorizedReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsAuthorizedReq.Merge(m, src)
}
func (m *IsAuthorizedReq) XXX_Size() int {
	return xxx_messageInfo_IsAuthorizedReq.Size(m)
}
func (m *IsAuthorizedReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IsAuthorizedReq.DiscardUnknown(m)
}

var xxx_messageInfo_IsAuthorizedReq proto.InternalMessageInfo

func (m *IsAuthorizedReq) GetSubjects() []string {
	if m != nil {
		return m.Subjects
	}
	return nil
}

func (m *IsAuthorizedReq) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *IsAuthorizedReq) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

type IsAuthorizedResp struct {
	Authorized           bool     `protobuf:"varint,1,opt,name=authorized,proto3" json:"authorized,omitempty" toml:"authorized,omitempty" mapstructure:"authorized,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *IsAuthorizedResp) Reset()         { *m = IsAuthorizedResp{} }
func (m *IsAuthorizedResp) String() string { return proto.CompactTextString(m) }
func (*IsAuthorizedResp) ProtoMessage()    {}
func (*IsAuthorizedResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_744c9d667935614c, []int{1}
}

func (m *IsAuthorizedResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsAuthorizedResp.Unmarshal(m, b)
}
func (m *IsAuthorizedResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsAuthorizedResp.Marshal(b, m, deterministic)
}
func (m *IsAuthorizedResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsAuthorizedResp.Merge(m, src)
}
func (m *IsAuthorizedResp) XXX_Size() int {
	return xxx_messageInfo_IsAuthorizedResp.Size(m)
}
func (m *IsAuthorizedResp) XXX_DiscardUnknown() {
	xxx_messageInfo_IsAuthorizedResp.DiscardUnknown(m)
}

var xxx_messageInfo_IsAuthorizedResp proto.InternalMessageInfo

func (m *IsAuthorizedResp) GetAuthorized() bool {
	if m != nil {
		return m.Authorized
	}
	return false
}

type ProjectsAuthorizedReq struct {
	Subjects             []string `protobuf:"bytes,1,rep,name=subjects,proto3" json:"subjects,omitempty" toml:"subjects,omitempty" mapstructure:"subjects,omitempty"`
	Resource             string   `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty" toml:"resource,omitempty" mapstructure:"resource,omitempty"`
	Action               string   `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty" toml:"action,omitempty" mapstructure:"action,omitempty"`
	ProjectsFilter       []string `protobuf:"bytes,4,rep,name=projects_filter,json=projectsFilter,proto3" json:"projects_filter,omitempty" toml:"projects_filter,omitempty" mapstructure:"projects_filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ProjectsAuthorizedReq) Reset()         { *m = ProjectsAuthorizedReq{} }
func (m *ProjectsAuthorizedReq) String() string { return proto.CompactTextString(m) }
func (*ProjectsAuthorizedReq) ProtoMessage()    {}
func (*ProjectsAuthorizedReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_744c9d667935614c, []int{2}
}

func (m *ProjectsAuthorizedReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectsAuthorizedReq.Unmarshal(m, b)
}
func (m *ProjectsAuthorizedReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectsAuthorizedReq.Marshal(b, m, deterministic)
}
func (m *ProjectsAuthorizedReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectsAuthorizedReq.Merge(m, src)
}
func (m *ProjectsAuthorizedReq) XXX_Size() int {
	return xxx_messageInfo_ProjectsAuthorizedReq.Size(m)
}
func (m *ProjectsAuthorizedReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectsAuthorizedReq.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectsAuthorizedReq proto.InternalMessageInfo

func (m *ProjectsAuthorizedReq) GetSubjects() []string {
	if m != nil {
		return m.Subjects
	}
	return nil
}

func (m *ProjectsAuthorizedReq) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *ProjectsAuthorizedReq) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *ProjectsAuthorizedReq) GetProjectsFilter() []string {
	if m != nil {
		return m.ProjectsFilter
	}
	return nil
}

type ProjectsAuthorizedResp struct {
	Projects             []string `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty" toml:"projects,omitempty" mapstructure:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ProjectsAuthorizedResp) Reset()         { *m = ProjectsAuthorizedResp{} }
func (m *ProjectsAuthorizedResp) String() string { return proto.CompactTextString(m) }
func (*ProjectsAuthorizedResp) ProtoMessage()    {}
func (*ProjectsAuthorizedResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_744c9d667935614c, []int{3}
}

func (m *ProjectsAuthorizedResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectsAuthorizedResp.Unmarshal(m, b)
}
func (m *ProjectsAuthorizedResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectsAuthorizedResp.Marshal(b, m, deterministic)
}
func (m *ProjectsAuthorizedResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectsAuthorizedResp.Merge(m, src)
}
func (m *ProjectsAuthorizedResp) XXX_Size() int {
	return xxx_messageInfo_ProjectsAuthorizedResp.Size(m)
}
func (m *ProjectsAuthorizedResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectsAuthorizedResp.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectsAuthorizedResp proto.InternalMessageInfo

func (m *ProjectsAuthorizedResp) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type FilterAuthorizedPairsReq struct {
	Subjects             []string `protobuf:"bytes,1,rep,name=subjects,proto3" json:"subjects,omitempty" toml:"subjects,omitempty" mapstructure:"subjects,omitempty"`
	Pairs                []*Pair  `protobuf:"bytes,2,rep,name=pairs,proto3" json:"pairs,omitempty" toml:"pairs,omitempty" mapstructure:"pairs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *FilterAuthorizedPairsReq) Reset()         { *m = FilterAuthorizedPairsReq{} }
func (m *FilterAuthorizedPairsReq) String() string { return proto.CompactTextString(m) }
func (*FilterAuthorizedPairsReq) ProtoMessage()    {}
func (*FilterAuthorizedPairsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_744c9d667935614c, []int{4}
}

func (m *FilterAuthorizedPairsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterAuthorizedPairsReq.Unmarshal(m, b)
}
func (m *FilterAuthorizedPairsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterAuthorizedPairsReq.Marshal(b, m, deterministic)
}
func (m *FilterAuthorizedPairsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterAuthorizedPairsReq.Merge(m, src)
}
func (m *FilterAuthorizedPairsReq) XXX_Size() int {
	return xxx_messageInfo_FilterAuthorizedPairsReq.Size(m)
}
func (m *FilterAuthorizedPairsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterAuthorizedPairsReq.DiscardUnknown(m)
}

var xxx_messageInfo_FilterAuthorizedPairsReq proto.InternalMessageInfo

func (m *FilterAuthorizedPairsReq) GetSubjects() []string {
	if m != nil {
		return m.Subjects
	}
	return nil
}

func (m *FilterAuthorizedPairsReq) GetPairs() []*Pair {
	if m != nil {
		return m.Pairs
	}
	return nil
}

type FilterAuthorizedPairsResp struct {
	Pairs                []*Pair  `protobuf:"bytes,1,rep,name=pairs,proto3" json:"pairs,omitempty" toml:"pairs,omitempty" mapstructure:"pairs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *FilterAuthorizedPairsResp) Reset()         { *m = FilterAuthorizedPairsResp{} }
func (m *FilterAuthorizedPairsResp) String() string { return proto.CompactTextString(m) }
func (*FilterAuthorizedPairsResp) ProtoMessage()    {}
func (*FilterAuthorizedPairsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_744c9d667935614c, []int{5}
}

func (m *FilterAuthorizedPairsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterAuthorizedPairsResp.Unmarshal(m, b)
}
func (m *FilterAuthorizedPairsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterAuthorizedPairsResp.Marshal(b, m, deterministic)
}
func (m *FilterAuthorizedPairsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterAuthorizedPairsResp.Merge(m, src)
}
func (m *FilterAuthorizedPairsResp) XXX_Size() int {
	return xxx_messageInfo_FilterAuthorizedPairsResp.Size(m)
}
func (m *FilterAuthorizedPairsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterAuthorizedPairsResp.DiscardUnknown(m)
}

var xxx_messageInfo_FilterAuthorizedPairsResp proto.InternalMessageInfo

func (m *FilterAuthorizedPairsResp) GetPairs() []*Pair {
	if m != nil {
		return m.Pairs
	}
	return nil
}

type Pair struct {
	Resource             string   `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty" toml:"resource,omitempty" mapstructure:"resource,omitempty"`
	Action               string   `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty" toml:"action,omitempty" mapstructure:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Pair) Reset()         { *m = Pair{} }
func (m *Pair) String() string { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()    {}
func (*Pair) Descriptor() ([]byte, []int) {
	return fileDescriptor_744c9d667935614c, []int{6}
}

func (m *Pair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pair.Unmarshal(m, b)
}
func (m *Pair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pair.Marshal(b, m, deterministic)
}
func (m *Pair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pair.Merge(m, src)
}
func (m *Pair) XXX_Size() int {
	return xxx_messageInfo_Pair.Size(m)
}
func (m *Pair) XXX_DiscardUnknown() {
	xxx_messageInfo_Pair.DiscardUnknown(m)
}

var xxx_messageInfo_Pair proto.InternalMessageInfo

func (m *Pair) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *Pair) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

type FilterAuthorizedProjectsReq struct {
	Subjects             []string `protobuf:"bytes,1,rep,name=subjects,proto3" json:"subjects,omitempty" toml:"subjects,omitempty" mapstructure:"subjects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *FilterAuthorizedProjectsReq) Reset()         { *m = FilterAuthorizedProjectsReq{} }
func (m *FilterAuthorizedProjectsReq) String() string { return proto.CompactTextString(m) }
func (*FilterAuthorizedProjectsReq) ProtoMessage()    {}
func (*FilterAuthorizedProjectsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_744c9d667935614c, []int{7}
}

func (m *FilterAuthorizedProjectsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterAuthorizedProjectsReq.Unmarshal(m, b)
}
func (m *FilterAuthorizedProjectsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterAuthorizedProjectsReq.Marshal(b, m, deterministic)
}
func (m *FilterAuthorizedProjectsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterAuthorizedProjectsReq.Merge(m, src)
}
func (m *FilterAuthorizedProjectsReq) XXX_Size() int {
	return xxx_messageInfo_FilterAuthorizedProjectsReq.Size(m)
}
func (m *FilterAuthorizedProjectsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterAuthorizedProjectsReq.DiscardUnknown(m)
}

var xxx_messageInfo_FilterAuthorizedProjectsReq proto.InternalMessageInfo

func (m *FilterAuthorizedProjectsReq) GetSubjects() []string {
	if m != nil {
		return m.Subjects
	}
	return nil
}

type FilterAuthorizedProjectsResp struct {
	Projects             []string `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty" toml:"projects,omitempty" mapstructure:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *FilterAuthorizedProjectsResp) Reset()         { *m = FilterAuthorizedProjectsResp{} }
func (m *FilterAuthorizedProjectsResp) String() string { return proto.CompactTextString(m) }
func (*FilterAuthorizedProjectsResp) ProtoMessage()    {}
func (*FilterAuthorizedProjectsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_744c9d667935614c, []int{8}
}

func (m *FilterAuthorizedProjectsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterAuthorizedProjectsResp.Unmarshal(m, b)
}
func (m *FilterAuthorizedProjectsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterAuthorizedProjectsResp.Marshal(b, m, deterministic)
}
func (m *FilterAuthorizedProjectsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterAuthorizedProjectsResp.Merge(m, src)
}
func (m *FilterAuthorizedProjectsResp) XXX_Size() int {
	return xxx_messageInfo_FilterAuthorizedProjectsResp.Size(m)
}
func (m *FilterAuthorizedProjectsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterAuthorizedProjectsResp.DiscardUnknown(m)
}

var xxx_messageInfo_FilterAuthorizedProjectsResp proto.InternalMessageInfo

func (m *FilterAuthorizedProjectsResp) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type ValidateProjectAssignmentReq struct {
	Subjects             []string `protobuf:"bytes,1,rep,name=subjects,proto3" json:"subjects,omitempty" toml:"subjects,omitempty" mapstructure:"subjects,omitempty"`
	ProjectIds           []string `protobuf:"bytes,2,rep,name=project_ids,json=projectIds,proto3" json:"project_ids,omitempty" toml:"project_ids,omitempty" mapstructure:"project_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ValidateProjectAssignmentReq) Reset()         { *m = ValidateProjectAssignmentReq{} }
func (m *ValidateProjectAssignmentReq) String() string { return proto.CompactTextString(m) }
func (*ValidateProjectAssignmentReq) ProtoMessage()    {}
func (*ValidateProjectAssignmentReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_744c9d667935614c, []int{9}
}

func (m *ValidateProjectAssignmentReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateProjectAssignmentReq.Unmarshal(m, b)
}
func (m *ValidateProjectAssignmentReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateProjectAssignmentReq.Marshal(b, m, deterministic)
}
func (m *ValidateProjectAssignmentReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateProjectAssignmentReq.Merge(m, src)
}
func (m *ValidateProjectAssignmentReq) XXX_Size() int {
	return xxx_messageInfo_ValidateProjectAssignmentReq.Size(m)
}
func (m *ValidateProjectAssignmentReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateProjectAssignmentReq.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateProjectAssignmentReq proto.InternalMessageInfo

func (m *ValidateProjectAssignmentReq) GetSubjects() []string {
	if m != nil {
		return m.Subjects
	}
	return nil
}

func (m *ValidateProjectAssignmentReq) GetProjectIds() []string {
	if m != nil {
		return m.ProjectIds
	}
	return nil
}

type ValidateProjectAssignmentResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ValidateProjectAssignmentResp) Reset()         { *m = ValidateProjectAssignmentResp{} }
func (m *ValidateProjectAssignmentResp) String() string { return proto.CompactTextString(m) }
func (*ValidateProjectAssignmentResp) ProtoMessage()    {}
func (*ValidateProjectAssignmentResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_744c9d667935614c, []int{10}
}

func (m *ValidateProjectAssignmentResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateProjectAssignmentResp.Unmarshal(m, b)
}
func (m *ValidateProjectAssignmentResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateProjectAssignmentResp.Marshal(b, m, deterministic)
}
func (m *ValidateProjectAssignmentResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateProjectAssignmentResp.Merge(m, src)
}
func (m *ValidateProjectAssignmentResp) XXX_Size() int {
	return xxx_messageInfo_ValidateProjectAssignmentResp.Size(m)
}
func (m *ValidateProjectAssignmentResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateProjectAssignmentResp.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateProjectAssignmentResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*IsAuthorizedReq)(nil), "chef.automate.domain.authz.v2.IsAuthorizedReq")
	proto.RegisterType((*IsAuthorizedResp)(nil), "chef.automate.domain.authz.v2.IsAuthorizedResp")
	proto.RegisterType((*ProjectsAuthorizedReq)(nil), "chef.automate.domain.authz.v2.ProjectsAuthorizedReq")
	proto.RegisterType((*ProjectsAuthorizedResp)(nil), "chef.automate.domain.authz.v2.ProjectsAuthorizedResp")
	proto.RegisterType((*FilterAuthorizedPairsReq)(nil), "chef.automate.domain.authz.v2.FilterAuthorizedPairsReq")
	proto.RegisterType((*FilterAuthorizedPairsResp)(nil), "chef.automate.domain.authz.v2.FilterAuthorizedPairsResp")
	proto.RegisterType((*Pair)(nil), "chef.automate.domain.authz.v2.Pair")
	proto.RegisterType((*FilterAuthorizedProjectsReq)(nil), "chef.automate.domain.authz.v2.FilterAuthorizedProjectsReq")
	proto.RegisterType((*FilterAuthorizedProjectsResp)(nil), "chef.automate.domain.authz.v2.FilterAuthorizedProjectsResp")
	proto.RegisterType((*ValidateProjectAssignmentReq)(nil), "chef.automate.domain.authz.v2.ValidateProjectAssignmentReq")
	proto.RegisterType((*ValidateProjectAssignmentResp)(nil), "chef.automate.domain.authz.v2.ValidateProjectAssignmentResp")
}

func init() {
	proto.RegisterFile("api/interservice/authz/v2/authz.proto", fileDescriptor_744c9d667935614c)
}

var fileDescriptor_744c9d667935614c = []byte{
	// 750 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x56, 0x4d, 0x4f, 0x13, 0x4f,
	0x1c, 0xfe, 0x4f, 0x0b, 0xa4, 0x0c, 0x7f, 0xc5, 0x4c, 0x00, 0x97, 0xe5, 0xad, 0xae, 0x95, 0x6c,
	0xab, 0xdd, 0xd5, 0x15, 0x5f, 0x58, 0x8d, 0x84, 0x1e, 0x24, 0xc4, 0x18, 0xc9, 0x26, 0x82, 0x42,
	0x5a, 0x1c, 0xda, 0x81, 0xae, 0xb6, 0xdd, 0x65, 0x66, 0xdb, 0x43, 0x29, 0x07, 0xc3, 0xd5, 0xc4,
	0xa4, 0x89, 0x89, 0x27, 0x3f, 0x88, 0x27, 0xcf, 0x7c, 0x13, 0xcf, 0x7e, 0x01, 0x33, 0xdb, 0xd9,
	0x52, 0xe8, 0x5b, 0xe8, 0x45, 0x0f, 0xde, 0xba, 0xd3, 0x79, 0x9e, 0x79, 0x9e, 0xdf, 0xf3, 0xfb,
	0xcd, 0x2e, 0xbc, 0x85, 0x5d, 0x5b, 0xb7, 0x4b, 0x1e, 0xa1, 0x8c, 0xd0, 0x8a, 0x9d, 0x25, 0x3a,
	0x2e, 0x7b, 0xf9, 0xaa, 0x5e, 0x31, 0x1a, 0x3f, 0x34, 0x97, 0x3a, 0x9e, 0x83, 0xe6, 0xb2, 0x79,
	0xb2, 0xaf, 0xe1, 0xb2, 0xe7, 0x14, 0xb1, 0x47, 0xb4, 0x9c, 0x53, 0xc4, 0x76, 0x49, 0x6b, 0xec,
	0xa8, 0x18, 0xf2, 0xf5, 0x0a, 0x2e, 0xd8, 0x39, 0xec, 0x11, 0x3d, 0xf8, 0xd1, 0xc0, 0x29, 0xdf,
	0x42, 0x70, 0x7c, 0x9d, 0xad, 0x96, 0xbd, 0xbc, 0x43, 0xed, 0x2a, 0xc9, 0x59, 0xe4, 0x10, 0x7d,
	0x04, 0x30, 0xc2, 0xca, 0x7b, 0xef, 0x49, 0xd6, 0x63, 0x12, 0x88, 0x86, 0xd5, 0xd1, 0x14, 0xf9,
	0xfe, 0xf3, 0x47, 0xf8, 0x5d, 0x1d, 0xa4, 0x23, 0x40, 0x79, 0x4b, 0xb7, 0x8c, 0xd7, 0x19, 0x75,
	0xc5, 0xf4, 0x08, 0x2e, 0xd6, 0xca, 0x8c, 0xd0, 0xb8, 0xa9, 0xae, 0x98, 0x05, 0x27, 0x8b, 0x0b,
	0xb5, 0x42, 0x0e, 0xbb, 0x35, 0x86, 0x8b, 0x85, 0xb8, 0xb9, 0x93, 0x31, 0x13, 0xe9, 0xdb, 0xb1,
	0x5a, 0xc6, 0x73, 0x3e, 0x90, 0x52, 0xcb, 0x63, 0x81, 0x99, 0xc2, 0x8b, 0x58, 0x0c, 0xfe, 0xb3,
	0x9a, 0xc7, 0xa2, 0x67, 0x30, 0x42, 0x09, 0x73, 0xca, 0x34, 0x4b, 0xa4, 0x50, 0x14, 0xa8, 0xa3,
	0x29, 0x85, 0x4b, 0x98, 0xa3, 0x33, 0xc6, 0x74, 0x66, 0x07, 0x27, 0xab, 0x69, 0x1f, 0x93, 0x50,
	0x57, 0x4c, 0x81, 0x8e, 0x27, 0x62, 0x56, 0x13, 0x83, 0xd6, 0xe0, 0x08, 0xce, 0x7a, 0xb6, 0x53,
	0x92, 0xc2, 0x3e, 0x5a, 0xe7, 0xe8, 0x04, 0x55, 0x8d, 0x45, 0x81, 0xc6, 0xc9, 0xea, 0x6a, 0x72,
	0x5b, 0x10, 0x9c, 0x5b, 0x89, 0x1f, 0x19, 0xc7, 0x31, 0x4b, 0xc0, 0x15, 0x03, 0x5e, 0x3b, 0x5f,
	0x1f, 0xe6, 0xa2, 0x79, 0x08, 0x71, 0x73, 0x45, 0x02, 0x51, 0xa0, 0x46, 0xac, 0x96, 0x15, 0xe5,
	0x57, 0x08, 0x4e, 0x6e, 0x50, 0xc7, 0x77, 0xf2, 0xaf, 0xb4, 0xdd, 0x4a, 0x8b, 0x5e, 0xc2, 0x71,
	0x57, 0x54, 0x69, 0x77, 0xdf, 0x2e, 0x78, 0x84, 0x4a, 0x43, 0x7e, 0x49, 0x62, 0x9c, 0x71, 0xa1,
	0x0e, 0x66, 0x25, 0xa0, 0x48, 0x74, 0xca, 0x98, 0xf0, 0x89, 0xef, 0x26, 0x97, 0xd5, 0x78, 0x72,
	0x37, 0x7d, 0x74, 0xef, 0xce, 0xc3, 0xa5, 0xe3, 0x98, 0x75, 0x35, 0x00, 0x3f, 0xf7, 0xb1, 0xca,
	0x1b, 0x38, 0xd5, 0xa9, 0xe8, 0xcc, 0xe5, 0x8e, 0x83, 0xbd, 0xa2, 0xe8, 0x0d, 0xc7, 0x75, 0x20,
	0x4b, 0x40, 0x99, 0xa2, 0x13, 0x06, 0x0a, 0x4e, 0x68, 0xe1, 0x6f, 0x62, 0x94, 0x53, 0x00, 0xa5,
	0xc6, 0x21, 0x67, 0xc4, 0x1b, 0xd8, 0xa6, 0x8c, 0x47, 0xca, 0xda, 0x12, 0xdd, 0xe2, 0xe4, 0x56,
	0x1d, 0xbc, 0x8a, 0x00, 0xe5, 0x05, 0x5d, 0x37, 0xd6, 0x78, 0xa2, 0x7d, 0x43, 0xad, 0xf9, 0x59,
	0xd6, 0xda, 0x23, 0x8c, 0x77, 0xc8, 0x70, 0x19, 0x0e, 0xbb, 0x5c, 0x80, 0x14, 0x8a, 0x86, 0xd5,
	0x31, 0xe3, 0xa6, 0xd6, 0x73, 0xfc, 0x35, 0x2e, 0xd6, 0x6a, 0x20, 0x94, 0x4d, 0x38, 0xdd, 0xc5,
	0x0b, 0x73, 0xcf, 0x78, 0xc1, 0xa5, 0x79, 0x3f, 0x03, 0x38, 0xc4, 0x9f, 0xff, 0x9e, 0xd1, 0xad,
	0x03, 0x38, 0xd3, 0x66, 0x55, 0x64, 0xfa, 0xa7, 0x92, 0x53, 0x4c, 0x38, 0xdb, 0x5d, 0x13, 0x73,
	0x91, 0x7c, 0xb1, 0x57, 0x5b, 0xfa, 0xf0, 0x04, 0xc0, 0xd9, 0x4d, 0x71, 0x7f, 0x0b, 0xd0, 0x2a,
	0x63, 0xf6, 0x41, 0xa9, 0x48, 0x4a, 0x1e, 0x77, 0x24, 0x5f, 0x74, 0xd4, 0xd2, 0x32, 0x6b, 0x70,
	0x4c, 0x10, 0xed, 0xda, 0xb9, 0x46, 0xe3, 0x8c, 0xa6, 0x16, 0xb9, 0xe1, 0x1b, 0x75, 0x30, 0x1f,
	0x01, 0x3d, 0x67, 0x0d, 0x0a, 0xe8, 0x7a, 0x8e, 0x29, 0x0b, 0x70, 0xae, 0x87, 0x08, 0xe6, 0x1a,
	0xa7, 0xc3, 0xf0, 0x4a, 0xe0, 0x0e, 0xfb, 0x93, 0x7e, 0x08, 0xff, 0x6f, 0xbd, 0x44, 0x91, 0xd6,
	0xa7, 0xaf, 0x2e, 0xbc, 0x91, 0x64, 0xfd, 0x52, 0xfb, 0x99, 0xab, 0xfc, 0x87, 0x3e, 0x01, 0x38,
	0xd9, 0xb1, 0xcf, 0xd1, 0xa3, 0x3e, 0x64, 0xdd, 0x26, 0x5d, 0x7e, 0x3c, 0x18, 0xd0, 0x97, 0xf3,
	0xa5, 0xd3, 0x15, 0x22, 0x72, 0x45, 0xe6, 0x65, 0x89, 0xcf, 0x9a, 0x58, 0x7e, 0x32, 0x30, 0xd6,
	0xd7, 0x75, 0x02, 0x20, 0x6a, 0xbf, 0x35, 0xd1, 0x52, 0xbf, 0xc1, 0xef, 0xf4, 0x76, 0x93, 0x1f,
	0x0c, 0x80, 0xf2, 0x55, 0x7c, 0x05, 0x70, 0xba, 0x6b, 0x4f, 0xa1, 0x7e, 0x16, 0x7b, 0x8d, 0x84,
	0xfc, 0x74, 0x70, 0x30, 0x97, 0x96, 0x5a, 0xda, 0x36, 0x0e, 0x6c, 0x2f, 0x5f, 0xde, 0xd3, 0xb2,
	0x4e, 0x51, 0xe7, 0x5c, 0x7a, 0xc0, 0xa5, 0x77, 0xfd, 0x34, 0xdb, 0x1b, 0xf1, 0xbf, 0xae, 0xee,
	0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x26, 0xe5, 0x5d, 0x53, 0xbe, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthorizationClient is the client API for Authorization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorizationClient interface {
	IsAuthorized(ctx context.Context, in *IsAuthorizedReq, opts ...grpc.CallOption) (*IsAuthorizedResp, error)
	FilterAuthorizedPairs(ctx context.Context, in *FilterAuthorizedPairsReq, opts ...grpc.CallOption) (*FilterAuthorizedPairsResp, error)
	FilterAuthorizedProjects(ctx context.Context, in *FilterAuthorizedProjectsReq, opts ...grpc.CallOption) (*FilterAuthorizedProjectsResp, error)
	ProjectsAuthorized(ctx context.Context, in *ProjectsAuthorizedReq, opts ...grpc.CallOption) (*ProjectsAuthorizedResp, error)
	ValidateProjectAssignment(ctx context.Context, in *ValidateProjectAssignmentReq, opts ...grpc.CallOption) (*ValidateProjectAssignmentResp, error)
}

type authorizationClient struct {
	cc *grpc.ClientConn
}

func NewAuthorizationClient(cc *grpc.ClientConn) AuthorizationClient {
	return &authorizationClient{cc}
}

func (c *authorizationClient) IsAuthorized(ctx context.Context, in *IsAuthorizedReq, opts ...grpc.CallOption) (*IsAuthorizedResp, error) {
	out := new(IsAuthorizedResp)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.authz.v2.Authorization/IsAuthorized", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) FilterAuthorizedPairs(ctx context.Context, in *FilterAuthorizedPairsReq, opts ...grpc.CallOption) (*FilterAuthorizedPairsResp, error) {
	out := new(FilterAuthorizedPairsResp)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.authz.v2.Authorization/FilterAuthorizedPairs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) FilterAuthorizedProjects(ctx context.Context, in *FilterAuthorizedProjectsReq, opts ...grpc.CallOption) (*FilterAuthorizedProjectsResp, error) {
	out := new(FilterAuthorizedProjectsResp)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.authz.v2.Authorization/FilterAuthorizedProjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) ProjectsAuthorized(ctx context.Context, in *ProjectsAuthorizedReq, opts ...grpc.CallOption) (*ProjectsAuthorizedResp, error) {
	out := new(ProjectsAuthorizedResp)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.authz.v2.Authorization/ProjectsAuthorized", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) ValidateProjectAssignment(ctx context.Context, in *ValidateProjectAssignmentReq, opts ...grpc.CallOption) (*ValidateProjectAssignmentResp, error) {
	out := new(ValidateProjectAssignmentResp)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.authz.v2.Authorization/ValidateProjectAssignment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationServer is the server API for Authorization service.
type AuthorizationServer interface {
	IsAuthorized(context.Context, *IsAuthorizedReq) (*IsAuthorizedResp, error)
	FilterAuthorizedPairs(context.Context, *FilterAuthorizedPairsReq) (*FilterAuthorizedPairsResp, error)
	FilterAuthorizedProjects(context.Context, *FilterAuthorizedProjectsReq) (*FilterAuthorizedProjectsResp, error)
	ProjectsAuthorized(context.Context, *ProjectsAuthorizedReq) (*ProjectsAuthorizedResp, error)
	ValidateProjectAssignment(context.Context, *ValidateProjectAssignmentReq) (*ValidateProjectAssignmentResp, error)
}

// UnimplementedAuthorizationServer can be embedded to have forward compatible implementations.
type UnimplementedAuthorizationServer struct {
}

func (*UnimplementedAuthorizationServer) IsAuthorized(ctx context.Context, req *IsAuthorizedReq) (*IsAuthorizedResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAuthorized not implemented")
}
func (*UnimplementedAuthorizationServer) FilterAuthorizedPairs(ctx context.Context, req *FilterAuthorizedPairsReq) (*FilterAuthorizedPairsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FilterAuthorizedPairs not implemented")
}
func (*UnimplementedAuthorizationServer) FilterAuthorizedProjects(ctx context.Context, req *FilterAuthorizedProjectsReq) (*FilterAuthorizedProjectsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FilterAuthorizedProjects not implemented")
}
func (*UnimplementedAuthorizationServer) ProjectsAuthorized(ctx context.Context, req *ProjectsAuthorizedReq) (*ProjectsAuthorizedResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProjectsAuthorized not implemented")
}
func (*UnimplementedAuthorizationServer) ValidateProjectAssignment(ctx context.Context, req *ValidateProjectAssignmentReq) (*ValidateProjectAssignmentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateProjectAssignment not implemented")
}

func RegisterAuthorizationServer(s *grpc.Server, srv AuthorizationServer) {
	s.RegisterService(&_Authorization_serviceDesc, srv)
}

func _Authorization_IsAuthorized_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAuthorizedReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).IsAuthorized(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.authz.v2.Authorization/IsAuthorized",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).IsAuthorized(ctx, req.(*IsAuthorizedReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_FilterAuthorizedPairs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterAuthorizedPairsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).FilterAuthorizedPairs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.authz.v2.Authorization/FilterAuthorizedPairs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).FilterAuthorizedPairs(ctx, req.(*FilterAuthorizedPairsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_FilterAuthorizedProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterAuthorizedProjectsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).FilterAuthorizedProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.authz.v2.Authorization/FilterAuthorizedProjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).FilterAuthorizedProjects(ctx, req.(*FilterAuthorizedProjectsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_ProjectsAuthorized_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectsAuthorizedReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).ProjectsAuthorized(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.authz.v2.Authorization/ProjectsAuthorized",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).ProjectsAuthorized(ctx, req.(*ProjectsAuthorizedReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_ValidateProjectAssignment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateProjectAssignmentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).ValidateProjectAssignment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.authz.v2.Authorization/ValidateProjectAssignment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).ValidateProjectAssignment(ctx, req.(*ValidateProjectAssignmentReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authorization_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.domain.authz.v2.Authorization",
	HandlerType: (*AuthorizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsAuthorized",
			Handler:    _Authorization_IsAuthorized_Handler,
		},
		{
			MethodName: "FilterAuthorizedPairs",
			Handler:    _Authorization_FilterAuthorizedPairs_Handler,
		},
		{
			MethodName: "FilterAuthorizedProjects",
			Handler:    _Authorization_FilterAuthorizedProjects_Handler,
		},
		{
			MethodName: "ProjectsAuthorized",
			Handler:    _Authorization_ProjectsAuthorized_Handler,
		},
		{
			MethodName: "ValidateProjectAssignment",
			Handler:    _Authorization_ValidateProjectAssignment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/interservice/authz/v2/authz.proto",
}
