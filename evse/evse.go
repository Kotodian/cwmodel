// Code generated by ent, DO NOT EDIT.

package evse

import (
	"github.com/Kotodian/gokit/datasource"
)

const (
	// Label holds the string label denoting the evse type in the database.
	Label = "evse"
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
	// FieldSerial holds the string denoting the serial field in the database.
	FieldSerial = "serial"
	// FieldConnectorNumber holds the string denoting the connector_number field in the database.
	FieldConnectorNumber = "connector_number"
	// EdgeEquipment holds the string denoting the equipment edge name in mutations.
	EdgeEquipment = "equipment"
	// EdgeConnector holds the string denoting the connector edge name in mutations.
	EdgeConnector = "connector"
	// Table holds the table name of the evse in the database.
	Table = "base_evse"
	// EquipmentTable is the table that holds the equipment relation/edge.
	EquipmentTable = "base_evse"
	// EquipmentInverseTable is the table name for the Equipment entity.
	// It exists in this package in order to avoid circular dependency with the "equipment" package.
	EquipmentInverseTable = "base_equipment"
	// EquipmentColumn is the table column denoting the equipment relation/edge.
	EquipmentColumn = "equipment_id"
	// ConnectorTable is the table that holds the connector relation/edge.
	ConnectorTable = "base_connector"
	// ConnectorInverseTable is the table name for the Connector entity.
	// It exists in this package in order to avoid circular dependency with the "connector" package.
	ConnectorInverseTable = "base_connector"
	// ConnectorColumn is the table column denoting the connector relation/edge.
	ConnectorColumn = "evse_id"
)

// Columns holds all SQL columns for evse fields.
var Columns = []string{
	FieldID,
	FieldVersion,
	FieldCreatedBy,
	FieldCreatedAt,
	FieldUpdatedBy,
	FieldUpdatedAt,
	FieldSerial,
	FieldConnectorNumber,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "base_evse"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"equipment_id",
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