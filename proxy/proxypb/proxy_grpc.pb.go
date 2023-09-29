// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proxypb/proxy.proto

package proxypb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CentrifugoProxy_Connect_FullMethodName    = "/centrifugal.centrifugo.proxy.CentrifugoProxy.CentrifugoProxy/Connect"
	CentrifugoProxy_Refresh_FullMethodName    = "/centrifugal.centrifugo.proxy.CentrifugoProxy.CentrifugoProxy/Refresh"
	CentrifugoProxy_Subscribe_FullMethodName  = "/centrifugal.centrifugo.proxy.CentrifugoProxy.CentrifugoProxy/Subscribe"
	CentrifugoProxy_Publish_FullMethodName    = "/centrifugal.centrifugo.proxy.CentrifugoProxy.CentrifugoProxy/Publish"
	CentrifugoProxy_RPC_FullMethodName        = "/centrifugal.centrifugo.proxy.CentrifugoProxy.CentrifugoProxy/RPC"
	CentrifugoProxy_SubRefresh_FullMethodName = "/centrifugal.centrifugo.proxy.CentrifugoProxy.CentrifugoProxy/SubRefresh"
)

// CentrifugoProxyClient is the client API for CentrifugoProxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CentrifugoProxyClient interface {
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error)
	Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error)
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscribeResponse, error)
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error)
	RPC(ctx context.Context, in *RPCRequest, opts ...grpc.CallOption) (*RPCResponse, error)
	SubRefresh(ctx context.Context, in *SubRefreshRequest, opts ...grpc.CallOption) (*SubRefreshResponse, error)
}

type centrifugoProxyClient struct {
	cc grpc.ClientConnInterface
}

func NewCentrifugoProxyClient(cc grpc.ClientConnInterface) CentrifugoProxyClient {
	return &centrifugoProxyClient{cc}
}

func (c *centrifugoProxyClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error) {
	out := new(ConnectResponse)
	err := c.cc.Invoke(ctx, CentrifugoProxy_Connect_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centrifugoProxyClient) Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error) {
	out := new(RefreshResponse)
	err := c.cc.Invoke(ctx, CentrifugoProxy_Refresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centrifugoProxyClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscribeResponse, error) {
	out := new(SubscribeResponse)
	err := c.cc.Invoke(ctx, CentrifugoProxy_Subscribe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centrifugoProxyClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error) {
	out := new(PublishResponse)
	err := c.cc.Invoke(ctx, CentrifugoProxy_Publish_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centrifugoProxyClient) RPC(ctx context.Context, in *RPCRequest, opts ...grpc.CallOption) (*RPCResponse, error) {
	out := new(RPCResponse)
	err := c.cc.Invoke(ctx, CentrifugoProxy_RPC_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centrifugoProxyClient) SubRefresh(ctx context.Context, in *SubRefreshRequest, opts ...grpc.CallOption) (*SubRefreshResponse, error) {
	out := new(SubRefreshResponse)
	err := c.cc.Invoke(ctx, CentrifugoProxy_SubRefresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CentrifugoProxyServer is the server API for CentrifugoProxy service.
// All implementations must embed UnimplementedCentrifugoProxyServer
// for forward compatibility
type CentrifugoProxyServer interface {
	Connect(context.Context, *ConnectRequest) (*ConnectResponse, error)
	Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error)
	Subscribe(context.Context, *SubscribeRequest) (*SubscribeResponse, error)
	Publish(context.Context, *PublishRequest) (*PublishResponse, error)
	RPC(context.Context, *RPCRequest) (*RPCResponse, error)
	SubRefresh(context.Context, *SubRefreshRequest) (*SubRefreshResponse, error)
	mustEmbedUnimplementedCentrifugoProxyServer()
}

// UnimplementedCentrifugoProxyServer must be embedded to have forward compatible implementations.
type UnimplementedCentrifugoProxyServer struct {
}

func (UnimplementedCentrifugoProxyServer) Connect(context.Context, *ConnectRequest) (*ConnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedCentrifugoProxyServer) Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (UnimplementedCentrifugoProxyServer) Subscribe(context.Context, *SubscribeRequest) (*SubscribeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedCentrifugoProxyServer) Publish(context.Context, *PublishRequest) (*PublishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedCentrifugoProxyServer) RPC(context.Context, *RPCRequest) (*RPCResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RPC not implemented")
}
func (UnimplementedCentrifugoProxyServer) SubRefresh(context.Context, *SubRefreshRequest) (*SubRefreshResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubRefresh not implemented")
}
func (UnimplementedCentrifugoProxyServer) mustEmbedUnimplementedCentrifugoProxyServer() {}

// UnsafeCentrifugoProxyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CentrifugoProxyServer will
// result in compilation errors.
type UnsafeCentrifugoProxyServer interface {
	mustEmbedUnimplementedCentrifugoProxyServer()
}

func RegisterCentrifugoProxyServer(s grpc.ServiceRegistrar, srv CentrifugoProxyServer) {
	s.RegisterService(&CentrifugoProxy_ServiceDesc, srv)
}

func _CentrifugoProxy_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentrifugoProxyServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CentrifugoProxy_Connect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentrifugoProxyServer).Connect(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CentrifugoProxy_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentrifugoProxyServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CentrifugoProxy_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentrifugoProxyServer).Refresh(ctx, req.(*RefreshRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CentrifugoProxy_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentrifugoProxyServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CentrifugoProxy_Subscribe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentrifugoProxyServer).Subscribe(ctx, req.(*SubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CentrifugoProxy_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentrifugoProxyServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CentrifugoProxy_Publish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentrifugoProxyServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CentrifugoProxy_RPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentrifugoProxyServer).RPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CentrifugoProxy_RPC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentrifugoProxyServer).RPC(ctx, req.(*RPCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CentrifugoProxy_SubRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubRefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentrifugoProxyServer).SubRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CentrifugoProxy_SubRefresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentrifugoProxyServer).SubRefresh(ctx, req.(*SubRefreshRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CentrifugoProxy_ServiceDesc is the grpc.ServiceDesc for CentrifugoProxy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CentrifugoProxy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "centrifugal.centrifugo.proxy.CentrifugoProxy.CentrifugoProxy",
	HandlerType: (*CentrifugoProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _CentrifugoProxy_Connect_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _CentrifugoProxy_Refresh_Handler,
		},
		{
			MethodName: "Subscribe",
			Handler:    _CentrifugoProxy_Subscribe_Handler,
		},
		{
			MethodName: "Publish",
			Handler:    _CentrifugoProxy_Publish_Handler,
		},
		{
			MethodName: "RPC",
			Handler:    _CentrifugoProxy_RPC_Handler,
		},
		{
			MethodName: "SubRefresh",
			Handler:    _CentrifugoProxy_SubRefresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proxypb/proxy.proto",
}
