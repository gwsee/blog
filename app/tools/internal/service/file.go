package service

import (
	"blog/api/global"
	"context"

	"blog/api/tools/v1"
	"blog/app/tools/internal/biz"
)

// ToolsService is a greeter service.
type ToolsService struct {
	v1.UnimplementedToolsServer

	uc *biz.ToolsUsecase
}

// NewToolsService new a greeter service.
func NewToolsService(uc *biz.ToolsUsecase) *ToolsService {
	return &ToolsService{uc: uc}
}

func (s *ToolsService) UploadFile(ctx context.Context, in *v1.UploadFileRequest) (*v1.UploadFileReply, error) {
	return s.uc.UploadFile(ctx, in)
}
func (s *ToolsService) UploadFileByStream(ctx context.Context, in *v1.StreamRequest) (*v1.UploadFileReply, error) {
	return s.uc.UploadFileByStream(ctx, in)
}
func (s *ToolsService) Files(ctx context.Context, in *global.IDStr) (*global.IDStr, error) {
	return s.uc.Files(ctx, in)
}
func (s *ToolsService) UploadOssToken(ctx context.Context, in *global.IDStr) (*v1.UploadOssTokenReply, error) {
	return s.uc.UploadOssToken(ctx, in)
}
func (s *ToolsService) UploadOssSave(ctx context.Context, req *v1.UploadOssSaveReq) (*global.Empty, error) {
	return s.uc.UploadOssSave(ctx, req)
}
