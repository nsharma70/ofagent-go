// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voltha_protos/adapter.proto

package voltha

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	common "github.com/opencord/voltha-protos/v2/go/common"
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

type AdapterConfig struct {
	// Common adapter config attributes here
	LogLevel common.LogLevel_LogLevel `protobuf:"varint,1,opt,name=log_level,json=logLevel,proto3,enum=common.LogLevel_LogLevel" json:"log_level,omitempty"`
	// Custom (vendor-specific) configuration attributes
	AdditionalConfig     *any.Any `protobuf:"bytes,64,opt,name=additional_config,json=additionalConfig,proto3" json:"additional_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdapterConfig) Reset()         { *m = AdapterConfig{} }
func (m *AdapterConfig) String() string { return proto.CompactTextString(m) }
func (*AdapterConfig) ProtoMessage()    {}
func (*AdapterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e998ce153307274, []int{0}
}

func (m *AdapterConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdapterConfig.Unmarshal(m, b)
}
func (m *AdapterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdapterConfig.Marshal(b, m, deterministic)
}
func (m *AdapterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdapterConfig.Merge(m, src)
}
func (m *AdapterConfig) XXX_Size() int {
	return xxx_messageInfo_AdapterConfig.Size(m)
}
func (m *AdapterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AdapterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AdapterConfig proto.InternalMessageInfo

func (m *AdapterConfig) GetLogLevel() common.LogLevel_LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return common.LogLevel_DEBUG
}

func (m *AdapterConfig) GetAdditionalConfig() *any.Any {
	if m != nil {
		return m.AdditionalConfig
	}
	return nil
}

// Adapter (software plugin)
type Adapter struct {
	// Unique name of adapter, matching the python package name under
	// voltha/adapters.
	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Vendor  string `protobuf:"bytes,2,opt,name=vendor,proto3" json:"vendor,omitempty"`
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// Adapter configuration
	Config *AdapterConfig `protobuf:"bytes,16,opt,name=config,proto3" json:"config,omitempty"`
	// Custom descriptors and custom configuration
	AdditionalDescription *any.Any `protobuf:"bytes,64,opt,name=additional_description,json=additionalDescription,proto3" json:"additional_description,omitempty"`
	LogicalDeviceIds      []string `protobuf:"bytes,4,rep,name=logical_device_ids,json=logicalDeviceIds,proto3" json:"logical_device_ids,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *Adapter) Reset()         { *m = Adapter{} }
func (m *Adapter) String() string { return proto.CompactTextString(m) }
func (*Adapter) ProtoMessage()    {}
func (*Adapter) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e998ce153307274, []int{1}
}

func (m *Adapter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Adapter.Unmarshal(m, b)
}
func (m *Adapter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Adapter.Marshal(b, m, deterministic)
}
func (m *Adapter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Adapter.Merge(m, src)
}
func (m *Adapter) XXX_Size() int {
	return xxx_messageInfo_Adapter.Size(m)
}
func (m *Adapter) XXX_DiscardUnknown() {
	xxx_messageInfo_Adapter.DiscardUnknown(m)
}

var xxx_messageInfo_Adapter proto.InternalMessageInfo

func (m *Adapter) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Adapter) GetVendor() string {
	if m != nil {
		return m.Vendor
	}
	return ""
}

func (m *Adapter) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Adapter) GetConfig() *AdapterConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *Adapter) GetAdditionalDescription() *any.Any {
	if m != nil {
		return m.AdditionalDescription
	}
	return nil
}

func (m *Adapter) GetLogicalDeviceIds() []string {
	if m != nil {
		return m.LogicalDeviceIds
	}
	return nil
}

