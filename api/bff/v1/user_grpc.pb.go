// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.4
// source: api/bff/v1/user.proto

package v1

import (
	global "blog/api/global"
	v1 "blog/api/user/v1"
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
	User_SaveUser_FullMethodName         = "/api.bff.v1.User/SaveUser"
	User_GetUser_FullMethodName          = "/api.bff.v1.User/GetUser"
	User_SaveProject_FullMethodName      = "/api.bff.v1.User/SaveProject"
	User_DeleteProject_FullMethodName    = "/api.bff.v1.User/DeleteProject"
	User_GetProject_FullMethodName       = "/api.bff.v1.User/GetProject"
	User_ListProject_FullMethodName      = "/api.bff.v1.User/ListProject"
	User_SaveExperience_FullMethodName   = "/api.bff.v1.User/SaveExperience"
	User_DeleteExperience_FullMethodName = "/api.bff.v1.User/DeleteExperience"
	User_GetExperience_FullMethodName    = "/api.bff.v1.User/GetExperience"
	User_ListExperience_FullMethodName   = "/api.bff.v1.User/ListExperience"
	User_Photos_FullMethodName           = "/api.bff.v1.User/Photos"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	// 个人信息管理
	SaveUser(ctx context.Context, in *v1.SaveUserRequest, opts ...grpc.CallOption) (*global.Empty, error)
	GetUser(ctx context.Context, in *global.Empty, opts ...grpc.CallOption) (*v1.GetUserReply, error)
	// 项目管理
	SaveProject(ctx context.Context, in *v1.SaveProjectRequest, opts ...grpc.CallOption) (*global.Empty, error)
	DeleteProject(ctx context.Context, in *global.ID, opts ...grpc.CallOption) (*global.Empty, error)
	GetProject(ctx context.Context, in *global.ID, opts ...grpc.CallOption) (*v1.GetProjectReply, error)
	ListProject(ctx context.Context, in *v1.ListProjectRequest, opts ...grpc.CallOption) (*v1.ListProjectReply, error)
	// 工作经验管理
	SaveExperience(ctx context.Context, in *v1.SaveExperienceRequest, opts ...grpc.CallOption) (*global.Empty, error)
	DeleteExperience(ctx context.Context, in *global.ID, opts ...grpc.CallOption) (*global.Empty, error)
	GetExperience(ctx context.Context, in *global.ID, opts ...grpc.CallOption) (*v1.GetExperienceReply, error)
	ListExperience(ctx context.Context, in *v1.ListExperienceRequest, opts ...grpc.CallOption) (*v1.ListExperienceReply, error)
	Photos(ctx context.Context, in *v1.PhotosReq, opts ...grpc.CallOption) (*v1.PhotosReply, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) SaveUser(ctx context.Context, in *v1.SaveUserRequest, opts ...grpc.CallOption) (*global.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(global.Empty)
	err := c.cc.Invoke(ctx, User_SaveUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUser(ctx context.Context, in *global.Empty, opts ...grpc.CallOption) (*v1.GetUserReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.GetUserReply)
	err := c.cc.Invoke(ctx, User_GetUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SaveProject(ctx context.Context, in *v1.SaveProjectRequest, opts ...grpc.CallOption) (*global.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(global.Empty)
	err := c.cc.Invoke(ctx, User_SaveProject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DeleteProject(ctx context.Context, in *global.ID, opts ...grpc.CallOption) (*global.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(global.Empty)
	err := c.cc.Invoke(ctx, User_DeleteProject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetProject(ctx context.Context, in *global.ID, opts ...grpc.CallOption) (*v1.GetProjectReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.GetProjectReply)
	err := c.cc.Invoke(ctx, User_GetProject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ListProject(ctx context.Context, in *v1.ListProjectRequest, opts ...grpc.CallOption) (*v1.ListProjectReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.ListProjectReply)
	err := c.cc.Invoke(ctx, User_ListProject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SaveExperience(ctx context.Context, in *v1.SaveExperienceRequest, opts ...grpc.CallOption) (*global.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(global.Empty)
	err := c.cc.Invoke(ctx, User_SaveExperience_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DeleteExperience(ctx context.Context, in *global.ID, opts ...grpc.CallOption) (*global.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(global.Empty)
	err := c.cc.Invoke(ctx, User_DeleteExperience_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetExperience(ctx context.Context, in *global.ID, opts ...grpc.CallOption) (*v1.GetExperienceReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.GetExperienceReply)
	err := c.cc.Invoke(ctx, User_GetExperience_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ListExperience(ctx context.Context, in *v1.ListExperienceRequest, opts ...grpc.CallOption) (*v1.ListExperienceReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.ListExperienceReply)
	err := c.cc.Invoke(ctx, User_ListExperience_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Photos(ctx context.Context, in *v1.PhotosReq, opts ...grpc.CallOption) (*v1.PhotosReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.PhotosReply)
	err := c.cc.Invoke(ctx, User_Photos_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility.
type UserServer interface {
	// 个人信息管理
	SaveUser(context.Context, *v1.SaveUserRequest) (*global.Empty, error)
	GetUser(context.Context, *global.Empty) (*v1.GetUserReply, error)
	// 项目管理
	SaveProject(context.Context, *v1.SaveProjectRequest) (*global.Empty, error)
	DeleteProject(context.Context, *global.ID) (*global.Empty, error)
	GetProject(context.Context, *global.ID) (*v1.GetProjectReply, error)
	ListProject(context.Context, *v1.ListProjectRequest) (*v1.ListProjectReply, error)
	// 工作经验管理
	SaveExperience(context.Context, *v1.SaveExperienceRequest) (*global.Empty, error)
	DeleteExperience(context.Context, *global.ID) (*global.Empty, error)
	GetExperience(context.Context, *global.ID) (*v1.GetExperienceReply, error)
	ListExperience(context.Context, *v1.ListExperienceRequest) (*v1.ListExperienceReply, error)
	Photos(context.Context, *v1.PhotosReq) (*v1.PhotosReply, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServer struct{}

func (UnimplementedUserServer) SaveUser(context.Context, *v1.SaveUserRequest) (*global.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveUser not implemented")
}
func (UnimplementedUserServer) GetUser(context.Context, *global.Empty) (*v1.GetUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServer) SaveProject(context.Context, *v1.SaveProjectRequest) (*global.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveProject not implemented")
}
func (UnimplementedUserServer) DeleteProject(context.Context, *global.ID) (*global.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProject not implemented")
}
func (UnimplementedUserServer) GetProject(context.Context, *global.ID) (*v1.GetProjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProject not implemented")
}
func (UnimplementedUserServer) ListProject(context.Context, *v1.ListProjectRequest) (*v1.ListProjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProject not implemented")
}
func (UnimplementedUserServer) SaveExperience(context.Context, *v1.SaveExperienceRequest) (*global.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveExperience not implemented")
}
func (UnimplementedUserServer) DeleteExperience(context.Context, *global.ID) (*global.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteExperience not implemented")
}
func (UnimplementedUserServer) GetExperience(context.Context, *global.ID) (*v1.GetExperienceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExperience not implemented")
}
func (UnimplementedUserServer) ListExperience(context.Context, *v1.ListExperienceRequest) (*v1.ListExperienceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListExperience not implemented")
}
func (UnimplementedUserServer) Photos(context.Context, *v1.PhotosReq) (*v1.PhotosReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Photos not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}
func (UnimplementedUserServer) testEmbeddedByValue()              {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	// If the following call pancis, it indicates UnimplementedUserServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_SaveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.SaveUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SaveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_SaveUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SaveUser(ctx, req.(*v1.SaveUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(global.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUser(ctx, req.(*global.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SaveProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.SaveProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SaveProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_SaveProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SaveProject(ctx, req.(*v1.SaveProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DeleteProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(global.ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DeleteProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_DeleteProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DeleteProject(ctx, req.(*global.ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(global.ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetProject(ctx, req.(*global.ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ListProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.ListProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ListProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_ListProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ListProject(ctx, req.(*v1.ListProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SaveExperience_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.SaveExperienceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SaveExperience(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_SaveExperience_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SaveExperience(ctx, req.(*v1.SaveExperienceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DeleteExperience_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(global.ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DeleteExperience(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_DeleteExperience_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DeleteExperience(ctx, req.(*global.ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetExperience_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(global.ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetExperience(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetExperience_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetExperience(ctx, req.(*global.ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ListExperience_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.ListExperienceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ListExperience(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_ListExperience_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ListExperience(ctx, req.(*v1.ListExperienceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Photos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.PhotosReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Photos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_Photos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Photos(ctx, req.(*v1.PhotosReq))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.bff.v1.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveUser",
			Handler:    _User_SaveUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _User_GetUser_Handler,
		},
		{
			MethodName: "SaveProject",
			Handler:    _User_SaveProject_Handler,
		},
		{
			MethodName: "DeleteProject",
			Handler:    _User_DeleteProject_Handler,
		},
		{
			MethodName: "GetProject",
			Handler:    _User_GetProject_Handler,
		},
		{
			MethodName: "ListProject",
			Handler:    _User_ListProject_Handler,
		},
		{
			MethodName: "SaveExperience",
			Handler:    _User_SaveExperience_Handler,
		},
		{
			MethodName: "DeleteExperience",
			Handler:    _User_DeleteExperience_Handler,
		},
		{
			MethodName: "GetExperience",
			Handler:    _User_GetExperience_Handler,
		},
		{
			MethodName: "ListExperience",
			Handler:    _User_ListExperience_Handler,
		},
		{
			MethodName: "Photos",
			Handler:    _User_Photos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/bff/v1/user.proto",
}
