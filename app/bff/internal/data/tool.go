package data

import (
	"blog/api/global"
	v1 "blog/api/tools/v1"
	"blog/app/bff/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type toolsRepo struct {
	data *Data
	log  *log.Helper
}

func NewToolsRepo(data *Data, logger log.Logger) biz.ToolRepo {
	return &toolsRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (o *toolsRepo) UploadFileByStream(ctx context.Context, in *v1.StreamRequest) (*v1.UploadFileReply, error) {
	return o.data.t.UploadFileByStream(ctx, in)
}
func (o *toolsRepo) UploadFile(ctx context.Context, in *v1.UploadFileRequest) (*v1.UploadFileReply, error) {
	return o.data.t.UploadFile(ctx, in)
}
func (o *toolsRepo) Files(ctx context.Context, in *global.IDStr) (*global.Byte, error) {
	return o.data.t.Files(ctx, in)
}
