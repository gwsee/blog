package biz

import (
	pb "blog/api/blogs/v1"
	"blog/api/global"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type BlogsComment struct {
	Id        int64
	BlogId    int64
	TopId     int64
	ParentId  int64
	Level     int64
	Total     int64
	Status    int64
	Content   string
	AccountId int64
	CreatedAt int64
}

type BlogsCommentQuery struct {
	*global.PageIndex
	BlogId    int64
	ParentId  int64
	Content   string
	AccountId int64
}

type BlogsCommentRepo interface {
	CreateBlogsComment(context.Context, *BlogsComment) error
	UpdateBlogsComment(context.Context, *BlogsComment) error
	DeleteBlogsComment(context.Context, int64) error
	GetBlogsComment(context.Context, int64) (*BlogsComment, error)
	ListBlogsComment(context.Context, *BlogsCommentQuery) (int64, []*BlogsComment, error)
}

type BlogsCommentUseCase struct {
	repo BlogsCommentRepo
	log  *log.Helper
}

func NewBlogsCommentUseCase(repo BlogsCommentRepo, logger log.Logger) *BlogsCommentUseCase {
	return &BlogsCommentUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "app/blogs-comment-service")),
	}
}

func (s *BlogsCommentUseCase) CreateBlogsComment(ctx context.Context, req *pb.CreateBlogsCommentRequest) (*global.Empty, error) {
	return &global.Empty{}, s.repo.CreateBlogsComment(ctx, &BlogsComment{
		BlogId:    req.BlogId,
		TopId:     req.TopId,
		ParentId:  req.ParentId,
		Level:     req.Level,
		Total:     req.Total,
		Status:    req.Status,
		Content:   req.Content,
		AccountId: 0,
	})
}
func (s *BlogsCommentUseCase) UpdateBlogsComment(ctx context.Context, req *pb.UpdateBlogsCommentRequest) (*global.Empty, error) {
	return &global.Empty{}, s.repo.UpdateBlogsComment(ctx, &BlogsComment{
		Id:      req.Id,
		Status:  req.Status,
		Content: req.Content,
	})
}
func (s *BlogsCommentUseCase) DeleteBlogsComment(ctx context.Context, req *global.ID) (*global.Empty, error) {
	return &global.Empty{}, s.repo.DeleteBlogsComment(ctx, req.ID)
}
func (s *BlogsCommentUseCase) GetBlogsComment(ctx context.Context, req *global.ID) (*pb.GetBlogsCommentReply, error) {
	comment, err := s.repo.GetBlogsComment(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &pb.GetBlogsCommentReply{
		Id:        comment.Id,
		BlogId:    comment.BlogId,
		TopId:     comment.TopId,
		ParentId:  comment.ParentId,
		Level:     comment.Level,
		Total:     comment.Total,
		Status:    comment.Status,
		Content:   comment.Content,
		AccountId: comment.AccountId,
		CreatedAt: comment.CreatedAt,
	}, nil
}
func (s *BlogsCommentUseCase) ListBlogsComment(ctx context.Context, req *pb.ListBlogsCommentRequest) (*pb.ListBlogsCommentReply, error) {
	end, list, err := s.repo.ListBlogsComment(ctx, &BlogsCommentQuery{
		PageIndex: req.Page,
		BlogId:    req.BlogId,
		ParentId:  req.ParentId,
		Content:   req.Content,
		AccountId: req.AccountId,
	})
	if err != nil {
		return nil, err
	}
	info := &pb.ListBlogsCommentReply{
		End:  end,
		List: make([]*pb.GetBlogsCommentReply, 0, len(list)),
	}
	for _, comment := range list {
		info.List = append(info.List, &pb.GetBlogsCommentReply{
			Id:        comment.Id,
			BlogId:    comment.BlogId,
			TopId:     comment.TopId,
			ParentId:  comment.ParentId,
			Level:     comment.Level,
			Total:     comment.Total,
			Status:    comment.Status,
			Content:   comment.Content,
			AccountId: comment.AccountId,
			CreatedAt: comment.CreatedAt,
		})
	}
	return info, nil
}
