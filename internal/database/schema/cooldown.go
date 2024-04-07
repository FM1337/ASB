package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Cooldown holds the schema definition for the Cooldown entity.
type Cooldown struct {
	ent.Schema
}

// Fields of the Cooldown.
func (Cooldown) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").Comment("the discord id of the user"),
		field.String("hash").Comment("the hash of the message being potentially spammed"),
		field.Int("count").Default(1).Comment("The amount of times the message has been seen"),
		field.Time("resets_at").Default(time.Now).Comment("The time the cooldown expires at"),
	}
}

func (Cooldown) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Cooldown) Edges() []ent.Edge {
	// return nil
	return []ent.Edge{
		edge.From("server", Server.Type).Ref("cooldown").Unique().Required(),
	}
}

func (Cooldown) Annotations() []schema.Annotation {
	return nil
}