type Adapters struct {
	Items                []*Adapter `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Adapters) Reset()         { *m = Adapters{} }
func (m *Adapters) String() string { return proto.CompactTextString(m) }
func (*Adapters) ProtoMessage()    {}
func (*Adapters) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e998ce153307274, []int{2}
}

func (m *Adapters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Adapters.Unmarshal(m, b)
}
func (m *Adapters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Adapters.Marshal(b, m, deterministic)
}
func (m *Adapters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Adapters.Merge(m, src)
}
func (m *Adapters) XXX_Size() int {
	return xxx_messageInfo_Adapters.Size(m)
}
func (m *Adapters) XXX_DiscardUnknown() {
	xxx_messageInfo_Adapters.DiscardUnknown(m)
}

var xxx_messageInfo_Adapters proto.InternalMessageInfo

func (m *Adapters) GetItems() []*Adapter {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*AdapterConfig)(nil), "voltha.AdapterConfig")
	proto.RegisterType((*Adapter)(nil), "voltha.Adapter")
	proto.RegisterType((*Adapters)(nil), "voltha.Adapters")
}

func init() { proto.RegisterFile("voltha_protos/adapter.proto", fileDescriptor_7e998ce153307274) }

var fileDescriptor_7e998ce153307274 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x8e, 0xda, 0x30,
	0x10, 0x86, 0x95, 0x50, 0x02, 0x18, 0xb5, 0xa5, 0x56, 0xa9, 0x02, 0x15, 0x6a, 0x84, 0x54, 0x29,
	0x87, 0xe2, 0xb4, 0xa9, 0xd4, 0x73, 0xa1, 0x5c, 0x2a, 0x71, 0xca, 0xb1, 0x97, 0x28, 0xc4, 0xc6,
	0x58, 0x72, 0x3c, 0x51, 0x12, 0x22, 0xf1, 0x0a, 0xbd, 0xed, 0x83, 0xed, 0x7b, 0xec, 0x13, 0xec,
	0x79, 0x85, 0x6d, 0x16, 0xd8, 0xc3, 0xde, 0xec, 0xff, 0x9b, 0x99, 0xff, 0xf7, 0x24, 0xe8, 0x73,
	0x0b, 0xb2, 0xd9, 0x67, 0x69, 0x59, 0x41, 0x03, 0x75, 0x94, 0xd1, 0xac, 0x6c, 0x58, 0x45, 0xf4,
	0x15, 0x7b, 0x06, 0x4e, 0x27, 0x1c, 0x80, 0x4b, 0x16, 0x69, 0x75, 0x7b, 0xd8, 0x45, 0x99, 0x3a,
	0x9a, 0x92, 0xe9, 0xf4, 0xb6, 0x3f, 0x87, 0xa2, 0x00, 0x65, 0x99, 0x7f, 0xcb, 0x0a, 0xd6, 0x64,
	0x86, 0xcc, 0xff, 0x3b, 0xe8, 0xed, 0xd2, 0x58, 0xfd, 0x01, 0xb5, 0x13, 0x1c, 0xff, 0x42, 0x03,
	0x09, 0x3c, 0x95, 0xac, 0x65, 0xd2, 0x77, 0x02, 0x27, 0x7c, 0x17, 0x4f, 0x88, 0x9d, 0xb6, 0x01,
	0xbe, 0x39, 0xe9, 0xcf, 0x87, 0xa4, 0x2f, 0xed, 0x09, 0x2f, 0xd1, 0x87, 0x8c, 0x52, 0xd1, 0x08,
	0x50, 0x99, 0x4c, 0x73, 0x3d, 0xcc, 0xff, 0x1d, 0x38, 0xe1, 0x30, 0xfe, 0x48, 0x4c, 0x6c, 0x72,
	0x8e, 0x4d, 0x96, 0xea, 0x98, 0x8c, 0x2e, 0xe5, 0xc6, 0x7a, 0x7e, 0xe7, 0xa2, 0x9e, 0x0d, 0x83,
	0xc7, 0xc8, 0x15, 0x54, 0xfb, 0x0f, 0x56, 0xdd, 0x87, 0xc7, 0xfb, 0x99, 0x93, 0xb8, 0x82, 0xe2,
	0x19, 0xf2, 0x5a, 0xa6, 0x28, 0x54, 0xbe, 0x7b, 0x8d, 0xac, 0x88, 0xbf, 0xa0, 0x5e, 0xcb, 0xaa,
	0x5a, 0x80, 0xf2, 0x3b, 0xd7, 0xfc, 0xac, 0xe2, 0x05, 0xf2, 0x6c, 0xb4, 0x91, 0x8e, 0x36, 0x26,
	0x66, 0x35, 0xe4, 0x66, 0x09, 0x89, 0x2d, 0xc2, 0x09, 0xfa, 0x74, 0xf5, 0x28, 0xca, 0xea, 0xbc,
	0x12, 0xe5, 0xe9, 0xf6, 0xda, 0xcb, 0xce, 0xa6, 0xe3, 0x4b, 0xeb, 0xfa, 0xd2, 0x89, 0xbf, 0x21,
	0x2c, 0x81, 0x8b, 0x5c, 0x0f, 0x6c, 0x45, 0xce, 0x52, 0x41, 0x6b, 0xff, 0x4d, 0xd0, 0x09, 0x07,
	0xc9, 0xc8, 0x92, 0xb5, 0x06, 0x7f, 0x69, 0x3d, 0xff, 0x81, 0xfa, 0x36, 0x5a, 0x8d, 0xbf, 0xa2,
	0xae, 0x68, 0x58, 0x51, 0xfb, 0x4e, 0xd0, 0x09, 0x87, 0xf1, 0xfb, 0x17, 0xd9, 0x13, 0x43, 0x57,
	0xdf, 0xff, 0x11, 0x2e, 0x9a, 0xfd, 0x61, 0x7b, 0xfa, 0x6c, 0x11, 0x94, 0x4c, 0xe5, 0x50, 0xd1,
	0xc8, 0x14, 0x2f, 0xec, 0x3f, 0xd0, 0xc6, 0x11, 0x07, 0xab, 0x6d, 0x3d, 0x2d, 0xfe, 0x7c, 0x0a,
	0x00, 0x00, 0xff, 0xff, 0x63, 0x35, 0x2c, 0x02, 0x84, 0x02, 0x00, 0x00,
}