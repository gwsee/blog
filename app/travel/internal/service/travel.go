package service

import (
	"blog/api/travel/v1"
	"blog/app/travel/internal/biz"
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
