// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team03/app/ent/dentalexpense"
	"github.com/team03/app/ent/medicalfile"
	"github.com/team03/app/ent/nurse"
	"github.com/team03/app/ent/pricetype"
)

// Dentalexpense is the model entity for the Dentalexpense schema.
type Dentalexpense struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "Name" field.
	Name string `json:"Name,omitempty"`
	// Phone holds the value of the "Phone" field.
	Phone string `json:"Phone,omitempty"`
	// AddedTime holds the value of the "AddedTime" field.
	AddedTime time.Time `json:"AddedTime,omitempty"`
	// Rates holds the value of the "Rates" field.
	Rates float64 `json:"Rates,omitempty"`
	// Tax holds the value of the "Tax" field.
	Tax string `json:"Tax,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DentalexpenseQuery when eager-loading is set.
	Edges          DentalexpenseEdges `json:"edges"`
	medicalfile_id *int
	nurse_id       *int
	pricetype_id   *int
}

// DentalexpenseEdges holds the relations/edges for other nodes in the graph.
type DentalexpenseEdges struct {
	// Nurse holds the value of the nurse edge.
	Nurse *Nurse
	// Medicalfile holds the value of the medicalfile edge.
	Medicalfile *Medicalfile
	// Pricetype holds the value of the pricetype edge.
	Pricetype *Pricetype
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// NurseOrErr returns the Nurse value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DentalexpenseEdges) NurseOrErr() (*Nurse, error) {
	if e.loadedTypes[0] {
		if e.Nurse == nil {
			// The edge nurse was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: nurse.Label}
		}
		return e.Nurse, nil
	}
	return nil, &NotLoadedError{edge: "nurse"}
}

// MedicalfileOrErr returns the Medicalfile value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DentalexpenseEdges) MedicalfileOrErr() (*Medicalfile, error) {
	if e.loadedTypes[1] {
		if e.Medicalfile == nil {
			// The edge medicalfile was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: medicalfile.Label}
		}
		return e.Medicalfile, nil
	}
	return nil, &NotLoadedError{edge: "medicalfile"}
}

// PricetypeOrErr returns the Pricetype value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DentalexpenseEdges) PricetypeOrErr() (*Pricetype, error) {
	if e.loadedTypes[2] {
		if e.Pricetype == nil {
			// The edge pricetype was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: pricetype.Label}
		}
		return e.Pricetype, nil
	}
	return nil, &NotLoadedError{edge: "pricetype"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Dentalexpense) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},   // id
		&sql.NullString{},  // Name
		&sql.NullString{},  // Phone
		&sql.NullTime{},    // AddedTime
		&sql.NullFloat64{}, // Rates
		&sql.NullString{},  // Tax
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Dentalexpense) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // medicalfile_id
		&sql.NullInt64{}, // nurse_id
		&sql.NullInt64{}, // pricetype_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Dentalexpense fields.
func (d *Dentalexpense) assignValues(values ...interface{}) error {
	if m, n := len(values), len(dentalexpense.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	d.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Name", values[0])
	} else if value.Valid {
		d.Name = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Phone", values[1])
	} else if value.Valid {
		d.Phone = value.String
	}
	if value, ok := values[2].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field AddedTime", values[2])
	} else if value.Valid {
		d.AddedTime = value.Time
	}
	if value, ok := values[3].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field Rates", values[3])
	} else if value.Valid {
		d.Rates = value.Float64
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Tax", values[4])
	} else if value.Valid {
		d.Tax = value.String
	}
	values = values[5:]
	if len(values) == len(dentalexpense.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field medicalfile_id", value)
		} else if value.Valid {
			d.medicalfile_id = new(int)
			*d.medicalfile_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field nurse_id", value)
		} else if value.Valid {
			d.nurse_id = new(int)
			*d.nurse_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field pricetype_id", value)
		} else if value.Valid {
			d.pricetype_id = new(int)
			*d.pricetype_id = int(value.Int64)
		}
	}
	return nil
}

// QueryNurse queries the nurse edge of the Dentalexpense.
func (d *Dentalexpense) QueryNurse() *NurseQuery {
	return (&DentalexpenseClient{config: d.config}).QueryNurse(d)
}

// QueryMedicalfile queries the medicalfile edge of the Dentalexpense.
func (d *Dentalexpense) QueryMedicalfile() *MedicalfileQuery {
	return (&DentalexpenseClient{config: d.config}).QueryMedicalfile(d)
}

// QueryPricetype queries the pricetype edge of the Dentalexpense.
func (d *Dentalexpense) QueryPricetype() *PricetypeQuery {
	return (&DentalexpenseClient{config: d.config}).QueryPricetype(d)
}

// Update returns a builder for updating this Dentalexpense.
// Note that, you need to call Dentalexpense.Unwrap() before calling this method, if this Dentalexpense
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Dentalexpense) Update() *DentalexpenseUpdateOne {
	return (&DentalexpenseClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (d *Dentalexpense) Unwrap() *Dentalexpense {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Dentalexpense is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Dentalexpense) String() string {
	var builder strings.Builder
	builder.WriteString("Dentalexpense(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", Name=")
	builder.WriteString(d.Name)
	builder.WriteString(", Phone=")
	builder.WriteString(d.Phone)
	builder.WriteString(", AddedTime=")
	builder.WriteString(d.AddedTime.Format(time.ANSIC))
	builder.WriteString(", Rates=")
	builder.WriteString(fmt.Sprintf("%v", d.Rates))
	builder.WriteString(", Tax=")
	builder.WriteString(d.Tax)
	builder.WriteByte(')')
	return builder.String()
}

// Dentalexpenses is a parsable slice of Dentalexpense.
type Dentalexpenses []*Dentalexpense

func (d Dentalexpenses) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
