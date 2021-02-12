package schema

import (
	"errors"
	"regexp"

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
		field.String("AppointID").Validate(func(s string) error {
			match, _ := regexp.MatchString("[A]\\d{5}", s)
			if !match {
				return errors.New("รูปแบบรหัสการนัดหมายไม่ถูกต้อง")
			}
			return nil
		}),
		field.String("Detail").MinLen(5),
		field.Time("Datetime"),
		field.String("Remark").MaxLen(30),
	}
}

// Edges of the Appointment.
func (Appointment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("patient", Patient.Type).Ref("appointment").Unique(),
		edge.From("room", Room.Type).Ref("appointment").Unique(),
		edge.From("dentist", Dentist.Type).Ref("appointment").Unique(),
		edge.From("nurse", Nurse.Type).Ref("appointment").Unique(),
	}
}
