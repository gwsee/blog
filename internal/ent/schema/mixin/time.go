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
		field.Int64("created_at").Immutable().Default(constx.Now().Unix()).Comment("创建时间").Immutable(),
		field.String("created_by").Default("").Comment("创建人").Immutable(),
		field.Int64("updated_at").Default(constx.Now().Unix()).UpdateDefault(constx.Now().Unix).Comment("更新时间"),
		field.String("updated_by").Default("").Comment("更新人"),
	}
}

// Hooks of the Default Time with User.
func (d Time) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
					mx, ok := m.(interface {
						CreatedBy() (r string, exists bool)
						SetCreatedBy(s string)
						UpdatedBy() (r string, exists bool)
						SetUpdatedBy(s string)
						SetCompanyID(i int)
					})
					if ok {
						if ex, _ := mx.CreatedBy(); ex == constx.EmptyUser {
							mx.SetCreatedBy(constx.DefaultUser.Default(ctx).Name)
						}
						if ex, _ := mx.UpdatedBy(); ex == constx.EmptyUser {
							mx.SetUpdatedBy(constx.DefaultUser.Default(ctx).Name)
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
						UpdatedBy() (r string, exists bool)
						SetUpdatedBy(s string)
					})
					if ok {
						if ex, _ := mx.UpdatedBy(); ex == constx.EmptyUser {
							mx.SetUpdatedBy(constx.DefaultUser.Default(ctx).Name)
						}
					}
					return next.Mutate(ctx, m)
				})
			},
			ent.OpUpdate|ent.OpUpdateOne,
		),
	}
}
