// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/FM1337/ASB/internal/ent/predicate"
	"github.com/FM1337/ASB/internal/ent/server"
	"github.com/FM1337/ASB/internal/ent/wordblacklist"
)

// WordBlacklistUpdate is the builder for updating WordBlacklist entities.
type WordBlacklistUpdate struct {
	config
	hooks     []Hook
	mutation  *WordBlacklistMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the WordBlacklistUpdate builder.
func (wbu *WordBlacklistUpdate) Where(ps ...predicate.WordBlacklist) *WordBlacklistUpdate {
	wbu.mutation.Where(ps...)
	return wbu
}

// SetUpdateTime sets the "update_time" field.
func (wbu *WordBlacklistUpdate) SetUpdateTime(t time.Time) *WordBlacklistUpdate {
	wbu.mutation.SetUpdateTime(t)
	return wbu
}

// SetWord sets the "word" field.
func (wbu *WordBlacklistUpdate) SetWord(s string) *WordBlacklistUpdate {
	wbu.mutation.SetWord(s)
	return wbu
}

// SetNillableWord sets the "word" field if the given value is not nil.
func (wbu *WordBlacklistUpdate) SetNillableWord(s *string) *WordBlacklistUpdate {
	if s != nil {
		wbu.SetWord(*s)
	}
	return wbu
}

// AddServerIDs adds the "server" edge to the Server entity by IDs.
func (wbu *WordBlacklistUpdate) AddServerIDs(ids ...int) *WordBlacklistUpdate {
	wbu.mutation.AddServerIDs(ids...)
	return wbu
}

// AddServer adds the "server" edges to the Server entity.
func (wbu *WordBlacklistUpdate) AddServer(s ...*Server) *WordBlacklistUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return wbu.AddServerIDs(ids...)
}

// Mutation returns the WordBlacklistMutation object of the builder.
func (wbu *WordBlacklistUpdate) Mutation() *WordBlacklistMutation {
	return wbu.mutation
}

// ClearServer clears all "server" edges to the Server entity.
func (wbu *WordBlacklistUpdate) ClearServer() *WordBlacklistUpdate {
	wbu.mutation.ClearServer()
	return wbu
}

// RemoveServerIDs removes the "server" edge to Server entities by IDs.
func (wbu *WordBlacklistUpdate) RemoveServerIDs(ids ...int) *WordBlacklistUpdate {
	wbu.mutation.RemoveServerIDs(ids...)
	return wbu
}

