// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/FM1337/ASB/internal/ent/wordblacklist"
)

// WordBlacklist is the model entity for the WordBlacklist schema.
type WordBlacklist struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Word holds the value of the "word" field.
	Word string `json:"word,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WordBlacklistQuery when eager-loading is set.
	Edges        WordBlacklistEdges `json:"edges"`
	selectValues sql.SelectValues
}

// WordBlacklistEdges holds the relations/edges for other nodes in the graph.
type WordBlacklistEdges struct {
	// Server holds the value of the server edge.
	Server []*Server `json:"server,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ServerOrErr returns the Server value or an error if the edge
// was not loaded in eager-loading.
func (e WordBlacklistEdges) ServerOrErr() ([]*Server, error) {
	if e.loadedTypes[0] {
		return e.Server, nil
	}
	return nil, &NotLoadedError{edge: "server"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WordBlacklist) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case wordblacklist.FieldID:
			values[i] = new(sql.NullInt64)
		case wordblacklist.FieldWord:
			values[i] = new(sql.NullString)
		case wordblacklist.FieldCreateTime, wordblacklist.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WordBlacklist fields.
func (wb *WordBlacklist) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case wordblacklist.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			wb.ID = int(value.Int64)
		case wordblacklist.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				wb.CreateTime = value.Time
			}
		case wordblacklist.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				wb.UpdateTime = value.Time
			}
		case wordblacklist.FieldWord:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field word", values[i])
			} else if value.Valid {
				wb.Word = value.String
			}
		default:
			wb.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the WordBlacklist.
// This includes values selected through modifiers, order, etc.
func (wb *WordBlacklist) Value(name string) (ent.Value, error) {
	return wb.selectValues.Get(name)
}

// QueryServer queries the "server" edge of the WordBlacklist entity.
func (wb *WordBlacklist) QueryServer() *ServerQuery {
	return NewWordBlacklistClient(wb.config).QueryServer(wb)
}

// Update returns a builder for updating this WordBlacklist.
// Note that you need to call WordBlacklist.Unwrap() before calling this method if this WordBlacklist
// was returned from a transaction, and the transaction was committed or rolled back.
func (wb *WordBlacklist) Update() *WordBlacklistUpdateOne {
	return NewWordBlacklistClient(wb.config).UpdateOne(wb)
}

// Unwrap unwraps the WordBlacklist entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wb *WordBlacklist) Unwrap() *WordBlacklist {
	_tx, ok := wb.config.driver.(*txDriver)
	if !ok {
		panic("ent: WordBlacklist is not a transactional entity")
	}
	wb.config.driver = _tx.drv
	return wb
}

// String implements the fmt.Stringer.
func (wb *WordBlacklist) String() string {
	var builder strings.Builder
	builder.WriteString("WordBlacklist(")
	builder.WriteString(fmt.Sprintf("id=%v, ", wb.ID))
	builder.WriteString("create_time=")
	builder.WriteString(wb.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(wb.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("word=")
	builder.WriteString(wb.Word)
	builder.WriteByte(')')
	return builder.String()
}

// WordBlacklists is a parsable slice of WordBlacklist.
type WordBlacklists []*WordBlacklist
