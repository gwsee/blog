// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.2
// - protoc             v3.19.4
// source: api/user/v1/user.proto

package v1

import (
	global "blog/api/global"
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationUserDeleteExperience = "/api.user.v1.User/DeleteExperience"
const OperationUserDeleteProject = "/api.user.v1.User/DeleteProject"
const OperationUserGetExperience = "/api.user.v1.User/GetExperience"
const OperationUserGetProject = "/api.user.v1.User/GetProject"
const OperationUserGetUser = "/api.user.v1.User/GetUser"
const OperationUserListExperience = "/api.user.v1.User/ListExperience"
const OperationUserListProject = "/api.user.v1.User/ListProject"
const OperationUserSaveExperience = "/api.user.v1.User/SaveExperience"
const OperationUserSaveProject = "/api.user.v1.User/SaveProject"
const OperationUserSaveUser = "/api.user.v1.User/SaveUser"

type UserHTTPServer interface {
	DeleteExperience(context.Context, *global.ID) (*global.Empty, error)
	DeleteProject(context.Context, *global.ID) (*global.Empty, error)
	GetExperience(context.Context, *global.ID) (*GetExperienceReply, error)
	GetProject(context.Context, *global.ID) (*GetProjectReply, error)
	GetUser(context.Context, *global.Empty) (*GetUserReply, error)
	ListExperience(context.Context, *ListExperienceRequest) (*ListExperienceReply, error)
	ListProject(context.Context, *ListProjectRequest) (*ListProjectReply, error)
	SaveExperience(context.Context, *SaveExperienceRequest) (*global.Empty, error)
	SaveProject(context.Context, *SaveProjectRequest) (*global.Empty, error)
	SaveUser(context.Context, *SaveUserRequest) (*global.Empty, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.POST("/api.user.v1.User/SaveUser", _User_SaveUser0_HTTP_Handler(srv))
	r.POST("/api.user.v1.User/GetUser", _User_GetUser0_HTTP_Handler(srv))
	r.POST("/api.user.v1.User/SaveProject", _User_SaveProject0_HTTP_Handler(srv))
	r.POST("/api.user.v1.User/DeleteProject", _User_DeleteProject0_HTTP_Handler(srv))
	r.POST("/api.user.v1.User/GetProject", _User_GetProject0_HTTP_Handler(srv))
	r.POST("/api.user.v1.User/ListProject", _User_ListProject0_HTTP_Handler(srv))
	r.POST("/api.user.v1.User/SaveExperience", _User_SaveExperience0_HTTP_Handler(srv))
	r.POST("/api.user.v1.User/DeleteExperience", _User_DeleteExperience0_HTTP_Handler(srv))
	r.POST("/api.user.v1.User/GetExperience", _User_GetExperience0_HTTP_Handler(srv))
	r.POST("/api.user.v1.User/ListExperience", _User_ListExperience0_HTTP_Handler(srv))
}

func _User_SaveUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SaveUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserSaveUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SaveUser(ctx, req.(*SaveUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*global.Empty)
		return ctx.Result(200, reply)
	}
}

func _User_GetUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in global.Empty
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUser(ctx, req.(*global.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_SaveProject0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SaveProjectRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserSaveProject)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SaveProject(ctx, req.(*SaveProjectRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*global.Empty)
		return ctx.Result(200, reply)
	}
}

func _User_DeleteProject0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in global.ID
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserDeleteProject)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteProject(ctx, req.(*global.ID))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*global.Empty)
		return ctx.Result(200, reply)
	}
}

func _User_GetProject0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in global.ID
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetProject)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetProject(ctx, req.(*global.ID))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetProjectReply)
		return ctx.Result(200, reply)
	}
}

func _User_ListProject0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListProjectRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserListProject)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListProject(ctx, req.(*ListProjectRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListProjectReply)
		return ctx.Result(200, reply)
	}
}

func _User_SaveExperience0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SaveExperienceRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserSaveExperience)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SaveExperience(ctx, req.(*SaveExperienceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*global.Empty)
		return ctx.Result(200, reply)
	}
}

func _User_DeleteExperience0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in global.ID
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserDeleteExperience)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteExperience(ctx, req.(*global.ID))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*global.Empty)
		return ctx.Result(200, reply)
	}
}

