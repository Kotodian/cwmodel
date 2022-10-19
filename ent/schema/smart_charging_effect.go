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

type SmartChargingEffect struct {
	ent.Schema
}

func (SmartChargingEffect) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ModelMixin{},
	}
}

func (SmartChargingEffect) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "smart_charging_effect"},
		edge.Annotation{StructTag: `json:"-"`},
	}
}

func (SmartChargingEffect) Fields() []ent.Field {
	return []ent.Field{
		// field.Uint64("equipment_id").GoType(datasource.UUID(0)).Comment("桩端id"),
		// field.Uint64("connector_id").GoType(datasource.UUID(0)).Comment("枪id"),
		// field.Uint64("order_id").Optional().GoType(datasource.UUID(0)).Comment("订单id"),
		field.Int64("smart_id").Comment("智慧id"),
		field.Uint64("pid").GoType(datasource.UUID(0)).Comment("父id"),
		field.String("unit").Comment("单位(W或者A)"),
		field.String("equipment_sn").Comment("桩sn号"),
		field.Int64("valid_from").Optional().Comment("有效开始时间"),
		field.Int64("valid_to").Optional().Comment("有效结束时间"),
		field.JSON("spec", []types.ChargingSchedulePeriod{}).Comment("时间间隔"),
	}
}

func (SmartChargingEffect) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipment", Equipment.Type).Ref("smart_charging_effect").Unique().Required(),
		edge.From("connector", Connector.Type).Ref("smart_charging_effect").Unique().Required(),
		edge.From("order_info", OrderInfo.Type).Ref("smart_charging_effect").Unique(),
	}
}
