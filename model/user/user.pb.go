// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user/user.proto

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	user/user.proto

It has these top-level messages:
	User
	LoginRequest
	OtpRequest
	SetPasswordRequest
	ReponseMessage
	EmailRequest
	EmailVerifyRequest
	PasswordRequest
	ResetPasswordRequest
*/
package user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type User struct {
	Id        int32                      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	FirstName string                     `protobuf:"bytes,2,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName  string                     `protobuf:"bytes,3,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	Email     string                     `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
	Phone     string                     `protobuf:"bytes,5,opt,name=phone" json:"phone,omitempty"`
	Type      int32                      `protobuf:"varint,6,opt,name=type" json:"type,omitempty"`
	Password  string                     `protobuf:"bytes,7,opt,name=password" json:"password,omitempty"`
	Otp       string                     `protobuf:"bytes,8,opt,name=otp" json:"otp,omitempty"`
	IsActive  bool                       `protobuf:"varint,9,opt,name=is_active,json=isActive" json:"is_active,omitempty"`
	IsDeleted bool                       `protobuf:"varint,10,opt,name=is_deleted,json=isDeleted" json:"is_deleted,omitempty"`
	UpdatedAt *google_protobuf.Timestamp `protobuf:"bytes,11,opt,name=UpdatedAt" json:"UpdatedAt,omitempty"`
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,12,opt,name=CreatedAt" json:"CreatedAt,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *User) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetOtp() string {
	if m != nil {
		return m.Otp
	}
	return ""
}

func (m *User) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *User) GetIsDeleted() bool {
	if m != nil {
		return m.IsDeleted
	}
	return false
}

func (m *User) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *User) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type LoginRequest struct {
	Email    string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	Phone    string `protobuf:"bytes,2,opt,name=phone" json:"phone,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LoginRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type OtpRequest struct {
	Phone string `protobuf:"bytes,1,opt,name=phone" json:"phone,omitempty"`
	Otp   string `protobuf:"bytes,2,opt,name=otp" json:"otp,omitempty"`
}

func (m *OtpRequest) Reset()                    { *m = OtpRequest{} }
func (m *OtpRequest) String() string            { return proto.CompactTextString(m) }
func (*OtpRequest) ProtoMessage()               {}
func (*OtpRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *OtpRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *OtpRequest) GetOtp() string {
	if m != nil {
		return m.Otp
	}
	return ""
}

type SetPasswordRequest struct {
	Id              int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Password        string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	ConfirmPassword string `protobuf:"bytes,3,opt,name=confirmPassword" json:"confirmPassword,omitempty"`
	NewPassword     string `protobuf:"bytes,4,opt,name=newPassword" json:"newPassword,omitempty"`
}

func (m *SetPasswordRequest) Reset()                    { *m = SetPasswordRequest{} }
func (m *SetPasswordRequest) String() string            { return proto.CompactTextString(m) }
func (*SetPasswordRequest) ProtoMessage()               {}
func (*SetPasswordRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SetPasswordRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SetPasswordRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *SetPasswordRequest) GetConfirmPassword() string {
	if m != nil {
		return m.ConfirmPassword
	}
	return ""
}

func (m *SetPasswordRequest) GetNewPassword() string {
	if m != nil {
		return m.NewPassword
	}
	return ""
}

type ReponseMessage struct {
	Code    int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Status  string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	Data    *User  `protobuf:"bytes,4,opt,name=data" json:"data,omitempty"`
	Token   string `protobuf:"bytes,5,opt,name=token" json:"token,omitempty"`
}

func (m *ReponseMessage) Reset()                    { *m = ReponseMessage{} }
func (m *ReponseMessage) String() string            { return proto.CompactTextString(m) }
func (*ReponseMessage) ProtoMessage()               {}
func (*ReponseMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ReponseMessage) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ReponseMessage) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ReponseMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ReponseMessage) GetData() *User {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ReponseMessage) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type EmailRequest struct {
	Id    int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
}

func (m *EmailRequest) Reset()                    { *m = EmailRequest{} }
func (m *EmailRequest) String() string            { return proto.CompactTextString(m) }
func (*EmailRequest) ProtoMessage()               {}
func (*EmailRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *EmailRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *EmailRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type EmailVerifyRequest struct {
	Id    int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Token string `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
}

