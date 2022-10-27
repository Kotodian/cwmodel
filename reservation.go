// Code generated by ent, DO NOT EDIT.

package cwmodel

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Kotodian/cwmodel/connector"
	"github.com/Kotodian/cwmodel/equipment"
	"github.com/Kotodian/cwmodel/reservation"
	"github.com/Kotodian/gokit/datasource"
)

// Reservation is the model entity for the Reservation schema.
type Reservation struct {
	config `json:"-"`
	// ID of the ent.
	// 主键
	ID datasource.UUID `json:"id"`
	// 乐观锁
	Version int64 `json:"version"`
	// 创建者
	CreatedBy datasource.UUID `json:"-"`
	// 创建时间
	CreatedAt int64 `json:"-"`
	// 修改者
	UpdatedBy datasource.UUID `json:"-"`
	// 修改时间
	UpdatedAt int64 `json:"-"`
	// 桩id
	EquipmentID datasource.UUID `json:"equipment_id,omitempty"`
	// 枪id
	ConnectorID datasource.UUID `json:"connector_id,omitempty"`
	// 预约id
	ReservationID int64 `json:"reservation_id,omitempty"`
	// 授权模式
	AuthorizationMode int `json:"authorization_mode,omitempty"`
	// 授权id
	AuthorizationID string `json:"authorization_id,omitempty"`
	// 额外授权id
	Additional *string `json:"additional,omitempty"`
	// 客户id
	CustomerID *string `json:"customer_id,omitempty"`
	// 过期时间
	Expired int64 `json:"expired,omitempty"`
	// 预约状态
	State int `json:"state,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ReservationQuery when eager-loading is set.
	Edges ReservationEdges `json:"-"`
}

// ReservationEdges holds the relations/edges for other nodes in the graph.
type ReservationEdges struct {
	// Equipment holds the value of the equipment edge.
	Equipment *Equipment `json:"equipment,omitempty"`
	// Connector holds the value of the connector edge.
	Connector *Connector `json:"connector,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// EquipmentOrErr returns the Equipment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ReservationEdges) EquipmentOrErr() (*Equipment, error) {
	if e.loadedTypes[0] {
		if e.Equipment == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: equipment.Label}
		}
		return e.Equipment, nil
	}
	return nil, &NotLoadedError{edge: "equipment"}
}

// ConnectorOrErr returns the Connector value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ReservationEdges) ConnectorOrErr() (*Connector, error) {
	if e.loadedTypes[1] {
		if e.Connector == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: connector.Label}
		}
		return e.Connector, nil
	}
	return nil, &NotLoadedError{edge: "connector"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Reservation) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case reservation.FieldID, reservation.FieldVersion, reservation.FieldCreatedBy, reservation.FieldCreatedAt, reservation.FieldUpdatedBy, reservation.FieldUpdatedAt, reservation.FieldEquipmentID, reservation.FieldConnectorID, reservation.FieldReservationID, reservation.FieldAuthorizationMode, reservation.FieldExpired, reservation.FieldState:
			values[i] = new(sql.NullInt64)
		case reservation.FieldAuthorizationID, reservation.FieldAdditional, reservation.FieldCustomerID:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Reservation", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Reservation fields.
func (r *Reservation) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case reservation.FieldID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				r.ID = datasource.UUID(value.Int64)
			}
		case reservation.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				r.Version = value.Int64
			}
		case reservation.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				r.CreatedBy = datasource.UUID(value.Int64)
			}
		case reservation.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Int64
			}
		case reservation.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				r.UpdatedBy = datasource.UUID(value.Int64)
			}
		case reservation.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				r.UpdatedAt = value.Int64
			}
		case reservation.FieldEquipmentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field equipment_id", values[i])
			} else if value.Valid {
				r.EquipmentID = datasource.UUID(value.Int64)
			}
		case reservation.FieldConnectorID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field connector_id", values[i])
			} else if value.Valid {
				r.ConnectorID = datasource.UUID(value.Int64)
			}
		case reservation.FieldReservationID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field reservation_id", values[i])
			} else if value.Valid {
				r.ReservationID = value.Int64
			}
		case reservation.FieldAuthorizationMode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field authorization_mode", values[i])
			} else if value.Valid {
				r.AuthorizationMode = int(value.Int64)
			}
		case reservation.FieldAuthorizationID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field authorization_id", values[i])
			} else if value.Valid {
				r.AuthorizationID = value.String
			}
		case reservation.FieldAdditional:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field additional", values[i])
			} else if value.Valid {
				r.Additional = new(string)
				*r.Additional = value.String
			}
		case reservation.FieldCustomerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field customer_id", values[i])
			} else if value.Valid {
				r.CustomerID = new(string)
				*r.CustomerID = value.String
			}
		case reservation.FieldExpired:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field expired", values[i])
			} else if value.Valid {
				r.Expired = value.Int64
			}
		case reservation.FieldState:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				r.State = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryEquipment queries the "equipment" edge of the Reservation entity.
func (r *Reservation) QueryEquipment() *EquipmentQuery {
	return (&ReservationClient{config: r.config}).QueryEquipment(r)
}

// QueryConnector queries the "connector" edge of the Reservation entity.
func (r *Reservation) QueryConnector() *ConnectorQuery {
	return (&ReservationClient{config: r.config}).QueryConnector(r)
}

// Update returns a builder for updating this Reservation.
// Note that you need to call Reservation.Unwrap() before calling this method if this Reservation
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Reservation) Update() *ReservationUpdateOne {
	return (&ReservationClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Reservation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Reservation) Unwrap() *Reservation {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("cwmodel: Reservation is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Reservation) String() string {
	var builder strings.Builder
	builder.WriteString("Reservation(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", r.Version))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", r.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", r.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", r.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", r.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("equipment_id=")
	builder.WriteString(fmt.Sprintf("%v", r.EquipmentID))
	builder.WriteString(", ")
	builder.WriteString("connector_id=")
	builder.WriteString(fmt.Sprintf("%v", r.ConnectorID))
	builder.WriteString(", ")
	builder.WriteString("reservation_id=")
	builder.WriteString(fmt.Sprintf("%v", r.ReservationID))
	builder.WriteString(", ")
	builder.WriteString("authorization_mode=")
	builder.WriteString(fmt.Sprintf("%v", r.AuthorizationMode))
	builder.WriteString(", ")
	builder.WriteString("authorization_id=")
	builder.WriteString(r.AuthorizationID)
	builder.WriteString(", ")
	if v := r.Additional; v != nil {
		builder.WriteString("additional=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := r.CustomerID; v != nil {
		builder.WriteString("customer_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("expired=")
	builder.WriteString(fmt.Sprintf("%v", r.Expired))
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(fmt.Sprintf("%v", r.State))
	builder.WriteByte(')')
	return builder.String()
}

// Reservations is a parsable slice of Reservation.
type Reservations []*Reservation

func (r Reservations) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
