package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type EquipmentAlarm struct {
	ent.Schema
}

func (EquipmentAlarm) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ModelMixin{},
	}
}

func (EquipmentAlarm) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "base_equipment_alarm"},
		edge.Annotation{StructTag: `json:"-"`},
	}
}

// Fields of the Connector.
func (EquipmentAlarm) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("dtc_code").Comment("告警代码"),
		field.String("remote_address").Comment("ip地址"),
		field.Int64("trigger_time").Optional().Nillable().Comment("触发时间"),
		field.Int64("final_time").Optional().Nillable().Comment("结束时间"),
		field.Int("count").Default(0).Comment("数量"),
	}
}

// Edges of the Connector.
func (EquipmentAlarm) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipment", Equipment.Type).
			Ref("equipment_alarm").
			Unique().
			Required(),
	}
}
