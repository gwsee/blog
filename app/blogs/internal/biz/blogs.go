package biz

import (
	pb "blog/api/blogs/v1"
	"blog/api/global"
	"blog/internal/constx"
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
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
}

type BlogsQuery struct {
	*global.PageInfo
	Title     string
	AccountId int64 //当前登陆人ID
	Tags      []string
}

type BlogsRepo interface {
	CreateBlogs(context.Context, *Blogs) error
	UpdateBlogs(context.Context, *Blogs) error
	DeleteBlogs(context.Context, *Blogs) error
	GetBlogs(context.Context, *Blogs) (*Blogs, error)
	ListBlogs(context.Context, *BlogsQuery) (int64, []*Blogs, error)
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
func (s *BlogsUseCase) CreateBlogs(ctx context.Context, req *pb.CreateBlogsRequest) (*global.Empty, error) {
	if req.Title == "" || req.Content == "" {
		return nil, errors.New("title or content is empty")
	}
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	return &global.Empty{}, s.repo.CreateBlogs(ctx, &Blogs{
		IsHidden:    int64(req.IsHidden),
		AccountId:   u.Id,
		Title:       req.Title,
		Description: req.Description,
		Tags:        req.Tags,
		Cover:       req.Cover,
		Content:     req.Content,
	})
}
func (s *BlogsUseCase) UpdateBlogs(ctx context.Context, req *pb.UpdateBlogsRequest) (*global.Empty, error) {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}

	return &global.Empty{}, s.repo.UpdateBlogs(ctx, &Blogs{
		Id:          req.Id,
		IsHidden:    int64(req.IsHidden),
		Title:       req.Title,
		Description: req.Description,
		AccountId:   u.Id,
		Tags:        req.Tags,
		Cover:       req.Cover,
		Content:     req.Content,
	})
}
func (s *BlogsUseCase) DeleteBlogs(ctx context.Context, req *global.ID) (*global.Empty, error) {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	return &global.Empty{}, s.repo.DeleteBlogs(ctx, &Blogs{
		Id:        req.ID,
		AccountId: u.Id,
	})
}
func (s *BlogsUseCase) GetBlogs(ctx context.Context, req *global.ID) (*pb.GetBlogsReply, error) {
	u := constx.DefaultUser.Default(ctx)
	info, err := s.repo.GetBlogs(ctx, &Blogs{
		Id:        req.ID,
		AccountId: u.Id,
	})
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
		},
		Content: info.Content,
	}, nil
}
func (s *BlogsUseCase) ListBlogs(ctx context.Context, req *pb.ListBlogsRequest) (*pb.ListBlogsReply, error) {
	user := constx.DefaultUser.Default(ctx)
	total, list, err := s.repo.ListBlogs(ctx, &BlogsQuery{
		PageInfo:  req.Page,
		Title:     req.Title,
		AccountId: user.Id,
		Tags:      req.Tags,
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
		})
	}
	return res, nil
}
