package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/Kotodian/gokit/datasource"
	"github.com/Kotodian/gokit/id"
)

type ModelMixin struct {
	mixin.Schema
}

func (ModelMixin) Annotations() []schema.Annotation {
	return []schema.Annotation{}
}

func (ModelMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").GoType(datasource.UUID(0)).DefaultFunc(id.Next().Uint64()).Immutable().StructTag(`json:"id"`).Comment("主键"),
		field.Int64("version").Default(1).StructTag(`json:"version"`).Comment("乐观锁"),
		field.Uint64("created_by").Default(1).GoType(datasource.UUID(0)).Immutable().StructTag(`json:"-"`).Comment("创建者"),
		field.Int64("created_at").Default(time.Now().Unix()).Immutable().StructTag(`json:"-"`).Comment("创建时间"),
		field.Uint64("updated_by").Default(1).GoType(datasource.UUID(0)).StructTag(`json:"-"`).Comment("修改者"),
		field.Int64("updated_at").UpdateDefault(time.Now().Unix).Default(time.Now().Unix()).StructTag(`json:"-"`).Comment("修改时间"),
	}
}