func _User_GetExperience0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in global.ID
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetExperience)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetExperience(ctx, req.(*global.ID))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetExperienceReply)
		return ctx.Result(200, reply)
	}
}

func _User_ListExperience0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListExperienceRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserListExperience)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListExperience(ctx, req.(*ListExperienceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListExperienceReply)
		return ctx.Result(200, reply)
	}
}

type UserHTTPClient interface {
	DeleteExperience(ctx context.Context, req *global.ID, opts ...http.CallOption) (rsp *global.Empty, err error)
	DeleteProject(ctx context.Context, req *global.ID, opts ...http.CallOption) (rsp *global.Empty, err error)
	GetExperience(ctx context.Context, req *global.ID, opts ...http.CallOption) (rsp *GetExperienceReply, err error)
	GetProject(ctx context.Context, req *global.ID, opts ...http.CallOption) (rsp *GetProjectReply, err error)
	GetUser(ctx context.Context, req *global.Empty, opts ...http.CallOption) (rsp *GetUserReply, err error)
	ListExperience(ctx context.Context, req *ListExperienceRequest, opts ...http.CallOption) (rsp *ListExperienceReply, err error)
	ListProject(ctx context.Context, req *ListProjectRequest, opts ...http.CallOption) (rsp *ListProjectReply, err error)
	SaveExperience(ctx context.Context, req *SaveExperienceRequest, opts ...http.CallOption) (rsp *global.Empty, err error)
	SaveProject(ctx context.Context, req *SaveProjectRequest, opts ...http.CallOption) (rsp *global.Empty, err error)
	SaveUser(ctx context.Context, req *SaveUserRequest, opts ...http.CallOption) (rsp *global.Empty, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) DeleteExperience(ctx context.Context, in *global.ID, opts ...http.CallOption) (*global.Empty, error) {
	var out global.Empty
	pattern := "/api.user.v1.User/DeleteExperience"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserDeleteExperience))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserHTTPClientImpl) DeleteProject(ctx context.Context, in *global.ID, opts ...http.CallOption) (*global.Empty, error) {
	var out global.Empty
	pattern := "/api.user.v1.User/DeleteProject"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserDeleteProject))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserHTTPClientImpl) GetExperience(ctx context.Context, in *global.ID, opts ...http.CallOption) (*GetExperienceReply, error) {
	var out GetExperienceReply
	pattern := "/api.user.v1.User/GetExperience"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserGetExperience))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserHTTPClientImpl) GetProject(ctx context.Context, in *global.ID, opts ...http.CallOption) (*GetProjectReply, error) {
	var out GetProjectReply
	pattern := "/api.user.v1.User/GetProject"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserGetProject))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserHTTPClientImpl) GetUser(ctx context.Context, in *global.Empty, opts ...http.CallOption) (*GetUserReply, error) {
	var out GetUserReply
	pattern := "/api.user.v1.User/GetUser"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserGetUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserHTTPClientImpl) ListExperience(ctx context.Context, in *ListExperienceRequest, opts ...http.CallOption) (*ListExperienceReply, error) {
	var out ListExperienceReply
	pattern := "/api.user.v1.User/ListExperience"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserListExperience))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserHTTPClientImpl) ListProject(ctx context.Context, in *ListProjectRequest, opts ...http.CallOption) (*ListProjectReply, error) {
	var out ListProjectReply
	pattern := "/api.user.v1.User/ListProject"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserListProject))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserHTTPClientImpl) SaveExperience(ctx context.Context, in *SaveExperienceRequest, opts ...http.CallOption) (*global.Empty, error) {
	var out global.Empty
	pattern := "/api.user.v1.User/SaveExperience"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserSaveExperience))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserHTTPClientImpl) SaveProject(ctx context.Context, in *SaveProjectRequest, opts ...http.CallOption) (*global.Empty, error) {
	var out global.Empty
	pattern := "/api.user.v1.User/SaveProject"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserSaveProject))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserHTTPClientImpl) SaveUser(ctx context.Context, in *SaveUserRequest, opts ...http.CallOption) (*global.Empty, error) {
	var out global.Empty
	pattern := "/api.user.v1.User/SaveUser"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserSaveUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
