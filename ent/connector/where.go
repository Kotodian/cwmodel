// Code generated by entc, DO NOT EDIT.

package connector

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Kotodian/ent-practice/ent/enums"
	"github.com/Kotodian/ent-practice/ent/predicate"
	"github.com/Kotodian/gokit/datasource"
)

// ID filters vertices based on their ID field.
func ID(id datasource.UUID) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id datasource.UUID) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id datasource.UUID) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...datasource.UUID) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...datasource.UUID) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id datasource.UUID) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id datasource.UUID) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id datasource.UUID) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id datasource.UUID) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// EquipmentSn applies equality check predicate on the "equipment_sn" field. It's identical to EquipmentSnEQ.
func EquipmentSn(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEquipmentSn), v))
	})
}

// EvseSerial applies equality check predicate on the "evse_serial" field. It's identical to EvseSerialEQ.
func EvseSerial(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEvseSerial), v))
	})
}

// Serial applies equality check predicate on the "serial" field. It's identical to SerialEQ.
func Serial(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSerial), v))
	})
}

// EquipmentSnEQ applies the EQ predicate on the "equipment_sn" field.
func EquipmentSnEQ(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEquipmentSn), v))
	})
}

// EquipmentSnNEQ applies the NEQ predicate on the "equipment_sn" field.
func EquipmentSnNEQ(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEquipmentSn), v))
	})
}

// EquipmentSnIn applies the In predicate on the "equipment_sn" field.
func EquipmentSnIn(vs ...string) predicate.Connector {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connector(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEquipmentSn), v...))
	})
}

// EquipmentSnNotIn applies the NotIn predicate on the "equipment_sn" field.
func EquipmentSnNotIn(vs ...string) predicate.Connector {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connector(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEquipmentSn), v...))
	})
}

// EquipmentSnGT applies the GT predicate on the "equipment_sn" field.
func EquipmentSnGT(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEquipmentSn), v))
	})
}

// EquipmentSnGTE applies the GTE predicate on the "equipment_sn" field.
func EquipmentSnGTE(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEquipmentSn), v))
	})
}

// EquipmentSnLT applies the LT predicate on the "equipment_sn" field.
func EquipmentSnLT(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEquipmentSn), v))
	})
}

// EquipmentSnLTE applies the LTE predicate on the "equipment_sn" field.
func EquipmentSnLTE(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEquipmentSn), v))
	})
}

// EquipmentSnContains applies the Contains predicate on the "equipment_sn" field.
func EquipmentSnContains(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEquipmentSn), v))
	})
}

// EquipmentSnHasPrefix applies the HasPrefix predicate on the "equipment_sn" field.
func EquipmentSnHasPrefix(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEquipmentSn), v))
	})
}

// EquipmentSnHasSuffix applies the HasSuffix predicate on the "equipment_sn" field.
func EquipmentSnHasSuffix(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEquipmentSn), v))
	})
}

// EquipmentSnEqualFold applies the EqualFold predicate on the "equipment_sn" field.
func EquipmentSnEqualFold(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEquipmentSn), v))
	})
}

// EquipmentSnContainsFold applies the ContainsFold predicate on the "equipment_sn" field.
func EquipmentSnContainsFold(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEquipmentSn), v))
	})
}

// EvseSerialEQ applies the EQ predicate on the "evse_serial" field.
func EvseSerialEQ(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEvseSerial), v))
	})
}

// EvseSerialNEQ applies the NEQ predicate on the "evse_serial" field.
func EvseSerialNEQ(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEvseSerial), v))
	})
}

// EvseSerialIn applies the In predicate on the "evse_serial" field.
func EvseSerialIn(vs ...string) predicate.Connector {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connector(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEvseSerial), v...))
	})
}

// EvseSerialNotIn applies the NotIn predicate on the "evse_serial" field.
func EvseSerialNotIn(vs ...string) predicate.Connector {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connector(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEvseSerial), v...))
	})
}

// EvseSerialGT applies the GT predicate on the "evse_serial" field.
func EvseSerialGT(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEvseSerial), v))
	})
}

// EvseSerialGTE applies the GTE predicate on the "evse_serial" field.
func EvseSerialGTE(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEvseSerial), v))
	})
}

// EvseSerialLT applies the LT predicate on the "evse_serial" field.
func EvseSerialLT(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEvseSerial), v))
	})
}

// EvseSerialLTE applies the LTE predicate on the "evse_serial" field.
func EvseSerialLTE(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEvseSerial), v))
	})
}

// EvseSerialContains applies the Contains predicate on the "evse_serial" field.
func EvseSerialContains(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEvseSerial), v))
	})
}

