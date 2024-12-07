package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type BlogsContent struct {
	ent.Schema
}

func (BlogsContent) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Comment("博客ID"),
		field.String("content").NotEmpty().Comment("内容").
			SchemaType(map[string]string{
				dialect.MySQL: "mediumtext",
			}),
		field.JSON("files", []string{}).Comment("博客内容可以关联的文件"),
	}
}

func (BlogsContent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "blogs_content"},
		entsql.WithComments(true),
		schema.Comment("博客内容"),
	}
}
