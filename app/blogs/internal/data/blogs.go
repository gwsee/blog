package data

import (
	"blog/app/blogs/internal/biz"
	"blog/internal/ent"
	"blog/internal/ent/blogs"
	"blog/internal/ent/blogscontent"
	"blog/internal/ent/migrate"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type blogsRepo struct {
	data *Data
	log  *log.Helper
}

func NewBlogsRepo(data *Data, logger log.Logger) biz.BlogsRepo {
	if err := data.db.Schema.Create(context.TODO(), migrate.WithForeignKeys(false),
		migrate.WithDropColumn(true)); err != nil {
		panic(err)
	}
	return &blogsRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (o *blogsRepo) CreateBlogs(ctx context.Context, data *biz.Blogs) error {
	tx, err := o.data.db.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	blog, err := tx.Blogs.Create().
		SetTitle(data.Title).
		SetDescription(data.Description).
		SetAccountID(int(data.AccountId)).
		SetIsHidden(int8(data.IsHidden)).
		SetTags(data.Tags).
		SetCover(data.Cover).
		Save(ctx)
	if err != nil {
		return err
	}
	_, err = tx.BlogsComment.Create().SetBlogID(blog.ID).SetContent(data.Content).Save(ctx)
	return err
}
func (o *blogsRepo) UpdateBlogs(ctx context.Context, data *biz.Blogs) error {
	tx, err := o.data.db.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	_, err = tx.Blogs.UpdateOneID(int(data.Id)).
		SetTitle(data.Title).
		SetDescription(data.Description).
		SetIsHidden(int8(data.IsHidden)).
		SetTags(data.Tags).
		SetCover(data.Cover).
		Save(ctx)
	if err != nil {
		return err
	}
	_, err = tx.BlogsComment.UpdateOneID(int(data.Id)).SetContent(data.Content).Save(ctx)
	return err
}
func (o *blogsRepo) DeleteBlogs(ctx context.Context, id int64) error {
	_, err := o.data.db.Blogs.Delete().Where(blogs.ID(int(id))).Exec(ctx)
	return err
}
func (o *blogsRepo) GetBlogs(ctx context.Context, id int64) (*biz.Blogs, error) {
	blog, err := o.data.db.Blogs.Query().Where(blogs.ID(int(id))).First(ctx)
	if err != nil {
		return nil, err
	}
	content, err := o.data.db.BlogsContent.Query().Where(blogscontent.ID(int(id))).First(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Blogs{
		CreatedAt:   blog.CreatedAt,
		Id:          int64(blog.ID),
		IsHidden:    int64(blog.IsHidden),
		AccountId:   int64(blog.AccountID),
		Title:       blog.Title,
		Description: blog.Description,
		Tags:        blog.Tags,
		Cover:       blog.Cover,
		Content:     content.Content,
	}, err
}
func (o *blogsRepo) ListBlogs(ctx context.Context, query *biz.BlogsQuery) (int64, []*biz.Blogs, error) {
	tx := o.data.db.Blogs.Query().Where(blogs.IsHidden(0)) // todo  查看自己的或者查询公共的
	if query.Title != "" {
		tx = tx.Where(blogs.TitleContains(query.Title))
	}
	if len(query.Tags) > 0 {
		//todo 查询标签
		//tx = tx.Where(blogs.Tags)
	}
	total, err := tx.Count(ctx)
	if err != nil {
		return 0, nil, err
	}
	list, err := tx.Limit(int(query.GetPageSize())).Offset(int(query.GetOffset())).
		Order(ent.Desc(blogs.FieldID)).All(ctx)
	if err != nil {
		return 0, nil, err
	}
	data := make([]*biz.Blogs, 0, len(list))
	for _, blog := range list {
		data = append(data, &biz.Blogs{
			CreatedAt:   blog.CreatedAt,
			Id:          int64(blog.ID),
			IsHidden:    int64(blog.IsHidden),
			AccountId:   int64(blog.AccountID),
			Title:       blog.Title,
			Description: blog.Description,
			Tags:        blog.Tags,
			Cover:       blog.Cover,
		})
	}
	return int64(total), data, err
}
