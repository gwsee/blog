package schema

import (
	"blog/internal/ent/schema/mixin"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Files struct {
	ent.Schema
}

/**
关于文件系统的设计
	1：文件相同的内容只存储一份
	2：不是所有文件都能够被你看见，要这个文件有对应的权限才行--所以有个files_relation表;
		通过 files_relation 的is_hidden字段，如果存在一个为false 就代表这个文件是公开的，
		否则只能是自己上传的文件可以看 来判断是否有权限
	3：访问方式是 域名/files/md5  的形式走服务器来访问对应的文件 在接口files中 通过redis的方式缓存对应的2的内容 时效为1小时; 来获取对应的内容

	todo  优化文件显示
		利用短期签名 URL
		引入 CDN 加速

	对于自建的存储服务
	对于阿里云OSS存储
*/

func (Files) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Comment("文件的ID(内容MD5)"),
		field.String("type").Default("").Comment("文件类型").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(50)",
			}),
		field.Int64("size").Default(0).Comment("文件大小"),
		field.String("name").Default("").Comment("文件名称"),
		field.String("path").Default("").Comment("文件真实存放地址").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(250)",
			}),
	}
}
func (Files) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		mixin.SoftDelete{},
	}
}
func (Files) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "files"},
		entsql.WithComments(true),
		schema.Comment("文件"),
	}
}