// EvseSerialHasPrefix applies the HasPrefix predicate on the "evse_serial" field.
func EvseSerialHasPrefix(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEvseSerial), v))
	})
}

// EvseSerialHasSuffix applies the HasSuffix predicate on the "evse_serial" field.
func EvseSerialHasSuffix(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEvseSerial), v))
	})
}

// EvseSerialEqualFold applies the EqualFold predicate on the "evse_serial" field.
func EvseSerialEqualFold(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEvseSerial), v))
	})
}

// EvseSerialContainsFold applies the ContainsFold predicate on the "evse_serial" field.
func EvseSerialContainsFold(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEvseSerial), v))
	})
}

// SerialEQ applies the EQ predicate on the "serial" field.
func SerialEQ(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSerial), v))
	})
}

// SerialNEQ applies the NEQ predicate on the "serial" field.
func SerialNEQ(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSerial), v))
	})
}

// SerialIn applies the In predicate on the "serial" field.
func SerialIn(vs ...string) predicate.Connector {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connector(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSerial), v...))
	})
}

// SerialNotIn applies the NotIn predicate on the "serial" field.
func SerialNotIn(vs ...string) predicate.Connector {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connector(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSerial), v...))
	})
}

// SerialGT applies the GT predicate on the "serial" field.
func SerialGT(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSerial), v))
	})
}

// SerialGTE applies the GTE predicate on the "serial" field.
func SerialGTE(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSerial), v))
	})
}

// SerialLT applies the LT predicate on the "serial" field.
func SerialLT(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSerial), v))
	})
}

// SerialLTE applies the LTE predicate on the "serial" field.
func SerialLTE(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSerial), v))
	})
}

// SerialContains applies the Contains predicate on the "serial" field.
func SerialContains(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSerial), v))
	})
}

// SerialHasPrefix applies the HasPrefix predicate on the "serial" field.
func SerialHasPrefix(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSerial), v))
	})
}

// SerialHasSuffix applies the HasSuffix predicate on the "serial" field.
func SerialHasSuffix(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSerial), v))
	})
}

// SerialEqualFold applies the EqualFold predicate on the "serial" field.
func SerialEqualFold(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSerial), v))
	})
}

// SerialContainsFold applies the ContainsFold predicate on the "serial" field.
func SerialContainsFold(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSerial), v))
	})
}

// CurrentStateEQ applies the EQ predicate on the "current_state" field.
func CurrentStateEQ(v enums.ConnectorState) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCurrentState), v))
	})
}

// CurrentStateNEQ applies the NEQ predicate on the "current_state" field.
func CurrentStateNEQ(v enums.ConnectorState) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCurrentState), v))
	})
}

// CurrentStateIn applies the In predicate on the "current_state" field.
func CurrentStateIn(vs ...enums.ConnectorState) predicate.Connector {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connector(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCurrentState), v...))
	})
}

// CurrentStateNotIn applies the NotIn predicate on the "current_state" field.
func CurrentStateNotIn(vs ...enums.ConnectorState) predicate.Connector {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connector(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCurrentState), v...))
	})
}

// BeforeStateEQ applies the EQ predicate on the "before_state" field.
func BeforeStateEQ(v enums.ConnectorState) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBeforeState), v))
	})
}

// BeforeStateNEQ applies the NEQ predicate on the "before_state" field.
func BeforeStateNEQ(v enums.ConnectorState) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBeforeState), v))
	})
}

// BeforeStateIn applies the In predicate on the "before_state" field.
func BeforeStateIn(vs ...enums.ConnectorState) predicate.Connector {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connector(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBeforeState), v...))
	})
}

// BeforeStateNotIn applies the NotIn predicate on the "before_state" field.
func BeforeStateNotIn(vs ...enums.ConnectorState) predicate.Connector {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connector(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBeforeState), v...))
	})
}

// HasEvse applies the HasEdge predicate on the "evse" edge.
func HasEvse() predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EvseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EvseTable, EvseColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEvseWith applies the HasEdge predicate on the "evse" edge with a given conditions (other predicates).
func HasEvseWith(preds ...predicate.Evse) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EvseInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EvseTable, EvseColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEquipment applies the HasEdge predicate on the "equipment" edge.
func HasEquipment() predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EquipmentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EquipmentTable, EquipmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEquipmentWith applies the HasEdge predicate on the "equipment" edge with a given conditions (other predicates).
func HasEquipmentWith(preds ...predicate.Equipment) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EquipmentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EquipmentTable, EquipmentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Connector) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Connector) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
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
func Not(p predicate.Connector) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		p(s.Not())
	})
}
