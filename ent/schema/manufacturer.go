package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Manufacturer struct {
	ent.Schema
}

func (Manufacturer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ModelMixin{},
	}
}

func (Manufacturer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "equip_manufacturer"},
	}
}

func (Manufacturer) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").Comment("产商代码"),
		field.String("name").Optional().Comment("产商名称"),
	}
}
