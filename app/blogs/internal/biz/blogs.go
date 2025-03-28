package biz

import (
	pb "blog/api/blogs/v1"
	"blog/api/global"
	"blog/internal/common/array"
	"blog/internal/constx"
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"strings"
)

type Blogs struct {
	CreatedAt   int64
	UpdatedAt   int64
	Id          int64
	IsHidden    int64
	AccountId   int64
	Title       string
	Description string
	Tags        []string
	Cover       string
	Content     string
	Files       []string
	BrowseNum   int64
	ThumbNum    int64
	CollectNum  int64
}

type BlogTags struct {
	TagId int64  `json:"tag_id"`
	Name  string `json:"name"`
	Count int64  `json:"count"`
}

type BlogsQuery struct {
	*global.PageInfo
	Title     string
	AccountId int64 //当前登陆人ID
	Tags      []string
	Mine      bool
}

type BlogsRepo interface {
	CreateBlogs(context.Context, *Blogs) error
	UpdateBlogs(context.Context, *Blogs) error
	DeleteBlogs(context.Context, *global.ID) error
	GetBlogs(context.Context, *global.ID) (*Blogs, error)
	ListBlogs(context.Context, *BlogsQuery) (int64, []*Blogs, error)
	HotBlogs(context.Context, int, int) ([]*Blogs, error)
	Thumb(context.Context, *global.Action) error
	Collect(context.Context, *global.Action) error
	ListBlogTags(context.Context, *BlogsQuery) (int64, []BlogTags, error)
}

type BlogsUseCase struct {
	repo BlogsRepo
	log  *log.Helper
}

