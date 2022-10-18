package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/gokit/datasource"
)

type OrderEvent struct {
	ent.Schema
}

func (OrderEvent) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ModelMixin{},
	}
}

func (OrderEvent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order_event"},
		edge.Annotation{StructTag: `json:"-"`},
	}
}

func (OrderEvent) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("order_id").GoType(datasource.UUID(0)).Comment("订单id"),
		field.String("content").Comment("事件内容"),
		field.Int64("occurrence").Comment("事件发生时间"),
	}
}

func (OrderEvent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order_info", OrderInfo.Type).Ref("order_event").Unique(),
	}
}
