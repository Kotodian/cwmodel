// Code generated by ent, DO NOT EDIT.

package cwmodel

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Kotodian/cwmodel/connector"
	"github.com/Kotodian/cwmodel/equipment"
	"github.com/Kotodian/cwmodel/evse"
	"github.com/Kotodian/gokit/datasource"
)

// Connector is the model entity for the Connector schema.
type Connector struct {
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
	EquipmentID datasource.UUID `json:"equipment_id"`
	// 设备id
	EvseID datasource.UUID `json:"evse_id"`
	// 桩序列号
	EquipmentSn string `json:"equipment_sn"`
	// 设备序列号
	EvseSerial string `json:"evse_serial"`
	// 枪序列号
	Serial string `json:"serial"`
	// 当前状态
	CurrentState int `json:"current_state"`
	// 之前状态
	BeforeState int `json:"before_state"`
	// 充电状态
	ChargingState int `json:"charging_state"`
	// 预约id
	ReservationID *datasource.UUID `json:"reservation_id"`
	// 订单id
	OrderID *datasource.UUID `json:"order_id"`
	// 停车编号
	ParkNo string `json:"park_no"`
	// 二维码
	QrCode *string `json:"qr_code"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ConnectorQuery when eager-loading is set.
	Edges ConnectorEdges `json:"-"`

	PushInterval int             `json:"push_interval"`
	LastPushTime int64           `json:"last_push_time"`
	OrderState   *int            `json:"order_state"`
	StationID    datasource.UUID `json:"station_id"`
}

// ConnectorEdges holds the relations/edges for other nodes in the graph.
type ConnectorEdges struct {
	// Evse holds the value of the evse edge.
	Evse *Evse `json:"evse,omitempty"`
	// Equipment holds the value of the equipment edge.
	Equipment *Equipment `json:"equipment,omitempty"`
	// OrderInfo holds the value of the order_info edge.
	OrderInfo []*OrderInfo `json:"order_info,omitempty"`
	// Reservation holds the value of the reservation edge.
	Reservation []*Reservation `json:"reservation,omitempty"`
	// SmartChargingEffect holds the value of the smart_charging_effect edge.
	SmartChargingEffect []*SmartChargingEffect `json:"smart_charging_effect,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// EvseOrErr returns the Evse value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ConnectorEdges) EvseOrErr() (*Evse, error) {
	if e.loadedTypes[0] {
		if e.Evse == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: evse.Label}
		}
		return e.Evse, nil
	}
	return nil, &NotLoadedError{edge: "evse"}
}

// EquipmentOrErr returns the Equipment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ConnectorEdges) EquipmentOrErr() (*Equipment, error) {
	if e.loadedTypes[1] {
		if e.Equipment == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: equipment.Label}
		}
		return e.Equipment, nil
	}
	return nil, &NotLoadedError{edge: "equipment"}
}

// OrderInfoOrErr returns the OrderInfo value or an error if the edge
// was not loaded in eager-loading.
func (e ConnectorEdges) OrderInfoOrErr() ([]*OrderInfo, error) {
	if e.loadedTypes[2] {
		return e.OrderInfo, nil
	}
	return nil, &NotLoadedError{edge: "order_info"}
}

// ReservationOrErr returns the Reservation value or an error if the edge
// was not loaded in eager-loading.
func (e ConnectorEdges) ReservationOrErr() ([]*Reservation, error) {
	if e.loadedTypes[3] {
		return e.Reservation, nil
	}
	return nil, &NotLoadedError{edge: "reservation"}
}

