package service

import (
	"context"

	pb "blog/app/blogs/api"
)

type BlogsService struct {
	pb.UnimplementedBlogsServer
}

func NewBlogsService() *BlogsService {
	return &BlogsService{}
}

func (s *BlogsService) CreateBlogs(ctx context.Context, req *pb.CreateBlogsRequest) (*pb.CreateBlogsReply, error) {
	return &pb.CreateBlogsReply{}, nil
}
func (s *BlogsService) UpdateBlogs(ctx context.Context, req *pb.UpdateBlogsRequest) (*pb.UpdateBlogsReply, error) {
	return &pb.UpdateBlogsReply{}, nil
}
func (s *BlogsService) DeleteBlogs(ctx context.Context, req *pb.DeleteBlogsRequest) (*pb.DeleteBlogsReply, error) {
	return &pb.DeleteBlogsReply{}, nil
}
func (s *BlogsService) GetBlogs(ctx context.Context, req *pb.GetBlogsRequest) (*pb.GetBlogsReply, error) {
	return &pb.GetBlogsReply{}, nil
}
func (s *BlogsService) ListBlogs(ctx context.Context, req *pb.ListBlogsRequest) (*pb.ListBlogsReply, error) {
	return &pb.ListBlogsReply{}, nil
}
