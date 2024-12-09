package service

import (
	pb "blog/api/bff/v1"
	"blog/api/global"
	v1 "blog/api/travel/v1"
	"blog/app/bff/internal/biz"
	"context"
)

type TravelService struct {
	pb.UnimplementedTravelsServer
	biz *biz.TravelUseCase
}

func NewTravelService(biz *biz.TravelUseCase) *TravelService {
	return &TravelService{biz: biz}
}
func (o *TravelService) SaveTravel(ctx context.Context, req *v1.SaveTravelRequest) (*global.Empty, error) {
	return o.biz.SaveTravel(ctx, req)
}
func (o *TravelService) DeleteTravel(ctx context.Context, req *global.ID) (*global.Empty, error) {
	return o.biz.DeleteTravel(ctx, req)
}
func (o *TravelService) GetTravel(ctx context.Context, req *global.ID) (*v1.GetTravelReply, error) {
	return o.biz.GetTravel(ctx, req)
}
func (o *TravelService) Thumb(ctx context.Context, req *global.Action) (*global.Empty, error) {
	return o.biz.Thumb(ctx, req)
}
func (o *TravelService) Collect(ctx context.Context, req *global.Action) (*global.Empty, error) {
	return o.biz.Collect(ctx, req)
}
func (o *TravelService) ListTravel(ctx context.Context, req *v1.ListTravelRequest) (*v1.ListTravelReply, error) {
	return o.biz.ListTravel(ctx, req)
}
