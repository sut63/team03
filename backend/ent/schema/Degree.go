package schema

import (
   "github.com/facebookincubator/ent"
   "github.com/facebookincubator/ent/schema/field"
   "github.com/facebookincubator/ent/schema/edge"
)

// Degree holds the schema definition for the Degree entity.
type Degree struct {
   ent.Schema
}

// Fields of the Degree.
func (Degree) Fields() []ent.Field {
   return []ent.Field{
		field.String("name").NotEmpty(),


   }
}

// Edges of the Degree.
func (Degree) Edges() []ent.Edge {
   return []ent.Edge{
   
      edge.To("dentists", Dentist.Type).StorageKey(edge.Column("degree_id")),
   }
} 