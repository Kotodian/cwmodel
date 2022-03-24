// Code generated by entc, DO NOT EDIT.

package connector

import (
	"fmt"

	"github.com/Kotodian/ent-practice/ent/enums"
	"github.com/Kotodian/gokit/datasource"
)

const (
	// Label holds the string label denoting the connector type in the database.
	Label = "connector"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
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
	EvseColumn = "evse_connectors"
	// EquipmentTable is the table that holds the equipment relation/edge.
	EquipmentTable = "base_connector"
	// EquipmentInverseTable is the table name for the Equipment entity.
	// It exists in this package in order to avoid circular dependency with the "equipment" package.
	EquipmentInverseTable = "base_equipment"
	// EquipmentColumn is the table column denoting the equipment relation/edge.
	EquipmentColumn = "equipment_connectors"
)

// Columns holds all SQL columns for connector fields.
var Columns = []string{
	FieldID,
	FieldEquipmentSn,
	FieldEvseSerial,
	FieldSerial,
	FieldCurrentState,
	FieldBeforeState,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "base_connector"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"equipment_connectors",
	"evse_connectors",
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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID datasource.UUID
)

// CurrentStateValidator is a validator for the "current_state" field enum values. It is called by the builders before save.
func CurrentStateValidator(cs enums.ConnectorState) error {
	switch cs.String() {
	case "unavailable", "available", "occcupied", "reserved", "faulted":
		return nil
	default:
		return fmt.Errorf("connector: invalid enum value for current_state field: %q", cs)
	}
}

// BeforeStateValidator is a validator for the "before_state" field enum values. It is called by the builders before save.
func BeforeStateValidator(bs enums.ConnectorState) error {
	switch bs.String() {
	case "unavailable", "available", "occcupied", "reserved", "faulted":
		return nil
	default:
		return fmt.Errorf("connector: invalid enum value for before_state field: %q", bs)
	}
}
