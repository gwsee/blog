package biz

import (
	pb "blog/api/account/v1"
	"blog/api/global"
	v1 "blog/api/travel/v1"
	"blog/internal/constx"
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// Travel is a Travel model.
type Travel struct {
	TravelBase
	TravelAccount
	TravelExtend
}

type TravelBase struct {
	Id          int      `json:"id"`
	Title       string   `json:"title"`
	IsHidden    bool     `json:"is_hidden"`
	Description string   `json:"description"`
	Video       string   `json:"video"`
	Photos      []string `json:"photos"`
	AccountId   int      `json:"account_id"`
	CreatedAt   int64    `json:"created_at"`
	UpdatedAt   int64    `json:"updated_at"`
}

type TravelAccount struct {
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
}

type TravelExtend struct {
	IsThumb    bool `json:"is_thumb"`
	ThumbNum   int  `json:"thumb_num"`
	IsCollect  bool `json:"is_collect"`
	CollectNum int  `json:"collect_num"`
	BrowseNum  int  `json:"browse_num"`
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
func (o *TravelUsecase) Save(ctx context.Context, req *v1.SaveTravelRequest) (*global.Empty, error) {
	if req.Title == "" {
		return nil, errors.New("title is empty")
	}
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	return &global.Empty{}, o.repo.Save(ctx, &TravelBase{
		Id:          int(req.Id),
		Title:       req.Title,
		IsHidden:    req.IsHidden,
		Description: req.Description,
		Video:       req.Video,
		Photos:      req.Photo,
		AccountId:   int(u.Id),
	})
}
func (o *TravelUsecase) Get(ctx context.Context, req *global.ID) (*v1.GetTravelReply, error) {
	u := constx.DefaultUser.Default(ctx)
	info, err := o.repo.Get(ctx, &TravelBase{
		Id:        int(req.Id),
		AccountId: int(u.Id),
	})
	if err != nil {
		return nil, err
	}
	return &v1.GetTravelReply{
		Id:          int64(info.Id),
		Title:       info.Title,
		IsHidden:    info.IsHidden,
		Description: info.TravelBase.Description,
		Video:       info.Video,
		Photos:      info.Photos,
		BrowseNum:   int64(info.BrowseNum),
		ThumbNum:    int64(info.ThumbNum),
		IsThumb:     info.IsThumb,
		CollectNum:  int64(info.CollectNum),
		IsCollect:   info.IsCollect,
		AccountId:   int64(info.AccountId),
		Account: &pb.AccountInfoReply{
			Id:          int64(info.AccountId),
			Nickname:    info.Nickname,
			Avatar:      info.Avatar,
			Description: info.TravelAccount.Description,
		},
		CreatedAt: info.CreatedAt,
		UpdatedAt: info.UpdatedAt,
	}, nil
}
func (o *TravelUsecase) List(ctx context.Context, req *v1.ListTravelRequest) (*v1.ListTravelReply, error) {
	u := constx.DefaultUser.Default(ctx)
	total, list, err := o.repo.List(ctx, &TravelQuery{
		PageInfo:    req.Page,
		Title:       req.Title,
		AccountId:   u.Id,
		Description: req.Description,
		My:          req.My,
		MyCollect:   req.MyCollect,
		MyThumb:     req.MyThumb,
		Sort:        req.Sort,
	})
	if err != nil {
		return nil, err
	}
	resp := &v1.ListTravelReply{
		Total: total,
		List:  make([]*v1.ListTravel, 0, len(list)),
	}
	_rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	for _, info := range list {
		one := &v1.ListTravel{
			Id:          int64(info.Id),
			Title:       info.Title,
			Description: info.TravelBase.Description,
			Video:       info.Video,
			BrowseNum:   int64(info.BrowseNum),
			ThumbNum:    int64(info.ThumbNum),
			IsThumb:     info.IsThumb,
			CollectNum:  int64(info.CollectNum),
			IsCollect:   info.IsCollect,
			AccountId:   int64(info.AccountId),
			Account: &pb.AccountInfoReply{
				Id:       int64(info.AccountId),
				Nickname: info.Nickname,
				Avatar:   info.Avatar,
			},
			CreatedAt: info.CreatedAt,
			UpdatedAt: info.UpdatedAt,
		}
		if len(info.Photos) > 0 {
			index := _rand.Intn(len(info.Photos))
			one.Photo = info.Photos[index]
		}
		resp.List = append(resp.List, one)
	}
	return resp, nil
}
func (o *TravelUsecase) Delete(ctx context.Context, req *global.ID) (*global.Empty, error) {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	return &global.Empty{}, o.repo.Delete(ctx, &TravelBase{
		Id:        int(req.Id),
		AccountId: int(u.Id),
	})
}
func (o *TravelUsecase) Thumb(ctx context.Context, req *global.Action) (*global.Empty, error) {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	return &global.Empty{}, o.repo.Thumb(ctx, &TravelDo{
		Id:        int(req.Id),
		AccountId: int(u.Id),
		Do:        req.GetDo(),
	})
}
func (o *TravelUsecase) Collect(ctx context.Context, req *global.Action) (*global.Empty, error) {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	return &global.Empty{}, o.repo.Collect(ctx, &TravelDo{
		Id:        int(req.Id),
		AccountId: int(u.Id),
		Do:        req.GetDo(),
	})
}
