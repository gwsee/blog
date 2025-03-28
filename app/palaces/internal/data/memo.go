package data

import (
	"blog/api/global"
	"blog/app/palaces/internal/biz"
	"blog/internal/ent"
	"blog/internal/ent/palacesmemo"
	"context"
)

func (p palacesRepo) SaveMemo(ctx context.Context, memo *biz.Memo) error {
	if memo.Id > 0 {
		return p.data.db.PalacesMemo.UpdateOneID(memo.Id).
			Where(palacesmemo.AccountIDEQ(memo.AccountId)).
			SetContent(memo.Content).SetName(memo.Name).
			SetStatus(memo.Status).
			Exec(ctx)
	} else {
		return p.data.db.PalacesMemo.Create().
			SetAccountID(memo.AccountId).
			SetContent(memo.Content).SetName(memo.Name).
			SetStatus(memo.Status).
			Exec(ctx)
	}
}

func (p palacesRepo) DoneMemo(ctx context.Context, id *global.State) error {
	return p.data.db.PalacesMemo.UpdateOneID(int(id.Id)).Where(palacesmemo.AccountIDEQ(int(id.GetAccountId()))).
		SetStatus(int8(id.State)).
		Exec(ctx)
}

func (p palacesRepo) DeleteMemo(ctx context.Context, id *global.ID) error {
	return p.data.db.PalacesMemo.DeleteOneID(int(id.Id)).Where(palacesmemo.AccountIDEQ(int(id.GetAccountId()))).Exec(ctx)
}

func (p palacesRepo) ListMemo(ctx context.Context, query *biz.MemoQuery) (int64, []*biz.Memo, error) {
	tx := p.data.db.PalacesMemo.Query().Where(palacesmemo.AccountIDEQ(query.AccountId))
	if query.Name != "" {
		tx = tx.Where(palacesmemo.NameContains(query.Name))
	}
	if query.Status != nil {
		tx = tx.Where(palacesmemo.StatusEQ(int8(*query.Status)))
	}
	total, err := tx.Count(ctx)
	if err != nil {
		return 0, nil, err
	}
	if total == 0 {
		return 0, []*biz.Memo{}, nil
	}
	tx = tx.Order(ent.Asc(palacesmemo.FieldStatus), ent.Desc(palacesmemo.FieldUpdatedAt))
	resp := make([]*biz.Memo, 0, query.GetPageSize())
	res, err := tx.Offset(int(query.GetOffset())).Limit(int(query.GetPageSize())).All(ctx)
	if err != nil {
		return 0, nil, err
	}
	for _, m := range res {
		resp = append(resp, &biz.Memo{
			Id:        m.ID,
			Name:      m.Name,
			Content:   m.Content,
			Status:    m.Status,
			UpdatedAt: m.UpdatedAt,
			CreatedAt: m.CreatedAt,
		})
	}
	return int64(total), resp, nil
}
