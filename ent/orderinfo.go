// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Kotodian/ent-practice/ent/connector"
	"github.com/Kotodian/ent-practice/ent/equipment"
	"github.com/Kotodian/ent-practice/ent/orderinfo"
	"github.com/Kotodian/gokit/datasource"
)

// OrderInfo is the model entity for the OrderInfo schema.
type OrderInfo struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// RemoteStartID holds the value of the "remote_start_id" field.
	// 远程启动id
	RemoteStartID int64 `json:"remote_start_id,omitempty"`
	// TransactionID holds the value of the "transaction_id" field.
	// 桩端订单id
	TransactionID string `json:"transaction_id,omitempty"`
	// AuthorizationID holds the value of the "authorization_id" field.
	// 授权id
	AuthorizationID string `json:"authorization_id,omitempty"`
	// CustomerID holds the value of the "customer_id" field.
	// 客户id
	CustomerID string `json:"customer_id,omitempty"`
	// CallerOrderID holds the value of the "caller_order_id" field.
	// 第三方订单id
	CallerOrderID string `json:"caller_order_id,omitempty"`
	// TotalElectricity holds the value of the "total_electricity" field.
	// 总充电电量
	TotalElectricity float64 `json:"total_electricity,omitempty"`
	// ChargeStartElectricity holds the value of the "charge_start_electricity" field.
	// 起始电量
	ChargeStartElectricity float64 `json:"charge_start_electricity,omitempty"`
	// ChargeStopElectricity holds the value of the "charge_stop_electricity" field.
	// 结束电量
	ChargeStopElectricity float64 `json:"charge_stop_electricity,omitempty"`
	// SharpElectricity holds the value of the "sharp_electricity" field.
	// 尖电量
	SharpElectricity float64 `json:"sharp_electricity,omitempty"`
	// PeakElectricity holds the value of the "peak_electricity" field.
	// 峰电量
	PeakElectricity float64 `json:"peak_electricity,omitempty"`
	// FlatElectricity holds the value of the "flat_electricity" field.
	// 平电量
	FlatElectricity float64 `json:"flat_electricity,omitempty"`
	// ValleyElectricity holds the value of the "valley_electricity" field.
	// 谷电量
	ValleyElectricity float64 `json:"valley_electricity,omitempty"`
	// StopReasonCode holds the value of the "stop_reason_code" field.
	// 停止原因代码
	StopReasonCode int32 `json:"stop_reason_code,omitempty"`
	// Offline holds the value of the "offline" field.
	// 是否为离线订单
	Offline bool `json:"offline,omitempty"`
	// PriceSchemeReleaseID holds the value of the "price_scheme_release_id" field.
	// 计费模板id
	PriceSchemeReleaseID int64 `json:"price_scheme_release_id,omitempty"`
	// OrderStartTime holds the value of the "order_start_time" field.
	// 订单开始时间
	OrderStartTime int64 `json:"order_start_time,omitempty"`
	// OrderFinalTime holds the value of the "order_final_time" field.
	// 订单结束时间
	OrderFinalTime int64 `json:"order_final_time,omitempty"`
	// ChargeStartTime holds the value of the "charge_start_time" field.
	// 充电开始时间
	ChargeStartTime int64 `json:"charge_start_time,omitempty"`
	// ChargeFinalTime holds the value of the "charge_final_time" field.
	// 充电结束时间
	ChargeFinalTime int64 `json:"charge_final_time,omitempty"`
	// IntellectID holds the value of the "intellect_id" field.
	// 智慧充电id
	IntellectID int64 `json:"intellect_id,omitempty"`
	// StationID holds the value of the "station_id" field.
	// 场站id
	StationID datasource.UUID `json:"station_id,omitempty"`
	// OperatorID holds the value of the "operator_id" field.
	// 运营商id
	OperatorID datasource.UUID `json:"operator_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderInfoQuery when eager-loading is set.
	Edges                OrderInfoEdges `json:"edges"`
	connector_order_info *int
	equipment_order_info *int
}

// OrderInfoEdges holds the relations/edges for other nodes in the graph.
type OrderInfoEdges struct {
	// Connector holds the value of the connector edge.
	Connector *Connector `json:"connector,omitempty"`
	// Equipment holds the value of the equipment edge.
	Equipment *Equipment `json:"equipment,omitempty"`
	// OrderEvent holds the value of the order_event edge.
	OrderEvent []*OrderEvent `json:"order_event,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ConnectorOrErr returns the Connector value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderInfoEdges) ConnectorOrErr() (*Connector, error) {
	if e.loadedTypes[0] {
		if e.Connector == nil {
			// The edge connector was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: connector.Label}
		}
		return e.Connector, nil
	}
	return nil, &NotLoadedError{edge: "connector"}
}

// EquipmentOrErr returns the Equipment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderInfoEdges) EquipmentOrErr() (*Equipment, error) {
	if e.loadedTypes[1] {
		if e.Equipment == nil {
			// The edge equipment was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: equipment.Label}
		}
		return e.Equipment, nil
	}
	return nil, &NotLoadedError{edge: "equipment"}
}

