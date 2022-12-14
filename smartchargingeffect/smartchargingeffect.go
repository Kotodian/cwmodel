// Code generated by ent, DO NOT EDIT.

package smartchargingeffect

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
	// FieldEquipmentID holds the string denoting the equipment_id field in the database.
	FieldEquipmentID = "equipment_id"
	// FieldConnectorID holds the string denoting the connector_id field in the database.
	FieldConnectorID = "connector_id"
	// FieldOrderID holds the string denoting the order_id field in the database.
	FieldOrderID = "order_id"
	// FieldSmartID holds the string denoting the smart_id field in the database.
	FieldSmartID = "smart_id"
	// FieldStartTime holds the string denoting the start_time field in the database.
	FieldStartTime = "start_time"
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
)

// Columns holds all SQL columns for smartchargingeffect fields.
var Columns = []string{
	FieldID,
	FieldVersion,
	FieldCreatedBy,
	FieldCreatedAt,
	FieldUpdatedBy,
	FieldUpdatedAt,
	FieldEquipmentID,
	FieldConnectorID,
	FieldOrderID,
	FieldSmartID,
	FieldStartTime,
	FieldPid,
	FieldUnit,
	FieldEquipmentSn,
	FieldValidFrom,
	FieldValidTo,
	FieldSpec,
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
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt int64
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() int64
)
