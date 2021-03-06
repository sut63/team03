package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
 )

// Nurse holds the schema definition for the Nurse entity.
type Nurse struct {
	ent.Schema
}

// Fields of the Nurse.
func (Nurse) Fields() []ent.Field {
	return []ent.Field{
		field.String("nurse_name").NotEmpty(),
		field.Int("nurse_age"),
		field.String("email").Unique(),
		field.String("password"),
		
		
	}
}

// Edges of the Nurse.
func (Nurse) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("queue", Queue.Type).StorageKey(edge.Column("nurse_id")),
		edge.To("medicalfiles", Medicalfile.Type).StorageKey(edge.Column("nurse_id")),
		edge.To("dentalexpenses", Dentalexpense.Type).StorageKey(edge.Column("nurse_id")),
		edge.To("patients", Patient.Type).StorageKey(edge.Column("nurse_id")),
		edge.To("dentists", Dentist.Type).StorageKey(edge.Column("nurse_id")),
		edge.To("appointment", Appointment.Type).StorageKey(edge.Column("nurse_id")),
		
	}
}
