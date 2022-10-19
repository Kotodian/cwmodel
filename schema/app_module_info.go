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

type AppModuleInfo struct {
	ent.Schema
}

func (AppModuleInfo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "app_module_info"},
		edge.Annotation{StructTag: `json:"-"`},
	}
}

// Fields of the Connector.
func (AppModuleInfo) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").GoType(datasource.UUID(0)).DefaultFunc(id.Next().Uint64()).Immutable().Comment("主键"),
		field.String("name").Comment("名称"),
		field.String("desc").Comment("描述"),
	}
}
