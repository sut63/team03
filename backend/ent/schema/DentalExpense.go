package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// DentalExpense holds the schema definition for the DentalExpense entity.
type DentalExpense struct {
	ent.Schema
}

// Fields of the DentalExpense.
func (DentalExpense) Fields() []ent.Field {
	return []ent.Field{
		field.String("tax").NotEmpty().Unique(),
		field.String("name").NotEmpty(),
		field.Int("rates").Positive(),
		field.String("phone").NotEmpty(),
		field.Time("added_time"),
	}
}

// Edges of the DentalExpense.
func (DentalExpense) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("nurse", Nurse.Type).Ref("dentalexpenses").Unique(),
		edge.From("medicalfile", Medicalfile.Type).Ref("dentalexpenses").Unique(),
		edge.From("pricetype", PriceType.Type).Ref("dentalexpenses").Unique(),
	}
}
