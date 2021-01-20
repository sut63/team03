// Code generated by entc, DO NOT EDIT.

package dentalexpense

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team03/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
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
func IDGT(id int) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Tax applies equality check predicate on the "tax" field. It's identical to TaxEQ.
func Tax(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTax), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Rates applies equality check predicate on the "rates" field. It's identical to RatesEQ.
func Rates(v int) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRates), v))
	})
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhone), v))
	})
}

// AddedTime applies equality check predicate on the "added_time" field. It's identical to AddedTimeEQ.
func AddedTime(v time.Time) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddedTime), v))
	})
}

// TaxEQ applies the EQ predicate on the "tax" field.
func TaxEQ(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTax), v))
	})
}

// TaxNEQ applies the NEQ predicate on the "tax" field.
func TaxNEQ(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTax), v))
	})
}

// TaxIn applies the In predicate on the "tax" field.
func TaxIn(vs ...string) predicate.DentalExpense {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DentalExpense(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTax), v...))
	})
}

// TaxNotIn applies the NotIn predicate on the "tax" field.
func TaxNotIn(vs ...string) predicate.DentalExpense {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DentalExpense(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTax), v...))
	})
}

// TaxGT applies the GT predicate on the "tax" field.
func TaxGT(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTax), v))
	})
}

// TaxGTE applies the GTE predicate on the "tax" field.
func TaxGTE(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTax), v))
	})
}

// TaxLT applies the LT predicate on the "tax" field.
func TaxLT(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTax), v))
	})
}

// TaxLTE applies the LTE predicate on the "tax" field.
func TaxLTE(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTax), v))
	})
}

// TaxContains applies the Contains predicate on the "tax" field.
func TaxContains(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTax), v))
	})
}

// TaxHasPrefix applies the HasPrefix predicate on the "tax" field.
func TaxHasPrefix(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTax), v))
	})
}

// TaxHasSuffix applies the HasSuffix predicate on the "tax" field.
func TaxHasSuffix(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTax), v))
	})
}

// TaxEqualFold applies the EqualFold predicate on the "tax" field.
func TaxEqualFold(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTax), v))
	})
}

// TaxContainsFold applies the ContainsFold predicate on the "tax" field.
func TaxContainsFold(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTax), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.DentalExpense {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DentalExpense(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.DentalExpense {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DentalExpense(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// RatesEQ applies the EQ predicate on the "rates" field.
func RatesEQ(v int) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRates), v))
	})
}

// RatesNEQ applies the NEQ predicate on the "rates" field.
func RatesNEQ(v int) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRates), v))
	})
}

// RatesIn applies the In predicate on the "rates" field.
func RatesIn(vs ...int) predicate.DentalExpense {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DentalExpense(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRates), v...))
	})
}

// RatesNotIn applies the NotIn predicate on the "rates" field.
func RatesNotIn(vs ...int) predicate.DentalExpense {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DentalExpense(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRates), v...))
	})
}

// RatesGT applies the GT predicate on the "rates" field.
func RatesGT(v int) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRates), v))
	})
}

// RatesGTE applies the GTE predicate on the "rates" field.
func RatesGTE(v int) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRates), v))
	})
}

// RatesLT applies the LT predicate on the "rates" field.
func RatesLT(v int) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRates), v))
	})
}

// RatesLTE applies the LTE predicate on the "rates" field.
func RatesLTE(v int) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRates), v))
	})
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhone), v))
	})
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPhone), v))
	})
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.DentalExpense {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DentalExpense(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPhone), v...))
	})
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.DentalExpense {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DentalExpense(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPhone), v...))
	})
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPhone), v))
	})
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPhone), v))
	})
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPhone), v))
	})
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPhone), v))
	})
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPhone), v))
	})
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPhone), v))
	})
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPhone), v))
	})
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPhone), v))
	})
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPhone), v))
	})
}

// AddedTimeEQ applies the EQ predicate on the "added_time" field.
func AddedTimeEQ(v time.Time) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddedTime), v))
	})
}

// AddedTimeNEQ applies the NEQ predicate on the "added_time" field.
func AddedTimeNEQ(v time.Time) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAddedTime), v))
	})
}

// AddedTimeIn applies the In predicate on the "added_time" field.
func AddedTimeIn(vs ...time.Time) predicate.DentalExpense {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DentalExpense(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAddedTime), v...))
	})
}

// AddedTimeNotIn applies the NotIn predicate on the "added_time" field.
func AddedTimeNotIn(vs ...time.Time) predicate.DentalExpense {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DentalExpense(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAddedTime), v...))
	})
}

// AddedTimeGT applies the GT predicate on the "added_time" field.
func AddedTimeGT(v time.Time) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAddedTime), v))
	})
}

// AddedTimeGTE applies the GTE predicate on the "added_time" field.
func AddedTimeGTE(v time.Time) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAddedTime), v))
	})
}

// AddedTimeLT applies the LT predicate on the "added_time" field.
func AddedTimeLT(v time.Time) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAddedTime), v))
	})
}

// AddedTimeLTE applies the LTE predicate on the "added_time" field.
func AddedTimeLTE(v time.Time) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAddedTime), v))
	})
}

// HasNurse applies the HasEdge predicate on the "nurse" edge.
func HasNurse() predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NurseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NurseTable, NurseColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNurseWith applies the HasEdge predicate on the "nurse" edge with a given conditions (other predicates).
func HasNurseWith(preds ...predicate.Nurse) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
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

// HasMedicalfile applies the HasEdge predicate on the "medicalfile" edge.
func HasMedicalfile() predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MedicalfileTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MedicalfileTable, MedicalfileColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMedicalfileWith applies the HasEdge predicate on the "medicalfile" edge with a given conditions (other predicates).
func HasMedicalfileWith(preds ...predicate.Medicalfile) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MedicalfileInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MedicalfileTable, MedicalfileColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPricetype applies the HasEdge predicate on the "pricetype" edge.
func HasPricetype() predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PricetypeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PricetypeTable, PricetypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPricetypeWith applies the HasEdge predicate on the "pricetype" edge with a given conditions (other predicates).
func HasPricetypeWith(preds ...predicate.PriceType) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PricetypeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PricetypeTable, PricetypeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.DentalExpense) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.DentalExpense) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
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
func Not(p predicate.DentalExpense) predicate.DentalExpense {
	return predicate.DentalExpense(func(s *sql.Selector) {
		p(s.Not())
	})
}
