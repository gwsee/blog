package data

import (
	"blog/internal/ent"
	"blog/internal/ent/account"
	"blog/internal/ent/filesextend"
	"blog/internal/ent/travel"
	"blog/internal/ent/travelextend"
	"context"
	"errors"

	"blog/app/travel/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type travelRepo struct {
	data *Data
	log  *log.Helper
}

// NewTravelRepo .
func NewTravelRepo(data *Data, logger log.Logger) biz.TravelRepo {
	return &travelRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *travelRepo) Save(ctx context.Context, g *biz.TravelBase) (err error) {
	tx, err := r.data.db.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	var travelObj *ent.Travel
	if g.Id == 0 {
		travelObj, err = tx.Travel.Create().SetTitle(g.Title).SetDescription(g.Description).SetVideo(g.Video).
			SetIsHidden(g.IsHidden).SetAccountID(int(g.AccountId)).SetPhotos(g.Photos).Save(ctx)
	} else {
		travelObj, err = tx.Travel.UpdateOneID(int(g.Id)).SetTitle(g.Title).SetDescription(g.Description).SetVideo(g.Video).
			SetIsHidden(g.IsHidden).SetAccountID(int(g.AccountId)).SetPhotos(g.Photos).Save(ctx)
	}
	if err != nil {
		return
	}

	_, err = tx.FilesExtend.Delete().Where(filesextend.FromEQ(travel.Table),
		filesextend.FromIDEQ(int(travelObj.ID)), filesextend.FileIDNotIn(g.Photos...)).Exec(ctx)
	if err != nil {
		return
	}
	var builders []*ent.FilesExtendCreate
	for _, v := range g.Photos {
		builders = append(builders, tx.FilesExtend.Create().
			SetFileID(v).SetFrom(travel.Table).SetFromID(travelObj.ID).SetUserID(int(g.AccountId)).
			SetIsHidden(g.IsHidden).SetDeletedAt(0),
		)
	}
	err = tx.FilesExtend.CreateBulk(
		builders...,
	).OnConflictColumns(filesextend.FieldFromID, filesextend.FieldFrom, filesextend.FieldFileID).Exec(ctx)
	return
}

func (r *travelRepo) Delete(ctx context.Context, g *biz.TravelBase) (err error) {
	tx, err := r.data.db.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	row, err := tx.Travel.Delete().Where(travel.IDEQ(int(g.Id)), travel.AccountIDEQ(int(g.AccountId))).Exec(ctx)
	if err != nil {
		return err
	}
	if row == 0 {
		return errors.New("travel not found")
	}
	_, err = tx.FilesExtend.Delete().Where(filesextend.FromEQ(travel.Table),
		filesextend.FromIDEQ(int(g.Id))).Exec(ctx)
	return
}
func (r *travelRepo) Thumb(ctx context.Context, g *biz.TravelDo) (err error) {
	tx, err := r.data.db.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	id, err := tx.TravelExtend.Create().SetAccountID(int(g.AccountId)).SetTravelID(int(g.Id)).
		SetIsThumb(g.Do).
		OnConflictColumns(travelextend.FieldAccountID, travelextend.FieldTravelID).
		UpdateIsThumb().ID(ctx)
	if err != nil {
		return
	}
	if id == 0 {
		return errors.New("travel extend not found")
	}
	i := 1
	if !g.Do {
		i = -1
	}
	_, err = tx.Travel.UpdateOneID(int(g.Id)).AddThumbNum(i).Save(ctx)
	return
}
func (r *travelRepo) Collect(ctx context.Context, g *biz.TravelDo) (err error) {
	tx, err := r.data.db.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	id, err := tx.TravelExtend.Create().SetAccountID(int(g.AccountId)).SetTravelID(int(g.Id)).
		SetIsCollect(g.Do).
		OnConflictColumns(travelextend.FieldAccountID, travelextend.FieldTravelID).
		UpdateIsCollect().
		ID(ctx)
	if err != nil {
		return
	}
	if id == 0 {
		return errors.New("travel extend not found")
	}
	i := 1
	if !g.Do {
		i = -1
	}
	_, err = tx.Travel.UpdateOneID(int(g.Id)).AddCollectNum(i).Save(ctx)
	return
}

