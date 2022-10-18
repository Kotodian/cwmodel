package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type AppModuleInfo struct {
	ent.Schema
}

func (AppModuleInfo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "app_module_info"},
	}
}

// Fields of the Connector.
func (AppModuleInfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称"),
		field.String("desc").Comment("描述"),
	}
}
