package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/gokit/datasource"
)

// Connector holds the schema definition for the Connector entity.
type Connector struct {
	ent.Schema
}

func (Connector) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ModelMixin{},
	}
}

func (Connector) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "base_connector"},
		edge.Annotation{StructTag: `json:"-"`},
	}
}

// Fields of the Connector.
func (Connector) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("equipment_id").GoType(datasource.UUID(0)).Comment("桩id"),
		field.Uint64("evse_id").GoType(datasource.UUID(0)).Comment("设备id"),
		field.String("equipment_sn").Comment("桩序列号"),
		field.String("evse_serial").Comment("设备序列号"),
		field.String("serial").Comment("枪序列号"),
		field.Int("current_state").Comment("当前状态"),
		field.Int("before_state").Comment("之前状态"),
		field.Int("charging_state").Optional().Comment("充电状态"),
		field.Uint64("reservation_id").Optional().Nillable().GoType(datasource.UUID(0)).Comment("预约id"),
		field.Uint64("order_id").Optional().Nillable().GoType(datasource.UUID(0)).Comment("订单id"),
		field.String("park_no").Default("0001").Comment("停车编号"),
	}
}

// Edges of the Connector.
func (Connector) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("evse", Evse.Type).Field("evse_id").Ref("connector").Unique().Required(),
		edge.From("equipment", Equipment.Type).Field("equipment_id").Ref("connector").Unique().Required(),
		edge.To("order_info", OrderInfo.Type),
		edge.To("reservation", Reservation.Type),
		edge.To("smart_charging_effect", SmartChargingEffect.Type),
	}
}
