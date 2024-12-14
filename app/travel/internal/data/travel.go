package data

import (
	"blog/app/travel/internal/biz"
	"blog/internal/ent"
	"blog/internal/ent/account"
	"blog/internal/ent/filesextend"
	"blog/internal/ent/travelextends"
	"blog/internal/ent/travels"
	"context"
	"errors"

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
	var travelObj *ent.Travels
	if g.Id == 0 {
		travelObj, err = tx.Travels.Create().SetTitle(g.Title).SetDescription(g.Description).SetVideo(g.Video).
			SetIsHidden(g.IsHidden).SetAccountID(int(g.AccountId)).SetPhotos(g.Photos).Save(ctx)
	} else {
		travelObj, err = tx.Travels.UpdateOneID(int(g.Id)).SetTitle(g.Title).SetDescription(g.Description).SetVideo(g.Video).
			SetIsHidden(g.IsHidden).SetAccountID(int(g.AccountId)).SetPhotos(g.Photos).Save(ctx)
	}
	if err != nil {
		return
	}

	_, err = tx.FilesExtend.Delete().Where(filesextend.FromEQ(travels.Table),
		filesextend.FromIDEQ(int(travelObj.ID)), filesextend.FileIDNotIn(g.Photos...)).Exec(ctx)
	if err != nil {
		return
	}
	var builders []*ent.FilesExtendCreate
	for _, v := range g.Photos {
		builders = append(builders, tx.FilesExtend.Create().
			SetFileID(v).SetFrom(travels.Table).SetFromID(travelObj.ID).SetUserID(int(g.AccountId)).
			SetIsHidden(g.IsHidden).SetDeletedAt(0),
		)
	}
	err = tx.FilesExtend.CreateBulk(
		builders...,
	).OnConflictColumns(filesextend.FieldFromID, filesextend.FieldFrom, filesextend.FieldFileID).DoNothing().Exec(ctx)
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
	row, err := tx.Travels.Delete().Where(travels.IDEQ(int(g.Id)), travels.AccountIDEQ(int(g.AccountId))).Exec(ctx)
	if err != nil {
		return err
	}
	if row == 0 {
		return errors.New("travel not found")
	}
	_, err = tx.FilesExtend.Delete().Where(filesextend.FromEQ(travels.Table),
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
	id, err := tx.TravelExtends.Create().SetAccountID(int(g.AccountId)).SetTravelID(int(g.Id)).
		SetIsThumb(g.Do).
		OnConflictColumns(travelextends.FieldAccountID, travelextends.FieldTravelID).
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
	_, err = tx.Travels.UpdateOneID(int(g.Id)).AddThumbNum(i).Save(ctx)
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
	id, err := tx.TravelExtends.Create().SetAccountID(int(g.AccountId)).SetTravelID(int(g.Id)).
		SetIsCollect(g.Do).
		OnConflictColumns(travelextends.FieldAccountID, travelextends.FieldTravelID).
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
	_, err = tx.Travels.UpdateOneID(int(g.Id)).AddCollectNum(i).Save(ctx)
	return
}

func (r *travelRepo) Get(ctx context.Context, g *biz.TravelBase) (resp *biz.Travel, err error) {
	info, err := r.data.db.Travels.Query().Where(travels.IDEQ(int(g.Id))).
		Where(travels.Or(travels.AccountIDEQ(int(g.AccountId)), travels.IsHidden(false))).
		WithTravelAccount(func(query *ent.AccountQuery) {
			query.Select(account.FieldNickname, account.FieldAvatar, account.FieldDescription)
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
	if info.Edges.TravelAccount != nil {
		resp.Nickname = info.Edges.TravelAccount.Nickname
		resp.Avatar = info.Edges.TravelAccount.Avatar
		resp.TravelAccount.Description = info.Edges.TravelAccount.Description
	}
	if g.AccountId == 0 {
		return
	}
	extend, err := r.data.db.TravelExtends.Query().Where(travelextends.TravelIDEQ(info.ID), travelextends.AccountIDEQ(int(g.AccountId))).Only(ctx)
	if err != nil {
		return resp, nil
	}
	resp.IsCollect = extend.IsCollect
	resp.IsThumb = extend.IsThumb
	return
}

func (r *travelRepo) List(ctx context.Context, g *biz.TravelQuery) (int64, []*biz.Travel, error) {
	tx := r.data.db.Travels.Query().Where(travels.Or(travels.AccountIDEQ(int(g.AccountId)), travels.IsHidden(false)))
	tx = tx.WithTravelAccount(func(query *ent.AccountQuery) {
		query.Select(account.FieldNickname, account.FieldAvatar)
	})
	if g.My {
		tx = tx.Where(travels.AccountIDEQ(int(g.AccountId)))
	}
	if g.AccountId > 0 {
		tx = tx.WithTravelExtends(func(query *ent.TravelExtendsQuery) {
			query = query.Select(travelextends.FieldIsThumb, travelextends.FieldIsCollect).Where(travelextends.AccountIDEQ(int(g.AccountId)))
			if g.MyThumb {
				query = query.Where(travelextends.IsThumbEQ(true))
			}
			if g.MyCollect {
				query = query.Where(travelextends.IsCollectEQ(true))
			}
		})
	}
	if len(g.Title) > 0 {
		tx.Where(travels.TitleContains(g.Title))
	}
	if len(g.Description) > 0 {
		tx.Where(travels.DescriptionContains(g.Description))
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
	tx = tx.Order(ent.Desc(travels.FieldID))
	resp := make([]*biz.Travel, 0, g.GetPageSize())
	// account.FieldNickname, account.FieldAvatar,
	res, err := tx.Offset(int(g.GetOffset())).Limit(int(g.GetPageSize())).All(ctx) //.Scan(ctx, &resp)
	for _, info := range res {
		one := &biz.Travel{
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
			TravelExtend: biz.TravelExtend{
				BrowseNum:  info.BrowseNum,
				ThumbNum:   info.ThumbNum,
				CollectNum: info.CollectNum,
			},
		}
		if info.Edges.TravelAccount != nil {
			one.TravelAccount.Nickname = info.Edges.TravelAccount.Nickname
			one.TravelAccount.Avatar = info.Edges.TravelAccount.Avatar
		}
		if info.Edges.TravelExtends != nil && len(info.Edges.TravelExtends) == 1 {
			item := info.Edges.TravelExtends[0]
			one.TravelExtend.IsThumb = item.IsThumb
			one.TravelExtend.IsCollect = item.IsCollect
		}
		resp = append(resp, one)
	}
	return int64(total), resp, err
}
