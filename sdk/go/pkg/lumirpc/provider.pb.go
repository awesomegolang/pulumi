// Code generated by protoc-gen-go.
// source: provider.proto
// DO NOT EDIT!

package lumirpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
import google_protobuf "github.com/golang/protobuf/ptypes/struct"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CheckRequest struct {
	Type       string                  `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Properties *google_protobuf.Struct `protobuf:"bytes,2,opt,name=properties" json:"properties,omitempty"`
	Unknowns   map[string]bool         `protobuf:"bytes,3,rep,name=unknowns" json:"unknowns,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *CheckRequest) Reset()                    { *m = CheckRequest{} }
func (m *CheckRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckRequest) ProtoMessage()               {}
func (*CheckRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *CheckRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CheckRequest) GetProperties() *google_protobuf.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *CheckRequest) GetUnknowns() map[string]bool {
	if m != nil {
		return m.Unknowns
	}
	return nil
}

type CheckResponse struct {
	Failures []*CheckFailure `protobuf:"bytes,1,rep,name=failures" json:"failures,omitempty"`
}

func (m *CheckResponse) Reset()                    { *m = CheckResponse{} }
func (m *CheckResponse) String() string            { return proto.CompactTextString(m) }
func (*CheckResponse) ProtoMessage()               {}
func (*CheckResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *CheckResponse) GetFailures() []*CheckFailure {
	if m != nil {
		return m.Failures
	}
	return nil
}

type CheckFailure struct {
	Property string `protobuf:"bytes,1,opt,name=property" json:"property,omitempty"`
	Reason   string `protobuf:"bytes,2,opt,name=reason" json:"reason,omitempty"`
}

func (m *CheckFailure) Reset()                    { *m = CheckFailure{} }
func (m *CheckFailure) String() string            { return proto.CompactTextString(m) }
func (*CheckFailure) ProtoMessage()               {}
func (*CheckFailure) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *CheckFailure) GetProperty() string {
	if m != nil {
		return m.Property
	}
	return ""
}

func (m *CheckFailure) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type NameRequest struct {
	Type       string                  `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Properties *google_protobuf.Struct `protobuf:"bytes,2,opt,name=properties" json:"properties,omitempty"`
	Unknowns   map[string]bool         `protobuf:"bytes,3,rep,name=unknowns" json:"unknowns,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *NameRequest) Reset()                    { *m = NameRequest{} }
func (m *NameRequest) String() string            { return proto.CompactTextString(m) }
func (*NameRequest) ProtoMessage()               {}
func (*NameRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *NameRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *NameRequest) GetProperties() *google_protobuf.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *NameRequest) GetUnknowns() map[string]bool {
	if m != nil {
		return m.Unknowns
	}
	return nil
}

