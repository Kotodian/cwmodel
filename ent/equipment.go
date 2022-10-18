// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Kotodian/ent-practice/ent/equipment"
	"github.com/Kotodian/ent-practice/ent/equipmentinfo"
	"github.com/Kotodian/ent-practice/ent/equipmentiot"
	"github.com/Kotodian/gokit/datasource"
)

// Equipment is the model entity for the Equipment schema.
type Equipment struct {
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
	// 桩序列号
	Sn string `json:"sn"`
	// 运营商id
	OperatorID datasource.UUID `json:"operatorId"`
	// 站点id
	StationID datasource.UUID `json:"stationId"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EquipmentQuery when eager-loading is set.
	Edges EquipmentEdges `json:"-"`

	EquipmentInfo *EquipmentInfo `json:"equipmentInfo"`
}

// EquipmentEdges holds the relations/edges for other nodes in the graph.
type EquipmentEdges struct {
	// EquipmentInfo holds the value of the equipment_info edge.
	EquipmentInfo *EquipmentInfo `json:"equipmentId"`
	// Evse holds the value of the evse edge.
	Evse []*Evse `json:"evse,omitempty"`
	// Connector holds the value of the connector edge.
	Connector []*Connector `json:"connector,omitempty"`
	// EquipmentAlarm holds the value of the equipment_alarm edge.
	EquipmentAlarm []*EquipmentAlarm `json:"equipment_alarm,omitempty"`
	// EquipmentIot holds the value of the equipment_iot edge.
	EquipmentIot *EquipmentIot `json:"equipment_iot,omitempty"`
	// EquipmentFirmwareEffect holds the value of the equipment_firmware_effect edge.
	EquipmentFirmwareEffect []*EquipmentFirmwareEffect `json:"equipment_firmware_effect,omitempty"`
	// OrderInfo holds the value of the order_info edge.
	OrderInfo []*OrderInfo `json:"order_info,omitempty"`
	// Reservation holds the value of the reservation edge.
	Reservation []*Reservation `json:"reservation,omitempty"`
	// EquipmentLog holds the value of the equipment_log edge.
	EquipmentLog []*EquipmentLog `json:"equipment_log,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [9]bool
}

// EquipmentInfoOrErr returns the EquipmentInfo value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EquipmentEdges) EquipmentInfoOrErr() (*EquipmentInfo, error) {
	if e.loadedTypes[0] {
		if e.EquipmentInfo == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: equipmentinfo.Label}
		}
		return e.EquipmentInfo, nil
	}
	return nil, &NotLoadedError{edge: "equipment_info"}
}

// EvseOrErr returns the Evse value or an error if the edge
// was not loaded in eager-loading.
func (e EquipmentEdges) EvseOrErr() ([]*Evse, error) {
	if e.loadedTypes[1] {
		return e.Evse, nil
	}
	return nil, &NotLoadedError{edge: "evse"}
}

// ConnectorOrErr returns the Connector value or an error if the edge
// was not loaded in eager-loading.
func (e EquipmentEdges) ConnectorOrErr() ([]*Connector, error) {
	if e.loadedTypes[2] {
		return e.Connector, nil
	}
	return nil, &NotLoadedError{edge: "connector"}
}

// EquipmentAlarmOrErr returns the EquipmentAlarm value or an error if the edge
// was not loaded in eager-loading.
func (e EquipmentEdges) EquipmentAlarmOrErr() ([]*EquipmentAlarm, error) {
	if e.loadedTypes[3] {
		return e.EquipmentAlarm, nil
	}
	return nil, &NotLoadedError{edge: "equipment_alarm"}
}

// EquipmentIotOrErr returns the EquipmentIot value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EquipmentEdges) EquipmentIotOrErr() (*EquipmentIot, error) {
	if e.loadedTypes[4] {
		if e.EquipmentIot == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: equipmentiot.Label}
		}
		return e.EquipmentIot, nil
	}
	return nil, &NotLoadedError{edge: "equipment_iot"}
}

// EquipmentFirmwareEffectOrErr returns the EquipmentFirmwareEffect value or an error if the edge
// was not loaded in eager-loading.
func (e EquipmentEdges) EquipmentFirmwareEffectOrErr() ([]*EquipmentFirmwareEffect, error) {
	if e.loadedTypes[5] {
		return e.EquipmentFirmwareEffect, nil
	}
	return nil, &NotLoadedError{edge: "equipment_firmware_effect"}
}

// OrderInfoOrErr returns the OrderInfo value or an error if the edge
// was not loaded in eager-loading.
func (e EquipmentEdges) OrderInfoOrErr() ([]*OrderInfo, error) {
	if e.loadedTypes[6] {
		return e.OrderInfo, nil
	}
	return nil, &NotLoadedError{edge: "order_info"}
}

// ReservationOrErr returns the Reservation value or an error if the edge
// was not loaded in eager-loading.
func (e EquipmentEdges) ReservationOrErr() ([]*Reservation, error) {
	if e.loadedTypes[7] {
		return e.Reservation, nil
	}
	return nil, &NotLoadedError{edge: "reservation"}
}

