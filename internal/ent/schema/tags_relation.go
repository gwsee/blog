package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type TagsRelation struct {
	ent.Schema
}

func (TagsRelation) Fields() []ent.Field {
	return []ent.Field{
		field.Int("tag_id").Comment("标签ID"),
		field.String("relation"), //标签关联的表 目前只做blog
		field.Int("relation_id"),
	}
}
func (TagsRelation) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("tag_id", "relation_id").Unique(),
	}
}
func (TagsRelation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("blog", Blogs.Type).
			Unique().
			Required().
			Field("relation_id"),
		edge.To("tag", Tags.Type).
			Unique().
			Required().
			Field("tag_id"),
	}
}
func (TagsRelation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tags_relation"},
		entsql.WithComments(true),
		schema.Comment("标签-关联表信息"),
	}
}
