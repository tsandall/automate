// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/config/authz/config_request.proto

package authz

import (
	fmt "fmt"
	shared "github.com/chef/automate/api/config/shared"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-a2-config/api/a2conf"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type ConfigRequest struct {
	V1                   *ConfigRequest_V1 `protobuf:"bytes,3,opt,name=v1,proto3" json:"v1,omitempty" toml:"v1,omitempty" mapstructure:"v1,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte            `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32             `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest) Reset()         { *m = ConfigRequest{} }
func (m *ConfigRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest) ProtoMessage()    {}
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6d3fc06ac00b631, []int{0}
}

func (m *ConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest.Unmarshal(m, b)
}
func (m *ConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest.Marshal(b, m, deterministic)
}
func (m *ConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest.Merge(m, src)
}
func (m *ConfigRequest) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest.Size(m)
}
func (m *ConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest proto.InternalMessageInfo

func (m *ConfigRequest) GetV1() *ConfigRequest_V1 {
	if m != nil {
		return m.V1
	}
	return nil
}

type ConfigRequest_V1 struct {
	Sys                  *ConfigRequest_V1_System  `protobuf:"bytes,1,opt,name=sys,proto3" json:"sys,omitempty" toml:"sys,omitempty" mapstructure:"sys,omitempty"`
	Svc                  *ConfigRequest_V1_Service `protobuf:"bytes,2,opt,name=svc,proto3" json:"svc,omitempty" toml:"svc,omitempty" mapstructure:"svc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                    `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                     `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1) Reset()         { *m = ConfigRequest_V1{} }
func (m *ConfigRequest_V1) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1) ProtoMessage()    {}
func (*ConfigRequest_V1) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6d3fc06ac00b631, []int{0, 0}
}

func (m *ConfigRequest_V1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1.Unmarshal(m, b)
}
func (m *ConfigRequest_V1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1.Merge(m, src)
}
func (m *ConfigRequest_V1) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1.Size(m)
}
func (m *ConfigRequest_V1) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1 proto.InternalMessageInfo

func (m *ConfigRequest_V1) GetSys() *ConfigRequest_V1_System {
	if m != nil {
		return m.Sys
	}
	return nil
}

func (m *ConfigRequest_V1) GetSvc() *ConfigRequest_V1_Service {
	if m != nil {
		return m.Svc
	}
	return nil
}

type ConfigRequest_V1_System struct {
	Mlsa                 *shared.Mlsa                     `protobuf:"bytes,1,opt,name=mlsa,proto3" json:"mlsa,omitempty" toml:"mlsa,omitempty" mapstructure:"mlsa,omitempty"`
	Tls                  *shared.TLSCredentials           `protobuf:"bytes,2,opt,name=tls,proto3" json:"tls,omitempty" toml:"tls,omitempty" mapstructure:"tls,omitempty"`
	Service              *ConfigRequest_V1_System_Service `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty" toml:"service,omitempty" mapstructure:"service,omitempty"`
	Logger               *ConfigRequest_V1_System_Logger  `protobuf:"bytes,4,opt,name=logger,proto3" json:"logger,omitempty" toml:"logger,omitempty" mapstructure:"logger,omitempty"`
	Storage              *ConfigRequest_V1_System_Storage `protobuf:"bytes,5,opt,name=storage,proto3" json:"storage,omitempty" toml:"storage,omitempty" mapstructure:"storage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                           `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                            `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System) Reset()         { *m = ConfigRequest_V1_System{} }
func (m *ConfigRequest_V1_System) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System) ProtoMessage()    {}
func (*ConfigRequest_V1_System) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6d3fc06ac00b631, []int{0, 0, 0}
}

func (m *ConfigRequest_V1_System) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System.Merge(m, src)
}
func (m *ConfigRequest_V1_System) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System.Size(m)
}
func (m *ConfigRequest_V1_System) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System proto.InternalMessageInfo

func (m *ConfigRequest_V1_System) GetMlsa() *shared.Mlsa {
	if m != nil {
		return m.Mlsa
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetTls() *shared.TLSCredentials {
	if m != nil {
		return m.Tls
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetService() *ConfigRequest_V1_System_Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetLogger() *ConfigRequest_V1_System_Logger {
	if m != nil {
		return m.Logger
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetStorage() *ConfigRequest_V1_System_Storage {
	if m != nil {
		return m.Storage
	}
	return nil
}

type ConfigRequest_V1_System_Service struct {
	Host                 *wrappers.StringValue `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty" toml:"host,omitempty" mapstructure:"host,omitempty"`
	Port                 *wrappers.Int32Value  `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty" toml:"port,omitempty" mapstructure:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Service) Reset()         { *m = ConfigRequest_V1_System_Service{} }
func (m *ConfigRequest_V1_System_Service) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Service) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6d3fc06ac00b631, []int{0, 0, 0, 0}
}

func (m *ConfigRequest_V1_System_Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Service.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Service.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System_Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Service.Merge(m, src)
}
func (m *ConfigRequest_V1_System_Service) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Service.Size(m)
}
func (m *ConfigRequest_V1_System_Service) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Service.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Service proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Service) GetHost() *wrappers.StringValue {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetPort() *wrappers.Int32Value {
	if m != nil {
		return m.Port
	}
	return nil
}

type ConfigRequest_V1_System_Logger struct {
	Format               *wrappers.StringValue `protobuf:"bytes,1,opt,name=format,proto3" json:"format,omitempty" toml:"format,omitempty" mapstructure:"format,omitempty"`
	Level                *wrappers.StringValue `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty" toml:"level,omitempty" mapstructure:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Logger) Reset()         { *m = ConfigRequest_V1_System_Logger{} }
func (m *ConfigRequest_V1_System_Logger) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Logger) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Logger) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6d3fc06ac00b631, []int{0, 0, 0, 1}
}

