// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/FM1337/ASB/internal/ent/server"
	"github.com/FM1337/ASB/internal/ent/serverconfig"
	"github.com/FM1337/ASB/internal/ent/wordblacklist"
)

// ServerCreate is the builder for creating a Server entity.
type ServerCreate struct {
	config
	mutation *ServerMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (sc *ServerCreate) SetCreateTime(t time.Time) *ServerCreate {
	sc.mutation.SetCreateTime(t)
	return sc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (sc *ServerCreate) SetNillableCreateTime(t *time.Time) *ServerCreate {
	if t != nil {
		sc.SetCreateTime(*t)
	}
	return sc
}

// SetUpdateTime sets the "update_time" field.
func (sc *ServerCreate) SetUpdateTime(t time.Time) *ServerCreate {
	sc.mutation.SetUpdateTime(t)
	return sc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (sc *ServerCreate) SetNillableUpdateTime(t *time.Time) *ServerCreate {
	if t != nil {
		sc.SetUpdateTime(*t)
	}
	return sc
}

// SetOwnerID sets the "owner_id" field.
func (sc *ServerCreate) SetOwnerID(s string) *ServerCreate {
	sc.mutation.SetOwnerID(s)
	return sc
}

// SetEnabled sets the "enabled" field.
func (sc *ServerCreate) SetEnabled(b bool) *ServerCreate {
	sc.mutation.SetEnabled(b)
	return sc
}

// SetNillableEnabled sets the "enabled" field if the given value is not nil.
func (sc *ServerCreate) SetNillableEnabled(b *bool) *ServerCreate {
	if b != nil {
		sc.SetEnabled(*b)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *ServerCreate) SetID(s string) *ServerCreate {
	sc.mutation.SetID(s)
	return sc
}

// SetConfigurationID sets the "configuration" edge to the ServerConfig entity by ID.
func (sc *ServerCreate) SetConfigurationID(id int) *ServerCreate {
	sc.mutation.SetConfigurationID(id)
	return sc
}

// SetConfiguration sets the "configuration" edge to the ServerConfig entity.
func (sc *ServerCreate) SetConfiguration(s *ServerConfig) *ServerCreate {
	return sc.SetConfigurationID(s.ID)
}

// AddWordBlacklistIDs adds the "word_blacklist" edge to the WordBlacklist entity by IDs.
func (sc *ServerCreate) AddWordBlacklistIDs(ids ...int) *ServerCreate {
	sc.mutation.AddWordBlacklistIDs(ids...)
	return sc
}

// AddWordBlacklist adds the "word_blacklist" edges to the WordBlacklist entity.
func (sc *ServerCreate) AddWordBlacklist(w ...*WordBlacklist) *ServerCreate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return sc.AddWordBlacklistIDs(ids...)
}

// Mutation returns the ServerMutation object of the builder.
func (sc *ServerCreate) Mutation() *ServerMutation {
	return sc.mutation
}

// Save creates the Server in the database.
func (sc *ServerCreate) Save(ctx context.Context) (*Server, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ServerCreate) SaveX(ctx context.Context) *Server {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *ServerCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *ServerCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *ServerCreate) defaults() {
	if _, ok := sc.mutation.CreateTime(); !ok {
		v := server.DefaultCreateTime()
		sc.mutation.SetCreateTime(v)
	}
	if _, ok := sc.mutation.UpdateTime(); !ok {
		v := server.DefaultUpdateTime()
		sc.mutation.SetUpdateTime(v)
	}
	if _, ok := sc.mutation.Enabled(); !ok {
		v := server.DefaultEnabled
		sc.mutation.SetEnabled(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *ServerCreate) check() error {
	if _, ok := sc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Server.create_time"`)}
	}
	if _, ok := sc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Server.update_time"`)}
	}
	if _, ok := sc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner_id", err: errors.New(`ent: missing required field "Server.owner_id"`)}
	}
	if _, ok := sc.mutation.Enabled(); !ok {
		return &ValidationError{Name: "enabled", err: errors.New(`ent: missing required field "Server.enabled"`)}
	}
	if _, ok := sc.mutation.ConfigurationID(); !ok {
		return &ValidationError{Name: "configuration", err: errors.New(`ent: missing required edge "Server.configuration"`)}
	}
	return nil
}

func (sc *ServerCreate) sqlSave(ctx context.Context) (*Server, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Server.ID type: %T", _spec.ID.Value)
		}
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *ServerCreate) createSpec() (*Server, *sqlgraph.CreateSpec) {
	var (
		_node = &Server{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(server.Table, sqlgraph.NewFieldSpec(server.FieldID, field.TypeString))
	)
	_spec.OnConflict = sc.conflict
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.CreateTime(); ok {
		_spec.SetField(server.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := sc.mutation.UpdateTime(); ok {
		_spec.SetField(server.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := sc.mutation.OwnerID(); ok {
		_spec.SetField(server.FieldOwnerID, field.TypeString, value)
		_node.OwnerID = value
	}
	if value, ok := sc.mutation.Enabled(); ok {
		_spec.SetField(server.FieldEnabled, field.TypeBool, value)
		_node.Enabled = value
	}
	if nodes := sc.mutation.ConfigurationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   server.ConfigurationTable,
			Columns: []string{server.ConfigurationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(serverconfig.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.server_configuration = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.WordBlacklistIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   server.WordBlacklistTable,
			Columns: server.WordBlacklistPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(wordblacklist.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Server.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ServerUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (sc *ServerCreate) OnConflict(opts ...sql.ConflictOption) *ServerUpsertOne {
	sc.conflict = opts
	return &ServerUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Server.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sc *ServerCreate) OnConflictColumns(columns ...string) *ServerUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &ServerUpsertOne{
		create: sc,
	}
}

type (
	// ServerUpsertOne is the builder for "upsert"-ing
	//  one Server node.
	ServerUpsertOne struct {
		create *ServerCreate
	}

	// ServerUpsert is the "OnConflict" setter.
	ServerUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *ServerUpsert) SetUpdateTime(v time.Time) *ServerUpsert {
	u.Set(server.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *ServerUpsert) UpdateUpdateTime() *ServerUpsert {
	u.SetExcluded(server.FieldUpdateTime)
	return u
}

// SetOwnerID sets the "owner_id" field.
func (u *ServerUpsert) SetOwnerID(v string) *ServerUpsert {
	u.Set(server.FieldOwnerID, v)
	return u
}

// UpdateOwnerID sets the "owner_id" field to the value that was provided on create.
func (u *ServerUpsert) UpdateOwnerID() *ServerUpsert {
	u.SetExcluded(server.FieldOwnerID)
	return u
}

// SetEnabled sets the "enabled" field.
func (u *ServerUpsert) SetEnabled(v bool) *ServerUpsert {
	u.Set(server.FieldEnabled, v)
	return u
}

// UpdateEnabled sets the "enabled" field to the value that was provided on create.
func (u *ServerUpsert) UpdateEnabled() *ServerUpsert {
	u.SetExcluded(server.FieldEnabled)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Server.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(server.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ServerUpsertOne) UpdateNewValues() *ServerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(server.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(server.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Server.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ServerUpsertOne) Ignore() *ServerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ServerUpsertOne) DoNothing() *ServerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ServerCreate.OnConflict
// documentation for more info.
func (u *ServerUpsertOne) Update(set func(*ServerUpsert)) *ServerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ServerUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *ServerUpsertOne) SetUpdateTime(v time.Time) *ServerUpsertOne {
	return u.Update(func(s *ServerUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *ServerUpsertOne) UpdateUpdateTime() *ServerUpsertOne {
	return u.Update(func(s *ServerUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetOwnerID sets the "owner_id" field.
func (u *ServerUpsertOne) SetOwnerID(v string) *ServerUpsertOne {
	return u.Update(func(s *ServerUpsert) {
		s.SetOwnerID(v)
	})
}

// UpdateOwnerID sets the "owner_id" field to the value that was provided on create.
func (u *ServerUpsertOne) UpdateOwnerID() *ServerUpsertOne {
	return u.Update(func(s *ServerUpsert) {
		s.UpdateOwnerID()
	})
}

// SetEnabled sets the "enabled" field.
func (u *ServerUpsertOne) SetEnabled(v bool) *ServerUpsertOne {
	return u.Update(func(s *ServerUpsert) {
		s.SetEnabled(v)
	})
}

// UpdateEnabled sets the "enabled" field to the value that was provided on create.
func (u *ServerUpsertOne) UpdateEnabled() *ServerUpsertOne {
	return u.Update(func(s *ServerUpsert) {
		s.UpdateEnabled()
	})
}

// Exec executes the query.
func (u *ServerUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ServerCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ServerUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ServerUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ServerUpsertOne.ID is not supported by MySQL driver. Use ServerUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ServerUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ServerCreateBulk is the builder for creating many Server entities in bulk.
type ServerCreateBulk struct {
	config
	err      error
	builders []*ServerCreate
	conflict []sql.ConflictOption
}

// Save creates the Server entities in the database.
func (scb *ServerCreateBulk) Save(ctx context.Context) ([]*Server, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Server, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ServerMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = scb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *ServerCreateBulk) SaveX(ctx context.Context) []*Server {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *ServerCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *ServerCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Server.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ServerUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (scb *ServerCreateBulk) OnConflict(opts ...sql.ConflictOption) *ServerUpsertBulk {
	scb.conflict = opts
	return &ServerUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Server.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (scb *ServerCreateBulk) OnConflictColumns(columns ...string) *ServerUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &ServerUpsertBulk{
		create: scb,
	}
}

// ServerUpsertBulk is the builder for "upsert"-ing
// a bulk of Server nodes.
type ServerUpsertBulk struct {
	create *ServerCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Server.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(server.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ServerUpsertBulk) UpdateNewValues() *ServerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(server.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(server.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Server.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ServerUpsertBulk) Ignore() *ServerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ServerUpsertBulk) DoNothing() *ServerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ServerCreateBulk.OnConflict
// documentation for more info.
func (u *ServerUpsertBulk) Update(set func(*ServerUpsert)) *ServerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ServerUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *ServerUpsertBulk) SetUpdateTime(v time.Time) *ServerUpsertBulk {
	return u.Update(func(s *ServerUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *ServerUpsertBulk) UpdateUpdateTime() *ServerUpsertBulk {
	return u.Update(func(s *ServerUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetOwnerID sets the "owner_id" field.
func (u *ServerUpsertBulk) SetOwnerID(v string) *ServerUpsertBulk {
	return u.Update(func(s *ServerUpsert) {
		s.SetOwnerID(v)
	})
}

// UpdateOwnerID sets the "owner_id" field to the value that was provided on create.
func (u *ServerUpsertBulk) UpdateOwnerID() *ServerUpsertBulk {
	return u.Update(func(s *ServerUpsert) {
		s.UpdateOwnerID()
	})
}

// SetEnabled sets the "enabled" field.
func (u *ServerUpsertBulk) SetEnabled(v bool) *ServerUpsertBulk {
	return u.Update(func(s *ServerUpsert) {
		s.SetEnabled(v)
	})
}

// UpdateEnabled sets the "enabled" field to the value that was provided on create.
func (u *ServerUpsertBulk) UpdateEnabled() *ServerUpsertBulk {
	return u.Update(func(s *ServerUpsert) {
		s.UpdateEnabled()
	})
}

// Exec executes the query.
func (u *ServerUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ServerCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ServerCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ServerUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
