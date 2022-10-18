package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/ent-practice/ent/types"
	"github.com/Kotodian/gokit/datasource"
)

// Equipment holds the schema definition for the Equipment entity.
type Equipment struct {
	ent.Schema
}

func (Equipment) Mixins() []ent.Mixin {
	return []ent.Mixin{
		ModelMixin{},
	}
}

func (Equipment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "base_equipment"},
	}
}

// Fields of the Equipment.
func (Equipment) Fields() []ent.Field {
	return []ent.Field{
		field.String("sn").NotEmpty().Comment("桩序列号"),
		field.String("sn2").NotEmpty().Comment("桩序列号2"),
		field.Enum("category").GoType(types.EquipmentCategory(0)).Comment("桩类型"),
		field.Uint64("operator_id").GoType(datasource.UUID(0)).Comment("运营商id"),
		field.Uint64("station_id").GoType(datasource.UUID(0)).Comment("站点id"),
	}
}

// Edges of the Equipment.
func (Equipment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("model", Model.Type).Ref("equipment").Unique(),
		edge.From("firmware", Firmware.Type).Ref("equipment").Unique(),
		edge.From("manufacturer", Manufacturer.Type).Ref("equipment").Unique(),
		edge.To("equipment_info", EquipmentInfo.Type).Unique(),
		edge.To("evse", Evse.Type),
		edge.To("connector", Connector.Type),
		edge.To("equipment_alarm", EquipmentAlarm.Type),
		edge.To("equipment_iot", EquipmentIot.Type).Unique(),
		edge.To("equipment_firmware_effect", EquipmentFirmwareEffect.Type),
		edge.To("order_info", OrderInfo.Type),
		edge.To("reservation", Reservation.Type),
	}
}
