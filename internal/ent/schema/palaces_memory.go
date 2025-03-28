package schema

import (
	"blog/internal/ent/schema/mixin"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type PalacesMemory struct {
	ent.Schema
}

func (PalacesMemory) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Positive().Comment("ID"),
		//field.String("name"), //名
		//field.String("content").NotEmpty().Comment("内容").
		//	SchemaType(map[string]string{
		//		dialect.MySQL: "mediumtext",
		//	}),
		field.Int8("status"),
	}
}
func (PalacesMemory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		mixin.SoftDelete{},
	}
}
func (PalacesMemory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "palaces_memory"},
		entsql.WithComments(true),
		schema.Comment("记忆宫殿"),
	}
}
