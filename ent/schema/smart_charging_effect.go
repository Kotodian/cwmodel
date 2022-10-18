package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/ent-practice/ent/types"
	"github.com/Kotodian/gokit/datasource"
)

type SmartChargingEvent struct {
	ent.Schema
}

func (SmartChargingEvent) Mixins() []ent.Mixin {
	return []ent.Mixin{
		ModelMixin{},
	}
}

func (SmartChargingEvent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "smart_charging_effect"},
	}
}

func (SmartChargingEvent) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("smart_id").GoType(datasource.UUID(0)).Comment("智慧id"),
		field.Uint64("equipment_id").GoType(datasource.UUID(0)).Comment("桩端id"),
		field.Uint64("connector_id").GoType(datasource.UUID(0)).Comment("枪id"),
		field.Uint64("order_id").Optional().GoType(datasource.UUID(0)).Comment("订单id"),
		field.String("unit").Comment("单位(W或者A)"),
		field.Int64("valid_from").Optional().Comment("有效开始时间"),
		field.Int64("valid_to").Optional().Comment("有效结束时间"),
		field.JSON("spec", []types.ChargingSchedulePeriod{}).Comment("时间间隔"),
	}
}
