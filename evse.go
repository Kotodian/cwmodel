// Code generated by ent, DO NOT EDIT.

package cwmodel

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Kotodian/cwmodel/equipment"
	"github.com/Kotodian/cwmodel/evse"
	"github.com/Kotodian/gokit/datasource"
)

// Evse is the model entity for the Evse schema.
type Evse struct {
	config `json:"-"`
	// ID of the ent.
	// 主键
	ID datasource.UUID `json:"id,omitempty"`
	// 乐观锁
	Version int64 `json:"version,omitempty"`
	// 创建者
	CreatedBy datasource.UUID `json:"created_by,omitempty"`
	// 创建时间
	CreatedAt int64 `json:"created_at,omitempty"`
	// 修改者
	UpdatedBy datasource.UUID `json:"updated_by,omitempty"`
	// 修改时间
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// 设备序列号
	Serial string `json:"serial,omitempty"`
	// 枪数量
	ConnectorNumber int `json:"connector_number,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EvseQuery when eager-loading is set.
	Edges        EvseEdges `json:"-"`
	equipment_id *datasource.UUID
}

// EvseEdges holds the relations/edges for other nodes in the graph.
type EvseEdges struct {
	// Equipment holds the value of the equipment edge.
	Equipment *Equipment `json:"equipment,omitempty"`
	// Connector holds the value of the connector edge.
	Connector []*Connector `json:"connector,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// EquipmentOrErr returns the Equipment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EvseEdges) EquipmentOrErr() (*Equipment, error) {
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
// was not loaded in eager-loading.
func (e EvseEdges) ConnectorOrErr() ([]*Connector, error) {
	if e.loadedTypes[1] {
		return e.Connector, nil
	}
	return nil, &NotLoadedError{edge: "connector"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Evse) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case evse.FieldID, evse.FieldVersion, evse.FieldCreatedBy, evse.FieldCreatedAt, evse.FieldUpdatedBy, evse.FieldUpdatedAt, evse.FieldConnectorNumber:
			values[i] = new(sql.NullInt64)
		case evse.FieldSerial:
			values[i] = new(sql.NullString)
		case evse.ForeignKeys[0]: // equipment_id
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Evse", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Evse fields.
func (e *Evse) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case evse.FieldID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				e.ID = datasource.UUID(value.Int64)
			}
		case evse.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				e.Version = value.Int64
			}
		case evse.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				e.CreatedBy = datasource.UUID(value.Int64)
			}
		case evse.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				e.CreatedAt = value.Int64
			}
		case evse.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				e.UpdatedBy = datasource.UUID(value.Int64)
			}
		case evse.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				e.UpdatedAt = value.Int64
			}
		case evse.FieldSerial:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field serial", values[i])
			} else if value.Valid {
				e.Serial = value.String
			}
		case evse.FieldConnectorNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field connector_number", values[i])
			} else if value.Valid {
				e.ConnectorNumber = int(value.Int64)
			}
		case evse.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field equipment_id", values[i])
			} else if value.Valid {
				e.equipment_id = new(datasource.UUID)
				*e.equipment_id = datasource.UUID(value.Int64)
			}
		}
	}
	return nil
}

// QueryEquipment queries the "equipment" edge of the Evse entity.
func (e *Evse) QueryEquipment() *EquipmentQuery {
	return (&EvseClient{config: e.config}).QueryEquipment(e)
}

// QueryConnector queries the "connector" edge of the Evse entity.
func (e *Evse) QueryConnector() *ConnectorQuery {
	return (&EvseClient{config: e.config}).QueryConnector(e)
}

// Update returns a builder for updating this Evse.
// Note that you need to call Evse.Unwrap() before calling this method if this Evse
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Evse) Update() *EvseUpdateOne {
	return (&EvseClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the Evse entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Evse) Unwrap() *Evse {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("cwmodel: Evse is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Evse) String() string {
	var builder strings.Builder
	builder.WriteString("Evse(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", e.Version))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", e.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", e.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", e.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", e.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("serial=")
	builder.WriteString(e.Serial)
	builder.WriteString(", ")
	builder.WriteString("connector_number=")
	builder.WriteString(fmt.Sprintf("%v", e.ConnectorNumber))
	builder.WriteByte(')')
	return builder.String()
}

// Evses is a parsable slice of Evse.
type Evses []*Evse

func (e Evses) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}