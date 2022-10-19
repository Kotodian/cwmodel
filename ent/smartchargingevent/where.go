// Code generated by ent, DO NOT EDIT.

package smartchargingevent

import (
	"entgo.io/ent/dialect/sql"
	"github.com/Kotodian/ent-practice/ent/predicate"
	"github.com/Kotodian/gokit/datasource"
)

// ID filters vertices based on their ID field.
func ID(id datasource.UUID) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id datasource.UUID) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id datasource.UUID) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...datasource.UUID) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...datasource.UUID) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id datasource.UUID) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id datasource.UUID) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id datasource.UUID) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id datasource.UUID) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVersion), v))
	})
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedBy), vc))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedBy), vc))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// SmartID applies equality check predicate on the "smart_id" field. It's identical to SmartIDEQ.
func SmartID(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSmartID), vc))
	})
}

// EquipmentID applies equality check predicate on the "equipment_id" field. It's identical to EquipmentIDEQ.
func EquipmentID(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEquipmentID), vc))
	})
}

// ConnectorID applies equality check predicate on the "connector_id" field. It's identical to ConnectorIDEQ.
func ConnectorID(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldConnectorID), vc))
	})
}

// OrderID applies equality check predicate on the "order_id" field. It's identical to OrderIDEQ.
func OrderID(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderID), vc))
	})
}

// Unit applies equality check predicate on the "unit" field. It's identical to UnitEQ.
func Unit(v string) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUnit), v))
	})
}

// ValidFrom applies equality check predicate on the "valid_from" field. It's identical to ValidFromEQ.
func ValidFrom(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValidFrom), v))
	})
}

// ValidTo applies equality check predicate on the "valid_to" field. It's identical to ValidToEQ.
func ValidTo(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValidTo), v))
	})
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVersion), v))
	})
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldVersion), v))
	})
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...int64) predicate.SmartChargingEvent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldVersion), v...))
	})
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...int64) predicate.SmartChargingEvent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldVersion), v...))
	})
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldVersion), v))
	})
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldVersion), v))
	})
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldVersion), v))
	})
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldVersion), v))
	})
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedBy), vc))
	})
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedBy), vc))
	})
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...datasource.UUID) predicate.SmartChargingEvent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedBy), v...))
	})
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...datasource.UUID) predicate.SmartChargingEvent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedBy), v...))
	})
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedBy), vc))
	})
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedBy), vc))
	})
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedBy), vc))
	})
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedBy), vc))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...int64) predicate.SmartChargingEvent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...int64) predicate.SmartChargingEvent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedBy), vc))
	})
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedBy), vc))
	})
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...datasource.UUID) predicate.SmartChargingEvent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedBy), v...))
	})
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...datasource.UUID) predicate.SmartChargingEvent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedBy), v...))
	})
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedBy), vc))
	})
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedBy), vc))
	})
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedBy), vc))
	})
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedBy), vc))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...int64) predicate.SmartChargingEvent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...int64) predicate.SmartChargingEvent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// SmartIDEQ applies the EQ predicate on the "smart_id" field.
func SmartIDEQ(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSmartID), vc))
	})
}

// SmartIDNEQ applies the NEQ predicate on the "smart_id" field.
func SmartIDNEQ(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSmartID), vc))
	})
}

// SmartIDIn applies the In predicate on the "smart_id" field.
func SmartIDIn(vs ...datasource.UUID) predicate.SmartChargingEvent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSmartID), v...))
	})
}

// SmartIDNotIn applies the NotIn predicate on the "smart_id" field.
func SmartIDNotIn(vs ...datasource.UUID) predicate.SmartChargingEvent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSmartID), v...))
	})
}

// SmartIDGT applies the GT predicate on the "smart_id" field.
func SmartIDGT(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSmartID), vc))
	})
}

// SmartIDGTE applies the GTE predicate on the "smart_id" field.
func SmartIDGTE(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSmartID), vc))
	})
}

// SmartIDLT applies the LT predicate on the "smart_id" field.
func SmartIDLT(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSmartID), vc))
	})
}

// SmartIDLTE applies the LTE predicate on the "smart_id" field.
func SmartIDLTE(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSmartID), vc))
	})
}

// EquipmentIDEQ applies the EQ predicate on the "equipment_id" field.
func EquipmentIDEQ(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEquipmentID), vc))
	})
}

// EquipmentIDNEQ applies the NEQ predicate on the "equipment_id" field.
func EquipmentIDNEQ(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEquipmentID), vc))
	})
}

