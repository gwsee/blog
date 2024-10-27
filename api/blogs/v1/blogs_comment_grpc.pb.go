// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: api/blogs/v1/blogs_comment.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BlogsComment_CreateBlogsComment_FullMethodName = "/api.blogs.v1.BlogsComment/CreateBlogsComment"
	BlogsComment_UpdateBlogsComment_FullMethodName = "/api.blogs.v1.BlogsComment/UpdateBlogsComment"
	BlogsComment_DeleteBlogsComment_FullMethodName = "/api.blogs.v1.BlogsComment/DeleteBlogsComment"
	BlogsComment_GetBlogsComment_FullMethodName    = "/api.blogs.v1.BlogsComment/GetBlogsComment"
	BlogsComment_ListBlogsComment_FullMethodName   = "/api.blogs.v1.BlogsComment/ListBlogsComment"
)

// BlogsCommentClient is the client API for BlogsComment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlogsCommentClient interface {
	CreateBlogsComment(ctx context.Context, in *CreateBlogsCommentRequest, opts ...grpc.CallOption) (*CreateBlogsCommentReply, error)
	UpdateBlogsComment(ctx context.Context, in *UpdateBlogsCommentRequest, opts ...grpc.CallOption) (*UpdateBlogsCommentReply, error)
	DeleteBlogsComment(ctx context.Context, in *DeleteBlogsCommentRequest, opts ...grpc.CallOption) (*DeleteBlogsCommentReply, error)
	GetBlogsComment(ctx context.Context, in *GetBlogsCommentRequest, opts ...grpc.CallOption) (*GetBlogsCommentReply, error)
	ListBlogsComment(ctx context.Context, in *ListBlogsCommentRequest, opts ...grpc.CallOption) (*ListBlogsCommentReply, error)
}

type blogsCommentClient struct {
	cc grpc.ClientConnInterface
}

func NewBlogsCommentClient(cc grpc.ClientConnInterface) BlogsCommentClient {
	return &blogsCommentClient{cc}
}

