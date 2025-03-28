package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type BlogsExtend struct {
	ent.Schema
}

func (BlogsExtend) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Comment("ID"),
		field.Int("blog_id").Comment("博客ID"),
		field.Int("account_id").Comment("账户ID"),
		field.Int("browse_num").Default(0).Comment("浏览量"),
		field.Int64("browse_at").Default(0).Comment("最后的浏览时间"),
		field.Bool("collect").Default(false).Comment("是否收藏"),
		field.Int64("collect_at").Default(0).Comment("收藏时间"),
		field.Bool("love").Default(false).Comment("是否点赞"),
		field.Int64("love_at").Default(0).Comment("点赞时间"),
	}
}

func (BlogsExtend) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "blogs_extend"},
		entsql.WithComments(true),
		schema.Comment("博客关联信息"),
	}
}

// Indexes 联合唯一
func (BlogsExtend) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("blog_id", "account_id").Unique(),
	}
}