// SmartChargingEffectOrErr returns the SmartChargingEffect value or an error if the edge
// was not loaded in eager-loading.
func (e ConnectorEdges) SmartChargingEffectOrErr() ([]*SmartChargingEffect, error) {
	if e.loadedTypes[4] {
		return e.SmartChargingEffect, nil
	}
	return nil, &NotLoadedError{edge: "smart_charging_effect"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Connector) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case connector.FieldID, connector.FieldVersion, connector.FieldCreatedBy, connector.FieldCreatedAt, connector.FieldUpdatedBy, connector.FieldUpdatedAt, connector.FieldEquipmentID, connector.FieldEvseID, connector.FieldCurrentState, connector.FieldBeforeState, connector.FieldChargingState, connector.FieldReservationID, connector.FieldOrderID:
			values[i] = new(sql.NullInt64)
		case connector.FieldEquipmentSn, connector.FieldEvseSerial, connector.FieldSerial, connector.FieldParkNo, connector.FieldQrCode:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Connector", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Connector fields.
func (c *Connector) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case connector.FieldID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				c.ID = datasource.UUID(value.Int64)
			}
		case connector.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				c.Version = value.Int64
			}
		case connector.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				c.CreatedBy = datasource.UUID(value.Int64)
			}
		case connector.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Int64
			}
		case connector.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				c.UpdatedBy = datasource.UUID(value.Int64)
			}
		case connector.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Int64
			}
		case connector.FieldEquipmentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field equipment_id", values[i])
			} else if value.Valid {
				c.EquipmentID = datasource.UUID(value.Int64)
			}
		case connector.FieldEvseID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field evse_id", values[i])
			} else if value.Valid {
				c.EvseID = datasource.UUID(value.Int64)
			}
		case connector.FieldEquipmentSn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field equipment_sn", values[i])
			} else if value.Valid {
				c.EquipmentSn = value.String
			}
		case connector.FieldEvseSerial:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field evse_serial", values[i])
			} else if value.Valid {
				c.EvseSerial = value.String
			}
		case connector.FieldSerial:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field serial", values[i])
			} else if value.Valid {
				c.Serial = value.String
			}
		case connector.FieldCurrentState:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field current_state", values[i])
			} else if value.Valid {
				c.CurrentState = int(value.Int64)
			}
		case connector.FieldBeforeState:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field before_state", values[i])
			} else if value.Valid {
				c.BeforeState = int(value.Int64)
			}
		case connector.FieldChargingState:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field charging_state", values[i])
			} else if value.Valid {
				c.ChargingState = int(value.Int64)
			}
		case connector.FieldReservationID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field reservation_id", values[i])
			} else if value.Valid {
				c.ReservationID = new(datasource.UUID)
				*c.ReservationID = datasource.UUID(value.Int64)
			}
		case connector.FieldOrderID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value.Valid {
				c.OrderID = new(datasource.UUID)
				*c.OrderID = datasource.UUID(value.Int64)
			}
		case connector.FieldParkNo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field park_no", values[i])
			} else if value.Valid {
				c.ParkNo = value.String
			}
		case connector.FieldQrCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field qr_code", values[i])
			} else if value.Valid {
				c.QrCode = new(string)
				*c.QrCode = value.String
			}
		}
	}
	return nil
}

// QueryEvse queries the "evse" edge of the Connector entity.
func (c *Connector) QueryEvse() *EvseQuery {
	return (&ConnectorClient{config: c.config}).QueryEvse(c)
}

// QueryEquipment queries the "equipment" edge of the Connector entity.
func (c *Connector) QueryEquipment() *EquipmentQuery {
	return (&ConnectorClient{config: c.config}).QueryEquipment(c)
}

// QueryOrderInfo queries the "order_info" edge of the Connector entity.
func (c *Connector) QueryOrderInfo() *OrderInfoQuery {
	return (&ConnectorClient{config: c.config}).QueryOrderInfo(c)
}

// QueryReservation queries the "reservation" edge of the Connector entity.
func (c *Connector) QueryReservation() *ReservationQuery {
	return (&ConnectorClient{config: c.config}).QueryReservation(c)
}

// QuerySmartChargingEffect queries the "smart_charging_effect" edge of the Connector entity.
func (c *Connector) QuerySmartChargingEffect() *SmartChargingEffectQuery {
	return (&ConnectorClient{config: c.config}).QuerySmartChargingEffect(c)
}

// Update returns a builder for updating this Connector.
// Note that you need to call Connector.Unwrap() before calling this method if this Connector
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Connector) Update() *ConnectorUpdateOne {
	return (&ConnectorClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Connector entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Connector) Unwrap() *Connector {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("cwmodel: Connector is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Connector) String() string {
	var builder strings.Builder
	builder.WriteString("Connector(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", c.Version))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", c.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", c.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", c.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", c.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("equipment_id=")
	builder.WriteString(fmt.Sprintf("%v", c.EquipmentID))
	builder.WriteString(", ")
	builder.WriteString("evse_id=")
	builder.WriteString(fmt.Sprintf("%v", c.EvseID))
	builder.WriteString(", ")
	builder.WriteString("equipment_sn=")
	builder.WriteString(c.EquipmentSn)
	builder.WriteString(", ")
	builder.WriteString("evse_serial=")
	builder.WriteString(c.EvseSerial)
	builder.WriteString(", ")
	builder.WriteString("serial=")
	builder.WriteString(c.Serial)
	builder.WriteString(", ")
	builder.WriteString("current_state=")
	builder.WriteString(fmt.Sprintf("%v", c.CurrentState))
	builder.WriteString(", ")
	builder.WriteString("before_state=")
	builder.WriteString(fmt.Sprintf("%v", c.BeforeState))
	builder.WriteString(", ")
	builder.WriteString("charging_state=")
	builder.WriteString(fmt.Sprintf("%v", c.ChargingState))
	builder.WriteString(", ")
	if v := c.ReservationID; v != nil {
		builder.WriteString("reservation_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := c.OrderID; v != nil {
		builder.WriteString("order_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("park_no=")
	builder.WriteString(c.ParkNo)
	builder.WriteString(", ")
	if v := c.QrCode; v != nil {
		builder.WriteString("qr_code=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Connectors is a parsable slice of Connector.
type Connectors []*Connector

func (c Connectors) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
