package v1

import (
	v1 "blog/api/tools/v1"
	"blog/internal/constx"
	"context"
	"github.com/go-kratos/kratos/v2/transport/http"
	"io"
)

func _Tools_UploadFile0_HTTP_Handler(srv ToolsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.UploadFileRequest
		req := ctx.Request()
		file, header, err := req.FormFile(constx.FileName)
		if err != nil {
			return err
		}
		defer file.Close()
		in.Filename = header.Filename
		in.Content, err = io.ReadAll(file)
		if err != nil {
			return err
		}
		//if err := ctx.Bind(&in); err != nil {
		//	return err
		//}
		//if err := ctx.BindQuery(&in); err != nil {
		//	return err
		//}
		http.SetOperation(ctx, OperationToolsUploadFile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UploadFile(ctx, req.(*v1.UploadFileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.UploadFileReply)
		return ctx.Result(200, reply)
	}
}
