package schema

import (
	"blog/internal/ent/schema/mixin"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type PalacesMemo struct {
	ent.Schema
}

// 备忘录：就是纯备忘罢了

func (PalacesMemo) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Positive().Comment("ID"),
		field.Int("account_id").Comment("账户ID"),
		field.String("name"), //名
		field.String("content").NotEmpty().Comment("内容").
			SchemaType(map[string]string{
				dialect.MySQL: "mediumtext",
			}),
		field.Int8("status").Comment("不同数值代办的不同记忆情况 越大代表越不需要记忆 为后续数据列表做铺垫"),
	}
}
func (PalacesMemo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		mixin.SoftDelete{},
	}
}
func (PalacesMemo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "palaces_memo"},
		entsql.WithComments(true),
		schema.Comment("备忘录"),
	}
}
