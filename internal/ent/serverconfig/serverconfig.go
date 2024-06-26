// Code generated by ent, DO NOT EDIT.

package serverconfig

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the serverconfig type in the database.
	Label = "server_config"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldRemoveRoles holds the string denoting the remove_roles field in the database.
	FieldRemoveRoles = "remove_roles"
	// FieldGiveRole holds the string denoting the give_role field in the database.
	FieldGiveRole = "give_role"
	// FieldTimeout holds the string denoting the timeout field in the database.
	FieldTimeout = "timeout"
	// FieldKick holds the string denoting the kick field in the database.
	FieldKick = "kick"
	// FieldBan holds the string denoting the ban field in the database.
	FieldBan = "ban"
	// FieldCheckInvites holds the string denoting the check_invites field in the database.
	FieldCheckInvites = "check_invites"
	// FieldCheckLinks holds the string denoting the check_links field in the database.
	FieldCheckLinks = "check_links"
	// FieldRatelimit holds the string denoting the ratelimit field in the database.
	FieldRatelimit = "ratelimit"
	// FieldAlerts holds the string denoting the alerts field in the database.
	FieldAlerts = "alerts"
	// FieldFlagLinks holds the string denoting the flag_links field in the database.
	FieldFlagLinks = "flag_links"
	// FieldLogChannel holds the string denoting the log_channel field in the database.
	FieldLogChannel = "log_channel"
	// FieldGivenRole holds the string denoting the given_role field in the database.
	FieldGivenRole = "given_role"
	// FieldExcludedChannels holds the string denoting the excluded_channels field in the database.
	FieldExcludedChannels = "excluded_channels"
	// FieldExcludedRoles holds the string denoting the excluded_roles field in the database.
	FieldExcludedRoles = "excluded_roles"
	// FieldExcludedUsers holds the string denoting the excluded_users field in the database.
	FieldExcludedUsers = "excluded_users"
	// FieldRatelimitMessage holds the string denoting the ratelimit_message field in the database.
	FieldRatelimitMessage = "ratelimit_message"
	// FieldRatelimitTime holds the string denoting the ratelimit_time field in the database.
	FieldRatelimitTime = "ratelimit_time"
	// FieldTimeoutTime holds the string denoting the timeout_time field in the database.
	FieldTimeoutTime = "timeout_time"
	// FieldBanDeleteMessageTime holds the string denoting the ban_delete_message_time field in the database.
	FieldBanDeleteMessageTime = "ban_delete_message_time"
	// EdgeServer holds the string denoting the server edge name in mutations.
	EdgeServer = "server"
	// Table holds the table name of the serverconfig in the database.
	Table = "server_configs"
	// ServerTable is the table that holds the server relation/edge.
	ServerTable = "server_configs"
	// ServerInverseTable is the table name for the Server entity.
	// It exists in this package in order to avoid circular dependency with the "server" package.
	ServerInverseTable = "servers"
	// ServerColumn is the table column denoting the server relation/edge.
	ServerColumn = "server_configuration"
)

