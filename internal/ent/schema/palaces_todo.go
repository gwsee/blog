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

type PalacesTodo struct {
	ent.Schema
}

/**
代办事项分为
	1：按次数执行 执行到具体次数就完结了
	2：每日任务，从开始到结束 每日需要执行的；到具体结束日期就自动完结
*/

func (PalacesTodo) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Positive().Comment("ID"),
		field.Int("account_id").Comment("账户ID"),
		field.String("theme"), //主题
		field.Int8("type").Default(1).Comment("类型(1:按次,2：固定日期,2无固定日期)"),
		field.Int64("from"), //代办事项 开始日期
		field.Int64("to"),   //代办事项 结束日期 如果没有日期 就代办一直执行
		field.Int64("num"),  //代办事项 重复执行次数
		field.Int64("sort"), //数值越大越重要
		//field.Int64("pid"),    //父代办是什么  我们可以把代办事项拆分为具体明细项进行执行 如果存在子项的话 就是代办子项完成了 就代办他也完成了
		field.String("content").NotEmpty().Comment("内容").
			SchemaType(map[string]string{
				dialect.MySQL: "mediumtext",
			}),
		field.Int8("status").Comment("状态 0 代办,1 已做"),
	}
}
func (PalacesTodo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		mixin.SoftDelete{},
	}
}
func (PalacesTodo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("dones", PalacesTodoDone.Type),
	}
}
func (PalacesTodo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "palaces_todo"},
		entsql.WithComments(true),
		schema.Comment("代办事项"),
	}
}
