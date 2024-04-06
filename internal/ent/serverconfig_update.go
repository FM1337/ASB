// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/FM1337/ASB/internal/ent/predicate"
	"github.com/FM1337/ASB/internal/ent/server"
	"github.com/FM1337/ASB/internal/ent/serverconfig"
)

// ServerConfigUpdate is the builder for updating ServerConfig entities.
type ServerConfigUpdate struct {
	config
	hooks     []Hook
	mutation  *ServerConfigMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ServerConfigUpdate builder.
func (scu *ServerConfigUpdate) Where(ps ...predicate.ServerConfig) *ServerConfigUpdate {
	scu.mutation.Where(ps...)
	return scu
}

// SetUpdateTime sets the "update_time" field.
func (scu *ServerConfigUpdate) SetUpdateTime(t time.Time) *ServerConfigUpdate {
	scu.mutation.SetUpdateTime(t)
	return scu
}

// SetRemoveRoles sets the "remove_roles" field.
func (scu *ServerConfigUpdate) SetRemoveRoles(b bool) *ServerConfigUpdate {
	scu.mutation.SetRemoveRoles(b)
	return scu
}

// SetNillableRemoveRoles sets the "remove_roles" field if the given value is not nil.
func (scu *ServerConfigUpdate) SetNillableRemoveRoles(b *bool) *ServerConfigUpdate {
	if b != nil {
		scu.SetRemoveRoles(*b)
	}
	return scu
}

// SetGiveRole sets the "give_role" field.
func (scu *ServerConfigUpdate) SetGiveRole(b bool) *ServerConfigUpdate {
	scu.mutation.SetGiveRole(b)
	return scu
}

// SetNillableGiveRole sets the "give_role" field if the given value is not nil.
func (scu *ServerConfigUpdate) SetNillableGiveRole(b *bool) *ServerConfigUpdate {
	if b != nil {
		scu.SetGiveRole(*b)
	}
	return scu
}

// SetTimeout sets the "timeout" field.
func (scu *ServerConfigUpdate) SetTimeout(b bool) *ServerConfigUpdate {
	scu.mutation.SetTimeout(b)
	return scu
}

// SetNillableTimeout sets the "timeout" field if the given value is not nil.
func (scu *ServerConfigUpdate) SetNillableTimeout(b *bool) *ServerConfigUpdate {
	if b != nil {
		scu.SetTimeout(*b)
	}
	return scu
}

// SetKick sets the "kick" field.
func (scu *ServerConfigUpdate) SetKick(b bool) *ServerConfigUpdate {
	scu.mutation.SetKick(b)
	return scu
}

// SetNillableKick sets the "kick" field if the given value is not nil.
func (scu *ServerConfigUpdate) SetNillableKick(b *bool) *ServerConfigUpdate {
	if b != nil {
		scu.SetKick(*b)
	}
	return scu
}

// SetBan sets the "ban" field.
func (scu *ServerConfigUpdate) SetBan(b bool) *ServerConfigUpdate {
	scu.mutation.SetBan(b)
	return scu
}

// SetNillableBan sets the "ban" field if the given value is not nil.
func (scu *ServerConfigUpdate) SetNillableBan(b *bool) *ServerConfigUpdate {
	if b != nil {
		scu.SetBan(*b)
	}
	return scu
}

// SetCheckInvites sets the "check_invites" field.
func (scu *ServerConfigUpdate) SetCheckInvites(b bool) *ServerConfigUpdate {
	scu.mutation.SetCheckInvites(b)
	return scu
}

// SetNillableCheckInvites sets the "check_invites" field if the given value is not nil.
func (scu *ServerConfigUpdate) SetNillableCheckInvites(b *bool) *ServerConfigUpdate {
	if b != nil {
		scu.SetCheckInvites(*b)
	}
	return scu
}

// SetCheckLinks sets the "check_links" field.
func (scu *ServerConfigUpdate) SetCheckLinks(b bool) *ServerConfigUpdate {
	scu.mutation.SetCheckLinks(b)
	return scu
}

// SetNillableCheckLinks sets the "check_links" field if the given value is not nil.
func (scu *ServerConfigUpdate) SetNillableCheckLinks(b *bool) *ServerConfigUpdate {
	if b != nil {
		scu.SetCheckLinks(*b)
	}
	return scu
}

// SetRatelimit sets the "ratelimit" field.
func (scu *ServerConfigUpdate) SetRatelimit(b bool) *ServerConfigUpdate {
	scu.mutation.SetRatelimit(b)
	return scu
}

// SetNillableRatelimit sets the "ratelimit" field if the given value is not nil.
func (scu *ServerConfigUpdate) SetNillableRatelimit(b *bool) *ServerConfigUpdate {
	if b != nil {
		scu.SetRatelimit(*b)
	}
	return scu
}

// SetAlerts sets the "alerts" field.
func (scu *ServerConfigUpdate) SetAlerts(b bool) *ServerConfigUpdate {
	scu.mutation.SetAlerts(b)
	return scu
}

// SetNillableAlerts sets the "alerts" field if the given value is not nil.
func (scu *ServerConfigUpdate) SetNillableAlerts(b *bool) *ServerConfigUpdate {
	if b != nil {
		scu.SetAlerts(*b)
	}
	return scu
}

// SetFlagLinks sets the "flag_links" field.
func (scu *ServerConfigUpdate) SetFlagLinks(b bool) *ServerConfigUpdate {
	scu.mutation.SetFlagLinks(b)
	return scu
}

// SetNillableFlagLinks sets the "flag_links" field if the given value is not nil.
func (scu *ServerConfigUpdate) SetNillableFlagLinks(b *bool) *ServerConfigUpdate {
	if b != nil {
		scu.SetFlagLinks(*b)
	}
	return scu
}

// SetLogChannel sets the "log_channel" field.
func (scu *ServerConfigUpdate) SetLogChannel(s string) *ServerConfigUpdate {
	scu.mutation.SetLogChannel(s)
	return scu
}

// SetNillableLogChannel sets the "log_channel" field if the given value is not nil.
func (scu *ServerConfigUpdate) SetNillableLogChannel(s *string) *ServerConfigUpdate {
	if s != nil {
		scu.SetLogChannel(*s)
	}
	return scu
}

// SetExcludedChannels sets the "excluded_channels" field.
func (scu *ServerConfigUpdate) SetExcludedChannels(s []string) *ServerConfigUpdate {
	scu.mutation.SetExcludedChannels(s)
	return scu
}

// AppendExcludedChannels appends s to the "excluded_channels" field.
func (scu *ServerConfigUpdate) AppendExcludedChannels(s []string) *ServerConfigUpdate {
	scu.mutation.AppendExcludedChannels(s)
	return scu
}

// SetExcludedRoles sets the "excluded_roles" field.
func (scu *ServerConfigUpdate) SetExcludedRoles(s []string) *ServerConfigUpdate {
	scu.mutation.SetExcludedRoles(s)
	return scu
}

// AppendExcludedRoles appends s to the "excluded_roles" field.
func (scu *ServerConfigUpdate) AppendExcludedRoles(s []string) *ServerConfigUpdate {
	scu.mutation.AppendExcludedRoles(s)
	return scu
}

// SetExcludedUsers sets the "excluded_users" field.
func (scu *ServerConfigUpdate) SetExcludedUsers(s []string) *ServerConfigUpdate {
	scu.mutation.SetExcludedUsers(s)
	return scu
}

// AppendExcludedUsers appends s to the "excluded_users" field.
func (scu *ServerConfigUpdate) AppendExcludedUsers(s []string) *ServerConfigUpdate {
	scu.mutation.AppendExcludedUsers(s)
	return scu
}

// SetGivenRole sets the "given_role" field.
func (scu *ServerConfigUpdate) SetGivenRole(s string) *ServerConfigUpdate {
	scu.mutation.SetGivenRole(s)
	return scu
}

// SetNillableGivenRole sets the "given_role" field if the given value is not nil.
func (scu *ServerConfigUpdate) SetNillableGivenRole(s *string) *ServerConfigUpdate {
	if s != nil {
		scu.SetGivenRole(*s)
	}
	return scu
}

// SetRatelimitMessage sets the "ratelimit_message" field.
func (scu *ServerConfigUpdate) SetRatelimitMessage(i int) *ServerConfigUpdate {
	scu.mutation.ResetRatelimitMessage()
	scu.mutation.SetRatelimitMessage(i)
	return scu
}

// SetNillableRatelimitMessage sets the "ratelimit_message" field if the given value is not nil.
func (scu *ServerConfigUpdate) SetNillableRatelimitMessage(i *int) *ServerConfigUpdate {
	if i != nil {
		scu.SetRatelimitMessage(*i)
	}
	return scu
}

// AddRatelimitMessage adds i to the "ratelimit_message" field.
func (scu *ServerConfigUpdate) AddRatelimitMessage(i int) *ServerConfigUpdate {
	scu.mutation.AddRatelimitMessage(i)
	return scu
}

// SetRatelimitTime sets the "ratelimit_time" field.
func (scu *ServerConfigUpdate) SetRatelimitTime(st serverconfig.RatelimitTime) *ServerConfigUpdate {
	scu.mutation.SetRatelimitTime(st)
	return scu
}

// SetNillableRatelimitTime sets the "ratelimit_time" field if the given value is not nil.
func (scu *ServerConfigUpdate) SetNillableRatelimitTime(st *serverconfig.RatelimitTime) *ServerConfigUpdate {
	if st != nil {
		scu.SetRatelimitTime(*st)
	}
	return scu
}

// SetTimeoutTime sets the "timeout_time" field.
func (scu *ServerConfigUpdate) SetTimeoutTime(st serverconfig.TimeoutTime) *ServerConfigUpdate {
	scu.mutation.SetTimeoutTime(st)
	return scu
}

// SetNillableTimeoutTime sets the "timeout_time" field if the given value is not nil.
func (scu *ServerConfigUpdate) SetNillableTimeoutTime(st *serverconfig.TimeoutTime) *ServerConfigUpdate {
	if st != nil {
		scu.SetTimeoutTime(*st)
	}
	return scu
}

// SetBanDeleteMessageTime sets the "ban_delete_message_time" field.
func (scu *ServerConfigUpdate) SetBanDeleteMessageTime(sdmt serverconfig.BanDeleteMessageTime) *ServerConfigUpdate {
	scu.mutation.SetBanDeleteMessageTime(sdmt)
	return scu
}

// SetNillableBanDeleteMessageTime sets the "ban_delete_message_time" field if the given value is not nil.
func (scu *ServerConfigUpdate) SetNillableBanDeleteMessageTime(sdmt *serverconfig.BanDeleteMessageTime) *ServerConfigUpdate {
	if sdmt != nil {
		scu.SetBanDeleteMessageTime(*sdmt)
	}
	return scu
}

// AddServerIDs adds the "server" edge to the Server entity by IDs.
func (scu *ServerConfigUpdate) AddServerIDs(ids ...string) *ServerConfigUpdate {
	scu.mutation.AddServerIDs(ids...)
	return scu
}

// AddServer adds the "server" edges to the Server entity.
func (scu *ServerConfigUpdate) AddServer(s ...*Server) *ServerConfigUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return scu.AddServerIDs(ids...)
}