func NewBlogsUseCase(repo BlogsRepo, logger log.Logger) *BlogsUseCase {
	return &BlogsUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "app/blogs-blogs-service")),
	}
}
func (s *BlogsUseCase) validData(data *Blogs) error {
	data.Title = strings.TrimSpace(data.Title)
	data.Content = strings.TrimSpace(data.Content)
	data.Description = strings.TrimSpace(data.Description)
	var tags []string
	for _, tag := range data.Tags {
		it := strings.ToLower(strings.TrimSpace(tag))
		if it == "" {
			continue
		}
		if array.In(it, tags) {
			continue
		}
		tags = append(tags, it)
	}
	data.Tags = tags
	if data.Title == "" || data.Content == "" {
		return errors.New("title or content is empty")
	}
	if len(data.Tags) == 0 {
		return errors.New("tags is empty")
	}
	return nil
}
func (s *BlogsUseCase) CreateBlogs(ctx context.Context, req *pb.CreateBlogsRequest) (*global.Empty, error) {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	data := &Blogs{
		IsHidden:    int64(req.IsHidden),
		AccountId:   u.Id,
		Title:       req.Title,
		Description: req.Description,
		Tags:        req.Tags,
		Cover:       req.Cover,
		Content:     req.Content,
		Files:       req.Files,
	}
	if err = s.validData(data); err != nil {
		return nil, err
	}
	return &global.Empty{}, s.repo.CreateBlogs(ctx, data)
}
func (s *BlogsUseCase) UpdateBlogs(ctx context.Context, req *pb.UpdateBlogsRequest) (*global.Empty, error) {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	data := &Blogs{
		Id:          req.Id,
		IsHidden:    int64(req.IsHidden),
		Title:       req.Title,
		Description: req.Description,
		AccountId:   u.Id,
		Tags:        req.Tags,
		Cover:       req.Cover,
		Content:     req.Content,
		Files:       req.Files,
	}
	if err = s.validData(data); err != nil {
		return nil, err
	}
	return &global.Empty{}, s.repo.UpdateBlogs(ctx, data)
}
func (s *BlogsUseCase) DeleteBlogs(ctx context.Context, req *global.ID) (*global.Empty, error) {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	req.AccountId = &u.Id
	return &global.Empty{}, s.repo.DeleteBlogs(ctx, req)
}
func (s *BlogsUseCase) GetBlogs(ctx context.Context, req *global.ID) (*pb.GetBlogsReply, error) {
	u := constx.DefaultUser.Default(ctx)
	req.AccountId = &u.Id
	info, err := s.repo.GetBlogs(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.GetBlogsReply{
		Header: &pb.BlogsHeader{
			Id:          info.Id,
			Title:       info.Title,
			Description: info.Description,
			IsHidden:    info.IsHidden,
			Tags:        info.Tags,
			Cover:       info.Cover,
			AccountId:   info.AccountId,
			CreatedAt:   info.CreatedAt,
			UpdatedAt:   info.UpdatedAt,
			BrowseNum:   info.BrowseNum,
			CollectNum:  info.CollectNum,
			ThumbNum:    info.ThumbNum,
		},
		Content: info.Content,
		Files:   info.Files,
	}, nil
}
func (s *BlogsUseCase) ListBlogs(ctx context.Context, req *pb.ListBlogsRequest) (*pb.ListBlogsReply, error) {
	user := constx.DefaultUser.Default(ctx)
	total, list, err := s.repo.ListBlogs(ctx, &BlogsQuery{
		PageInfo:  req.Page,
		Title:     req.Title,
		AccountId: user.Id,
		Tags:      req.Tags,
		Mine:      req.Mine,
	})
	if err != nil {
		return nil, err
	}
	res := &pb.ListBlogsReply{
		Total: total,
		List:  make([]*pb.BlogsHeader, 0, len(list)),
	}
	for _, info := range list {
		res.List = append(res.List, &pb.BlogsHeader{
			Id:          info.Id,
			Title:       info.Title,
			Description: info.Description,
			IsHidden:    info.IsHidden,
			Tags:        info.Tags,
			Cover:       info.Cover,
			AccountId:   info.AccountId,
			CreatedAt:   info.CreatedAt,
			UpdatedAt:   info.UpdatedAt,
			BrowseNum:   info.BrowseNum,
			CollectNum:  info.CollectNum,
			ThumbNum:    info.ThumbNum,
		})
	}
	return res, nil
}
func (s *BlogsUseCase) HotBlogs(ctx context.Context, req *global.PageInfo) (*pb.ListBlogsReply, error) {
	user := constx.DefaultUser.Default(ctx)
	list, err := s.repo.HotBlogs(ctx, int(user.Id), int(req.PageSize))
	if err != nil {
		return nil, err
	}
	res := &pb.ListBlogsReply{
		List: make([]*pb.BlogsHeader, 0, len(list)),
	}
	for _, info := range list {
		res.List = append(res.List, &pb.BlogsHeader{
			Id:          info.Id,
			Title:       info.Title,
			Description: info.Description,
			IsHidden:    info.IsHidden,
			Tags:        info.Tags,
			Cover:       info.Cover,
			AccountId:   info.AccountId,
			CreatedAt:   info.CreatedAt,
			UpdatedAt:   info.UpdatedAt,
			BrowseNum:   info.BrowseNum,
			CollectNum:  info.CollectNum,
			ThumbNum:    info.ThumbNum,
		})
	}
	return res, nil
}
func (s *BlogsUseCase) Thumb(ctx context.Context, req *global.Action) (*global.Empty, error) {
	user, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	req.AccountId = &user.Id
	return &global.Empty{}, s.repo.Thumb(ctx, req)
}
func (s *BlogsUseCase) Collect(ctx context.Context, req *global.Action) (*global.Empty, error) {
	user, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	req.AccountId = &user.Id
	return &global.Empty{}, s.repo.Collect(ctx, req)
}
func (s *BlogsUseCase) ListBlogTags(ctx context.Context, req *pb.ListBlogsRequest) (*pb.ListBlogTagsReply, error) {
	user := constx.DefaultUser.Default(ctx)
	total, list, err := s.repo.ListBlogTags(ctx, &BlogsQuery{
		PageInfo:  req.Page,
		Title:     req.Title,
		AccountId: user.Id,
		Tags:      req.Tags,
	})
	if err != nil {
		return nil, err
	}
	res := &pb.ListBlogTagsReply{
		Total: total,
	}
	for _, info := range list {
		res.Tags = append(res.Tags, &pb.BlogTags{
			Name: info.Name,
			Num:  info.Count,
		})
	}
	return res, nil
}
