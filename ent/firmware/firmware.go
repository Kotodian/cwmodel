// Code generated by entc, DO NOT EDIT.

package firmware

const (
	// Label holds the string label denoting the firmware type in the database.
	Label = "firmware"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEquipVersion holds the string denoting the equip_version field in the database.
	FieldEquipVersion = "equip_version"
	// EdgeEquipment holds the string denoting the equipment edge name in mutations.
	EdgeEquipment = "equipment"
	// EdgeEquipmentFirmwareEffect holds the string denoting the equipment_firmware_effect edge name in mutations.
	EdgeEquipmentFirmwareEffect = "equipment_firmware_effect"
	// Table holds the table name of the firmware in the database.
	Table = "equip_firmware_template"
	// EquipmentTable is the table that holds the equipment relation/edge.
	EquipmentTable = "base_equipment"
	// EquipmentInverseTable is the table name for the Equipment entity.
	// It exists in this package in order to avoid circular dependency with the "equipment" package.
	EquipmentInverseTable = "base_equipment"
	// EquipmentColumn is the table column denoting the equipment relation/edge.
	EquipmentColumn = "firmware_equipment"
	// EquipmentFirmwareEffectTable is the table that holds the equipment_firmware_effect relation/edge.
	EquipmentFirmwareEffectTable = "equipment_firmware_effects"
	// EquipmentFirmwareEffectInverseTable is the table name for the EquipmentFirmwareEffect entity.
	// It exists in this package in order to avoid circular dependency with the "equipmentfirmwareeffect" package.
	EquipmentFirmwareEffectInverseTable = "equipment_firmware_effects"
	// EquipmentFirmwareEffectColumn is the table column denoting the equipment_firmware_effect relation/edge.
	EquipmentFirmwareEffectColumn = "firmware_equipment_firmware_effect"
)

// Columns holds all SQL columns for firmware fields.
var Columns = []string{
	FieldID,
	FieldEquipVersion,
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
