package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Appointment holds the schema definition for the Appointment entity.
type Appointment struct {
	ent.Schema
}

// Fields of the Appointment.
func (Appointment) Fields() []ent.Field {
	return []ent.Field{
		field.String("detail").NotEmpty(),
		field.Time("datetime"),
	}
}

// Edges of the Appointment.
func (Appointment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("patient", Patient.Type).Ref("appointment").Unique(),
		edge.From("room", Room.Type).Ref("appointment").Unique(),
		edge.From("dentist", Dentist.Type).Ref("appointment").Unique(),
	}
}
