package schema

import (
	"blog/internal/ent/schema/mixin"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type BlogsComment struct {
	ent.Schema
}

func (BlogsComment) Fields() []ent.Field {
	return []ent.Field{
		field.Int("account_id").Positive().Comment("账户ID"),
		field.Int("blog_id").Positive().Comment("博客ID"),
		field.Int("id").Positive().Unique().Comment("评论ID"),
		field.Int("top_id").Default(0).Comment("顶级ID"),
		field.Int("parent_id").Default(0).Comment("父评论"),
		field.Int("level").Default(0).Comment("第几楼"),
		field.Int("total").Default(0).Comment("当前级有多少数据;状态未0的"),
		field.Int8("status").Default(0).Comment("0显示,1隐藏").
			SchemaType(map[string]string{
				dialect.MySQL: "tinyint",
			}),
		field.String("content").Comment("评论内容"),
	}
}
func (BlogsComment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		mixin.SoftDelete{},
	}
}
func (BlogsComment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "blogs_comments"},
		entsql.WithComments(true),
		schema.Comment("博客评论"),
	}
}