// RemoveServer removes "server" edges to Server entities.
func (wbu *WordBlacklistUpdate) RemoveServer(s ...*Server) *WordBlacklistUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return wbu.RemoveServerIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wbu *WordBlacklistUpdate) Save(ctx context.Context) (int, error) {
	wbu.defaults()
	return withHooks(ctx, wbu.sqlSave, wbu.mutation, wbu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wbu *WordBlacklistUpdate) SaveX(ctx context.Context) int {
	affected, err := wbu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wbu *WordBlacklistUpdate) Exec(ctx context.Context) error {
	_, err := wbu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wbu *WordBlacklistUpdate) ExecX(ctx context.Context) {
	if err := wbu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wbu *WordBlacklistUpdate) defaults() {
	if _, ok := wbu.mutation.UpdateTime(); !ok {
		v := wordblacklist.UpdateDefaultUpdateTime()
		wbu.mutation.SetUpdateTime(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (wbu *WordBlacklistUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *WordBlacklistUpdate {
	wbu.modifiers = append(wbu.modifiers, modifiers...)
	return wbu
}

func (wbu *WordBlacklistUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(wordblacklist.Table, wordblacklist.Columns, sqlgraph.NewFieldSpec(wordblacklist.FieldID, field.TypeInt))
	if ps := wbu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wbu.mutation.UpdateTime(); ok {
		_spec.SetField(wordblacklist.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := wbu.mutation.Word(); ok {
		_spec.SetField(wordblacklist.FieldWord, field.TypeString, value)
	}
	if wbu.mutation.ServerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   wordblacklist.ServerTable,
			Columns: wordblacklist.ServerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(server.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wbu.mutation.RemovedServerIDs(); len(nodes) > 0 && !wbu.mutation.ServerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   wordblacklist.ServerTable,
			Columns: wordblacklist.ServerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(server.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wbu.mutation.ServerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   wordblacklist.ServerTable,
			Columns: wordblacklist.ServerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(server.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(wbu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, wbu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wordblacklist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wbu.mutation.done = true
	return n, nil
}

// WordBlacklistUpdateOne is the builder for updating a single WordBlacklist entity.
type WordBlacklistUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *WordBlacklistMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdateTime sets the "update_time" field.
func (wbuo *WordBlacklistUpdateOne) SetUpdateTime(t time.Time) *WordBlacklistUpdateOne {
	wbuo.mutation.SetUpdateTime(t)
	return wbuo
}

// SetWord sets the "word" field.
func (wbuo *WordBlacklistUpdateOne) SetWord(s string) *WordBlacklistUpdateOne {
	wbuo.mutation.SetWord(s)
	return wbuo
}

// SetNillableWord sets the "word" field if the given value is not nil.
func (wbuo *WordBlacklistUpdateOne) SetNillableWord(s *string) *WordBlacklistUpdateOne {
	if s != nil {
		wbuo.SetWord(*s)
	}
	return wbuo
}

// AddServerIDs adds the "server" edge to the Server entity by IDs.
func (wbuo *WordBlacklistUpdateOne) AddServerIDs(ids ...int) *WordBlacklistUpdateOne {
	wbuo.mutation.AddServerIDs(ids...)
	return wbuo
}

// AddServer adds the "server" edges to the Server entity.
func (wbuo *WordBlacklistUpdateOne) AddServer(s ...*Server) *WordBlacklistUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return wbuo.AddServerIDs(ids...)
}

// Mutation returns the WordBlacklistMutation object of the builder.
func (wbuo *WordBlacklistUpdateOne) Mutation() *WordBlacklistMutation {
	return wbuo.mutation
}

// ClearServer clears all "server" edges to the Server entity.
func (wbuo *WordBlacklistUpdateOne) ClearServer() *WordBlacklistUpdateOne {
	wbuo.mutation.ClearServer()
	return wbuo
}

// RemoveServerIDs removes the "server" edge to Server entities by IDs.
func (wbuo *WordBlacklistUpdateOne) RemoveServerIDs(ids ...int) *WordBlacklistUpdateOne {
	wbuo.mutation.RemoveServerIDs(ids...)
	return wbuo
}

// RemoveServer removes "server" edges to Server entities.
func (wbuo *WordBlacklistUpdateOne) RemoveServer(s ...*Server) *WordBlacklistUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return wbuo.RemoveServerIDs(ids...)
}

// Where appends a list predicates to the WordBlacklistUpdate builder.
func (wbuo *WordBlacklistUpdateOne) Where(ps ...predicate.WordBlacklist) *WordBlacklistUpdateOne {
	wbuo.mutation.Where(ps...)
	return wbuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wbuo *WordBlacklistUpdateOne) Select(field string, fields ...string) *WordBlacklistUpdateOne {
	wbuo.fields = append([]string{field}, fields...)
	return wbuo
}

// Save executes the query and returns the updated WordBlacklist entity.
func (wbuo *WordBlacklistUpdateOne) Save(ctx context.Context) (*WordBlacklist, error) {
	wbuo.defaults()
	return withHooks(ctx, wbuo.sqlSave, wbuo.mutation, wbuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wbuo *WordBlacklistUpdateOne) SaveX(ctx context.Context) *WordBlacklist {
	node, err := wbuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wbuo *WordBlacklistUpdateOne) Exec(ctx context.Context) error {
	_, err := wbuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wbuo *WordBlacklistUpdateOne) ExecX(ctx context.Context) {
	if err := wbuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wbuo *WordBlacklistUpdateOne) defaults() {
	if _, ok := wbuo.mutation.UpdateTime(); !ok {
		v := wordblacklist.UpdateDefaultUpdateTime()
		wbuo.mutation.SetUpdateTime(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (wbuo *WordBlacklistUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *WordBlacklistUpdateOne {
	wbuo.modifiers = append(wbuo.modifiers, modifiers...)
	return wbuo
}

func (wbuo *WordBlacklistUpdateOne) sqlSave(ctx context.Context) (_node *WordBlacklist, err error) {
	_spec := sqlgraph.NewUpdateSpec(wordblacklist.Table, wordblacklist.Columns, sqlgraph.NewFieldSpec(wordblacklist.FieldID, field.TypeInt))
	id, ok := wbuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "WordBlacklist.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wbuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, wordblacklist.FieldID)
		for _, f := range fields {
			if !wordblacklist.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != wordblacklist.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wbuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wbuo.mutation.UpdateTime(); ok {
		_spec.SetField(wordblacklist.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := wbuo.mutation.Word(); ok {
		_spec.SetField(wordblacklist.FieldWord, field.TypeString, value)
	}
	if wbuo.mutation.ServerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   wordblacklist.ServerTable,
			Columns: wordblacklist.ServerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(server.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wbuo.mutation.RemovedServerIDs(); len(nodes) > 0 && !wbuo.mutation.ServerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   wordblacklist.ServerTable,
			Columns: wordblacklist.ServerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(server.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wbuo.mutation.ServerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   wordblacklist.ServerTable,
			Columns: wordblacklist.ServerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(server.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(wbuo.modifiers...)
	_node = &WordBlacklist{config: wbuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wbuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wordblacklist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wbuo.mutation.done = true
	return _node, nil
}
