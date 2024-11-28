package service

import (
	"blog/api/global"
	"blog/app/blogs/internal/biz"
	"context"

	pb "blog/api/blogs/v1"
)

type BlogsService struct {
	pb.UnimplementedBlogsServer
	bc *biz.BlogsUseCase
}

func NewBlogsService(bc *biz.BlogsUseCase) *BlogsService {
	return &BlogsService{bc: bc}
}

func (s *BlogsService) CreateBlogs(ctx context.Context, req *pb.CreateBlogsRequest) (*global.Empty, error) {
	return s.bc.CreateBlogs(ctx, req)
}
func (s *BlogsService) UpdateBlogs(ctx context.Context, req *pb.UpdateBlogsRequest) (*global.Empty, error) {
	return s.bc.UpdateBlogs(ctx, req)
}
func (s *BlogsService) DeleteBlogs(ctx context.Context, req *global.ID) (*global.Empty, error) {
	return s.bc.DeleteBlogs(ctx, req)
}
func (s *BlogsService) GetBlogs(ctx context.Context, req *global.ID) (*pb.GetBlogsReply, error) {
	return s.bc.GetBlogs(ctx, req)
}
func (s *BlogsService) ListBlogs(ctx context.Context, req *pb.ListBlogsRequest) (*pb.ListBlogsReply, error) {
	return s.bc.ListBlogs(ctx, req)
}
