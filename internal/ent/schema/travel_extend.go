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

type TravelExtends struct {
	ent.Schema
}

func (TravelExtends) Fields() []ent.Field {
	return []ent.Field{
		field.Int("account_id").Positive().Comment("账户ID"),
		field.Int("travel_id").Comment("旅行的ID"),
		field.Bool("is_thumb").Comment("是否点赞"),
		field.Bool("is_collect").Comment("收藏量"),
	}
}
func (TravelExtends) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("account_id", "travel_id").Unique(),
	}
}

func (TravelExtends) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("extends", Travels.Type).Ref("travel_extends").Unique(),
	}
}
func (TravelExtends) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		mixin.SoftDelete{},
	}
}
func (TravelExtends) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "travel_extends"},
		entsql.WithComments(true),
		schema.Comment("旅行关联关系"),
	}
}
