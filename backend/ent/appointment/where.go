// Code generated by entc, DO NOT EDIT.

package appointment

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team03/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// AppointID applies equality check predicate on the "AppointID" field. It's identical to AppointIDEQ.
func AppointID(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppointID), v))
	})
}

// Detail applies equality check predicate on the "Detail" field. It's identical to DetailEQ.
func Detail(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDetail), v))
	})
}

// Datetime applies equality check predicate on the "Datetime" field. It's identical to DatetimeEQ.
func Datetime(v time.Time) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDatetime), v))
	})
}

// Remark applies equality check predicate on the "Remark" field. It's identical to RemarkEQ.
func Remark(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRemark), v))
	})
}

// AppointIDEQ applies the EQ predicate on the "AppointID" field.
func AppointIDEQ(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppointID), v))
	})
}

// AppointIDNEQ applies the NEQ predicate on the "AppointID" field.
func AppointIDNEQ(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppointID), v))
	})
}

// AppointIDIn applies the In predicate on the "AppointID" field.
func AppointIDIn(vs ...string) predicate.Appointment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Appointment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAppointID), v...))
	})
}

// AppointIDNotIn applies the NotIn predicate on the "AppointID" field.
func AppointIDNotIn(vs ...string) predicate.Appointment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Appointment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAppointID), v...))
	})
}

// AppointIDGT applies the GT predicate on the "AppointID" field.
func AppointIDGT(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppointID), v))
	})
}

// AppointIDGTE applies the GTE predicate on the "AppointID" field.
func AppointIDGTE(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppointID), v))
	})
}

// AppointIDLT applies the LT predicate on the "AppointID" field.
func AppointIDLT(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppointID), v))
	})
}

// AppointIDLTE applies the LTE predicate on the "AppointID" field.
func AppointIDLTE(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppointID), v))
	})
}

// AppointIDContains applies the Contains predicate on the "AppointID" field.
func AppointIDContains(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAppointID), v))
	})
}

// AppointIDHasPrefix applies the HasPrefix predicate on the "AppointID" field.
func AppointIDHasPrefix(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAppointID), v))
	})
}

// AppointIDHasSuffix applies the HasSuffix predicate on the "AppointID" field.
func AppointIDHasSuffix(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAppointID), v))
	})
}

// AppointIDEqualFold applies the EqualFold predicate on the "AppointID" field.
func AppointIDEqualFold(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAppointID), v))
	})
}

// AppointIDContainsFold applies the ContainsFold predicate on the "AppointID" field.
func AppointIDContainsFold(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAppointID), v))
	})
}

// DetailEQ applies the EQ predicate on the "Detail" field.
func DetailEQ(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDetail), v))
	})
}

// DetailNEQ applies the NEQ predicate on the "Detail" field.
func DetailNEQ(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDetail), v))
	})
}

// DetailIn applies the In predicate on the "Detail" field.
func DetailIn(vs ...string) predicate.Appointment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Appointment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDetail), v...))
	})
}

// DetailNotIn applies the NotIn predicate on the "Detail" field.
func DetailNotIn(vs ...string) predicate.Appointment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Appointment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDetail), v...))
	})
}

// DetailGT applies the GT predicate on the "Detail" field.
func DetailGT(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDetail), v))
	})
}

// DetailGTE applies the GTE predicate on the "Detail" field.
func DetailGTE(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDetail), v))
	})
}

// DetailLT applies the LT predicate on the "Detail" field.
func DetailLT(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDetail), v))
	})
}

// DetailLTE applies the LTE predicate on the "Detail" field.
func DetailLTE(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDetail), v))
	})
}

// DetailContains applies the Contains predicate on the "Detail" field.
func DetailContains(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDetail), v))
	})
}

// DetailHasPrefix applies the HasPrefix predicate on the "Detail" field.
func DetailHasPrefix(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDetail), v))
	})
}

// DetailHasSuffix applies the HasSuffix predicate on the "Detail" field.
func DetailHasSuffix(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDetail), v))
	})
}

// DetailEqualFold applies the EqualFold predicate on the "Detail" field.
func DetailEqualFold(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDetail), v))
	})
}

// DetailContainsFold applies the ContainsFold predicate on the "Detail" field.
func DetailContainsFold(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDetail), v))
	})
}