// OrderEventOrErr returns the OrderEvent value or an error if the edge
// was not loaded in eager-loading.
func (e OrderInfoEdges) OrderEventOrErr() ([]*OrderEvent, error) {
	if e.loadedTypes[2] {
		return e.OrderEvent, nil
	}
	return nil, &NotLoadedError{edge: "order_event"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrderInfo) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case orderinfo.FieldOffline:
			values[i] = new(sql.NullBool)
		case orderinfo.FieldTotalElectricity, orderinfo.FieldChargeStartElectricity, orderinfo.FieldChargeStopElectricity, orderinfo.FieldSharpElectricity, orderinfo.FieldPeakElectricity, orderinfo.FieldFlatElectricity, orderinfo.FieldValleyElectricity:
			values[i] = new(sql.NullFloat64)
		case orderinfo.FieldID, orderinfo.FieldRemoteStartID, orderinfo.FieldStopReasonCode, orderinfo.FieldPriceSchemeReleaseID, orderinfo.FieldOrderStartTime, orderinfo.FieldOrderFinalTime, orderinfo.FieldChargeStartTime, orderinfo.FieldChargeFinalTime, orderinfo.FieldIntellectID, orderinfo.FieldStationID, orderinfo.FieldOperatorID:
			values[i] = new(sql.NullInt64)
		case orderinfo.FieldTransactionID, orderinfo.FieldAuthorizationID, orderinfo.FieldCustomerID, orderinfo.FieldCallerOrderID:
			values[i] = new(sql.NullString)
		case orderinfo.ForeignKeys[0]: // connector_order_info
			values[i] = new(sql.NullInt64)
		case orderinfo.ForeignKeys[1]: // equipment_order_info
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OrderInfo", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrderInfo fields.
func (oi *OrderInfo) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orderinfo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			oi.ID = int(value.Int64)
		case orderinfo.FieldRemoteStartID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field remote_start_id", values[i])
			} else if value.Valid {
				oi.RemoteStartID = value.Int64
			}
		case orderinfo.FieldTransactionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field transaction_id", values[i])
			} else if value.Valid {
				oi.TransactionID = value.String
			}
		case orderinfo.FieldAuthorizationID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field authorization_id", values[i])
			} else if value.Valid {
				oi.AuthorizationID = value.String
			}
		case orderinfo.FieldCustomerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field customer_id", values[i])
			} else if value.Valid {
				oi.CustomerID = value.String
			}
		case orderinfo.FieldCallerOrderID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field caller_order_id", values[i])
			} else if value.Valid {
				oi.CallerOrderID = value.String
			}
		case orderinfo.FieldTotalElectricity:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field total_electricity", values[i])
			} else if value.Valid {
				oi.TotalElectricity = value.Float64
			}
		case orderinfo.FieldChargeStartElectricity:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field charge_start_electricity", values[i])
			} else if value.Valid {
				oi.ChargeStartElectricity = value.Float64
			}
		case orderinfo.FieldChargeStopElectricity:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field charge_stop_electricity", values[i])
			} else if value.Valid {
				oi.ChargeStopElectricity = value.Float64
			}
		case orderinfo.FieldSharpElectricity:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field sharp_electricity", values[i])
			} else if value.Valid {
				oi.SharpElectricity = value.Float64
			}
		case orderinfo.FieldPeakElectricity:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field peak_electricity", values[i])
			} else if value.Valid {
				oi.PeakElectricity = value.Float64
			}
		case orderinfo.FieldFlatElectricity:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field flat_electricity", values[i])
			} else if value.Valid {
				oi.FlatElectricity = value.Float64
			}
		case orderinfo.FieldValleyElectricity:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field valley_electricity", values[i])
			} else if value.Valid {
				oi.ValleyElectricity = value.Float64
			}
		case orderinfo.FieldStopReasonCode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field stop_reason_code", values[i])
			} else if value.Valid {
				oi.StopReasonCode = int32(value.Int64)
			}
		case orderinfo.FieldOffline:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field offline", values[i])
			} else if value.Valid {
				oi.Offline = value.Bool
			}
		case orderinfo.FieldPriceSchemeReleaseID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field price_scheme_release_id", values[i])
			} else if value.Valid {
				oi.PriceSchemeReleaseID = value.Int64
			}
		case orderinfo.FieldOrderStartTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_start_time", values[i])
			} else if value.Valid {
				oi.OrderStartTime = value.Int64
			}
		case orderinfo.FieldOrderFinalTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_final_time", values[i])
			} else if value.Valid {
				oi.OrderFinalTime = value.Int64
			}
		case orderinfo.FieldChargeStartTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field charge_start_time", values[i])
			} else if value.Valid {
				oi.ChargeStartTime = value.Int64
			}
		case orderinfo.FieldChargeFinalTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field charge_final_time", values[i])
			} else if value.Valid {
				oi.ChargeFinalTime = value.Int64
			}
		case orderinfo.FieldIntellectID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field intellect_id", values[i])
			} else if value.Valid {
				oi.IntellectID = value.Int64
			}
		case orderinfo.FieldStationID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field station_id", values[i])
			} else if value.Valid {
				oi.StationID = datasource.UUID(value.Int64)
			}
		case orderinfo.FieldOperatorID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field operator_id", values[i])
			} else if value.Valid {
				oi.OperatorID = datasource.UUID(value.Int64)
			}
		case orderinfo.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field connector_order_info", value)
			} else if value.Valid {
				oi.connector_order_info = new(int)
				*oi.connector_order_info = int(value.Int64)
			}
		case orderinfo.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field equipment_order_info", value)
			} else if value.Valid {
				oi.equipment_order_info = new(int)
				*oi.equipment_order_info = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryConnector queries the "connector" edge of the OrderInfo entity.
func (oi *OrderInfo) QueryConnector() *ConnectorQuery {
	return (&OrderInfoClient{config: oi.config}).QueryConnector(oi)
}

// QueryEquipment queries the "equipment" edge of the OrderInfo entity.
func (oi *OrderInfo) QueryEquipment() *EquipmentQuery {
	return (&OrderInfoClient{config: oi.config}).QueryEquipment(oi)
}

// QueryOrderEvent queries the "order_event" edge of the OrderInfo entity.
func (oi *OrderInfo) QueryOrderEvent() *OrderEventQuery {
	return (&OrderInfoClient{config: oi.config}).QueryOrderEvent(oi)
}

// Update returns a builder for updating this OrderInfo.
// Note that you need to call OrderInfo.Unwrap() before calling this method if this OrderInfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (oi *OrderInfo) Update() *OrderInfoUpdateOne {
	return (&OrderInfoClient{config: oi.config}).UpdateOne(oi)
}

