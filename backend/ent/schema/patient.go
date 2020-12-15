package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Patient holds the schema definition for the Patient entity.
type Patient struct {
	ent.Schema
}

// Fields of the Patient.
func (Patient) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("cardID").NotEmpty(),
		field.String("tel").NotEmpty(),
		field.Int("age"),
		field.Time("birthday"),
	}
}

// Edges of the Patient.
func (Patient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("gender", Gender.Type).Ref("patients").Unique(),          
		edge.From("medicalcare", MedicalCare.Type).Ref("patients").Unique(), 
		edge.From("employee", Employee.Type).Ref("patients").Unique(),     
		edge.From("disease", Disease.Type).Ref("patients").Unique(),   
	}