// Code generated by entc, DO NOT EDIT.

package medicalfile

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team03/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// DrugAllergy applies equality check predicate on the "DrugAllergy" field. It's identical to DrugAllergyEQ.
func DrugAllergy(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDrugAllergy), v))
	})
}

// Detial applies equality check predicate on the "Detial" field. It's identical to DetialEQ.
func Detial(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDetial), v))
	})
}

// AddedTime applies equality check predicate on the "AddedTime" field. It's identical to AddedTimeEQ.
func AddedTime(v time.Time) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddedTime), v))
	})
}

// Medno applies equality check predicate on the "Medno" field. It's identical to MednoEQ.
func Medno(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMedno), v))
	})
}

// DrugAllergyEQ applies the EQ predicate on the "DrugAllergy" field.
func DrugAllergyEQ(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDrugAllergy), v))
	})
}

// DrugAllergyNEQ applies the NEQ predicate on the "DrugAllergy" field.
func DrugAllergyNEQ(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDrugAllergy), v))
	})
}

// DrugAllergyIn applies the In predicate on the "DrugAllergy" field.
func DrugAllergyIn(vs ...string) predicate.Medicalfile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Medicalfile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDrugAllergy), v...))
	})
}

// DrugAllergyNotIn applies the NotIn predicate on the "DrugAllergy" field.
func DrugAllergyNotIn(vs ...string) predicate.Medicalfile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Medicalfile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDrugAllergy), v...))
	})
}

// DrugAllergyGT applies the GT predicate on the "DrugAllergy" field.
func DrugAllergyGT(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDrugAllergy), v))
	})
}

// DrugAllergyGTE applies the GTE predicate on the "DrugAllergy" field.
func DrugAllergyGTE(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDrugAllergy), v))
	})
}

// DrugAllergyLT applies the LT predicate on the "DrugAllergy" field.
func DrugAllergyLT(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDrugAllergy), v))
	})
}

// DrugAllergyLTE applies the LTE predicate on the "DrugAllergy" field.
func DrugAllergyLTE(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDrugAllergy), v))
	})
}

// DrugAllergyContains applies the Contains predicate on the "DrugAllergy" field.
func DrugAllergyContains(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDrugAllergy), v))
	})
}

// DrugAllergyHasPrefix applies the HasPrefix predicate on the "DrugAllergy" field.
func DrugAllergyHasPrefix(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDrugAllergy), v))
	})
}

// DrugAllergyHasSuffix applies the HasSuffix predicate on the "DrugAllergy" field.
func DrugAllergyHasSuffix(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDrugAllergy), v))
	})
}

// DrugAllergyEqualFold applies the EqualFold predicate on the "DrugAllergy" field.
func DrugAllergyEqualFold(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDrugAllergy), v))
	})
}

// DrugAllergyContainsFold applies the ContainsFold predicate on the "DrugAllergy" field.
func DrugAllergyContainsFold(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDrugAllergy), v))
	})
}

// DetialEQ applies the EQ predicate on the "Detial" field.
func DetialEQ(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDetial), v))
	})
}

// DetialNEQ applies the NEQ predicate on the "Detial" field.
func DetialNEQ(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDetial), v))
	})
}

// DetialIn applies the In predicate on the "Detial" field.
func DetialIn(vs ...string) predicate.Medicalfile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Medicalfile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDetial), v...))
	})
}

// DetialNotIn applies the NotIn predicate on the "Detial" field.
func DetialNotIn(vs ...string) predicate.Medicalfile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Medicalfile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDetial), v...))
	})
}

// DetialGT applies the GT predicate on the "Detial" field.
func DetialGT(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDetial), v))
	})
}

// DetialGTE applies the GTE predicate on the "Detial" field.
func DetialGTE(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDetial), v))
	})
}

// DetialLT applies the LT predicate on the "Detial" field.
func DetialLT(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDetial), v))
	})
}

// DetialLTE applies the LTE predicate on the "Detial" field.
func DetialLTE(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDetial), v))
	})
}

// DetialContains applies the Contains predicate on the "Detial" field.
func DetialContains(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDetial), v))
	})
}

// DetialHasPrefix applies the HasPrefix predicate on the "Detial" field.
func DetialHasPrefix(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDetial), v))
	})
}

// DetialHasSuffix applies the HasSuffix predicate on the "Detial" field.
func DetialHasSuffix(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDetial), v))
	})
}

// DetialEqualFold applies the EqualFold predicate on the "Detial" field.
func DetialEqualFold(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDetial), v))
	})
}

// DetialContainsFold applies the ContainsFold predicate on the "Detial" field.
func DetialContainsFold(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDetial), v))
	})
}

