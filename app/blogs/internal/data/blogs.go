package data

import (
	"blog/app/blogs/internal/biz"
	"blog/internal/ent"
	"blog/internal/ent/blogs"
	"blog/internal/ent/blogscontent"
	"blog/internal/ent/migrate"
	"context"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
)

const (
	NotHidden = iota
	IsHidden
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
	_, err = tx.BlogsContent.Create().SetID(blog.ID).SetContent(data.Content).Save(ctx)
	if err != nil {
		return err
	}
	_, err = tx.Account.UpdateOneID(int(data.AccountId)).AddBlogNum(1).Save(ctx)
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
	_, err = tx.BlogsContent.UpdateOneID(int(data.Id)).SetContent(data.Content).Save(ctx)
	return err
}
func (o *blogsRepo) DeleteBlogs(ctx context.Context, data *biz.Blogs) error {
	blog, err := o.data.db.Blogs.Query().Where(blogs.ID(int(data.Id))).First(ctx)
	if err != nil {
		return err
	}
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
	row, err := tx.Blogs.Delete().Where(blogs.ID(int(data.Id))).Where(blogs.AccountIDEQ(int(data.AccountId))).Exec(ctx)
	if err != nil {
		return err
	}
	if row == 0 {
		return errors.New("blog not found")
	}
	_, err = tx.Account.UpdateOneID(int(blog.AccountID)).AddBlogNum(-1).Save(ctx)
	return err
}
func (o *blogsRepo) GetBlogs(ctx context.Context, data *biz.Blogs) (*biz.Blogs, error) {
	blog, err := o.data.db.Blogs.Query().Where(blogs.ID(int(data.Id))).
		Where(blogs.Or(blogs.AccountIDEQ(int(data.AccountId)), blogs.IsHiddenEQ(NotHidden))).First(ctx)
	if err != nil {
		return nil, err
	}
	content, err := o.data.db.BlogsContent.Query().Where(blogscontent.ID(int(data.Id))).First(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Blogs{
		CreatedAt:   blog.CreatedAt,
		UpdatedAt:   blog.UpdatedAt,
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
	tx := o.data.db.Blogs.Query().Where(blogs.Or(
		blogs.AccountIDEQ(int(query.AccountId)),
		blogs.IsHiddenEQ(NotHidden),
	))
	if query.Title != "" {
		tx = tx.Where(blogs.TitleContains(query.Title))
	}
	if len(query.Tags) > 0 {
		tx = tx.Where(func(selector *sql.Selector) {
			selector.Where(sqljson.ValueContains(blogs.FieldTags, query.Tags))
		})
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
			UpdatedAt:   blog.UpdatedAt,
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
