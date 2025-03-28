package data

import (
	"blog/api/global"
	"blog/app/blogs/internal/biz"
	"blog/internal/constx"
	"blog/internal/ent"
	"blog/internal/ent/blogs"
	"blog/internal/ent/blogscontent"
	"blog/internal/ent/blogsextend"
	"blog/internal/ent/migrate"
	"blog/internal/ent/predicate"
	"blog/internal/ent/tags"
	"blog/internal/ent/tagsrelation"
	"context"
	"entgo.io/ent/dialect/sql"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"time"
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
	var bulkTags []int
	for _, name := range data.Tags {
		var id int
		id, err = tx.Tags.
			Create().
			SetName(name).
			OnConflictColumns("name").DoNothing(). // 冲突时不做操作
			ID(ctx)
		if err != nil {
			return err
		}
		bulkTags = append(bulkTags, id)
	}

	blog, err := tx.Blogs.Create().
		SetTitle(data.Title).
		SetDescription(data.Description).
		SetAccountID(int(data.AccountId)).
		SetIsHidden(int8(data.IsHidden)).
		SetTags(data.Tags).
		//AddTag(tagData...).
		SetCover(data.Cover).
		Save(ctx)
	if err != nil {
		return err
	}
	var bulkTagRelations []*ent.TagsRelationCreate
	for _, tag := range bulkTags {
		one := tx.TagsRelation.Create()
		one.SetBlogID(blog.ID)
		one.SetRelation(constx.Blog)
		one.SetTagID(tag)
		bulkTagRelations = append(bulkTagRelations, one)
	}
	err = tx.TagsRelation.CreateBulk(bulkTagRelations...).Exec(ctx)
	if err != nil {
		return err
	}
	_, err = tx.BlogsContent.Create().SetID(blog.ID).SetContent(data.Content).SetFiles(data.Files).Save(ctx)
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
	var bulkTags []int
	for _, name := range data.Tags {
		var id int
		id, err = tx.Tags.
			Create().
			SetName(name).
			OnConflictColumns("name").DoNothing(). // 冲突时不做操作
			ID(ctx)
		if err != nil {
			return err
		}
		bulkTags = append(bulkTags, id)
	}

	_, err = tx.TagsRelation.Delete().Where(tagsrelation.RelationEQ(constx.Blog)).Where(tagsrelation.RelationIDEQ(int(data.Id))).Exec(ctx)
	if err != nil {
		return err
	}
	var bulkTagRelations []*ent.TagsRelationCreate
	for _, tag := range bulkTags {
		one := tx.TagsRelation.Create()
		one.SetBlogID(int(data.Id))
		one.SetRelation(constx.Blog)
		one.SetTagID(tag)
		one.OnConflict().DoNothing()
		bulkTagRelations = append(bulkTagRelations, one)
	}
	err = tx.TagsRelation.CreateBulk(bulkTagRelations...).Exec(ctx)
	if err != nil {
		return err
	}
	_, err = tx.Blogs.UpdateOneID(int(data.Id)).Where(blogs.AccountIDEQ(int(data.AccountId))).
		SetTitle(data.Title).
		SetDescription(data.Description).
		SetIsHidden(int8(data.IsHidden)).
		SetTags(data.Tags).
		SetCover(data.Cover).
		Save(ctx)
	if err != nil {
		return err
	}
	_, err = tx.BlogsContent.UpdateOneID(int(data.Id)).SetContent(data.Content).SetFiles(data.Files).Save(ctx)
	return err
}
func (o *blogsRepo) DeleteBlogs(ctx context.Context, data *global.ID) error {
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
	row, err := tx.Blogs.Delete().Where(blogs.IDEQ(int(data.Id))).Where(blogs.AccountIDEQ(int(data.GetAccountId()))).Exec(ctx)
	if err != nil {
		return err
	}
	if row == 0 {
		return errors.New("blog not found")
	}
	_, err = tx.Account.UpdateOneID(int(blog.AccountID)).AddBlogNum(-1).Save(ctx)
	return err
}
func (o *blogsRepo) browse(g *global.ID) {
	ctx := context.Background()
	tx, err := o.data.db.Tx(ctx)
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = tx.BlogsExtend.Create().
		SetAccountID(int(g.GetAccountId())).SetBlogID(int(g.Id)).
		SetBrowseAt(time.Now().Unix()).SetBrowseNum(1).
		OnConflictColumns(blogsextend.FieldAccountID, blogsextend.FieldBlogID).
		AddBrowseNum(1).SetBrowseAt(time.Now().Unix()).Exec(ctx)
	if err != nil {
		return
	}
	err = tx.Blogs.UpdateOneID(int(g.Id)).AddBrowseNum(1).Exec(ctx)

}
func (o *blogsRepo) GetBlogs(ctx context.Context, data *global.ID) (*biz.Blogs, error) {
	blog, err := o.data.db.Blogs.Query().Where(blogs.ID(int(data.Id))).
		Where(blogs.Or(blogs.AccountIDEQ(int(data.GetAccountId())), blogs.IsHiddenEQ(constx.BlogNotHidden))).WithTag().First(ctx)
	if err != nil {
		return nil, err
	}
	content, err := o.data.db.BlogsContent.Query().Where(blogscontent.ID(int(data.Id))).First(ctx)
	if err != nil {
		return nil, err
	}
	go o.browse(data)
	resp := &biz.Blogs{
		CreatedAt:   blog.CreatedAt,
		UpdatedAt:   blog.UpdatedAt,
		Id:          int64(blog.ID),
		IsHidden:    int64(blog.IsHidden),
		AccountId:   int64(blog.AccountID),
		Title:       blog.Title,
		Description: blog.Description,
		//Tags:        blog.Tags,
		Cover:      blog.Cover,
		Content:    content.Content,
		Files:      content.Files,
		ThumbNum:   int64(blog.LoveNum),
		CollectNum: int64(blog.CollectNum),
		BrowseNum:  int64(blog.BrowseNum),
	}
	if len(blog.Edges.Tag) > 0 {
		for _, v := range blog.Edges.Tag {
			resp.Tags = append(resp.Tags, v.Name)
		}
	}
	return resp, err
}
func (o *blogsRepo) ListBlogs(ctx context.Context, query *biz.BlogsQuery) (int64, []*biz.Blogs, error) {
	tx := o.data.db.Blogs.Query().Where(blogs.Or(
		blogs.AccountIDEQ(int(query.AccountId)),
		blogs.IsHiddenEQ(constx.BlogNotHidden),
	))
	if query.Title != "" {
		tx = tx.Where(blogs.TitleContains(query.Title))
	}
	if len(query.Tags) > 0 {
		//tx = tx.Where(func(selector *sql.Selector) {
		//	selector.Where(sqljson.ValueContains(blogs.FieldTags, query.Tags))
		//})
		tx = tx.Where(blogs.HasTagRelationWith(
			tagsrelation.RelationEQ(constx.Blog),
			tagsrelation.HasTagWith(
				tags.NameIn(query.Tags...),
			),
		))
		//tx = tx.Where(blogs.Tags)
	}
	if query.Mine {
		tx = tx.Where(blogs.AccountIDEQ(int(query.AccountId)))
	}
	total, err := tx.Count(ctx)
	if err != nil {
		return 0, nil, err
	}
	list, err := tx.Limit(int(query.GetPageSize())).Offset(int(query.GetOffset())).
		Order(ent.Desc(blogs.FieldID)).WithTag().All(ctx)
	if err != nil {
		return 0, nil, err
	}
	data := make([]*biz.Blogs, 0, len(list))
	for _, blog := range list {
		resp := &biz.Blogs{
			CreatedAt:   blog.CreatedAt,
			UpdatedAt:   blog.UpdatedAt,
			Id:          int64(blog.ID),
			IsHidden:    int64(blog.IsHidden),
			AccountId:   int64(blog.AccountID),
			Title:       blog.Title,
			Description: blog.Description,
			//Tags:        blog.Tags,
			Cover:      blog.Cover,
			ThumbNum:   int64(blog.LoveNum),
			CollectNum: int64(blog.CollectNum),
			BrowseNum:  int64(blog.BrowseNum),
		}
		if len(blog.Edges.Tag) > 0 {
			for _, v := range blog.Edges.Tag {
				resp.Tags = append(resp.Tags, v.Name)
			}
		}
		data = append(data, resp)
	}
	return int64(total), data, err
}

