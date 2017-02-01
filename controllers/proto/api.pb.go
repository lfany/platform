// Code generated by protoc-gen-go.
// source: controllers/proto/api.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	controllers/proto/api.proto

It has these top-level messages:
	ContactResponse
	AuthorizeContactRequest
	ConnectServiceAndMirror
	MirrorFindRequest
	ListRequest
	MirrorGetResponse
	ServiceGetRequest
	ServiceGetResponse
	Contact
	Service
	Mirror
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gengo/grpc-gateway/third_party/googleapis/google/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type ContactResponse struct {
	Token   string   `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Contact *Contact `protobuf:"bytes,2,opt,name=contact" json:"contact,omitempty"`
}

func (m *ContactResponse) Reset()                    { *m = ContactResponse{} }
func (m *ContactResponse) String() string            { return proto1.CompactTextString(m) }
func (*ContactResponse) ProtoMessage()               {}
func (*ContactResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ContactResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ContactResponse) GetContact() *Contact {
	if m != nil {
		return m.Contact
	}
	return nil
}

type AuthorizeContactRequest struct {
	Email    string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *AuthorizeContactRequest) Reset()                    { *m = AuthorizeContactRequest{} }
func (m *AuthorizeContactRequest) String() string            { return proto1.CompactTextString(m) }
func (*AuthorizeContactRequest) ProtoMessage()               {}
func (*AuthorizeContactRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AuthorizeContactRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AuthorizeContactRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type ConnectServiceAndMirror struct {
	MirrorId   int32 `protobuf:"varint,1,opt,name=mirror_id,json=mirrorId" json:"mirror_id,omitempty"`
	EndpointId int32 `protobuf:"varint,2,opt,name=endpoint_id,json=endpointId" json:"endpoint_id,omitempty"`
}

func (m *ConnectServiceAndMirror) Reset()                    { *m = ConnectServiceAndMirror{} }
func (m *ConnectServiceAndMirror) String() string            { return proto1.CompactTextString(m) }
func (*ConnectServiceAndMirror) ProtoMessage()               {}
func (*ConnectServiceAndMirror) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ConnectServiceAndMirror) GetMirrorId() int32 {
	if m != nil {
		return m.MirrorId
	}
	return 0
}

func (m *ConnectServiceAndMirror) GetEndpointId() int32 {
	if m != nil {
		return m.EndpointId
	}
	return 0
}

type MirrorFindRequest struct {
	ClientToken string `protobuf:"bytes,1,opt,name=client_token,json=clientToken" json:"client_token,omitempty"`
}

func (m *MirrorFindRequest) Reset()                    { *m = MirrorFindRequest{} }
func (m *MirrorFindRequest) String() string            { return proto1.CompactTextString(m) }
func (*MirrorFindRequest) ProtoMessage()               {}
func (*MirrorFindRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MirrorFindRequest) GetClientToken() string {
	if m != nil {
		return m.ClientToken
	}
	return ""
}

type ListRequest struct {
	Limit  int32 `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
	Offset int32 `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto1.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ListRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type MirrorGetResponse struct {
	Mirrors []*Mirror `protobuf:"bytes,1,rep,name=mirrors" json:"mirrors,omitempty"`
}

func (m *MirrorGetResponse) Reset()                    { *m = MirrorGetResponse{} }
func (m *MirrorGetResponse) String() string            { return proto1.CompactTextString(m) }
func (*MirrorGetResponse) ProtoMessage()               {}
func (*MirrorGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *MirrorGetResponse) GetMirrors() []*Mirror {
	if m != nil {
		return m.Mirrors
	}
	return nil
}

type ServiceGetRequest struct {
	Limit  int32 `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
	Offset int32 `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
}

func (m *ServiceGetRequest) Reset()                    { *m = ServiceGetRequest{} }
func (m *ServiceGetRequest) String() string            { return proto1.CompactTextString(m) }
func (*ServiceGetRequest) ProtoMessage()               {}
func (*ServiceGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ServiceGetRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ServiceGetRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ServiceGetResponse struct {
	Services []*Service `protobuf:"bytes,1,rep,name=services" json:"services,omitempty"`
}

func (m *ServiceGetResponse) Reset()                    { *m = ServiceGetResponse{} }
func (m *ServiceGetResponse) String() string            { return proto1.CompactTextString(m) }
func (*ServiceGetResponse) ProtoMessage()               {}
func (*ServiceGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ServiceGetResponse) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

type Contact struct {
	Id       int32     `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name     string    `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Email    string    `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Password string    `protobuf:"bytes,4,opt,name=password" json:"password,omitempty"`
	Admin    bool      `protobuf:"varint,5,opt,name=admin" json:"admin,omitempty"`
	Mirrors  []*Mirror `protobuf:"bytes,6,rep,name=mirrors" json:"mirrors,omitempty"`
}

