package biz

import (
	"blog/api/global"
	v1 "blog/api/travel/v1"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type TravelUseCase struct {
	repo TravelRepo
	log  *log.Helper
}
type TravelRepo interface {
	SaveTravel(context.Context, *v1.SaveTravelRequest) (*global.Empty, error)
	DeleteTravel(context.Context, *global.ID) (*global.Empty, error)
	GetTravel(context.Context, *global.ID) (*v1.GetTravelReply, error)
	ListTravel(context.Context, *v1.ListTravelRequest) (*v1.ListTravelReply, error)
	Thumb(context.Context, *global.Action) (*global.Empty, error)
	Collect(context.Context, *global.Action) (*global.Empty, error)
}

func NewTravelUseCase(repo TravelRepo, logger log.Logger) *TravelUseCase {
	return &TravelUseCase{repo: repo, log: log.NewHelper(logger)}
}
func (o *TravelUseCase) SaveTravel(ctx context.Context, req *v1.SaveTravelRequest) (*global.Empty, error) {
	return o.repo.SaveTravel(ctx, req)
}
func (o *TravelUseCase) DeleteTravel(ctx context.Context, req *global.ID) (*global.Empty, error) {
	return o.repo.DeleteTravel(ctx, req)
}
func (o *TravelUseCase) GetTravel(ctx context.Context, req *global.ID) (*v1.GetTravelReply, error) {
	return o.repo.GetTravel(ctx, req)
}
func (o *TravelUseCase) Thumb(ctx context.Context, req *global.Action) (*global.Empty, error) {
	return o.repo.Thumb(ctx, req)
}
func (o *TravelUseCase) Collect(ctx context.Context, req *global.Action) (*global.Empty, error) {
	return o.repo.Collect(ctx, req)
}
func (o *TravelUseCase) ListTravel(ctx context.Context, req *v1.ListTravelRequest) (*v1.ListTravelReply, error) {
	return o.repo.ListTravel(ctx, req)
}
