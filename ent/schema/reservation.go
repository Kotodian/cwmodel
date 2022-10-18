package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Reservation struct {
	ent.Schema
}

func (Reservation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ModelMixin{},
	}
}

func (Reservation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "reservation_charging_release"},
	}
}

func (Reservation) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("reservation_id").Comment("预约id"),
		field.Int("authorization_mode").Comment("授权模式"),
		field.String("authorization_id").Comment("授权id"),
		field.String("additional").Optional().Comment("额外授权id"),
		field.String("customer_id").Optional().Comment("客户id"),
		field.Int64("expired").Comment("过期时间"),
		field.Int("state").Comment("预约状态"),
	}
}

func (Reservation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipment", Equipment.Type).Ref("reservation").Unique(),
		edge.From("connector", Connector.Type).Ref("reservation").Unique(),
	}
}
