// Code generated by ent, DO NOT EDIT.

package cwmodel

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Kotodian/cwmodel/manufacturer"
	"github.com/Kotodian/gokit/datasource"
)

// Manufacturer is the model entity for the Manufacturer schema.
type Manufacturer struct {
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
	// 产商代码
	Code string `json:"code,omitempty"`
	// 产商名称
	Name *string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ManufacturerQuery when eager-loading is set.
	Edges ManufacturerEdges `json:"-"`
}

// ManufacturerEdges holds the relations/edges for other nodes in the graph.
type ManufacturerEdges struct {
	// Firmware holds the value of the firmware edge.
	Firmware []*Firmware `json:"firmware,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// FirmwareOrErr returns the Firmware value or an error if the edge
// was not loaded in eager-loading.
func (e ManufacturerEdges) FirmwareOrErr() ([]*Firmware, error) {
	if e.loadedTypes[0] {
		return e.Firmware, nil
	}
	return nil, &NotLoadedError{edge: "firmware"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Manufacturer) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case manufacturer.FieldID, manufacturer.FieldVersion, manufacturer.FieldCreatedBy, manufacturer.FieldCreatedAt, manufacturer.FieldUpdatedBy, manufacturer.FieldUpdatedAt:
			values[i] = new(sql.NullInt64)
		case manufacturer.FieldCode, manufacturer.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Manufacturer", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Manufacturer fields.
func (m *Manufacturer) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case manufacturer.FieldID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				m.ID = datasource.UUID(value.Int64)
			}
		case manufacturer.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				m.Version = value.Int64
			}
		case manufacturer.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				m.CreatedBy = datasource.UUID(value.Int64)
			}
		case manufacturer.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				m.CreatedAt = value.Int64
			}
		case manufacturer.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				m.UpdatedBy = datasource.UUID(value.Int64)
			}
		case manufacturer.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				m.UpdatedAt = value.Int64
			}
		case manufacturer.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				m.Code = value.String
			}
		case manufacturer.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				m.Name = new(string)
				*m.Name = value.String
			}
		}
	}
	return nil
}

// QueryFirmware queries the "firmware" edge of the Manufacturer entity.
func (m *Manufacturer) QueryFirmware() *FirmwareQuery {
	return (&ManufacturerClient{config: m.config}).QueryFirmware(m)
}

// Update returns a builder for updating this Manufacturer.
// Note that you need to call Manufacturer.Unwrap() before calling this method if this Manufacturer
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Manufacturer) Update() *ManufacturerUpdateOne {
	return (&ManufacturerClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the Manufacturer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Manufacturer) Unwrap() *Manufacturer {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("cwmodel: Manufacturer is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Manufacturer) String() string {
	var builder strings.Builder
	builder.WriteString("Manufacturer(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", m.Version))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", m.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", m.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", m.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", m.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(m.Code)
	builder.WriteString(", ")
	if v := m.Name; v != nil {
		builder.WriteString("name=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Manufacturers is a parsable slice of Manufacturer.
type Manufacturers []*Manufacturer

func (m Manufacturers) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
