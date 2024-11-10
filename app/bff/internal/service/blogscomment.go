package service

import (
	"blog/api/global"
	"blog/app/bff/internal/biz"
	"context"

	pb "blog/api/bff/v1"
	blog "blog/api/blogs/v1"
)

type BlogsCommentService struct {
	pb.UnimplementedBlogsCommentServer
	bc *biz.BlogsUseCase
}

func NewBlogsCommentService(bc *biz.BlogsUseCase) *BlogsCommentService {
	return &BlogsCommentService{bc: bc}
}

func (s *BlogsCommentService) CreateBlogsComment(ctx context.Context, req *blog.CreateBlogsCommentRequest) (*global.Response, error) {
	return s.bc.CreateBlogsComment(ctx, req)
}
func (s *BlogsCommentService) UpdateBlogsComment(ctx context.Context, req *blog.UpdateBlogsCommentRequest) (*global.Response, error) {
	return s.bc.UpdateBlogsComment(ctx, req)
}
func (s *BlogsCommentService) DeleteBlogsComment(ctx context.Context, req *blog.DeleteBlogsCommentRequest) (*global.Response, error) {
	return s.bc.DeleteBlogsComment(ctx, req)
}
func (s *BlogsCommentService) GetBlogsComment(ctx context.Context, req *blog.GetBlogsCommentRequest) (*global.Response, error) {
	return s.bc.GetBlogsComment(ctx, req)
}
func (s *BlogsCommentService) ListBlogsComment(ctx context.Context, req *blog.ListBlogsCommentRequest) (*global.Response, error) {
	return s.bc.ListBlogsComment(ctx, req)
}