// Mutation returns the ServerConfigMutation object of the builder.
func (scu *ServerConfigUpdate) Mutation() *ServerConfigMutation {
	return scu.mutation
}

// ClearServer clears all "server" edges to the Server entity.
func (scu *ServerConfigUpdate) ClearServer() *ServerConfigUpdate {
	scu.mutation.ClearServer()
	return scu
}

// RemoveServerIDs removes the "server" edge to Server entities by IDs.
func (scu *ServerConfigUpdate) RemoveServerIDs(ids ...string) *ServerConfigUpdate {
	scu.mutation.RemoveServerIDs(ids...)
	return scu
}

// RemoveServer removes "server" edges to Server entities.
func (scu *ServerConfigUpdate) RemoveServer(s ...*Server) *ServerConfigUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return scu.RemoveServerIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (scu *ServerConfigUpdate) Save(ctx context.Context) (int, error) {
	scu.defaults()
	return withHooks(ctx, scu.sqlSave, scu.mutation, scu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (scu *ServerConfigUpdate) SaveX(ctx context.Context) int {
	affected, err := scu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (scu *ServerConfigUpdate) Exec(ctx context.Context) error {
	_, err := scu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scu *ServerConfigUpdate) ExecX(ctx context.Context) {
	if err := scu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (scu *ServerConfigUpdate) defaults() {
	if _, ok := scu.mutation.UpdateTime(); !ok {
		v := serverconfig.UpdateDefaultUpdateTime()
		scu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (scu *ServerConfigUpdate) check() error {
	if v, ok := scu.mutation.RatelimitTime(); ok {
		if err := serverconfig.RatelimitTimeValidator(v); err != nil {
			return &ValidationError{Name: "ratelimit_time", err: fmt.Errorf(`ent: validator failed for field "ServerConfig.ratelimit_time": %w`, err)}
		}
	}
	if v, ok := scu.mutation.TimeoutTime(); ok {
		if err := serverconfig.TimeoutTimeValidator(v); err != nil {
			return &ValidationError{Name: "timeout_time", err: fmt.Errorf(`ent: validator failed for field "ServerConfig.timeout_time": %w`, err)}
		}
	}
	if v, ok := scu.mutation.BanDeleteMessageTime(); ok {
		if err := serverconfig.BanDeleteMessageTimeValidator(v); err != nil {
			return &ValidationError{Name: "ban_delete_message_time", err: fmt.Errorf(`ent: validator failed for field "ServerConfig.ban_delete_message_time": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (scu *ServerConfigUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ServerConfigUpdate {
	scu.modifiers = append(scu.modifiers, modifiers...)
	return scu
}

func (scu *ServerConfigUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := scu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(serverconfig.Table, serverconfig.Columns, sqlgraph.NewFieldSpec(serverconfig.FieldID, field.TypeInt))
	if ps := scu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := scu.mutation.UpdateTime(); ok {
		_spec.SetField(serverconfig.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := scu.mutation.RemoveRoles(); ok {
		_spec.SetField(serverconfig.FieldRemoveRoles, field.TypeBool, value)
	}
	if value, ok := scu.mutation.GiveRole(); ok {
		_spec.SetField(serverconfig.FieldGiveRole, field.TypeBool, value)
	}
	if value, ok := scu.mutation.Timeout(); ok {
		_spec.SetField(serverconfig.FieldTimeout, field.TypeBool, value)
	}
	if value, ok := scu.mutation.Kick(); ok {
		_spec.SetField(serverconfig.FieldKick, field.TypeBool, value)
	}
	if value, ok := scu.mutation.Ban(); ok {
		_spec.SetField(serverconfig.FieldBan, field.TypeBool, value)
	}
	if value, ok := scu.mutation.CheckInvites(); ok {
		_spec.SetField(serverconfig.FieldCheckInvites, field.TypeBool, value)
	}
	if value, ok := scu.mutation.CheckLinks(); ok {
		_spec.SetField(serverconfig.FieldCheckLinks, field.TypeBool, value)
	}
	if value, ok := scu.mutation.Ratelimit(); ok {
		_spec.SetField(serverconfig.FieldRatelimit, field.TypeBool, value)
	}
	if value, ok := scu.mutation.Alerts(); ok {
		_spec.SetField(serverconfig.FieldAlerts, field.TypeBool, value)
	}
	if value, ok := scu.mutation.FlagLinks(); ok {
		_spec.SetField(serverconfig.FieldFlagLinks, field.TypeBool, value)
	}
	if value, ok := scu.mutation.LogChannel(); ok {
		_spec.SetField(serverconfig.FieldLogChannel, field.TypeString, value)
	}
	if value, ok := scu.mutation.ExcludedChannels(); ok {
		_spec.SetField(serverconfig.FieldExcludedChannels, field.TypeJSON, value)
	}
	if value, ok := scu.mutation.AppendedExcludedChannels(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, serverconfig.FieldExcludedChannels, value)
		})
	}
	if value, ok := scu.mutation.ExcludedRoles(); ok {
		_spec.SetField(serverconfig.FieldExcludedRoles, field.TypeJSON, value)
	}
	if value, ok := scu.mutation.AppendedExcludedRoles(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, serverconfig.FieldExcludedRoles, value)
		})
	}
	if value, ok := scu.mutation.ExcludedUsers(); ok {
		_spec.SetField(serverconfig.FieldExcludedUsers, field.TypeJSON, value)
	}
	if value, ok := scu.mutation.AppendedExcludedUsers(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, serverconfig.FieldExcludedUsers, value)
		})
	}
	if value, ok := scu.mutation.GivenRole(); ok {
		_spec.SetField(serverconfig.FieldGivenRole, field.TypeString, value)
	}
	if value, ok := scu.mutation.RatelimitMessage(); ok {
		_spec.SetField(serverconfig.FieldRatelimitMessage, field.TypeInt, value)
	}
	if value, ok := scu.mutation.AddedRatelimitMessage(); ok {
		_spec.AddField(serverconfig.FieldRatelimitMessage, field.TypeInt, value)
	}
	if value, ok := scu.mutation.RatelimitTime(); ok {
		_spec.SetField(serverconfig.FieldRatelimitTime, field.TypeEnum, value)
	}
	if value, ok := scu.mutation.TimeoutTime(); ok {
		_spec.SetField(serverconfig.FieldTimeoutTime, field.TypeEnum, value)
	}
	if value, ok := scu.mutation.BanDeleteMessageTime(); ok {
		_spec.SetField(serverconfig.FieldBanDeleteMessageTime, field.TypeEnum, value)
	}
	if scu.mutation.ServerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   serverconfig.ServerTable,
			Columns: []string{serverconfig.ServerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(server.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scu.mutation.RemovedServerIDs(); len(nodes) > 0 && !scu.mutation.ServerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   serverconfig.ServerTable,
			Columns: []string{serverconfig.ServerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(server.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scu.mutation.ServerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   serverconfig.ServerTable,
			Columns: []string{serverconfig.ServerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(server.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(scu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, scu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{serverconfig.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	scu.mutation.done = true
	return n, nil
}

// ServerConfigUpdateOne is the builder for updating a single ServerConfig entity.
type ServerConfigUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ServerConfigMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdateTime sets the "update_time" field.
func (scuo *ServerConfigUpdateOne) SetUpdateTime(t time.Time) *ServerConfigUpdateOne {
	scuo.mutation.SetUpdateTime(t)
	return scuo
}

// SetRemoveRoles sets the "remove_roles" field.
func (scuo *ServerConfigUpdateOne) SetRemoveRoles(b bool) *ServerConfigUpdateOne {
	scuo.mutation.SetRemoveRoles(b)
	return scuo
}

// SetNillableRemoveRoles sets the "remove_roles" field if the given value is not nil.
func (scuo *ServerConfigUpdateOne) SetNillableRemoveRoles(b *bool) *ServerConfigUpdateOne {
	if b != nil {
		scuo.SetRemoveRoles(*b)
	}
	return scuo
}

// SetGiveRole sets the "give_role" field.
func (scuo *ServerConfigUpdateOne) SetGiveRole(b bool) *ServerConfigUpdateOne {
	scuo.mutation.SetGiveRole(b)
	return scuo
}

// SetNillableGiveRole sets the "give_role" field if the given value is not nil.
func (scuo *ServerConfigUpdateOne) SetNillableGiveRole(b *bool) *ServerConfigUpdateOne {
	if b != nil {
		scuo.SetGiveRole(*b)
	}
	return scuo
}

// SetTimeout sets the "timeout" field.
func (scuo *ServerConfigUpdateOne) SetTimeout(b bool) *ServerConfigUpdateOne {
	scuo.mutation.SetTimeout(b)
	return scuo
}

// SetNillableTimeout sets the "timeout" field if the given value is not nil.
func (scuo *ServerConfigUpdateOne) SetNillableTimeout(b *bool) *ServerConfigUpdateOne {
	if b != nil {
		scuo.SetTimeout(*b)
	}
	return scuo
}

// SetKick sets the "kick" field.
func (scuo *ServerConfigUpdateOne) SetKick(b bool) *ServerConfigUpdateOne {
	scuo.mutation.SetKick(b)
	return scuo
}

// SetNillableKick sets the "kick" field if the given value is not nil.
func (scuo *ServerConfigUpdateOne) SetNillableKick(b *bool) *ServerConfigUpdateOne {
	if b != nil {
		scuo.SetKick(*b)
	}
	return scuo
}

// SetBan sets the "ban" field.
func (scuo *ServerConfigUpdateOne) SetBan(b bool) *ServerConfigUpdateOne {
	scuo.mutation.SetBan(b)
	return scuo
}

// SetNillableBan sets the "ban" field if the given value is not nil.
func (scuo *ServerConfigUpdateOne) SetNillableBan(b *bool) *ServerConfigUpdateOne {
	if b != nil {
		scuo.SetBan(*b)
	}
	return scuo
}

// SetCheckInvites sets the "check_invites" field.
func (scuo *ServerConfigUpdateOne) SetCheckInvites(b bool) *ServerConfigUpdateOne {
	scuo.mutation.SetCheckInvites(b)
	return scuo
}

// SetNillableCheckInvites sets the "check_invites" field if the given value is not nil.
func (scuo *ServerConfigUpdateOne) SetNillableCheckInvites(b *bool) *ServerConfigUpdateOne {
	if b != nil {
		scuo.SetCheckInvites(*b)
	}
	return scuo
}

// SetCheckLinks sets the "check_links" field.
func (scuo *ServerConfigUpdateOne) SetCheckLinks(b bool) *ServerConfigUpdateOne {
	scuo.mutation.SetCheckLinks(b)
	return scuo
}

// SetNillableCheckLinks sets the "check_links" field if the given value is not nil.
func (scuo *ServerConfigUpdateOne) SetNillableCheckLinks(b *bool) *ServerConfigUpdateOne {
	if b != nil {
		scuo.SetCheckLinks(*b)
	}
	return scuo
}

// SetRatelimit sets the "ratelimit" field.
func (scuo *ServerConfigUpdateOne) SetRatelimit(b bool) *ServerConfigUpdateOne {
	scuo.mutation.SetRatelimit(b)
	return scuo
}

// SetNillableRatelimit sets the "ratelimit" field if the given value is not nil.
func (scuo *ServerConfigUpdateOne) SetNillableRatelimit(b *bool) *ServerConfigUpdateOne {
	if b != nil {
		scuo.SetRatelimit(*b)
	}
	return scuo
}

// SetAlerts sets the "alerts" field.
func (scuo *ServerConfigUpdateOne) SetAlerts(b bool) *ServerConfigUpdateOne {
	scuo.mutation.SetAlerts(b)
	return scuo
}

// SetNillableAlerts sets the "alerts" field if the given value is not nil.
func (scuo *ServerConfigUpdateOne) SetNillableAlerts(b *bool) *ServerConfigUpdateOne {
	if b != nil {
		scuo.SetAlerts(*b)
	}
	return scuo
}

// SetFlagLinks sets the "flag_links" field.
func (scuo *ServerConfigUpdateOne) SetFlagLinks(b bool) *ServerConfigUpdateOne {
	scuo.mutation.SetFlagLinks(b)
	return scuo
}

// SetNillableFlagLinks sets the "flag_links" field if the given value is not nil.
func (scuo *ServerConfigUpdateOne) SetNillableFlagLinks(b *bool) *ServerConfigUpdateOne {
	if b != nil {
		scuo.SetFlagLinks(*b)
	}
	return scuo
}

// SetLogChannel sets the "log_channel" field.
func (scuo *ServerConfigUpdateOne) SetLogChannel(s string) *ServerConfigUpdateOne {
	scuo.mutation.SetLogChannel(s)
	return scuo
}

// SetNillableLogChannel sets the "log_channel" field if the given value is not nil.
func (scuo *ServerConfigUpdateOne) SetNillableLogChannel(s *string) *ServerConfigUpdateOne {
	if s != nil {
		scuo.SetLogChannel(*s)
	}
	return scuo
}

// SetExcludedChannels sets the "excluded_channels" field.
func (scuo *ServerConfigUpdateOne) SetExcludedChannels(s []string) *ServerConfigUpdateOne {
	scuo.mutation.SetExcludedChannels(s)
	return scuo
}

// AppendExcludedChannels appends s to the "excluded_channels" field.
func (scuo *ServerConfigUpdateOne) AppendExcludedChannels(s []string) *ServerConfigUpdateOne {
	scuo.mutation.AppendExcludedChannels(s)
	return scuo
}

// SetExcludedRoles sets the "excluded_roles" field.
func (scuo *ServerConfigUpdateOne) SetExcludedRoles(s []string) *ServerConfigUpdateOne {
	scuo.mutation.SetExcludedRoles(s)
	return scuo
}

// AppendExcludedRoles appends s to the "excluded_roles" field.
func (scuo *ServerConfigUpdateOne) AppendExcludedRoles(s []string) *ServerConfigUpdateOne {
	scuo.mutation.AppendExcludedRoles(s)
	return scuo
}

// SetExcludedUsers sets the "excluded_users" field.
func (scuo *ServerConfigUpdateOne) SetExcludedUsers(s []string) *ServerConfigUpdateOne {
	scuo.mutation.SetExcludedUsers(s)
	return scuo
}

// AppendExcludedUsers appends s to the "excluded_users" field.
func (scuo *ServerConfigUpdateOne) AppendExcludedUsers(s []string) *ServerConfigUpdateOne {
	scuo.mutation.AppendExcludedUsers(s)
	return scuo
}

// SetGivenRole sets the "given_role" field.
func (scuo *ServerConfigUpdateOne) SetGivenRole(s string) *ServerConfigUpdateOne {
	scuo.mutation.SetGivenRole(s)
	return scuo
}

// SetNillableGivenRole sets the "given_role" field if the given value is not nil.
func (scuo *ServerConfigUpdateOne) SetNillableGivenRole(s *string) *ServerConfigUpdateOne {
	if s != nil {
		scuo.SetGivenRole(*s)
	}
	return scuo
}

// SetRatelimitMessage sets the "ratelimit_message" field.
func (scuo *ServerConfigUpdateOne) SetRatelimitMessage(i int) *ServerConfigUpdateOne {
	scuo.mutation.ResetRatelimitMessage()
	scuo.mutation.SetRatelimitMessage(i)
	return scuo
}

// SetNillableRatelimitMessage sets the "ratelimit_message" field if the given value is not nil.
func (scuo *ServerConfigUpdateOne) SetNillableRatelimitMessage(i *int) *ServerConfigUpdateOne {
	if i != nil {
		scuo.SetRatelimitMessage(*i)
	}
	return scuo
}

// AddRatelimitMessage adds i to the "ratelimit_message" field.
func (scuo *ServerConfigUpdateOne) AddRatelimitMessage(i int) *ServerConfigUpdateOne {
	scuo.mutation.AddRatelimitMessage(i)
	return scuo
}

// SetRatelimitTime sets the "ratelimit_time" field.
func (scuo *ServerConfigUpdateOne) SetRatelimitTime(st serverconfig.RatelimitTime) *ServerConfigUpdateOne {
	scuo.mutation.SetRatelimitTime(st)
	return scuo
}

// SetNillableRatelimitTime sets the "ratelimit_time" field if the given value is not nil.
func (scuo *ServerConfigUpdateOne) SetNillableRatelimitTime(st *serverconfig.RatelimitTime) *ServerConfigUpdateOne {
	if st != nil {
		scuo.SetRatelimitTime(*st)
	}
	return scuo
}

// SetTimeoutTime sets the "timeout_time" field.
func (scuo *ServerConfigUpdateOne) SetTimeoutTime(st serverconfig.TimeoutTime) *ServerConfigUpdateOne {
	scuo.mutation.SetTimeoutTime(st)
	return scuo
}

// SetNillableTimeoutTime sets the "timeout_time" field if the given value is not nil.
func (scuo *ServerConfigUpdateOne) SetNillableTimeoutTime(st *serverconfig.TimeoutTime) *ServerConfigUpdateOne {
	if st != nil {
		scuo.SetTimeoutTime(*st)
	}
	return scuo
}

// SetBanDeleteMessageTime sets the "ban_delete_message_time" field.
func (scuo *ServerConfigUpdateOne) SetBanDeleteMessageTime(sdmt serverconfig.BanDeleteMessageTime) *ServerConfigUpdateOne {
	scuo.mutation.SetBanDeleteMessageTime(sdmt)
	return scuo
}

// SetNillableBanDeleteMessageTime sets the "ban_delete_message_time" field if the given value is not nil.
func (scuo *ServerConfigUpdateOne) SetNillableBanDeleteMessageTime(sdmt *serverconfig.BanDeleteMessageTime) *ServerConfigUpdateOne {
	if sdmt != nil {
		scuo.SetBanDeleteMessageTime(*sdmt)
	}
	return scuo
}

// AddServerIDs adds the "server" edge to the Server entity by IDs.
func (scuo *ServerConfigUpdateOne) AddServerIDs(ids ...string) *ServerConfigUpdateOne {
	scuo.mutation.AddServerIDs(ids...)
	return scuo
}

// AddServer adds the "server" edges to the Server entity.
func (scuo *ServerConfigUpdateOne) AddServer(s ...*Server) *ServerConfigUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return scuo.AddServerIDs(ids...)
}

// Mutation returns the ServerConfigMutation object of the builder.
func (scuo *ServerConfigUpdateOne) Mutation() *ServerConfigMutation {
	return scuo.mutation
}

// ClearServer clears all "server" edges to the Server entity.
func (scuo *ServerConfigUpdateOne) ClearServer() *ServerConfigUpdateOne {
	scuo.mutation.ClearServer()
	return scuo
}

// RemoveServerIDs removes the "server" edge to Server entities by IDs.
func (scuo *ServerConfigUpdateOne) RemoveServerIDs(ids ...string) *ServerConfigUpdateOne {
	scuo.mutation.RemoveServerIDs(ids...)
	return scuo
}

// RemoveServer removes "server" edges to Server entities.
func (scuo *ServerConfigUpdateOne) RemoveServer(s ...*Server) *ServerConfigUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return scuo.RemoveServerIDs(ids...)
}

// Where appends a list predicates to the ServerConfigUpdate builder.
func (scuo *ServerConfigUpdateOne) Where(ps ...predicate.ServerConfig) *ServerConfigUpdateOne {
	scuo.mutation.Where(ps...)
	return scuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (scuo *ServerConfigUpdateOne) Select(field string, fields ...string) *ServerConfigUpdateOne {
	scuo.fields = append([]string{field}, fields...)
	return scuo
}

// Save executes the query and returns the updated ServerConfig entity.
func (scuo *ServerConfigUpdateOne) Save(ctx context.Context) (*ServerConfig, error) {
	scuo.defaults()
	return withHooks(ctx, scuo.sqlSave, scuo.mutation, scuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (scuo *ServerConfigUpdateOne) SaveX(ctx context.Context) *ServerConfig {
	node, err := scuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (scuo *ServerConfigUpdateOne) Exec(ctx context.Context) error {
	_, err := scuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scuo *ServerConfigUpdateOne) ExecX(ctx context.Context) {
	if err := scuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (scuo *ServerConfigUpdateOne) defaults() {
	if _, ok := scuo.mutation.UpdateTime(); !ok {
		v := serverconfig.UpdateDefaultUpdateTime()
		scuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (scuo *ServerConfigUpdateOne) check() error {
	if v, ok := scuo.mutation.RatelimitTime(); ok {
		if err := serverconfig.RatelimitTimeValidator(v); err != nil {
			return &ValidationError{Name: "ratelimit_time", err: fmt.Errorf(`ent: validator failed for field "ServerConfig.ratelimit_time": %w`, err)}
		}
	}
	if v, ok := scuo.mutation.TimeoutTime(); ok {
		if err := serverconfig.TimeoutTimeValidator(v); err != nil {
			return &ValidationError{Name: "timeout_time", err: fmt.Errorf(`ent: validator failed for field "ServerConfig.timeout_time": %w`, err)}
		}
	}
	if v, ok := scuo.mutation.BanDeleteMessageTime(); ok {
		if err := serverconfig.BanDeleteMessageTimeValidator(v); err != nil {
			return &ValidationError{Name: "ban_delete_message_time", err: fmt.Errorf(`ent: validator failed for field "ServerConfig.ban_delete_message_time": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (scuo *ServerConfigUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ServerConfigUpdateOne {
	scuo.modifiers = append(scuo.modifiers, modifiers...)
	return scuo
}

func (scuo *ServerConfigUpdateOne) sqlSave(ctx context.Context) (_node *ServerConfig, err error) {
	if err := scuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(serverconfig.Table, serverconfig.Columns, sqlgraph.NewFieldSpec(serverconfig.FieldID, field.TypeInt))
	id, ok := scuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ServerConfig.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := scuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, serverconfig.FieldID)
		for _, f := range fields {
			if !serverconfig.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != serverconfig.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := scuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := scuo.mutation.UpdateTime(); ok {
		_spec.SetField(serverconfig.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := scuo.mutation.RemoveRoles(); ok {
		_spec.SetField(serverconfig.FieldRemoveRoles, field.TypeBool, value)
	}
	if value, ok := scuo.mutation.GiveRole(); ok {
		_spec.SetField(serverconfig.FieldGiveRole, field.TypeBool, value)
	}
	if value, ok := scuo.mutation.Timeout(); ok {
		_spec.SetField(serverconfig.FieldTimeout, field.TypeBool, value)
	}
	if value, ok := scuo.mutation.Kick(); ok {
		_spec.SetField(serverconfig.FieldKick, field.TypeBool, value)
	}
	if value, ok := scuo.mutation.Ban(); ok {
		_spec.SetField(serverconfig.FieldBan, field.TypeBool, value)
	}
	if value, ok := scuo.mutation.CheckInvites(); ok {
		_spec.SetField(serverconfig.FieldCheckInvites, field.TypeBool, value)
	}
	if value, ok := scuo.mutation.CheckLinks(); ok {
		_spec.SetField(serverconfig.FieldCheckLinks, field.TypeBool, value)
	}
	if value, ok := scuo.mutation.Ratelimit(); ok {
		_spec.SetField(serverconfig.FieldRatelimit, field.TypeBool, value)
	}
	if value, ok := scuo.mutation.Alerts(); ok {
		_spec.SetField(serverconfig.FieldAlerts, field.TypeBool, value)
	}
	if value, ok := scuo.mutation.FlagLinks(); ok {
		_spec.SetField(serverconfig.FieldFlagLinks, field.TypeBool, value)
	}
	if value, ok := scuo.mutation.LogChannel(); ok {
		_spec.SetField(serverconfig.FieldLogChannel, field.TypeString, value)
	}
	if value, ok := scuo.mutation.ExcludedChannels(); ok {
		_spec.SetField(serverconfig.FieldExcludedChannels, field.TypeJSON, value)
	}
	if value, ok := scuo.mutation.AppendedExcludedChannels(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, serverconfig.FieldExcludedChannels, value)
		})
	}
	if value, ok := scuo.mutation.ExcludedRoles(); ok {
		_spec.SetField(serverconfig.FieldExcludedRoles, field.TypeJSON, value)
	}
	if value, ok := scuo.mutation.AppendedExcludedRoles(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, serverconfig.FieldExcludedRoles, value)
		})
	}
	if value, ok := scuo.mutation.ExcludedUsers(); ok {
		_spec.SetField(serverconfig.FieldExcludedUsers, field.TypeJSON, value)
	}
	if value, ok := scuo.mutation.AppendedExcludedUsers(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, serverconfig.FieldExcludedUsers, value)
		})
	}
	if value, ok := scuo.mutation.GivenRole(); ok {
		_spec.SetField(serverconfig.FieldGivenRole, field.TypeString, value)
	}
	if value, ok := scuo.mutation.RatelimitMessage(); ok {
		_spec.SetField(serverconfig.FieldRatelimitMessage, field.TypeInt, value)
	}
	if value, ok := scuo.mutation.AddedRatelimitMessage(); ok {
		_spec.AddField(serverconfig.FieldRatelimitMessage, field.TypeInt, value)
	}
	if value, ok := scuo.mutation.RatelimitTime(); ok {
		_spec.SetField(serverconfig.FieldRatelimitTime, field.TypeEnum, value)
	}
	if value, ok := scuo.mutation.TimeoutTime(); ok {
		_spec.SetField(serverconfig.FieldTimeoutTime, field.TypeEnum, value)
	}
	if value, ok := scuo.mutation.BanDeleteMessageTime(); ok {
		_spec.SetField(serverconfig.FieldBanDeleteMessageTime, field.TypeEnum, value)
	}
	if scuo.mutation.ServerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   serverconfig.ServerTable,
			Columns: []string{serverconfig.ServerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(server.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scuo.mutation.RemovedServerIDs(); len(nodes) > 0 && !scuo.mutation.ServerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   serverconfig.ServerTable,
			Columns: []string{serverconfig.ServerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(server.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scuo.mutation.ServerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   serverconfig.ServerTable,
			Columns: []string{serverconfig.ServerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(server.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(scuo.modifiers...)
	_node = &ServerConfig{config: scuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, scuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{serverconfig.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	scuo.mutation.done = true
	return _node, nil
}