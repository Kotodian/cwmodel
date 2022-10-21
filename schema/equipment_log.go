package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/gokit/datasource"
)

type EquipmentLog struct {
	ent.Schema
}

func (EquipmentLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ModelMixin{},
	}
}

func (EquipmentLog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "equip_logs"},
		edge.Annotation{StructTag: `json:"-"`},
	}
}

func (EquipmentLog) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("equipment_id").GoType(datasource.UUID(0)).Comment("桩id").StructTag(`json:"equipmentId"`),
		field.Int64("request_id").Comment("请求id").StructTag(`json:"requestId"`),
		field.Int("state").Comment("状态").StructTag(`json:"state"`),
		field.Uint64("data_link").GoType(datasource.UUID(0)).Comment("data link").StructTag(`json:"dataLink"`),
	}
}

func (EquipmentLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipment", Equipment.Type).Field("equipment_id").Ref("equipment_log").Unique().Required(),
	}
}