// Columns holds all SQL columns for serverconfig fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldRemoveRoles,
	FieldGiveRole,
	FieldTimeout,
	FieldKick,
	FieldBan,
	FieldCheckInvites,
	FieldCheckLinks,
	FieldRatelimit,
	FieldAlerts,
	FieldFlagLinks,
	FieldLogChannel,
	FieldGivenRole,
	FieldExcludedChannels,
	FieldExcludedRoles,
	FieldExcludedUsers,
	FieldRatelimitMessage,
	FieldRatelimitTime,
	FieldTimeoutTime,
	FieldBanDeleteMessageTime,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "server_configs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"server_configuration",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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
	// DefaultRemoveRoles holds the default value on creation for the "remove_roles" field.
	DefaultRemoveRoles bool
	// DefaultGiveRole holds the default value on creation for the "give_role" field.
	DefaultGiveRole bool
	// DefaultTimeout holds the default value on creation for the "timeout" field.
	DefaultTimeout bool
	// DefaultKick holds the default value on creation for the "kick" field.
	DefaultKick bool
	// DefaultBan holds the default value on creation for the "ban" field.
	DefaultBan bool
	// DefaultCheckInvites holds the default value on creation for the "check_invites" field.
	DefaultCheckInvites bool
	// DefaultCheckLinks holds the default value on creation for the "check_links" field.
	DefaultCheckLinks bool
	// DefaultRatelimit holds the default value on creation for the "ratelimit" field.
	DefaultRatelimit bool
	// DefaultAlerts holds the default value on creation for the "alerts" field.
	DefaultAlerts bool
	// DefaultFlagLinks holds the default value on creation for the "flag_links" field.
	DefaultFlagLinks bool
	// DefaultLogChannel holds the default value on creation for the "log_channel" field.
	DefaultLogChannel string
	// DefaultGivenRole holds the default value on creation for the "given_role" field.
	DefaultGivenRole string
	// DefaultExcludedChannels holds the default value on creation for the "excluded_channels" field.
	DefaultExcludedChannels []string
	// DefaultExcludedRoles holds the default value on creation for the "excluded_roles" field.
	DefaultExcludedRoles []string
	// DefaultExcludedUsers holds the default value on creation for the "excluded_users" field.
	DefaultExcludedUsers []string
	// DefaultRatelimitMessage holds the default value on creation for the "ratelimit_message" field.
	DefaultRatelimitMessage int
)

// RatelimitTime defines the type for the "ratelimit_time" enum field.
type RatelimitTime string

// RatelimitTime5m is the default value of the RatelimitTime enum.
const DefaultRatelimitTime = RatelimitTime5m

// RatelimitTime values.
const (
	RatelimitTime30s RatelimitTime = "30s"
	RatelimitTime1m  RatelimitTime = "1m"
	RatelimitTime2m  RatelimitTime = "2m"
	RatelimitTime3m  RatelimitTime = "3m"
	RatelimitTime4m  RatelimitTime = "4m"
	RatelimitTime5m  RatelimitTime = "5m"
)

func (rt RatelimitTime) String() string {
	return string(rt)
}

// RatelimitTimeValidator is a validator for the "ratelimit_time" field enum values. It is called by the builders before save.
func RatelimitTimeValidator(rt RatelimitTime) error {
	switch rt {
	case RatelimitTime30s, RatelimitTime1m, RatelimitTime2m, RatelimitTime3m, RatelimitTime4m, RatelimitTime5m:
		return nil
	default:
		return fmt.Errorf("serverconfig: invalid enum value for ratelimit_time field: %q", rt)
	}
}

// TimeoutTime defines the type for the "timeout_time" enum field.
type TimeoutTime string

// TimeoutTime1h is the default value of the TimeoutTime enum.
const DefaultTimeoutTime = TimeoutTime1h

// TimeoutTime values.
const (
	TimeoutTime60s TimeoutTime = "60s"
	TimeoutTime5m  TimeoutTime = "5m"
	TimeoutTime10m TimeoutTime = "10m"
	TimeoutTime1h  TimeoutTime = "1h"
	TimeoutTime1d  TimeoutTime = "1d"
	TimeoutTime1w  TimeoutTime = "1w"
)

func (tt TimeoutTime) String() string {
	return string(tt)
}

// TimeoutTimeValidator is a validator for the "timeout_time" field enum values. It is called by the builders before save.
func TimeoutTimeValidator(tt TimeoutTime) error {
	switch tt {
	case TimeoutTime60s, TimeoutTime5m, TimeoutTime10m, TimeoutTime1h, TimeoutTime1d, TimeoutTime1w:
		return nil
	default:
		return fmt.Errorf("serverconfig: invalid enum value for timeout_time field: %q", tt)
	}
}

