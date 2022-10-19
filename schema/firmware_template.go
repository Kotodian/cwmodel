package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Firmware struct {
	ent.Schema
}

func (Firmware) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ModelMixin{},
	}
}

func (Firmware) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "equip_firmware_template"},
		edge.Annotation{StructTag: `json:"-"`},
	}
}

func (Firmware) Fields() []ent.Field {
	return []ent.Field{
		field.String("equip_version").Comment("固件版本"),
	}
}

func (Firmware) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("equipment_firmware_effect", EquipmentFirmwareEffect.Type).StorageKey(edge.Column("firmware_id")),
	}
}
