// Code generated by ent, DO NOT EDIT.

package smartchargingeffect

import (
	"github.com/Kotodian/gokit/datasource"
)

const (
	// Label holds the string label denoting the smartchargingeffect type in the database.
	Label = "smart_charging_effect"
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
	// FieldSmartID holds the string denoting the smart_id field in the database.
	FieldSmartID = "smart_id"
	// FieldPid holds the string denoting the pid field in the database.
	FieldPid = "pid"
	// FieldUnit holds the string denoting the unit field in the database.
	FieldUnit = "unit"
	// FieldEquipmentSn holds the string denoting the equipment_sn field in the database.
	FieldEquipmentSn = "equipment_sn"
	// FieldValidFrom holds the string denoting the valid_from field in the database.
	FieldValidFrom = "valid_from"
	// FieldValidTo holds the string denoting the valid_to field in the database.
	FieldValidTo = "valid_to"
	// FieldSpec holds the string denoting the spec field in the database.
	FieldSpec = "spec"
	// EdgeEquipment holds the string denoting the equipment edge name in mutations.
	EdgeEquipment = "equipment"
	// EdgeConnector holds the string denoting the connector edge name in mutations.
	EdgeConnector = "connector"
	// EdgeOrderInfo holds the string denoting the order_info edge name in mutations.
	EdgeOrderInfo = "order_info"
	// Table holds the table name of the smartchargingeffect in the database.
	Table = "smart_charging_effect"
	// EquipmentTable is the table that holds the equipment relation/edge.
	EquipmentTable = "smart_charging_effect"
	// EquipmentInverseTable is the table name for the Equipment entity.
	// It exists in this package in order to avoid circular dependency with the "equipment" package.
	EquipmentInverseTable = "base_equipment"
	// EquipmentColumn is the table column denoting the equipment relation/edge.
	EquipmentColumn = "equipment_id"
	// ConnectorTable is the table that holds the connector relation/edge.
	ConnectorTable = "smart_charging_effect"
	// ConnectorInverseTable is the table name for the Connector entity.
	// It exists in this package in order to avoid circular dependency with the "connector" package.
	ConnectorInverseTable = "base_connector"
	// ConnectorColumn is the table column denoting the connector relation/edge.
	ConnectorColumn = "connector_id"
	// OrderInfoTable is the table that holds the order_info relation/edge.
	OrderInfoTable = "smart_charging_effect"
	// OrderInfoInverseTable is the table name for the OrderInfo entity.
	// It exists in this package in order to avoid circular dependency with the "orderinfo" package.
	OrderInfoInverseTable = "order_info"
	// OrderInfoColumn is the table column denoting the order_info relation/edge.
	OrderInfoColumn = "order_id"
)

// Columns holds all SQL columns for smartchargingeffect fields.
var Columns = []string{
	FieldID,
	FieldVersion,
	FieldCreatedBy,
	FieldCreatedAt,
	FieldUpdatedBy,
	FieldUpdatedAt,
	FieldSmartID,
	FieldPid,
	FieldUnit,
	FieldEquipmentSn,
	FieldValidFrom,
	FieldValidTo,
	FieldSpec,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "smart_charging_effect"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"connector_id",
	"equipment_id",
	"order_id",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID datasource.UUID
)
