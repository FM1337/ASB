// Code generated by ent, DO NOT EDIT.

package serverconfig

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/FM1337/ASB/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldUpdateTime, v))
}

// RemoveRoles applies equality check predicate on the "remove_roles" field. It's identical to RemoveRolesEQ.
func RemoveRoles(v bool) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldRemoveRoles, v))
}

// GiveRole applies equality check predicate on the "give_role" field. It's identical to GiveRoleEQ.
func GiveRole(v bool) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldGiveRole, v))
}

// Timeout applies equality check predicate on the "timeout" field. It's identical to TimeoutEQ.
func Timeout(v bool) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldTimeout, v))
}

// Kick applies equality check predicate on the "kick" field. It's identical to KickEQ.
func Kick(v bool) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldKick, v))
}

// Ban applies equality check predicate on the "ban" field. It's identical to BanEQ.
func Ban(v bool) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldBan, v))
}

// CheckInvites applies equality check predicate on the "check_invites" field. It's identical to CheckInvitesEQ.
func CheckInvites(v bool) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldCheckInvites, v))
}

// CheckLinks applies equality check predicate on the "check_links" field. It's identical to CheckLinksEQ.
func CheckLinks(v bool) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldCheckLinks, v))
}

// Ratelimit applies equality check predicate on the "ratelimit" field. It's identical to RatelimitEQ.
func Ratelimit(v bool) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldRatelimit, v))
}

// Alerts applies equality check predicate on the "alerts" field. It's identical to AlertsEQ.
func Alerts(v bool) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldAlerts, v))
}

// FlagLinks applies equality check predicate on the "flag_links" field. It's identical to FlagLinksEQ.
func FlagLinks(v bool) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldFlagLinks, v))
}

// LogChannel applies equality check predicate on the "log_channel" field. It's identical to LogChannelEQ.
func LogChannel(v string) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldLogChannel, v))
}

// GivenRole applies equality check predicate on the "given_role" field. It's identical to GivenRoleEQ.
func GivenRole(v string) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldGivenRole, v))
}

// RatelimitMessage applies equality check predicate on the "ratelimit_message" field. It's identical to RatelimitMessageEQ.
func RatelimitMessage(v int) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldRatelimitMessage, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldLTE(FieldUpdateTime, v))
}

// RemoveRolesEQ applies the EQ predicate on the "remove_roles" field.
func RemoveRolesEQ(v bool) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldRemoveRoles, v))
}

// RemoveRolesNEQ applies the NEQ predicate on the "remove_roles" field.
func RemoveRolesNEQ(v bool) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldNEQ(FieldRemoveRoles, v))
}

// GiveRoleEQ applies the EQ predicate on the "give_role" field.
func GiveRoleEQ(v bool) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldGiveRole, v))
}

// GiveRoleNEQ applies the NEQ predicate on the "give_role" field.
func GiveRoleNEQ(v bool) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldNEQ(FieldGiveRole, v))
}

// TimeoutEQ applies the EQ predicate on the "timeout" field.
func TimeoutEQ(v bool) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldTimeout, v))
}

// TimeoutNEQ applies the NEQ predicate on the "timeout" field.
func TimeoutNEQ(v bool) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldNEQ(FieldTimeout, v))
}

// KickEQ applies the EQ predicate on the "kick" field.
func KickEQ(v bool) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldKick, v))
}

// KickNEQ applies the NEQ predicate on the "kick" field.
func KickNEQ(v bool) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldNEQ(FieldKick, v))
}

// BanEQ applies the EQ predicate on the "ban" field.
func BanEQ(v bool) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldBan, v))
}

// BanNEQ applies the NEQ predicate on the "ban" field.
func BanNEQ(v bool) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldNEQ(FieldBan, v))
}

// CheckInvitesEQ applies the EQ predicate on the "check_invites" field.
func CheckInvitesEQ(v bool) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldCheckInvites, v))
}

// CheckInvitesNEQ applies the NEQ predicate on the "check_invites" field.
func CheckInvitesNEQ(v bool) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldNEQ(FieldCheckInvites, v))
}

// CheckLinksEQ applies the EQ predicate on the "check_links" field.
func CheckLinksEQ(v bool) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldCheckLinks, v))
}

// CheckLinksNEQ applies the NEQ predicate on the "check_links" field.
func CheckLinksNEQ(v bool) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldNEQ(FieldCheckLinks, v))
}

// RatelimitEQ applies the EQ predicate on the "ratelimit" field.
func RatelimitEQ(v bool) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldRatelimit, v))
}

// RatelimitNEQ applies the NEQ predicate on the "ratelimit" field.
func RatelimitNEQ(v bool) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldNEQ(FieldRatelimit, v))
}

// AlertsEQ applies the EQ predicate on the "alerts" field.
func AlertsEQ(v bool) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldAlerts, v))
}

