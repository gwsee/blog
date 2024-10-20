package mixin

import (
	"blog/internal/constx"
	gen "blog/internal/ent"
	"blog/internal/ent/hook"
	"blog/internal/ent/intercept"
	"context"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"fmt"
)

// SoftDelete implements the soft delete pattern for schemas.
type SoftDelete struct {
	mixin.Schema
}

// Fields of the SoftDelete.
func (SoftDelete) Fields() []ent.Field {
	return []ent.Field{
		field.Uint8("is_deleted").Default(constx.NotDelete).Comment("是否删除;0：正常，1：删除"),
		field.Int64("deleted_at").Default(0).Comment("软删除时间"),
		field.String("deleted_by").Default(constx.EmptyUser).Comment("删除人"),
	}
}

type softDeleteKey struct{}

// SkipSoftDelete 返回一个context，表示在修改和查询阶段跳过软删除逻辑。
func SkipSoftDelete(parent context.Context) context.Context {
	return context.WithValue(parent, softDeleteKey{}, true)
}

// Interceptors of the SoftDelete.
func (d SoftDelete) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		intercept.TraverseFunc(func(ctx context.Context, q intercept.Query) error {
			// Skip soft-delete, means include soft-deleted entities.
			if skip, _ := ctx.Value(softDeleteKey{}).(bool); skip {
				return nil
			}
			d.P(q)
			return nil
		}),
	}
}

// Hooks of the SoftDelete.
func (d SoftDelete) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
					if skip, _ := ctx.Value(softDeleteKey{}).(bool); skip {
						return next.Mutate(ctx, m)
					}
					mx, ok := m.(interface {
						SetOp(ent.Op)
						Client() *gen.Client
						SetIsDeleted(uint8)
						SetDeletedBy(s string)
						SetDeletedAt(int642 int64)
						WhereP(...func(*sql.Selector))
					})
					if !ok {
						return nil, fmt.Errorf("unexpected mutation type %T", m)
					}
					d.P(mx)
					mx.SetOp(ent.OpUpdate)
					mx.SetIsDeleted(constx.IsDeleted)
					mx.SetDeletedBy(constx.DefaultUser.Default(ctx).Name)
					mx.SetDeletedAt(constx.Now().Unix())
					return mx.Client().Mutate(ctx, m)
				})
			},
			ent.OpDeleteOne|ent.OpDelete,
		),
	}
}

// P 方法增加另一个更强的筛选校验，只删除和查询is_deleted = 0的数据。 // 如果出现 deleted_at is null则代表数据有问题
func (d SoftDelete) P(w interface{ WhereP(...func(*sql.Selector)) }) {
	w.WhereP(
		sql.FieldEQ(d.Fields()[0].Descriptor().Name, constx.NotDelete),
	)
}
