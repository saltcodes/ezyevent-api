// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: internal/proto/location.proto

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

// LocationDataServiceClient is the client API for LocationDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LocationDataServiceClient interface {
	LocationData(ctx context.Context, in *LocationObject, opts ...grpc.CallOption) (*emptypb.Empty, error)
	FindEventsWithin(ctx context.Context, in *LocationQueryObject, opts ...grpc.CallOption) (*EventsLists, error)
}

type locationDataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLocationDataServiceClient(cc grpc.ClientConnInterface) LocationDataServiceClient {
	return &locationDataServiceClient{cc}
}

func (c *locationDataServiceClient) LocationData(ctx context.Context, in *LocationObject, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.LocationDataService/LocationData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationDataServiceClient) FindEventsWithin(ctx context.Context, in *LocationQueryObject, opts ...grpc.CallOption) (*EventsLists, error) {
	out := new(EventsLists)
	err := c.cc.Invoke(ctx, "/proto.LocationDataService/FindEventsWithin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocationDataServiceServer is the server API for LocationDataService service.
// All implementations must embed UnimplementedLocationDataServiceServer
// for forward compatibility
type LocationDataServiceServer interface {
	LocationData(context.Context, *LocationObject) (*emptypb.Empty, error)
	FindEventsWithin(context.Context, *LocationQueryObject) (*EventsLists, error)
	mustEmbedUnimplementedLocationDataServiceServer()
}

// UnimplementedLocationDataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLocationDataServiceServer struct {
}

func (UnimplementedLocationDataServiceServer) LocationData(context.Context, *LocationObject) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LocationData not implemented")
}
func (UnimplementedLocationDataServiceServer) FindEventsWithin(context.Context, *LocationQueryObject) (*EventsLists, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEventsWithin not implemented")
}
func (UnimplementedLocationDataServiceServer) mustEmbedUnimplementedLocationDataServiceServer() {}

// UnsafeLocationDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LocationDataServiceServer will
// result in compilation errors.
type UnsafeLocationDataServiceServer interface {
	mustEmbedUnimplementedLocationDataServiceServer()
}

func RegisterLocationDataServiceServer(s grpc.ServiceRegistrar, srv LocationDataServiceServer) {
	s.RegisterService(&LocationDataService_ServiceDesc, srv)
}

func _LocationDataService_LocationData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocationObject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationDataServiceServer).LocationData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LocationDataService/LocationData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationDataServiceServer).LocationData(ctx, req.(*LocationObject))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationDataService_FindEventsWithin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocationQueryObject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationDataServiceServer).FindEventsWithin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LocationDataService/FindEventsWithin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationDataServiceServer).FindEventsWithin(ctx, req.(*LocationQueryObject))
	}
	return interceptor(ctx, in, info, handler)
}

// LocationDataService_ServiceDesc is the grpc.ServiceDesc for LocationDataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LocationDataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.LocationDataService",
	HandlerType: (*LocationDataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LocationData",
			Handler:    _LocationDataService_LocationData_Handler,
		},
		{
			MethodName: "FindEventsWithin",
			Handler:    _LocationDataService_FindEventsWithin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/proto/location.proto",
}
