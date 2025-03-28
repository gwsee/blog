package data

import (
	blogs "blog/api/blogs/v1"
	"blog/api/global"
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

func (o *blogsRepo) CreateBlogs(ctx context.Context, data *blogs.CreateBlogsRequest) (*global.Empty, error) {
	return o.data.bc.CreateBlogs(ctx, data)
}
func (o *blogsRepo) UpdateBlogs(ctx context.Context, data *blogs.UpdateBlogsRequest) (*global.Empty, error) {
	return o.data.bc.UpdateBlogs(ctx, data)
}
func (o *blogsRepo) DeleteBlogs(ctx context.Context, data *global.ID) (*global.Empty, error) {
	return o.data.bc.DeleteBlogs(ctx, data)
}
func (o *blogsRepo) GetBlogs(ctx context.Context, data *global.ID) (*blogs.GetBlogsReply, error) {
	return o.data.bc.GetBlogs(ctx, data)
}
func (o *blogsRepo) ListBlogs(ctx context.Context, data *blogs.ListBlogsRequest) (*blogs.ListBlogsReply, error) {
	return o.data.bc.ListBlogs(ctx, data)
}
func (o *blogsRepo) HotBlogs(ctx context.Context, data *global.PageInfo) (*blogs.ListBlogsReply, error) {
	return o.data.bc.HotBlogs(ctx, data)
}
func (o *blogsRepo) Thumb(ctx context.Context, data *global.Action) (*global.Empty, error) {
	return o.data.bc.Thumb(ctx, data)
}
func (o *blogsRepo) Collect(ctx context.Context, data *global.Action) (*global.Empty, error) {
	return o.data.bc.Collect(ctx, data)
}
func (o *blogsRepo) ListBlogTags(ctx context.Context, data *blogs.ListBlogsRequest) (*blogs.ListBlogTagsReply, error) {
	return o.data.bc.ListBlogTags(ctx, data)
}
func (o *blogsRepo) CreateBlogsComment(ctx context.Context, data *blogs.CreateBlogsCommentRequest) (*global.Empty, error) {
	return o.data.bcm.CreateBlogsComment(ctx, data)
}
func (o *blogsRepo) UpdateBlogsComment(ctx context.Context, data *blogs.UpdateBlogsCommentRequest) (*global.Empty, error) {
	return o.data.bcm.UpdateBlogsComment(ctx, data)
}
func (o *blogsRepo) DeleteBlogsComment(ctx context.Context, data *global.ID) (*global.Empty, error) {
	return o.data.bcm.DeleteBlogsComment(ctx, data)
}
func (o *blogsRepo) GetBlogsComment(ctx context.Context, data *global.ID) (*blogs.GetBlogsCommentReply, error) {
	return o.data.bcm.GetBlogsComment(ctx, data)
}
func (o *blogsRepo) ListBlogsComment(ctx context.Context, data *blogs.ListBlogsCommentRequest) (*blogs.ListBlogsCommentReply, error) {
	return o.data.bcm.ListBlogsComment(ctx, data)
}
