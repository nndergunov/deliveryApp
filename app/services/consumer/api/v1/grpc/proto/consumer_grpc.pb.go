// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: proto/consumer.proto

package consumer_proto

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

// ConsumerClient is the client API for Consumer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConsumerClient interface {
	// Sends a greeting
	InsertNewConsumer(ctx context.Context, in *NewConsumerRequest, opts ...grpc.CallOption) (*ConsumerResponse, error)
	GetAllConsumer(ctx context.Context, in *SearchParam, opts ...grpc.CallOption) (*ConsumerListResponse, error)
	DeleteConsumer(ctx context.Context, in *ConsumerID, opts ...grpc.CallOption) (*ConsumerDeleteResponse, error)
	UpdateConsumer(ctx context.Context, in *UpdateConsumerRequest, opts ...grpc.CallOption) (*ConsumerResponse, error)
	GetConsumer(ctx context.Context, in *ConsumerID, opts ...grpc.CallOption) (*ConsumerResponse, error)
	InsertNewConsumerLocation(ctx context.Context, in *Location, opts ...grpc.CallOption) (*Location, error)
	UpdateConsumerLocation(ctx context.Context, in *Location, opts ...grpc.CallOption) (*Location, error)
	GetConsumerLocation(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Location, error)
}

type consumerClient struct {
	cc grpc.ClientConnInterface
}

func NewConsumerClient(cc grpc.ClientConnInterface) ConsumerClient {
	return &consumerClient{cc}
}

