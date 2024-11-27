package schema

import (
	"blog/internal/ent/schema/mixin"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Travel struct {
	ent.Schema
}

func (Travel) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Comment("旅行记录的ID"),
		field.String("title").Default("").Comment("标题").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(200)",
			}),
		field.String("description").NotEmpty().Comment("旅行简介").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}),
		field.String("video").NotEmpty().Comment("旅行视频").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(200)",
			}),
		field.JSON("photos", []string{}).Comment("旅行的照片"),
		field.Int("browse_num").Comment("浏览量"),
		field.Int("thumb_num").Comment("点赞量"),
		field.Int("collect_num").Comment("收藏量"),
	}
}
func (Travel) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		mixin.SoftDelete{},
	}
}
func (Travel) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "travel"},
		entsql.WithComments(true),
		schema.Comment("旅行"),
	}
}
