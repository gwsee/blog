package biz

import (
	"blog/api/global"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// Travel is a Travel model.
type Travel struct {
	TravelBase
	TravelAccount
	TravelExtend
}

type TravelBase struct {
	Id          int
	Title       string
	IsHidden    bool
	Description string
	Video       string
	Photos      []string
	AccountId   int
	CreatedAt   int64
	UpdatedAt   int64
}

type TravelAccount struct {
	Nickname string
	Avatar   string
}

type TravelExtend struct {
	IsThumb    bool
	ThumbNum   int
	IsCollect  bool
	CollectNum int
	BrowseNum  int
}
type TravelDo struct {
	Id        int
	AccountId int
	Do        bool //是否收藏 或者 是否点赞
}
type TravelQuery struct {
	*global.PageInfo
	Title       string
	AccountId   int64 //当前登陆人ID
	Description string
	My          bool
	MyCollect   bool
	MyThumb     bool
	Sort        string
}

// TravelRepo is a Greater repo.
type TravelRepo interface {
	Save(context.Context, *TravelBase) error
	Delete(context.Context, *TravelBase) error
	Thumb(context.Context, *TravelDo) error
	Collect(context.Context, *TravelDo) error
	Get(context.Context, *TravelBase) (*Travel, error)
	List(context.Context, *TravelQuery) (int64, []*Travel, error)
}

// TravelUsecase is a Travel usecase.
type TravelUsecase struct {
	repo TravelRepo
	log  *log.Helper
}

// NewTravelUsecase new a Travel usecase.
func NewTravelUsecase(repo TravelRepo, logger log.Logger) *TravelUsecase {
	return &TravelUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateTravel creates a Travel, and returns the new Travel.
func (uc *TravelUsecase) CreateTravel(ctx context.Context, g *Travel) (*Travel, error) {

}
