package data

import (
	blogs "blog/api/blogs/v1"
	"blog/app/bff/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type blogsRepo struct {
	data *Data
	log  *log.Helper
}

func NewBlogsRepo(data *Data, logger log.Logger) biz.BlogsRepo {
	return &blogsRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (o *blogsRepo) CreateBlogs(ctx context.Context, data *blogs.CreateBlogsRequest) (*blogs.CreateBlogsReply, error) {
	return o.data.bc.CreateBlogs(ctx, data)
}
func (o *blogsRepo) UpdateBlogs(ctx context.Context, data *blogs.UpdateBlogsRequest) (*blogs.UpdateBlogsReply, error) {
	return o.data.bc.UpdateBlogs(ctx, data)
}
func (o *blogsRepo) DeleteBlogs(ctx context.Context, data *blogs.DeleteBlogsRequest) (*blogs.DeleteBlogsReply, error) {
	return o.data.bc.DeleteBlogs(ctx, data)
}
func (o *blogsRepo) GetBlogs(ctx context.Context, data *blogs.GetBlogsRequest) (*blogs.GetBlogsReply, error) {
	return o.data.bc.GetBlogs(ctx, data)
}
func (o *blogsRepo) ListBlogs(ctx context.Context, data *blogs.ListBlogsRequest) (*blogs.ListBlogsReply, error) {
	return o.data.bc.ListBlogs(ctx, data)
}

func (o *blogsRepo) CreateBlogsComment(ctx context.Context, data *blogs.CreateBlogsCommentRequest) (*blogs.CreateBlogsCommentReply, error) {
	return o.data.bcm.CreateBlogsComment(ctx, data)
}
func (o *blogsRepo) UpdateBlogsComment(ctx context.Context, data *blogs.UpdateBlogsCommentRequest) (*blogs.UpdateBlogsCommentReply, error) {
	return o.data.bcm.UpdateBlogsComment(ctx, data)
}
func (o *blogsRepo) DeleteBlogsComment(ctx context.Context, data *blogs.DeleteBlogsCommentRequest) (*blogs.DeleteBlogsCommentReply, error) {
	return o.data.bcm.DeleteBlogsComment(ctx, data)
}
func (o *blogsRepo) GetBlogsComment(ctx context.Context, data *blogs.GetBlogsCommentRequest) (*blogs.GetBlogsCommentReply, error) {
	return o.data.bcm.GetBlogsComment(ctx, data)
}
func (o *blogsRepo) ListBlogsComment(ctx context.Context, data *blogs.ListBlogsCommentRequest) (*blogs.ListBlogsCommentReply, error) {
	return o.data.bcm.ListBlogsComment(ctx, data)
}
