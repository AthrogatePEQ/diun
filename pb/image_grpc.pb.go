// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// ImageServiceClient is the client API for ImageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageServiceClient interface {
	ImageList(ctx context.Context, in *ImageListRequest, opts ...grpc.CallOption) (*ImageListResponse, error)
	ImageInspect(ctx context.Context, in *ImageInspectRequest, opts ...grpc.CallOption) (*ImageInspectResponse, error)
	ImageRemove(ctx context.Context, in *ImageRemoveRequest, opts ...grpc.CallOption) (*ImageRemoveResponse, error)
}

type imageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImageServiceClient(cc grpc.ClientConnInterface) ImageServiceClient {
	return &imageServiceClient{cc}
}

func (c *imageServiceClient) ImageList(ctx context.Context, in *ImageListRequest, opts ...grpc.CallOption) (*ImageListResponse, error) {
	out := new(ImageListResponse)
	err := c.cc.Invoke(ctx, "/pb.ImageService/ImageList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageServiceClient) ImageInspect(ctx context.Context, in *ImageInspectRequest, opts ...grpc.CallOption) (*ImageInspectResponse, error) {
	out := new(ImageInspectResponse)
	err := c.cc.Invoke(ctx, "/pb.ImageService/ImageInspect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageServiceClient) ImageRemove(ctx context.Context, in *ImageRemoveRequest, opts ...grpc.CallOption) (*ImageRemoveResponse, error) {
	out := new(ImageRemoveResponse)
	err := c.cc.Invoke(ctx, "/pb.ImageService/ImageRemove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageServiceServer is the server API for ImageService service.
// All implementations must embed UnimplementedImageServiceServer
// for forward compatibility
type ImageServiceServer interface {
	ImageList(context.Context, *ImageListRequest) (*ImageListResponse, error)
	ImageInspect(context.Context, *ImageInspectRequest) (*ImageInspectResponse, error)
	ImageRemove(context.Context, *ImageRemoveRequest) (*ImageRemoveResponse, error)
	mustEmbedUnimplementedImageServiceServer()
}

// UnimplementedImageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedImageServiceServer struct {
}

func (UnimplementedImageServiceServer) ImageList(context.Context, *ImageListRequest) (*ImageListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImageList not implemented")
}
func (UnimplementedImageServiceServer) ImageInspect(context.Context, *ImageInspectRequest) (*ImageInspectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImageInspect not implemented")
}
func (UnimplementedImageServiceServer) ImageRemove(context.Context, *ImageRemoveRequest) (*ImageRemoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImageRemove not implemented")
}
func (UnimplementedImageServiceServer) mustEmbedUnimplementedImageServiceServer() {}

// UnsafeImageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageServiceServer will
// result in compilation errors.
type UnsafeImageServiceServer interface {
	mustEmbedUnimplementedImageServiceServer()
}

func RegisterImageServiceServer(s grpc.ServiceRegistrar, srv ImageServiceServer) {
	s.RegisterService(&ImageService_ServiceDesc, srv)
}

func _ImageService_ImageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).ImageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ImageService/ImageList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).ImageList(ctx, req.(*ImageListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageService_ImageInspect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageInspectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).ImageInspect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ImageService/ImageInspect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).ImageInspect(ctx, req.(*ImageInspectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageService_ImageRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageRemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).ImageRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ImageService/ImageRemove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).ImageRemove(ctx, req.(*ImageRemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImageService_ServiceDesc is the grpc.ServiceDesc for ImageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ImageService",
	HandlerType: (*ImageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ImageList",
			Handler:    _ImageService_ImageList_Handler,
		},
		{
			MethodName: "ImageInspect",
			Handler:    _ImageService_ImageInspect_Handler,
		},
		{
			MethodName: "ImageRemove",
			Handler:    _ImageService_ImageRemove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "image.proto",
}
