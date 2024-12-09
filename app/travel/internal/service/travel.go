package service

import (
	"blog/api/global"
	"blog/api/travel/v1"
	"blog/app/travel/internal/biz"
	"context"
)

// TravelService is a travel service.
type TravelService struct {
	v1.UnimplementedTravelServer

	uc *biz.TravelUsecase
}

// NewTravelService new a travel service.
func NewTravelService(uc *biz.TravelUsecase) *TravelService {
	return &TravelService{uc: uc}
}
func (s *TravelService) SaveTravel(ctx context.Context, req *v1.SaveTravelRequest) (*global.Empty, error) {
	return s.uc.Save(ctx, req)
}
func (s *TravelService) GetTravel(ctx context.Context, req *global.ID) (*v1.GetTravelReply, error) {
	return s.uc.Get(ctx, req)
}
func (s *TravelService) ListTravel(ctx context.Context, req *v1.ListTravelRequest) (*v1.ListTravelReply, error) {
	return s.uc.List(ctx, req)
}
func (s *TravelService) Thumb(ctx context.Context, req *global.Action) (*global.Empty, error) {
	return s.uc.Thumb(ctx, req)
}
func (s *TravelService) DeleteTravel(ctx context.Context, req *global.ID) (*global.Empty, error) {
	return s.uc.Delete(ctx, req)
}
func (s *TravelService) Collect(ctx context.Context, req *global.Action) (*global.Empty, error) {
	return s.uc.Collect(ctx, req)
}
