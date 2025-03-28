package schema

import (
	"blog/internal/ent/schema/mixin"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type UserExperience struct {
	ent.Schema
}

func (UserExperience) Fields() []ent.Field {
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
				dialect.MySQL: "mediumtext",
			}),
		field.String("responsibilities").Comment("主要职责").
			SchemaType(map[string]string{
				dialect.MySQL: "mediumtext",
			}),
		field.String("achievements").Comment("工作成就").
			SchemaType(map[string]string{
				dialect.MySQL: "mediumtext",
			}),
		field.JSON("skills", []string{}).Comment("使用技能"),
		field.Int("project").Default(0).Comment("项目数"),
		field.String("image").Default("863f7821fa42eb9d61091b5c6df1c4b0").Comment("公司名称").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(255)",
			}),
	}
}
func (UserExperience) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		mixin.SoftDelete{},
	}
}
func (UserExperience) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "users_experience"},
		entsql.WithComments(true),
		schema.Comment("经历表"),
	}
}
