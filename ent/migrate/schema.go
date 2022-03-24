// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BaseConnectorColumns holds the columns for the "base_connector" table.
	BaseConnectorColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "equipment_sn", Type: field.TypeString},
		{Name: "evse_serial", Type: field.TypeString},
		{Name: "serial", Type: field.TypeString},
		{Name: "current_state", Type: field.TypeEnum, Enums: []string{"unavailable", "available", "occcupied", "reserved", "faulted"}},
		{Name: "before_state", Type: field.TypeEnum, Enums: []string{"unavailable", "available", "occcupied", "reserved", "faulted"}},
		{Name: "equipment_connectors", Type: field.TypeUint64},
		{Name: "evse_connectors", Type: field.TypeUint64},
	}
	// BaseConnectorTable holds the schema information for the "base_connector" table.
	BaseConnectorTable = &schema.Table{
		Name:       "base_connector",
		Columns:    BaseConnectorColumns,
		PrimaryKey: []*schema.Column{BaseConnectorColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "base_connector_base_equipment_connectors",
				Columns:    []*schema.Column{BaseConnectorColumns[6]},
				RefColumns: []*schema.Column{BaseEquipmentColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "base_connector_base_evse_connectors",
				Columns:    []*schema.Column{BaseConnectorColumns[7]},
				RefColumns: []*schema.Column{BaseEvseColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "connector_serial_evse_connectors",
				Unique:  true,
				Columns: []*schema.Column{BaseConnectorColumns[3], BaseConnectorColumns[7]},
			},
		},
	}
	// BaseEquipmentColumns holds the columns for the "base_equipment" table.
	BaseEquipmentColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "sn", Type: field.TypeString},
		{Name: "category", Type: field.TypeEnum, Enums: []string{"private", "public"}},
		{Name: "operator_id", Type: field.TypeUint64},
		{Name: "station_id", Type: field.TypeUint64},
	}
	// BaseEquipmentTable holds the schema information for the "base_equipment" table.
	BaseEquipmentTable = &schema.Table{
		Name:       "base_equipment",
		Columns:    BaseEquipmentColumns,
		PrimaryKey: []*schema.Column{BaseEquipmentColumns[0]},
	}
	// BaseEquipmentExtraColumns holds the columns for the "base_equipment_extra" table.
	BaseEquipmentExtraColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "equipment_sn", Type: field.TypeString},
		{Name: "access_pod", Type: field.TypeString},
		{Name: "state", Type: field.TypeBool},
		{Name: "equipment_equipment_info", Type: field.TypeUint64, Unique: true},
	}
	// BaseEquipmentExtraTable holds the schema information for the "base_equipment_extra" table.
	BaseEquipmentExtraTable = &schema.Table{
		Name:       "base_equipment_extra",
		Columns:    BaseEquipmentExtraColumns,
		PrimaryKey: []*schema.Column{BaseEquipmentExtraColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "base_equipment_extra_base_equipment_equipment_info",
				Columns:    []*schema.Column{BaseEquipmentExtraColumns[4]},
				RefColumns: []*schema.Column{BaseEquipmentColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// BaseEvseColumns holds the columns for the "base_evse" table.
	BaseEvseColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "serial", Type: field.TypeString},
		{Name: "connector_number", Type: field.TypeInt},
		{Name: "equipment_evses", Type: field.TypeUint64},
	}
	// BaseEvseTable holds the schema information for the "base_evse" table.
	BaseEvseTable = &schema.Table{
		Name:       "base_evse",
		Columns:    BaseEvseColumns,
		PrimaryKey: []*schema.Column{BaseEvseColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "base_evse_base_equipment_evses",
				Columns:    []*schema.Column{BaseEvseColumns[3]},
				RefColumns: []*schema.Column{BaseEquipmentColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "evse_serial_equipment_evses",
				Unique:  true,
				Columns: []*schema.Column{BaseEvseColumns[1], BaseEvseColumns[3]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BaseConnectorTable,
		BaseEquipmentTable,
		BaseEquipmentExtraTable,
		BaseEvseTable,
	}
)

func init() {
	BaseConnectorTable.ForeignKeys[0].RefTable = BaseEquipmentTable
	BaseConnectorTable.ForeignKeys[1].RefTable = BaseEvseTable
	BaseConnectorTable.Annotation = &entsql.Annotation{
		Table: "base_connector",
	}
	BaseEquipmentTable.Annotation = &entsql.Annotation{
		Table: "base_equipment",
	}
	BaseEquipmentExtraTable.ForeignKeys[0].RefTable = BaseEquipmentTable
	BaseEquipmentExtraTable.Annotation = &entsql.Annotation{
		Table: "base_equipment_extra",
	}
	BaseEvseTable.ForeignKeys[0].RefTable = BaseEquipmentTable
	BaseEvseTable.Annotation = &entsql.Annotation{
		Table: "base_evse",
	}
}