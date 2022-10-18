// Code generated by ent, DO NOT EDIT.

package equipment

import (
	"github.com/Kotodian/gokit/datasource"
)

const (
	// Label holds the string label denoting the equipment type in the database.
	Label = "equipment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldSn holds the string denoting the sn field in the database.
	FieldSn = "sn"
	// FieldOperatorID holds the string denoting the operator_id field in the database.
	FieldOperatorID = "operator_id"
	// FieldStationID holds the string denoting the station_id field in the database.
	FieldStationID = "station_id"
	// EdgeEquipmentInfo holds the string denoting the equipment_info edge name in mutations.
	EdgeEquipmentInfo = "equipment_info"
	// EdgeEvse holds the string denoting the evse edge name in mutations.
	EdgeEvse = "evse"
	// EdgeConnector holds the string denoting the connector edge name in mutations.
	EdgeConnector = "connector"
	// EdgeEquipmentAlarm holds the string denoting the equipment_alarm edge name in mutations.
	EdgeEquipmentAlarm = "equipment_alarm"
	// EdgeEquipmentIot holds the string denoting the equipment_iot edge name in mutations.
	EdgeEquipmentIot = "equipment_iot"
	// EdgeEquipmentFirmwareEffect holds the string denoting the equipment_firmware_effect edge name in mutations.
	EdgeEquipmentFirmwareEffect = "equipment_firmware_effect"
	// EdgeOrderInfo holds the string denoting the order_info edge name in mutations.
	EdgeOrderInfo = "order_info"
	// EdgeReservation holds the string denoting the reservation edge name in mutations.
	EdgeReservation = "reservation"
	// EdgeEquipmentLog holds the string denoting the equipment_log edge name in mutations.
	EdgeEquipmentLog = "equipment_log"
	// Table holds the table name of the equipment in the database.
	Table = "base_equipment"
	// EquipmentInfoTable is the table that holds the equipment_info relation/edge.
	EquipmentInfoTable = "base_equipment_extra"
	// EquipmentInfoInverseTable is the table name for the EquipmentInfo entity.
	// It exists in this package in order to avoid circular dependency with the "equipmentinfo" package.
	EquipmentInfoInverseTable = "base_equipment_extra"
	// EquipmentInfoColumn is the table column denoting the equipment_info relation/edge.
	EquipmentInfoColumn = "equipment_id"
	// EvseTable is the table that holds the evse relation/edge.
	EvseTable = "base_evse"
	// EvseInverseTable is the table name for the Evse entity.
	// It exists in this package in order to avoid circular dependency with the "evse" package.
	EvseInverseTable = "base_evse"
	// EvseColumn is the table column denoting the evse relation/edge.
	EvseColumn = "equipment_id"
	// ConnectorTable is the table that holds the connector relation/edge.
	ConnectorTable = "base_connector"
	// ConnectorInverseTable is the table name for the Connector entity.
	// It exists in this package in order to avoid circular dependency with the "connector" package.
	ConnectorInverseTable = "base_connector"
	// ConnectorColumn is the table column denoting the connector relation/edge.
	ConnectorColumn = "equipment_id"
	// EquipmentAlarmTable is the table that holds the equipment_alarm relation/edge.
	EquipmentAlarmTable = "base_equipment_alarm"
	// EquipmentAlarmInverseTable is the table name for the EquipmentAlarm entity.
	// It exists in this package in order to avoid circular dependency with the "equipmentalarm" package.
	EquipmentAlarmInverseTable = "base_equipment_alarm"
	// EquipmentAlarmColumn is the table column denoting the equipment_alarm relation/edge.
	EquipmentAlarmColumn = "equipment_id"
	// EquipmentIotTable is the table that holds the equipment_iot relation/edge.
	EquipmentIotTable = "equip_iot"
	// EquipmentIotInverseTable is the table name for the EquipmentIot entity.
	// It exists in this package in order to avoid circular dependency with the "equipmentiot" package.
	EquipmentIotInverseTable = "equip_iot"
	// EquipmentIotColumn is the table column denoting the equipment_iot relation/edge.
	EquipmentIotColumn = "equipment_id"
	// EquipmentFirmwareEffectTable is the table that holds the equipment_firmware_effect relation/edge.
	EquipmentFirmwareEffectTable = "equip_firmware_effect"
	// EquipmentFirmwareEffectInverseTable is the table name for the EquipmentFirmwareEffect entity.
	// It exists in this package in order to avoid circular dependency with the "equipmentfirmwareeffect" package.
	EquipmentFirmwareEffectInverseTable = "equip_firmware_effect"
	// EquipmentFirmwareEffectColumn is the table column denoting the equipment_firmware_effect relation/edge.
	EquipmentFirmwareEffectColumn = "equipment_id"
	// OrderInfoTable is the table that holds the order_info relation/edge.
	OrderInfoTable = "order_info"
	// OrderInfoInverseTable is the table name for the OrderInfo entity.
	// It exists in this package in order to avoid circular dependency with the "orderinfo" package.
	OrderInfoInverseTable = "order_info"
	// OrderInfoColumn is the table column denoting the order_info relation/edge.
	OrderInfoColumn = "equipment_id"
	// ReservationTable is the table that holds the reservation relation/edge.
	ReservationTable = "reservation_charging_release"
	// ReservationInverseTable is the table name for the Reservation entity.
	// It exists in this package in order to avoid circular dependency with the "reservation" package.
	ReservationInverseTable = "reservation_charging_release"
	// ReservationColumn is the table column denoting the reservation relation/edge.
	ReservationColumn = "equipment_id"
	// EquipmentLogTable is the table that holds the equipment_log relation/edge.
	EquipmentLogTable = "equip_logs"
	// EquipmentLogInverseTable is the table name for the EquipmentLog entity.
	// It exists in this package in order to avoid circular dependency with the "equipmentlog" package.
	EquipmentLogInverseTable = "equip_logs"
	// EquipmentLogColumn is the table column denoting the equipment_log relation/edge.
	EquipmentLogColumn = "equipment_id"
)

// Columns holds all SQL columns for equipment fields.
var Columns = []string{
	FieldID,
	FieldVersion,
	FieldCreatedBy,
	FieldCreatedAt,
	FieldUpdatedBy,
	FieldUpdatedAt,
	FieldSn,
	FieldOperatorID,
	FieldStationID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultVersion holds the default value on creation for the "version" field.
	DefaultVersion int64
	// DefaultCreatedBy holds the default value on creation for the "created_by" field.
	DefaultCreatedBy datasource.UUID
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt int64
	// DefaultUpdatedBy holds the default value on creation for the "updated_by" field.
	DefaultUpdatedBy datasource.UUID
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt int64
	// SnValidator is a validator for the "sn" field. It is called by the builders before save.
	SnValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID datasource.UUID
)
