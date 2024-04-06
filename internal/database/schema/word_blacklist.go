package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// WordBlacklist holds the schema definition for the Word Blacklist entity.
type WordBlacklist struct {
	ent.Schema
}

// Fields of the Word Blacklist.
func (WordBlacklist) Fields() []ent.Field {
	return []ent.Field{
		field.String("word"),
	}
}

func (WordBlacklist) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (WordBlacklist) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("server", Server.Type).Ref("word_blacklist"),
	}
}

func (WordBlacklist) Annotations() []schema.Annotation {
	return nil
}
