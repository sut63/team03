package schema

import (
    "github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/field"
    "github.com/facebookincubator/ent/schema/edge"
)

// Disease holds the schema definition for the Disease entity.
type Disease struct {
    ent.Schema
}

// Fields of the Disease.
func (Disease) Fields() []ent.Field {
    return []ent.Field{
        field.String("name").NotEmpty(),
    }
}

// Edges of the Disease.
func (Disease) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("patients", Patient.Type).StorageKey(edge.Column("disease_id")),
    }
}