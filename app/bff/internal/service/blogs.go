package service

import (
	"blog/api/global"
	"blog/app/bff/internal/biz"
	"context"

	pb "blog/api/bff/v1"
	blog "blog/api/blogs/v1"
)

type BlogsService struct {
	pb.UnimplementedBlogsServer
	bc *biz.BlogsUseCase
}

func NewBlogsService(bc *biz.BlogsUseCase) *BlogsService {
	return &BlogsService{bc: bc}
}

func (s *BlogsService) CreateBlogs(ctx context.Context, req *blog.CreateBlogsRequest) (*global.Empty, error) {
	return s.bc.CreateBlogs(ctx, req)
}
func (s *BlogsService) UpdateBlogs(ctx context.Context, req *blog.UpdateBlogsRequest) (*global.Empty, error) {
	return s.bc.UpdateBlogs(ctx, req)
}
func (s *BlogsService) DeleteBlogs(ctx context.Context, req *global.ID) (*global.Empty, error) {
	return s.bc.DeleteBlogs(ctx, req)
}
func (s *BlogsService) GetBlogs(ctx context.Context, req *global.ID) (*blog.GetBlogsReply, error) {
	return s.bc.GetBlogs(ctx, req)
}
func (s *BlogsService) ListBlogs(ctx context.Context, req *blog.ListBlogsRequest) (*blog.ListBlogsReply, error) {
	return s.bc.ListBlogs(ctx, req)
}
