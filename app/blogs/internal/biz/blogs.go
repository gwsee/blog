package biz

import (
	pb "blog/api/blogs/v1"
	"blog/api/global"
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
)

type Blogs struct {
	CreatedAt   int64
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
	DeleteBlogs(context.Context, int64) error
	GetBlogs(context.Context, int64) (*Blogs, error)
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
func (s *BlogsUseCase) CreateBlogs(ctx context.Context, req *pb.CreateBlogsRequest) (*pb.CreateBlogsReply, error) {
	if req.Title == "" || req.Content == "" {
		return nil, errors.New("title or content is empty")
	}
	//todo 获取当前登陆人信息
	return &pb.CreateBlogsReply{}, s.repo.CreateBlogs(ctx, &Blogs{
		IsHidden:    int64(req.IsHidden),
		AccountId:   0,
		Title:       req.Title,
		Description: req.Description,
		Tags:        req.Tags,
		Cover:       req.Cover,
		Content:     req.Content,
	})
}
func (s *BlogsUseCase) UpdateBlogs(ctx context.Context, req *pb.UpdateBlogsRequest) (*pb.UpdateBlogsReply, error) {
	return &pb.UpdateBlogsReply{}, s.repo.UpdateBlogs(ctx, &Blogs{
		Id:          req.Id,
		IsHidden:    int64(req.IsHidden),
		Title:       req.Title,
		Description: req.Description,
		Tags:        req.Tags,
		Cover:       req.Cover,
		Content:     req.Content,
	})
}
func (s *BlogsUseCase) DeleteBlogs(ctx context.Context, req *pb.DeleteBlogsRequest) (*pb.DeleteBlogsReply, error) {
	return &pb.DeleteBlogsReply{}, s.repo.DeleteBlogs(ctx, req.Id)
}
func (s *BlogsUseCase) GetBlogs(ctx context.Context, req *pb.GetBlogsRequest) (*pb.GetBlogsReply, error) {
	info, err := s.repo.GetBlogs(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.GetBlogsReply{
		Header: &pb.BlogsHeader{
			Id:          info.Id,
			Title:       info.Title,
			Description: info.Description,
			IsHidden:    pb.HiddenType(info.IsHidden),
			Tags:        info.Tags,
			Cover:       info.Cover,
			AccountId:   info.AccountId,
			CreatedAt:   info.CreatedAt,
		},
		Content: info.Content,
	}, nil
}
func (s *BlogsUseCase) ListBlogs(ctx context.Context, req *pb.ListBlogsRequest) (*pb.ListBlogsReply, error) {
	//todo 获取当前登陆人信息
	total, list, err := s.repo.ListBlogs(ctx, &BlogsQuery{
		PageInfo:  req.Page,
		Title:     req.Title,
		AccountId: 0,
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
			IsHidden:    pb.HiddenType(info.IsHidden),
			Tags:        info.Tags,
			Cover:       info.Cover,
			AccountId:   info.AccountId,
			CreatedAt:   info.CreatedAt,
		})
	}
	return res, nil
}