// AlertsNEQ applies the NEQ predicate on the "alerts" field.
func AlertsNEQ(v bool) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldNEQ(FieldAlerts, v))
}

// FlagLinksEQ applies the EQ predicate on the "flag_links" field.
func FlagLinksEQ(v bool) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldFlagLinks, v))
}

// FlagLinksNEQ applies the NEQ predicate on the "flag_links" field.
func FlagLinksNEQ(v bool) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldNEQ(FieldFlagLinks, v))
}

// LogChannelEQ applies the EQ predicate on the "log_channel" field.
func LogChannelEQ(v string) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldLogChannel, v))
}

// LogChannelNEQ applies the NEQ predicate on the "log_channel" field.
func LogChannelNEQ(v string) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldNEQ(FieldLogChannel, v))
}

// LogChannelIn applies the In predicate on the "log_channel" field.
func LogChannelIn(vs ...string) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldIn(FieldLogChannel, vs...))
}

// LogChannelNotIn applies the NotIn predicate on the "log_channel" field.
func LogChannelNotIn(vs ...string) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldNotIn(FieldLogChannel, vs...))
}

// LogChannelGT applies the GT predicate on the "log_channel" field.
func LogChannelGT(v string) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldGT(FieldLogChannel, v))
}

// LogChannelGTE applies the GTE predicate on the "log_channel" field.
func LogChannelGTE(v string) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldGTE(FieldLogChannel, v))
}

// LogChannelLT applies the LT predicate on the "log_channel" field.
func LogChannelLT(v string) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldLT(FieldLogChannel, v))
}

// LogChannelLTE applies the LTE predicate on the "log_channel" field.
func LogChannelLTE(v string) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldLTE(FieldLogChannel, v))
}

// LogChannelContains applies the Contains predicate on the "log_channel" field.
func LogChannelContains(v string) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldContains(FieldLogChannel, v))
}

// LogChannelHasPrefix applies the HasPrefix predicate on the "log_channel" field.
func LogChannelHasPrefix(v string) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldHasPrefix(FieldLogChannel, v))
}

// LogChannelHasSuffix applies the HasSuffix predicate on the "log_channel" field.
func LogChannelHasSuffix(v string) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldHasSuffix(FieldLogChannel, v))
}

// LogChannelEqualFold applies the EqualFold predicate on the "log_channel" field.
func LogChannelEqualFold(v string) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEqualFold(FieldLogChannel, v))
}

// LogChannelContainsFold applies the ContainsFold predicate on the "log_channel" field.
func LogChannelContainsFold(v string) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldContainsFold(FieldLogChannel, v))
}

// GivenRoleEQ applies the EQ predicate on the "given_role" field.
func GivenRoleEQ(v string) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldGivenRole, v))
}

// GivenRoleNEQ applies the NEQ predicate on the "given_role" field.
func GivenRoleNEQ(v string) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldNEQ(FieldGivenRole, v))
}

// GivenRoleIn applies the In predicate on the "given_role" field.
func GivenRoleIn(vs ...string) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldIn(FieldGivenRole, vs...))
}

// GivenRoleNotIn applies the NotIn predicate on the "given_role" field.
func GivenRoleNotIn(vs ...string) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldNotIn(FieldGivenRole, vs...))
}

// GivenRoleGT applies the GT predicate on the "given_role" field.
func GivenRoleGT(v string) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldGT(FieldGivenRole, v))
}

// GivenRoleGTE applies the GTE predicate on the "given_role" field.
func GivenRoleGTE(v string) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldGTE(FieldGivenRole, v))
}

// GivenRoleLT applies the LT predicate on the "given_role" field.
func GivenRoleLT(v string) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldLT(FieldGivenRole, v))
}

// GivenRoleLTE applies the LTE predicate on the "given_role" field.
func GivenRoleLTE(v string) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldLTE(FieldGivenRole, v))
}

// GivenRoleContains applies the Contains predicate on the "given_role" field.
func GivenRoleContains(v string) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldContains(FieldGivenRole, v))
}

// GivenRoleHasPrefix applies the HasPrefix predicate on the "given_role" field.
func GivenRoleHasPrefix(v string) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldHasPrefix(FieldGivenRole, v))
}

// GivenRoleHasSuffix applies the HasSuffix predicate on the "given_role" field.
func GivenRoleHasSuffix(v string) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldHasSuffix(FieldGivenRole, v))
}

// GivenRoleEqualFold applies the EqualFold predicate on the "given_role" field.
func GivenRoleEqualFold(v string) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEqualFold(FieldGivenRole, v))
}

// GivenRoleContainsFold applies the ContainsFold predicate on the "given_role" field.
func GivenRoleContainsFold(v string) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldContainsFold(FieldGivenRole, v))
}

