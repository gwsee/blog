package schema

import (
	"blog/internal/ent/schema/mixin"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Comment struct {
	ent.Schema
}

func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("account_id").Positive().Comment("账户ID"),
		field.Int8("kind").Default(0).Comment("对应类型;0评论"),
		field.Int32("kind_id").Positive().Comment("对应类型的ID"),

		field.Int32("id").Positive().Unique().Comment("评论ID"),
		field.Int32("top_id").Default(0).Comment("顶级ID"),
		field.Int32("parent_id").Default(0).Comment("父评论"),
		field.Int8("status").Default(0).Comment("0显示,1隐藏"),
		field.Int32("level").Default(0).Comment("第几楼"),
		field.Int32("total").Default(0).Comment("当前级有多少数据;状态未0的"),
		field.String("content").Comment("评论内容"),
	}
}
func (Comment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		mixin.SoftDelete{},
	}
}
func (Comment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "comments"},
		entsql.WithComments(true),
		schema.Comment("博客评论"),
	}
}
