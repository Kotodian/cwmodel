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

func (OrderInfo) Mixins() []ent.Mixin {
	return []ent.Mixin{
		ModelMixin{},
	}
}

func (OrderInfo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order_info"},
	}
}

func (OrderInfo) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("remote_start_id").Optional().Comment("远程启动id"),
		field.String("transaction_id").Comment("桩端订单id"),
		field.String("authorization_id").Optional().Comment("授权id"),
		field.String("customer_id").Optional().Comment("客户id"),
		field.String("caller_order_id").Optional().Comment("第三方订单id"),
		field.Float("total_electricity").Optional().Comment("总充电电量"),
		field.Float("charge_start_electricity").Optional().Comment("起始电量"),
		field.Float("charge_stop_electricity").Optional().Comment("结束电量"),
		field.Float("sharp_electricity").Optional().Comment("尖电量"),
		field.Float("peak_electricity").Optional().Comment("峰电量"),
		field.Float("flat_electricity").Optional().Comment("平电量"),
		field.Float("valley_electricity").Optional().Comment("谷电量"),
		field.Int32("stop_reason_code").Optional().Comment("停止原因代码"),
		field.Bool("offline").Comment("是否为离线订单"),
		field.Int64("price_scheme_release_id").Comment("计费模板id"),
		field.Int64("order_start_time").Optional().Comment("订单开始时间"),
		field.Int64("order_final_time").Optional().Comment("订单结束时间"),
		field.Int64("charge_start_time").Optional().Comment("充电开始时间"),
		field.Int64("charge_final_time").Optional().Comment("充电结束时间"),
		field.Int64("intellect_id").Optional().Comment("智慧充电id"),
		field.Uint64("station_id").GoType(datasource.UUID(0)).Optional().Comment("场站id"),
		field.Uint64("operator_id").GoType(datasource.UUID(0)).Optional().Comment("运营商id"),
	}
}

func (OrderInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("connector", Connector.Type).Ref("order_info").Unique(),
		edge.From("equipment", Equipment.Type).Ref("order_info").Unique(),
		edge.To("order_event", OrderEvent.Type),
	}
}
