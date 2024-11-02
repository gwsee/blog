package global

import (
	"github.com/go-kratos/kratos/v2/transport/http"
)

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