func (o *blogsRepo) HotBlogs(ctx context.Context, userId, limit int) ([]*biz.Blogs, error) {
	tx := o.data.db.Blogs.Query().Where(blogs.Or(
		blogs.AccountIDEQ(userId),
		blogs.IsHiddenEQ(constx.BlogNotHidden),
	))
	tx = tx.Where(blogs.LoveNumGTE(constx.BlogHotNum)).Where(blogs.UpdatedAtGTE(time.Now().Add(-(constx.RecentSevenDay * time.Second)).Unix()))
	list, err := tx.Limit(limit).Order(ent.Desc(blogs.FieldLoveNum)).WithTag().All(ctx)
	if err != nil {
		return nil, err
	}
	data := make([]*biz.Blogs, 0, len(list))
	var existIds []int
	for _, blog := range list {
		existIds = append(existIds, int(blog.ID))
		resp := &biz.Blogs{
			CreatedAt:   blog.CreatedAt,
			UpdatedAt:   blog.UpdatedAt,
			Id:          int64(blog.ID),
			IsHidden:    int64(blog.IsHidden),
			AccountId:   int64(blog.AccountID),
			Title:       blog.Title,
			Description: blog.Description,
			//Tags:        blog.Tags,
			Cover:      blog.Cover,
			ThumbNum:   int64(blog.LoveNum),
			CollectNum: int64(blog.CollectNum),
			BrowseNum:  int64(blog.BrowseNum),
		}
		if len(blog.Edges.Tag) > 0 {
			for _, v := range blog.Edges.Tag {
				resp.Tags = append(resp.Tags, v.Name)
			}
		}
		data = append(data, resp)
	}
	if len(data) == limit {
		return data, nil
	}
	tx = o.data.db.Blogs.Query().Where(blogs.Or(
		blogs.AccountIDEQ(userId),
		blogs.IsHiddenEQ(constx.BlogNotHidden),
	)).Where(blogs.IDNotIn(existIds...))
	list, err = tx.Limit(limit-len(data)).Order(ent.Desc(blogs.FieldBrowseNum), ent.Desc(blogs.FieldID)).WithTag().All(ctx)
	if err != nil {
		return nil, err
	}
	for _, blog := range list {
		resp := &biz.Blogs{
			CreatedAt:   blog.CreatedAt,
			UpdatedAt:   blog.UpdatedAt,
			Id:          int64(blog.ID),
			IsHidden:    int64(blog.IsHidden),
			AccountId:   int64(blog.AccountID),
			Title:       blog.Title,
			Description: blog.Description,
			//Tags:        blog.Tags,
			Cover:      blog.Cover,
			ThumbNum:   int64(blog.LoveNum),
			CollectNum: int64(blog.CollectNum),
			BrowseNum:  int64(blog.BrowseNum),
		}
		if len(blog.Edges.Tag) > 0 {
			for _, v := range blog.Edges.Tag {
				resp.Tags = append(resp.Tags, v.Name)
			}
		}
		data = append(data, resp)
	}
	return data, nil
}

