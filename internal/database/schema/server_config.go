package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// ServerConfig holds the schema definition for the Bot Server Config entity.
type ServerConfig struct {
	ent.Schema
}

// Fields of the Server Config.
func (ServerConfig) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("remove_roles").Default(false).Comment("Indicates whether a spammer should have their roles removed"),
		field.Bool("give_role").Default(false).Comment("Indicates if a role should be assigned to a spammer"),
		field.Bool("timeout").Default(false).Comment("Indicates if a spammer should be timed out"),
		field.Bool("kick").Default(false).Comment("Indicates if a spammer should be kicked (takes priority over timeout)"),
		field.Bool("ban").Default(false).Comment("Indicates if a spammer should be banned (takes priority over kick)"),
		field.Bool("check_invites").Default(false).Comment("Indicate if information about an invite should be checked against the blacklist"),
		field.Bool("check_links").Default(false).Comment("Indicates if a url should be investigated"),
		field.Bool("ratelimit").Default(false).Comment("Indicates if ratelimiting is enabled to identify spammers"),
		field.Bool("alerts").Default(false).Comment("Indicates if an alert should be posted to the alert channel for a server when a spammer is detected"),
		field.Bool("flag_links").Default(false).Comment("Indicates if a user should be monitored if they post a message containing a URL"),

		field.String("log_channel").Default("").Comment("The alert/log channel for a server"),
		field.String("given_role").Default("").Comment("The role to give to a spammer if give_role is true"),
		field.Strings("excluded_channels").Default([]string{}).Comment("Channels listed should not be monitored for spammers"),
		field.Strings("excluded_roles").Default([]string{}).Comment("Roles that should be ignored for spam checks"),
		field.Strings("excluded_users").Default([]string{}).Comment("Users that should be ignored for spam checks"),

		field.Int("ratelimit_message").Default(3).Comment("The amount of times the same message can be sent within the rate limit period before being considered spam"),

		field.Enum("ratelimit_time").Values().Values("30s", "1m", "2m", "3m", "4m", "5m").Default("2m").Comment("The ratelimit cooldown time, message tracking will be reset after this time period"),
		field.Enum("timeout_time").Values().Values("60s", "5m", "10m", "1h", "1d", "1w").Default("1h").Comment("The discord timeout time assigned to a spammer"),
		field.Enum("ban_delete_message_time").Values("1h", "6h", "12h", "1d", "3d", "1w").Default("1h").Comment("The discord time to remove messages sent by a spammer"),
	}
}

func (ServerConfig) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (ServerConfig) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("server", Server.Type).Ref("configuration").Unique(),
	}
}

func (ServerConfig) Annotations() []schema.Annotation {
	return nil
}
