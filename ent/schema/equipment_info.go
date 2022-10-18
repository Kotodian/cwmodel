package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/gokit/datasource"
)

// EquipmentInfo holds the schema definition for the EquipmentInfo entity.
type EquipmentInfo struct {
	ent.Schema
}

func (EquipmentInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ModelMixin{},
	}
}

func (EquipmentInfo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "base_equipment_extra"},
		edge.Annotation{StructTag: `json:"-"`},
	}
}

// Fields of the EquipmentInfo.
func (EquipmentInfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("equipment_sn").Comment("桩序列号").StructTag(`json:"equipmentSN"`),
		field.Uint64("model_id").GoType(datasource.UUID(0)).Comment("型号id").StructTag(`json:"modelId"`),
		field.Uint64("manufacturer_id").GoType(datasource.UUID(0)).Comment("产商id").StructTag(`json:"manufacturerId"`),
		field.Uint64("firmware_id").GoType(datasource.UUID(0)).Comment("固件id").StructTag(`json:"firmwareId"`),
		field.String("access_pod").Comment("所连接的pod").StructTag(`json:"accessPod"`),
		field.Bool("state").Comment("在线或者离线").StructTag(`json:"state"`),
		field.Uint("evse_number").Comment("evse数量").StructTag(`json:"evseNumber"`),
		field.Uint("alarm_number").Comment("告警数量").StructTag(`json:"alarmNums"`),
		field.Int64("register_datetime").Comment("注册时间").StructTag(`json:"registerDatetime"`),
		field.Int64("remote_address").Comment("远程ip地址").StructTag(`json:"remoteAddress"`),
	}
}

// Edges of the EquipmentInfo.
func (EquipmentInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipment", Equipment.Type).Ref("equipment_info").Unique().Required(),
	}
}