func (m *ConfigRequest_V1_System_Logger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Logger.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Logger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Logger.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System_Logger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Logger.Merge(m, src)
}
func (m *ConfigRequest_V1_System_Logger) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Logger.Size(m)
}
func (m *ConfigRequest_V1_System_Logger) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Logger.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Logger proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Logger) GetFormat() *wrappers.StringValue {
	if m != nil {
		return m.Format
	}
	return nil
}

func (m *ConfigRequest_V1_System_Logger) GetLevel() *wrappers.StringValue {
	if m != nil {
		return m.Level
	}
	return nil
}

type ConfigRequest_V1_System_Storage struct {
	Database             *wrappers.StringValue `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty" toml:"database,omitempty" mapstructure:"database,omitempty"`
	User                 *wrappers.StringValue `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty" toml:"user,omitempty" mapstructure:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Storage) Reset()         { *m = ConfigRequest_V1_System_Storage{} }
func (m *ConfigRequest_V1_System_Storage) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Storage) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Storage) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6d3fc06ac00b631, []int{0, 0, 0, 2}
}

func (m *ConfigRequest_V1_System_Storage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Storage.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Storage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Storage.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System_Storage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Storage.Merge(m, src)
}
func (m *ConfigRequest_V1_System_Storage) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Storage.Size(m)
}
func (m *ConfigRequest_V1_System_Storage) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Storage.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Storage proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Storage) GetDatabase() *wrappers.StringValue {
	if m != nil {
		return m.Database
	}
	return nil
}

func (m *ConfigRequest_V1_System_Storage) GetUser() *wrappers.StringValue {
	if m != nil {
		return m.User
	}
	return nil
}

type ConfigRequest_V1_Service struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_Service) Reset()         { *m = ConfigRequest_V1_Service{} }
func (m *ConfigRequest_V1_Service) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_Service) ProtoMessage()    {}
func (*ConfigRequest_V1_Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6d3fc06ac00b631, []int{0, 0, 1}
}

func (m *ConfigRequest_V1_Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_Service.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_Service.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_Service.Merge(m, src)
}
func (m *ConfigRequest_V1_Service) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_Service.Size(m)
}
func (m *ConfigRequest_V1_Service) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_Service.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_Service proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ConfigRequest)(nil), "chef.automate.domain.authz.ConfigRequest")
	proto.RegisterType((*ConfigRequest_V1)(nil), "chef.automate.domain.authz.ConfigRequest.V1")
	proto.RegisterType((*ConfigRequest_V1_System)(nil), "chef.automate.domain.authz.ConfigRequest.V1.System")
	proto.RegisterType((*ConfigRequest_V1_System_Service)(nil), "chef.automate.domain.authz.ConfigRequest.V1.System.Service")
	proto.RegisterType((*ConfigRequest_V1_System_Logger)(nil), "chef.automate.domain.authz.ConfigRequest.V1.System.Logger")
	proto.RegisterType((*ConfigRequest_V1_System_Storage)(nil), "chef.automate.domain.authz.ConfigRequest.V1.System.Storage")
	proto.RegisterType((*ConfigRequest_V1_Service)(nil), "chef.automate.domain.authz.ConfigRequest.V1.Service")
}

func init() {
	proto.RegisterFile("api/config/authz/config_request.proto", fileDescriptor_c6d3fc06ac00b631)
}

var fileDescriptor_c6d3fc06ac00b631 = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x5d, 0x6b, 0xd4, 0x40,
	0x14, 0x86, 0xd9, 0x4d, 0xf6, 0xc3, 0x91, 0x62, 0x19, 0x41, 0x42, 0x2a, 0xa5, 0x08, 0x82, 0x5f,
	0x3b, 0x71, 0x3f, 0x2e, 0xa4, 0xf6, 0xaa, 0x45, 0xc5, 0xa5, 0x22, 0x64, 0x75, 0x2f, 0xbc, 0x91,
	0xb3, 0xd9, 0xd9, 0xd9, 0xc0, 0x64, 0x26, 0xce, 0x4c, 0x56, 0xaa, 0x77, 0xe2, 0x2f, 0xf0, 0xc6,
	0xff, 0xd3, 0xbf, 0x54, 0xf0, 0x5a, 0x32, 0x99, 0x2c, 0xda, 0x62, 0xd9, 0xd6, 0xcb, 0x70, 0xce,
	0xfb, 0xcc, 0x7b, 0xde, 0x33, 0x19, 0x74, 0x1f, 0xf2, 0x34, 0x4a, 0xa4, 0x58, 0xa4, 0x2c, 0x82,
	0xc2, 0x2c, 0xbf, 0xb8, 0x8f, 0x8f, 0x8a, 0x7e, 0x2a, 0xa8, 0x36, 0x24, 0x57, 0xd2, 0x48, 0x1c,
	0x26, 0x4b, 0xba, 0x20, 0x50, 0x18, 0x99, 0x81, 0xa1, 0x64, 0x2e, 0x33, 0x48, 0x05, 0xb1, 0x82,
	0x70, 0xf7, 0x0f, 0x84, 0x5e, 0x82, 0xa2, 0xf3, 0x88, 0x71, 0x39, 0x03, 0x5e, 0x69, 0xc3, 0x9d,
	0x8b, 0x75, 0xc3, 0xb5, 0x2b, 0x8e, 0x13, 0x99, 0xe5, 0x52, 0x50, 0x61, 0x74, 0x54, 0xe3, 0x7b,
	0x4c, 0xe5, 0x49, 0x64, 0xeb, 0x49, 0x8f, 0x51, 0xd1, 0x83, 0x41, 0xaf, 0xb6, 0x98, 0xa7, 0x11,
	0x0c, 0xca, 0x8f, 0x08, 0x84, 0x90, 0x06, 0x4c, 0x2a, 0x45, 0xcd, 0xda, 0x65, 0x52, 0x32, 0x4e,
	0x2b, 0xe5, 0xac, 0x58, 0x44, 0x9f, 0x15, 0xe4, 0x39, 0x55, 0xae, 0x7e, 0xef, 0x5b, 0x17, 0x6d,
	0x1d, 0x59, 0x4e, 0x5c, 0x0d, 0x87, 0x0f, 0x50, 0x73, 0xd5, 0x0f, 0xbc, 0xbd, 0xc6, 0x83, 0x9b,
	0x83, 0x27, 0xe4, 0xdf, 0x33, 0x92, 0xbf, 0x64, 0x64, 0xda, 0x8f, 0x9b, 0xab, 0x7e, 0xf8, 0xb3,
	0x83, 0x9a, 0xd3, 0x3e, 0x7e, 0x81, 0x3c, 0x7d, 0xa2, 0x83, 0x86, 0xa5, 0x0c, 0xaf, 0x42, 0x21,
	0x93, 0x13, 0x6d, 0x68, 0x16, 0x97, 0x7a, 0xfc, 0x12, 0x79, 0x7a, 0x95, 0x04, 0x4d, 0x8b, 0x19,
	0x5d, 0x0d, 0x43, 0xd5, 0x2a, 0x4d, 0x68, 0x5c, 0x02, 0xc2, 0x5f, 0x2d, 0xd4, 0xae, 0xb8, 0x78,
	0x84, 0xfc, 0x8c, 0x6b, 0x70, 0xd6, 0xf6, 0xce, 0x31, 0x53, 0xb1, 0x50, 0x40, 0xaa, 0x60, 0xc9,
	0x1b, 0xae, 0x21, 0xb6, 0xdd, 0xf8, 0x00, 0x79, 0x86, 0x6b, 0x67, 0xe4, 0xd1, 0x65, 0xa2, 0x77,
	0xc7, 0x93, 0x23, 0x45, 0xe7, 0x54, 0x98, 0x14, 0xb8, 0x8e, 0x4b, 0x19, 0x7e, 0x8f, 0x3a, 0xba,
	0xb2, 0xe3, 0x72, 0x7d, 0x7e, 0x8d, 0x44, 0xd6, 0x13, 0xd5, 0x2c, 0x1c, 0xa3, 0x36, 0x97, 0x8c,
	0x51, 0x15, 0xf8, 0x96, 0xba, 0x7f, 0x1d, 0xea, 0xb1, 0x25, 0xc4, 0x8e, 0x64, 0xad, 0x1a, 0xa9,
	0x80, 0xd1, 0xa0, 0xf5, 0x1f, 0x56, 0x2b, 0x44, 0x5c, 0xb3, 0xc2, 0xef, 0x0d, 0xd4, 0x71, 0xfe,
	0xf1, 0x53, 0xe4, 0x2f, 0xa5, 0x36, 0x6e, 0x03, 0x77, 0x49, 0x75, 0x43, 0x49, 0x7d, 0x43, 0xc9,
	0xc4, 0xa8, 0x54, 0xb0, 0x29, 0xf0, 0x82, 0xc6, 0xb6, 0x13, 0xbf, 0x42, 0x7e, 0x2e, 0x95, 0x71,
	0xf1, 0xef, 0x5c, 0x50, 0xbc, 0x16, 0x66, 0x38, 0xb0, 0x82, 0xc3, 0x3b, 0xa7, 0x67, 0x01, 0x5e,
	0xc7, 0xbd, 0xfd, 0xe3, 0x6d, 0xe8, 0x97, 0x7f, 0x4e, 0x6c, 0x01, 0xa1, 0x42, 0xed, 0x6a, 0x5e,
	0x3c, 0x42, 0xed, 0x85, 0x54, 0x19, 0x6c, 0x66, 0xc3, 0xf5, 0xe2, 0x01, 0x6a, 0x71, 0xba, 0xa2,
	0xdc, 0x39, 0xb9, 0x5c, 0x54, 0xb5, 0x86, 0x5f, 0x51, 0xc7, 0xc5, 0x81, 0x9f, 0xa1, 0xee, 0x1c,
	0x0c, 0xcc, 0x40, 0xd3, 0x8d, 0x8e, 0x5d, 0x77, 0x97, 0x99, 0x15, 0x9a, 0xaa, 0x8d, 0xce, 0xb5,
	0x9d, 0x63, 0xbf, 0xeb, 0x6d, 0xb7, 0xc2, 0x1b, 0xeb, 0xd8, 0xf7, 0x6f, 0x9f, 0x9e, 0x05, 0xb7,
	0xd0, 0x96, 0x5d, 0x5c, 0xcf, 0xe5, 0x33, 0xf6, 0xbb, 0x8d, 0x6d, 0xef, 0xf0, 0xf1, 0x87, 0x87,
	0x2c, 0x35, 0xcb, 0x62, 0x46, 0x12, 0x99, 0x45, 0xe5, 0xbe, 0xd7, 0xef, 0x4e, 0x74, 0xfe, 0x2d,
	0x9c, 0xb5, 0xed, 0xa1, 0xc3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x76, 0x4c, 0xac, 0x1c, 0x26,
	0x05, 0x00, 0x00,
}
