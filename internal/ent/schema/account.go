package schema

import (
	"blog/internal/ent/schema/mixin"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Account struct {
	ent.Schema
}

func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Positive().Comment("账户ID"),
		field.String("nickname").Default("好运连连").Comment("昵称"),
		field.String("account").Unique().NotEmpty().Comment("账户").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(100)",
			}),
		field.String("password").NotEmpty().Comment("密码").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(200)",
			}),
		field.String("email").Default("").Comment("邮箱").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(100)",
			}),
		field.String("description").Comment("想说啥?").
			SchemaType(map[string]string{
				dialect.MySQL: "tinytext",
			}),
		field.String("avatar").Default("").Comment("头像"),
		field.Int("blog_num").Default(0).Comment("博客数量"),
		field.Int8("status").Default(1).Comment("状态:0失效,1正常").
			SchemaType(map[string]string{
				dialect.MySQL: "tinyint",
			}),
	}
}
func (Account) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		mixin.SoftDelete{},
	}
}
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("travels", Travels.Type),
	}
}
func (Account) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "account"},
		entsql.WithComments(true),
		schema.Comment("站点"),
	}
}
