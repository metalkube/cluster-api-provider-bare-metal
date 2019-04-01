// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/securitycenter/v1/source.proto

package securitycenter // import "google.golang.org/genproto/googleapis/cloud/securitycenter/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Cloud Security Command Center's (Cloud SCC) finding source. A finding source
// is an entity or a mechanism that can produce a finding. A source is like a
// container of findings that come from the same scanner, logger, monitor, etc.
type Source struct {
	// The relative resource name of this source. See:
	// https://cloud.google.com/apis/design/resource_names#relative_resource_name
	// Example:
	// "organizations/123/sources/456"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The source’s display name.
	// A source’s display name must be unique amongst its siblings, for example,
	// two sources with the same parent can't share the same display name.
	// The display name must start and end with a letter or digit, may contain
	// letters, digits, spaces, hyphens, and underscores, and can be no longer
	// than 32 characters. This is captured by the regular expression:
	// [\p{L}\p{N}]({\p{L}\p{N}_- ]{0,30}[\p{L}\p{N}])?.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The description of the source (max of 1024 characters).
	// Example:
	// "Cloud Security Scanner is a web security scanner for common
	// vulnerabilities in App Engine applications. It can automatically
	// scan and detect four common vulnerabilities, including cross-site-scripting
	// (XSS), Flash injection, mixed content (HTTP in HTTPS), and
	// outdated/insecure libraries."
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Source) Reset()         { *m = Source{} }
func (m *Source) String() string { return proto.CompactTextString(m) }
func (*Source) ProtoMessage()    {}
func (*Source) Descriptor() ([]byte, []int) {
	return fileDescriptor_source_9289be609709d799, []int{0}
}
func (m *Source) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Source.Unmarshal(m, b)
}
func (m *Source) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Source.Marshal(b, m, deterministic)
}
func (dst *Source) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Source.Merge(dst, src)
}
func (m *Source) XXX_Size() int {
	return xxx_messageInfo_Source.Size(m)
}
func (m *Source) XXX_DiscardUnknown() {
	xxx_messageInfo_Source.DiscardUnknown(m)
}

var xxx_messageInfo_Source proto.InternalMessageInfo

func (m *Source) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Source) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Source) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*Source)(nil), "google.cloud.securitycenter.v1.Source")
}

func init() {
	proto.RegisterFile("google/cloud/securitycenter/v1/source.proto", fileDescriptor_source_9289be609709d799)
}

var fileDescriptor_source_9289be609709d799 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0x03, 0x31,
	0x18, 0x84, 0xd9, 0xba, 0x16, 0x4d, 0x3d, 0xc8, 0x9e, 0x8a, 0x48, 0xa9, 0x3d, 0x09, 0x42, 0xc2,
	0xe2, 0xd1, 0x9b, 0x3d, 0x08, 0x22, 0x22, 0x16, 0x7a, 0x90, 0x05, 0x89, 0xd9, 0x9f, 0x10, 0xd8,
	0xe6, 0x0f, 0x49, 0xb6, 0xd0, 0x57, 0xf2, 0x01, 0x7c, 0x08, 0x9f, 0x4a, 0xfa, 0x27, 0x22, 0x7b,
	0xb0, 0xb7, 0x30, 0xf3, 0x65, 0x66, 0xf8, 0xd9, 0x8d, 0x46, 0xd4, 0x1d, 0x08, 0xd5, 0x61, 0xdf,
	0x8a, 0x00, 0xaa, 0xf7, 0x26, 0xee, 0x14, 0xd8, 0x08, 0x5e, 0x6c, 0x6b, 0x11, 0xb0, 0xf7, 0x0a,
	0xb8, 0xf3, 0x18, 0xb1, 0x9a, 0x25, 0x98, 0x13, 0xcc, 0x87, 0x30, 0xdf, 0xd6, 0x17, 0x97, 0x39,
	0x4c, 0x3a, 0x23, 0xa4, 0xb5, 0x18, 0x65, 0x34, 0x68, 0x43, 0xfa, 0xbd, 0xd0, 0x6c, 0xbc, 0xa2,
	0xb4, 0xaa, 0x62, 0xa5, 0x95, 0x1b, 0x98, 0x16, 0xf3, 0xe2, 0xfa, 0xf4, 0x95, 0xde, 0xd5, 0x15,
	0x3b, 0x6b, 0x4d, 0x70, 0x9d, 0xdc, 0xbd, 0x93, 0x37, 0x22, 0x6f, 0x92, 0xb5, 0xe7, 0x3d, 0x32,
	0x67, 0x93, 0x16, 0x82, 0xf2, 0xc6, 0xed, 0x63, 0xa7, 0x47, 0x99, 0xf8, 0x93, 0x1e, 0xcb, 0x93,
	0xf2, 0xfc, 0xf8, 0xfe, 0xab, 0x60, 0x0b, 0x85, 0x1b, 0x7e, 0x78, 0xed, 0x4b, 0xf1, 0xf6, 0x94,
	0x09, 0x8d, 0x9d, 0xb4, 0x9a, 0xa3, 0xd7, 0x42, 0x83, 0xa5, 0xb5, 0x22, 0x59, 0xd2, 0x99, 0xf0,
	0xdf, 0x6d, 0xee, 0x86, 0xca, 0xe7, 0x68, 0xf6, 0x90, 0xe2, 0x96, 0x54, 0xb8, 0xca, 0xee, 0x32,
	0x15, 0xae, 0xeb, 0xef, 0x5f, 0xa0, 0x21, 0xa0, 0x19, 0x02, 0xcd, 0xba, 0xfe, 0x18, 0x53, 0xf5,
	0xed, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x47, 0x06, 0x4b, 0x57, 0x95, 0x01, 0x00, 0x00,
}
