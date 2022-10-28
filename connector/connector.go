// Code generated by ent, DO NOT EDIT.

package connector

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
	// EdgeOrderInfo holds the string denoting the order_info edge name in mutations.
	EdgeOrderInfo = "order_info"
	// EdgeReservation holds the string denoting the reservation edge name in mutations.
	EdgeReservation = "reservation"
	// EdgeSmartChargingEffect holds the string denoting the smart_charging_effect edge name in mutations.
	EdgeSmartChargingEffect = "smart_charging_effect"
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
	// OrderInfoTable is the table that holds the order_info relation/edge.
	OrderInfoTable = "order_info"
	// OrderInfoInverseTable is the table name for the OrderInfo entity.
	// It exists in this package in order to avoid circular dependency with the "orderinfo" package.
	OrderInfoInverseTable = "order_info"
	// OrderInfoColumn is the table column denoting the order_info relation/edge.
	OrderInfoColumn = "connector_id"
	// ReservationTable is the table that holds the reservation relation/edge.
	ReservationTable = "reservation_charging_release"
	// ReservationInverseTable is the table name for the Reservation entity.
	// It exists in this package in order to avoid circular dependency with the "reservation" package.
	ReservationInverseTable = "reservation_charging_release"
	// ReservationColumn is the table column denoting the reservation relation/edge.
	ReservationColumn = "connector_id"
	// SmartChargingEffectTable is the table that holds the smart_charging_effect relation/edge.
	SmartChargingEffectTable = "smart_charging_effect"
	// SmartChargingEffectInverseTable is the table name for the SmartChargingEffect entity.
	// It exists in this package in order to avoid circular dependency with the "smartchargingeffect" package.
	SmartChargingEffectInverseTable = "smart_charging_effect"
	// SmartChargingEffectColumn is the table column denoting the smart_charging_effect relation/edge.
	SmartChargingEffectColumn = "connector_id"
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
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt int64
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() int64
	// DefaultParkNo holds the default value on creation for the "park_no" field.
	DefaultParkNo string
)
