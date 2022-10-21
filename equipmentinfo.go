// Code generated by ent, DO NOT EDIT.

package cwmodel

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Kotodian/cwmodel/equipment"
	"github.com/Kotodian/cwmodel/equipmentinfo"
	"github.com/Kotodian/gokit/datasource"
)

// EquipmentInfo is the model entity for the EquipmentInfo schema.
type EquipmentInfo struct {
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
	// 桩id
	EquipmentID datasource.UUID `json:"equipmentId"`
	// 桩序列号
	EquipmentSn string `json:"equipmentSN"`
	// 型号id
	ModelID datasource.UUID `json:"modelId"`
	// 产商id
	ManufacturerID datasource.UUID `json:"manufacturerId"`
	// 固件id
	FirmwareID datasource.UUID `json:"firmwareId"`
	// 所连接的pod
	AccessPod string `json:"accessPod"`
	// 在线或者离线
	State bool `json:"state"`
	// evse数量
	EvseNumber uint `json:"evseNumber"`
	// 告警数量
	AlarmNumber uint `json:"alarmNums"`
	// 注册时间
	RegisterDatetime int64 `json:"registerDatetime"`
	// 远程ip地址
	RemoteAddress string `json:"remoteAddress"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EquipmentInfoQuery when eager-loading is set.
	Edges EquipmentInfoEdges `json:"-"`
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
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: equipment.Label}
		}
		return e.Equipment, nil
	}
	return nil, &NotLoadedError{edge: "equipment"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EquipmentInfo) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case equipmentinfo.FieldState:
			values[i] = new(sql.NullBool)
		case equipmentinfo.FieldID, equipmentinfo.FieldVersion, equipmentinfo.FieldCreatedBy, equipmentinfo.FieldCreatedAt, equipmentinfo.FieldUpdatedBy, equipmentinfo.FieldUpdatedAt, equipmentinfo.FieldEquipmentID, equipmentinfo.FieldModelID, equipmentinfo.FieldManufacturerID, equipmentinfo.FieldFirmwareID, equipmentinfo.FieldEvseNumber, equipmentinfo.FieldAlarmNumber, equipmentinfo.FieldRegisterDatetime:
			values[i] = new(sql.NullInt64)
		case equipmentinfo.FieldEquipmentSn, equipmentinfo.FieldAccessPod, equipmentinfo.FieldRemoteAddress:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type EquipmentInfo", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EquipmentInfo fields.
func (ei *EquipmentInfo) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case equipmentinfo.FieldID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				ei.ID = datasource.UUID(value.Int64)
			}
		case equipmentinfo.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				ei.Version = value.Int64
			}
		case equipmentinfo.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				ei.CreatedBy = datasource.UUID(value.Int64)
			}
		case equipmentinfo.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ei.CreatedAt = value.Int64
			}
		case equipmentinfo.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				ei.UpdatedBy = datasource.UUID(value.Int64)
			}
		case equipmentinfo.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ei.UpdatedAt = value.Int64
			}
		case equipmentinfo.FieldEquipmentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field equipment_id", values[i])
			} else if value.Valid {
				ei.EquipmentID = datasource.UUID(value.Int64)
			}
		case equipmentinfo.FieldEquipmentSn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field equipment_sn", values[i])
			} else if value.Valid {
				ei.EquipmentSn = value.String
			}
		case equipmentinfo.FieldModelID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field model_id", values[i])
			} else if value.Valid {
				ei.ModelID = datasource.UUID(value.Int64)
			}
		case equipmentinfo.FieldManufacturerID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field manufacturer_id", values[i])
			} else if value.Valid {
				ei.ManufacturerID = datasource.UUID(value.Int64)
			}
		case equipmentinfo.FieldFirmwareID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field firmware_id", values[i])
			} else if value.Valid {
				ei.FirmwareID = datasource.UUID(value.Int64)
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
		case equipmentinfo.FieldEvseNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field evse_number", values[i])
			} else if value.Valid {
				ei.EvseNumber = uint(value.Int64)
			}
		case equipmentinfo.FieldAlarmNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field alarm_number", values[i])
			} else if value.Valid {
				ei.AlarmNumber = uint(value.Int64)
			}
		case equipmentinfo.FieldRegisterDatetime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field register_datetime", values[i])
			} else if value.Valid {
				ei.RegisterDatetime = value.Int64
			}
		case equipmentinfo.FieldRemoteAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remote_address", values[i])
			} else if value.Valid {
				ei.RemoteAddress = value.String
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
	_tx, ok := ei.config.driver.(*txDriver)
	if !ok {
		panic("cwmodel: EquipmentInfo is not a transactional entity")
	}
	ei.config.driver = _tx.drv
	return ei
}

// String implements the fmt.Stringer.
func (ei *EquipmentInfo) String() string {
	var builder strings.Builder
	builder.WriteString("EquipmentInfo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ei.ID))
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", ei.Version))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", ei.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", ei.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", ei.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", ei.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("equipment_id=")
	builder.WriteString(fmt.Sprintf("%v", ei.EquipmentID))
	builder.WriteString(", ")
	builder.WriteString("equipment_sn=")
	builder.WriteString(ei.EquipmentSn)
	builder.WriteString(", ")
	builder.WriteString("model_id=")
	builder.WriteString(fmt.Sprintf("%v", ei.ModelID))
	builder.WriteString(", ")
	builder.WriteString("manufacturer_id=")
	builder.WriteString(fmt.Sprintf("%v", ei.ManufacturerID))
	builder.WriteString(", ")
	builder.WriteString("firmware_id=")
	builder.WriteString(fmt.Sprintf("%v", ei.FirmwareID))
	builder.WriteString(", ")
	builder.WriteString("access_pod=")
	builder.WriteString(ei.AccessPod)
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(fmt.Sprintf("%v", ei.State))
	builder.WriteString(", ")
	builder.WriteString("evse_number=")
	builder.WriteString(fmt.Sprintf("%v", ei.EvseNumber))
	builder.WriteString(", ")
	builder.WriteString("alarm_number=")
	builder.WriteString(fmt.Sprintf("%v", ei.AlarmNumber))
	builder.WriteString(", ")
	builder.WriteString("register_datetime=")
	builder.WriteString(fmt.Sprintf("%v", ei.RegisterDatetime))
	builder.WriteString(", ")
	builder.WriteString("remote_address=")
	builder.WriteString(ei.RemoteAddress)
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
