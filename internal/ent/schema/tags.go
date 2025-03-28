package schema

import (
	"blog/internal/ent/schema/mixin"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Tags struct {
	ent.Schema
}

func (Tags) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		mixin.SoftDelete{},
	}
}
func (Tags) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Positive().Comment("标签ID"),
		field.String("name").Unique().NotEmpty(), //标签名
	}
}

func (Tags) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("blogs", Blogs.Type).Ref("tag").
			Through("tag_relation", TagsRelation.Type),
	}
}
func (Tags) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tags"},
		entsql.WithComments(true),
		schema.Comment("标签信息"),
	}
}
