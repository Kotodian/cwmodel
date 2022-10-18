// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Kotodian/ent-practice/ent/firmware"
	"github.com/Kotodian/gokit/datasource"
)

// Firmware is the model entity for the Firmware schema.
type Firmware struct {
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
	// 固件版本
	EquipVersion string `json:"equip_version,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FirmwareQuery when eager-loading is set.
	Edges FirmwareEdges `json:"-"`
}

// FirmwareEdges holds the relations/edges for other nodes in the graph.
type FirmwareEdges struct {
	// EquipmentFirmwareEffect holds the value of the equipment_firmware_effect edge.
	EquipmentFirmwareEffect []*EquipmentFirmwareEffect `json:"equipment_firmware_effect,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EquipmentFirmwareEffectOrErr returns the EquipmentFirmwareEffect value or an error if the edge
// was not loaded in eager-loading.
func (e FirmwareEdges) EquipmentFirmwareEffectOrErr() ([]*EquipmentFirmwareEffect, error) {
	if e.loadedTypes[0] {
		return e.EquipmentFirmwareEffect, nil
	}
	return nil, &NotLoadedError{edge: "equipment_firmware_effect"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Firmware) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case firmware.FieldID, firmware.FieldVersion, firmware.FieldCreatedBy, firmware.FieldCreatedAt, firmware.FieldUpdatedBy, firmware.FieldUpdatedAt:
			values[i] = new(sql.NullInt64)
		case firmware.FieldEquipVersion:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Firmware", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Firmware fields.
func (f *Firmware) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case firmware.FieldID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				f.ID = datasource.UUID(value.Int64)
			}
		case firmware.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				f.Version = value.Int64
			}
		case firmware.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				f.CreatedBy = datasource.UUID(value.Int64)
			}
		case firmware.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				f.CreatedAt = value.Int64
			}
		case firmware.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				f.UpdatedBy = datasource.UUID(value.Int64)
			}
		case firmware.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				f.UpdatedAt = value.Int64
			}
		case firmware.FieldEquipVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field equip_version", values[i])
			} else if value.Valid {
				f.EquipVersion = value.String
			}
		}
	}
	return nil
}

// QueryEquipmentFirmwareEffect queries the "equipment_firmware_effect" edge of the Firmware entity.
func (f *Firmware) QueryEquipmentFirmwareEffect() *EquipmentFirmwareEffectQuery {
	return (&FirmwareClient{config: f.config}).QueryEquipmentFirmwareEffect(f)
}

// Update returns a builder for updating this Firmware.
// Note that you need to call Firmware.Unwrap() before calling this method if this Firmware
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Firmware) Update() *FirmwareUpdateOne {
	return (&FirmwareClient{config: f.config}).UpdateOne(f)
}

// Unwrap unwraps the Firmware entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Firmware) Unwrap() *Firmware {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Firmware is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Firmware) String() string {
	var builder strings.Builder
	builder.WriteString("Firmware(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", f.Version))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", f.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", f.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", f.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", f.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("equip_version=")
	builder.WriteString(f.EquipVersion)
	builder.WriteByte(')')
	return builder.String()
}

// Firmwares is a parsable slice of Firmware.
type Firmwares []*Firmware

func (f Firmwares) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}
