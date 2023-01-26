// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: discovery.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TelemetryServiceClient is the client API for TelemetryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TelemetryServiceClient interface {
	EstablishConnection(ctx context.Context, in *Connection, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Status(ctx context.Context, in *NodeStatus, opts ...grpc.CallOption) (*emptypb.Empty, error)
	FindClosestPeers(ctx context.Context, in *FindClosestPeersRequest, opts ...grpc.CallOption) (*FindClosestPeersResponse, error)
}

type telemetryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTelemetryServiceClient(cc grpc.ClientConnInterface) TelemetryServiceClient {
	return &telemetryServiceClient{cc}
}

func (c *telemetryServiceClient) EstablishConnection(ctx context.Context, in *Connection, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/odana.discovery.TelemetryService/EstablishConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telemetryServiceClient) Status(ctx context.Context, in *NodeStatus, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/odana.discovery.TelemetryService/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telemetryServiceClient) FindClosestPeers(ctx context.Context, in *FindClosestPeersRequest, opts ...grpc.CallOption) (*FindClosestPeersResponse, error) {
	out := new(FindClosestPeersResponse)
	err := c.cc.Invoke(ctx, "/odana.discovery.TelemetryService/FindClosestPeers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TelemetryServiceServer is the server API for TelemetryService service.
// All implementations must embed UnimplementedTelemetryServiceServer
// for forward compatibility
type TelemetryServiceServer interface {
	EstablishConnection(context.Context, *Connection) (*emptypb.Empty, error)
	Status(context.Context, *NodeStatus) (*emptypb.Empty, error)
	FindClosestPeers(context.Context, *FindClosestPeersRequest) (*FindClosestPeersResponse, error)
	mustEmbedUnimplementedTelemetryServiceServer()
}

// UnimplementedTelemetryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTelemetryServiceServer struct {
}

func (UnimplementedTelemetryServiceServer) EstablishConnection(context.Context, *Connection) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EstablishConnection not implemented")
}
func (UnimplementedTelemetryServiceServer) Status(context.Context, *NodeStatus) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedTelemetryServiceServer) FindClosestPeers(context.Context, *FindClosestPeersRequest) (*FindClosestPeersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindClosestPeers not implemented")
}
func (UnimplementedTelemetryServiceServer) mustEmbedUnimplementedTelemetryServiceServer() {}

// UnsafeTelemetryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TelemetryServiceServer will
// result in compilation errors.
type UnsafeTelemetryServiceServer interface {
	mustEmbedUnimplementedTelemetryServiceServer()
}

func RegisterTelemetryServiceServer(s grpc.ServiceRegistrar, srv TelemetryServiceServer) {
	s.RegisterService(&TelemetryService_ServiceDesc, srv)
}

func _TelemetryService_EstablishConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Connection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServiceServer).EstablishConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/odana.discovery.TelemetryService/EstablishConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServiceServer).EstablishConnection(ctx, req.(*Connection))
	}
	return interceptor(ctx, in, info, handler)
}

func _TelemetryService_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServiceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/odana.discovery.TelemetryService/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServiceServer).Status(ctx, req.(*NodeStatus))
	}
	return interceptor(ctx, in, info, handler)
}

func _TelemetryService_FindClosestPeers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindClosestPeersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServiceServer).FindClosestPeers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/odana.discovery.TelemetryService/FindClosestPeers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServiceServer).FindClosestPeers(ctx, req.(*FindClosestPeersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TelemetryService_ServiceDesc is the grpc.ServiceDesc for TelemetryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TelemetryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "odana.discovery.TelemetryService",
	HandlerType: (*TelemetryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EstablishConnection",
			Handler:    _TelemetryService_EstablishConnection_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _TelemetryService_Status_Handler,
		},
		{
			MethodName: "FindClosestPeers",
			Handler:    _TelemetryService_FindClosestPeers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "discovery.proto",
}