func (c *blogsCommentClient) CreateBlogsComment(ctx context.Context, in *CreateBlogsCommentRequest, opts ...grpc.CallOption) (*CreateBlogsCommentReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateBlogsCommentReply)
	err := c.cc.Invoke(ctx, BlogsComment_CreateBlogsComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogsCommentClient) UpdateBlogsComment(ctx context.Context, in *UpdateBlogsCommentRequest, opts ...grpc.CallOption) (*UpdateBlogsCommentReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateBlogsCommentReply)
	err := c.cc.Invoke(ctx, BlogsComment_UpdateBlogsComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogsCommentClient) DeleteBlogsComment(ctx context.Context, in *DeleteBlogsCommentRequest, opts ...grpc.CallOption) (*DeleteBlogsCommentReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteBlogsCommentReply)
	err := c.cc.Invoke(ctx, BlogsComment_DeleteBlogsComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogsCommentClient) GetBlogsComment(ctx context.Context, in *GetBlogsCommentRequest, opts ...grpc.CallOption) (*GetBlogsCommentReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBlogsCommentReply)
	err := c.cc.Invoke(ctx, BlogsComment_GetBlogsComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogsCommentClient) ListBlogsComment(ctx context.Context, in *ListBlogsCommentRequest, opts ...grpc.CallOption) (*ListBlogsCommentReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListBlogsCommentReply)
	err := c.cc.Invoke(ctx, BlogsComment_ListBlogsComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlogsCommentServer is the server API for BlogsComment service.
// All implementations must embed UnimplementedBlogsCommentServer
// for forward compatibility.
type BlogsCommentServer interface {
	CreateBlogsComment(context.Context, *CreateBlogsCommentRequest) (*CreateBlogsCommentReply, error)
	UpdateBlogsComment(context.Context, *UpdateBlogsCommentRequest) (*UpdateBlogsCommentReply, error)
	DeleteBlogsComment(context.Context, *DeleteBlogsCommentRequest) (*DeleteBlogsCommentReply, error)
	GetBlogsComment(context.Context, *GetBlogsCommentRequest) (*GetBlogsCommentReply, error)
	ListBlogsComment(context.Context, *ListBlogsCommentRequest) (*ListBlogsCommentReply, error)
	mustEmbedUnimplementedBlogsCommentServer()
}

// UnimplementedBlogsCommentServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBlogsCommentServer struct{}

func (UnimplementedBlogsCommentServer) CreateBlogsComment(context.Context, *CreateBlogsCommentRequest) (*CreateBlogsCommentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlogsComment not implemented")
}
func (UnimplementedBlogsCommentServer) UpdateBlogsComment(context.Context, *UpdateBlogsCommentRequest) (*UpdateBlogsCommentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBlogsComment not implemented")
}
func (UnimplementedBlogsCommentServer) DeleteBlogsComment(context.Context, *DeleteBlogsCommentRequest) (*DeleteBlogsCommentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBlogsComment not implemented")
}
func (UnimplementedBlogsCommentServer) GetBlogsComment(context.Context, *GetBlogsCommentRequest) (*GetBlogsCommentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlogsComment not implemented")
}
func (UnimplementedBlogsCommentServer) ListBlogsComment(context.Context, *ListBlogsCommentRequest) (*ListBlogsCommentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBlogsComment not implemented")
}
func (UnimplementedBlogsCommentServer) mustEmbedUnimplementedBlogsCommentServer() {}
func (UnimplementedBlogsCommentServer) testEmbeddedByValue()                      {}

// UnsafeBlogsCommentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlogsCommentServer will
// result in compilation errors.
type UnsafeBlogsCommentServer interface {
	mustEmbedUnimplementedBlogsCommentServer()
}

func RegisterBlogsCommentServer(s grpc.ServiceRegistrar, srv BlogsCommentServer) {
	// If the following call pancis, it indicates UnimplementedBlogsCommentServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BlogsComment_ServiceDesc, srv)
}

func _BlogsComment_CreateBlogsComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlogsCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogsCommentServer).CreateBlogsComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlogsComment_CreateBlogsComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogsCommentServer).CreateBlogsComment(ctx, req.(*CreateBlogsCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogsComment_UpdateBlogsComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBlogsCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogsCommentServer).UpdateBlogsComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlogsComment_UpdateBlogsComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogsCommentServer).UpdateBlogsComment(ctx, req.(*UpdateBlogsCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogsComment_DeleteBlogsComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBlogsCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogsCommentServer).DeleteBlogsComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlogsComment_DeleteBlogsComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogsCommentServer).DeleteBlogsComment(ctx, req.(*DeleteBlogsCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogsComment_GetBlogsComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlogsCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogsCommentServer).GetBlogsComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlogsComment_GetBlogsComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogsCommentServer).GetBlogsComment(ctx, req.(*GetBlogsCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogsComment_ListBlogsComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBlogsCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogsCommentServer).ListBlogsComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlogsComment_ListBlogsComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogsCommentServer).ListBlogsComment(ctx, req.(*ListBlogsCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BlogsComment_ServiceDesc is the grpc.ServiceDesc for BlogsComment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlogsComment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.blogs.v1.BlogsComment",
	HandlerType: (*BlogsCommentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBlogsComment",
			Handler:    _BlogsComment_CreateBlogsComment_Handler,
		},
		{
			MethodName: "UpdateBlogsComment",
			Handler:    _BlogsComment_UpdateBlogsComment_Handler,
		},
		{
			MethodName: "DeleteBlogsComment",
			Handler:    _BlogsComment_DeleteBlogsComment_Handler,
		},
		{
			MethodName: "GetBlogsComment",
			Handler:    _BlogsComment_GetBlogsComment_Handler,
		},
		{
			MethodName: "ListBlogsComment",
			Handler:    _BlogsComment_ListBlogsComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/blogs/v1/blogs_comment.proto",
}