func (m *EmailVerifyRequest) Reset()                    { *m = EmailVerifyRequest{} }
func (m *EmailVerifyRequest) String() string            { return proto.CompactTextString(m) }
func (*EmailVerifyRequest) ProtoMessage()               {}
func (*EmailVerifyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *EmailVerifyRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *EmailVerifyRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *EmailVerifyRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type PasswordRequest struct {
	Email string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
}

func (m *PasswordRequest) Reset()                    { *m = PasswordRequest{} }
func (m *PasswordRequest) String() string            { return proto.CompactTextString(m) }
func (*PasswordRequest) ProtoMessage()               {}
func (*PasswordRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *PasswordRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type ResetPasswordRequest struct {
	Email           string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	Token           string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	Password        string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	ConfirmPassword string `protobuf:"bytes,4,opt,name=confirmPassword" json:"confirmPassword,omitempty"`
}

func (m *ResetPasswordRequest) Reset()                    { *m = ResetPasswordRequest{} }
func (m *ResetPasswordRequest) String() string            { return proto.CompactTextString(m) }
func (*ResetPasswordRequest) ProtoMessage()               {}
func (*ResetPasswordRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ResetPasswordRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ResetPasswordRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ResetPasswordRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *ResetPasswordRequest) GetConfirmPassword() string {
	if m != nil {
		return m.ConfirmPassword
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "user.user")
	proto.RegisterType((*LoginRequest)(nil), "user.loginRequest")
	proto.RegisterType((*OtpRequest)(nil), "user.otpRequest")
	proto.RegisterType((*SetPasswordRequest)(nil), "user.setPasswordRequest")
	proto.RegisterType((*ReponseMessage)(nil), "user.reponseMessage")
	proto.RegisterType((*EmailRequest)(nil), "user.emailRequest")
	proto.RegisterType((*EmailVerifyRequest)(nil), "user.emailVerifyRequest")
	proto.RegisterType((*PasswordRequest)(nil), "user.passwordRequest")
	proto.RegisterType((*ResetPasswordRequest)(nil), "user.resetPasswordRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for HarryUserService service

type HarryUserServiceClient interface {
	RegisterSocialUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*ReponseMessage, error)
	RegisterMobileUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*ReponseMessage, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*ReponseMessage, error)
	VerifyOTP(ctx context.Context, in *OtpRequest, opts ...grpc.CallOption) (*ReponseMessage, error)
	UpdateProfile(ctx context.Context, in *User, opts ...grpc.CallOption) (*ReponseMessage, error)
	SetPassword(ctx context.Context, in *SetPasswordRequest, opts ...grpc.CallOption) (*ReponseMessage, error)
	ChangePassword(ctx context.Context, in *SetPasswordRequest, opts ...grpc.CallOption) (*ReponseMessage, error)
	SetEmail(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*ReponseMessage, error)
	EmailVerification(ctx context.Context, in *EmailVerifyRequest, opts ...grpc.CallOption) (*ReponseMessage, error)
	ForgotPassword(ctx context.Context, in *PasswordRequest, opts ...grpc.CallOption) (*ReponseMessage, error)
	ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ReponseMessage, error)
}

type harryUserServiceClient struct {
	cc *grpc.ClientConn
}

func NewHarryUserServiceClient(cc *grpc.ClientConn) HarryUserServiceClient {
	return &harryUserServiceClient{cc}
}

func (c *harryUserServiceClient) RegisterSocialUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*ReponseMessage, error) {
	out := new(ReponseMessage)
	err := grpc.Invoke(ctx, "/user.HarryUserService/RegisterSocialUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *harryUserServiceClient) RegisterMobileUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*ReponseMessage, error) {
	out := new(ReponseMessage)
	err := grpc.Invoke(ctx, "/user.HarryUserService/RegisterMobileUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *harryUserServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*ReponseMessage, error) {
	out := new(ReponseMessage)
	err := grpc.Invoke(ctx, "/user.HarryUserService/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *harryUserServiceClient) VerifyOTP(ctx context.Context, in *OtpRequest, opts ...grpc.CallOption) (*ReponseMessage, error) {
	out := new(ReponseMessage)
	err := grpc.Invoke(ctx, "/user.HarryUserService/VerifyOTP", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *harryUserServiceClient) UpdateProfile(ctx context.Context, in *User, opts ...grpc.CallOption) (*ReponseMessage, error) {
	out := new(ReponseMessage)
	err := grpc.Invoke(ctx, "/user.HarryUserService/UpdateProfile", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *harryUserServiceClient) SetPassword(ctx context.Context, in *SetPasswordRequest, opts ...grpc.CallOption) (*ReponseMessage, error) {
	out := new(ReponseMessage)
	err := grpc.Invoke(ctx, "/user.HarryUserService/SetPassword", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *harryUserServiceClient) ChangePassword(ctx context.Context, in *SetPasswordRequest, opts ...grpc.CallOption) (*ReponseMessage, error) {
	out := new(ReponseMessage)
	err := grpc.Invoke(ctx, "/user.HarryUserService/ChangePassword", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *harryUserServiceClient) SetEmail(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*ReponseMessage, error) {
	out := new(ReponseMessage)
	err := grpc.Invoke(ctx, "/user.HarryUserService/SetEmail", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *harryUserServiceClient) EmailVerification(ctx context.Context, in *EmailVerifyRequest, opts ...grpc.CallOption) (*ReponseMessage, error) {
	out := new(ReponseMessage)
	err := grpc.Invoke(ctx, "/user.HarryUserService/EmailVerification", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *harryUserServiceClient) ForgotPassword(ctx context.Context, in *PasswordRequest, opts ...grpc.CallOption) (*ReponseMessage, error) {
	out := new(ReponseMessage)
	err := grpc.Invoke(ctx, "/user.HarryUserService/ForgotPassword", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *harryUserServiceClient) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ReponseMessage, error) {
	out := new(ReponseMessage)
	err := grpc.Invoke(ctx, "/user.HarryUserService/ResetPassword", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HarryUserService service

type HarryUserServiceServer interface {
	RegisterSocialUser(context.Context, *User) (*ReponseMessage, error)
	RegisterMobileUser(context.Context, *User) (*ReponseMessage, error)
	Login(context.Context, *LoginRequest) (*ReponseMessage, error)
	VerifyOTP(context.Context, *OtpRequest) (*ReponseMessage, error)
	UpdateProfile(context.Context, *User) (*ReponseMessage, error)
	SetPassword(context.Context, *SetPasswordRequest) (*ReponseMessage, error)
	ChangePassword(context.Context, *SetPasswordRequest) (*ReponseMessage, error)
	SetEmail(context.Context, *EmailRequest) (*ReponseMessage, error)
	EmailVerification(context.Context, *EmailVerifyRequest) (*ReponseMessage, error)
	ForgotPassword(context.Context, *PasswordRequest) (*ReponseMessage, error)
	ResetPassword(context.Context, *ResetPasswordRequest) (*ReponseMessage, error)
}

func RegisterHarryUserServiceServer(s *grpc.Server, srv HarryUserServiceServer) {
	s.RegisterService(&_HarryUserService_serviceDesc, srv)
}

func _HarryUserService_RegisterSocialUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarryUserServiceServer).RegisterSocialUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.HarryUserService/RegisterSocialUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarryUserServiceServer).RegisterSocialUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _HarryUserService_RegisterMobileUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarryUserServiceServer).RegisterMobileUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.HarryUserService/RegisterMobileUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarryUserServiceServer).RegisterMobileUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _HarryUserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarryUserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.HarryUserService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarryUserServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HarryUserService_VerifyOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OtpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarryUserServiceServer).VerifyOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.HarryUserService/VerifyOTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarryUserServiceServer).VerifyOTP(ctx, req.(*OtpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HarryUserService_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarryUserServiceServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.HarryUserService/UpdateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarryUserServiceServer).UpdateProfile(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _HarryUserService_SetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarryUserServiceServer).SetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.HarryUserService/SetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarryUserServiceServer).SetPassword(ctx, req.(*SetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HarryUserService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarryUserServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.HarryUserService/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarryUserServiceServer).ChangePassword(ctx, req.(*SetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HarryUserService_SetEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarryUserServiceServer).SetEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.HarryUserService/SetEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarryUserServiceServer).SetEmail(ctx, req.(*EmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HarryUserService_EmailVerification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailVerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarryUserServiceServer).EmailVerification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.HarryUserService/EmailVerification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarryUserServiceServer).EmailVerification(ctx, req.(*EmailVerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HarryUserService_ForgotPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarryUserServiceServer).ForgotPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.HarryUserService/ForgotPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarryUserServiceServer).ForgotPassword(ctx, req.(*PasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HarryUserService_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarryUserServiceServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.HarryUserService/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarryUserServiceServer).ResetPassword(ctx, req.(*ResetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HarryUserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.HarryUserService",
	HandlerType: (*HarryUserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterSocialUser",
			Handler:    _HarryUserService_RegisterSocialUser_Handler,
		},
		{
			MethodName: "RegisterMobileUser",
			Handler:    _HarryUserService_RegisterMobileUser_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _HarryUserService_Login_Handler,
		},
		{
			MethodName: "VerifyOTP",
			Handler:    _HarryUserService_VerifyOTP_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _HarryUserService_UpdateProfile_Handler,
		},
		{
			MethodName: "SetPassword",
			Handler:    _HarryUserService_SetPassword_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _HarryUserService_ChangePassword_Handler,
		},
		{
			MethodName: "SetEmail",
			Handler:    _HarryUserService_SetEmail_Handler,
		},
		{
			MethodName: "EmailVerification",
			Handler:    _HarryUserService_EmailVerification_Handler,
		},
		{
			MethodName: "ForgotPassword",
			Handler:    _HarryUserService_ForgotPassword_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _HarryUserService_ResetPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}

func init() { proto.RegisterFile("user/user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 683 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0x96, 0x9d, 0xa4, 0x8d, 0x4f, 0xda, 0xb4, 0x77, 0xd4, 0x7b, 0x65, 0xe5, 0x0a, 0x88, 0xbc,
	0x21, 0xab, 0x54, 0xb4, 0x15, 0x62, 0x83, 0x44, 0x55, 0x8a, 0x58, 0x50, 0x88, 0x9c, 0xb6, 0xdb,
	0x6a, 0x1a, 0x9f, 0xb8, 0x23, 0x6c, 0x8f, 0x99, 0x99, 0xb4, 0xea, 0x1b, 0xb0, 0xe0, 0x0d, 0x78,
	0x49, 0x1e, 0x01, 0x79, 0xc6, 0x4e, 0xa6, 0x69, 0x0c, 0x81, 0x8d, 0x75, 0xfe, 0xbe, 0x39, 0x67,
	0xbe, 0xf3, 0x8d, 0x61, 0x67, 0x26, 0x51, 0xec, 0x17, 0x9f, 0x61, 0x2e, 0xb8, 0xe2, 0xa4, 0x59,
	0xd8, 0xbd, 0x67, 0x31, 0xe7, 0x71, 0x82, 0xfb, 0x3a, 0x76, 0x3d, 0x9b, 0xee, 0x2b, 0x96, 0xa2,
	0x54, 0x34, 0xcd, 0x4d, 0x59, 0xf0, 0xc3, 0x05, 0x5d, 0x49, 0xba, 0xe0, 0xb2, 0xc8, 0x77, 0xfa,
	0xce, 0xa0, 0x15, 0xba, 0x2c, 0x22, 0x4f, 0x00, 0xa6, 0x4c, 0x48, 0x75, 0x95, 0xd1, 0x14, 0x7d,
	0xb7, 0xef, 0x0c, 0xbc, 0xd0, 0xd3, 0x91, 0x8f, 0x34, 0x45, 0xf2, 0x3f, 0x78, 0x09, 0xad, 0xb2,
	0x0d, 0x9d, 0x6d, 0x17, 0x01, 0x9d, 0xdc, 0x83, 0x16, 0xa6, 0x94, 0x25, 0x7e, 0x53, 0x27, 0x8c,
	0x53, 0x44, 0xf3, 0x1b, 0x9e, 0xa1, 0xdf, 0x32, 0x51, 0xed, 0x10, 0x02, 0x4d, 0x75, 0x9f, 0xa3,
	0xbf, 0xa1, 0x3b, 0x6b, 0x9b, 0xf4, 0xa0, 0x9d, 0x53, 0x29, 0xef, 0xb8, 0x88, 0xfc, 0x4d, 0x73,
	0x76, 0xe5, 0x93, 0x5d, 0x68, 0x70, 0x95, 0xfb, 0x6d, 0x1d, 0x2e, 0xcc, 0x62, 0x14, 0x26, 0xaf,
	0xe8, 0x44, 0xb1, 0x5b, 0xf4, 0xbd, 0xbe, 0x33, 0x68, 0x87, 0x6d, 0x26, 0x8f, 0xb5, 0x5f, 0x5c,
	0x83, 0xc9, 0xab, 0x08, 0x13, 0x54, 0x18, 0xf9, 0xa0, 0xb3, 0x1e, 0x93, 0x6f, 0x4d, 0x80, 0xbc,
	0x02, 0xef, 0x22, 0x8f, 0xa8, 0xc2, 0xe8, 0x58, 0xf9, 0x9d, 0xbe, 0x33, 0xe8, 0x1c, 0xf4, 0x86,
	0x86, 0xb3, 0x61, 0xc5, 0xd9, 0xf0, 0xbc, 0xe2, 0x2c, 0x5c, 0x14, 0x17, 0xc8, 0x13, 0x81, 0x25,
	0x72, 0xeb, 0xf7, 0xc8, 0x79, 0x71, 0x70, 0x09, 0x5b, 0x09, 0x8f, 0x59, 0x16, 0xe2, 0x97, 0x19,
	0x4a, 0xb5, 0x60, 0xcb, 0x59, 0xc9, 0x96, 0x6b, 0xb3, 0x65, 0x33, 0xd3, 0x78, 0xc8, 0x4c, 0x70,
	0x04, 0xc0, 0x55, 0x6e, 0x9d, 0x6a, 0xf0, 0x8e, 0x8d, 0x2f, 0xd9, 0x73, 0xe7, 0xec, 0x05, 0xdf,
	0x1c, 0x20, 0x12, 0xd5, 0xa8, 0x3c, 0xa5, 0x82, 0x2f, 0xcb, 0xc1, 0x6e, 0xec, 0x2e, 0xad, 0x64,
	0x00, 0x3b, 0x13, 0x9e, 0x4d, 0x99, 0x48, 0x47, 0x0f, 0x67, 0x5b, 0x0e, 0x93, 0x3e, 0x74, 0x32,
	0xbc, 0x9b, 0x57, 0x19, 0x79, 0xd8, 0xa1, 0x62, 0x9c, 0xae, 0xc0, 0x9c, 0x67, 0x12, 0xcf, 0x50,
	0x4a, 0x1a, 0x6b, 0x85, 0x4c, 0x78, 0x84, 0xe5, 0x30, 0xda, 0x26, 0xff, 0xc1, 0x86, 0x54, 0x54,
	0xcd, 0x64, 0x39, 0x4c, 0xe9, 0x11, 0x1f, 0x36, 0x53, 0x03, 0x2b, 0x47, 0xa8, 0x5c, 0xf2, 0x14,
	0x9a, 0x11, 0x55, 0x54, 0xf7, 0xec, 0x1c, 0xc0, 0x50, 0x3f, 0x95, 0xe2, 0x13, 0xea, 0x78, 0xc1,
	0x97, 0xe2, 0x9f, 0x31, 0xab, 0xd4, 0xa9, 0x9d, 0xe0, 0x08, 0xb6, 0xf4, 0x3a, 0xea, 0x68, 0x99,
	0xef, 0xce, 0xb5, 0x76, 0x17, 0x8c, 0x80, 0x68, 0xe3, 0x12, 0x05, 0x9b, 0xde, 0xff, 0x11, 0x76,
	0x31, 0x47, 0xc3, 0x9e, 0xe3, 0x39, 0xec, 0xe4, 0x4b, 0x1b, 0x5a, 0x29, 0x9b, 0xe0, 0xab, 0x03,
	0x7b, 0x02, 0x57, 0x2c, 0xb4, 0x56, 0x65, 0xa6, 0x9b, 0x6b, 0x75, 0xfb, 0x95, 0xca, 0x56, 0x2d,
	0xbb, 0xb9, 0x72, 0xd9, 0x07, 0xdf, 0x5b, 0xb0, 0xfb, 0x9e, 0x0a, 0x71, 0x7f, 0x21, 0x51, 0x8c,
	0x51, 0xdc, 0xb2, 0x09, 0x92, 0x97, 0x40, 0x42, 0x8c, 0x99, 0x54, 0x28, 0xc6, 0x7c, 0xc2, 0x68,
	0x52, 0x24, 0x89, 0xb5, 0x8e, 0xde, 0x9e, 0xb1, 0x97, 0x44, 0x60, 0xe1, 0xce, 0xf8, 0x35, 0x4b,
	0x70, 0x4d, 0xdc, 0x0b, 0x68, 0x7d, 0x28, 0x1e, 0x1b, 0x21, 0x26, 0x6d, 0xbf, 0xbc, 0x1a, 0xc8,
	0x21, 0x78, 0x66, 0x71, 0x9f, 0xce, 0x47, 0x64, 0xd7, 0x94, 0x2c, 0x1e, 0x56, 0x6d, 0x9f, 0x6d,
	0xf3, 0x6f, 0x18, 0x09, 0x3e, 0x65, 0x09, 0xae, 0x31, 0xda, 0x6b, 0xe8, 0x8c, 0x17, 0x7b, 0x22,
	0xbe, 0x29, 0x7a, 0xbc, 0xba, 0x1a, 0xf8, 0x1b, 0xe8, 0x9e, 0xdc, 0xd0, 0x2c, 0xc6, 0xbf, 0x3e,
	0xe1, 0x08, 0xda, 0x63, 0x54, 0xa7, 0x5a, 0x08, 0x25, 0x3d, 0xb6, 0xd8, 0x6b, 0x50, 0x27, 0xf0,
	0xcf, 0xe9, 0x5c, 0xdc, 0x6c, 0x42, 0x15, 0xe3, 0x59, 0xd5, 0xfa, 0xb1, 0xea, 0x6b, 0xef, 0xde,
	0x7d, 0xc7, 0x45, 0xcc, 0x17, 0xd7, 0xff, 0xd7, 0xd4, 0xe5, 0x6b, 0x4d, 0x7e, 0x0c, 0xdb, 0xa1,
	0x2d, 0x72, 0xd2, 0xab, 0xca, 0xd6, 0xbd, 0xfc, 0xf5, 0x86, 0xfe, 0x49, 0x1f, 0xfe, 0x0c, 0x00,
	0x00, 0xff, 0xff, 0xf6, 0x90, 0x9c, 0x04, 0x39, 0x07, 0x00, 0x00,
}
