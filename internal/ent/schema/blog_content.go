package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type BlogContent struct {
	ent.Schema
}

func (BlogContent) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Comment("博客ID"),
		field.String("content").NotEmpty().Comment("内容").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}),
	}
}

func (BlogContent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "blog_content"},
		entsql.WithComments(true),
		schema.Comment("博客内容"),
	}
}
