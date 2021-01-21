package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Pricetype holds the schema definition for the Pricetype entity.
type Pricetype struct {
	ent.Schema
}

// Fields of the Pricetype.
func (Pricetype) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Pricetype.
func (Pricetype) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("dentalexpenses", Dentalexpense.Type).StorageKey(edge.Column("pricetype_id")),
	}
}
