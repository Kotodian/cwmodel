package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type EquipmentIot struct {
	ent.Schema
}

func (EquipmentIot) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ModelMixin{},
	}
}

func (EquipmentIot) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "equip_iot"},
		edge.Annotation{StructTag: `json:"-"`},
	}
}

func (EquipmentIot) Fields() []ent.Field {
	return []ent.Field{
		field.String("iccid").Optional().Nillable().Comment("iccid"),
		field.String("imei").Optional().Nillable().Comment("imei"),
		field.String("remote_address").Optional().Nillable().Comment("remote_address"),
	}
}

func (EquipmentIot) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipment", Equipment.Type).Ref("equipment_iot").Unique(),
	}
}
