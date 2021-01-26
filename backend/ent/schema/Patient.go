package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	//"errors"
	//"regexp"
)

// Patient holds the schema definition for the Patient entity.
type Patient struct {
	ent.Schema
}

// Fields of the Patient.
func (Patient) Fields() []ent.Field {
	return []ent.Field{
		field.String("PatientID").Unique().MaxLen(6).MinLen(6),
		field.String("Name").NotEmpty(),
		field.String("CardID").Unique().MaxLen(13).MinLen(13),
		field.String("Tel").Unique().MaxLen(10).MinLen(10),
		field.Int("Age").Min(0),
		field.Time("Birthday"),
	}
}

// Edges of the Patient.
func (Patient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("gender", Gender.Type).Ref("patients").Unique(),
		edge.From("medicalcare", MedicalCare.Type).Ref("patients").Unique(),
		edge.From("nurse", Nurse.Type).Ref("patients").Unique(), 
		edge.From("disease", Disease.Type).Ref("patients").Unique(),

		edge.To("medicalfiles", Medicalfile.Type).StorageKey(edge.Column("patient_id")),
		edge.To("queue", Queue.Type).StorageKey(edge.Column("patient_id")),
		edge.To("appointment", Appointment.Type).StorageKey(edge.Column("patient_id")),
		
		
	}

}
