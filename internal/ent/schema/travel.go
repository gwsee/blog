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

type Travels struct {
	ent.Schema
}

func (Travels) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Comment("旅行记录的ID"),
		field.String("title").NotEmpty().Comment("标题").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(200)",
			}),
		field.String("description").NotEmpty().Comment("旅行简介").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}),
		field.String("video").Default("").Comment("旅行视频").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(200)",
			}),
		field.Bool("is_hidden").Default(false).Comment("是否隐藏:0否,1是"),
		field.Int("account_id").Optional().Comment("账户ID"),
		field.JSON("photos", []string{}).Comment("旅行的照片"),
		field.Int("browse_num").Default(0).Comment("浏览量"),
		field.Int("thumb_num").Default(0).Comment("点赞量"),
		field.Int("collect_num").Default(0).Comment("收藏量"),
	}
}
func (Travels) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		mixin.SoftDelete{},
	}
}
func (Travels) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("travel_extends", TravelExtends.Type),
		edge.From("travel_account", Account.Type).Field("account_id").Ref("travels").Unique(),
	}
}
func (Travels) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "travels"},
		entsql.WithComments(true),
		schema.Comment("旅行"),
	}
}
