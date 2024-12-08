package schema

import (
	"blog/internal/ent/schema/mixin"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Positive().Comment("用户ID=账户ID"),
		field.String("name").Default("").Comment("姓名").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(100)",
			}),
		field.String("avatar").Default("").Comment("头像").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(200)",
			}),
		field.String("email").Default("").Comment("工作邮箱").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(100)",
			}),

		field.String("professional").Default("").Comment("职称").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(100)",
			}),
		field.String("address").Default("").Comment("地址").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(200)",
			}),
		field.JSON("skills", []string{}).Comment("技能"),
		field.String("description").NotEmpty().Comment("个人简介").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}),
		field.Int("experience").Default(0).Comment("经历数"),
		field.Int("project").Default(0).Comment("项目数"),
	}
}
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		mixin.SoftDelete{},
	}
}
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user"},
		entsql.WithComments(true),
		schema.Comment("用户"),
	}
}
