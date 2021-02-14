package schema

import (
   "errors"
	"regexp"
   "github.com/facebookincubator/ent"
   "github.com/facebookincubator/ent/schema/field"
   "github.com/facebookincubator/ent/schema/edge"
)

// Dentalexpense holds the schema definition for the Dentalexpense entity.
type Dentalexpense struct {
   ent.Schema
}

// Fields of the Dentalexpense.
func (Dentalexpense) Fields() []ent.Field {
   return []ent.Field{
        field.String("Name").Validate(func(s string) error {
			match, _ := regexp.MatchString("^[ก-๏\\s]+$",s)
			if !match {
				return errors.New("กรุณากรอกเฉพาะภาษาไทยเท่านั้น")
			}
			return nil
		}),
	    field.String("Phone").MaxLen(10).MinLen(10),
	    field.Time("AddedTime"),
		field.Int("Amount").Validate(func(s int) error{
			if (s < 0) {
				return errors.New("Amount have to be Positive")
			}
			return nil
		}),
        field.String("Tax").Validate(func(s string) error {
			match, _ := regexp.MatchString("[R]\\d{10}",s)
			if !match {
				return errors.New("ระบุเลขกำกับภาษีไม่ถูกต้อง")
			}
			return nil
		}),
   }
}

// Edges of the Dentalexpense.
func (Dentalexpense) Edges() []ent.Edge {
   return []ent.Edge{
        edge.From("nurse", Nurse.Type).Ref("dentalexpenses").Unique(),
	    edge.From("medicalfile", Medicalfile.Type).Ref("dentalexpenses").Unique(),
	    edge.From("pricetype", Pricetype.Type).Ref("dentalexpenses").Unique(),
   }
} 





