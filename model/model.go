// Code generated by ent, DO NOT EDIT.

package model

const (
	// Label holds the string label denoting the model type in the database.
	Label = "model"
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
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldConnectorNumber holds the string denoting the connector_number field in the database.
	FieldConnectorNumber = "connector_number"
	// FieldPhaseCategory holds the string denoting the phase_category field in the database.
	FieldPhaseCategory = "phase_category"
	// FieldCurrentCategory holds the string denoting the current_category field in the database.
	FieldCurrentCategory = "current_category"
	// FieldConnectorCategory holds the string denoting the connector_category field in the database.
	FieldConnectorCategory = "connector_category"
	// EdgeFirmware holds the string denoting the firmware edge name in mutations.
	EdgeFirmware = "firmware"
	// Table holds the table name of the model in the database.
	Table = "equip_model"
	// FirmwareTable is the table that holds the firmware relation/edge.
	FirmwareTable = "equip_firmware_template"
	// FirmwareInverseTable is the table name for the Firmware entity.
	// It exists in this package in order to avoid circular dependency with the "firmware" package.
	FirmwareInverseTable = "equip_firmware_template"
	// FirmwareColumn is the table column denoting the firmware relation/edge.
	FirmwareColumn = "model_id"
)

// Columns holds all SQL columns for model fields.
var Columns = []string{
	FieldID,
	FieldVersion,
	FieldCreatedBy,
	FieldCreatedAt,
	FieldUpdatedBy,
	FieldUpdatedAt,
	FieldCode,
	FieldName,
	FieldConnectorNumber,
	FieldPhaseCategory,
	FieldCurrentCategory,
	FieldConnectorCategory,
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
