package data

import (
	"blog/api/global"
	v1 "blog/api/travel/v1"
	"blog/app/bff/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type travelRepo struct {
	data *Data
	log  *log.Helper
}

func NewTravelRepo(data *Data, logger log.Logger) biz.TravelRepo {
	return &travelRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (o *travelRepo) SaveTravel(ctx context.Context, in *v1.SaveTravelRequest) (*global.Empty, error) {
	return o.data.tc.SaveTravel(ctx, in)
}
func (o *travelRepo) DeleteTravel(ctx context.Context, in *global.ID) (*global.Empty, error) {
	return o.data.tc.DeleteTravel(ctx, in)
}
func (o *travelRepo) GetTravel(ctx context.Context, in *global.ID) (*v1.GetTravelReply, error) {
	return o.data.tc.GetTravel(ctx, in)
}
func (o *travelRepo) ListTravel(ctx context.Context, in *v1.ListTravelRequest) (*v1.ListTravelReply, error) {
	return o.data.tc.ListTravel(ctx, in)
}
func (o *travelRepo) Thumb(ctx context.Context, in *global.Action) (*global.Empty, error) {
	return o.data.tc.Thumb(ctx, in)
}
func (o *travelRepo) Collect(ctx context.Context, in *global.Action) (*global.Empty, error) {
	return o.data.tc.Collect(ctx, in)
}
