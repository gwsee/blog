package data

import (
	"blog/api/global"
	"blog/app/palaces/internal/biz"
	"blog/internal/common"
	"blog/internal/constx"
	"blog/internal/ent"
	"blog/internal/ent/palacestodo"
	"blog/internal/ent/palacestododone"
	"context"
	"entgo.io/ent/dialect/sql"
	"errors"
	"time"
)

func (p palacesRepo) ValidTodo(ctx context.Context, todo *biz.Todo) (err error) {
	old := new(ent.PalacesTodo)
	if todo.Id > 0 {
		old, err = p.data.db.PalacesTodo.Get(ctx, todo.Id)
		if err != nil {
			return
		}
		if old.Type != int8(todo.Type) {
			var count int
			count, err = p.data.db.PalacesTodoDone.Query().Where(palacestododone.TodoIDEQ(todo.Id)).Count(ctx)
			if err != nil {
				return
			}
			if count > 0 {
				err = errors.New("todo类型不能更换")
			}
		}
		if old.Status == constx.PlacesTodoDone {
			err = errors.New("todo 已完成 无需修改")
			return
		}
	}
	if todo.Type == constx.PalacesTodoNum {
		if todo.Num <= 0 {
			todo.Num = constx.PalacesTodoNumMin
		}
		if old.ID > 0 && old.Num != todo.Num {
			var count int
			count, err = p.data.db.PalacesTodoDone.Query().Where(palacestododone.TodoIDEQ(todo.Id)).Count(ctx)
			if err != nil {
				return
			}
			if count >= int(todo.Num) {
				err = errors.New("todo 按次时 调整需调整次数大于已完成次数")
				return
			}
			if todo.Status == constx.PlacesTodoDone {
				todo.Num = int64(count)
			}
		}
		return
	}
	if todo.From <= 0 {
		todo.From = common.GetZeroAt(time.Now().Unix())
	}
	if todo.Type == constx.PalacesTodoShort {
		if todo.To <= 0 {
			err = errors.New("短期todo 需要结束时间")
		}
		if todo.To < todo.From {
			err = errors.New("短期todo 结束时间不能小于开始时间")
		}
		todo.To = common.GetNextZeroAt(todo.To) - 1
		if todo.To < time.Now().Unix() {
			err = errors.New("短期todo 截止日期不能小于此刻")
			return
		}
	} else if todo.Type == constx.PalacesTodoLong {
		todo.To = 0
	} else {
		err = errors.New("无效类型")
		return
	}
	if old.ID > 0 && (old.From != todo.From || old.To != todo.To) {
		//当将todo的开始时间向此刻调整时 就是时间范围缩短时候
		if old.From < todo.From {
			var count int
			count, err = p.data.db.PalacesTodoDone.Query().Where(palacestododone.TodoIDEQ(todo.Id)).Count(ctx)
			if err != nil {
				return
			}
			if count > 0 {
				err = errors.New("已完成的todo 不能更换开始日期")
				return
			}
		}
		//当时间范围缩短的时候  要判断是否是当前时间
		current := common.GetNextZeroAt(time.Now().Unix()) - 1
		if old.To > todo.To && current > todo.To {
			err = errors.New("todo的截止日期 不能小于今天")
			return
		}
		if old.Status == constx.PlacesTodoDone && old.To > 0 {
			old.To = common.GetNextZeroAt(time.Now().Unix()) - 1
		}
	}
	return
}
func (p palacesRepo) SaveTodo(ctx context.Context, todo *biz.Todo) error {
	if todo.Id > 0 {
		return p.data.db.PalacesTodo.UpdateOneID(int(todo.Id)).Where(palacestodo.AccountIDEQ(int(todo.AccountId))).
			Where(palacestodo.StatusEQ(0)).
			SetTheme(todo.Theme).SetType(int8(todo.Type)).
			SetFrom(todo.From).SetTo(todo.To).SetNum(todo.Num).
			SetSort(todo.Sort).SetContent(todo.Content).
			SetStatus(int8(todo.Status)).
			Exec(ctx)
	} else {
		return p.data.db.PalacesTodo.Create().
			SetAccountID(todo.AccountId).
			SetTheme(todo.Theme).SetType(int8(todo.Type)).
			SetFrom(todo.From).SetTo(todo.To).SetNum(todo.Num).
			SetSort(todo.Sort).SetContent(todo.Content).
			SetStatus(int8(todo.Status)).
			Exec(ctx)
	}
}

func (p palacesRepo) DoneTodo(ctx context.Context, id *global.ID) (err error) {
	old, err := p.data.db.PalacesTodo.Get(ctx, int(id.GetId()))
	if err != nil {
		return err
	}
	if old.Status == 1 {
		err = errors.New("todo 已完成 无需继续完成")
		return err
	}
	tx, err := p.data.db.Tx(ctx)
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
	var count int
	count, err = tx.PalacesTodoDone.Query().
		Where(palacestododone.CreatedAtGTE(common.GetZeroAt(time.Now().Unix()))).
		Where(palacestododone.TodoIDEQ(int(id.Id))).Count(ctx)
	if err != nil {
		return
	}
	if count > 0 {
		err = errors.New("今日已打卡 请勿重复打卡")
		return
	}

	err = tx.PalacesTodoDone.Create().SetTodoID(int(id.Id)).Exec(ctx)
	if err != nil {
		return err
	}
	switch old.Type {
	case constx.PalacesTodoNum:
		count, err = tx.PalacesTodoDone.Query().Where(palacestododone.TodoIDEQ(int(id.Id))).Count(ctx)
		if err != nil {
			return
		}
		if count >= int(old.Num) {
			err = tx.PalacesTodo.UpdateOneID(int(id.Id)).SetStatus(constx.PlacesTodoDone).Exec(ctx)
		}
	case constx.PalacesTodoShort:
		end := common.GetNextZeroAt(time.Now().Unix())
		if old.To <= end {
			err = tx.PalacesTodo.UpdateOneID(int(id.Id)).SetStatus(constx.PlacesTodoDone).Exec(ctx)
		}
	case constx.PalacesTodoLong:
		//长期的就不需要管了
	}
	return

}

