// Code generated by ent, DO NOT EDIT.

package cooldown

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/FM1337/ASB/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldEQ(FieldUpdateTime, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldEQ(FieldUserID, v))
}

// Hash applies equality check predicate on the "hash" field. It's identical to HashEQ.
func Hash(v string) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldEQ(FieldHash, v))
}

// Count applies equality check predicate on the "count" field. It's identical to CountEQ.
func Count(v int) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldEQ(FieldCount, v))
}

// ResetsAt applies equality check predicate on the "resets_at" field. It's identical to ResetsAtEQ.
func ResetsAt(v time.Time) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldEQ(FieldResetsAt, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldLTE(FieldUpdateTime, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldContainsFold(FieldUserID, v))
}

// HashEQ applies the EQ predicate on the "hash" field.
func HashEQ(v string) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldEQ(FieldHash, v))
}

// HashNEQ applies the NEQ predicate on the "hash" field.
func HashNEQ(v string) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldNEQ(FieldHash, v))
}

// HashIn applies the In predicate on the "hash" field.
func HashIn(vs ...string) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldIn(FieldHash, vs...))
}

// HashNotIn applies the NotIn predicate on the "hash" field.
func HashNotIn(vs ...string) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldNotIn(FieldHash, vs...))
}

// HashGT applies the GT predicate on the "hash" field.
func HashGT(v string) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldGT(FieldHash, v))
}

// HashGTE applies the GTE predicate on the "hash" field.
func HashGTE(v string) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldGTE(FieldHash, v))
}

// HashLT applies the LT predicate on the "hash" field.
func HashLT(v string) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldLT(FieldHash, v))
}

// HashLTE applies the LTE predicate on the "hash" field.
func HashLTE(v string) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldLTE(FieldHash, v))
}

// HashContains applies the Contains predicate on the "hash" field.
func HashContains(v string) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldContains(FieldHash, v))
}

// HashHasPrefix applies the HasPrefix predicate on the "hash" field.
func HashHasPrefix(v string) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldHasPrefix(FieldHash, v))
}

// HashHasSuffix applies the HasSuffix predicate on the "hash" field.
func HashHasSuffix(v string) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldHasSuffix(FieldHash, v))
}

// HashEqualFold applies the EqualFold predicate on the "hash" field.
func HashEqualFold(v string) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldEqualFold(FieldHash, v))
}

// HashContainsFold applies the ContainsFold predicate on the "hash" field.
func HashContainsFold(v string) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldContainsFold(FieldHash, v))
}

// CountEQ applies the EQ predicate on the "count" field.
func CountEQ(v int) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldEQ(FieldCount, v))
}

// CountNEQ applies the NEQ predicate on the "count" field.
func CountNEQ(v int) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldNEQ(FieldCount, v))
}

// CountIn applies the In predicate on the "count" field.
func CountIn(vs ...int) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldIn(FieldCount, vs...))
}

// CountNotIn applies the NotIn predicate on the "count" field.
func CountNotIn(vs ...int) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldNotIn(FieldCount, vs...))
}

// CountGT applies the GT predicate on the "count" field.
func CountGT(v int) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldGT(FieldCount, v))
}

// CountGTE applies the GTE predicate on the "count" field.
func CountGTE(v int) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldGTE(FieldCount, v))
}

// CountLT applies the LT predicate on the "count" field.
func CountLT(v int) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldLT(FieldCount, v))
}

// CountLTE applies the LTE predicate on the "count" field.
func CountLTE(v int) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldLTE(FieldCount, v))
}

// ResetsAtEQ applies the EQ predicate on the "resets_at" field.
func ResetsAtEQ(v time.Time) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldEQ(FieldResetsAt, v))
}

// ResetsAtNEQ applies the NEQ predicate on the "resets_at" field.
func ResetsAtNEQ(v time.Time) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldNEQ(FieldResetsAt, v))
}

// ResetsAtIn applies the In predicate on the "resets_at" field.
func ResetsAtIn(vs ...time.Time) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldIn(FieldResetsAt, vs...))
}

// ResetsAtNotIn applies the NotIn predicate on the "resets_at" field.
func ResetsAtNotIn(vs ...time.Time) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldNotIn(FieldResetsAt, vs...))
}

// ResetsAtGT applies the GT predicate on the "resets_at" field.
func ResetsAtGT(v time.Time) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldGT(FieldResetsAt, v))
}

// ResetsAtGTE applies the GTE predicate on the "resets_at" field.
func ResetsAtGTE(v time.Time) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldGTE(FieldResetsAt, v))
}

// ResetsAtLT applies the LT predicate on the "resets_at" field.
func ResetsAtLT(v time.Time) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldLT(FieldResetsAt, v))
}

// ResetsAtLTE applies the LTE predicate on the "resets_at" field.
func ResetsAtLTE(v time.Time) predicate.Cooldown {
	return predicate.Cooldown(sql.FieldLTE(FieldResetsAt, v))
}

// HasServer applies the HasEdge predicate on the "server" edge.
func HasServer() predicate.Cooldown {
	return predicate.Cooldown(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ServerTable, ServerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasServerWith applies the HasEdge predicate on the "server" edge with a given conditions (other predicates).
func HasServerWith(preds ...predicate.Server) predicate.Cooldown {
	return predicate.Cooldown(func(s *sql.Selector) {
		step := newServerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Cooldown) predicate.Cooldown {
	return predicate.Cooldown(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Cooldown) predicate.Cooldown {
	return predicate.Cooldown(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Cooldown) predicate.Cooldown {
	return predicate.Cooldown(sql.NotPredicates(p))
}