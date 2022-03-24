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

// EquipmentInfo holds the schema definition for the EquipmentInfo entity.
type EquipmentInfo struct {
	ent.Schema
}

func (EquipmentInfo) Mixins() []ent.Mixin {
	return []ent.Mixin{
		ModelMixin{},
	}
}

func (EquipmentInfo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "base_equipment_extra"},
	}
}

// Fields of the EquipmentInfo.
func (EquipmentInfo) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").GoType(datasource.UUID(0)).DefaultFunc(id.Next()).Immutable().Comment("主键"),
		field.String("equipment_sn").Comment("桩序列号"),
		field.String("access_pod").Comment("所连接的pod"),
		field.Bool("state").Comment("在线或者离线"),
	}
}

// Edges of the EquipmentInfo.
func (EquipmentInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipment", Equipment.Type).Ref("equipment_info").Unique().Required(),
	}
}
