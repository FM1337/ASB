package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Spammer holds the schema definition for the Spammer entity.
type Spammer struct {
	ent.Schema
}

// Fields of the Spammer.
func (Spammer) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").Comment("the discord id of the user"),
		field.Strings("removed_roles").Optional().Comment("a list of removed roles on the user if remove roles is enabled"),
		field.Time("last_flagged").Default(time.Now).Comment("The last time the user was flagged by the bot as a spammer"),
	}
}

func (Spammer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Spammer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("server", Server.Type).Ref("spammer").Unique().Required(),
	}
}

func (Spammer) Annotations() []schema.Annotation {
	return nil
}