// BanDeleteMessageTime defines the type for the "ban_delete_message_time" enum field.
type BanDeleteMessageTime string

// BanDeleteMessageTime1h is the default value of the BanDeleteMessageTime enum.
const DefaultBanDeleteMessageTime = BanDeleteMessageTime1h

// BanDeleteMessageTime values.
const (
	BanDeleteMessageTime1h  BanDeleteMessageTime = "1h"
	BanDeleteMessageTime6h  BanDeleteMessageTime = "6h"
	BanDeleteMessageTime12h BanDeleteMessageTime = "12h"
	BanDeleteMessageTime1d  BanDeleteMessageTime = "1d"
	BanDeleteMessageTime3d  BanDeleteMessageTime = "3d"
	BanDeleteMessageTime1w  BanDeleteMessageTime = "1w"
)

func (bdmt BanDeleteMessageTime) String() string {
	return string(bdmt)
}

// BanDeleteMessageTimeValidator is a validator for the "ban_delete_message_time" field enum values. It is called by the builders before save.
func BanDeleteMessageTimeValidator(bdmt BanDeleteMessageTime) error {
	switch bdmt {
	case BanDeleteMessageTime1h, BanDeleteMessageTime6h, BanDeleteMessageTime12h, BanDeleteMessageTime1d, BanDeleteMessageTime3d, BanDeleteMessageTime1w:
		return nil
	default:
		return fmt.Errorf("serverconfig: invalid enum value for ban_delete_message_time field: %q", bdmt)
	}
}

// OrderOption defines the ordering options for the ServerConfig queries.
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

// ByRemoveRoles orders the results by the remove_roles field.
func ByRemoveRoles(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRemoveRoles, opts...).ToFunc()
}

// ByGiveRole orders the results by the give_role field.
func ByGiveRole(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGiveRole, opts...).ToFunc()
}

// ByTimeout orders the results by the timeout field.
func ByTimeout(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTimeout, opts...).ToFunc()
}

// ByKick orders the results by the kick field.
func ByKick(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKick, opts...).ToFunc()
}

// ByBan orders the results by the ban field.
func ByBan(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBan, opts...).ToFunc()
}

// ByCheckInvites orders the results by the check_invites field.
func ByCheckInvites(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCheckInvites, opts...).ToFunc()
}

// ByCheckLinks orders the results by the check_links field.
func ByCheckLinks(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCheckLinks, opts...).ToFunc()
}

// ByRatelimit orders the results by the ratelimit field.
func ByRatelimit(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRatelimit, opts...).ToFunc()
}

// ByAlerts orders the results by the alerts field.
func ByAlerts(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAlerts, opts...).ToFunc()
}

// ByFlagLinks orders the results by the flag_links field.
func ByFlagLinks(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFlagLinks, opts...).ToFunc()
}

// ByLogChannel orders the results by the log_channel field.
func ByLogChannel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLogChannel, opts...).ToFunc()
}

// ByGivenRole orders the results by the given_role field.
func ByGivenRole(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGivenRole, opts...).ToFunc()
}

// ByRatelimitMessage orders the results by the ratelimit_message field.
func ByRatelimitMessage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRatelimitMessage, opts...).ToFunc()
}

// ByRatelimitTime orders the results by the ratelimit_time field.
func ByRatelimitTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRatelimitTime, opts...).ToFunc()
}

// ByTimeoutTime orders the results by the timeout_time field.
func ByTimeoutTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTimeoutTime, opts...).ToFunc()
}

// ByBanDeleteMessageTime orders the results by the ban_delete_message_time field.
func ByBanDeleteMessageTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBanDeleteMessageTime, opts...).ToFunc()
}

// ByServerField orders the results by server field.
func ByServerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newServerStep(), sql.OrderByField(field, opts...))
	}
}
func newServerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ServerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, ServerTable, ServerColumn),
	)
}
