package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/Kotodian/cwmodel/types"
)

type Model struct {
	ent.Schema
}

func (Model) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ModelMixin{},
	}
}

func (Model) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "equip_model"},
		edge.Annotation{StructTag: `json:"-"`},
	}
}

func (Model) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").Comment("型号代码"),
		field.String("name").Comment("型号名称"),
		field.String("phase_category").Comment("相位类型"),
		field.String("current_category").Comment("电流类型"),
		field.Int("connector_category").GoType(types.ConnectorType(0)).Comment("充电设备接口类型"),
	}
}

func (Model) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("firmware", Firmware.Type),
	}
}
