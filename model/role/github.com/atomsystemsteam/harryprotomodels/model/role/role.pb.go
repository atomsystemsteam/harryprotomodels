// Code generated by protoc-gen-go. DO NOT EDIT.
// source: role.proto

/*
Package role is a generated protocol buffer package.

It is generated from these files:
	role.proto

It has these top-level messages:
	Role
*/
package role

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Role struct {
	Id        int32                      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	RoleName  string                     `protobuf:"bytes,2,opt,name=role_name,json=roleName" json:"role_name,omitempty"`
	IsActive  bool                       `protobuf:"varint,3,opt,name=is_active,json=isActive" json:"is_active,omitempty"`
	IsDeleted bool                       `protobuf:"varint,4,opt,name=is_deleted,json=isDeleted" json:"is_deleted,omitempty"`
	UpdatedAt *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=UpdatedAt" json:"UpdatedAt,omitempty"`
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=CreatedAt" json:"CreatedAt,omitempty"`
}

func (m *Role) Reset()                    { *m = Role{} }
func (m *Role) String() string            { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()               {}
func (*Role) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Role) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Role) GetRoleName() string {
	if m != nil {
		return m.RoleName
	}
	return ""
}

func (m *Role) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *Role) GetIsDeleted() bool {
	if m != nil {
		return m.IsDeleted
	}
	return false
}

func (m *Role) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Role) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Role)(nil), "role.Role")
}

func init() { proto.RegisterFile("role.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0x3f, 0x4b, 0x03, 0x31,
	0x18, 0xc6, 0x49, 0xbd, 0x96, 0xbb, 0x08, 0x0e, 0x99, 0x42, 0x45, 0x3c, 0x9c, 0x6e, 0x4a, 0x40,
	0x41, 0xba, 0x56, 0x9d, 0x1d, 0x82, 0x2e, 0x2e, 0x47, 0xae, 0x79, 0xbd, 0x06, 0x12, 0x73, 0x24,
	0x6f, 0x85, 0x7e, 0x5e, 0xbf, 0x88, 0x24, 0x67, 0x75, 0x74, 0x4a, 0xf2, 0x7b, 0xfe, 0xf0, 0x84,
	0xd2, 0x18, 0x1c, 0x88, 0x29, 0x06, 0x0c, 0xac, 0xca, 0xf7, 0xf5, 0xf5, 0x18, 0xc2, 0xe8, 0x40,
	0x16, 0x36, 0x1c, 0xde, 0x25, 0x5a, 0x0f, 0x09, 0xb5, 0x9f, 0x66, 0xdb, 0xcd, 0x17, 0xa1, 0x95,
	0x0a, 0x0e, 0xd8, 0x05, 0x5d, 0x58, 0xc3, 0x49, 0x4b, 0xba, 0xa5, 0x5a, 0x58, 0xc3, 0x2e, 0x69,
	0x93, 0x1b, 0xfa, 0x0f, 0xed, 0x81, 0x2f, 0x5a, 0xd2, 0x35, 0xaa, 0xce, 0xe0, 0x59, 0x7b, 0xc8,
	0xa2, 0x4d, 0xbd, 0xde, 0xa1, 0xfd, 0x04, 0x7e, 0xd6, 0x92, 0xae, 0x56, 0xb5, 0x4d, 0xdb, 0xf2,
	0x66, 0x57, 0x94, 0xda, 0xd4, 0x1b, 0x70, 0x80, 0x60, 0x78, 0x55, 0xd4, 0xc6, 0xa6, 0xa7, 0x19,
	0xb0, 0x0d, 0x6d, 0x5e, 0x27, 0xa3, 0x11, 0xcc, 0x16, 0xf9, 0xb2, 0x25, 0xdd, 0xf9, 0xed, 0x5a,
	0xcc, 0x33, 0xc5, 0x69, 0xa6, 0x78, 0x39, 0xcd, 0x54, 0x7f, 0xe6, 0x9c, 0x7c, 0x8c, 0xf0, 0x93,
	0x5c, 0xfd, 0x9f, 0xfc, 0x35, 0x3f, 0x6c, 0xde, 0xee, 0x47, 0x8b, 0xfb, 0xc3, 0x20, 0x76, 0xc1,
	0x4b, 0x8d, 0xc1, 0xa7, 0x63, 0x42, 0xf0, 0x09, 0x41, 0x7b, 0xb9, 0xd7, 0x31, 0x1e, 0x4b, 0x81,
	0x0f, 0x06, 0x5c, 0x92, 0xe5, 0x90, 0xf9, 0xb7, 0xc3, 0xaa, 0xf0, 0xbb, 0xef, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xe7, 0x62, 0x73, 0xed, 0x5b, 0x01, 0x00, 0x00,
}