// Unwrap unwraps the OrderInfo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oi *OrderInfo) Unwrap() *OrderInfo {
	tx, ok := oi.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrderInfo is not a transactional entity")
	}
	oi.config.driver = tx.drv
	return oi
}

// String implements the fmt.Stringer.
func (oi *OrderInfo) String() string {
	var builder strings.Builder
	builder.WriteString("OrderInfo(")
	builder.WriteString(fmt.Sprintf("id=%v", oi.ID))
	builder.WriteString(", remote_start_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.RemoteStartID))
	builder.WriteString(", transaction_id=")
	builder.WriteString(oi.TransactionID)
	builder.WriteString(", authorization_id=")
	builder.WriteString(oi.AuthorizationID)
	builder.WriteString(", customer_id=")
	builder.WriteString(oi.CustomerID)
	builder.WriteString(", caller_order_id=")
	builder.WriteString(oi.CallerOrderID)
	builder.WriteString(", total_electricity=")
	builder.WriteString(fmt.Sprintf("%v", oi.TotalElectricity))
	builder.WriteString(", charge_start_electricity=")
	builder.WriteString(fmt.Sprintf("%v", oi.ChargeStartElectricity))
	builder.WriteString(", charge_stop_electricity=")
	builder.WriteString(fmt.Sprintf("%v", oi.ChargeStopElectricity))
	builder.WriteString(", sharp_electricity=")
	builder.WriteString(fmt.Sprintf("%v", oi.SharpElectricity))
	builder.WriteString(", peak_electricity=")
	builder.WriteString(fmt.Sprintf("%v", oi.PeakElectricity))
	builder.WriteString(", flat_electricity=")
	builder.WriteString(fmt.Sprintf("%v", oi.FlatElectricity))
	builder.WriteString(", valley_electricity=")
	builder.WriteString(fmt.Sprintf("%v", oi.ValleyElectricity))
	builder.WriteString(", stop_reason_code=")
	builder.WriteString(fmt.Sprintf("%v", oi.StopReasonCode))
	builder.WriteString(", offline=")
	builder.WriteString(fmt.Sprintf("%v", oi.Offline))
	builder.WriteString(", price_scheme_release_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.PriceSchemeReleaseID))
	builder.WriteString(", order_start_time=")
	builder.WriteString(fmt.Sprintf("%v", oi.OrderStartTime))
	builder.WriteString(", order_final_time=")
	builder.WriteString(fmt.Sprintf("%v", oi.OrderFinalTime))
	builder.WriteString(", charge_start_time=")
	builder.WriteString(fmt.Sprintf("%v", oi.ChargeStartTime))
	builder.WriteString(", charge_final_time=")
	builder.WriteString(fmt.Sprintf("%v", oi.ChargeFinalTime))
	builder.WriteString(", intellect_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.IntellectID))
	builder.WriteString(", station_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.StationID))
	builder.WriteString(", operator_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.OperatorID))
	builder.WriteByte(')')
	return builder.String()
}

// OrderInfos is a parsable slice of OrderInfo.
type OrderInfos []*OrderInfo

func (oi OrderInfos) config(cfg config) {
	for _i := range oi {
		oi[_i].config = cfg
	}
}