// EquipmentLogOrErr returns the EquipmentLog value or an error if the edge
// was not loaded in eager-loading.
func (e EquipmentEdges) EquipmentLogOrErr() ([]*EquipmentLog, error) {
	if e.loadedTypes[8] {
		return e.EquipmentLog, nil
	}
	return nil, &NotLoadedError{edge: "equipment_log"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Equipment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case equipment.FieldID, equipment.FieldVersion, equipment.FieldCreatedBy, equipment.FieldCreatedAt, equipment.FieldUpdatedBy, equipment.FieldUpdatedAt, equipment.FieldOperatorID, equipment.FieldStationID:
			values[i] = new(sql.NullInt64)
		case equipment.FieldSn:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Equipment", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Equipment fields.
func (e *Equipment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case equipment.FieldID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				e.ID = datasource.UUID(value.Int64)
			}
		case equipment.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				e.Version = value.Int64
			}
		case equipment.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				e.CreatedBy = datasource.UUID(value.Int64)
			}
		case equipment.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				e.CreatedAt = value.Int64
			}
		case equipment.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				e.UpdatedBy = datasource.UUID(value.Int64)
			}
		case equipment.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				e.UpdatedAt = value.Int64
			}
		case equipment.FieldSn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sn", values[i])
			} else if value.Valid {
				e.Sn = value.String
			}
		case equipment.FieldOperatorID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field operator_id", values[i])
			} else if value.Valid {
				e.OperatorID = datasource.UUID(value.Int64)
			}
		case equipment.FieldStationID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field station_id", values[i])
			} else if value.Valid {
				e.StationID = datasource.UUID(value.Int64)
			}
		}
	}
	return nil
}

// QueryEquipmentInfo queries the "equipment_info" edge of the Equipment entity.
func (e *Equipment) QueryEquipmentInfo() *EquipmentInfoQuery {
	return (&EquipmentClient{config: e.config}).QueryEquipmentInfo(e)
}

// QueryEvse queries the "evse" edge of the Equipment entity.
func (e *Equipment) QueryEvse() *EvseQuery {
	return (&EquipmentClient{config: e.config}).QueryEvse(e)
}

// QueryConnector queries the "connector" edge of the Equipment entity.
func (e *Equipment) QueryConnector() *ConnectorQuery {
	return (&EquipmentClient{config: e.config}).QueryConnector(e)
}

// QueryEquipmentAlarm queries the "equipment_alarm" edge of the Equipment entity.
func (e *Equipment) QueryEquipmentAlarm() *EquipmentAlarmQuery {
	return (&EquipmentClient{config: e.config}).QueryEquipmentAlarm(e)
}

// QueryEquipmentIot queries the "equipment_iot" edge of the Equipment entity.
func (e *Equipment) QueryEquipmentIot() *EquipmentIotQuery {
	return (&EquipmentClient{config: e.config}).QueryEquipmentIot(e)
}

// QueryEquipmentFirmwareEffect queries the "equipment_firmware_effect" edge of the Equipment entity.
func (e *Equipment) QueryEquipmentFirmwareEffect() *EquipmentFirmwareEffectQuery {
	return (&EquipmentClient{config: e.config}).QueryEquipmentFirmwareEffect(e)
}

// QueryOrderInfo queries the "order_info" edge of the Equipment entity.
func (e *Equipment) QueryOrderInfo() *OrderInfoQuery {
	return (&EquipmentClient{config: e.config}).QueryOrderInfo(e)
}

// QueryReservation queries the "reservation" edge of the Equipment entity.
func (e *Equipment) QueryReservation() *ReservationQuery {
	return (&EquipmentClient{config: e.config}).QueryReservation(e)
}

// QueryEquipmentLog queries the "equipment_log" edge of the Equipment entity.
func (e *Equipment) QueryEquipmentLog() *EquipmentLogQuery {
	return (&EquipmentClient{config: e.config}).QueryEquipmentLog(e)
}

// Update returns a builder for updating this Equipment.
// Note that you need to call Equipment.Unwrap() before calling this method if this Equipment
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Equipment) Update() *EquipmentUpdateOne {
	return (&EquipmentClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the Equipment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Equipment) Unwrap() *Equipment {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Equipment is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Equipment) String() string {
	var builder strings.Builder
	builder.WriteString("Equipment(")
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
	builder.WriteString("sn=")
	builder.WriteString(e.Sn)
	builder.WriteString(", ")
	builder.WriteString("operator_id=")
	builder.WriteString(fmt.Sprintf("%v", e.OperatorID))
	builder.WriteString(", ")
	builder.WriteString("station_id=")
	builder.WriteString(fmt.Sprintf("%v", e.StationID))
	builder.WriteByte(')')
	return builder.String()
}

// EquipmentSlice is a parsable slice of Equipment.
type EquipmentSlice []*Equipment

func (e EquipmentSlice) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}
