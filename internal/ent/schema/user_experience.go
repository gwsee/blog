package schema

import (
	"blog/internal/ent/schema/mixin"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type AccountExperience struct {
	ent.Schema
}

func (AccountExperience) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Positive().Comment("ID"),
		field.Int("user_id").Comment("用户ID"),
		field.String("company").Comment("公司名称").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(100)",
			}),
		field.String("role").Comment("职位名称").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(100)",
			}),
		field.String("location").Comment("公司地址").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(200)",
			}),
		field.Int64("start").Comment("开始时间"),
		field.Int64("end").Comment("结束时间"),
		field.String("description").Comment("职位描述").
			SchemaType(map[string]string{
				dialect.MySQL: "tinytext",
			}),
		field.String("responsibilities").Comment("主要职责").
			SchemaType(map[string]string{
				dialect.MySQL: "tinytext",
			}),
		field.String("achievements").Comment("工作成就").
			SchemaType(map[string]string{
				dialect.MySQL: "tinytext",
			}),
		field.JSON("skills", []string{}).Comment("使用技能"),
	}
}
func (AccountExperience) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		mixin.SoftDelete{},
	}
}
func (AccountExperience) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "account_experience"},
		entsql.WithComments(true),
		schema.Comment("经历表"),
	}
}
