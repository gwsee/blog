package biz

import (
	blogs "blog/api/blogs/v1"
	"blog/api/global"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type BlogsUseCase struct {
	repo BlogsRepo
	log  *log.Helper
}

type BlogsRepo interface {
	CreateBlogs(ctx context.Context, data *blogs.CreateBlogsRequest) (*global.Empty, error)
	UpdateBlogs(ctx context.Context, data *blogs.UpdateBlogsRequest) (*global.Empty, error)
	DeleteBlogs(ctx context.Context, data *global.ID) (*global.Empty, error)
	GetBlogs(ctx context.Context, data *global.ID) (*blogs.GetBlogsReply, error)
	ListBlogs(ctx context.Context, data *blogs.ListBlogsRequest) (*blogs.ListBlogsReply, error)
	HotBlogs(ctx context.Context, data *global.PageInfo) (*blogs.ListBlogsReply, error)
	Thumb(ctx context.Context, data *global.Action) (*global.Empty, error)
	Collect(ctx context.Context, data *global.Action) (*global.Empty, error)
	ListBlogTags(ctx context.Context, data *blogs.ListBlogsRequest) (*blogs.ListBlogTagsReply, error)
	CreateBlogsComment(ctx context.Context, data *blogs.CreateBlogsCommentRequest) (*global.Empty, error)
	UpdateBlogsComment(ctx context.Context, data *blogs.UpdateBlogsCommentRequest) (*global.Empty, error)
	DeleteBlogsComment(ctx context.Context, data *global.ID) (*global.Empty, error)
	GetBlogsComment(ctx context.Context, data *global.ID) (*blogs.GetBlogsCommentReply, error)
	ListBlogsComment(ctx context.Context, data *blogs.ListBlogsCommentRequest) (*blogs.ListBlogsCommentReply, error)
}

func NewBlogsUseCase(repo BlogsRepo, logger log.Logger) *BlogsUseCase {
	return &BlogsUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (l *BlogsUseCase) CreateBlogs(ctx context.Context, data *blogs.CreateBlogsRequest) (*global.Empty, error) {
	return l.repo.CreateBlogs(ctx, data)
	//if err != nil {
	//	return nil, err
	//}
	//anyData, err := anypb.New(res)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to marshal response data: %w", err)
	//}
	//return &global.Response{
	//	Data: anyData,
	//}, nil
}
func (l *BlogsUseCase) UpdateBlogs(ctx context.Context, data *blogs.UpdateBlogsRequest) (*global.Empty, error) {
	return l.repo.UpdateBlogs(ctx, data)
	//if err != nil {
	//	return nil, err
	//}
	//anyData, err := anypb.New(res)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to marshal response data: %w", err)
	//}
	//return &global.Response{
	//	Data: anyData,
	//}, nil
}
func (l *BlogsUseCase) DeleteBlogs(ctx context.Context, data *global.ID) (*global.Empty, error) {
	return l.repo.DeleteBlogs(ctx, data)
	//if err != nil {
	//	return nil, err
	//}
	//anyData, err := anypb.New(res)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to marshal response data: %w", err)
	//}
	//return &global.Response{
	//	Data: anyData,
	//}, nil
}
func (l *BlogsUseCase) GetBlogs(ctx context.Context, data *global.ID) (*blogs.GetBlogsReply, error) {
	return l.repo.GetBlogs(ctx, data)
	//if err != nil {
	//	return nil, err
	//}
	//anyData, err := anypb.New(res)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to marshal response data: %w", err)
	//}
	//return &global.Response{
	//	Data: anyData,
	//}, nil
}
func (l *BlogsUseCase) ListBlogs(ctx context.Context, data *blogs.ListBlogsRequest) (*blogs.ListBlogsReply, error) {
	return l.repo.ListBlogs(ctx, data)
	//if err != nil {
	//	return nil, err
	//}
	//anyData, err := anypb.New(res)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to marshal response data: %w", err)
	//}
	//return &global.Response{
	//	Data: anyData,
	//}, nil
}
func (l *BlogsUseCase) HotBlogs(ctx context.Context, data *global.PageInfo) (*blogs.ListBlogsReply, error) {
	return l.repo.HotBlogs(ctx, data)
}
func (l *BlogsUseCase) Thumb(ctx context.Context, data *global.Action) (*global.Empty, error) {
	return l.repo.Thumb(ctx, data)
}
func (l *BlogsUseCase) Collect(ctx context.Context, data *global.Action) (*global.Empty, error) {
	return l.repo.Collect(ctx, data)
}
func (l *BlogsUseCase) ListBlogTags(ctx context.Context, data *blogs.ListBlogsRequest) (*blogs.ListBlogTagsReply, error) {
	return l.repo.ListBlogTags(ctx, data)
}

func (l *BlogsUseCase) CreateBlogsComment(ctx context.Context, data *blogs.CreateBlogsCommentRequest) (*global.Empty, error) {
	return l.repo.CreateBlogsComment(ctx, data)
	//if err != nil {
	//	return nil, err
	//}
	//anyData, err := anypb.New(res)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to marshal response data: %w", err)
	//}
	//return &global.Response{
	//	Data: anyData,
	//}, nil
}
func (l *BlogsUseCase) UpdateBlogsComment(ctx context.Context, data *blogs.UpdateBlogsCommentRequest) (*global.Empty, error) {
	return l.repo.UpdateBlogsComment(ctx, data)
	//if err != nil {
	//	return nil, err
	//}
	//anyData, err := anypb.New(res)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to marshal response data: %w", err)
	//}
	//return &global.Response{
	//	Data: anyData,
	//}, nil
}
func (l *BlogsUseCase) DeleteBlogsComment(ctx context.Context, data *global.ID) (*global.Empty, error) {
	return l.repo.DeleteBlogsComment(ctx, data)
	//if err != nil {
	//	return nil, err
	//}
	//anyData, err := anypb.New(res)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to marshal response data: %w", err)
	//}
	//return &global.Response{
	//	Data: anyData,
	//}, nil
}
func (l *BlogsUseCase) GetBlogsComment(ctx context.Context, data *global.ID) (*blogs.GetBlogsCommentReply, error) {
	return l.repo.GetBlogsComment(ctx, data)
	//if err != nil {
	//	return nil, err
	//}
	//anyData, err := anypb.New(res)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to marshal response data: %w", err)
	//}
	//return &global.Response{
	//	Data: anyData,
	//}, nil
}
func (l *BlogsUseCase) ListBlogsComment(ctx context.Context, data *blogs.ListBlogsCommentRequest) (*blogs.ListBlogsCommentReply, error) {
	return l.repo.ListBlogsComment(ctx, data)
	//if err != nil {
	//	return nil, err
	//}
	//anyData, err := anypb.New(res)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to marshal response data: %w", err)
	//}
	//return &global.Response{
	//	Data: anyData,
	//}, nil
}
