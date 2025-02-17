// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v5.28.2
// source: api/bff/v1/account.proto

package v1

import (
	v1 "blog/api/account/v1"
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

const OperationAccountCreateAccount = "/api.bff.v1.Account/CreateAccount"
const OperationAccountInfo = "/api.bff.v1.Account/Info"
const OperationAccountLoginByAccount = "/api.bff.v1.Account/LoginByAccount"
const OperationAccountResetPassword = "/api.bff.v1.Account/ResetPassword"
const OperationAccountUpdateAccount = "/api.bff.v1.Account/UpdateAccount"

type AccountHTTPServer interface {
	CreateAccount(context.Context, *v1.CreateAccountRequest) (*global.Empty, error)
	Info(context.Context, *global.Empty) (*v1.AccountInfoReply, error)
	LoginByAccount(context.Context, *v1.LoginByAccountRequest) (*v1.LoginByAccountReply, error)
	ResetPassword(context.Context, *v1.ResetPasswordRequest) (*global.Empty, error)
	UpdateAccount(context.Context, *v1.UpdateAccountRequest) (*global.Empty, error)
}

func RegisterAccountHTTPServer(s *http.Server, srv AccountHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/register", _Account_CreateAccount0_HTTP_Handler(srv))
	r.POST("/v1/reset_password", _Account_ResetPassword0_HTTP_Handler(srv))
	r.POST("/v1/login", _Account_LoginByAccount0_HTTP_Handler(srv))
	r.POST("/v1/info", _Account_Info0_HTTP_Handler(srv))
	r.POST("/v1/update_account", _Account_UpdateAccount0_HTTP_Handler(srv))
}

func _Account_CreateAccount0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.CreateAccountRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAccountCreateAccount)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateAccount(ctx, req.(*v1.CreateAccountRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*global.Empty)
		return ctx.Result(200, reply)
	}
}

func _Account_ResetPassword0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.ResetPasswordRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAccountResetPassword)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ResetPassword(ctx, req.(*v1.ResetPasswordRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*global.Empty)
		return ctx.Result(200, reply)
	}
}

func _Account_LoginByAccount0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.LoginByAccountRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAccountLoginByAccount)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LoginByAccount(ctx, req.(*v1.LoginByAccountRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.LoginByAccountReply)
		return ctx.Result(200, reply)
	}
}

func _Account_Info0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in global.Empty
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAccountInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Info(ctx, req.(*global.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.AccountInfoReply)
		return ctx.Result(200, reply)
	}
}

func _Account_UpdateAccount0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.UpdateAccountRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAccountUpdateAccount)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateAccount(ctx, req.(*v1.UpdateAccountRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*global.Empty)
		return ctx.Result(200, reply)
	}
}

type AccountHTTPClient interface {
	CreateAccount(ctx context.Context, req *v1.CreateAccountRequest, opts ...http.CallOption) (rsp *global.Empty, err error)
	Info(ctx context.Context, req *global.Empty, opts ...http.CallOption) (rsp *v1.AccountInfoReply, err error)
	LoginByAccount(ctx context.Context, req *v1.LoginByAccountRequest, opts ...http.CallOption) (rsp *v1.LoginByAccountReply, err error)
	ResetPassword(ctx context.Context, req *v1.ResetPasswordRequest, opts ...http.CallOption) (rsp *global.Empty, err error)
	UpdateAccount(ctx context.Context, req *v1.UpdateAccountRequest, opts ...http.CallOption) (rsp *global.Empty, err error)
}

type AccountHTTPClientImpl struct {
	cc *http.Client
}

func NewAccountHTTPClient(client *http.Client) AccountHTTPClient {
	return &AccountHTTPClientImpl{client}
}

func (c *AccountHTTPClientImpl) CreateAccount(ctx context.Context, in *v1.CreateAccountRequest, opts ...http.CallOption) (*global.Empty, error) {
	var out global.Empty
	pattern := "/v1/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAccountCreateAccount))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AccountHTTPClientImpl) Info(ctx context.Context, in *global.Empty, opts ...http.CallOption) (*v1.AccountInfoReply, error) {
	var out v1.AccountInfoReply
	pattern := "/v1/info"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAccountInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AccountHTTPClientImpl) LoginByAccount(ctx context.Context, in *v1.LoginByAccountRequest, opts ...http.CallOption) (*v1.LoginByAccountReply, error) {
	var out v1.LoginByAccountReply
	pattern := "/v1/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAccountLoginByAccount))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AccountHTTPClientImpl) ResetPassword(ctx context.Context, in *v1.ResetPasswordRequest, opts ...http.CallOption) (*global.Empty, error) {
	var out global.Empty
	pattern := "/v1/reset_password"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAccountResetPassword))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AccountHTTPClientImpl) UpdateAccount(ctx context.Context, in *v1.UpdateAccountRequest, opts ...http.CallOption) (*global.Empty, error) {
	var out global.Empty
	pattern := "/v1/update_account"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAccountUpdateAccount))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
