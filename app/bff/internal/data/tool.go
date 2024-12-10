package data

import (
	"blog/api/global"
	v1 "blog/api/tools/v1"
	"blog/app/bff/internal/biz"
	"blog/internal/constx"
	"context"
	"fmt"
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
func (o *toolsRepo) Files(ctx context.Context, in *global.IDStr) (*global.IDStr, error) {
	res, err := o.data.redisCli.Get(ctx, constx.FileCachePrefix+in.Id).Result()
	if err != nil {
		//获取缓存失败
		fmt.Println("获取缓存失败")
		return o.data.t.Files(ctx, in)
	}
	return &global.IDStr{Id: res}, nil
}
