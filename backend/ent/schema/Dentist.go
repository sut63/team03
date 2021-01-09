package schema

import (
   "github.com/facebookincubator/ent"
   "github.com/facebookincubator/ent/schema/field"
   "github.com/facebookincubator/ent/schema/edge"
)

// Dentist holds the schema definition for the Dentist entity.
type Dentist struct {
   ent.Schema
}

// Fields of the Dentist.
func (Dentist) Fields() []ent.Field {
   return []ent.Field{
	field.String("name").NotEmpty(),
	field.Int("age").Positive(),
	field.String("cardid").NotEmpty(),
	field.Time("birthday"),
	field.String("experience").NotEmpty(),
	field.String("tel").NotEmpty(),
	field.String("email").NotEmpty(),
	field.String("password").NotEmpty(),

   }
}

// Edges of the Dentist.
func (Dentist) Edges() []ent.Edge {
   return []ent.Edge{
	edge.From("nurse", Nurse.Type).Ref("dentists").Unique(),
	edge.From("degree", Degree.Type).Ref("dentists").Unique(),
	edge.From("expert", Expert.Type).Ref("dentists").Unique(),
	edge.From("gender", Gender.Type).Ref("dentists").Unique(),

	edge.To("medicalfiles", Medicalfile.Type).StorageKey(edge.Column("dentist_id")),
	edge.To("queue", Queue.Type).StorageKey(edge.Column("dentist_id")),
	edge.To("appointment", Appointment.Type).StorageKey(edge.Column("dentist_id")),
	
   }
} 