func (o *blogsRepo) Thumb(ctx context.Context, g *global.Action) (err error) {
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
	id, err := tx.BlogsExtend.Create().
		SetAccountID(int(g.GetAccountId())).SetBlogID(int(g.Id)).
		SetLoveAt(time.Now().Unix()).SetLove(g.Do).
		OnConflictColumns(blogsextend.FieldAccountID, blogsextend.FieldBlogID).
		UpdateLove().UpdateLoveAt().ID(ctx)
	if err != nil {
		return
	}
	if id == 0 {
		return errors.New("record not found")
	}
	i := 1
	if !g.Do {
		i = -1
	}
	err = tx.Blogs.UpdateOneID(int(g.Id)).AddLoveNum(i).Exec(ctx)
	return err
}

func (o *blogsRepo) Collect(ctx context.Context, g *global.Action) (err error) {
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
	id, err := tx.BlogsExtend.Create().
		SetAccountID(int(g.GetAccountId())).SetBlogID(int(g.Id)).
		SetCollectAt(time.Now().Unix()).SetCollect(g.Do).
		OnConflictColumns(blogsextend.FieldAccountID, blogsextend.FieldBlogID).
		UpdateCollect().UpdateCollectAt().ID(ctx)
	if err != nil {
		return
	}
	if id == 0 {
		return errors.New("record not found")
	}
	i := 1
	if !g.Do {
		i = -1
	}
	err = tx.Blogs.UpdateOneID(int(g.Id)).AddCollectNum(i).Exec(ctx)
	return err
}

func (o *blogsRepo) ListBlogTags(ctx context.Context, query *biz.BlogsQuery) (int64, []biz.BlogTags, error) {
	tx := o.data.db.TagsRelation.Query().Where(tagsrelation.RelationEQ(constx.Blog))
	var scopes []predicate.Blogs
	scopes = append(scopes, blogs.Or(
		blogs.AccountIDEQ(int(query.AccountId)),
		blogs.IsHiddenEQ(constx.BlogNotHidden),
	))
	if query.Title != "" {
		scopes = append(scopes, blogs.TitleContains(query.Title))
	}
	if query.Mine {
		scopes = append(scopes, blogs.AccountIDEQ(int(query.AccountId)))
	}

	tx = tx.Where(tagsrelation.HasBlogWith(scopes...))
	if len(query.Tags) > 0 {
		tx = tx.Where(tagsrelation.HasBlogWith(blogs.HasTagRelationWith(
			tagsrelation.RelationEQ(constx.Blog),
			tagsrelation.HasTagWith(
				tags.NameIn(query.Tags...),
			),
		)))
	}
	//datax, err := tx.GroupBy(tagsrelation.FieldTagID).Aggregate(ent.Count()).Int(ctx)
	//fmt.Println(datax, err)
	total, err := tx.Clone().Select(tagsrelation.FieldTagID).Unique(true).Count(ctx)
	if err != nil {
		return 0, nil, err
	}
	var aggrates []ent.AggregateFunc
	//aggrates = append(aggrates, func(selector *sql.Selector) string {
	//	return ""
	//})
	//aggrates = append(aggrates, func(selector *sql.Selector) string {
	//	return ""
	//})
	aggrates = append(aggrates, func(s *sql.Selector) string {
		t := sql.Table(tags.Table)
		return s.Join(t).On(t.C(tags.FieldID), s.C(tagsrelation.FieldTagID)).
			Select(tagsrelation.FieldTagID, tags.FieldName, "count(*) as count").String()
	})
	resp := make([]biz.BlogTags, 0, query.PageSize)
	err = tx.Order(func(selector *sql.Selector) {
		selector.OrderBy("count(*) desc,tag_id desc")
	}).Limit(int(query.PageSize)).
		GroupBy(tagsrelation.FieldTagID).Aggregate(aggrates...).Scan(ctx, &resp)
	return int64(total), resp, err
}
