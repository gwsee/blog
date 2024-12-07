package schema

import (
	"blog/internal/ent/schema/mixin"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type FilesExtend struct {
	ent.Schema
}

func (FilesExtend) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Positive().Comment("当前账户关联的文件的ID"),
		field.String("file_id").Default("").Comment("文件的ID").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(200)",
			}),
		field.Int("user_id").Default(0).Comment("所属的用户"),
		field.String("from").Comment("文件来源的表"),   //account-avatar;blog;travel;project
		field.Int("from_id").Comment("文件来源表的ID"), //account-avatar;blog;travel;project
		field.Int8("is_hidden").Comment("是否隐藏"),
	}
}

func (FilesExtend) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		mixin.SoftDelete{},
	}
}

func (FilesExtend) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "files_extend"},
		entsql.WithComments(true),
		schema.Comment("文件关系表"),
	}
}
