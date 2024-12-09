package biz

import (
	"blog/api/global"
	v1 "blog/api/tools/v1"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type ToolUseCase struct {
	repo ToolRepo
	log  *log.Helper
}

type ToolRepo interface {
	UploadFileByStream(context.Context, *v1.StreamRequest) (*v1.UploadFileReply, error)
	UploadFile(context.Context, *v1.UploadFileRequest) (*v1.UploadFileReply, error)
	Files(context.Context, *global.IDStr) (*global.Byte, error)
}

func NewToolUseCase(repo ToolRepo, logger log.Logger) *ToolUseCase {
	return &ToolUseCase{repo: repo, log: log.NewHelper(logger)}
}
func (o *ToolUseCase) UploadFileByStream(ctx context.Context, req *v1.StreamRequest) (*v1.UploadFileReply, error) {
	return o.repo.UploadFileByStream(ctx, req)
}
func (o *ToolUseCase) UploadFile(ctx context.Context, req *v1.UploadFileRequest) (*v1.UploadFileReply, error) {
	return o.repo.UploadFile(ctx, req)
}
func (o *ToolUseCase) Files(ctx context.Context, req *global.IDStr) (*global.Byte, error) {
	return o.repo.Files(ctx, req)
}
