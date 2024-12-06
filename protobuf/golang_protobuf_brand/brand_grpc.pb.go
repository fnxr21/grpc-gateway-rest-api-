// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0
// source: protobuf/brand.proto

package golang_protobuf_brand

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Crud_Create_FullMethodName  = "/golang_protobuf_brand.Crud/Create"
	Crud_GetList_FullMethodName = "/golang_protobuf_brand.Crud/GetList"
	Crud_GetOne_FullMethodName  = "/golang_protobuf_brand.Crud/GetOne"
	Crud_Update_FullMethodName  = "/golang_protobuf_brand.Crud/Update"
	Crud_Delete_FullMethodName  = "/golang_protobuf_brand.Crud/Delete"
)

// CrudClient is the client API for Crud service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CrudClient interface {
	Create(ctx context.Context, in *ProtoBrandRepo_ProtoBrand, opts ...grpc.CallOption) (*ProtoBrandRepo_ProtoBrand, error)
	GetList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ProtoBrandRepo_ProtoBrand], error)
	GetOne(ctx context.Context, in *wrapperspb.Int64Value, opts ...grpc.CallOption) (*ProtoBrandRepo_ProtoBrand, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*ProtoBrandRepo_ProtoBrand, error)
	Delete(ctx context.Context, in *wrapperspb.Int64Value, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error)
}

type crudClient struct {
	cc grpc.ClientConnInterface
}

func NewCrudClient(cc grpc.ClientConnInterface) CrudClient {
	return &crudClient{cc}
}

func (c *crudClient) Create(ctx context.Context, in *ProtoBrandRepo_ProtoBrand, opts ...grpc.CallOption) (*ProtoBrandRepo_ProtoBrand, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProtoBrandRepo_ProtoBrand)
	err := c.cc.Invoke(ctx, Crud_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crudClient) GetList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ProtoBrandRepo_ProtoBrand], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Crud_ServiceDesc.Streams[0], Crud_GetList_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[emptypb.Empty, ProtoBrandRepo_ProtoBrand]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Crud_GetListClient = grpc.ServerStreamingClient[ProtoBrandRepo_ProtoBrand]

func (c *crudClient) GetOne(ctx context.Context, in *wrapperspb.Int64Value, opts ...grpc.CallOption) (*ProtoBrandRepo_ProtoBrand, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProtoBrandRepo_ProtoBrand)
	err := c.cc.Invoke(ctx, Crud_GetOne_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crudClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*ProtoBrandRepo_ProtoBrand, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProtoBrandRepo_ProtoBrand)
	err := c.cc.Invoke(ctx, Crud_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crudClient) Delete(ctx context.Context, in *wrapperspb.Int64Value, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(wrapperspb.BoolValue)
	err := c.cc.Invoke(ctx, Crud_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CrudServer is the server API for Crud service.
// All implementations should embed UnimplementedCrudServer
// for forward compatibility.
type CrudServer interface {
	Create(context.Context, *ProtoBrandRepo_ProtoBrand) (*ProtoBrandRepo_ProtoBrand, error)
	GetList(*emptypb.Empty, grpc.ServerStreamingServer[ProtoBrandRepo_ProtoBrand]) error
	GetOne(context.Context, *wrapperspb.Int64Value) (*ProtoBrandRepo_ProtoBrand, error)
	Update(context.Context, *UpdateRequest) (*ProtoBrandRepo_ProtoBrand, error)
	Delete(context.Context, *wrapperspb.Int64Value) (*wrapperspb.BoolValue, error)
}

// UnimplementedCrudServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCrudServer struct{}

func (UnimplementedCrudServer) Create(context.Context, *ProtoBrandRepo_ProtoBrand) (*ProtoBrandRepo_ProtoBrand, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCrudServer) GetList(*emptypb.Empty, grpc.ServerStreamingServer[ProtoBrandRepo_ProtoBrand]) error {
	return status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedCrudServer) GetOne(context.Context, *wrapperspb.Int64Value) (*ProtoBrandRepo_ProtoBrand, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOne not implemented")
}
func (UnimplementedCrudServer) Update(context.Context, *UpdateRequest) (*ProtoBrandRepo_ProtoBrand, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCrudServer) Delete(context.Context, *wrapperspb.Int64Value) (*wrapperspb.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCrudServer) testEmbeddedByValue() {}

// UnsafeCrudServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CrudServer will
// result in compilation errors.
type UnsafeCrudServer interface {
	mustEmbedUnimplementedCrudServer()
}

func RegisterCrudServer(s grpc.ServiceRegistrar, srv CrudServer) {
	// If the following call pancis, it indicates UnimplementedCrudServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Crud_ServiceDesc, srv)
}

func _Crud_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProtoBrandRepo_ProtoBrand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Crud_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudServer).Create(ctx, req.(*ProtoBrandRepo_ProtoBrand))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crud_GetList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CrudServer).GetList(m, &grpc.GenericServerStream[emptypb.Empty, ProtoBrandRepo_ProtoBrand]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Crud_GetListServer = grpc.ServerStreamingServer[ProtoBrandRepo_ProtoBrand]

func _Crud_GetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.Int64Value)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudServer).GetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Crud_GetOne_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudServer).GetOne(ctx, req.(*wrapperspb.Int64Value))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crud_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Crud_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crud_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.Int64Value)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Crud_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudServer).Delete(ctx, req.(*wrapperspb.Int64Value))
	}
	return interceptor(ctx, in, info, handler)
}

// Crud_ServiceDesc is the grpc.ServiceDesc for Crud service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Crud_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "golang_protobuf_brand.Crud",
	HandlerType: (*CrudServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Crud_Create_Handler,
		},
		{
			MethodName: "GetOne",
			Handler:    _Crud_GetOne_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Crud_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Crud_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetList",
			Handler:       _Crud_GetList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protobuf/brand.proto",
}