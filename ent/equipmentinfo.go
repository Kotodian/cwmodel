// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Kotodian/ent-practice/ent/equipment"
	"github.com/Kotodian/ent-practice/ent/equipmentinfo"
)

// EquipmentInfo is the model entity for the EquipmentInfo schema.
type EquipmentInfo struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// EquipmentSn holds the value of the "equipment_sn" field.
	// 桩序列号
	EquipmentSn string `json:"equipment_sn,omitempty"`
	// AccessPod holds the value of the "access_pod" field.
	// 所连接的pod
	AccessPod string `json:"access_pod,omitempty"`
	// State holds the value of the "state" field.
	// 在线或者离线
	State bool `json:"state,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EquipmentInfoQuery when eager-loading is set.
	Edges                    EquipmentInfoEdges `json:"edges"`
	equipment_equipment_info *int
}

// EquipmentInfoEdges holds the relations/edges for other nodes in the graph.
type EquipmentInfoEdges struct {
	// Equipment holds the value of the equipment edge.
	Equipment *Equipment `json:"equipment,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EquipmentOrErr returns the Equipment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EquipmentInfoEdges) EquipmentOrErr() (*Equipment, error) {
	if e.loadedTypes[0] {
		if e.Equipment == nil {
			// The edge equipment was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: equipment.Label}
		}
		return e.Equipment, nil
	}
	return nil, &NotLoadedError{edge: "equipment"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EquipmentInfo) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case equipmentinfo.FieldState:
			values[i] = new(sql.NullBool)
		case equipmentinfo.FieldID:
			values[i] = new(sql.NullInt64)
		case equipmentinfo.FieldEquipmentSn, equipmentinfo.FieldAccessPod:
			values[i] = new(sql.NullString)
		case equipmentinfo.ForeignKeys[0]: // equipment_equipment_info
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type EquipmentInfo", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EquipmentInfo fields.
func (ei *EquipmentInfo) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case equipmentinfo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ei.ID = int(value.Int64)
		case equipmentinfo.FieldEquipmentSn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field equipment_sn", values[i])
			} else if value.Valid {
				ei.EquipmentSn = value.String
			}
		case equipmentinfo.FieldAccessPod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field access_pod", values[i])
			} else if value.Valid {
				ei.AccessPod = value.String
			}
		case equipmentinfo.FieldState:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				ei.State = value.Bool
			}
		case equipmentinfo.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field equipment_equipment_info", value)
			} else if value.Valid {
				ei.equipment_equipment_info = new(int)
				*ei.equipment_equipment_info = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryEquipment queries the "equipment" edge of the EquipmentInfo entity.
func (ei *EquipmentInfo) QueryEquipment() *EquipmentQuery {
	return (&EquipmentInfoClient{config: ei.config}).QueryEquipment(ei)
}

// Update returns a builder for updating this EquipmentInfo.
// Note that you need to call EquipmentInfo.Unwrap() before calling this method if this EquipmentInfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (ei *EquipmentInfo) Update() *EquipmentInfoUpdateOne {
	return (&EquipmentInfoClient{config: ei.config}).UpdateOne(ei)
}

// Unwrap unwraps the EquipmentInfo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ei *EquipmentInfo) Unwrap() *EquipmentInfo {
	tx, ok := ei.config.driver.(*txDriver)
	if !ok {
		panic("ent: EquipmentInfo is not a transactional entity")
	}
	ei.config.driver = tx.drv
	return ei
}

// String implements the fmt.Stringer.
func (ei *EquipmentInfo) String() string {
	var builder strings.Builder
	builder.WriteString("EquipmentInfo(")
	builder.WriteString(fmt.Sprintf("id=%v", ei.ID))
	builder.WriteString(", equipment_sn=")
	builder.WriteString(ei.EquipmentSn)
	builder.WriteString(", access_pod=")
	builder.WriteString(ei.AccessPod)
	builder.WriteString(", state=")
	builder.WriteString(fmt.Sprintf("%v", ei.State))
	builder.WriteByte(')')
	return builder.String()
}

// EquipmentInfos is a parsable slice of EquipmentInfo.
type EquipmentInfos []*EquipmentInfo

func (ei EquipmentInfos) config(cfg config) {
	for _i := range ei {
		ei[_i].config = cfg
	}
}