func (p palacesRepo) DeleteTodo(ctx context.Context, id *global.ID) error {
	return p.data.db.PalacesTodo.DeleteOneID(int(id.Id)).Where(palacestodo.AccountIDEQ(int(id.GetAccountId()))).Exec(ctx)
}

func (p palacesRepo) DeleteTodoDone(ctx context.Context, id *global.ID) error {
	return p.data.db.PalacesTodoDone.DeleteOneID(int(id.Id)).
		Where(palacestododone.HasOwnerWith(palacestodo.AccountIDEQ(int(id.GetAccountId())))).
		Exec(ctx)
}

func (p palacesRepo) DeleteTodoRecord(ctx context.Context, id *global.ID) error {
	_, err := p.data.db.PalacesTodoDone.Delete().Where(palacestododone.TodoIDEQ(int(id.Id))).
		Where(palacestododone.HasOwnerWith(palacestodo.AccountIDEQ(int(id.GetAccountId())))).
		Exec(ctx)
	return err
}

func (p palacesRepo) ListTodo(ctx context.Context, query *biz.TodoQuery) (int64, []*biz.Todo, error) {
	tx := p.data.db.PalacesTodo.Query().Where(palacestodo.AccountIDEQ(int(query.AccountId)))
	if query.Status != nil {
		tx = tx.Where(palacestodo.Status(int8(*query.Status)))
	}
	if query.From > 0 {
		tx = tx.Where(palacestodo.FromGTE(query.From))
	}
	if query.Type > 0 {
		tx = tx.Where(palacestodo.TypeEQ(int8(query.Type)))
	}
	if query.Theme != "" {
		tx = tx.Where(palacestodo.ThemeContains(query.Theme))
	}
	tx = tx.Where(palacestodo.And(func(s *sql.Selector) {
		t := sql.Table(palacestododone.Table)
		s.Where(
			sql.NotExists(
				sql.Select(t.C(palacestododone.FieldTodoID)).
					From(t).
					Where(
						sql.And(
							sql.ColumnsEQ(s.C(palacestodo.FieldID), t.C(palacestododone.FieldTodoID)),
							sql.GTE(t.C(palacestododone.FieldCreatedAt), common.GetZeroAt(time.Now().Unix())),
							sql.EQ(palacestodo.FieldAccountID, query.AccountId),
							sql.EQ(palacestododone.FieldDeletedAt, 0),
						),
					),
			),
		)
	}))
	total, err := tx.Count(ctx)
	if err != nil {
		return 0, nil, err
	}
	if total == 0 {
		return 0, []*biz.Todo{}, nil
	}
	tx = tx.Order(ent.Asc(palacestodo.FieldStatus), ent.Desc(palacestodo.FieldUpdatedAt))
	res, err := tx.Offset(int(query.GetOffset())).Limit(int(query.GetPageSize())).All(ctx)
	if err != nil {
		return 0, nil, err
	}
	resp := make([]*biz.Todo, 0, query.GetPageSize())
	for _, m := range res {
		resp = append(resp, &biz.Todo{
			Id:        m.ID,
			Theme:     m.Theme,
			Type:      int64(m.Type),
			From:      m.From,
			To:        m.To,
			Num:       m.Num,
			Sort:      m.Sort,
			Content:   m.Content,
			Status:    int64(m.Status),
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		})
	}
	return int64(total), resp, nil
}

func (p palacesRepo) ListTodoDone(ctx context.Context, query *biz.TodoQuery) (int64, []*biz.Todo, error) {
	db := p.data.db.PalacesTodoDone.Query().WithOwner(func(tx *ent.PalacesTodoQuery) {
		tx = tx.Where(palacestodo.AccountIDEQ(query.AccountId)) // palacestodo.Status(constx.PlacesTodoDone)
		if query.From > 0 {
			tx = tx.Where(palacestodo.FromGTE(query.From))
		}
		if query.Type > 0 {
			tx = tx.Where(palacestodo.TypeEQ(int8(query.Type)))
		}
		if query.Theme != "" {
			tx = tx.Where(palacestodo.ThemeContains(query.Theme))
		}
		return
	})
	total, err := db.Count(ctx)
	if err != nil {
		return 0, nil, err
	}
	res, err := db.Offset(int(query.GetOffset())).Limit(int(query.GetPageSize())).All(ctx)
	if err != nil {
		return 0, nil, err
	}
	resp := make([]*biz.Todo, 0, query.GetPageSize())
	for _, x := range res {
		m := x.Edges.Owner
		if m == nil {
			continue
		}
		resp = append(resp, &biz.Todo{
			Id:        m.ID,
			Theme:     m.Theme,
			Type:      int64(m.Type),
			From:      m.From,
			To:        m.To,
			Num:       m.Num,
			Sort:      m.Sort,
			Content:   m.Content,
			Status:    int64(m.Status),
			CreatedAt: x.CreatedAt,
			UpdatedAt: x.UpdatedAt,
			DoneId:    int64(x.ID),
		})
	}
	return int64(total), resp, nil
}
