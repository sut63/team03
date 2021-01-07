package schema

import (
    "github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/field"
    "github.com/facebookincubator/ent/schema/edge"
)

// MedicalCare holds the schema definition for the MedicalCare entity.
type MedicalCare struct {
    ent.Schema
}

// Fields of the MedicalCare.
func (MedicalCare) Fields() []ent.Field {
    return []ent.Field{
        field.String("name").NotEmpty(),
    }
}

// Edges of the MedicalCare.
func (MedicalCare) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("patients", Patient.Type).StorageKey(edge.Column("medicalcare_id")),
    }
}