package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type EquipmentFirmwareEffect struct {
	ent.Schema
}

func (EquipmentFirmwareEffect) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ModelMixin{},
	}
}

func (EquipmentFirmwareEffect) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("request_id").Comment("请求id"),
		field.Int("state").Comment("状态"),
	}
}

func (EquipmentFirmwareEffect) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipment", Equipment.Type).Ref("equipment_firmware_effect").Unique(),
		edge.From("firmware", Firmware.Type).Ref("equipment_firmware_effect").Unique(),
	}
}
