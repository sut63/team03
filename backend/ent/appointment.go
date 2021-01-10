// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team03/app/ent/appointment"
	"github.com/team03/app/ent/dentist"
	"github.com/team03/app/ent/patient"
	"github.com/team03/app/ent/room"
)

// Appointment is the model entity for the Appointment schema.
type Appointment struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Detail holds the value of the "detail" field.
	Detail string `json:"detail,omitempty"`
	// Datetime holds the value of the "datetime" field.
	Datetime time.Time `json:"datetime,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AppointmentQuery when eager-loading is set.
	Edges      AppointmentEdges `json:"edges"`
	dentist_id *int
	patient_id *int
	room_id    *int
}

// AppointmentEdges holds the relations/edges for other nodes in the graph.
type AppointmentEdges struct {
	// Patient holds the value of the patient edge.
	Patient *Patient
	// Room holds the value of the room edge.
	Room *Room
	// Dentist holds the value of the dentist edge.
	Dentist *Dentist
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// PatientOrErr returns the Patient value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AppointmentEdges) PatientOrErr() (*Patient, error) {
	if e.loadedTypes[0] {
		if e.Patient == nil {
			// The edge patient was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: patient.Label}
		}
		return e.Patient, nil
	}
	return nil, &NotLoadedError{edge: "patient"}
}

// RoomOrErr returns the Room value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AppointmentEdges) RoomOrErr() (*Room, error) {
	if e.loadedTypes[1] {
		if e.Room == nil {
			// The edge room was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: room.Label}
		}
		return e.Room, nil
	}
	return nil, &NotLoadedError{edge: "room"}
}

// DentistOrErr returns the Dentist value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AppointmentEdges) DentistOrErr() (*Dentist, error) {
	if e.loadedTypes[2] {
		if e.Dentist == nil {
			// The edge dentist was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: dentist.Label}
		}
		return e.Dentist, nil
	}
	return nil, &NotLoadedError{edge: "dentist"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Appointment) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // detail
		&sql.NullTime{},   // datetime
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Appointment) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // dentist_id
		&sql.NullInt64{}, // patient_id
		&sql.NullInt64{}, // room_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Appointment fields.
func (a *Appointment) assignValues(values ...interface{}) error {
	if m, n := len(values), len(appointment.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	a.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field detail", values[0])
	} else if value.Valid {
		a.Detail = value.String
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field datetime", values[1])
	} else if value.Valid {
		a.Datetime = value.Time
	}
	values = values[2:]
	if len(values) == len(appointment.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field dentist_id", value)
		} else if value.Valid {
			a.dentist_id = new(int)
			*a.dentist_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field patient_id", value)
		} else if value.Valid {
			a.patient_id = new(int)
			*a.patient_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field room_id", value)
		} else if value.Valid {
			a.room_id = new(int)
			*a.room_id = int(value.Int64)
		}
	}
	return nil
}

// QueryPatient queries the patient edge of the Appointment.
func (a *Appointment) QueryPatient() *PatientQuery {
	return (&AppointmentClient{config: a.config}).QueryPatient(a)
}

// QueryRoom queries the room edge of the Appointment.
func (a *Appointment) QueryRoom() *RoomQuery {
	return (&AppointmentClient{config: a.config}).QueryRoom(a)
}

// QueryDentist queries the dentist edge of the Appointment.
func (a *Appointment) QueryDentist() *DentistQuery {
	return (&AppointmentClient{config: a.config}).QueryDentist(a)
}

// Update returns a builder for updating this Appointment.
// Note that, you need to call Appointment.Unwrap() before calling this method, if this Appointment
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Appointment) Update() *AppointmentUpdateOne {
	return (&AppointmentClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (a *Appointment) Unwrap() *Appointment {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Appointment is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Appointment) String() string {
	var builder strings.Builder
	builder.WriteString("Appointment(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", detail=")
	builder.WriteString(a.Detail)
	builder.WriteString(", datetime=")
	builder.WriteString(a.Datetime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Appointments is a parsable slice of Appointment.
type Appointments []*Appointment

func (a Appointments) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
