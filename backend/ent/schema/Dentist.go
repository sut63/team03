package schema

import (
	"errors"
	"regexp"
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
	field.Int("age").Min(0),

	field.String("cardid").Validate(func(s string) error {
	match, _ := regexp.MatchString("[0123456789]\\d{12}", s)
	if !match {
		return errors.New("รูปแบบเลขบัตรประชาชน 13 หลัก ผิดพลาด")
	}
	return nil
	}),
	field.Time("birthday"),
	field.String("experience").NotEmpty(),

    field.String("tel").Validate(func(a string) error {
	match, _ := regexp.MatchString("[0]\\d{9}", a)
	if !match {
		return errors.New("รูปแบบหมายเลขโทรศัพท์ ผิดพลาด")
	}
	return nil
	}),

    field.String("email").Validate(func(c string) error {
	match, _ := regexp.MatchString("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$", c)
	if !match {
		return errors.New("รูปแบบอีเมล ผิดพลาด")
	}
	return nil
	}),
	field.String("password").NotEmpty(),
	
	//	field.String("cardid").Match(regexp.MustCompile("[0123456789]\\d{12}")),
	//field.String("email").Match(regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")),
	//	field.String("tel").Match(regexp.MustCompile("[0]\\d{9}")),
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