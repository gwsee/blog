package schema

import (
	"blog/internal/ent/schema/mixin"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Blog struct {
	ent.Schema
}

func (Blog) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Positive().Comment("博客ID"),
		field.Int("account_id").Comment("账户ID"),
		field.String("title").Comment("标题").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(100)",
			}),
		field.String("description").Comment("描述").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(255)",
			}),
		field.Int8("is_hidden").Default(0).Comment("是否隐藏:0否,1是").
			SchemaType(map[string]string{
				dialect.MySQL: "tinyint",
			}),
		field.JSON("tags", []string{}).Comment("标签"),
		field.String("cover").Comment("封面").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(200)",
			}),
		field.String("content").NotEmpty().Comment("内容").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}),
	}
}

func (Blog) Edges() []ent.Edge {
	return []ent.Edge{}
}
func (Blog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		mixin.SoftDelete{},
	}
}
func (Blog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "blogs"},
		entsql.WithComments(true),
		schema.Comment("博客"),
	}
}
