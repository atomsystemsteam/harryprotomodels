// Code generated by protoc-gen-go. DO NOT EDIT.
// source: admin/admin.proto

package admin

import (
	context "context"
	fmt "fmt"
	role "github.com/atomsystemsteam/harryprotomodels/model/role"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Admin struct {
	Id                   int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string               `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string               `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email                string               `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string               `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	Password             string               `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	AdminType            string               `protobuf:"bytes,7,opt,name=admin_type,json=adminType,proto3" json:"admin_type,omitempty"`
	IsActive             bool                 `protobuf:"varint,8,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	IsDeleted            bool                 `protobuf:"varint,9,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,10,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,11,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	Modules              []*role.Role         `protobuf:"bytes,12,rep,name=modules,proto3" json:"modules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Admin) Reset()         { *m = Admin{} }
func (m *Admin) String() string { return proto.CompactTextString(m) }
func (*Admin) ProtoMessage()    {}
func (*Admin) Descriptor() ([]byte, []int) {
	return fileDescriptor_8595c8dce2486799, []int{0}
}

func (m *Admin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Admin.Unmarshal(m, b)
}
func (m *Admin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Admin.Marshal(b, m, deterministic)
}
func (m *Admin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Admin.Merge(m, src)
}
func (m *Admin) XXX_Size() int {
	return xxx_messageInfo_Admin.Size(m)
}
func (m *Admin) XXX_DiscardUnknown() {
	xxx_messageInfo_Admin.DiscardUnknown(m)
}

var xxx_messageInfo_Admin proto.InternalMessageInfo

func (m *Admin) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Admin) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Admin) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Admin) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Admin) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Admin) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Admin) GetAdminType() string {
	if m != nil {
		return m.AdminType
	}
	return ""
}

func (m *Admin) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *Admin) GetIsDeleted() bool {
	if m != nil {
		return m.IsDeleted
	}
	return false
}

func (m *Admin) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Admin) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Admin) GetModules() []*role.Role {
	if m != nil {
		return m.Modules
	}
	return nil
}

type Request struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_8595c8dce2486799, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Request) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

type LoginRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8595c8dce2486799, []int{2}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Code                 int32    `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Token                string   `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8595c8dce2486799, []int{3}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *LoginResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *LoginResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type DeleteResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Code                 int32    `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8595c8dce2486799, []int{4}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *DeleteResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*Admin)(nil), "admin.admin")
	proto.RegisterType((*Request)(nil), "admin.request")
	proto.RegisterType((*LoginRequest)(nil), "admin.loginRequest")
	proto.RegisterType((*LoginResponse)(nil), "admin.loginResponse")
	proto.RegisterType((*DeleteResponse)(nil), "admin.deleteResponse")
}

func init() { proto.RegisterFile("admin/admin.proto", fileDescriptor_8595c8dce2486799) }

