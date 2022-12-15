// Code generated by ent, DO NOT EDIT.

package manufacturer

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Kotodian/cwmodel/predicate"
	"github.com/Kotodian/gokit/datasource"
)

// ID filters vertices based on their ID field.
func ID(id datasource.UUID) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id datasource.UUID) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id datasource.UUID) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...datasource.UUID) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...datasource.UUID) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id datasource.UUID) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id datasource.UUID) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id datasource.UUID) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id datasource.UUID) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v int64) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVersion), v))
	})
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v datasource.UUID) predicate.Manufacturer {
	vc := uint64(v)
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedBy), vc))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v int64) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v datasource.UUID) predicate.Manufacturer {
	vc := uint64(v)
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedBy), vc))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v int64) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCode), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v int64) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVersion), v))
	})
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v int64) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldVersion), v))
	})
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...int64) predicate.Manufacturer {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldVersion), v...))
	})
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...int64) predicate.Manufacturer {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldVersion), v...))
	})
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v int64) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldVersion), v))
	})
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v int64) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldVersion), v))
	})
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v int64) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldVersion), v))
	})
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v int64) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldVersion), v))
	})
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v datasource.UUID) predicate.Manufacturer {
	vc := uint64(v)
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedBy), vc))
	})
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v datasource.UUID) predicate.Manufacturer {
	vc := uint64(v)
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedBy), vc))
	})
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...datasource.UUID) predicate.Manufacturer {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedBy), v...))
	})
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...datasource.UUID) predicate.Manufacturer {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedBy), v...))
	})
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v datasource.UUID) predicate.Manufacturer {
	vc := uint64(v)
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedBy), vc))
	})
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v datasource.UUID) predicate.Manufacturer {
	vc := uint64(v)
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedBy), vc))
	})
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v datasource.UUID) predicate.Manufacturer {
	vc := uint64(v)
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedBy), vc))
	})
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v datasource.UUID) predicate.Manufacturer {
	vc := uint64(v)
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedBy), vc))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v int64) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v int64) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...int64) predicate.Manufacturer {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...int64) predicate.Manufacturer {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v int64) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v int64) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v int64) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v int64) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v datasource.UUID) predicate.Manufacturer {
	vc := uint64(v)
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedBy), vc))
	})
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v datasource.UUID) predicate.Manufacturer {
	vc := uint64(v)
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedBy), vc))
	})
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...datasource.UUID) predicate.Manufacturer {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedBy), v...))
	})
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...datasource.UUID) predicate.Manufacturer {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedBy), v...))
	})
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v datasource.UUID) predicate.Manufacturer {
	vc := uint64(v)
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedBy), vc))
	})
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v datasource.UUID) predicate.Manufacturer {
	vc := uint64(v)
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedBy), vc))
	})
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v datasource.UUID) predicate.Manufacturer {
	vc := uint64(v)
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedBy), vc))
	})
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v datasource.UUID) predicate.Manufacturer {
	vc := uint64(v)
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedBy), vc))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v int64) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v int64) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...int64) predicate.Manufacturer {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...int64) predicate.Manufacturer {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v int64) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v int64) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v int64) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v int64) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCode), v))
	})
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCode), v))
	})
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.Manufacturer {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCode), v...))
	})
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.Manufacturer {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCode), v...))
	})
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCode), v))
	})
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCode), v))
	})
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCode), v))
	})
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCode), v))
	})
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCode), v))
	})
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCode), v))
	})
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCode), v))
	})
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCode), v))
	})
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCode), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Manufacturer {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Manufacturer {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldName)))
	})
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldName)))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// HasFirmware applies the HasEdge predicate on the "firmware" edge.
func HasFirmware() predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FirmwareTable, FirmwareColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFirmwareWith applies the HasEdge predicate on the "firmware" edge with a given conditions (other predicates).
func HasFirmwareWith(preds ...predicate.Firmware) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FirmwareInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FirmwareTable, FirmwareColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Manufacturer) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Manufacturer) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
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
func Not(p predicate.Manufacturer) predicate.Manufacturer {
	return predicate.Manufacturer(func(s *sql.Selector) {
		p(s.Not())
	})
}