// EquipmentIDIn applies the In predicate on the "equipment_id" field.
func EquipmentIDIn(vs ...datasource.UUID) predicate.SmartChargingEvent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEquipmentID), v...))
	})
}

// EquipmentIDNotIn applies the NotIn predicate on the "equipment_id" field.
func EquipmentIDNotIn(vs ...datasource.UUID) predicate.SmartChargingEvent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEquipmentID), v...))
	})
}

// EquipmentIDGT applies the GT predicate on the "equipment_id" field.
func EquipmentIDGT(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEquipmentID), vc))
	})
}

// EquipmentIDGTE applies the GTE predicate on the "equipment_id" field.
func EquipmentIDGTE(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEquipmentID), vc))
	})
}

// EquipmentIDLT applies the LT predicate on the "equipment_id" field.
func EquipmentIDLT(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEquipmentID), vc))
	})
}

// EquipmentIDLTE applies the LTE predicate on the "equipment_id" field.
func EquipmentIDLTE(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEquipmentID), vc))
	})
}

// ConnectorIDEQ applies the EQ predicate on the "connector_id" field.
func ConnectorIDEQ(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldConnectorID), vc))
	})
}

// ConnectorIDNEQ applies the NEQ predicate on the "connector_id" field.
func ConnectorIDNEQ(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldConnectorID), vc))
	})
}

// ConnectorIDIn applies the In predicate on the "connector_id" field.
func ConnectorIDIn(vs ...datasource.UUID) predicate.SmartChargingEvent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldConnectorID), v...))
	})
}

// ConnectorIDNotIn applies the NotIn predicate on the "connector_id" field.
func ConnectorIDNotIn(vs ...datasource.UUID) predicate.SmartChargingEvent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldConnectorID), v...))
	})
}

// ConnectorIDGT applies the GT predicate on the "connector_id" field.
func ConnectorIDGT(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldConnectorID), vc))
	})
}

// ConnectorIDGTE applies the GTE predicate on the "connector_id" field.
func ConnectorIDGTE(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldConnectorID), vc))
	})
}

// ConnectorIDLT applies the LT predicate on the "connector_id" field.
func ConnectorIDLT(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldConnectorID), vc))
	})
}

// ConnectorIDLTE applies the LTE predicate on the "connector_id" field.
func ConnectorIDLTE(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldConnectorID), vc))
	})
}

// OrderIDEQ applies the EQ predicate on the "order_id" field.
func OrderIDEQ(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderID), vc))
	})
}

// OrderIDNEQ applies the NEQ predicate on the "order_id" field.
func OrderIDNEQ(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrderID), vc))
	})
}

// OrderIDIn applies the In predicate on the "order_id" field.
func OrderIDIn(vs ...datasource.UUID) predicate.SmartChargingEvent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOrderID), v...))
	})
}

// OrderIDNotIn applies the NotIn predicate on the "order_id" field.
func OrderIDNotIn(vs ...datasource.UUID) predicate.SmartChargingEvent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOrderID), v...))
	})
}

// OrderIDGT applies the GT predicate on the "order_id" field.
func OrderIDGT(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrderID), vc))
	})
}

// OrderIDGTE applies the GTE predicate on the "order_id" field.
func OrderIDGTE(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrderID), vc))
	})
}

// OrderIDLT applies the LT predicate on the "order_id" field.
func OrderIDLT(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrderID), vc))
	})
}

// OrderIDLTE applies the LTE predicate on the "order_id" field.
func OrderIDLTE(v datasource.UUID) predicate.SmartChargingEvent {
	vc := uint64(v)
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrderID), vc))
	})
}

// OrderIDIsNil applies the IsNil predicate on the "order_id" field.
func OrderIDIsNil() predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOrderID)))
	})
}

// OrderIDNotNil applies the NotNil predicate on the "order_id" field.
func OrderIDNotNil() predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOrderID)))
	})
}

// UnitEQ applies the EQ predicate on the "unit" field.
func UnitEQ(v string) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUnit), v))
	})
}

// UnitNEQ applies the NEQ predicate on the "unit" field.
func UnitNEQ(v string) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUnit), v))
	})
}

// UnitIn applies the In predicate on the "unit" field.
func UnitIn(vs ...string) predicate.SmartChargingEvent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUnit), v...))
	})
}

