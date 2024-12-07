package schema

import (
	"blog/internal/ent/schema/mixin"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type TravelExtend struct {
	ent.Schema
}

func (TravelExtend) Fields() []ent.Field {
	return []ent.Field{
		field.Int("account_id").Positive().Comment("账户ID"),
		field.Int("travel_id").Comment("旅行的ID"),
		field.Bool("is_thumb").Comment("是否点赞"),
		field.Bool("is_collect").Comment("收藏量"),
	}
}
func (TravelExtend) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("account_id", "travel_id").Unique(),
	}
}
func (TravelExtend) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("travel", Travel.Type).Ref("travel_extend").Unique(),
	}
}
func (TravelExtend) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		mixin.SoftDelete{},
	}
}
func (TravelExtend) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "travel_extend"},
		entsql.WithComments(true),
		schema.Comment("旅行关联关系"),
	}
}
