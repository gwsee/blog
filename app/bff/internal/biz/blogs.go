package biz

import (
	blogs "blog/api/blogs/v1"
	"blog/api/global"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/anypb"
)

type BlogsUseCase struct {
	repo BlogsRepo
	log  *log.Helper
}

type BlogsRepo interface {
	CreateBlogs(ctx context.Context, data *blogs.CreateBlogsRequest) (*blogs.CreateBlogsReply, error)
	UpdateBlogs(ctx context.Context, data *blogs.UpdateBlogsRequest) (*blogs.UpdateBlogsReply, error)
	DeleteBlogs(ctx context.Context, data *blogs.DeleteBlogsRequest) (*blogs.DeleteBlogsReply, error)
	GetBlogs(ctx context.Context, data *blogs.GetBlogsRequest) (*blogs.GetBlogsReply, error)
	ListBlogs(ctx context.Context, data *blogs.ListBlogsRequest) (*blogs.ListBlogsReply, error)
	CreateBlogsComment(ctx context.Context, data *blogs.CreateBlogsCommentRequest) (*blogs.CreateBlogsCommentReply, error)
	UpdateBlogsComment(ctx context.Context, data *blogs.UpdateBlogsCommentRequest) (*blogs.UpdateBlogsCommentReply, error)
	DeleteBlogsComment(ctx context.Context, data *blogs.DeleteBlogsCommentRequest) (*blogs.DeleteBlogsCommentReply, error)
	GetBlogsComment(ctx context.Context, data *blogs.GetBlogsCommentRequest) (*blogs.GetBlogsCommentReply, error)
	ListBlogsComment(ctx context.Context, data *blogs.ListBlogsCommentRequest) (*blogs.ListBlogsCommentReply, error)
}

func NewBlogsUseCase(repo BlogsRepo, logger log.Logger) *BlogsUseCase {
	return &BlogsUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (l *BlogsUseCase) CreateBlogs(ctx context.Context, data *blogs.CreateBlogsRequest) (*global.Response, error) {
	res, err := l.repo.CreateBlogs(ctx, data)
	anyData, err := anypb.New(res)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}
	return &global.Response{
		Data: anyData,
	}, nil
}
func (l *BlogsUseCase) UpdateBlogs(ctx context.Context, data *blogs.UpdateBlogsRequest) (*global.Response, error) {
	res, err := l.repo.UpdateBlogs(ctx, data)
	anyData, err := anypb.New(res)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}
	return &global.Response{
		Data: anyData,
	}, nil
}
func (l *BlogsUseCase) DeleteBlogs(ctx context.Context, data *blogs.DeleteBlogsRequest) (*global.Response, error) {
	res, err := l.repo.DeleteBlogs(ctx, data)
	anyData, err := anypb.New(res)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}
	return &global.Response{
		Data: anyData,
	}, nil
}
func (l *BlogsUseCase) GetBlogs(ctx context.Context, data *blogs.GetBlogsRequest) (*global.Response, error) {
	res, err := l.repo.GetBlogs(ctx, data)
	anyData, err := anypb.New(res)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}
	return &global.Response{
		Data: anyData,
	}, nil
}
func (l *BlogsUseCase) ListBlogs(ctx context.Context, data *blogs.ListBlogsRequest) (*global.Response, error) {
	res, err := l.repo.ListBlogs(ctx, data)
	anyData, err := anypb.New(res)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}
	return &global.Response{
		Data: anyData,
	}, nil
}
func (l *BlogsUseCase) CreateBlogsComment(ctx context.Context, data *blogs.CreateBlogsCommentRequest) (*global.Response, error) {
	res, err := l.repo.CreateBlogsComment(ctx, data)
	anyData, err := anypb.New(res)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}
	return &global.Response{
		Data: anyData,
	}, nil
}
func (l *BlogsUseCase) UpdateBlogsComment(ctx context.Context, data *blogs.UpdateBlogsCommentRequest) (*global.Response, error) {
	res, err := l.repo.UpdateBlogsComment(ctx, data)
	anyData, err := anypb.New(res)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}
	return &global.Response{
		Data: anyData,
	}, nil
}
func (l *BlogsUseCase) DeleteBlogsComment(ctx context.Context, data *blogs.DeleteBlogsCommentRequest) (*global.Response, error) {
	res, err := l.repo.DeleteBlogsComment(ctx, data)
	anyData, err := anypb.New(res)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}
	return &global.Response{
		Data: anyData,
	}, nil
}
func (l *BlogsUseCase) GetBlogsComment(ctx context.Context, data *blogs.GetBlogsCommentRequest) (*global.Response, error) {
	res, err := l.repo.GetBlogsComment(ctx, data)
	anyData, err := anypb.New(res)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}
	return &global.Response{
		Data: anyData,
	}, nil
}
func (l *BlogsUseCase) ListBlogsComment(ctx context.Context, data *blogs.ListBlogsCommentRequest) (*global.Response, error) {
	res, err := l.repo.ListBlogsComment(ctx, data)
	if err != nil {
		return nil, err
	}
	anyData, err := anypb.New(res)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}
	return &global.Response{
		Data: anyData,
	}, nil
}