func (r *travelRepo) Get(ctx context.Context, g *biz.TravelBase) (resp *biz.Travel, err error) {
	info, err := r.data.db.Travel.Query().Where(travel.IDEQ(int(g.Id))).
		Where(travel.Or(travel.AccountIDEQ(int(g.AccountId)), travel.IsHidden(false))).
		WithAccountTravel(func(query *ent.AccountQuery) {
			query.Select(account.FieldNickname, account.FieldAvatar)
		}).
		Only(ctx)
	if err != nil {
		return
	}
	resp = &biz.Travel{
		TravelBase: biz.TravelBase{
			Id:          info.ID,
			Title:       info.Title,
			IsHidden:    info.IsHidden,
			Description: info.Description,
			Video:       info.Video,
			Photos:      info.Photos,
			AccountId:   info.AccountID,
			CreatedAt:   info.CreatedAt,
			UpdatedAt:   info.UpdatedAt,
		},
		TravelAccount: biz.TravelAccount{},
		TravelExtend: biz.TravelExtend{
			BrowseNum:  info.BrowseNum,
			ThumbNum:   info.ThumbNum,
			CollectNum: info.CollectNum,
		},
	}
	if info.Edges.AccountTravel != nil {
		resp.Nickname = info.Edges.AccountTravel.Nickname
		resp.Avatar = info.Edges.AccountTravel.Avatar
	}
	if g.AccountId == 0 {
		return
	}
	extend, err := r.data.db.TravelExtend.Query().Where(travelextend.TravelIDEQ(info.ID), travelextend.AccountIDEQ(int(g.AccountId))).Only(ctx)
	if err != nil {
		return resp, nil
	}
	resp.IsCollect = extend.IsCollect
	resp.IsThumb = extend.IsThumb
	return
}

func (r *travelRepo) List(ctx context.Context, g *biz.TravelQuery) (int64, []*biz.Travel, error) {
	tx := r.data.db.Travel.Query().Where(travel.Or(travel.AccountIDEQ(int(g.AccountId)), travel.IsHidden(false)))
	tx.WithAccountTravel(func(query *ent.AccountQuery) {
		query.Select(account.FieldNickname, account.FieldAvatar)
	})
	if g.My {
		tx.Where(travel.AccountIDEQ(int(g.AccountId)))
	}
	if g.MyThumb && g.AccountId > 0 {
		tx.WithTravelExtend(func(query *ent.TravelExtendQuery) {
			query.Select(travelextend.FieldIsThumb).Where(travelextend.AccountIDEQ(int(g.AccountId)), travelextend.IsThumbEQ(true))
		})
	}
	if g.MyCollect && g.AccountId > 0 {
		tx.WithTravelExtend(func(query *ent.TravelExtendQuery) {
			query.Select(travelextend.FieldIsCollect).Where(travelextend.AccountIDEQ(int(g.AccountId)), travelextend.IsCollectEQ(true))
		})
	}
	if len(g.Title) > 0 {
		tx.Where(travel.TitleContains(g.Title))
	}
	if len(g.Description) > 0 {
		tx.Where(travel.DescriptionContains(g.Description))
	}
	if len(g.Sort) > 0 {
		tx = tx.Order(ent.Desc(g.Sort))
	}
	total, err := tx.Count(ctx)
	if err != nil {
		return 0, nil, err
	}
	if total == 0 {
		return 0, []*biz.Travel{}, nil
	}
	tx = tx.Order(ent.Desc(travel.FieldID))
	resp := make([]*biz.Travel, 0, g.GetPageSize())
	err = tx.Select(travel.FieldID, travel.FieldTitle, travel.FieldDescription, travel.FieldVideo, travel.FieldPhotos, travel.FieldBrowseNum,
		travel.FieldThumbNum, travel.FieldCollectNum, travel.FieldAccountID, travel.FieldCreatedAt, travel.FieldCreatedAt,
	).Scan(ctx, &resp)
	return int64(total), resp, err
}
