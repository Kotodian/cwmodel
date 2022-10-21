package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/gokit/datasource"
)

// Evse holds the schema definition for the Evse entity.
type Evse struct {
	ent.Schema
}

func (Evse) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ModelMixin{},
	}
}

func (Evse) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "base_evse"},
		edge.Annotation{StructTag: `json:"-"`},
	}
}

// Fields of the Evse.
func (Evse) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("equipment_id").GoType(datasource.UUID(0)).Comment("桩id"),
		field.String("serial").Comment("设备序列号"),
		field.Int("connector_number").Comment("枪数量"),
	}
}

// Edges of the Evse.
func (Evse) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipment", Equipment.Type).Field("equipment_id").Ref("evse").Unique().Required(),
		edge.To("connector", Connector.Type).StorageKey(edge.Column("evse_id")),
	}
}
