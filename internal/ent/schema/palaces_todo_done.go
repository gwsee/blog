package schema

import (
	"blog/internal/ent/schema/mixin"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type PalacesTodoDone struct {
	ent.Schema
}

// 一天只能完成一次
func (PalacesTodoDone) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Positive().Comment("ID"),
		field.Int("todo_id").Optional().Comment("代办事项的ID"),
	}
}
func (PalacesTodoDone) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", PalacesTodo.Type).Ref("dones").Field("todo_id").Unique(),
	}
}
func (PalacesTodoDone) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		mixin.SoftDelete{},
	}
}
func (PalacesTodoDone) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "palaces_todo_done"},
		entsql.WithComments(true),
		schema.Comment("代办事项完成记录表"),
	}
}
