// Code generated by ent, DO NOT EDIT.

package cwmodel

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Kotodian/cwmodel/connector"
	"github.com/Kotodian/cwmodel/equipment"
	"github.com/Kotodian/cwmodel/orderinfo"
	"github.com/Kotodian/gokit/datasource"
)

// OrderInfo is the model entity for the OrderInfo schema.
type OrderInfo struct {
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
	// 远程启动id
	RemoteStartID *int64 `json:"remote_start_id,omitempty"`
	// 桩端订单id
	TransactionID *string `json:"transaction_id,omitempty"`
	// 授权id
	AuthorizationID *string `json:"authorization_id,omitempty"`
	// 授权模式
	AuthorizationMode *int `json:"authorization_mode,omitempty"`
	// 客户id
	CustomerID *string `json:"customer_id,omitempty"`
	// 第三方订单id
	CallerOrderID *string `json:"caller_order_id,omitempty"`
	// 总充电电量
	TotalElectricity *float64 `json:"total_electricity,omitempty"`
	// 起始电量
	ChargeStartElectricity *float64 `json:"charge_start_electricity,omitempty"`
	// 结束电量
	ChargeFinalElectricity *float64 `json:"charge_final_electricity,omitempty"`
	// 尖电量
	SharpElectricity *float64 `json:"sharp_electricity,omitempty"`
	// 峰电量
	PeakElectricity *float64 `json:"peak_electricity,omitempty"`
	// 平电量
	FlatElectricity *float64 `json:"flat_electricity,omitempty"`
	// 谷电量
	ValleyElectricity *float64 `json:"valley_electricity,omitempty"`
	// 停止原因代码
	StopReasonCode *int32 `json:"stop_reason_code,omitempty"`
	// 车辆VIN码
	VIN *string `json:"VIN,omitempty"`
	// 订单状态
	State int `json:"state,omitempty"`
	// 是否为离线订单
	Offline bool `json:"offline,omitempty"`
	// 计费模板id
	PriceSchemeReleaseID *int64 `json:"price_scheme_release_id,omitempty"`
	// 订单开始时间
	OrderStartTime *int64 `json:"order_start_time,omitempty"`
	// 订单结束时间
	OrderFinalTime *int64 `json:"order_final_time,omitempty"`
	// 充电开始时间
	ChargeStartTime *int64 `json:"charge_start_time,omitempty"`
	// 充电结束时间
	ChargeFinalTime *int64 `json:"charge_final_time,omitempty"`
	// 智慧充电id
	IntellectID *int64 `json:"intellect_id,omitempty"`
	// 场站id
	StationID *datasource.UUID `json:"station_id,omitempty"`
	// 运营商id
	OperatorID *datasource.UUID `json:"operator_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderInfoQuery when eager-loading is set.
	Edges OrderInfoEdges `json:"-"`
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
			// Edge was loaded but was not found.
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
			// Edge was loaded but was not found.
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
func (*OrderInfo) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case orderinfo.FieldOffline:
			values[i] = new(sql.NullBool)
		case orderinfo.FieldTotalElectricity, orderinfo.FieldChargeStartElectricity, orderinfo.FieldChargeFinalElectricity, orderinfo.FieldSharpElectricity, orderinfo.FieldPeakElectricity, orderinfo.FieldFlatElectricity, orderinfo.FieldValleyElectricity:
			values[i] = new(sql.NullFloat64)
		case orderinfo.FieldID, orderinfo.FieldVersion, orderinfo.FieldCreatedBy, orderinfo.FieldCreatedAt, orderinfo.FieldUpdatedBy, orderinfo.FieldUpdatedAt, orderinfo.FieldEquipmentID, orderinfo.FieldConnectorID, orderinfo.FieldRemoteStartID, orderinfo.FieldAuthorizationMode, orderinfo.FieldStopReasonCode, orderinfo.FieldState, orderinfo.FieldPriceSchemeReleaseID, orderinfo.FieldOrderStartTime, orderinfo.FieldOrderFinalTime, orderinfo.FieldChargeStartTime, orderinfo.FieldChargeFinalTime, orderinfo.FieldIntellectID, orderinfo.FieldStationID, orderinfo.FieldOperatorID:
			values[i] = new(sql.NullInt64)
		case orderinfo.FieldTransactionID, orderinfo.FieldAuthorizationID, orderinfo.FieldCustomerID, orderinfo.FieldCallerOrderID, orderinfo.FieldVIN:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OrderInfo", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrderInfo fields.
func (oi *OrderInfo) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orderinfo.FieldID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				oi.ID = datasource.UUID(value.Int64)
			}
		case orderinfo.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				oi.Version = value.Int64
			}
		case orderinfo.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				oi.CreatedBy = datasource.UUID(value.Int64)
			}
		case orderinfo.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				oi.CreatedAt = value.Int64
			}
		case orderinfo.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				oi.UpdatedBy = datasource.UUID(value.Int64)
			}
		case orderinfo.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				oi.UpdatedAt = value.Int64
			}
		case orderinfo.FieldEquipmentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field equipment_id", values[i])
			} else if value.Valid {
				oi.EquipmentID = datasource.UUID(value.Int64)
			}
		case orderinfo.FieldConnectorID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field connector_id", values[i])
			} else if value.Valid {
				oi.ConnectorID = datasource.UUID(value.Int64)
			}
		case orderinfo.FieldRemoteStartID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field remote_start_id", values[i])
			} else if value.Valid {
				oi.RemoteStartID = new(int64)
				*oi.RemoteStartID = value.Int64
			}
		case orderinfo.FieldTransactionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field transaction_id", values[i])
			} else if value.Valid {
				oi.TransactionID = new(string)
				*oi.TransactionID = value.String
			}
		case orderinfo.FieldAuthorizationID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field authorization_id", values[i])
			} else if value.Valid {
				oi.AuthorizationID = new(string)
				*oi.AuthorizationID = value.String
			}
		case orderinfo.FieldAuthorizationMode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field authorization_mode", values[i])
			} else if value.Valid {
				oi.AuthorizationMode = new(int)
				*oi.AuthorizationMode = int(value.Int64)
			}
		case orderinfo.FieldCustomerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field customer_id", values[i])
			} else if value.Valid {
				oi.CustomerID = new(string)
				*oi.CustomerID = value.String
			}
		case orderinfo.FieldCallerOrderID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field caller_order_id", values[i])
			} else if value.Valid {
				oi.CallerOrderID = new(string)
				*oi.CallerOrderID = value.String
			}
		case orderinfo.FieldTotalElectricity:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field total_electricity", values[i])
			} else if value.Valid {
				oi.TotalElectricity = new(float64)
				*oi.TotalElectricity = value.Float64
			}
		case orderinfo.FieldChargeStartElectricity:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field charge_start_electricity", values[i])
			} else if value.Valid {
				oi.ChargeStartElectricity = new(float64)
				*oi.ChargeStartElectricity = value.Float64
			}
		case orderinfo.FieldChargeFinalElectricity:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field charge_final_electricity", values[i])
			} else if value.Valid {
				oi.ChargeFinalElectricity = new(float64)
				*oi.ChargeFinalElectricity = value.Float64
			}
		case orderinfo.FieldSharpElectricity:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field sharp_electricity", values[i])
			} else if value.Valid {
				oi.SharpElectricity = new(float64)
				*oi.SharpElectricity = value.Float64
			}
		case orderinfo.FieldPeakElectricity:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field peak_electricity", values[i])
			} else if value.Valid {
				oi.PeakElectricity = new(float64)
				*oi.PeakElectricity = value.Float64
			}
		case orderinfo.FieldFlatElectricity:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field flat_electricity", values[i])
			} else if value.Valid {
				oi.FlatElectricity = new(float64)
				*oi.FlatElectricity = value.Float64
			}
		case orderinfo.FieldValleyElectricity:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field valley_electricity", values[i])
			} else if value.Valid {
				oi.ValleyElectricity = new(float64)
				*oi.ValleyElectricity = value.Float64
			}
		case orderinfo.FieldStopReasonCode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field stop_reason_code", values[i])
			} else if value.Valid {
				oi.StopReasonCode = new(int32)
				*oi.StopReasonCode = int32(value.Int64)
			}
		case orderinfo.FieldVIN:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field VIN", values[i])
			} else if value.Valid {
				oi.VIN = new(string)
				*oi.VIN = value.String
			}
		case orderinfo.FieldState:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				oi.State = int(value.Int64)
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
				oi.PriceSchemeReleaseID = new(int64)
				*oi.PriceSchemeReleaseID = value.Int64
			}
		case orderinfo.FieldOrderStartTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_start_time", values[i])
			} else if value.Valid {
				oi.OrderStartTime = new(int64)
				*oi.OrderStartTime = value.Int64
			}
		case orderinfo.FieldOrderFinalTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_final_time", values[i])
			} else if value.Valid {
				oi.OrderFinalTime = new(int64)
				*oi.OrderFinalTime = value.Int64
			}
		case orderinfo.FieldChargeStartTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field charge_start_time", values[i])
			} else if value.Valid {
				oi.ChargeStartTime = new(int64)
				*oi.ChargeStartTime = value.Int64
			}
		case orderinfo.FieldChargeFinalTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field charge_final_time", values[i])
			} else if value.Valid {
				oi.ChargeFinalTime = new(int64)
				*oi.ChargeFinalTime = value.Int64
			}
		case orderinfo.FieldIntellectID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field intellect_id", values[i])
			} else if value.Valid {
				oi.IntellectID = new(int64)
				*oi.IntellectID = value.Int64
			}
		case orderinfo.FieldStationID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field station_id", values[i])
			} else if value.Valid {
				oi.StationID = new(datasource.UUID)
				*oi.StationID = datasource.UUID(value.Int64)
			}
		case orderinfo.FieldOperatorID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field operator_id", values[i])
			} else if value.Valid {
				oi.OperatorID = new(datasource.UUID)
				*oi.OperatorID = datasource.UUID(value.Int64)
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
	_tx, ok := oi.config.driver.(*txDriver)
	if !ok {
		panic("cwmodel: OrderInfo is not a transactional entity")
	}
	oi.config.driver = _tx.drv
	return oi
}

// String implements the fmt.Stringer.
func (oi *OrderInfo) String() string {
	var builder strings.Builder
	builder.WriteString("OrderInfo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", oi.ID))
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", oi.Version))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", oi.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", oi.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", oi.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", oi.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("equipment_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.EquipmentID))
	builder.WriteString(", ")
	builder.WriteString("connector_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.ConnectorID))
	builder.WriteString(", ")
	if v := oi.RemoteStartID; v != nil {
		builder.WriteString("remote_start_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := oi.TransactionID; v != nil {
		builder.WriteString("transaction_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := oi.AuthorizationID; v != nil {
		builder.WriteString("authorization_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := oi.AuthorizationMode; v != nil {
		builder.WriteString("authorization_mode=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := oi.CustomerID; v != nil {
		builder.WriteString("customer_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := oi.CallerOrderID; v != nil {
		builder.WriteString("caller_order_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := oi.TotalElectricity; v != nil {
		builder.WriteString("total_electricity=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := oi.ChargeStartElectricity; v != nil {
		builder.WriteString("charge_start_electricity=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := oi.ChargeFinalElectricity; v != nil {
		builder.WriteString("charge_final_electricity=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := oi.SharpElectricity; v != nil {
		builder.WriteString("sharp_electricity=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := oi.PeakElectricity; v != nil {
		builder.WriteString("peak_electricity=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := oi.FlatElectricity; v != nil {
		builder.WriteString("flat_electricity=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := oi.ValleyElectricity; v != nil {
		builder.WriteString("valley_electricity=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := oi.StopReasonCode; v != nil {
		builder.WriteString("stop_reason_code=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := oi.VIN; v != nil {
		builder.WriteString("VIN=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(fmt.Sprintf("%v", oi.State))
	builder.WriteString(", ")
	builder.WriteString("offline=")
	builder.WriteString(fmt.Sprintf("%v", oi.Offline))
	builder.WriteString(", ")
	if v := oi.PriceSchemeReleaseID; v != nil {
		builder.WriteString("price_scheme_release_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := oi.OrderStartTime; v != nil {
		builder.WriteString("order_start_time=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := oi.OrderFinalTime; v != nil {
		builder.WriteString("order_final_time=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := oi.ChargeStartTime; v != nil {
		builder.WriteString("charge_start_time=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := oi.ChargeFinalTime; v != nil {
		builder.WriteString("charge_final_time=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := oi.IntellectID; v != nil {
		builder.WriteString("intellect_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := oi.StationID; v != nil {
		builder.WriteString("station_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := oi.OperatorID; v != nil {
		builder.WriteString("operator_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
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
