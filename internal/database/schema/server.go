package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Server holds the schema definition for the Discord Server entity.
type Server struct {
	ent.Schema
}

// Fields of the Server.
func (Server) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.String("owner_id"),
		field.Bool("enabled").Default(false),
	}
}

func (Server) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Server) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("configuration", ServerConfig.Type).Unique().Required(),
		edge.To("word_blacklist", WordBlacklist.Type),
	}
}

func (Server) Annotations() []schema.Annotation {
	return nil
}
