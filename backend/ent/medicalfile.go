// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/tanapon395/playlist-video/ent/dentist"
	"github.com/tanapon395/playlist-video/ent/medicalcare"
	"github.com/tanapon395/playlist-video/ent/medicalfile"
	"github.com/tanapon395/playlist-video/ent/nurse"
	"github.com/tanapon395/playlist-video/ent/patient"
)

// Medicalfile is the model entity for the Medicalfile schema.
type Medicalfile struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Detail holds the value of the "detail" field.
	Detail string `json:"detail,omitempty"`
	// AddedTime holds the value of the "added_time" field.
	AddedTime time.Time `json:"added_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MedicalfileQuery when eager-loading is set.
	Edges          MedicalfileEdges `json:"edges"`
	dentist_id     *int
	medicalcare_id *int
	nurse_id       *int
	patient_id     *int
}

// MedicalfileEdges holds the relations/edges for other nodes in the graph.
type MedicalfileEdges struct {
	// Dentist holds the value of the dentist edge.
	Dentist *Dentist
	// Patient holds the value of the patient edge.
	Patient *Patient
	// Nurse holds the value of the nurse edge.
	Nurse *Nurse
	// Medicalcare holds the value of the medicalcare edge.
	Medicalcare *MedicalCare
	// Dentalexpenses holds the value of the dentalexpenses edge.
	Dentalexpenses []*DentalExpense
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// DentistOrErr returns the Dentist value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MedicalfileEdges) DentistOrErr() (*Dentist, error) {
	if e.loadedTypes[0] {
		if e.Dentist == nil {
			// The edge dentist was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: dentist.Label}
		}
		return e.Dentist, nil
	}
	return nil, &NotLoadedError{edge: "dentist"}
}

// PatientOrErr returns the Patient value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MedicalfileEdges) PatientOrErr() (*Patient, error) {
	if e.loadedTypes[1] {
		if e.Patient == nil {
			// The edge patient was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: patient.Label}
		}
		return e.Patient, nil
	}
	return nil, &NotLoadedError{edge: "patient"}
}

// NurseOrErr returns the Nurse value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MedicalfileEdges) NurseOrErr() (*Nurse, error) {
	if e.loadedTypes[2] {
		if e.Nurse == nil {
			// The edge nurse was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: nurse.Label}
		}
		return e.Nurse, nil
	}
	return nil, &NotLoadedError{edge: "nurse"}
}

// MedicalcareOrErr returns the Medicalcare value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MedicalfileEdges) MedicalcareOrErr() (*MedicalCare, error) {
	if e.loadedTypes[3] {
		if e.Medicalcare == nil {
			// The edge medicalcare was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: medicalcare.Label}
		}
		return e.Medicalcare, nil
	}
	return nil, &NotLoadedError{edge: "medicalcare"}
}

// DentalexpensesOrErr returns the Dentalexpenses value or an error if the edge
// was not loaded in eager-loading.
func (e MedicalfileEdges) DentalexpensesOrErr() ([]*DentalExpense, error) {
	if e.loadedTypes[4] {
		return e.Dentalexpenses, nil
	}
	return nil, &NotLoadedError{edge: "dentalexpenses"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Medicalfile) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // detail
		&sql.NullTime{},   // added_time
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Medicalfile) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // dentist_id
		&sql.NullInt64{}, // medicalcare_id
		&sql.NullInt64{}, // nurse_id
		&sql.NullInt64{}, // patient_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Medicalfile fields.
func (m *Medicalfile) assignValues(values ...interface{}) error {
	if m, n := len(values), len(medicalfile.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	m.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field detail", values[0])
	} else if value.Valid {
		m.Detail = value.String
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field added_time", values[1])
	} else if value.Valid {
		m.AddedTime = value.Time
	}
	values = values[2:]
	if len(values) == len(medicalfile.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field dentist_id", value)
		} else if value.Valid {
			m.dentist_id = new(int)
			*m.dentist_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field medicalcare_id", value)
		} else if value.Valid {
			m.medicalcare_id = new(int)
			*m.medicalcare_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field nurse_id", value)
		} else if value.Valid {
			m.nurse_id = new(int)
			*m.nurse_id = int(value.Int64)
		}
		if value, ok := values[3].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field patient_id", value)
		} else if value.Valid {
			m.patient_id = new(int)
			*m.patient_id = int(value.Int64)
		}
	}
	return nil
}

// QueryDentist queries the dentist edge of the Medicalfile.
func (m *Medicalfile) QueryDentist() *DentistQuery {
	return (&MedicalfileClient{config: m.config}).QueryDentist(m)
}

// QueryPatient queries the patient edge of the Medicalfile.
func (m *Medicalfile) QueryPatient() *PatientQuery {
	return (&MedicalfileClient{config: m.config}).QueryPatient(m)
}

// QueryNurse queries the nurse edge of the Medicalfile.
func (m *Medicalfile) QueryNurse() *NurseQuery {
	return (&MedicalfileClient{config: m.config}).QueryNurse(m)
}

// QueryMedicalcare queries the medicalcare edge of the Medicalfile.
func (m *Medicalfile) QueryMedicalcare() *MedicalCareQuery {
	return (&MedicalfileClient{config: m.config}).QueryMedicalcare(m)
}

// QueryDentalexpenses queries the dentalexpenses edge of the Medicalfile.
func (m *Medicalfile) QueryDentalexpenses() *DentalExpenseQuery {
	return (&MedicalfileClient{config: m.config}).QueryDentalexpenses(m)
}

// Update returns a builder for updating this Medicalfile.
// Note that, you need to call Medicalfile.Unwrap() before calling this method, if this Medicalfile
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Medicalfile) Update() *MedicalfileUpdateOne {
	return (&MedicalfileClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (m *Medicalfile) Unwrap() *Medicalfile {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Medicalfile is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Medicalfile) String() string {
	var builder strings.Builder
	builder.WriteString("Medicalfile(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", detail=")
	builder.WriteString(m.Detail)
	builder.WriteString(", added_time=")
	builder.WriteString(m.AddedTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Medicalfiles is a parsable slice of Medicalfile.
type Medicalfiles []*Medicalfile

func (m Medicalfiles) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}