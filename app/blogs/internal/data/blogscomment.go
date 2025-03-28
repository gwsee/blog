package data

import (
	"blog/app/blogs/internal/biz"
	"blog/internal/constx"
	"blog/internal/ent"
	"blog/internal/ent/blogscomment"
	"blog/internal/ent/migrate"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type blogsCommentRepo struct {
	data *Data
	log  *log.Helper
}

func NewBlogsCommentRepo(data *Data, logger log.Logger) biz.BlogsCommentRepo {
	if err := data.db.Schema.Create(context.TODO(), migrate.WithForeignKeys(false),
		migrate.WithDropColumn(true)); err != nil {
		panic(err)
	}
	return &blogsCommentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (o *blogsCommentRepo) CreateBlogsComment(ctx context.Context, data *biz.BlogsComment) error {
	level := 1
	if data.ParentId > 0 {
		parent, _ := o.data.db.BlogsComment.Query().Where(blogscomment.ID(int(data.ParentId))).First(ctx)
		if parent != nil {
			level = parent.Level + 1
		}
	}
	_, err := o.data.db.BlogsComment.Create().
		SetAccountID(int(data.AccountId)).
		SetBlogID(int(data.BlogId)).
		SetTopID(int(data.TopId)).
		SetParentID(int(data.ParentId)).
		SetLevel(level).
		SetContent(data.Content).
		SetStatus(int8(data.Status)).
		Save(ctx)
	return err
}
func (o *blogsCommentRepo) UpdateBlogsComment(ctx context.Context, data *biz.BlogsComment) error {
	_, err := o.data.db.BlogsComment.UpdateOneID(int(data.Id)).
		SetContent(data.Content).
		SetStatus(int8(data.Status)).
		Save(ctx)
	return err
}
func (o *blogsCommentRepo) DeleteBlogsComment(ctx context.Context, id int64) error {
	return o.data.db.BlogsComment.DeleteOneID(int(id)).Exec(ctx)
}
func (o *blogsCommentRepo) GetBlogsComment(ctx context.Context, id int64) (*biz.BlogsComment, error) {
	comment, err := o.data.db.BlogsComment.Query().Where(blogscomment.ID(int(id))).First(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.BlogsComment{
		Id:        int64(comment.ID),
		BlogId:    int64(comment.BlogID),
		TopId:     int64(comment.TopID),
		ParentId:  int64(comment.ParentID),
		Level:     int64(comment.Level),
		Total:     int64(comment.Total),
		Status:    int64(comment.Status),
		Content:   comment.Content,
		AccountId: int64(comment.AccountID),
		CreatedAt: comment.CreatedAt,
	}, err
}
func (o *blogsCommentRepo) ListBlogsComment(ctx context.Context, query *biz.BlogsCommentQuery) (int64, []*biz.BlogsComment, error) {
	tx := o.data.db.BlogsComment.Query().Where(blogscomment.IDGT(int(query.GetStart()))).Where(blogscomment.Or(
		blogscomment.AccountIDEQ(int(query.AccountId)),
		blogscomment.StatusEQ(constx.BlogNotHidden),
	))
	list, err := tx.Limit(int(query.GetPageSize())).
		Order(ent.Asc(blogscomment.FieldID)).All(ctx)
	if err != nil {
		return 0, nil, err
	}
	if len(list) == 0 {
		return 0, nil, nil
	}
	data := make([]*biz.BlogsComment, 0, len(list))
	for _, comment := range list {
		data = append(data, &biz.BlogsComment{
			Id:        int64(comment.ID),
			BlogId:    int64(comment.BlogID),
			TopId:     int64(comment.TopID),
			ParentId:  int64(comment.ParentID),
			Level:     int64(comment.Level),
			Total:     int64(comment.Total),
			Status:    int64(comment.Status),
			Content:   comment.Content,
			AccountId: int64(comment.AccountID),
			CreatedAt: comment.CreatedAt,
		})
	}
	return int64(list[len(list)-1].ID), data, err
}
