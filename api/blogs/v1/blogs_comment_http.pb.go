// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v5.28.2
// source: api/blogs/v1/blogs_comment.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationBlogsCommentCreateBlogsComment = "/api.blogs.v1.BlogsComment/CreateBlogsComment"
const OperationBlogsCommentDeleteBlogsComment = "/api.blogs.v1.BlogsComment/DeleteBlogsComment"
const OperationBlogsCommentGetBlogsComment = "/api.blogs.v1.BlogsComment/GetBlogsComment"
const OperationBlogsCommentListBlogsComment = "/api.blogs.v1.BlogsComment/ListBlogsComment"
const OperationBlogsCommentUpdateBlogsComment = "/api.blogs.v1.BlogsComment/UpdateBlogsComment"

type BlogsCommentHTTPServer interface {
	CreateBlogsComment(context.Context, *CreateBlogsCommentRequest) (*CreateBlogsCommentReply, error)
	DeleteBlogsComment(context.Context, *DeleteBlogsCommentRequest) (*DeleteBlogsCommentReply, error)
	GetBlogsComment(context.Context, *GetBlogsCommentRequest) (*GetBlogsCommentReply, error)
	ListBlogsComment(context.Context, *ListBlogsCommentRequest) (*ListBlogsCommentReply, error)
	UpdateBlogsComment(context.Context, *UpdateBlogsCommentRequest) (*UpdateBlogsCommentReply, error)
}

func RegisterBlogsCommentHTTPServer(s *http.Server, srv BlogsCommentHTTPServer) {
	r := s.Route("/")
	r.POST("/api.blogs.v1.BlogsComment/CreateBlogsComment", _BlogsComment_CreateBlogsComment0_HTTP_Handler(srv))
	r.POST("/api.blogs.v1.BlogsComment/UpdateBlogsComment", _BlogsComment_UpdateBlogsComment0_HTTP_Handler(srv))
	r.POST("/api.blogs.v1.BlogsComment/DeleteBlogsComment", _BlogsComment_DeleteBlogsComment0_HTTP_Handler(srv))
	r.POST("/api.blogs.v1.BlogsComment/GetBlogsComment", _BlogsComment_GetBlogsComment0_HTTP_Handler(srv))
	r.POST("/api.blogs.v1.BlogsComment/ListBlogsComment", _BlogsComment_ListBlogsComment0_HTTP_Handler(srv))
}

func _BlogsComment_CreateBlogsComment0_HTTP_Handler(srv BlogsCommentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateBlogsCommentRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBlogsCommentCreateBlogsComment)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateBlogsComment(ctx, req.(*CreateBlogsCommentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateBlogsCommentReply)
		return ctx.Result(200, reply)
	}
}

func _BlogsComment_UpdateBlogsComment0_HTTP_Handler(srv BlogsCommentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateBlogsCommentRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBlogsCommentUpdateBlogsComment)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateBlogsComment(ctx, req.(*UpdateBlogsCommentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateBlogsCommentReply)
		return ctx.Result(200, reply)
	}
}

func _BlogsComment_DeleteBlogsComment0_HTTP_Handler(srv BlogsCommentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteBlogsCommentRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBlogsCommentDeleteBlogsComment)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteBlogsComment(ctx, req.(*DeleteBlogsCommentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteBlogsCommentReply)
		return ctx.Result(200, reply)
	}
}

func _BlogsComment_GetBlogsComment0_HTTP_Handler(srv BlogsCommentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetBlogsCommentRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBlogsCommentGetBlogsComment)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetBlogsComment(ctx, req.(*GetBlogsCommentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetBlogsCommentReply)
		return ctx.Result(200, reply)
	}
}

func _BlogsComment_ListBlogsComment0_HTTP_Handler(srv BlogsCommentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListBlogsCommentRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBlogsCommentListBlogsComment)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListBlogsComment(ctx, req.(*ListBlogsCommentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListBlogsCommentReply)
		return ctx.Result(200, reply)
	}
}

type BlogsCommentHTTPClient interface {
	CreateBlogsComment(ctx context.Context, req *CreateBlogsCommentRequest, opts ...http.CallOption) (rsp *CreateBlogsCommentReply, err error)
	DeleteBlogsComment(ctx context.Context, req *DeleteBlogsCommentRequest, opts ...http.CallOption) (rsp *DeleteBlogsCommentReply, err error)
	GetBlogsComment(ctx context.Context, req *GetBlogsCommentRequest, opts ...http.CallOption) (rsp *GetBlogsCommentReply, err error)
	ListBlogsComment(ctx context.Context, req *ListBlogsCommentRequest, opts ...http.CallOption) (rsp *ListBlogsCommentReply, err error)
	UpdateBlogsComment(ctx context.Context, req *UpdateBlogsCommentRequest, opts ...http.CallOption) (rsp *UpdateBlogsCommentReply, err error)
}

type BlogsCommentHTTPClientImpl struct {
	cc *http.Client
}

func NewBlogsCommentHTTPClient(client *http.Client) BlogsCommentHTTPClient {
	return &BlogsCommentHTTPClientImpl{client}
}

func (c *BlogsCommentHTTPClientImpl) CreateBlogsComment(ctx context.Context, in *CreateBlogsCommentRequest, opts ...http.CallOption) (*CreateBlogsCommentReply, error) {
	var out CreateBlogsCommentReply
	pattern := "/api.blogs.v1.BlogsComment/CreateBlogsComment"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBlogsCommentCreateBlogsComment))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BlogsCommentHTTPClientImpl) DeleteBlogsComment(ctx context.Context, in *DeleteBlogsCommentRequest, opts ...http.CallOption) (*DeleteBlogsCommentReply, error) {
	var out DeleteBlogsCommentReply
	pattern := "/api.blogs.v1.BlogsComment/DeleteBlogsComment"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBlogsCommentDeleteBlogsComment))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BlogsCommentHTTPClientImpl) GetBlogsComment(ctx context.Context, in *GetBlogsCommentRequest, opts ...http.CallOption) (*GetBlogsCommentReply, error) {
	var out GetBlogsCommentReply
	pattern := "/api.blogs.v1.BlogsComment/GetBlogsComment"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBlogsCommentGetBlogsComment))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BlogsCommentHTTPClientImpl) ListBlogsComment(ctx context.Context, in *ListBlogsCommentRequest, opts ...http.CallOption) (*ListBlogsCommentReply, error) {
	var out ListBlogsCommentReply
	pattern := "/api.blogs.v1.BlogsComment/ListBlogsComment"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBlogsCommentListBlogsComment))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BlogsCommentHTTPClientImpl) UpdateBlogsComment(ctx context.Context, in *UpdateBlogsCommentRequest, opts ...http.CallOption) (*UpdateBlogsCommentReply, error) {
	var out UpdateBlogsCommentReply
	pattern := "/api.blogs.v1.BlogsComment/UpdateBlogsComment"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBlogsCommentUpdateBlogsComment))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