// DatetimeEQ applies the EQ predicate on the "Datetime" field.
func DatetimeEQ(v time.Time) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDatetime), v))
	})
}

// DatetimeNEQ applies the NEQ predicate on the "Datetime" field.
func DatetimeNEQ(v time.Time) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDatetime), v))
	})
}

// DatetimeIn applies the In predicate on the "Datetime" field.
func DatetimeIn(vs ...time.Time) predicate.Appointment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Appointment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDatetime), v...))
	})
}

// DatetimeNotIn applies the NotIn predicate on the "Datetime" field.
func DatetimeNotIn(vs ...time.Time) predicate.Appointment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Appointment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDatetime), v...))
	})
}

// DatetimeGT applies the GT predicate on the "Datetime" field.
func DatetimeGT(v time.Time) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDatetime), v))
	})
}

// DatetimeGTE applies the GTE predicate on the "Datetime" field.
func DatetimeGTE(v time.Time) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDatetime), v))
	})
}

// DatetimeLT applies the LT predicate on the "Datetime" field.
func DatetimeLT(v time.Time) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDatetime), v))
	})
}

// DatetimeLTE applies the LTE predicate on the "Datetime" field.
func DatetimeLTE(v time.Time) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDatetime), v))
	})
}

// RemarkEQ applies the EQ predicate on the "Remark" field.
func RemarkEQ(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRemark), v))
	})
}

// RemarkNEQ applies the NEQ predicate on the "Remark" field.
func RemarkNEQ(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRemark), v))
	})
}

// RemarkIn applies the In predicate on the "Remark" field.
func RemarkIn(vs ...string) predicate.Appointment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Appointment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRemark), v...))
	})
}

// RemarkNotIn applies the NotIn predicate on the "Remark" field.
func RemarkNotIn(vs ...string) predicate.Appointment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Appointment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRemark), v...))
	})
}

// RemarkGT applies the GT predicate on the "Remark" field.
func RemarkGT(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRemark), v))
	})
}

// RemarkGTE applies the GTE predicate on the "Remark" field.
func RemarkGTE(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRemark), v))
	})
}

// RemarkLT applies the LT predicate on the "Remark" field.
func RemarkLT(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRemark), v))
	})
}

// RemarkLTE applies the LTE predicate on the "Remark" field.
func RemarkLTE(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRemark), v))
	})
}

// RemarkContains applies the Contains predicate on the "Remark" field.
func RemarkContains(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRemark), v))
	})
}

// RemarkHasPrefix applies the HasPrefix predicate on the "Remark" field.
func RemarkHasPrefix(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRemark), v))
	})
}

// RemarkHasSuffix applies the HasSuffix predicate on the "Remark" field.
func RemarkHasSuffix(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRemark), v))
	})
}

// RemarkEqualFold applies the EqualFold predicate on the "Remark" field.
func RemarkEqualFold(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRemark), v))
	})
}

// RemarkContainsFold applies the ContainsFold predicate on the "Remark" field.
func RemarkContainsFold(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRemark), v))
	})
}

// HasPatient applies the HasEdge predicate on the "patient" edge.
func HasPatient() predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PatientTable, PatientColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPatientWith applies the HasEdge predicate on the "patient" edge with a given conditions (other predicates).
func HasPatientWith(preds ...predicate.Patient) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PatientTable, PatientColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRoom applies the HasEdge predicate on the "room" edge.
func HasRoom() predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoomTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RoomTable, RoomColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoomWith applies the HasEdge predicate on the "room" edge with a given conditions (other predicates).
func HasRoomWith(preds ...predicate.Room) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoomInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RoomTable, RoomColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDentist applies the HasEdge predicate on the "dentist" edge.
func HasDentist() predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DentistTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DentistTable, DentistColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDentistWith applies the HasEdge predicate on the "dentist" edge with a given conditions (other predicates).
func HasDentistWith(preds ...predicate.Dentist) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DentistInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DentistTable, DentistColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNurse applies the HasEdge predicate on the "nurse" edge.
func HasNurse() predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NurseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NurseTable, NurseColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNurseWith applies the HasEdge predicate on the "nurse" edge with a given conditions (other predicates).
func HasNurseWith(preds ...predicate.Nurse) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NurseInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NurseTable, NurseColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Appointment) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Appointment) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
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
func Not(p predicate.Appointment) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		p(s.Not())
	})
}
