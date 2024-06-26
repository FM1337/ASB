// Code generated by ent, DO NOT EDIT.

package server

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the server type in the database.
	Label = "server"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldServerID holds the string denoting the server_id field in the database.
	FieldServerID = "server_id"
	// FieldOwnerID holds the string denoting the owner_id field in the database.
	FieldOwnerID = "owner_id"
	// FieldEnabled holds the string denoting the enabled field in the database.
	FieldEnabled = "enabled"
	// EdgeConfiguration holds the string denoting the configuration edge name in mutations.
	EdgeConfiguration = "configuration"
	// EdgeSpammer holds the string denoting the spammer edge name in mutations.
	EdgeSpammer = "spammer"
	// EdgeWordBlacklist holds the string denoting the word_blacklist edge name in mutations.
	EdgeWordBlacklist = "word_blacklist"
	// EdgeCooldown holds the string denoting the cooldown edge name in mutations.
	EdgeCooldown = "cooldown"
	// Table holds the table name of the server in the database.
	Table = "servers"
	// ConfigurationTable is the table that holds the configuration relation/edge.
	ConfigurationTable = "server_configs"
	// ConfigurationInverseTable is the table name for the ServerConfig entity.
	// It exists in this package in order to avoid circular dependency with the "serverconfig" package.
	ConfigurationInverseTable = "server_configs"
	// ConfigurationColumn is the table column denoting the configuration relation/edge.
	ConfigurationColumn = "server_configuration"
	// SpammerTable is the table that holds the spammer relation/edge.
	SpammerTable = "spammers"
	// SpammerInverseTable is the table name for the Spammer entity.
	// It exists in this package in order to avoid circular dependency with the "spammer" package.
	SpammerInverseTable = "spammers"
	// SpammerColumn is the table column denoting the spammer relation/edge.
	SpammerColumn = "server_spammer"
	// WordBlacklistTable is the table that holds the word_blacklist relation/edge. The primary key declared below.
	WordBlacklistTable = "server_word_blacklist"
	// WordBlacklistInverseTable is the table name for the WordBlacklist entity.
	// It exists in this package in order to avoid circular dependency with the "wordblacklist" package.
	WordBlacklistInverseTable = "word_blacklists"
	// CooldownTable is the table that holds the cooldown relation/edge.
	CooldownTable = "cooldowns"
	// CooldownInverseTable is the table name for the Cooldown entity.
	// It exists in this package in order to avoid circular dependency with the "cooldown" package.
	CooldownInverseTable = "cooldowns"
	// CooldownColumn is the table column denoting the cooldown relation/edge.
	CooldownColumn = "server_cooldown"
)

// Columns holds all SQL columns for server fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldServerID,
	FieldOwnerID,
	FieldEnabled,
}

var (
	// WordBlacklistPrimaryKey and WordBlacklistColumn2 are the table columns denoting the
	// primary key for the word_blacklist relation (M2M).
	WordBlacklistPrimaryKey = []string{"server_id", "word_blacklist_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// DefaultEnabled holds the default value on creation for the "enabled" field.
	DefaultEnabled bool
)

// OrderOption defines the ordering options for the Server queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByServerID orders the results by the server_id field.
func ByServerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldServerID, opts...).ToFunc()
}

// ByOwnerID orders the results by the owner_id field.
func ByOwnerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOwnerID, opts...).ToFunc()
}

// ByEnabled orders the results by the enabled field.
func ByEnabled(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEnabled, opts...).ToFunc()
}

// ByConfigurationField orders the results by configuration field.
func ByConfigurationField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newConfigurationStep(), sql.OrderByField(field, opts...))
	}
}

// BySpammerField orders the results by spammer field.
func BySpammerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSpammerStep(), sql.OrderByField(field, opts...))
	}
}

// ByWordBlacklistCount orders the results by word_blacklist count.
func ByWordBlacklistCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWordBlacklistStep(), opts...)
	}
}

// ByWordBlacklist orders the results by word_blacklist terms.
func ByWordBlacklist(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWordBlacklistStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCooldownCount orders the results by cooldown count.
func ByCooldownCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCooldownStep(), opts...)
	}
}

// ByCooldown orders the results by cooldown terms.
func ByCooldown(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCooldownStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newConfigurationStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ConfigurationInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, ConfigurationTable, ConfigurationColumn),
	)
}
func newSpammerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SpammerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, SpammerTable, SpammerColumn),
	)
}
func newWordBlacklistStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WordBlacklistInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, WordBlacklistTable, WordBlacklistPrimaryKey...),
	)
}
func newCooldownStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CooldownInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CooldownTable, CooldownColumn),
	)
}