func (c *consumerClient) InsertNewConsumer(ctx context.Context, in *NewConsumerRequest, opts ...grpc.CallOption) (*ConsumerResponse, error) {
	out := new(ConsumerResponse)
	err := c.cc.Invoke(ctx, "/consumer_proto.Consumer/insertNewConsumer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerClient) GetAllConsumer(ctx context.Context, in *SearchParam, opts ...grpc.CallOption) (*ConsumerListResponse, error) {
	out := new(ConsumerListResponse)
	err := c.cc.Invoke(ctx, "/consumer_proto.Consumer/getAllConsumer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerClient) DeleteConsumer(ctx context.Context, in *ConsumerID, opts ...grpc.CallOption) (*ConsumerDeleteResponse, error) {
	out := new(ConsumerDeleteResponse)
	err := c.cc.Invoke(ctx, "/consumer_proto.Consumer/deleteConsumer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerClient) UpdateConsumer(ctx context.Context, in *UpdateConsumerRequest, opts ...grpc.CallOption) (*ConsumerResponse, error) {
	out := new(ConsumerResponse)
	err := c.cc.Invoke(ctx, "/consumer_proto.Consumer/updateConsumer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerClient) GetConsumer(ctx context.Context, in *ConsumerID, opts ...grpc.CallOption) (*ConsumerResponse, error) {
	out := new(ConsumerResponse)
	err := c.cc.Invoke(ctx, "/consumer_proto.Consumer/getConsumer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerClient) InsertNewConsumerLocation(ctx context.Context, in *Location, opts ...grpc.CallOption) (*Location, error) {
	out := new(Location)
	err := c.cc.Invoke(ctx, "/consumer_proto.Consumer/insertNewConsumerLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerClient) UpdateConsumerLocation(ctx context.Context, in *Location, opts ...grpc.CallOption) (*Location, error) {
	out := new(Location)
	err := c.cc.Invoke(ctx, "/consumer_proto.Consumer/updateConsumerLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerClient) GetConsumerLocation(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Location, error) {
	out := new(Location)
	err := c.cc.Invoke(ctx, "/consumer_proto.Consumer/getConsumerLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConsumerServer is the server API for Consumer service.
// All implementations must embed UnimplementedConsumerServer
// for forward compatibility
type ConsumerServer interface {
	// Sends a greeting
	InsertNewConsumer(context.Context, *NewConsumerRequest) (*ConsumerResponse, error)
	GetAllConsumer(context.Context, *SearchParam) (*ConsumerListResponse, error)
	DeleteConsumer(context.Context, *ConsumerID) (*ConsumerDeleteResponse, error)
	UpdateConsumer(context.Context, *UpdateConsumerRequest) (*ConsumerResponse, error)
	GetConsumer(context.Context, *ConsumerID) (*ConsumerResponse, error)
	InsertNewConsumerLocation(context.Context, *Location) (*Location, error)
	UpdateConsumerLocation(context.Context, *Location) (*Location, error)
	GetConsumerLocation(context.Context, *UserID) (*Location, error)
	mustEmbedUnimplementedConsumerServer()
}

// UnimplementedConsumerServer must be embedded to have forward compatible implementations.
type UnimplementedConsumerServer struct {
}

func (UnimplementedConsumerServer) InsertNewConsumer(context.Context, *NewConsumerRequest) (*ConsumerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertNewConsumer not implemented")
}
func (UnimplementedConsumerServer) GetAllConsumer(context.Context, *SearchParam) (*ConsumerListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllConsumer not implemented")
}
func (UnimplementedConsumerServer) DeleteConsumer(context.Context, *ConsumerID) (*ConsumerDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConsumer not implemented")
}
func (UnimplementedConsumerServer) UpdateConsumer(context.Context, *UpdateConsumerRequest) (*ConsumerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConsumer not implemented")
}
func (UnimplementedConsumerServer) GetConsumer(context.Context, *ConsumerID) (*ConsumerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConsumer not implemented")
}
func (UnimplementedConsumerServer) InsertNewConsumerLocation(context.Context, *Location) (*Location, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertNewConsumerLocation not implemented")
}
func (UnimplementedConsumerServer) UpdateConsumerLocation(context.Context, *Location) (*Location, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConsumerLocation not implemented")
}
func (UnimplementedConsumerServer) GetConsumerLocation(context.Context, *UserID) (*Location, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConsumerLocation not implemented")
}
func (UnimplementedConsumerServer) mustEmbedUnimplementedConsumerServer() {}

// UnsafeConsumerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConsumerServer will
// result in compilation errors.
type UnsafeConsumerServer interface {
	mustEmbedUnimplementedConsumerServer()
}

func RegisterConsumerServer(s grpc.ServiceRegistrar, srv ConsumerServer) {
	s.RegisterService(&Consumer_ServiceDesc, srv)
}

func _Consumer_InsertNewConsumer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewConsumerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServer).InsertNewConsumer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consumer_proto.Consumer/insertNewConsumer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServer).InsertNewConsumer(ctx, req.(*NewConsumerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Consumer_GetAllConsumer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServer).GetAllConsumer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consumer_proto.Consumer/getAllConsumer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServer).GetAllConsumer(ctx, req.(*SearchParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _Consumer_DeleteConsumer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsumerID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServer).DeleteConsumer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consumer_proto.Consumer/deleteConsumer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServer).DeleteConsumer(ctx, req.(*ConsumerID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Consumer_UpdateConsumer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateConsumerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServer).UpdateConsumer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consumer_proto.Consumer/updateConsumer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServer).UpdateConsumer(ctx, req.(*UpdateConsumerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Consumer_GetConsumer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsumerID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServer).GetConsumer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consumer_proto.Consumer/getConsumer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServer).GetConsumer(ctx, req.(*ConsumerID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Consumer_InsertNewConsumerLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Location)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServer).InsertNewConsumerLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consumer_proto.Consumer/insertNewConsumerLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServer).InsertNewConsumerLocation(ctx, req.(*Location))
	}
	return interceptor(ctx, in, info, handler)
}

func _Consumer_UpdateConsumerLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Location)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServer).UpdateConsumerLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consumer_proto.Consumer/updateConsumerLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServer).UpdateConsumerLocation(ctx, req.(*Location))
	}
	return interceptor(ctx, in, info, handler)
}

func _Consumer_GetConsumerLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServer).GetConsumerLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consumer_proto.Consumer/getConsumerLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServer).GetConsumerLocation(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

// Consumer_ServiceDesc is the grpc.ServiceDesc for Consumer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Consumer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "consumer_proto.Consumer",
	HandlerType: (*ConsumerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "insertNewConsumer",
			Handler:    _Consumer_InsertNewConsumer_Handler,
		},
		{
			MethodName: "getAllConsumer",
			Handler:    _Consumer_GetAllConsumer_Handler,
		},
		{
			MethodName: "deleteConsumer",
			Handler:    _Consumer_DeleteConsumer_Handler,
		},
		{
			MethodName: "updateConsumer",
			Handler:    _Consumer_UpdateConsumer_Handler,
		},
		{
			MethodName: "getConsumer",
			Handler:    _Consumer_GetConsumer_Handler,
		},
		{
			MethodName: "insertNewConsumerLocation",
			Handler:    _Consumer_InsertNewConsumerLocation_Handler,
		},
		{
			MethodName: "updateConsumerLocation",
			Handler:    _Consumer_UpdateConsumerLocation_Handler,
		},
		{
			MethodName: "getConsumerLocation",
			Handler:    _Consumer_GetConsumerLocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/consumer.proto",
}
