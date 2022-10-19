package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/gokit/datasource"
	"github.com/Kotodian/gokit/id"
)

type OrderEvent struct {
	ent.Schema
}

// func (OrderEvent) Mixin() []ent.Mixin {
// 	return []ent.Mixin{
// 		ModelMixin{},
// 	}
// }

func (OrderEvent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order_event"},
		edge.Annotation{StructTag: `json:"-"`},
	}
}

func (OrderEvent) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").GoType(datasource.UUID(0)).DefaultFunc(id.Next().Uint64()).Immutable().Comment("主键"),
		field.String("content").Comment("事件内容"),
		field.Int64("occurrence").Comment("事件发生时间"),
	}
}

func (OrderEvent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order_info", OrderInfo.Type).Ref("order_event").Unique(),
	}
}
