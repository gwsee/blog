package service

import (
	pb "blog/api/bff/v1"
	"blog/api/global"
	v1 "blog/api/tools/v1"
	"blog/app/bff/internal/biz"
	"context"
)

type ToolsService struct {
	pb.UnimplementedToolsServer
	biz *biz.ToolUseCase
}

func NewToolsService(biz *biz.ToolUseCase) *ToolsService {
	return &ToolsService{biz: biz}
}
func (o *ToolsService) UploadFileByStream(ctx context.Context, in *v1.StreamRequest) (*v1.UploadFileReply, error) {
	return o.biz.UploadFileByStream(ctx, in)
}
func (o *ToolsService) UploadFile(ctx context.Context, in *v1.UploadFileRequest) (*v1.UploadFileReply, error) {
	return o.biz.UploadFile(ctx, in)
}
func (o *ToolsService) Files(ctx context.Context, in *global.IDStr) (*global.Byte, error) {
	return o.biz.Files(ctx, in)
}
