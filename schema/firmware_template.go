package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/gokit/datasource"
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
		field.Uint64("model_id").GoType(datasource.UUID(0)).Comment("型号id"),
		field.Uint64("manufacturer_id").GoType(datasource.UUID(0)).Comment("产商id"),
		field.String("equip_version").Comment("固件版本"),
	}
}

func (Firmware) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("equipment_firmware_effect", EquipmentFirmwareEffect.Type),
		edge.From("model", Model.Type).Field("model_id").Ref("firmware").Unique().Required(),
		edge.From("manufacturer", Manufacturer.Type).Field("manufacturer_id").Ref("firmware").Unique().Required(),
	}
}
