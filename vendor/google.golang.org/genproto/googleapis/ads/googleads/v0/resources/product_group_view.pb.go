// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/product_group_view.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v0/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A product group view.
type ProductGroupView struct {
	// The resource name of the product group view.
	// Product group view resource names have the form:
	//
	// `customers/{customer_id}/productGroupViews/{ad_group_id}_{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductGroupView) Reset()         { *m = ProductGroupView{} }
func (m *ProductGroupView) String() string { return proto.CompactTextString(m) }
func (*ProductGroupView) ProtoMessage()    {}
func (*ProductGroupView) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_group_view_f4c44c1580b30694, []int{0}
}
func (m *ProductGroupView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductGroupView.Unmarshal(m, b)
}
func (m *ProductGroupView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductGroupView.Marshal(b, m, deterministic)
}
func (dst *ProductGroupView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductGroupView.Merge(dst, src)
}
func (m *ProductGroupView) XXX_Size() int {
	return xxx_messageInfo_ProductGroupView.Size(m)
}
func (m *ProductGroupView) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductGroupView.DiscardUnknown(m)
}

var xxx_messageInfo_ProductGroupView proto.InternalMessageInfo

func (m *ProductGroupView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*ProductGroupView)(nil), "google.ads.googleads.v0.resources.ProductGroupView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/product_group_view.proto", fileDescriptor_product_group_view_f4c44c1580b30694)
}

var fileDescriptor_product_group_view_f4c44c1580b30694 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4a, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xa2, 0xd4,
	0xe2, 0xfc, 0xd2, 0xa2, 0xe4, 0xd4, 0x62, 0xfd, 0x82, 0xa2, 0xfc, 0x94, 0xd2, 0xe4, 0x92, 0xf8,
	0xf4, 0xa2, 0xfc, 0xd2, 0x82, 0xf8, 0xb2, 0xcc, 0xd4, 0x72, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c,
	0x21, 0x45, 0x88, 0x06, 0xbd, 0xc4, 0x94, 0x62, 0x3d, 0xb8, 0x5e, 0xbd, 0x32, 0x03, 0x3d, 0xb8,
	0x5e, 0x25, 0x73, 0x2e, 0x81, 0x00, 0x88, 0x76, 0x77, 0x90, 0xee, 0xb0, 0xcc, 0xd4, 0x72, 0x21,
	0x65, 0x2e, 0x5e, 0x98, 0x82, 0xf8, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce,
	0x20, 0x1e, 0x98, 0xa0, 0x5f, 0x62, 0x6e, 0xaa, 0x53, 0x13, 0x13, 0x97, 0x6a, 0x72, 0x7e, 0xae,
	0x1e, 0x41, 0x2b, 0x9c, 0x44, 0xd1, 0x2d, 0x08, 0x00, 0x39, 0x2e, 0x80, 0x31, 0xca, 0x0b, 0xaa,
	0x37, 0x3d, 0x3f, 0x27, 0x31, 0x2f, 0x5d, 0x2f, 0xbf, 0x28, 0x5d, 0x3f, 0x3d, 0x35, 0x0f, 0xec,
	0x74, 0x98, 0x57, 0x0b, 0x32, 0x8b, 0xf1, 0xf8, 0xdc, 0x1a, 0xce, 0x5a, 0xc4, 0xc4, 0xec, 0xee,
	0xe8, 0xb8, 0x8a, 0x49, 0xd1, 0x1d, 0x62, 0xa4, 0x63, 0x4a, 0xb1, 0x1e, 0x84, 0x09, 0x62, 0x85,
	0x19, 0xe8, 0x05, 0xc1, 0x54, 0x9e, 0x82, 0xa9, 0x89, 0x71, 0x4c, 0x29, 0x8e, 0x81, 0xab, 0x89,
	0x09, 0x33, 0x88, 0x81, 0xab, 0x79, 0xc5, 0xa4, 0x0a, 0x91, 0xb0, 0xb2, 0x72, 0x4c, 0x29, 0xb6,
	0xb2, 0x82, 0xab, 0xb2, 0xb2, 0x0a, 0x33, 0xb0, 0xb2, 0x82, 0xab, 0x4b, 0x62, 0x03, 0x3b, 0xd6,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x23, 0xa0, 0xf1, 0xa5, 0x01, 0x00, 0x00,
}
