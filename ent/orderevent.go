// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Kotodian/cwmodel/ent/orderevent"
	"github.com/Kotodian/cwmodel/ent/orderinfo"
	"github.com/Kotodian/gokit/datasource"
)

// OrderEvent is the model entity for the OrderEvent schema.
type OrderEvent struct {
	config `json:"-"`
	// ID of the ent.
	// 主键
	ID datasource.UUID `json:"id,omitempty"`
	// 事件内容
	Content string `json:"content,omitempty"`
	// 事件发生时间
	Occurrence int64 `json:"occurrence,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderEventQuery when eager-loading is set.
	Edges    OrderEventEdges `json:"-"`
	order_id *datasource.UUID
}

// OrderEventEdges holds the relations/edges for other nodes in the graph.
type OrderEventEdges struct {
	// OrderInfo holds the value of the order_info edge.
	OrderInfo *OrderInfo `json:"order_info,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OrderInfoOrErr returns the OrderInfo value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderEventEdges) OrderInfoOrErr() (*OrderInfo, error) {
	if e.loadedTypes[0] {
		if e.OrderInfo == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: orderinfo.Label}
		}
		return e.OrderInfo, nil
	}
	return nil, &NotLoadedError{edge: "order_info"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrderEvent) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case orderevent.FieldID, orderevent.FieldOccurrence:
			values[i] = new(sql.NullInt64)
		case orderevent.FieldContent:
			values[i] = new(sql.NullString)
		case orderevent.ForeignKeys[0]: // order_id
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OrderEvent", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrderEvent fields.
func (oe *OrderEvent) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orderevent.FieldID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				oe.ID = datasource.UUID(value.Int64)
			}
		case orderevent.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				oe.Content = value.String
			}
		case orderevent.FieldOccurrence:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field occurrence", values[i])
			} else if value.Valid {
				oe.Occurrence = value.Int64
			}
		case orderevent.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value.Valid {
				oe.order_id = new(datasource.UUID)
				*oe.order_id = datasource.UUID(value.Int64)
			}
		}
	}
	return nil
}

// QueryOrderInfo queries the "order_info" edge of the OrderEvent entity.
func (oe *OrderEvent) QueryOrderInfo() *OrderInfoQuery {
	return (&OrderEventClient{config: oe.config}).QueryOrderInfo(oe)
}

// Update returns a builder for updating this OrderEvent.
// Note that you need to call OrderEvent.Unwrap() before calling this method if this OrderEvent
// was returned from a transaction, and the transaction was committed or rolled back.
func (oe *OrderEvent) Update() *OrderEventUpdateOne {
	return (&OrderEventClient{config: oe.config}).UpdateOne(oe)
}

// Unwrap unwraps the OrderEvent entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oe *OrderEvent) Unwrap() *OrderEvent {
	_tx, ok := oe.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrderEvent is not a transactional entity")
	}
	oe.config.driver = _tx.drv
	return oe
}

// String implements the fmt.Stringer.
func (oe *OrderEvent) String() string {
	var builder strings.Builder
	builder.WriteString("OrderEvent(")
	builder.WriteString(fmt.Sprintf("id=%v, ", oe.ID))
	builder.WriteString("content=")
	builder.WriteString(oe.Content)
	builder.WriteString(", ")
	builder.WriteString("occurrence=")
	builder.WriteString(fmt.Sprintf("%v", oe.Occurrence))
	builder.WriteByte(')')
	return builder.String()
}

// OrderEvents is a parsable slice of OrderEvent.
type OrderEvents []*OrderEvent

func (oe OrderEvents) config(cfg config) {
	for _i := range oe {
		oe[_i].config = cfg
	}
}
