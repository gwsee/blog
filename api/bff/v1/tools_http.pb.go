// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v5.28.2
// source: api/bff/v1/tools.proto

package v1

import (
	global "blog/api/global"
	v1 "blog/api/tools/v1"
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationToolsFiles = "/api.bff.v1.Tools/Files"
const OperationToolsUploadFile = "/api.bff.v1.Tools/UploadFile"
const OperationToolsUploadFileByStream = "/api.bff.v1.Tools/UploadFileByStream"

type ToolsHTTPServer interface {
	Files(context.Context, *global.IDStr) (*global.IDStr, error)
	// UploadFile上传文件流的方式
	UploadFile(context.Context, *v1.UploadFileRequest) (*v1.UploadFileReply, error)
	UploadFileByStream(context.Context, *v1.StreamRequest) (*v1.UploadFileReply, error)
}

func RegisterToolsHTTPServer(s *http.Server, srv ToolsHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/upload/stream", _Tools_UploadFileByStream0_HTTP_Handler(srv))
	r.POST("/v1/upload/file", _Tools_UploadFile0_HTTP_Handler(srv))
	r.GET("/v1/file/{id}", _Tools_Files0_HTTP_Handler(srv))
}

func _Tools_UploadFileByStream0_HTTP_Handler(srv ToolsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.StreamRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationToolsUploadFileByStream)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UploadFileByStream(ctx, req.(*v1.StreamRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.UploadFileReply)
		return ctx.Result(200, reply)
	}
}




type ToolsHTTPClient interface {
	Files(ctx context.Context, req *global.IDStr, opts ...http.CallOption) (rsp *global.IDStr, err error)
	UploadFile(ctx context.Context, req *v1.UploadFileRequest, opts ...http.CallOption) (rsp *v1.UploadFileReply, err error)
	UploadFileByStream(ctx context.Context, req *v1.StreamRequest, opts ...http.CallOption) (rsp *v1.UploadFileReply, err error)
}

type ToolsHTTPClientImpl struct {
	cc *http.Client
}

func NewToolsHTTPClient(client *http.Client) ToolsHTTPClient {
	return &ToolsHTTPClientImpl{client}
}

func (c *ToolsHTTPClientImpl) Files(ctx context.Context, in *global.IDStr, opts ...http.CallOption) (*global.IDStr, error) {
	var out global.IDStr
	pattern := "/v1/file/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationToolsFiles))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ToolsHTTPClientImpl) UploadFile(ctx context.Context, in *v1.UploadFileRequest, opts ...http.CallOption) (*v1.UploadFileReply, error) {
	var out v1.UploadFileReply
	pattern := "/v1/upload/file"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationToolsUploadFile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ToolsHTTPClientImpl) UploadFileByStream(ctx context.Context, in *v1.StreamRequest, opts ...http.CallOption) (*v1.UploadFileReply, error) {
	var out v1.UploadFileReply
	pattern := "/v1/upload/stream"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationToolsUploadFileByStream))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
