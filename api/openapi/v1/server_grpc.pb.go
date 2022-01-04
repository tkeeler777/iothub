// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	v1 "github.com/tkeel-io/tkeel-interface/openapi/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OpenapiClient is the client API for Openapi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OpenapiClient interface {
	// Query identify.
	// TKEEL_COMMENT
	// {"response":{"raw_data":true}}
	Identify(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.IdentifyResponse, error)
	// Post addons identify.
	// TKEEL_COMMENT
	// {
	//  "response" :
	//    {
	//      "raw_data": true
	//    }
	// }
	AddonsIdentify(ctx context.Context, in *v1.AddonsIdentifyRequest, opts ...grpc.CallOption) (*v1.AddonsIdentifyResponse, error)
	// Post tenant bind.
	// TKEEL_COMMENT
	// {"response":{"raw_data":true}}
	TenantBind(ctx context.Context, in *v1.TenantBindRequst, opts ...grpc.CallOption) (*v1.TenantBindResponse, error)
	// Post tenant bind.
	// TKEEL_COMMENT
	// {"response":{"raw_data":true}}
	TenantUnbind(ctx context.Context, in *v1.TenantUnbindRequst, opts ...grpc.CallOption) (*v1.TenantUnbindResponse, error)
	// Query status.
	// TKEEL_COMMENT
	// {"response":{"raw_data":true}}
	Status(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.StatusResponse, error)
}

type openapiClient struct {
	cc grpc.ClientConnInterface
}

func NewOpenapiClient(cc grpc.ClientConnInterface) OpenapiClient {
	return &openapiClient{cc}
}

func (c *openapiClient) Identify(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.IdentifyResponse, error) {
	out := new(v1.IdentifyResponse)
	err := c.cc.Invoke(ctx, "/openapi.v1.Openapi/Identify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openapiClient) AddonsIdentify(ctx context.Context, in *v1.AddonsIdentifyRequest, opts ...grpc.CallOption) (*v1.AddonsIdentifyResponse, error) {
	out := new(v1.AddonsIdentifyResponse)
	err := c.cc.Invoke(ctx, "/openapi.v1.Openapi/AddonsIdentify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openapiClient) TenantBind(ctx context.Context, in *v1.TenantBindRequst, opts ...grpc.CallOption) (*v1.TenantBindResponse, error) {
	out := new(v1.TenantBindResponse)
	err := c.cc.Invoke(ctx, "/openapi.v1.Openapi/TenantBind", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openapiClient) TenantUnbind(ctx context.Context, in *v1.TenantUnbindRequst, opts ...grpc.CallOption) (*v1.TenantUnbindResponse, error) {
	out := new(v1.TenantUnbindResponse)
	err := c.cc.Invoke(ctx, "/openapi.v1.Openapi/TenantUnbind", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openapiClient) Status(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.StatusResponse, error) {
	out := new(v1.StatusResponse)
	err := c.cc.Invoke(ctx, "/openapi.v1.Openapi/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OpenapiServer is the server API for Openapi service.
// All implementations must embed UnimplementedOpenapiServer
// for forward compatibility
type OpenapiServer interface {
	// Query identify.
	// TKEEL_COMMENT
	// {"response":{"raw_data":true}}
	Identify(context.Context, *emptypb.Empty) (*v1.IdentifyResponse, error)
	// Post addons identify.
	// TKEEL_COMMENT
	// {
	//  "response" :
	//    {
	//      "raw_data": true
	//    }
	// }
	AddonsIdentify(context.Context, *v1.AddonsIdentifyRequest) (*v1.AddonsIdentifyResponse, error)
	// Post tenant bind.
	// TKEEL_COMMENT
	// {"response":{"raw_data":true}}
	TenantBind(context.Context, *v1.TenantBindRequst) (*v1.TenantBindResponse, error)
	// Post tenant bind.
	// TKEEL_COMMENT
	// {"response":{"raw_data":true}}
	TenantUnbind(context.Context, *v1.TenantUnbindRequst) (*v1.TenantUnbindResponse, error)
	// Query status.
	// TKEEL_COMMENT
	// {"response":{"raw_data":true}}
	Status(context.Context, *emptypb.Empty) (*v1.StatusResponse, error)
	mustEmbedUnimplementedOpenapiServer()
}

// UnimplementedOpenapiServer must be embedded to have forward compatible implementations.
type UnimplementedOpenapiServer struct {
}

func (UnimplementedOpenapiServer) Identify(context.Context, *emptypb.Empty) (*v1.IdentifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Identify not implemented")
}
func (UnimplementedOpenapiServer) AddonsIdentify(context.Context, *v1.AddonsIdentifyRequest) (*v1.AddonsIdentifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddonsIdentify not implemented")
}
func (UnimplementedOpenapiServer) TenantBind(context.Context, *v1.TenantBindRequst) (*v1.TenantBindResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TenantBind not implemented")
}
func (UnimplementedOpenapiServer) TenantUnbind(context.Context, *v1.TenantUnbindRequst) (*v1.TenantUnbindResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TenantUnbind not implemented")
}
func (UnimplementedOpenapiServer) Status(context.Context, *emptypb.Empty) (*v1.StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedOpenapiServer) mustEmbedUnimplementedOpenapiServer() {}

// UnsafeOpenapiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OpenapiServer will
// result in compilation errors.
type UnsafeOpenapiServer interface {
	mustEmbedUnimplementedOpenapiServer()
}

func RegisterOpenapiServer(s grpc.ServiceRegistrar, srv OpenapiServer) {
	s.RegisterService(&Openapi_ServiceDesc, srv)
}

func _Openapi_Identify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenapiServer).Identify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openapi.v1.Openapi/Identify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenapiServer).Identify(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Openapi_AddonsIdentify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.AddonsIdentifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenapiServer).AddonsIdentify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openapi.v1.Openapi/AddonsIdentify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenapiServer).AddonsIdentify(ctx, req.(*v1.AddonsIdentifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Openapi_TenantBind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.TenantBindRequst)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenapiServer).TenantBind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openapi.v1.Openapi/TenantBind",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenapiServer).TenantBind(ctx, req.(*v1.TenantBindRequst))
	}
	return interceptor(ctx, in, info, handler)
}

func _Openapi_TenantUnbind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.TenantUnbindRequst)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenapiServer).TenantUnbind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openapi.v1.Openapi/TenantUnbind",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenapiServer).TenantUnbind(ctx, req.(*v1.TenantUnbindRequst))
	}
	return interceptor(ctx, in, info, handler)
}

func _Openapi_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenapiServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openapi.v1.Openapi/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenapiServer).Status(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Openapi_ServiceDesc is the grpc.ServiceDesc for Openapi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Openapi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openapi.v1.Openapi",
	HandlerType: (*OpenapiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Identify",
			Handler:    _Openapi_Identify_Handler,
		},
		{
			MethodName: "AddonsIdentify",
			Handler:    _Openapi_AddonsIdentify_Handler,
		},
		{
			MethodName: "TenantBind",
			Handler:    _Openapi_TenantBind_Handler,
		},
		{
			MethodName: "TenantUnbind",
			Handler:    _Openapi_TenantUnbind_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Openapi_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/openapi/v1/server.proto",
}