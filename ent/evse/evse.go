// Code generated by entc, DO NOT EDIT.

package evse

const (
	// Label holds the string label denoting the evse type in the database.
	Label = "evse"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
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
	EquipmentColumn = "equipment_evse"
	// ConnectorTable is the table that holds the connector relation/edge.
	ConnectorTable = "base_connector"
	// ConnectorInverseTable is the table name for the Connector entity.
	// It exists in this package in order to avoid circular dependency with the "connector" package.
	ConnectorInverseTable = "base_connector"
	// ConnectorColumn is the table column denoting the connector relation/edge.
	ConnectorColumn = "evse_connector"
)

// Columns holds all SQL columns for evse fields.
var Columns = []string{
	FieldID,
	FieldSerial,
	FieldConnectorNumber,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "base_evse"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"equipment_evse",
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
