package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/gokit/datasource"
)

// Equipment holds the schema definition for the Equipment entity.
type Equipment struct {
	ent.Schema
}

func (Equipment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ModelMixin{},
	}
}

func (Equipment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "base_equipment"},
		edge.Annotation{StructTag: `json:"-"`},
	}
}

// Fields of the Equipment.
func (Equipment) Fields() []ent.Field {
	return []ent.Field{
		field.String("sn").NotEmpty().Comment("桩序列号").StructTag(`json:"sn"`),
		field.Uint64("operator_id").GoType(datasource.UUID(0)).Comment("运营商id").StructTag(`json:"operatorId"`),
		field.Uint64("station_id").GoType(datasource.UUID(0)).Comment("站点id").StructTag(`json:"stationId"`),
		// field.String("sn").NotEmpty().Comment("桩序列号"),
		// field.Uint64("operator_id").GoType(datasource.UUID(0)).Comment("运营商id"),
		// field.Uint64("station_id").GoType(datasource.UUID(0)).Comment("站点id"),
	}
}

// Edges of the Equipment.
func (Equipment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("equipment_info", EquipmentInfo.Type).Unique(),
		edge.To("evse", Evse.Type).StorageKey(edge.Column("equipment_id")),
		edge.To("connector", Connector.Type).StorageKey(edge.Column("equipment_id")),
		edge.To("equipment_alarm", EquipmentAlarm.Type),
		edge.To("equipment_iot", EquipmentIot.Type).StorageKey(edge.Column("equipment_id")).Unique(),
		edge.To("equipment_firmware_effect", EquipmentFirmwareEffect.Type).StorageKey(edge.Column("equipment_id")),
		edge.To("order_info", OrderInfo.Type).StorageKey(edge.Column("equipment_id")),
		edge.To("reservation", Reservation.Type).StorageKey(edge.Column("equipment_id")),
		edge.To("equipment_log", EquipmentLog.Type).StorageKey(edge.Column("equipment_id")).StructTag(`json:"equipmentId"`),
		edge.To("smart_charging_effect", SmartChargingEffect.Type).StorageKey(edge.Column("equipment_id")),
	}
}