func (m *Contact) Reset()                    { *m = Contact{} }
func (m *Contact) String() string            { return proto1.CompactTextString(m) }
func (*Contact) ProtoMessage()               {}
func (*Contact) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Contact) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Contact) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Contact) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Contact) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Contact) GetAdmin() bool {
	if m != nil {
		return m.Admin
	}
	return false
}

func (m *Contact) GetMirrors() []*Mirror {
	if m != nil {
		return m.Mirrors
	}
	return nil
}

type Service struct {
	Id                 int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name               string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Storage            int64  `protobuf:"varint,3,opt,name=storage" json:"storage,omitempty"`
	TrafficConsumption int64  `protobuf:"varint,4,opt,name=traffic_consumption,json=trafficConsumption" json:"traffic_consumption,omitempty"`
}

func (m *Service) Reset()                    { *m = Service{} }
func (m *Service) String() string            { return proto1.CompactTextString(m) }
func (*Service) ProtoMessage()               {}
func (*Service) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Service) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetStorage() int64 {
	if m != nil {
		return m.Storage
	}
	return 0
}

func (m *Service) GetTrafficConsumption() int64 {
	if m != nil {
		return m.TrafficConsumption
	}
	return 0
}

type Mirror struct {
	Id               int32    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	ContactId        int32    `protobuf:"varint,2,opt,name=contact_id,json=contactId" json:"contact_id,omitempty"`
	Ipv4             string   `protobuf:"bytes,3,opt,name=ipv4" json:"ipv4,omitempty"`
	Ipv6             string   `protobuf:"bytes,4,opt,name=ipv6" json:"ipv6,omitempty"`
	Name             string   `protobuf:"bytes,5,opt,name=name" json:"name,omitempty"`
	Domain           string   `protobuf:"bytes,6,opt,name=domain" json:"domain,omitempty"`
	ClientToken      string   `protobuf:"bytes,7,opt,name=client_token,json=clientToken" json:"client_token,omitempty"`
	CreatedAt        int64    `protobuf:"varint,8,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	Bandwidth        int64    `protobuf:"varint,9,opt,name=bandwidth" json:"bandwidth,omitempty"`
	Traffic          int64    `protobuf:"varint,10,opt,name=traffic" json:"traffic,omitempty"`
	AvailableStorage int64    `protobuf:"varint,11,opt,name=available_storage,json=availableStorage" json:"available_storage,omitempty"`
	Service          *Service `protobuf:"bytes,12,opt,name=service" json:"service,omitempty"`
	ServiceEndpoint  *Mirror  `protobuf:"bytes,13,opt,name=service_endpoint,json=serviceEndpoint" json:"service_endpoint,omitempty"`
	LocalDestination string   `protobuf:"bytes,14,opt,name=local_destination,json=localDestination" json:"local_destination,omitempty"`
	Storage          int64    `protobuf:"varint,15,opt,name=storage" json:"storage,omitempty"`
}

func (m *Mirror) Reset()                    { *m = Mirror{} }
func (m *Mirror) String() string            { return proto1.CompactTextString(m) }
func (*Mirror) ProtoMessage()               {}
func (*Mirror) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Mirror) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Mirror) GetContactId() int32 {
	if m != nil {
		return m.ContactId
	}
	return 0
}

func (m *Mirror) GetIpv4() string {
	if m != nil {
		return m.Ipv4
	}
	return ""
}

func (m *Mirror) GetIpv6() string {
	if m != nil {
		return m.Ipv6
	}
	return ""
}

func (m *Mirror) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Mirror) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Mirror) GetClientToken() string {
	if m != nil {
		return m.ClientToken
	}
	return ""
}

func (m *Mirror) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Mirror) GetBandwidth() int64 {
	if m != nil {
		return m.Bandwidth
	}
	return 0
}

func (m *Mirror) GetTraffic() int64 {
	if m != nil {
		return m.Traffic
	}
	return 0
}

func (m *Mirror) GetAvailableStorage() int64 {
	if m != nil {
		return m.AvailableStorage
	}
	return 0
}

func (m *Mirror) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *Mirror) GetServiceEndpoint() *Mirror {
	if m != nil {
		return m.ServiceEndpoint
	}
	return nil
}

func (m *Mirror) GetLocalDestination() string {
	if m != nil {
		return m.LocalDestination
	}
	return ""
}

func (m *Mirror) GetStorage() int64 {
	if m != nil {
		return m.Storage
	}
	return 0
}

func init() {
	proto1.RegisterType((*ContactResponse)(nil), "proto.ContactResponse")
	proto1.RegisterType((*AuthorizeContactRequest)(nil), "proto.AuthorizeContactRequest")
	proto1.RegisterType((*ConnectServiceAndMirror)(nil), "proto.ConnectServiceAndMirror")
	proto1.RegisterType((*MirrorFindRequest)(nil), "proto.MirrorFindRequest")
	proto1.RegisterType((*ListRequest)(nil), "proto.ListRequest")
	proto1.RegisterType((*MirrorGetResponse)(nil), "proto.MirrorGetResponse")
	proto1.RegisterType((*ServiceGetRequest)(nil), "proto.ServiceGetRequest")
	proto1.RegisterType((*ServiceGetResponse)(nil), "proto.ServiceGetResponse")
	proto1.RegisterType((*Contact)(nil), "proto.Contact")
	proto1.RegisterType((*Service)(nil), "proto.Service")
	proto1.RegisterType((*Mirror)(nil), "proto.Mirror")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ContactService service

type ContactServiceClient interface {
	Create(ctx context.Context, in *Contact, opts ...grpc.CallOption) (*ContactResponse, error)
	Update(ctx context.Context, in *Contact, opts ...grpc.CallOption) (*ContactResponse, error)
	Get(ctx context.Context, in *Contact, opts ...grpc.CallOption) (*ContactResponse, error)
	Authorize(ctx context.Context, in *AuthorizeContactRequest, opts ...grpc.CallOption) (*ContactResponse, error)
}

type contactServiceClient struct {
	cc *grpc.ClientConn
}

func NewContactServiceClient(cc *grpc.ClientConn) ContactServiceClient {
	return &contactServiceClient{cc}
}

func (c *contactServiceClient) Create(ctx context.Context, in *Contact, opts ...grpc.CallOption) (*ContactResponse, error) {
	out := new(ContactResponse)
	err := grpc.Invoke(ctx, "/proto.ContactService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) Update(ctx context.Context, in *Contact, opts ...grpc.CallOption) (*ContactResponse, error) {
	out := new(ContactResponse)
	err := grpc.Invoke(ctx, "/proto.ContactService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) Get(ctx context.Context, in *Contact, opts ...grpc.CallOption) (*ContactResponse, error) {
	out := new(ContactResponse)
	err := grpc.Invoke(ctx, "/proto.ContactService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) Authorize(ctx context.Context, in *AuthorizeContactRequest, opts ...grpc.CallOption) (*ContactResponse, error) {
	out := new(ContactResponse)
	err := grpc.Invoke(ctx, "/proto.ContactService/Authorize", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ContactService service

type ContactServiceServer interface {
	Create(context.Context, *Contact) (*ContactResponse, error)
	Update(context.Context, *Contact) (*ContactResponse, error)
	Get(context.Context, *Contact) (*ContactResponse, error)
	Authorize(context.Context, *AuthorizeContactRequest) (*ContactResponse, error)
}

func RegisterContactServiceServer(s *grpc.Server, srv ContactServiceServer) {
	s.RegisterService(&_ContactService_serviceDesc, srv)
}

func _ContactService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Contact)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ContactService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).Create(ctx, req.(*Contact))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Contact)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ContactService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).Update(ctx, req.(*Contact))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Contact)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ContactService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).Get(ctx, req.(*Contact))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_Authorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).Authorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ContactService/Authorize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).Authorize(ctx, req.(*AuthorizeContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContactService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ContactService",
	HandlerType: (*ContactServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ContactService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ContactService_Update_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ContactService_Get_Handler,
		},
		{
			MethodName: "Authorize",
			Handler:    _ContactService_Authorize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "controllers/proto/api.proto",
}

// Client API for MirrorService service

type MirrorServiceClient interface {
	Get(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*MirrorGetResponse, error)
	Find(ctx context.Context, in *MirrorFindRequest, opts ...grpc.CallOption) (*Mirror, error)
	FindById(ctx context.Context, in *Mirror, opts ...grpc.CallOption) (*Mirror, error)
	UpdateById(ctx context.Context, in *Mirror, opts ...grpc.CallOption) (*Mirror, error)
	Connect(ctx context.Context, in *ConnectServiceAndMirror, opts ...grpc.CallOption) (*Mirror, error)
	Create(ctx context.Context, in *Mirror, opts ...grpc.CallOption) (*Mirror, error)
}

type mirrorServiceClient struct {
	cc *grpc.ClientConn
}

func NewMirrorServiceClient(cc *grpc.ClientConn) MirrorServiceClient {
	return &mirrorServiceClient{cc}
}

func (c *mirrorServiceClient) Get(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*MirrorGetResponse, error) {
	out := new(MirrorGetResponse)
	err := grpc.Invoke(ctx, "/proto.MirrorService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mirrorServiceClient) Find(ctx context.Context, in *MirrorFindRequest, opts ...grpc.CallOption) (*Mirror, error) {
	out := new(Mirror)
	err := grpc.Invoke(ctx, "/proto.MirrorService/Find", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mirrorServiceClient) FindById(ctx context.Context, in *Mirror, opts ...grpc.CallOption) (*Mirror, error) {
	out := new(Mirror)
	err := grpc.Invoke(ctx, "/proto.MirrorService/FindById", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mirrorServiceClient) UpdateById(ctx context.Context, in *Mirror, opts ...grpc.CallOption) (*Mirror, error) {
	out := new(Mirror)
	err := grpc.Invoke(ctx, "/proto.MirrorService/UpdateById", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mirrorServiceClient) Connect(ctx context.Context, in *ConnectServiceAndMirror, opts ...grpc.CallOption) (*Mirror, error) {
	out := new(Mirror)
	err := grpc.Invoke(ctx, "/proto.MirrorService/Connect", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mirrorServiceClient) Create(ctx context.Context, in *Mirror, opts ...grpc.CallOption) (*Mirror, error) {
	out := new(Mirror)
	err := grpc.Invoke(ctx, "/proto.MirrorService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MirrorService service

type MirrorServiceServer interface {
	Get(context.Context, *ListRequest) (*MirrorGetResponse, error)
	Find(context.Context, *MirrorFindRequest) (*Mirror, error)
	FindById(context.Context, *Mirror) (*Mirror, error)
	UpdateById(context.Context, *Mirror) (*Mirror, error)
	Connect(context.Context, *ConnectServiceAndMirror) (*Mirror, error)
	Create(context.Context, *Mirror) (*Mirror, error)
}

func RegisterMirrorServiceServer(s *grpc.Server, srv MirrorServiceServer) {
	s.RegisterService(&_MirrorService_serviceDesc, srv)
}

func _MirrorService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MirrorServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MirrorService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MirrorServiceServer).Get(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MirrorService_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MirrorFindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MirrorServiceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MirrorService/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MirrorServiceServer).Find(ctx, req.(*MirrorFindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MirrorService_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Mirror)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MirrorServiceServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MirrorService/FindById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MirrorServiceServer).FindById(ctx, req.(*Mirror))
	}
	return interceptor(ctx, in, info, handler)
}

func _MirrorService_UpdateById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Mirror)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MirrorServiceServer).UpdateById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MirrorService/UpdateById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MirrorServiceServer).UpdateById(ctx, req.(*Mirror))
	}
	return interceptor(ctx, in, info, handler)
}

func _MirrorService_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectServiceAndMirror)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MirrorServiceServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MirrorService/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MirrorServiceServer).Connect(ctx, req.(*ConnectServiceAndMirror))
	}
	return interceptor(ctx, in, info, handler)
}

func _MirrorService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Mirror)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MirrorServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MirrorService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MirrorServiceServer).Create(ctx, req.(*Mirror))
	}
	return interceptor(ctx, in, info, handler)
}

var _MirrorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MirrorService",
	HandlerType: (*MirrorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _MirrorService_Get_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _MirrorService_Find_Handler,
		},
		{
			MethodName: "FindById",
			Handler:    _MirrorService_FindById_Handler,
		},
		{
			MethodName: "UpdateById",
			Handler:    _MirrorService_UpdateById_Handler,
		},
		{
			MethodName: "Connect",
			Handler:    _MirrorService_Connect_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _MirrorService_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "controllers/proto/api.proto",
}

// Client API for ServiceService service

type ServiceServiceClient interface {
	Get(ctx context.Context, in *ServiceGetRequest, opts ...grpc.CallOption) (*ServiceGetResponse, error)
	Update(ctx context.Context, in *Service, opts ...grpc.CallOption) (*Service, error)
	Find(ctx context.Context, in *Service, opts ...grpc.CallOption) (*Service, error)
	Create(ctx context.Context, in *Service, opts ...grpc.CallOption) (*Service, error)
}

type serviceServiceClient struct {
	cc *grpc.ClientConn
}

func NewServiceServiceClient(cc *grpc.ClientConn) ServiceServiceClient {
	return &serviceServiceClient{cc}
}

func (c *serviceServiceClient) Get(ctx context.Context, in *ServiceGetRequest, opts ...grpc.CallOption) (*ServiceGetResponse, error) {
	out := new(ServiceGetResponse)
	err := grpc.Invoke(ctx, "/proto.ServiceService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) Update(ctx context.Context, in *Service, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := grpc.Invoke(ctx, "/proto.ServiceService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) Find(ctx context.Context, in *Service, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := grpc.Invoke(ctx, "/proto.ServiceService/Find", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) Create(ctx context.Context, in *Service, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := grpc.Invoke(ctx, "/proto.ServiceService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ServiceService service

type ServiceServiceServer interface {
	Get(context.Context, *ServiceGetRequest) (*ServiceGetResponse, error)
	Update(context.Context, *Service) (*Service, error)
	Find(context.Context, *Service) (*Service, error)
	Create(context.Context, *Service) (*Service, error)
}

func RegisterServiceServiceServer(s *grpc.Server, srv ServiceServiceServer) {
	s.RegisterService(&_ServiceService_serviceDesc, srv)
}

func _ServiceService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ServiceService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).Get(ctx, req.(*ServiceGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Service)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ServiceService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).Update(ctx, req.(*Service))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Service)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ServiceService/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).Find(ctx, req.(*Service))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Service)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ServiceService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).Create(ctx, req.(*Service))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ServiceService",
	HandlerType: (*ServiceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ServiceService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ServiceService_Update_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _ServiceService_Find_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ServiceService_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "controllers/proto/api.proto",
}

func init() { proto1.RegisterFile("controllers/proto/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 957 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xcd, 0x72, 0xe3, 0x44,
	0x10, 0x2e, 0xdb, 0xf1, 0x5f, 0x3b, 0x71, 0xec, 0xc9, 0xe2, 0x68, 0x95, 0x00, 0x41, 0x17, 0x5c,
	0xa1, 0x2a, 0x82, 0x40, 0xa5, 0xa8, 0x40, 0x51, 0x78, 0x03, 0x1b, 0x5c, 0xfc, 0x14, 0x68, 0xa1,
	0x38, 0xba, 0x26, 0xd2, 0x38, 0x99, 0x42, 0x9e, 0x11, 0xd2, 0x24, 0x0b, 0x6c, 0xe5, 0xc2, 0x2b,
	0x70, 0xe7, 0x19, 0x78, 0x17, 0x2e, 0x5c, 0x38, 0xc1, 0x83, 0x50, 0x9a, 0xe9, 0x91, 0x25, 0x9b,
	0xec, 0x06, 0x4e, 0x9a, 0xee, 0x9e, 0xfe, 0xe6, 0xd3, 0xd7, 0x3d, 0x3d, 0xb0, 0x17, 0x4a, 0xa1,
	0x52, 0x19, 0xc7, 0x2c, 0xcd, 0xfc, 0x24, 0x95, 0x4a, 0xfa, 0x34, 0xe1, 0x47, 0x7a, 0x45, 0x9a,
	0xfa, 0xe3, 0xee, 0x5f, 0x4a, 0x79, 0x19, 0xb3, 0x3c, 0xe0, 0x53, 0x21, 0xa4, 0xa2, 0x8a, 0x4b,
	0x91, 0x99, 0x4d, 0xde, 0x57, 0xb0, 0x7d, 0x26, 0x85, 0xa2, 0xa1, 0x0a, 0x58, 0x96, 0x48, 0x91,
	0x31, 0xf2, 0x00, 0x9a, 0x4a, 0x7e, 0xc7, 0x84, 0x53, 0x3b, 0xa8, 0x8d, 0xbb, 0x81, 0x31, 0xc8,
	0x18, 0xda, 0xa1, 0xd9, 0xe8, 0xd4, 0x0f, 0x6a, 0xe3, 0xde, 0x71, 0xdf, 0x20, 0x1c, 0xd9, 0x74,
	0x1b, 0xf6, 0x3e, 0x85, 0xdd, 0xc9, 0xb5, 0xba, 0x92, 0x29, 0xff, 0x89, 0x15, 0xd8, 0xdf, 0x5f,
	0xb3, 0x4c, 0xe5, 0xd0, 0x6c, 0x41, 0x79, 0x6c, 0xa1, 0xb5, 0x41, 0x5c, 0xe8, 0x24, 0x34, 0xcb,
	0x9e, 0xca, 0x34, 0xd2, 0xd8, 0xdd, 0xa0, 0xb0, 0xbd, 0x6f, 0x61, 0xf7, 0x4c, 0x0a, 0xc1, 0x42,
	0xf5, 0x84, 0xa5, 0x37, 0x3c, 0x64, 0x13, 0x11, 0x7d, 0xce, 0xd3, 0x54, 0xa6, 0x64, 0x0f, 0xba,
	0x0b, 0xbd, 0x9a, 0xf1, 0x48, 0x03, 0x36, 0x83, 0x8e, 0x71, 0x4c, 0x23, 0xf2, 0x2a, 0xf4, 0x98,
	0x88, 0x12, 0xc9, 0x85, 0xca, 0xc3, 0x75, 0x1d, 0x06, 0xeb, 0x9a, 0x46, 0xde, 0x09, 0x0c, 0x0d,
	0xce, 0x63, 0x2e, 0x22, 0xcb, 0xef, 0x35, 0xd8, 0x0c, 0x63, 0xce, 0x84, 0x9a, 0x95, 0x15, 0xe8,
	0x19, 0xdf, 0xd7, 0xb9, 0xcb, 0x7b, 0x0f, 0x7a, 0x9f, 0xf1, 0xac, 0xfc, 0x47, 0x31, 0x5f, 0x70,
	0x85, 0x04, 0x8c, 0x41, 0x46, 0xd0, 0x92, 0xf3, 0x79, 0xc6, 0x14, 0x1e, 0x8c, 0x96, 0xf7, 0xbe,
	0x3d, 0xf4, 0x9c, 0x2d, 0xf5, 0x7e, 0x1d, 0xda, 0x86, 0x76, 0xe6, 0xd4, 0x0e, 0x1a, 0xe3, 0xde,
	0xf1, 0x16, 0x2a, 0x6b, 0xb6, 0x06, 0x36, 0xea, 0x4d, 0x60, 0x88, 0x22, 0xe8, 0xf4, 0xff, 0x43,
	0xe0, 0x43, 0x20, 0x65, 0x08, 0x64, 0x70, 0x08, 0x9d, 0xcc, 0x78, 0x2d, 0x05, 0x5b, 0x5c, 0xdc,
	0x1c, 0x14, 0x71, 0xef, 0xd7, 0x1a, 0xb4, 0xb1, 0xaa, 0xa4, 0x0f, 0xf5, 0x42, 0xfa, 0x3a, 0x8f,
	0x08, 0x81, 0x0d, 0x41, 0x17, 0x0c, 0x8b, 0xa8, 0xd7, 0xcb, 0x92, 0x37, 0xee, 0x2a, 0xf9, 0x46,
	0xb5, 0xe4, 0x79, 0x06, 0x8d, 0x16, 0x5c, 0x38, 0xcd, 0x83, 0xda, 0xb8, 0x13, 0x18, 0xa3, 0xac,
	0x52, 0xeb, 0xb9, 0x2a, 0xfd, 0x00, 0x6d, 0x64, 0x7d, 0x2f, 0x7e, 0x0e, 0xb4, 0x33, 0x25, 0x53,
	0x7a, 0xc9, 0x34, 0xc3, 0x46, 0x60, 0x4d, 0xe2, 0xc3, 0x8e, 0x4a, 0xe9, 0x7c, 0xce, 0xc3, 0x59,
	0x28, 0x45, 0x76, 0xbd, 0x48, 0xf2, 0x8b, 0xa3, 0xe9, 0x36, 0x02, 0x82, 0xa1, 0xb3, 0x65, 0xc4,
	0xfb, 0xb3, 0x01, 0x2d, 0xec, 0xcd, 0xd5, 0x93, 0x5f, 0x06, 0xc0, 0xeb, 0xb1, 0xec, 0xc6, 0x2e,
	0x7a, 0xa6, 0x9a, 0x18, 0x4f, 0x6e, 0xde, 0x41, 0x8d, 0xf4, 0x1a, 0x7d, 0x27, 0x28, 0x8f, 0x5e,
	0x17, 0x3f, 0xd0, 0x2c, 0xfd, 0xc0, 0x08, 0x5a, 0x91, 0x5c, 0x50, 0x2e, 0x9c, 0x96, 0xf6, 0xa2,
	0xb5, 0xd6, 0xcb, 0xed, 0xb5, 0x5e, 0xd6, 0xac, 0x52, 0x46, 0x15, 0x8b, 0x66, 0x54, 0x39, 0x1d,
	0xfd, 0x63, 0x5d, 0xf4, 0x4c, 0x14, 0xd9, 0x87, 0xee, 0x05, 0x15, 0xd1, 0x53, 0x1e, 0xa9, 0x2b,
	0xa7, 0x6b, 0xa2, 0x85, 0x23, 0x17, 0x0e, 0x35, 0x70, 0xc0, 0x08, 0x87, 0x26, 0x79, 0x03, 0x86,
	0xf4, 0x86, 0xf2, 0x98, 0x5e, 0xc4, 0x6c, 0x66, 0xc5, 0xed, 0xe9, 0x3d, 0x83, 0x22, 0xf0, 0x04,
	0x55, 0x1e, 0x43, 0x1b, 0x7b, 0xcb, 0xd9, 0xac, 0xcc, 0x15, 0xdb, 0x7a, 0x36, 0x4c, 0xde, 0x85,
	0x01, 0x2e, 0x67, 0xf6, 0x1e, 0x3b, 0x5b, 0x3a, 0x65, 0xa5, 0x15, 0xb6, 0x71, 0xdb, 0xc7, 0xb8,
	0x2b, 0x27, 0x14, 0xcb, 0x90, 0xc6, 0xb3, 0x88, 0x65, 0x8a, 0x0b, 0x3d, 0x00, 0x9d, 0xbe, 0xd6,
	0x63, 0xa0, 0x03, 0x1f, 0x2d, 0xfd, 0xe5, 0x86, 0xd8, 0xae, 0x34, 0xc4, 0xf1, 0x1f, 0x75, 0xe8,
	0x63, 0xeb, 0xdb, 0x0e, 0x9b, 0x42, 0xeb, 0x4c, 0xeb, 0x45, 0x56, 0xc6, 0xa1, 0x3b, 0x5a, 0x19,
	0x8f, 0x78, 0xd7, 0xbc, 0xdd, 0x9f, 0x7f, 0xff, 0xfb, 0x97, 0xfa, 0xd0, 0xdb, 0xf4, 0x6f, 0xde,
	0xf2, 0xb1, 0x03, 0xb2, 0xd3, 0xda, 0x21, 0xf9, 0x02, 0x5a, 0xdf, 0x24, 0xd1, 0x7f, 0x81, 0xda,
	0xd7, 0x50, 0x23, 0x77, 0x58, 0x86, 0xf2, 0x33, 0x16, 0xcf, 0x73, 0xbc, 0x4f, 0xa0, 0x71, 0xce,
	0xd4, 0xbd, 0xc1, 0x1e, 0x6a, 0xb0, 0x1d, 0xb2, 0x0e, 0x46, 0x28, 0x74, 0x8b, 0x81, 0x4e, 0x5e,
	0xc1, 0xfc, 0x3b, 0x46, 0xfc, 0x8b, 0xc8, 0x7a, 0x55, 0x7c, 0x7a, 0xad, 0xae, 0x4e, 0x6b, 0x87,
	0xc7, 0x7f, 0x35, 0x60, 0xcb, 0x54, 0xcf, 0x2a, 0x8b, 0xf4, 0x09, 0xc2, 0x95, 0x66, 0xae, 0xeb,
	0x54, 0xca, 0x5d, 0x1a, 0x64, 0xde, 0x8e, 0x3e, 0x64, 0x8b, 0xf4, 0xf2, 0x43, 0x70, 0x20, 0x90,
	0x29, 0x6c, 0xe4, 0x33, 0x9e, 0x54, 0xd3, 0x4a, 0x63, 0xdf, 0xad, 0xf6, 0x8f, 0xe7, 0x68, 0x14,
	0x42, 0x06, 0x25, 0x14, 0xa3, 0xc4, 0x23, 0xe8, 0xe4, 0x79, 0x8f, 0x7e, 0x9c, 0x46, 0xa4, 0x9a,
	0x74, 0x2f, 0x8c, 0x67, 0x3c, 0xba, 0x25, 0xe7, 0x00, 0xa6, 0xce, 0xf7, 0x40, 0xd9, 0xd3, 0x28,
	0x2f, 0x79, 0x6b, 0x28, 0x79, 0x81, 0x13, 0x3d, 0x88, 0xf3, 0xa7, 0xb1, 0x28, 0xca, 0x1d, 0x4f,
	0xe5, 0x2a, 0xec, 0x89, 0x86, 0x7d, 0xd3, 0x3d, 0xaa, 0xc0, 0x16, 0x8f, 0xe9, 0xad, 0x6f, 0xef,
	0x97, 0xff, 0xac, 0xf4, 0x88, 0xde, 0x92, 0x0f, 0x8a, 0x6e, 0x7f, 0x3e, 0xed, 0x91, 0xc6, 0x1f,
	0x78, 0xe5, 0x32, 0xe4, 0x55, 0xfe, 0xad, 0x0e, 0x7d, 0xe4, 0x66, 0xcb, 0xfc, 0xa5, 0x29, 0xb3,
	0x53, 0xbd, 0xf4, 0xcb, 0xf7, 0xcd, 0x7d, 0xf8, 0x2f, 0x11, 0xac, 0xf6, 0x03, 0x7d, 0x4c, 0x9f,
	0xe8, 0xab, 0x64, 0x1f, 0x28, 0xf2, 0x78, 0xed, 0x1e, 0x61, 0xaa, 0xbb, 0x62, 0x57, 0x5b, 0xd2,
	0xe6, 0x17, 0xf2, 0x4e, 0xb0, 0x6d, 0x5e, 0x84, 0x52, 0xb9, 0x38, 0x15, 0x14, 0x32, 0x59, 0x9b,
	0x0e, 0x77, 0x81, 0x54, 0xa6, 0x82, 0x05, 0x39, 0xad, 0x1d, 0x5e, 0xb4, 0xf4, 0xbe, 0xb7, 0xff,
	0x09, 0x00, 0x00, 0xff, 0xff, 0x61, 0x4f, 0x91, 0xa5, 0xea, 0x09, 0x00, 0x00,
}
