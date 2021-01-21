package schema

import (
   "errors"
	"regexp"
   "github.com/facebookincubator/ent"
   "github.com/facebookincubator/ent/schema/field"
   "github.com/facebookincubator/ent/schema/edge"
)

// Medicalfile holds the schema definition for the Medicalfile entity.
type Medicalfile struct {
   ent.Schema
}

// Fields of the Medicalfile.
func (Medicalfile) Fields() []ent.Field {
   return []ent.Field{
      field.String("DrugAllergy").NotEmpty().MaxLen(30),
		field.String("Detial").NotEmpty().MaxLen(30),
      field.Time("AddedTime"),
      field.String("Medno").Validate(func(s string) error {
			match, _ := regexp.MatchString("[M]\\d{3}", s)
			if !match {
				return errors.New("รูปแบบรหัสประวัติทันตกรรมไม่ถูกต้อง")
			}
			return nil
		}),
   }
}

// Edges of the Medicalfile.
func (Medicalfile) Edges() []ent.Edge {
   return []ent.Edge{
   edge.From("dentist", Dentist.Type).Ref("medicalfiles").Unique(),
   edge.From("patient", Patient.Type).Ref("medicalfiles").Unique(),
   edge.From("nurse", Nurse.Type).Ref("medicalfiles").Unique(),
   edge.To("dentalexpenses", Dentalexpense.Type).StorageKey(edge.Column("medicalfile_id")),
   }
} 