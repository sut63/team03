// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/tanapon395/playlist-video/ent/pricetype"
)

// PriceType is the model entity for the PriceType schema.
type PriceType struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PriceTypeQuery when eager-loading is set.
	Edges PriceTypeEdges `json:"edges"`
}

// PriceTypeEdges holds the relations/edges for other nodes in the graph.
type PriceTypeEdges struct {
	// Dentalexpenses holds the value of the dentalexpenses edge.
	Dentalexpenses []*DentalExpense
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// DentalexpensesOrErr returns the Dentalexpenses value or an error if the edge
// was not loaded in eager-loading.
func (e PriceTypeEdges) DentalexpensesOrErr() ([]*DentalExpense, error) {
	if e.loadedTypes[0] {
		return e.Dentalexpenses, nil
	}
	return nil, &NotLoadedError{edge: "dentalexpenses"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PriceType) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PriceType fields.
func (pt *PriceType) assignValues(values ...interface{}) error {
	if m, n := len(values), len(pricetype.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pt.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		pt.Name = value.String
	}
	return nil
}

// QueryDentalexpenses queries the dentalexpenses edge of the PriceType.
func (pt *PriceType) QueryDentalexpenses() *DentalExpenseQuery {
	return (&PriceTypeClient{config: pt.config}).QueryDentalexpenses(pt)
}

// Update returns a builder for updating this PriceType.
// Note that, you need to call PriceType.Unwrap() before calling this method, if this PriceType
// was returned from a transaction, and the transaction was committed or rolled back.
func (pt *PriceType) Update() *PriceTypeUpdateOne {
	return (&PriceTypeClient{config: pt.config}).UpdateOne(pt)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pt *PriceType) Unwrap() *PriceType {
	tx, ok := pt.config.driver.(*txDriver)
	if !ok {
		panic("ent: PriceType is not a transactional entity")
	}
	pt.config.driver = tx.drv
	return pt
}

// String implements the fmt.Stringer.
func (pt *PriceType) String() string {
	var builder strings.Builder
	builder.WriteString("PriceType(")
	builder.WriteString(fmt.Sprintf("id=%v", pt.ID))
	builder.WriteString(", name=")
	builder.WriteString(pt.Name)
	builder.WriteByte(')')
	return builder.String()
}

// PriceTypes is a parsable slice of PriceType.
type PriceTypes []*PriceType

func (pt PriceTypes) config(cfg config) {
	for _i := range pt {
		pt[_i].config = cfg
	}
}
