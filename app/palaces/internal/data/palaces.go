package data

import (
	"blog/api/global"
	"blog/app/palaces/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type palacesRepo struct {
	data *Data
	log  *log.Helper
}

// NewPalacesRepo .
func NewPalacesRepo(data *Data, logger log.Logger) biz.PalacesRepo {
	return &palacesRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (p palacesRepo) CreatePalaces(ctx context.Context, palaces *biz.Palaces) error {
	//TODO implement me
	panic("implement me")
}

func (p palacesRepo) UpdatePalaces(ctx context.Context, palaces *biz.Palaces) error {
	//TODO implement me
	panic("implement me")
}

func (p palacesRepo) DeletePalaces(ctx context.Context, id *global.ID) error {
	//TODO implement me
	panic("implement me")
}

func (p palacesRepo) GetPalaces(ctx context.Context, id *global.ID) (*biz.Palaces, error) {
	//TODO implement me
	panic("implement me")
}

func (p palacesRepo) ListPalaces(ctx context.Context, palaces *biz.Palaces) (int64, []*biz.Palaces, error) {
	//TODO implement me
	panic("implement me")
}
