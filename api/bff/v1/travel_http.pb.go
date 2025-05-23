// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v5.28.2
// source: api/bff/v1/travel.proto

package v1

import (
	global "blog/api/global"
	v1 "blog/api/travel/v1"
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationTravelsCollect = "/api.bff.v1.Travels/Collect"
const OperationTravelsDeleteTravel = "/api.bff.v1.Travels/DeleteTravel"
const OperationTravelsGetTravel = "/api.bff.v1.Travels/GetTravel"
const OperationTravelsListTravel = "/api.bff.v1.Travels/ListTravel"
const OperationTravelsSaveTravel = "/api.bff.v1.Travels/SaveTravel"
const OperationTravelsThumb = "/api.bff.v1.Travels/Thumb"

type TravelsHTTPServer interface {
	Collect(context.Context, *global.Action) (*global.Empty, error)
	DeleteTravel(context.Context, *global.ID) (*global.Empty, error)
	GetTravel(context.Context, *global.ID) (*v1.GetTravelReply, error)
	ListTravel(context.Context, *v1.ListTravelRequest) (*v1.ListTravelReply, error)
	SaveTravel(context.Context, *v1.SaveTravelRequest) (*global.Empty, error)
	Thumb(context.Context, *global.Action) (*global.Empty, error)
}

func RegisterTravelsHTTPServer(s *http.Server, srv TravelsHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/travel/save", _Travels_SaveTravel0_HTTP_Handler(srv))
	r.POST("/v1/travel/del", _Travels_DeleteTravel0_HTTP_Handler(srv))
	r.POST("/v1/travel/get", _Travels_GetTravel0_HTTP_Handler(srv))
	r.POST("/v1/travel/list", _Travels_ListTravel0_HTTP_Handler(srv))
	r.POST("/v1/travel/thumb", _Travels_Thumb0_HTTP_Handler(srv))
	r.POST("/v1/travel/collect", _Travels_Collect0_HTTP_Handler(srv))
}

func _Travels_SaveTravel0_HTTP_Handler(srv TravelsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.SaveTravelRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTravelsSaveTravel)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SaveTravel(ctx, req.(*v1.SaveTravelRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*global.Empty)
		return ctx.Result(200, reply)
	}
}

func _Travels_DeleteTravel0_HTTP_Handler(srv TravelsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in global.ID
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTravelsDeleteTravel)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteTravel(ctx, req.(*global.ID))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*global.Empty)
		return ctx.Result(200, reply)
	}
}

func _Travels_GetTravel0_HTTP_Handler(srv TravelsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in global.ID
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTravelsGetTravel)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTravel(ctx, req.(*global.ID))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.GetTravelReply)
		return ctx.Result(200, reply)
	}
}

func _Travels_ListTravel0_HTTP_Handler(srv TravelsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.ListTravelRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTravelsListTravel)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListTravel(ctx, req.(*v1.ListTravelRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.ListTravelReply)
		return ctx.Result(200, reply)
	}
}

func _Travels_Thumb0_HTTP_Handler(srv TravelsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in global.Action
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTravelsThumb)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Thumb(ctx, req.(*global.Action))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*global.Empty)
		return ctx.Result(200, reply)
	}
}

func _Travels_Collect0_HTTP_Handler(srv TravelsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in global.Action
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTravelsCollect)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Collect(ctx, req.(*global.Action))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*global.Empty)
		return ctx.Result(200, reply)
	}
}

type TravelsHTTPClient interface {
	Collect(ctx context.Context, req *global.Action, opts ...http.CallOption) (rsp *global.Empty, err error)
	DeleteTravel(ctx context.Context, req *global.ID, opts ...http.CallOption) (rsp *global.Empty, err error)
	GetTravel(ctx context.Context, req *global.ID, opts ...http.CallOption) (rsp *v1.GetTravelReply, err error)
	ListTravel(ctx context.Context, req *v1.ListTravelRequest, opts ...http.CallOption) (rsp *v1.ListTravelReply, err error)
	SaveTravel(ctx context.Context, req *v1.SaveTravelRequest, opts ...http.CallOption) (rsp *global.Empty, err error)
	Thumb(ctx context.Context, req *global.Action, opts ...http.CallOption) (rsp *global.Empty, err error)
}

type TravelsHTTPClientImpl struct {
	cc *http.Client
}

func NewTravelsHTTPClient(client *http.Client) TravelsHTTPClient {
	return &TravelsHTTPClientImpl{client}
}

func (c *TravelsHTTPClientImpl) Collect(ctx context.Context, in *global.Action, opts ...http.CallOption) (*global.Empty, error) {
	var out global.Empty
	pattern := "/v1/travel/collect"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTravelsCollect))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TravelsHTTPClientImpl) DeleteTravel(ctx context.Context, in *global.ID, opts ...http.CallOption) (*global.Empty, error) {
	var out global.Empty
	pattern := "/v1/travel/del"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTravelsDeleteTravel))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TravelsHTTPClientImpl) GetTravel(ctx context.Context, in *global.ID, opts ...http.CallOption) (*v1.GetTravelReply, error) {
	var out v1.GetTravelReply
	pattern := "/v1/travel/get"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTravelsGetTravel))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TravelsHTTPClientImpl) ListTravel(ctx context.Context, in *v1.ListTravelRequest, opts ...http.CallOption) (*v1.ListTravelReply, error) {
	var out v1.ListTravelReply
	pattern := "/v1/travel/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTravelsListTravel))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TravelsHTTPClientImpl) SaveTravel(ctx context.Context, in *v1.SaveTravelRequest, opts ...http.CallOption) (*global.Empty, error) {
	var out global.Empty
	pattern := "/v1/travel/save"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTravelsSaveTravel))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TravelsHTTPClientImpl) Thumb(ctx context.Context, in *global.Action, opts ...http.CallOption) (*global.Empty, error) {
	var out global.Empty
	pattern := "/v1/travel/thumb"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTravelsThumb))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
