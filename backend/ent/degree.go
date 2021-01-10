// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/tanapon395/playlist-video/ent/degree"
)

// Degree is the model entity for the Degree schema.
type Degree struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DegreeQuery when eager-loading is set.
	Edges DegreeEdges `json:"edges"`
}

// DegreeEdges holds the relations/edges for other nodes in the graph.
type DegreeEdges struct {
	// Dentists holds the value of the dentists edge.
	Dentists []*Dentist
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// DentistsOrErr returns the Dentists value or an error if the edge
// was not loaded in eager-loading.
func (e DegreeEdges) DentistsOrErr() ([]*Dentist, error) {
	if e.loadedTypes[0] {
		return e.Dentists, nil
	}
	return nil, &NotLoadedError{edge: "dentists"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Degree) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Degree fields.
func (d *Degree) assignValues(values ...interface{}) error {
	if m, n := len(values), len(degree.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	d.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		d.Name = value.String
	}
	return nil
}

// QueryDentists queries the dentists edge of the Degree.
func (d *Degree) QueryDentists() *DentistQuery {
	return (&DegreeClient{config: d.config}).QueryDentists(d)
}

// Update returns a builder for updating this Degree.
// Note that, you need to call Degree.Unwrap() before calling this method, if this Degree
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Degree) Update() *DegreeUpdateOne {
	return (&DegreeClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (d *Degree) Unwrap() *Degree {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Degree is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Degree) String() string {
	var builder strings.Builder
	builder.WriteString("Degree(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", name=")
	builder.WriteString(d.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Degrees is a parsable slice of Degree.
type Degrees []*Degree

func (d Degrees) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
