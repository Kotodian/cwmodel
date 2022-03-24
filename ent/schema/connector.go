package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/Kotodian/ent-practice/ent/enums"
	"github.com/Kotodian/gokit/datasource"
	"github.com/Kotodian/gokit/id"
)

// Connector holds the schema definition for the Connector entity.
type Connector struct {
	ent.Schema
}

func (Connector) Mixins() []ent.Mixin {
	return []ent.Mixin{
		ModelMixin{},
	}
}

func (Connector) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "base_connector"},
	}
}

// Fields of the Connector.
func (Connector) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").GoType(datasource.UUID(0)).DefaultFunc(id.Next()).Immutable().Comment("主键"),
		field.String("equipment_sn").Comment("桩序列号"),
		field.String("evse_serial").Comment("设备序列号"),
		field.String("serial").Comment("枪序列号"),
		field.Enum("current_state").GoType(enums.ConnectorState(0)).Comment("当前状态"),
		field.Enum("before_state").GoType(enums.ConnectorState(0)).Comment("之前状态"),
	}
}

// Edges of the Connector.
func (Connector) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("evse", Evse.Type).Ref("connectors").Unique().Required(),
		edge.From("equipment", Equipment.Type).Ref("connectors").Unique().Required(),
	}
}

func (Connector) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("equipment").Edges("evse").Fields("serial").Unique(),
	}
}
