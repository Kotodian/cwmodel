// Code generated by ent, DO NOT EDIT.

package connector

import (
	"github.com/Kotodian/gokit/datasource"
)

const (
	// Label holds the string label denoting the connector type in the database.
	Label = "connector"
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
	// FieldEquipmentID holds the string denoting the equipment_id field in the database.
	FieldEquipmentID = "equipment_id"
	// FieldEvseID holds the string denoting the evse_id field in the database.
	FieldEvseID = "evse_id"
	// FieldEquipmentSn holds the string denoting the equipment_sn field in the database.
	FieldEquipmentSn = "equipment_sn"
	// FieldEvseSerial holds the string denoting the evse_serial field in the database.
	FieldEvseSerial = "evse_serial"
	// FieldSerial holds the string denoting the serial field in the database.
	FieldSerial = "serial"
	// FieldCurrentState holds the string denoting the current_state field in the database.
	FieldCurrentState = "current_state"
	// FieldBeforeState holds the string denoting the before_state field in the database.
	FieldBeforeState = "before_state"
	// FieldChargingState holds the string denoting the charging_state field in the database.
	FieldChargingState = "charging_state"
	// FieldReservationID holds the string denoting the reservation_id field in the database.
	FieldReservationID = "reservation_id"
	// FieldOrderID holds the string denoting the order_id field in the database.
	FieldOrderID = "order_id"
	// FieldParkNo holds the string denoting the park_no field in the database.
	FieldParkNo = "park_no"
	// EdgeEvse holds the string denoting the evse edge name in mutations.
	EdgeEvse = "evse"
	// EdgeEquipment holds the string denoting the equipment edge name in mutations.
	EdgeEquipment = "equipment"
	// Table holds the table name of the connector in the database.
	Table = "base_connector"
	// EvseTable is the table that holds the evse relation/edge.
	EvseTable = "base_connector"
	// EvseInverseTable is the table name for the Evse entity.
	// It exists in this package in order to avoid circular dependency with the "evse" package.
	EvseInverseTable = "base_evse"
	// EvseColumn is the table column denoting the evse relation/edge.
	EvseColumn = "evse_id"
	// EquipmentTable is the table that holds the equipment relation/edge.
	EquipmentTable = "base_connector"
	// EquipmentInverseTable is the table name for the Equipment entity.
	// It exists in this package in order to avoid circular dependency with the "equipment" package.
	EquipmentInverseTable = "base_equipment"
	// EquipmentColumn is the table column denoting the equipment relation/edge.
	EquipmentColumn = "equipment_id"
)

// Columns holds all SQL columns for connector fields.
var Columns = []string{
	FieldID,
	FieldVersion,
	FieldCreatedBy,
	FieldCreatedAt,
	FieldUpdatedBy,
	FieldUpdatedAt,
	FieldEquipmentID,
	FieldEvseID,
	FieldEquipmentSn,
	FieldEvseSerial,
	FieldSerial,
	FieldCurrentState,
	FieldBeforeState,
	FieldChargingState,
	FieldReservationID,
	FieldOrderID,
	FieldParkNo,
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
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() int64
	// DefaultParkNo holds the default value on creation for the "park_no" field.
	DefaultParkNo string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID datasource.UUID
)
