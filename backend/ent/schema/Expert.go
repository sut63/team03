package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Expert holds the schema definition for the Expert entity.
type Expert struct {
	ent.Schema
}

// Fields of the Expert.
func (Expert) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
	}
}

// Edges of the Expert.
func (Expert) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("dentists", Dentist.Type).StorageKey(edge.Column("expert_id")),
	}
}