// UnitNotIn applies the NotIn predicate on the "unit" field.
func UnitNotIn(vs ...string) predicate.SmartChargingEvent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUnit), v...))
	})
}

// UnitGT applies the GT predicate on the "unit" field.
func UnitGT(v string) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUnit), v))
	})
}

// UnitGTE applies the GTE predicate on the "unit" field.
func UnitGTE(v string) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUnit), v))
	})
}

// UnitLT applies the LT predicate on the "unit" field.
func UnitLT(v string) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUnit), v))
	})
}

// UnitLTE applies the LTE predicate on the "unit" field.
func UnitLTE(v string) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUnit), v))
	})
}

// UnitContains applies the Contains predicate on the "unit" field.
func UnitContains(v string) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUnit), v))
	})
}

// UnitHasPrefix applies the HasPrefix predicate on the "unit" field.
func UnitHasPrefix(v string) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUnit), v))
	})
}

// UnitHasSuffix applies the HasSuffix predicate on the "unit" field.
func UnitHasSuffix(v string) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUnit), v))
	})
}

// UnitEqualFold applies the EqualFold predicate on the "unit" field.
func UnitEqualFold(v string) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUnit), v))
	})
}

// UnitContainsFold applies the ContainsFold predicate on the "unit" field.
func UnitContainsFold(v string) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUnit), v))
	})
}

// ValidFromEQ applies the EQ predicate on the "valid_from" field.
func ValidFromEQ(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValidFrom), v))
	})
}

// ValidFromNEQ applies the NEQ predicate on the "valid_from" field.
func ValidFromNEQ(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldValidFrom), v))
	})
}

// ValidFromIn applies the In predicate on the "valid_from" field.
func ValidFromIn(vs ...int64) predicate.SmartChargingEvent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldValidFrom), v...))
	})
}

// ValidFromNotIn applies the NotIn predicate on the "valid_from" field.
func ValidFromNotIn(vs ...int64) predicate.SmartChargingEvent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldValidFrom), v...))
	})
}

// ValidFromGT applies the GT predicate on the "valid_from" field.
func ValidFromGT(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldValidFrom), v))
	})
}

// ValidFromGTE applies the GTE predicate on the "valid_from" field.
func ValidFromGTE(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldValidFrom), v))
	})
}

// ValidFromLT applies the LT predicate on the "valid_from" field.
func ValidFromLT(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldValidFrom), v))
	})
}

// ValidFromLTE applies the LTE predicate on the "valid_from" field.
func ValidFromLTE(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldValidFrom), v))
	})
}

// ValidFromIsNil applies the IsNil predicate on the "valid_from" field.
func ValidFromIsNil() predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldValidFrom)))
	})
}

// ValidFromNotNil applies the NotNil predicate on the "valid_from" field.
func ValidFromNotNil() predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldValidFrom)))
	})
}

// ValidToEQ applies the EQ predicate on the "valid_to" field.
func ValidToEQ(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValidTo), v))
	})
}

// ValidToNEQ applies the NEQ predicate on the "valid_to" field.
func ValidToNEQ(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldValidTo), v))
	})
}

// ValidToIn applies the In predicate on the "valid_to" field.
func ValidToIn(vs ...int64) predicate.SmartChargingEvent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldValidTo), v...))
	})
}

// ValidToNotIn applies the NotIn predicate on the "valid_to" field.
func ValidToNotIn(vs ...int64) predicate.SmartChargingEvent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldValidTo), v...))
	})
}

// ValidToGT applies the GT predicate on the "valid_to" field.
func ValidToGT(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldValidTo), v))
	})
}

// ValidToGTE applies the GTE predicate on the "valid_to" field.
func ValidToGTE(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldValidTo), v))
	})
}

// ValidToLT applies the LT predicate on the "valid_to" field.
func ValidToLT(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldValidTo), v))
	})
}

// ValidToLTE applies the LTE predicate on the "valid_to" field.
func ValidToLTE(v int64) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldValidTo), v))
	})
}

// ValidToIsNil applies the IsNil predicate on the "valid_to" field.
func ValidToIsNil() predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldValidTo)))
	})
}

// ValidToNotNil applies the NotNil predicate on the "valid_to" field.
func ValidToNotNil() predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldValidTo)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SmartChargingEvent) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SmartChargingEvent) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SmartChargingEvent) predicate.SmartChargingEvent {
	return predicate.SmartChargingEvent(func(s *sql.Selector) {
		p(s.Not())
	})
}