package service

import (
	"context"

	pb "blog/app/blogs/api"
)

type BlogCommentService struct {
	pb.UnimplementedBlogCommentServer
}

func NewBlogCommentService() *BlogCommentService {
	return &BlogCommentService{}
}

func (s *BlogCommentService) CreateBlogComment(ctx context.Context, req *pb.CreateBlogCommentRequest) (*pb.CreateBlogCommentReply, error) {
	return &pb.CreateBlogCommentReply{}, nil
}
func (s *BlogCommentService) UpdateBlogComment(ctx context.Context, req *pb.UpdateBlogCommentRequest) (*pb.UpdateBlogCommentReply, error) {
	return &pb.UpdateBlogCommentReply{}, nil
}
func (s *BlogCommentService) DeleteBlogComment(ctx context.Context, req *pb.DeleteBlogCommentRequest) (*pb.DeleteBlogCommentReply, error) {
	return &pb.DeleteBlogCommentReply{}, nil
}
func (s *BlogCommentService) GetBlogComment(ctx context.Context, req *pb.GetBlogCommentRequest) (*pb.GetBlogCommentReply, error) {
	return &pb.GetBlogCommentReply{}, nil
}
func (s *BlogCommentService) ListBlogComment(ctx context.Context, req *pb.ListBlogCommentRequest) (*pb.ListBlogCommentReply, error) {
	return &pb.ListBlogCommentReply{}, nil
}
