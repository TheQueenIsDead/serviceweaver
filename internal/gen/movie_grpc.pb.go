// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: movie.proto

package gen

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

// MovieServiceClient is the client API for MovieService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MovieServiceClient interface {
	GetMovieByName(ctx context.Context, in *GetMovieByNameRequest, opts ...grpc.CallOption) (*GetMovieResponse, error)
	GetMovieByDirector(ctx context.Context, in *GetMovieByDirectorRequest, opts ...grpc.CallOption) (*GetMovieResponse, error)
}

type movieServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMovieServiceClient(cc grpc.ClientConnInterface) MovieServiceClient {
	return &movieServiceClient{cc}
}

func (c *movieServiceClient) GetMovieByName(ctx context.Context, in *GetMovieByNameRequest, opts ...grpc.CallOption) (*GetMovieResponse, error) {
	out := new(GetMovieResponse)
	err := c.cc.Invoke(ctx, "/api.MovieService/GetMovieByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieServiceClient) GetMovieByDirector(ctx context.Context, in *GetMovieByDirectorRequest, opts ...grpc.CallOption) (*GetMovieResponse, error) {
	out := new(GetMovieResponse)
	err := c.cc.Invoke(ctx, "/api.MovieService/GetMovieByDirector", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MovieServiceServer is the server API for MovieService service.
// All implementations should embed UnimplementedMovieServiceServer
// for forward compatibility
type MovieServiceServer interface {
	GetMovieByName(context.Context, *GetMovieByNameRequest) (*GetMovieResponse, error)
	GetMovieByDirector(context.Context, *GetMovieByDirectorRequest) (*GetMovieResponse, error)
}

// UnimplementedMovieServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMovieServiceServer struct {
}

func (UnimplementedMovieServiceServer) GetMovieByName(context.Context, *GetMovieByNameRequest) (*GetMovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovieByName not implemented")
}
func (UnimplementedMovieServiceServer) GetMovieByDirector(context.Context, *GetMovieByDirectorRequest) (*GetMovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovieByDirector not implemented")
}

// UnsafeMovieServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MovieServiceServer will
// result in compilation errors.
type UnsafeMovieServiceServer interface {
	mustEmbedUnimplementedMovieServiceServer()
}

func RegisterMovieServiceServer(s grpc.ServiceRegistrar, srv MovieServiceServer) {
	s.RegisterService(&MovieService_ServiceDesc, srv)
}

func _MovieService_GetMovieByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMovieByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).GetMovieByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MovieService/GetMovieByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).GetMovieByName(ctx, req.(*GetMovieByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieService_GetMovieByDirector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMovieByDirectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).GetMovieByDirector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MovieService/GetMovieByDirector",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).GetMovieByDirector(ctx, req.(*GetMovieByDirectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MovieService_ServiceDesc is the grpc.ServiceDesc for MovieService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MovieService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.MovieService",
	HandlerType: (*MovieServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMovieByName",
			Handler:    _MovieService_GetMovieByName_Handler,
		},
		{
			MethodName: "GetMovieByDirector",
			Handler:    _MovieService_GetMovieByDirector_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "movie.proto",
}
