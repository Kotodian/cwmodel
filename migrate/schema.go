// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AppModuleInfoColumns holds the columns for the "app_module_info" table.
	AppModuleInfoColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "desc", Type: field.TypeString},
	}
	// AppModuleInfoTable holds the schema information for the "app_module_info" table.
	AppModuleInfoTable = &schema.Table{
		Name:       "app_module_info",
		Columns:    AppModuleInfoColumns,
		PrimaryKey: []*schema.Column{AppModuleInfoColumns[0]},
	}
	// BaseConnectorColumns holds the columns for the "base_connector" table.
	BaseConnectorColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "version", Type: field.TypeInt64, Default: 1},
		{Name: "created_by", Type: field.TypeUint64, Default: 1},
		{Name: "created_at", Type: field.TypeInt64, Default: 1666233041},
		{Name: "updated_by", Type: field.TypeUint64, Default: 1},
		{Name: "updated_at", Type: field.TypeInt64, Default: 1666233041},
		{Name: "equipment_sn", Type: field.TypeString},
		{Name: "evse_serial", Type: field.TypeString},
		{Name: "serial", Type: field.TypeString},
		{Name: "current_state", Type: field.TypeInt},
		{Name: "before_state", Type: field.TypeInt},
		{Name: "charging_state", Type: field.TypeInt, Nullable: true},
		{Name: "reservation_id", Type: field.TypeUint64, Nullable: true},
		{Name: "order_id", Type: field.TypeUint64, Nullable: true},
		{Name: "park_no", Type: field.TypeString, Default: "0001"},
		{Name: "equipment_id", Type: field.TypeUint64},
		{Name: "evse_id", Type: field.TypeUint64},
	}
	// BaseConnectorTable holds the schema information for the "base_connector" table.
	BaseConnectorTable = &schema.Table{
		Name:       "base_connector",
		Columns:    BaseConnectorColumns,
		PrimaryKey: []*schema.Column{BaseConnectorColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "base_connector_base_equipment_connector",
				Columns:    []*schema.Column{BaseConnectorColumns[15]},
				RefColumns: []*schema.Column{BaseEquipmentColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "base_connector_base_evse_connector",
				Columns:    []*schema.Column{BaseConnectorColumns[16]},
				RefColumns: []*schema.Column{BaseEvseColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// BaseEquipmentColumns holds the columns for the "base_equipment" table.
	BaseEquipmentColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "version", Type: field.TypeInt64, Default: 1},
		{Name: "created_by", Type: field.TypeUint64, Default: 1},
		{Name: "created_at", Type: field.TypeInt64, Default: 1666233041},
		{Name: "updated_by", Type: field.TypeUint64, Default: 1},
		{Name: "updated_at", Type: field.TypeInt64, Default: 1666233041},
		{Name: "sn", Type: field.TypeString},
		{Name: "operator_id", Type: field.TypeUint64},
		{Name: "station_id", Type: field.TypeUint64},
	}
	// BaseEquipmentTable holds the schema information for the "base_equipment" table.
	BaseEquipmentTable = &schema.Table{
		Name:       "base_equipment",
		Columns:    BaseEquipmentColumns,
		PrimaryKey: []*schema.Column{BaseEquipmentColumns[0]},
	}
	// BaseEquipmentAlarmColumns holds the columns for the "base_equipment_alarm" table.
	BaseEquipmentAlarmColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "version", Type: field.TypeInt64, Default: 1},
		{Name: "created_by", Type: field.TypeUint64, Default: 1},
		{Name: "created_at", Type: field.TypeInt64, Default: 1666233041},
		{Name: "updated_by", Type: field.TypeUint64, Default: 1},
		{Name: "updated_at", Type: field.TypeInt64, Default: 1666233041},
		{Name: "dtc_code", Type: field.TypeInt64},
		{Name: "remote_address", Type: field.TypeString},
		{Name: "trigger_time", Type: field.TypeInt64, Nullable: true},
		{Name: "final_time", Type: field.TypeInt64, Nullable: true},
		{Name: "count", Type: field.TypeInt, Default: 0},
		{Name: "equipment_id", Type: field.TypeUint64},
	}
	// BaseEquipmentAlarmTable holds the schema information for the "base_equipment_alarm" table.
	BaseEquipmentAlarmTable = &schema.Table{
		Name:       "base_equipment_alarm",
		Columns:    BaseEquipmentAlarmColumns,
		PrimaryKey: []*schema.Column{BaseEquipmentAlarmColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "base_equipment_alarm_base_equipment_equipment_alarm",
				Columns:    []*schema.Column{BaseEquipmentAlarmColumns[11]},
				RefColumns: []*schema.Column{BaseEquipmentColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// EquipFirmwareEffectColumns holds the columns for the "equip_firmware_effect" table.
	EquipFirmwareEffectColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "version", Type: field.TypeInt64, Default: 1},
		{Name: "created_by", Type: field.TypeUint64, Default: 1},
		{Name: "created_at", Type: field.TypeInt64, Default: 1666233041},
		{Name: "updated_by", Type: field.TypeUint64, Default: 1},
		{Name: "updated_at", Type: field.TypeInt64, Default: 1666233041},
		{Name: "request_id", Type: field.TypeInt64},
		{Name: "state", Type: field.TypeInt},
		{Name: "equipment_id", Type: field.TypeUint64, Nullable: true},
		{Name: "firmware_id", Type: field.TypeUint64, Nullable: true},
	}
	// EquipFirmwareEffectTable holds the schema information for the "equip_firmware_effect" table.
	EquipFirmwareEffectTable = &schema.Table{
		Name:       "equip_firmware_effect",
		Columns:    EquipFirmwareEffectColumns,
		PrimaryKey: []*schema.Column{EquipFirmwareEffectColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "equip_firmware_effect_base_equipment_equipment_firmware_effect",
				Columns:    []*schema.Column{EquipFirmwareEffectColumns[8]},
				RefColumns: []*schema.Column{BaseEquipmentColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "equip_firmware_effect_equip_firmware_template_equipment_firmware_effect",
				Columns:    []*schema.Column{EquipFirmwareEffectColumns[9]},
				RefColumns: []*schema.Column{EquipFirmwareTemplateColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BaseEquipmentExtraColumns holds the columns for the "base_equipment_extra" table.
	BaseEquipmentExtraColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "version", Type: field.TypeInt64, Default: 1},
		{Name: "created_by", Type: field.TypeUint64, Default: 1},
		{Name: "created_at", Type: field.TypeInt64, Default: 1666233041},
		{Name: "updated_by", Type: field.TypeUint64, Default: 1},
		{Name: "updated_at", Type: field.TypeInt64, Default: 1666233041},
		{Name: "equipment_sn", Type: field.TypeString},
		{Name: "model_id", Type: field.TypeUint64},
		{Name: "manufacturer_id", Type: field.TypeUint64},
		{Name: "firmware_id", Type: field.TypeUint64},
		{Name: "access_pod", Type: field.TypeString},
		{Name: "state", Type: field.TypeBool},
		{Name: "evse_number", Type: field.TypeUint},
		{Name: "alarm_number", Type: field.TypeUint},
		{Name: "register_datetime", Type: field.TypeInt64},
		{Name: "remote_address", Type: field.TypeString},
		{Name: "equipment_id", Type: field.TypeUint64, Unique: true},
	}
	// BaseEquipmentExtraTable holds the schema information for the "base_equipment_extra" table.
	BaseEquipmentExtraTable = &schema.Table{
		Name:       "base_equipment_extra",
		Columns:    BaseEquipmentExtraColumns,
		PrimaryKey: []*schema.Column{BaseEquipmentExtraColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "base_equipment_extra_base_equipment_equipment_info",
				Columns:    []*schema.Column{BaseEquipmentExtraColumns[16]},
				RefColumns: []*schema.Column{BaseEquipmentColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// EquipIotColumns holds the columns for the "equip_iot" table.
	EquipIotColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "version", Type: field.TypeInt64, Default: 1},
		{Name: "created_by", Type: field.TypeUint64, Default: 1},
		{Name: "created_at", Type: field.TypeInt64, Default: 1666233041},
		{Name: "updated_by", Type: field.TypeUint64, Default: 1},
		{Name: "updated_at", Type: field.TypeInt64, Default: 1666233041},
		{Name: "iccid", Type: field.TypeString, Nullable: true},
		{Name: "imei", Type: field.TypeString, Nullable: true},
		{Name: "remote_address", Type: field.TypeString, Nullable: true},
		{Name: "equipment_id", Type: field.TypeUint64, Unique: true, Nullable: true},
	}
	// EquipIotTable holds the schema information for the "equip_iot" table.
	EquipIotTable = &schema.Table{
		Name:       "equip_iot",
		Columns:    EquipIotColumns,
		PrimaryKey: []*schema.Column{EquipIotColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "equip_iot_base_equipment_equipment_iot",
				Columns:    []*schema.Column{EquipIotColumns[9]},
				RefColumns: []*schema.Column{BaseEquipmentColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// EquipLogsColumns holds the columns for the "equip_logs" table.
	EquipLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "version", Type: field.TypeInt64, Default: 1},
		{Name: "created_by", Type: field.TypeUint64, Default: 1},
		{Name: "created_at", Type: field.TypeInt64, Default: 1666233041},
		{Name: "updated_by", Type: field.TypeUint64, Default: 1},
		{Name: "updated_at", Type: field.TypeInt64, Default: 1666233041},
		{Name: "request_id", Type: field.TypeInt64},
		{Name: "state", Type: field.TypeInt},
		{Name: "data_link", Type: field.TypeUint64},
		{Name: "equipment_id", Type: field.TypeUint64, Nullable: true},
	}
	// EquipLogsTable holds the schema information for the "equip_logs" table.
	EquipLogsTable = &schema.Table{
		Name:       "equip_logs",
		Columns:    EquipLogsColumns,
		PrimaryKey: []*schema.Column{EquipLogsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "equip_logs_base_equipment_equipment_log",
				Columns:    []*schema.Column{EquipLogsColumns[9]},
				RefColumns: []*schema.Column{BaseEquipmentColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BaseEvseColumns holds the columns for the "base_evse" table.
	BaseEvseColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "version", Type: field.TypeInt64, Default: 1},
		{Name: "created_by", Type: field.TypeUint64, Default: 1},
		{Name: "created_at", Type: field.TypeInt64, Default: 1666233041},
		{Name: "updated_by", Type: field.TypeUint64, Default: 1},
		{Name: "updated_at", Type: field.TypeInt64, Default: 1666233041},
		{Name: "serial", Type: field.TypeString},
		{Name: "connector_number", Type: field.TypeInt},
		{Name: "equipment_id", Type: field.TypeUint64},
	}
	// BaseEvseTable holds the schema information for the "base_evse" table.
	BaseEvseTable = &schema.Table{
		Name:       "base_evse",
		Columns:    BaseEvseColumns,
		PrimaryKey: []*schema.Column{BaseEvseColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "base_evse_base_equipment_evse",
				Columns:    []*schema.Column{BaseEvseColumns[8]},
				RefColumns: []*schema.Column{BaseEquipmentColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// EquipFirmwareTemplateColumns holds the columns for the "equip_firmware_template" table.
	EquipFirmwareTemplateColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "version", Type: field.TypeInt64, Default: 1},
		{Name: "created_by", Type: field.TypeUint64, Default: 1},
		{Name: "created_at", Type: field.TypeInt64, Default: 1666233041},
		{Name: "updated_by", Type: field.TypeUint64, Default: 1},
		{Name: "updated_at", Type: field.TypeInt64, Default: 1666233041},
		{Name: "equip_version", Type: field.TypeString},
		{Name: "manufacturer_id", Type: field.TypeUint64},
		{Name: "model_id", Type: field.TypeUint64},
	}
	// EquipFirmwareTemplateTable holds the schema information for the "equip_firmware_template" table.
	EquipFirmwareTemplateTable = &schema.Table{
		Name:       "equip_firmware_template",
		Columns:    EquipFirmwareTemplateColumns,
		PrimaryKey: []*schema.Column{EquipFirmwareTemplateColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "equip_firmware_template_equip_manufacturer_firmware",
				Columns:    []*schema.Column{EquipFirmwareTemplateColumns[7]},
				RefColumns: []*schema.Column{EquipManufacturerColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "equip_firmware_template_equip_model_firmware",
				Columns:    []*schema.Column{EquipFirmwareTemplateColumns[8]},
				RefColumns: []*schema.Column{EquipModelColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// EquipManufacturerColumns holds the columns for the "equip_manufacturer" table.
	EquipManufacturerColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "version", Type: field.TypeInt64, Default: 1},
		{Name: "created_by", Type: field.TypeUint64, Default: 1},
		{Name: "created_at", Type: field.TypeInt64, Default: 1666233041},
		{Name: "updated_by", Type: field.TypeUint64, Default: 1},
		{Name: "updated_at", Type: field.TypeInt64, Default: 1666233041},
		{Name: "code", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Nullable: true},
	}
	// EquipManufacturerTable holds the schema information for the "equip_manufacturer" table.
	EquipManufacturerTable = &schema.Table{
		Name:       "equip_manufacturer",
		Columns:    EquipManufacturerColumns,
		PrimaryKey: []*schema.Column{EquipManufacturerColumns[0]},
	}
	// EquipModelColumns holds the columns for the "equip_model" table.
	EquipModelColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "version", Type: field.TypeInt64, Default: 1},
		{Name: "created_by", Type: field.TypeUint64, Default: 1},
		{Name: "created_at", Type: field.TypeInt64, Default: 1666233041},
		{Name: "updated_by", Type: field.TypeUint64, Default: 1},
		{Name: "updated_at", Type: field.TypeInt64, Default: 1666233041},
		{Name: "code", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "phase_category", Type: field.TypeString},
		{Name: "current_category", Type: field.TypeString},
	}
	// EquipModelTable holds the schema information for the "equip_model" table.
	EquipModelTable = &schema.Table{
		Name:       "equip_model",
		Columns:    EquipModelColumns,
		PrimaryKey: []*schema.Column{EquipModelColumns[0]},
	}
	// OrderEventColumns holds the columns for the "order_event" table.
	OrderEventColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "content", Type: field.TypeString},
		{Name: "occurrence", Type: field.TypeInt64},
		{Name: "order_id", Type: field.TypeUint64, Nullable: true},
	}
	// OrderEventTable holds the schema information for the "order_event" table.
	OrderEventTable = &schema.Table{
		Name:       "order_event",
		Columns:    OrderEventColumns,
		PrimaryKey: []*schema.Column{OrderEventColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "order_event_order_info_order_event",
				Columns:    []*schema.Column{OrderEventColumns[3]},
				RefColumns: []*schema.Column{OrderInfoColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OrderInfoColumns holds the columns for the "order_info" table.
	OrderInfoColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "version", Type: field.TypeInt64, Default: 1},
		{Name: "created_by", Type: field.TypeUint64, Default: 1},
		{Name: "created_at", Type: field.TypeInt64, Default: 1666233041},
		{Name: "updated_by", Type: field.TypeUint64, Default: 1},
		{Name: "updated_at", Type: field.TypeInt64, Default: 1666233041},
		{Name: "remote_start_id", Type: field.TypeInt64, Nullable: true},
		{Name: "transaction_id", Type: field.TypeString},
		{Name: "authorization_id", Type: field.TypeString, Nullable: true},
		{Name: "customer_id", Type: field.TypeString, Nullable: true},
		{Name: "caller_order_id", Type: field.TypeString, Nullable: true},
		{Name: "total_electricity", Type: field.TypeFloat64, Nullable: true},
		{Name: "charge_start_electricity", Type: field.TypeFloat64, Nullable: true},
		{Name: "charge_final_electricity", Type: field.TypeFloat64, Nullable: true},
		{Name: "sharp_electricity", Type: field.TypeFloat64, Nullable: true},
		{Name: "peak_electricity", Type: field.TypeFloat64, Nullable: true},
		{Name: "flat_electricity", Type: field.TypeFloat64, Nullable: true},
		{Name: "valley_electricity", Type: field.TypeFloat64, Nullable: true},
		{Name: "stop_reason_code", Type: field.TypeInt32, Nullable: true},
		{Name: "offline", Type: field.TypeBool},
		{Name: "price_scheme_release_id", Type: field.TypeInt64},
		{Name: "order_start_time", Type: field.TypeInt64, Nullable: true},
		{Name: "order_final_time", Type: field.TypeInt64, Nullable: true},
		{Name: "charge_start_time", Type: field.TypeInt64, Nullable: true},
		{Name: "charge_final_time", Type: field.TypeInt64, Nullable: true},
		{Name: "intellect_id", Type: field.TypeInt64, Nullable: true},
		{Name: "station_id", Type: field.TypeUint64, Nullable: true},
		{Name: "operator_id", Type: field.TypeUint64, Nullable: true},
		{Name: "connector_id", Type: field.TypeUint64, Nullable: true},
		{Name: "equipment_id", Type: field.TypeUint64, Nullable: true},
	}
	// OrderInfoTable holds the schema information for the "order_info" table.
	OrderInfoTable = &schema.Table{
		Name:       "order_info",
		Columns:    OrderInfoColumns,
		PrimaryKey: []*schema.Column{OrderInfoColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "order_info_base_connector_order_info",
				Columns:    []*schema.Column{OrderInfoColumns[28]},
				RefColumns: []*schema.Column{BaseConnectorColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "order_info_base_equipment_order_info",
				Columns:    []*schema.Column{OrderInfoColumns[29]},
				RefColumns: []*schema.Column{BaseEquipmentColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ReservationChargingReleaseColumns holds the columns for the "reservation_charging_release" table.
	ReservationChargingReleaseColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "version", Type: field.TypeInt64, Default: 1},
		{Name: "created_by", Type: field.TypeUint64, Default: 1},
		{Name: "created_at", Type: field.TypeInt64, Default: 1666233041},
		{Name: "updated_by", Type: field.TypeUint64, Default: 1},
		{Name: "updated_at", Type: field.TypeInt64, Default: 1666233041},
		{Name: "reservation_id", Type: field.TypeInt64},
		{Name: "authorization_mode", Type: field.TypeInt},
		{Name: "authorization_id", Type: field.TypeString},
		{Name: "additional", Type: field.TypeString, Nullable: true},
		{Name: "customer_id", Type: field.TypeString, Nullable: true},
		{Name: "expired", Type: field.TypeInt64},
		{Name: "state", Type: field.TypeInt},
		{Name: "connector_id", Type: field.TypeUint64, Nullable: true},
		{Name: "equipment_id", Type: field.TypeUint64, Nullable: true},
	}
	// ReservationChargingReleaseTable holds the schema information for the "reservation_charging_release" table.
	ReservationChargingReleaseTable = &schema.Table{
		Name:       "reservation_charging_release",
		Columns:    ReservationChargingReleaseColumns,
		PrimaryKey: []*schema.Column{ReservationChargingReleaseColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "reservation_charging_release_base_connector_reservation",
				Columns:    []*schema.Column{ReservationChargingReleaseColumns[13]},
				RefColumns: []*schema.Column{BaseConnectorColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "reservation_charging_release_base_equipment_reservation",
				Columns:    []*schema.Column{ReservationChargingReleaseColumns[14]},
				RefColumns: []*schema.Column{BaseEquipmentColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SmartChargingEffectColumns holds the columns for the "smart_charging_effect" table.
	SmartChargingEffectColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "version", Type: field.TypeInt64, Default: 1},
		{Name: "created_by", Type: field.TypeUint64, Default: 1},
		{Name: "created_at", Type: field.TypeInt64, Default: 1666233041},
		{Name: "updated_by", Type: field.TypeUint64, Default: 1},
		{Name: "updated_at", Type: field.TypeInt64, Default: 1666233041},
		{Name: "smart_id", Type: field.TypeInt64},
		{Name: "start_time", Type: field.TypeInt64},
		{Name: "pid", Type: field.TypeUint64},
		{Name: "unit", Type: field.TypeString},
		{Name: "equipment_sn", Type: field.TypeString},
		{Name: "valid_from", Type: field.TypeInt64, Nullable: true},
		{Name: "valid_to", Type: field.TypeInt64, Nullable: true},
		{Name: "spec", Type: field.TypeJSON},
		{Name: "connector_id", Type: field.TypeUint64},
		{Name: "equipment_id", Type: field.TypeUint64},
		{Name: "order_id", Type: field.TypeUint64, Unique: true, Nullable: true},
	}
	// SmartChargingEffectTable holds the schema information for the "smart_charging_effect" table.
	SmartChargingEffectTable = &schema.Table{
		Name:       "smart_charging_effect",
		Columns:    SmartChargingEffectColumns,
		PrimaryKey: []*schema.Column{SmartChargingEffectColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "smart_charging_effect_base_connector_smart_charging_effect",
				Columns:    []*schema.Column{SmartChargingEffectColumns[14]},
				RefColumns: []*schema.Column{BaseConnectorColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "smart_charging_effect_base_equipment_smart_charging_effect",
				Columns:    []*schema.Column{SmartChargingEffectColumns[15]},
				RefColumns: []*schema.Column{BaseEquipmentColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "smart_charging_effect_order_info_smart_charging_effect",
				Columns:    []*schema.Column{SmartChargingEffectColumns[16]},
				RefColumns: []*schema.Column{OrderInfoColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AppModuleInfoTable,
		BaseConnectorTable,
		BaseEquipmentTable,
		BaseEquipmentAlarmTable,
		EquipFirmwareEffectTable,
		BaseEquipmentExtraTable,
		EquipIotTable,
		EquipLogsTable,
		BaseEvseTable,
		EquipFirmwareTemplateTable,
		EquipManufacturerTable,
		EquipModelTable,
		OrderEventTable,
		OrderInfoTable,
		ReservationChargingReleaseTable,
		SmartChargingEffectTable,
	}
)

func init() {
	AppModuleInfoTable.Annotation = &entsql.Annotation{
		Table: "app_module_info",
	}
	BaseConnectorTable.ForeignKeys[0].RefTable = BaseEquipmentTable
	BaseConnectorTable.ForeignKeys[1].RefTable = BaseEvseTable
	BaseConnectorTable.Annotation = &entsql.Annotation{
		Table: "base_connector",
	}
	BaseEquipmentTable.Annotation = &entsql.Annotation{
		Table: "base_equipment",
	}
	BaseEquipmentAlarmTable.ForeignKeys[0].RefTable = BaseEquipmentTable
	BaseEquipmentAlarmTable.Annotation = &entsql.Annotation{
		Table: "base_equipment_alarm",
	}
	EquipFirmwareEffectTable.ForeignKeys[0].RefTable = BaseEquipmentTable
	EquipFirmwareEffectTable.ForeignKeys[1].RefTable = EquipFirmwareTemplateTable
	EquipFirmwareEffectTable.Annotation = &entsql.Annotation{
		Table: "equip_firmware_effect",
	}
	BaseEquipmentExtraTable.ForeignKeys[0].RefTable = BaseEquipmentTable
	BaseEquipmentExtraTable.Annotation = &entsql.Annotation{
		Table: "base_equipment_extra",
	}
	EquipIotTable.ForeignKeys[0].RefTable = BaseEquipmentTable
	EquipIotTable.Annotation = &entsql.Annotation{
		Table: "equip_iot",
	}
	EquipLogsTable.ForeignKeys[0].RefTable = BaseEquipmentTable
	EquipLogsTable.Annotation = &entsql.Annotation{
		Table: "equip_logs",
	}
	BaseEvseTable.ForeignKeys[0].RefTable = BaseEquipmentTable
	BaseEvseTable.Annotation = &entsql.Annotation{
		Table: "base_evse",
	}
	EquipFirmwareTemplateTable.ForeignKeys[0].RefTable = EquipManufacturerTable
	EquipFirmwareTemplateTable.ForeignKeys[1].RefTable = EquipModelTable
	EquipFirmwareTemplateTable.Annotation = &entsql.Annotation{
		Table: "equip_firmware_template",
	}
	EquipManufacturerTable.Annotation = &entsql.Annotation{
		Table: "equip_manufacturer",
	}
	EquipModelTable.Annotation = &entsql.Annotation{
		Table: "equip_model",
	}
	OrderEventTable.ForeignKeys[0].RefTable = OrderInfoTable
	OrderEventTable.Annotation = &entsql.Annotation{
		Table: "order_event",
	}
	OrderInfoTable.ForeignKeys[0].RefTable = BaseConnectorTable
	OrderInfoTable.ForeignKeys[1].RefTable = BaseEquipmentTable
	OrderInfoTable.Annotation = &entsql.Annotation{
		Table: "order_info",
	}
	ReservationChargingReleaseTable.ForeignKeys[0].RefTable = BaseConnectorTable
	ReservationChargingReleaseTable.ForeignKeys[1].RefTable = BaseEquipmentTable
	ReservationChargingReleaseTable.Annotation = &entsql.Annotation{
		Table: "reservation_charging_release",
	}
	SmartChargingEffectTable.ForeignKeys[0].RefTable = BaseConnectorTable
	SmartChargingEffectTable.ForeignKeys[1].RefTable = BaseEquipmentTable
	SmartChargingEffectTable.ForeignKeys[2].RefTable = OrderInfoTable
	SmartChargingEffectTable.Annotation = &entsql.Annotation{
		Table: "smart_charging_effect",
	}
}
