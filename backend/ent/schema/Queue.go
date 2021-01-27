package schema

import (
	"errors"
	"regexp"
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Queue holds the schema definition for the Queue entity.
type Queue struct {
	ent.Schema
}

// Fields of the Queue.
func (Queue) Fields() []ent.Field {
	return []ent.Field{
        field.String("QueueID").Validate(func(s string) error {
			match, _ := regexp.MatchString("[Q]\\d{3}", s)
			if !match {
				return errors.New("รูปแบบรหัสคิวไม่ถูกต้อง")
			}
			return nil
		}),
		field.String("Phone").MaxLen(10).MinLen(10),
		field.String("Dental").NotEmpty().MaxLen(30),
		field.Time("QueueTime"),
	}
}

// Edges of the Queue.
func (Queue) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("dentist", Dentist.Type).Ref("queue").Unique(),
		edge.From("nurse", Nurse.Type).Ref("queue").Unique(),
		edge.From("patient", Patient.Type).Ref("queue").Unique(),
	}
}