// RatelimitMessageEQ applies the EQ predicate on the "ratelimit_message" field.
func RatelimitMessageEQ(v int) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldRatelimitMessage, v))
}

// RatelimitMessageNEQ applies the NEQ predicate on the "ratelimit_message" field.
func RatelimitMessageNEQ(v int) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldNEQ(FieldRatelimitMessage, v))
}

// RatelimitMessageIn applies the In predicate on the "ratelimit_message" field.
func RatelimitMessageIn(vs ...int) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldIn(FieldRatelimitMessage, vs...))
}

// RatelimitMessageNotIn applies the NotIn predicate on the "ratelimit_message" field.
func RatelimitMessageNotIn(vs ...int) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldNotIn(FieldRatelimitMessage, vs...))
}

// RatelimitMessageGT applies the GT predicate on the "ratelimit_message" field.
func RatelimitMessageGT(v int) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldGT(FieldRatelimitMessage, v))
}

// RatelimitMessageGTE applies the GTE predicate on the "ratelimit_message" field.
func RatelimitMessageGTE(v int) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldGTE(FieldRatelimitMessage, v))
}

// RatelimitMessageLT applies the LT predicate on the "ratelimit_message" field.
func RatelimitMessageLT(v int) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldLT(FieldRatelimitMessage, v))
}

// RatelimitMessageLTE applies the LTE predicate on the "ratelimit_message" field.
func RatelimitMessageLTE(v int) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldLTE(FieldRatelimitMessage, v))
}

// RatelimitTimeEQ applies the EQ predicate on the "ratelimit_time" field.
func RatelimitTimeEQ(v RatelimitTime) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldRatelimitTime, v))
}

// RatelimitTimeNEQ applies the NEQ predicate on the "ratelimit_time" field.
func RatelimitTimeNEQ(v RatelimitTime) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldNEQ(FieldRatelimitTime, v))
}

// RatelimitTimeIn applies the In predicate on the "ratelimit_time" field.
func RatelimitTimeIn(vs ...RatelimitTime) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldIn(FieldRatelimitTime, vs...))
}

// RatelimitTimeNotIn applies the NotIn predicate on the "ratelimit_time" field.
func RatelimitTimeNotIn(vs ...RatelimitTime) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldNotIn(FieldRatelimitTime, vs...))
}

// TimeoutTimeEQ applies the EQ predicate on the "timeout_time" field.
func TimeoutTimeEQ(v TimeoutTime) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldTimeoutTime, v))
}

// TimeoutTimeNEQ applies the NEQ predicate on the "timeout_time" field.
func TimeoutTimeNEQ(v TimeoutTime) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldNEQ(FieldTimeoutTime, v))
}

// TimeoutTimeIn applies the In predicate on the "timeout_time" field.
func TimeoutTimeIn(vs ...TimeoutTime) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldIn(FieldTimeoutTime, vs...))
}

// TimeoutTimeNotIn applies the NotIn predicate on the "timeout_time" field.
func TimeoutTimeNotIn(vs ...TimeoutTime) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldNotIn(FieldTimeoutTime, vs...))
}

// BanDeleteMessageTimeEQ applies the EQ predicate on the "ban_delete_message_time" field.
func BanDeleteMessageTimeEQ(v BanDeleteMessageTime) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldEQ(FieldBanDeleteMessageTime, v))
}

// BanDeleteMessageTimeNEQ applies the NEQ predicate on the "ban_delete_message_time" field.
func BanDeleteMessageTimeNEQ(v BanDeleteMessageTime) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldNEQ(FieldBanDeleteMessageTime, v))
}

// BanDeleteMessageTimeIn applies the In predicate on the "ban_delete_message_time" field.
func BanDeleteMessageTimeIn(vs ...BanDeleteMessageTime) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldIn(FieldBanDeleteMessageTime, vs...))
}

// BanDeleteMessageTimeNotIn applies the NotIn predicate on the "ban_delete_message_time" field.
func BanDeleteMessageTimeNotIn(vs ...BanDeleteMessageTime) predicate.ServerConfig {
	return predicate.ServerConfig(sql.FieldNotIn(FieldBanDeleteMessageTime, vs...))
}

// HasServer applies the HasEdge predicate on the "server" edge.
func HasServer() predicate.ServerConfig {
	return predicate.ServerConfig(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, ServerTable, ServerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasServerWith applies the HasEdge predicate on the "server" edge with a given conditions (other predicates).
func HasServerWith(preds ...predicate.Server) predicate.ServerConfig {
	return predicate.ServerConfig(func(s *sql.Selector) {
		step := newServerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ServerConfig) predicate.ServerConfig {
	return predicate.ServerConfig(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ServerConfig) predicate.ServerConfig {
	return predicate.ServerConfig(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ServerConfig) predicate.ServerConfig {
	return predicate.ServerConfig(sql.NotPredicates(p))
}
