package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/gokit/datasource"
)

type OrderInfo struct {
	ent.Schema
}

func (OrderInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ModelMixin{},
	}
}

func (OrderInfo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order_info"},
		edge.Annotation{StructTag: `json:"-"`},
	}
}

func (OrderInfo) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("equipment_id").GoType(datasource.UUID(0)).Comment("桩id"),
		field.Uint64("connector_id").GoType(datasource.UUID(0)).Comment("枪id"),
		field.Int64("remote_start_id").Optional().Nillable().Comment("远程启动id"),
		field.String("transaction_id").Comment("桩端订单id"),
		field.String("authorization_id").Optional().Nillable().Comment("授权id"),
		field.Int("authorization_mode").Optional().Nillable().Comment("授权模式"),
		field.String("customer_id").Optional().Nillable().Comment("客户id"),
		field.String("caller_order_id").Optional().Nillable().Comment("第三方订单id"),
		field.Float("total_electricity").Optional().Nillable().Comment("总充电电量"),
		field.Float("charge_start_electricity").Optional().Nillable().Comment("起始电量"),
		field.Float("charge_final_electricity").Optional().Nillable().Comment("结束电量"),
		field.Float("sharp_electricity").Optional().Nillable().Comment("尖电量"),
		field.Float("peak_electricity").Optional().Nillable().Comment("峰电量"),
		field.Float("flat_electricity").Optional().Nillable().Comment("平电量"),
		field.Float("valley_electricity").Optional().Nillable().Comment("谷电量"),
		field.Int32("stop_reason_code").Optional().Nillable().Comment("停止原因代码"),
		field.Int("state").Optional().Comment("订单状态"),
		field.Bool("offline").Comment("是否为离线订单"),
		field.Int64("price_scheme_release_id").Comment("计费模板id"),
		field.Int64("order_start_time").Optional().Nillable().Comment("订单开始时间"),
		field.Int64("order_final_time").Optional().Nillable().Comment("订单结束时间"),
		field.Int64("charge_start_time").Optional().Nillable().Comment("充电开始时间"),
		field.Int64("charge_final_time").Optional().Nillable().Comment("充电结束时间"),
		field.Int64("intellect_id").Optional().Nillable().Comment("智慧充电id"),
		field.Uint64("station_id").GoType(datasource.UUID(0)).Optional().Nillable().Comment("场站id"),
		field.Uint64("operator_id").GoType(datasource.UUID(0)).Optional().Nillable().Comment("运营商id"),
	}
}

func (OrderInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("connector", Connector.Type).Field("connector_id").Ref("order_info").Unique().Required(),
		edge.From("equipment", Equipment.Type).Field("equipment_id").Ref("order_info").Unique().Required(),
		edge.To("order_event", OrderEvent.Type).StorageKey(edge.Column("order_id")),
		// edge.To("smart_charging_effect", SmartChargingEffect.Type).StorageKey(edge.Column("order_id")).Unique(),
	}
}
