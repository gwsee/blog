package mixin

import (
	"blog/internal/constx"
	"blog/internal/ent/hook"
	"context"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type Time struct {
	mixin.Schema
}

func (Time) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("created_at").Immutable().Default(0).UpdateDefault(constx.NowUnix).Comment("创建时间").Immutable(),
		field.Int64("created_by").Default(0).Comment("创建人").Immutable(),
		field.Int64("updated_at").Default(0).UpdateDefault(constx.NowUnix).Comment("更新时间"),
		field.Int64("updated_by").Default(0).Comment("更新人"),
	}
}

// Hooks of the Default Time with User.
func (d Time) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
					mx, ok := m.(interface {
						CreatedAt() (r int64, exists bool)
						SetCreatedAt(s int64)
						UpdatedAt() (r int64, exists bool)
						SetUpdatedAt(s int64)
						CreatedBy() (r int64, exists bool)
						SetCreatedBy(s int64)
						UpdatedBy() (r int64, exists bool)
						SetUpdatedBy(s int64)
					})
					if ok {
						if ex, _ := mx.CreatedAt(); ex == constx.EmptyInt {
							mx.SetCreatedAt(constx.NowUnix())
						}
						if ex, _ := mx.UpdatedAt(); ex == constx.EmptyInt {
							mx.SetUpdatedAt(constx.NowUnix())
						}
						if ex, _ := mx.CreatedBy(); ex == constx.EmptyInt {
							mx.SetCreatedBy(constx.DefaultUser.Default(ctx).Id)
						}
						if ex, _ := mx.UpdatedBy(); ex == constx.EmptyInt {
							mx.SetUpdatedBy(constx.DefaultUser.Default(ctx).Id)
						}
					}
					return next.Mutate(ctx, m)
				})
			},
			ent.OpCreate,
		),
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
					mx, ok := m.(interface {
						UpdatedAt() (r int64, exists bool)
						SetUpdatedAt(s int64)
						UpdatedBy() (r int64, exists bool)
						SetUpdatedBy(s int64)
					})
					if ok {
						if ex, _ := mx.UpdatedAt(); ex == constx.EmptyInt {
							mx.SetUpdatedAt(constx.NowUnix())
						}
						if ex, _ := mx.UpdatedBy(); ex == constx.EmptyInt {
							mx.SetUpdatedBy(constx.DefaultUser.Default(ctx).Id)
						}
					}
					return next.Mutate(ctx, m)
				})
			},
			ent.OpUpdate|ent.OpUpdateOne,
		),
	}
}
