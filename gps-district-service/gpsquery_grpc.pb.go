// Code generated by protoc-gendemo-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gendemo-go-grpc v1.3.0
// - protoc             v4.23.3
// source: gps-district-service/gpsquery.proto

package gps_district_service

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
	GPSDistrictService_Query_FullMethodName = "/gps_district_service.GPSDistrictService/Query"
	GPSDistrictService_IsIn_FullMethodName  = "/gps_district_service.GPSDistrictService/IsIn"
)

// GPSDistrictServiceClient is the client API for GPSDistrictService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GPSDistrictServiceClient interface {
	// 返回的区划码和坐标顺序一一对应
	Query(ctx context.Context, in *CoordinateList, opts ...grpc.CallOption) (*AdcodeList, error)
	// 返回经纬度是否在某区划内 和坐标顺序一一对应
	IsIn(ctx context.Context, in *IsInQuery, opts ...grpc.CallOption) (*InBorderList, error)
}

type gPSDistrictServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGPSDistrictServiceClient(cc grpc.ClientConnInterface) GPSDistrictServiceClient {
	return &gPSDistrictServiceClient{cc}
}

func (c *gPSDistrictServiceClient) Query(ctx context.Context, in *CoordinateList, opts ...grpc.CallOption) (*AdcodeList, error) {
	out := new(AdcodeList)
	err := c.cc.Invoke(ctx, GPSDistrictService_Query_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gPSDistrictServiceClient) IsIn(ctx context.Context, in *IsInQuery, opts ...grpc.CallOption) (*InBorderList, error) {
	out := new(InBorderList)
	err := c.cc.Invoke(ctx, GPSDistrictService_IsIn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GPSDistrictServiceServer is the server API for GPSDistrictService service.
// All implementations must embed UnimplementedGPSDistrictServiceServer
// for forward compatibility
type GPSDistrictServiceServer interface {
	// 返回的区划码和坐标顺序一一对应
	Query(context.Context, *CoordinateList) (*AdcodeList, error)
	// 返回经纬度是否在某区划内 和坐标顺序一一对应
	IsIn(context.Context, *IsInQuery) (*InBorderList, error)
	mustEmbedUnimplementedGPSDistrictServiceServer()
}

// UnimplementedGPSDistrictServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGPSDistrictServiceServer struct {
}

func (UnimplementedGPSDistrictServiceServer) Query(context.Context, *CoordinateList) (*AdcodeList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedGPSDistrictServiceServer) IsIn(context.Context, *IsInQuery) (*InBorderList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsIn not implemented")
}
func (UnimplementedGPSDistrictServiceServer) mustEmbedUnimplementedGPSDistrictServiceServer() {}

// UnsafeGPSDistrictServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GPSDistrictServiceServer will
// result in compilation errors.
type UnsafeGPSDistrictServiceServer interface {
	mustEmbedUnimplementedGPSDistrictServiceServer()
}

func RegisterGPSDistrictServiceServer(s grpc.ServiceRegistrar, srv GPSDistrictServiceServer) {
	s.RegisterService(&GPSDistrictService_ServiceDesc, srv)
}

func _GPSDistrictService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoordinateList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GPSDistrictServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GPSDistrictService_Query_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GPSDistrictServiceServer).Query(ctx, req.(*CoordinateList))
	}
	return interceptor(ctx, in, info, handler)
}

func _GPSDistrictService_IsIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsInQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GPSDistrictServiceServer).IsIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GPSDistrictService_IsIn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GPSDistrictServiceServer).IsIn(ctx, req.(*IsInQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// GPSDistrictService_ServiceDesc is the grpc.ServiceDesc for GPSDistrictService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GPSDistrictService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gps_district_service.GPSDistrictService",
	HandlerType: (*GPSDistrictServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _GPSDistrictService_Query_Handler,
		},
		{
			MethodName: "IsIn",
			Handler:    _GPSDistrictService_IsIn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gps-district-service/gpsquery.proto",
}