// AddedTimeEQ applies the EQ predicate on the "AddedTime" field.
func AddedTimeEQ(v time.Time) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddedTime), v))
	})
}

// AddedTimeNEQ applies the NEQ predicate on the "AddedTime" field.
func AddedTimeNEQ(v time.Time) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAddedTime), v))
	})
}

// AddedTimeIn applies the In predicate on the "AddedTime" field.
func AddedTimeIn(vs ...time.Time) predicate.Medicalfile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Medicalfile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAddedTime), v...))
	})
}

// AddedTimeNotIn applies the NotIn predicate on the "AddedTime" field.
func AddedTimeNotIn(vs ...time.Time) predicate.Medicalfile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Medicalfile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAddedTime), v...))
	})
}

// AddedTimeGT applies the GT predicate on the "AddedTime" field.
func AddedTimeGT(v time.Time) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAddedTime), v))
	})
}

// AddedTimeGTE applies the GTE predicate on the "AddedTime" field.
func AddedTimeGTE(v time.Time) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAddedTime), v))
	})
}

// AddedTimeLT applies the LT predicate on the "AddedTime" field.
func AddedTimeLT(v time.Time) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAddedTime), v))
	})
}

// AddedTimeLTE applies the LTE predicate on the "AddedTime" field.
func AddedTimeLTE(v time.Time) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAddedTime), v))
	})
}

// MednoEQ applies the EQ predicate on the "Medno" field.
func MednoEQ(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMedno), v))
	})
}

// MednoNEQ applies the NEQ predicate on the "Medno" field.
func MednoNEQ(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMedno), v))
	})
}

// MednoIn applies the In predicate on the "Medno" field.
func MednoIn(vs ...string) predicate.Medicalfile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Medicalfile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMedno), v...))
	})
}

// MednoNotIn applies the NotIn predicate on the "Medno" field.
func MednoNotIn(vs ...string) predicate.Medicalfile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Medicalfile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMedno), v...))
	})
}

// MednoGT applies the GT predicate on the "Medno" field.
func MednoGT(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMedno), v))
	})
}

// MednoGTE applies the GTE predicate on the "Medno" field.
func MednoGTE(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMedno), v))
	})
}

// MednoLT applies the LT predicate on the "Medno" field.
func MednoLT(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMedno), v))
	})
}

// MednoLTE applies the LTE predicate on the "Medno" field.
func MednoLTE(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMedno), v))
	})
}

// MednoContains applies the Contains predicate on the "Medno" field.
func MednoContains(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMedno), v))
	})
}

// MednoHasPrefix applies the HasPrefix predicate on the "Medno" field.
func MednoHasPrefix(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMedno), v))
	})
}

// MednoHasSuffix applies the HasSuffix predicate on the "Medno" field.
func MednoHasSuffix(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMedno), v))
	})
}

// MednoEqualFold applies the EqualFold predicate on the "Medno" field.
func MednoEqualFold(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMedno), v))
	})
}

// MednoContainsFold applies the ContainsFold predicate on the "Medno" field.
func MednoContainsFold(v string) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMedno), v))
	})
}

// HasDentist applies the HasEdge predicate on the "dentist" edge.
func HasDentist() predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DentistTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DentistTable, DentistColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDentistWith applies the HasEdge predicate on the "dentist" edge with a given conditions (other predicates).
func HasDentistWith(preds ...predicate.Dentist) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
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

// HasPatient applies the HasEdge predicate on the "patient" edge.
func HasPatient() predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PatientTable, PatientColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPatientWith applies the HasEdge predicate on the "patient" edge with a given conditions (other predicates).
func HasPatientWith(preds ...predicate.Patient) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
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

// HasNurse applies the HasEdge predicate on the "nurse" edge.
func HasNurse() predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NurseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NurseTable, NurseColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNurseWith applies the HasEdge predicate on the "nurse" edge with a given conditions (other predicates).
func HasNurseWith(preds ...predicate.Nurse) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
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

// HasDentalexpenses applies the HasEdge predicate on the "dentalexpenses" edge.
func HasDentalexpenses() predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DentalexpensesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DentalexpensesTable, DentalexpensesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDentalexpensesWith applies the HasEdge predicate on the "dentalexpenses" edge with a given conditions (other predicates).
func HasDentalexpensesWith(preds ...predicate.Dentalexpense) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DentalexpensesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DentalexpensesTable, DentalexpensesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Medicalfile) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Medicalfile) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
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
func Not(p predicate.Medicalfile) predicate.Medicalfile {
	return predicate.Medicalfile(func(s *sql.Selector) {
		p(s.Not())
	})
}
