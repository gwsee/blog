package service

import (
	"blog/app/blogs/internal/biz"
	"context"

	pb "blog/api/blogs/v1"
)

type BlogsCommentService struct {
	pb.UnimplementedBlogsCommentServer
	bcc *biz.BlogsCommentUseCase
}

func NewBlogsCommentService(bcc *biz.BlogsCommentUseCase) *BlogsCommentService {
	return &BlogsCommentService{bcc: bcc}
}

func (s *BlogsCommentService) CreateBlogsComment(ctx context.Context, req *pb.CreateBlogsCommentRequest) (*pb.CreateBlogsCommentReply, error) {
	return s.bcc.CreateBlogsComment(ctx, req)
}
func (s *BlogsCommentService) UpdateBlogsComment(ctx context.Context, req *pb.UpdateBlogsCommentRequest) (*pb.UpdateBlogsCommentReply, error) {
	return s.bcc.UpdateBlogsComment(ctx, req)
}
func (s *BlogsCommentService) DeleteBlogsComment(ctx context.Context, req *pb.DeleteBlogsCommentRequest) (*pb.DeleteBlogsCommentReply, error) {
	return s.bcc.DeleteBlogsComment(ctx, req)
}
func (s *BlogsCommentService) GetBlogsComment(ctx context.Context, req *pb.GetBlogsCommentRequest) (*pb.GetBlogsCommentReply, error) {
	return s.bcc.GetBlogsComment(ctx, req)
}
func (s *BlogsCommentService) ListBlogsComment(ctx context.Context, req *pb.ListBlogsCommentRequest) (*pb.ListBlogsCommentReply, error) {
	return s.bcc.ListBlogsComment(ctx, req)
}
