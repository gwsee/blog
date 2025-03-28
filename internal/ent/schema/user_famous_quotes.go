package schema

import (
	"blog/internal/ent/schema/mixin"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type UserFamousQuotes struct {
	ent.Schema
}

func (UserFamousQuotes) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		mixin.SoftDelete{},
	}
}
func (UserFamousQuotes) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Positive().Comment("ID"),
		field.String("text").Comment("名言警句").
			SchemaType(map[string]string{
				dialect.MySQL: "mediumtext",
			}),
	}
}

func (UserFamousQuotes) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "users_famous_quotes"},
		entsql.WithComments(true),
		schema.Comment("名言警句"),
	}
}