var fileDescriptor_8595c8dce2486799 = []byte{
	// 540 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0xd3, 0xba, 0xb1, 0xc7, 0xa5, 0xa8, 0x4b, 0x41, 0x96, 0x11, 0x22, 0x8a, 0x38, 0xb8,
	0x07, 0x6c, 0x14, 0x0e, 0x14, 0x89, 0x03, 0x01, 0x24, 0x38, 0x20, 0x0e, 0xa6, 0x5c, 0xb8, 0x44,
	0x9b, 0xec, 0x34, 0x59, 0x75, 0xd7, 0x6b, 0xbc, 0x9b, 0xa2, 0xfc, 0x05, 0xdf, 0xc1, 0x57, 0xa2,
	0xdd, 0xb5, 0x4b, 0x52, 0x21, 0x15, 0xb8, 0x38, 0x79, 0xef, 0xcd, 0x4c, 0x46, 0xf3, 0x5e, 0x0c,
	0xc7, 0x94, 0x49, 0x5e, 0x97, 0xee, 0x59, 0x34, 0xad, 0x32, 0x8a, 0x84, 0x0e, 0x64, 0x8f, 0x97,
	0x4a, 0x2d, 0x05, 0x96, 0x8e, 0x9c, 0xaf, 0x2f, 0x4a, 0xc3, 0x25, 0x6a, 0x43, 0x65, 0xe3, 0xeb,
	0xb2, 0xbb, 0xad, 0x12, 0x58, 0xda, 0x87, 0x27, 0xc6, 0x3f, 0xf6, 0xc0, 0xf7, 0x92, 0x23, 0x18,
	0x70, 0x96, 0x06, 0xa3, 0x20, 0x0f, 0xab, 0x01, 0x67, 0xe4, 0x11, 0xc0, 0x05, 0x6f, 0xb5, 0x99,
	0xd5, 0x54, 0x62, 0x3a, 0x18, 0x05, 0x79, 0x5c, 0xc5, 0x8e, 0xf9, 0x44, 0x25, 0x92, 0x87, 0x10,
	0x0b, 0xda, 0xab, 0x7b, 0x4e, 0x8d, 0x2c, 0xe1, 0xc4, 0x13, 0x08, 0x51, 0x52, 0x2e, 0xd2, 0x7d,
	0x27, 0x78, 0x60, 0xd9, 0x66, 0xa5, 0x6a, 0x4c, 0x43, 0xcf, 0x3a, 0x40, 0x32, 0x88, 0x1a, 0xaa,
	0xf5, 0x77, 0xd5, 0xb2, 0xf4, 0xc0, 0xcf, 0xe9, 0xb1, 0xdd, 0xc1, 0x2d, 0x37, 0x33, 0x9b, 0x06,
	0xd3, 0xa1, 0xdf, 0xc1, 0x31, 0xe7, 0x9b, 0xc6, 0xed, 0xc0, 0xf5, 0x8c, 0x2e, 0x0c, 0xbf, 0xc2,
	0x34, 0x1a, 0x05, 0x79, 0x54, 0x45, 0x5c, 0x4f, 0x1d, 0xb6, 0xbd, 0x5c, 0xcf, 0x18, 0x0a, 0x34,
	0xc8, 0xd2, 0xd8, 0xa9, 0x31, 0xd7, 0xef, 0x3c, 0x41, 0xce, 0x20, 0xfe, 0xd2, 0x30, 0x6a, 0x90,
	0x4d, 0x4d, 0x0a, 0xa3, 0x20, 0x4f, 0x26, 0x59, 0xe1, 0xcf, 0x57, 0xf4, 0xe7, 0x2b, 0xce, 0xfb,
	0xf3, 0x55, 0xbf, 0x8b, 0x6d, 0xe7, 0xdb, 0x16, 0xbb, 0xce, 0xe4, 0xf6, 0xce, 0xeb, 0x62, 0xf2,
	0x04, 0x86, 0x52, 0xb1, 0xb5, 0x40, 0x9d, 0x1e, 0x8e, 0xf6, 0xf2, 0x64, 0x02, 0x85, 0xb3, 0xa2,
	0x52, 0x02, 0xab, 0x5e, 0x1a, 0x9f, 0xc1, 0xb0, 0xc5, 0x6f, 0x6b, 0xd4, 0xe6, 0x1f, 0x3d, 0x19,
	0xbf, 0x86, 0x43, 0xa1, 0x96, 0xbc, 0xae, 0xba, 0xf6, 0x6b, 0x1b, 0x82, 0x6d, 0x1b, 0xb6, 0x0f,
	0x3e, 0xd8, 0x3d, 0xf8, 0xf8, 0x12, 0xee, 0x74, 0x13, 0x74, 0xa3, 0x6a, 0x8d, 0xe4, 0x01, 0x1c,
	0x68, 0x43, 0xcd, 0x5a, 0x77, 0x33, 0x3a, 0x44, 0x08, 0xec, 0x2f, 0x14, 0xf3, 0x3b, 0x84, 0x95,
	0xfb, 0x4e, 0x52, 0x18, 0x4a, 0xd4, 0x9a, 0x2e, 0xfb, 0x40, 0xf4, 0xd0, 0x2e, 0x62, 0xd4, 0x25,
	0xd6, 0x7d, 0x1e, 0x1c, 0x18, 0xbf, 0x82, 0x23, 0x6f, 0xcf, 0xff, 0xfc, 0xda, 0xe4, 0xe7, 0x00,
	0x8e, 0x3f, 0xd0, 0xb6, 0xdd, 0x4c, 0x6d, 0x1e, 0x3e, 0x63, 0x7b, 0xc5, 0x17, 0x48, 0x72, 0x88,
	0xde, 0xa3, 0x99, 0xfa, 0x44, 0x17, 0xfe, 0x2f, 0xd2, 0x5d, 0x33, 0x3b, 0xec, 0xb0, 0xcf, 0xfb,
	0x53, 0x48, 0x6c, 0xa5, 0x10, 0x7f, 0x51, 0xfc, 0x2c, 0x20, 0xa7, 0x90, 0x78, 0x23, 0x7d, 0xf9,
	0x8e, 0x7c, 0x63, 0xf2, 0x29, 0x24, 0x3e, 0x2d, 0xb7, 0x97, 0x4e, 0x20, 0xf1, 0x81, 0xfc, 0x53,
	0xe9, 0xfd, 0x0e, 0xdd, 0x38, 0xd2, 0x04, 0xc2, 0x8f, 0xd6, 0x23, 0x72, 0xaf, 0xd3, 0xb7, 0x3d,
	0xcf, 0x4e, 0x76, 0x49, 0xdf, 0xf3, 0xe6, 0xe5, 0xd7, 0x17, 0x4b, 0x6e, 0x56, 0xeb, 0x79, 0xb1,
	0x50, 0xb2, 0xa4, 0x46, 0x49, 0xbd, 0xd1, 0x06, 0xa5, 0x36, 0x48, 0x65, 0xb9, 0xb2, 0x67, 0x74,
	0xd1, 0x95, 0x8a, 0xa1, 0xd0, 0xa5, 0xfb, 0xf0, 0x2f, 0x98, 0xf9, 0x81, 0x13, 0x9e, 0xff, 0x0a,
	0x00, 0x00, 0xff, 0xff, 0x4f, 0xd3, 0x94, 0x2c, 0x76, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HarryAdminServiceClient is the client API for HarryAdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HarryAdminServiceClient interface {
	GetAdmin(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Admin, error)
	GetAllAdmin(ctx context.Context, in *Request, opts ...grpc.CallOption) (HarryAdminService_GetAllAdminClient, error)
	CreateAdmin(ctx context.Context, in *Admin, opts ...grpc.CallOption) (*Admin, error)
	UpdateAdmin(ctx context.Context, in *Admin, opts ...grpc.CallOption) (*Admin, error)
	DeleteAdmin(ctx context.Context, in *Admin, opts ...grpc.CallOption) (*DeleteResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type harryAdminServiceClient struct {
	cc *grpc.ClientConn
}

func NewHarryAdminServiceClient(cc *grpc.ClientConn) HarryAdminServiceClient {
	return &harryAdminServiceClient{cc}
}

func (c *harryAdminServiceClient) GetAdmin(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Admin, error) {
	out := new(Admin)
	err := c.cc.Invoke(ctx, "/admin.HarryAdminService/GetAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *harryAdminServiceClient) GetAllAdmin(ctx context.Context, in *Request, opts ...grpc.CallOption) (HarryAdminService_GetAllAdminClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HarryAdminService_serviceDesc.Streams[0], "/admin.HarryAdminService/GetAllAdmin", opts...)
	if err != nil {
		return nil, err
	}
	x := &harryAdminServiceGetAllAdminClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HarryAdminService_GetAllAdminClient interface {
	Recv() (*Admin, error)
	grpc.ClientStream
}

type harryAdminServiceGetAllAdminClient struct {
	grpc.ClientStream
}

func (x *harryAdminServiceGetAllAdminClient) Recv() (*Admin, error) {
	m := new(Admin)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *harryAdminServiceClient) CreateAdmin(ctx context.Context, in *Admin, opts ...grpc.CallOption) (*Admin, error) {
	out := new(Admin)
	err := c.cc.Invoke(ctx, "/admin.HarryAdminService/CreateAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *harryAdminServiceClient) UpdateAdmin(ctx context.Context, in *Admin, opts ...grpc.CallOption) (*Admin, error) {
	out := new(Admin)
	err := c.cc.Invoke(ctx, "/admin.HarryAdminService/UpdateAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *harryAdminServiceClient) DeleteAdmin(ctx context.Context, in *Admin, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/admin.HarryAdminService/DeleteAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *harryAdminServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/admin.HarryAdminService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HarryAdminServiceServer is the server API for HarryAdminService service.
type HarryAdminServiceServer interface {
	GetAdmin(context.Context, *Request) (*Admin, error)
	GetAllAdmin(*Request, HarryAdminService_GetAllAdminServer) error
	CreateAdmin(context.Context, *Admin) (*Admin, error)
	UpdateAdmin(context.Context, *Admin) (*Admin, error)
	DeleteAdmin(context.Context, *Admin) (*DeleteResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
}

func RegisterHarryAdminServiceServer(s *grpc.Server, srv HarryAdminServiceServer) {
	s.RegisterService(&_HarryAdminService_serviceDesc, srv)
}

func _HarryAdminService_GetAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarryAdminServiceServer).GetAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.HarryAdminService/GetAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarryAdminServiceServer).GetAdmin(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _HarryAdminService_GetAllAdmin_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HarryAdminServiceServer).GetAllAdmin(m, &harryAdminServiceGetAllAdminServer{stream})
}

type HarryAdminService_GetAllAdminServer interface {
	Send(*Admin) error
	grpc.ServerStream
}

type harryAdminServiceGetAllAdminServer struct {
	grpc.ServerStream
}

func (x *harryAdminServiceGetAllAdminServer) Send(m *Admin) error {
	return x.ServerStream.SendMsg(m)
}

func _HarryAdminService_CreateAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Admin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarryAdminServiceServer).CreateAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.HarryAdminService/CreateAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarryAdminServiceServer).CreateAdmin(ctx, req.(*Admin))
	}
	return interceptor(ctx, in, info, handler)
}

func _HarryAdminService_UpdateAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Admin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarryAdminServiceServer).UpdateAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.HarryAdminService/UpdateAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarryAdminServiceServer).UpdateAdmin(ctx, req.(*Admin))
	}
	return interceptor(ctx, in, info, handler)
}

func _HarryAdminService_DeleteAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Admin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarryAdminServiceServer).DeleteAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.HarryAdminService/DeleteAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarryAdminServiceServer).DeleteAdmin(ctx, req.(*Admin))
	}
	return interceptor(ctx, in, info, handler)
}

func _HarryAdminService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarryAdminServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.HarryAdminService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarryAdminServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HarryAdminService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "admin.HarryAdminService",
	HandlerType: (*HarryAdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdmin",
			Handler:    _HarryAdminService_GetAdmin_Handler,
		},
		{
			MethodName: "CreateAdmin",
			Handler:    _HarryAdminService_CreateAdmin_Handler,
		},
		{
			MethodName: "UpdateAdmin",
			Handler:    _HarryAdminService_UpdateAdmin_Handler,
		},
		{
			MethodName: "DeleteAdmin",
			Handler:    _HarryAdminService_DeleteAdmin_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _HarryAdminService_Login_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllAdmin",
			Handler:       _HarryAdminService_GetAllAdmin_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "admin/admin.proto",
}
