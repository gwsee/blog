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
func (o *ToolsService) Files(ctx context.Context, in *global.IDStr) (*global.IDStr, error) {
	return o.biz.Files(ctx, in)
}
func (o *ToolsService) UploadOssToken(ctx context.Context, req *global.IDStr) (*v1.UploadOssTokenReply, error) {
	return o.biz.UploadOssToken(ctx, req)
}
func (o *ToolsService) UploadOssSave(ctx context.Context, req *v1.UploadOssSaveReq) (*global.Empty, error) {
	return o.biz.UploadOssSave(ctx, req)
}
