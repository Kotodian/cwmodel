package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/ent-practice/ent/enums"
	"github.com/Kotodian/gokit/datasource"
	"github.com/Kotodian/gokit/id"
)

// Equipment holds the schema definition for the Equipment entity.
type Equipment struct {
	ent.Schema
}

func (Equipment) Mixins() []ent.Mixin {
	return []ent.Mixin{
		ModelMixin{},
	}
}

func (Equipment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "base_equipment"},
	}
}

// Fields of the Equipment.
func (Equipment) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").GoType(datasource.UUID(0)).DefaultFunc(id.Next()).Immutable().Comment("主键"),
		field.String("sn").NotEmpty().Comment("桩序列号"),
		field.String("sn2").NotEmpty().Comment("桩序列号2"),
		field.Enum("category").GoType(enums.EquipmentCategory(0)).Comment("桩类型"),
		field.Uint64("operator_id").GoType(datasource.UUID(0)).Comment("运营商id"),
		field.Uint64("station_id").GoType(datasource.UUID(0)).Comment("站点id"),
	}
}

// Edges of the Equipment.
func (Equipment) Edges() []ent.Edge {
	return []ent.Edge{
		// 与info是一对一的关系
		edge.To("equipment_info", EquipmentInfo.Type).Unique(),
		// 与evse是一对多的关系
		edge.To("evses", Evse.Type),
		// 与connector是一对多的关系
		edge.To("connectors", Connector.Type),
	}
}
