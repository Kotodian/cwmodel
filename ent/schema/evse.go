package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/Kotodian/gokit/datasource"
	"github.com/Kotodian/gokit/id"
)

// Evse holds the schema definition for the Evse entity.
type Evse struct {
	ent.Schema
}

func (Evse) Mixins() []ent.Mixin {
	return []ent.Mixin{
		ModelMixin{},
	}
}

func (Evse) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "base_evse"},
	}
}

// Fields of the Evse.
func (Evse) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").GoType(datasource.UUID(0)).DefaultFunc(id.Next()).Immutable().Comment("主键"),
		field.String("serial").Comment("设备序列号"),
		field.Int("connector_number").Comment("枪数量"),
	}
}

// Edges of the Evse.
func (Evse) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipment", Equipment.Type).Ref("evses").Unique().Required(),
		edge.To("connectors", Connector.Type),
	}
}

func (Evse) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("equipment").Fields("serial").Unique(),
	}
}
