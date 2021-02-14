package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"errors"
	"regexp"
)

// Patient holds the schema definition for the Patient entity.
type Patient struct {
	ent.Schema
}

// Fields of the Patient.
func (Patient) Fields() []ent.Field {
	return []ent.Field{
		field.String("PatientID").Validate(func(s string) error {
			match, _ := regexp.MatchString("[P]\\d{6}", s)
			if !match {
				return errors.New("รูปแบบรหัสประจำตัวผู้ป่วยไม่ถูกต้อง")
			}
			return nil
		}),
		field.String("Name").NotEmpty(),
		field.String("CardID").Validate(func(s string) error {
			match, _ := regexp.MatchString("^[0-9]{13}", s)
			if !match {
				return errors.New("รูปแบบเลขบัตรประชาชน 13 หลักไม่ถูกต้อง")
			}
			return nil
		}),
		//field.String("CardID").MaxLen(13).MinLen(13),
		field.String("Tel").Validate(func(s string) error {
			match, _ := regexp.MatchString("[0]\\d{9}", s)
			if !match {
				return errors.New("รูปแบบหมายเลขโทรศัพท์ 10 หลักไม่ถูกต้อง")
			}
			return nil
			}),
		//field.String("Tel").MaxLen(10).MinLen(10),
		field.Int("Age").Range(1,200),
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
