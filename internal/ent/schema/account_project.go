package schema

import (
	"blog/internal/ent/schema/mixin"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type AccountProject struct {
	ent.Schema
}

func (AccountProject) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Positive().Comment("ID"),
		field.Int("account_id").Comment("账户ID"),
		field.Int("experience_id").Comment("经历ID"),
		field.String("title").Comment("项目名称"),
		field.String("description").Comment("项目描述"),
		field.JSON("skills", []string{}).Comment("使用技能"),
		field.Int64("start").Comment("开始时间"),
		field.Int64("end").Comment("结束时间"),
		field.String("link").Comment("项目地址"),
		field.JSON("pics", []string{}).Comment("项目快照"),
	}
}
func (AccountProject) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		mixin.SoftDelete{},
	}
}
func (AccountProject) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "account_project"},
		entsql.WithComments(true),
		schema.Comment("项目表"),
	}
}