type NameResponse struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *NameResponse) Reset()                    { *m = NameResponse{} }
func (m *NameResponse) String() string            { return proto.CompactTextString(m) }
func (*NameResponse) ProtoMessage()               {}
func (*NameResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *NameResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateRequest struct {
	Type       string                  `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Properties *google_protobuf.Struct `protobuf:"bytes,2,opt,name=properties" json:"properties,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *CreateRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CreateRequest) GetProperties() *google_protobuf.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

type CreateResponse struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *CreateResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetRequest struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *GetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type GetResponse struct {
	Properties *google_protobuf.Struct `protobuf:"bytes,1,opt,name=properties" json:"properties,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func (m *GetResponse) GetProperties() *google_protobuf.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

type InspectChangeRequest struct {
	Id       string                  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type     string                  `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Olds     *google_protobuf.Struct `protobuf:"bytes,3,opt,name=olds" json:"olds,omitempty"`
	News     *google_protobuf.Struct `protobuf:"bytes,4,opt,name=news" json:"news,omitempty"`
	Unknowns map[string]bool         `protobuf:"bytes,5,rep,name=unknowns" json:"unknowns,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *InspectChangeRequest) Reset()                    { *m = InspectChangeRequest{} }
func (m *InspectChangeRequest) String() string            { return proto.CompactTextString(m) }
func (*InspectChangeRequest) ProtoMessage()               {}
func (*InspectChangeRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func (m *InspectChangeRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *InspectChangeRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *InspectChangeRequest) GetOlds() *google_protobuf.Struct {
	if m != nil {
		return m.Olds
	}
	return nil
}

func (m *InspectChangeRequest) GetNews() *google_protobuf.Struct {
	if m != nil {
		return m.News
	}
	return nil
}

func (m *InspectChangeRequest) GetUnknowns() map[string]bool {
	if m != nil {
		return m.Unknowns
	}
	return nil
}

type InspectChangeResponse struct {
	Replaces []string                `protobuf:"bytes,1,rep,name=replaces" json:"replaces,omitempty"`
	Changes  *google_protobuf.Struct `protobuf:"bytes,2,opt,name=changes" json:"changes,omitempty"`
}

func (m *InspectChangeResponse) Reset()                    { *m = InspectChangeResponse{} }
func (m *InspectChangeResponse) String() string            { return proto.CompactTextString(m) }
func (*InspectChangeResponse) ProtoMessage()               {}
func (*InspectChangeResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{10} }

func (m *InspectChangeResponse) GetReplaces() []string {
	if m != nil {
		return m.Replaces
	}
	return nil
}

func (m *InspectChangeResponse) GetChanges() *google_protobuf.Struct {
	if m != nil {
		return m.Changes
	}
	return nil
}

type UpdateRequest struct {
	Id   string                  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type string                  `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Olds *google_protobuf.Struct `protobuf:"bytes,3,opt,name=olds" json:"olds,omitempty"`
	News *google_protobuf.Struct `protobuf:"bytes,4,opt,name=news" json:"news,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{11} }

func (m *UpdateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *UpdateRequest) GetOlds() *google_protobuf.Struct {
	if m != nil {
		return m.Olds
	}
	return nil
}

func (m *UpdateRequest) GetNews() *google_protobuf.Struct {
	if m != nil {
		return m.News
	}
	return nil
}

type DeleteRequest struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{12} }

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeleteRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*CheckRequest)(nil), "lumirpc.CheckRequest")
	proto.RegisterType((*CheckResponse)(nil), "lumirpc.CheckResponse")
	proto.RegisterType((*CheckFailure)(nil), "lumirpc.CheckFailure")
	proto.RegisterType((*NameRequest)(nil), "lumirpc.NameRequest")
	proto.RegisterType((*NameResponse)(nil), "lumirpc.NameResponse")
	proto.RegisterType((*CreateRequest)(nil), "lumirpc.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "lumirpc.CreateResponse")
	proto.RegisterType((*GetRequest)(nil), "lumirpc.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "lumirpc.GetResponse")
	proto.RegisterType((*InspectChangeRequest)(nil), "lumirpc.InspectChangeRequest")
	proto.RegisterType((*InspectChangeResponse)(nil), "lumirpc.InspectChangeResponse")
	proto.RegisterType((*UpdateRequest)(nil), "lumirpc.UpdateRequest")
	proto.RegisterType((*DeleteRequest)(nil), "lumirpc.DeleteRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ResourceProvider service

type ResourceProviderClient interface {
	// Check validates that the given property bag is valid for a resource of the given type.
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
	// Name names a given resource.  Sometimes this will be assigned by a developer, and so the provider
	// simply fetches it from the property bag; other times, the provider will assign this based on its own algorithm.
	// In any case, resources with the same name must be safe to use interchangeably with one another.
	Name(ctx context.Context, in *NameRequest, opts ...grpc.CallOption) (*NameResponse, error)
	// Create allocates a new instance of the provided resource and returns its unique ID afterwards.  (The input ID
	// must be blank.)  If this call fails, the resource must not have been created (i.e., it is "transacational").
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Get reads the instance state identified by ID, returning a populated resource object, or nil if not found.
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// InspectChange checks what impacts a hypothetical update will have on the resource's properties.
	InspectChange(ctx context.Context, in *InspectChangeRequest, opts ...grpc.CallOption) (*InspectChangeResponse, error)
	// Update updates an existing resource with new values.
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	// Delete tears down an existing resource with the given ID.  If it fails, the resource is assumed to still exist.
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type resourceProviderClient struct {
	cc *grpc.ClientConn
}

func NewResourceProviderClient(cc *grpc.ClientConn) ResourceProviderClient {
	return &resourceProviderClient{cc}
}

func (c *resourceProviderClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := grpc.Invoke(ctx, "/lumirpc.ResourceProvider/Check", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) Name(ctx context.Context, in *NameRequest, opts ...grpc.CallOption) (*NameResponse, error) {
	out := new(NameResponse)
	err := grpc.Invoke(ctx, "/lumirpc.ResourceProvider/Name", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := grpc.Invoke(ctx, "/lumirpc.ResourceProvider/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/lumirpc.ResourceProvider/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) InspectChange(ctx context.Context, in *InspectChangeRequest, opts ...grpc.CallOption) (*InspectChangeResponse, error) {
	out := new(InspectChangeResponse)
	err := grpc.Invoke(ctx, "/lumirpc.ResourceProvider/InspectChange", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/lumirpc.ResourceProvider/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/lumirpc.ResourceProvider/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ResourceProvider service

type ResourceProviderServer interface {
	// Check validates that the given property bag is valid for a resource of the given type.
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
	// Name names a given resource.  Sometimes this will be assigned by a developer, and so the provider
	// simply fetches it from the property bag; other times, the provider will assign this based on its own algorithm.
	// In any case, resources with the same name must be safe to use interchangeably with one another.
	Name(context.Context, *NameRequest) (*NameResponse, error)
	// Create allocates a new instance of the provided resource and returns its unique ID afterwards.  (The input ID
	// must be blank.)  If this call fails, the resource must not have been created (i.e., it is "transacational").
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Get reads the instance state identified by ID, returning a populated resource object, or nil if not found.
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// InspectChange checks what impacts a hypothetical update will have on the resource's properties.
	InspectChange(context.Context, *InspectChangeRequest) (*InspectChangeResponse, error)
	// Update updates an existing resource with new values.
	Update(context.Context, *UpdateRequest) (*google_protobuf1.Empty, error)
	// Delete tears down an existing resource with the given ID.  If it fails, the resource is assumed to still exist.
	Delete(context.Context, *DeleteRequest) (*google_protobuf1.Empty, error)
}

func RegisterResourceProviderServer(s *grpc.Server, srv ResourceProviderServer) {
	s.RegisterService(&_ResourceProvider_serviceDesc, srv)
}

func _ResourceProvider_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lumirpc.ResourceProvider/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_Name_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Name(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lumirpc.ResourceProvider/Name",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Name(ctx, req.(*NameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lumirpc.ResourceProvider/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lumirpc.ResourceProvider/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_InspectChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InspectChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).InspectChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lumirpc.ResourceProvider/InspectChange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).InspectChange(ctx, req.(*InspectChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lumirpc.ResourceProvider/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lumirpc.ResourceProvider/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourceProvider_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lumirpc.ResourceProvider",
	HandlerType: (*ResourceProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _ResourceProvider_Check_Handler,
		},
		{
			MethodName: "Name",
			Handler:    _ResourceProvider_Name_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ResourceProvider_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ResourceProvider_Get_Handler,
		},
		{
			MethodName: "InspectChange",
			Handler:    _ResourceProvider_InspectChange_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ResourceProvider_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ResourceProvider_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "provider.proto",
}

func init() { proto.RegisterFile("provider.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 617 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x55, 0xcd, 0x4e, 0xdb, 0x40,
	0x10, 0xae, 0x9d, 0x10, 0xc2, 0x04, 0x47, 0x68, 0x9b, 0x84, 0xc8, 0xfd, 0x51, 0xb4, 0xbd, 0x44,
	0x42, 0x32, 0x25, 0xa8, 0x2a, 0x02, 0xa9, 0x95, 0x48, 0x21, 0xea, 0xa5, 0x42, 0xae, 0xb8, 0xf5,
	0x62, 0x9c, 0x49, 0xb0, 0xe2, 0xd8, 0xee, 0x7a, 0x0d, 0xca, 0x43, 0xf4, 0x05, 0xfa, 0x30, 0x7d,
	0x8d, 0xaa, 0x6f, 0x53, 0x79, 0xbd, 0x59, 0xec, 0x24, 0x40, 0x39, 0x44, 0xbd, 0x79, 0x67, 0xbe,
	0x99, 0xf9, 0x66, 0xf6, 0x9b, 0x35, 0xd4, 0x23, 0x16, 0xde, 0x78, 0x43, 0x64, 0x56, 0xc4, 0x42,
	0x1e, 0x92, 0x4d, 0x3f, 0x99, 0x7a, 0x2c, 0x72, 0xcd, 0x17, 0xe3, 0x30, 0x1c, 0xfb, 0xb8, 0x2f,
	0xcc, 0x57, 0xc9, 0x68, 0x1f, 0xa7, 0x11, 0x9f, 0x65, 0x28, 0xf3, 0xe5, 0xa2, 0x33, 0xe6, 0x2c,
	0x71, 0x79, 0xe6, 0xa5, 0x7f, 0x34, 0xd8, 0xee, 0x5f, 0xa3, 0x3b, 0xb1, 0xf1, 0x7b, 0x82, 0x31,
	0x27, 0x04, 0xca, 0x7c, 0x16, 0x61, 0x5b, 0xeb, 0x68, 0xdd, 0x2d, 0x5b, 0x7c, 0x93, 0xf7, 0x00,
	0x11, 0x0b, 0x23, 0x64, 0xdc, 0xc3, 0xb8, 0xad, 0x77, 0xb4, 0x6e, 0xad, 0xb7, 0x6b, 0x65, 0x79,
	0xad, 0x79, 0x5e, 0xeb, 0xab, 0xc8, 0x6b, 0xe7, 0xa0, 0xe4, 0x23, 0x54, 0x93, 0x60, 0x12, 0x84,
	0xb7, 0x41, 0xdc, 0x2e, 0x75, 0x4a, 0xdd, 0x5a, 0xef, 0x8d, 0x25, 0x49, 0x5b, 0xf9, 0xaa, 0xd6,
	0xa5, 0x44, 0x9d, 0x05, 0x9c, 0xcd, 0x6c, 0x15, 0x64, 0x9e, 0x80, 0x51, 0x70, 0x91, 0x1d, 0x28,
	0x4d, 0x70, 0x26, 0xd9, 0xa5, 0x9f, 0xa4, 0x01, 0x1b, 0x37, 0x8e, 0x9f, 0xa0, 0xe0, 0x55, 0xb5,
	0xb3, 0xc3, 0xb1, 0x7e, 0xa4, 0xd1, 0x53, 0x30, 0x64, 0x91, 0x38, 0x0a, 0x83, 0x18, 0xc9, 0x01,
	0x54, 0x47, 0x8e, 0xe7, 0x27, 0x0c, 0xe3, 0xb6, 0x26, 0xe8, 0x34, 0x8b, 0x74, 0xce, 0x33, 0xaf,
	0xad, 0x60, 0xf4, 0x54, 0x8e, 0x47, 0x7a, 0x88, 0x09, 0x55, 0xd9, 0xdf, 0x9c, 0x84, 0x3a, 0x93,
	0x16, 0x54, 0x18, 0x3a, 0x71, 0x18, 0x08, 0x2a, 0x5b, 0xb6, 0x3c, 0xd1, 0xdf, 0x1a, 0xd4, 0xbe,
	0x38, 0x53, 0x5c, 0xcb, 0x88, 0x3f, 0x2c, 0x8d, 0x98, 0xaa, 0x9e, 0x72, 0x45, 0xd7, 0x33, 0x61,
	0x0a, 0xdb, 0x59, 0x0d, 0x39, 0x60, 0x02, 0xe5, 0xc0, 0x99, 0xaa, 0xce, 0xd2, 0x6f, 0xfa, 0x0d,
	0x8c, 0x3e, 0x43, 0x87, 0xaf, 0xa5, 0x7d, 0xda, 0x81, 0xfa, 0x3c, 0xbb, 0xe4, 0x50, 0x07, 0xdd,
	0x1b, 0xca, 0xe4, 0xba, 0x37, 0xa4, 0x6f, 0x01, 0x06, 0xc8, 0xe7, 0xc5, 0x17, 0xbc, 0x8a, 0x8c,
	0x7e, 0x47, 0x86, 0x9e, 0x43, 0x4d, 0x44, 0xc8, 0x84, 0x45, 0x6e, 0xda, 0xbf, 0x73, 0xfb, 0xa9,
	0x43, 0xe3, 0x73, 0x10, 0x47, 0xe8, 0xf2, 0xfe, 0xb5, 0x13, 0x8c, 0xf1, 0x09, 0x24, 0xc8, 0x1e,
	0x94, 0x43, 0x7f, 0x98, 0xde, 0xe9, 0x83, 0xf5, 0x04, 0x28, 0x05, 0x07, 0x78, 0x1b, 0xb7, 0xcb,
	0x8f, 0x80, 0x53, 0x10, 0x19, 0xe4, 0x14, 0xb3, 0x21, 0x14, 0xb3, 0xa7, 0x14, 0xb3, 0x8a, 0xee,
	0x7a, 0xa4, 0x33, 0x82, 0xe6, 0x42, 0x31, 0x39, 0x6e, 0x13, 0xaa, 0x0c, 0x23, 0xdf, 0x71, 0xe5,
	0x92, 0x6e, 0xd9, 0xea, 0x4c, 0x0e, 0x60, 0xd3, 0x15, 0xe8, 0x47, 0x35, 0x32, 0xc7, 0xd1, 0x1f,
	0x1a, 0x18, 0x97, 0xd1, 0x30, 0xa7, 0xbf, 0xff, 0x3a, 0x7d, 0x7a, 0x08, 0xc6, 0x27, 0xf4, 0xf1,
	0x49, 0x74, 0x7a, 0xbf, 0x4a, 0xb0, 0x63, 0x63, 0x1c, 0x26, 0xcc, 0xc5, 0x0b, 0xf9, 0x13, 0x20,
	0x47, 0xb0, 0x21, 0x9e, 0x26, 0xd2, 0x5c, 0xf9, 0xa6, 0x9a, 0xad, 0x45, 0x73, 0x36, 0x60, 0xfa,
	0x8c, 0xbc, 0x83, 0x72, 0xba, 0xb6, 0xa4, 0xb1, 0xea, 0xa5, 0x30, 0x9b, 0x0b, 0x56, 0x15, 0x76,
	0x02, 0x95, 0x6c, 0xd7, 0x48, 0x2e, 0x75, 0x7e, 0xb5, 0xcd, 0xdd, 0x25, 0xbb, 0x0a, 0xee, 0x41,
	0x69, 0x80, 0x9c, 0x3c, 0x57, 0x88, 0xbb, 0xa5, 0x34, 0x1b, 0x45, 0xa3, 0x8a, 0xb9, 0x00, 0xa3,
	0xa0, 0x11, 0xf2, 0xea, 0x41, 0xa1, 0x9a, 0xaf, 0xef, 0x73, 0xab, 0x8c, 0xc7, 0x50, 0xc9, 0xc4,
	0x90, 0x6b, 0xa1, 0xa0, 0x0e, 0xb3, 0xb5, 0x74, 0x7d, 0x67, 0xe9, 0xcf, 0x34, 0x8b, 0xcd, 0x6e,
	0x2e, 0x17, 0x5b, 0xb8, 0xca, 0xfb, 0x63, 0xaf, 0x2a, 0xc2, 0x72, 0xf8, 0x37, 0x00, 0x00, 0xff,
	0xff, 0x1d, 0xb0, 0x87, 0x68, 0xc3, 0x07, 0x00, 0x00,
}