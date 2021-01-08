package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// DentalExpense holds the schema definition for the DentalExpense entity.
type DentalExpense struct {
	ent.Schema
}

// Fields of the Dental_Expense.
func (DentalExpense) Fields() []ent.Field {
	return []ent.Field{
		field.Time("added_time"),
	}
}

// Edges of the Dental_Expense.
func (DentalExpense) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("nurse", Nurse.Type).Ref("dentalexpenses").Unique(),
		edge.From("medicalfile", Medicalfile.Type).Ref("dentalexpenses").Unique(),
		edge.From("pricetype", PriceType.Type).Ref("dentalexpenses").Unique(),
		
	}
}