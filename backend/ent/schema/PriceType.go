package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// PriceType holds the schema definition for the PriceType entity.
type PriceType struct {
	ent.Schema
}

// Fields of the PriceType.
func (PriceType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		
	}
}

// Edges of the PriceType.
func (PriceType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("dentalexpenses", DentalExpense.Type).StorageKey(edge.Column("pricetype_id")),
	}
}
