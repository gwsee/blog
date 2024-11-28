package global

import (
	"github.com/go-kratos/kratos/v2/transport/http"
)

type Response struct {
	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    any    `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (r *Response) Resp(ctx http.Context, err error) {
	if err != nil {
		r.Message = err.Error()
		_ = ctx.Result(200, r)
	} else {
		r.Code = 1
		_ = ctx.Result(200, r)
	}
	ctx.Done()
}
