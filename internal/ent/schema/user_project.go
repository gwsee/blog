package schema

import (
	"blog/internal/ent/schema/mixin"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type UserProject struct {
	ent.Schema
}

func (UserProject) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Positive().Comment("ID"),
		field.Int("user_id").Comment("用户ID"),
		field.Int("experience_id").Comment("经历ID"),
		field.String("title").Comment("项目名称").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(200)",
			}),
		field.String("description").Comment("项目描述").
			SchemaType(map[string]string{
				dialect.MySQL: "mediumtext",
			}),
		field.JSON("skills", []string{}).Comment("使用技能"),
		field.Int64("start").Comment("开始时间"),
		field.Int64("end").Comment("结束时间"),
		field.String("link").Comment("项目地址").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(200)",
			}),
		field.JSON("photos", []string{}).Comment("项目照片"),
	}
}
func (UserProject) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		mixin.SoftDelete{},
	}
}
func (UserProject) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "users_project"},
		entsql.WithComments(true),
		schema.Comment("项目表"),
	}
}
