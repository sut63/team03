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
		field.String("dentist_name").NotEmpty(),
		field.Int("dentist_age").Positive(),
		field.String("dentist_cardid").NotEmpty(),
		field.Time("dentist_birthday"),
		field.String("dentist_experience").NotEmpty(),
		field.String("dentist_tel").NotEmpty(),
		field.String("dentist_email").NotEmpty(),
		field.String("dentist_password").NotEmpty(),

   }
}

// Edges of the Dentist.
func (Dentist) Edges() []ent.Edge {
   return []ent.Edge{
	edge.From("nurse", Nurse.Type).Ref("dentist").Unique(),
	edge.From("degree", Degree.Type).Ref("dentist").Unique(),
	edge.From("expert", Expert.Type).Ref("dentist").Unique(),
	edge.From("gender", Gender.Type).Ref("dentist").Unique(),
   }
} 