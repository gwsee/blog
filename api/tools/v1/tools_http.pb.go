// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v5.28.2
// source: api/tools/v1/tools.proto

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

const OperationToolsUploadFile = "/api.tools.v1.Tools/UploadFile"
const OperationToolsUploadFileByStream = "/api.tools.v1.Tools/UploadFileByStream"

type ToolsHTTPServer interface {
	UploadFile(context.Context, *UploadFileRequest) (*UploadFileReply, error)
	UploadFileByStream(context.Context, *StreamRequest) (*UploadFileReply, error)
}

func RegisterToolsHTTPServer(s *http.Server, srv ToolsHTTPServer) {
	r := s.Route("/")
	r.POST("/api.tools.v1.Tools/UploadFileByStream", _Tools_UploadFileByStream0_HTTP_Handler(srv))
	r.POST("/api.tools.v1.Tools/UploadFile", _Tools_UploadFile0_HTTP_Handler(srv))
}

func _Tools_UploadFileByStream0_HTTP_Handler(srv ToolsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in StreamRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationToolsUploadFileByStream)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UploadFileByStream(ctx, req.(*StreamRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UploadFileReply)
		return ctx.Result(200, reply)
	}
}

func _Tools_UploadFile0_HTTP_Handler(srv ToolsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UploadFileRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationToolsUploadFile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UploadFile(ctx, req.(*UploadFileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UploadFileReply)
		return ctx.Result(200, reply)
	}
}

type ToolsHTTPClient interface {
	UploadFile(ctx context.Context, req *UploadFileRequest, opts ...http.CallOption) (rsp *UploadFileReply, err error)
	UploadFileByStream(ctx context.Context, req *StreamRequest, opts ...http.CallOption) (rsp *UploadFileReply, err error)
}

type ToolsHTTPClientImpl struct {
	cc *http.Client
}

func NewToolsHTTPClient(client *http.Client) ToolsHTTPClient {
	return &ToolsHTTPClientImpl{client}
}

func (c *ToolsHTTPClientImpl) UploadFile(ctx context.Context, in *UploadFileRequest, opts ...http.CallOption) (*UploadFileReply, error) {
	var out UploadFileReply
	pattern := "/api.tools.v1.Tools/UploadFile"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationToolsUploadFile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ToolsHTTPClientImpl) UploadFileByStream(ctx context.Context, in *StreamRequest, opts ...http.CallOption) (*UploadFileReply, error) {
	var out UploadFileReply
	pattern := "/api.tools.v1.Tools/UploadFileByStream"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationToolsUploadFileByStream))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
