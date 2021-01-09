// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/tanapon395/playlist-video/ent/disease"
	"github.com/tanapon395/playlist-video/ent/gender"
	"github.com/tanapon395/playlist-video/ent/medicalcare"
	"github.com/tanapon395/playlist-video/ent/nurse"
	"github.com/tanapon395/playlist-video/ent/patient"
)

// Patient is the model entity for the Patient schema.
type Patient struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// PatientID holds the value of the "patient_ID" field.
	PatientID string `json:"patient_ID,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CardID holds the value of the "cardID" field.
	CardID string `json:"cardID,omitempty"`
	// Tel holds the value of the "tel" field.
	Tel string `json:"tel,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Birthday holds the value of the "birthday" field.
	Birthday time.Time `json:"birthday,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PatientQuery when eager-loading is set.
	Edges          PatientEdges `json:"edges"`
	disease_id     *int
	gender_id      *int
	medicalcare_id *int
	nurse_id       *int
}

// PatientEdges holds the relations/edges for other nodes in the graph.
type PatientEdges struct {
	// Gender holds the value of the gender edge.
	Gender *Gender
	// Medicalcare holds the value of the medicalcare edge.
	Medicalcare *MedicalCare
	// Nurse holds the value of the nurse edge.
	Nurse *Nurse
	// Disease holds the value of the disease edge.
	Disease *Disease
	// Medicalfiles holds the value of the medicalfiles edge.
	Medicalfiles []*Medicalfile
	// Queue holds the value of the queue edge.
	Queue []*Queue
	// Appointment holds the value of the appointment edge.
	Appointment []*Appointment
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [7]bool
}

// GenderOrErr returns the Gender value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PatientEdges) GenderOrErr() (*Gender, error) {
	if e.loadedTypes[0] {
		if e.Gender == nil {
			// The edge gender was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: gender.Label}
		}
		return e.Gender, nil
	}
	return nil, &NotLoadedError{edge: "gender"}
}

// MedicalcareOrErr returns the Medicalcare value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PatientEdges) MedicalcareOrErr() (*MedicalCare, error) {
	if e.loadedTypes[1] {
		if e.Medicalcare == nil {
			// The edge medicalcare was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: medicalcare.Label}
		}
		return e.Medicalcare, nil
	}
	return nil, &NotLoadedError{edge: "medicalcare"}
}

// NurseOrErr returns the Nurse value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PatientEdges) NurseOrErr() (*Nurse, error) {
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

// DiseaseOrErr returns the Disease value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PatientEdges) DiseaseOrErr() (*Disease, error) {
	if e.loadedTypes[3] {
		if e.Disease == nil {
			// The edge disease was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: disease.Label}
		}
		return e.Disease, nil
	}
	return nil, &NotLoadedError{edge: "disease"}
}

// MedicalfilesOrErr returns the Medicalfiles value or an error if the edge
// was not loaded in eager-loading.
func (e PatientEdges) MedicalfilesOrErr() ([]*Medicalfile, error) {
	if e.loadedTypes[4] {
		return e.Medicalfiles, nil
	}
	return nil, &NotLoadedError{edge: "medicalfiles"}
}

// QueueOrErr returns the Queue value or an error if the edge
// was not loaded in eager-loading.
func (e PatientEdges) QueueOrErr() ([]*Queue, error) {
	if e.loadedTypes[5] {
		return e.Queue, nil
	}
	return nil, &NotLoadedError{edge: "queue"}
}

// AppointmentOrErr returns the Appointment value or an error if the edge
// was not loaded in eager-loading.
func (e PatientEdges) AppointmentOrErr() ([]*Appointment, error) {
	if e.loadedTypes[6] {
		return e.Appointment, nil
	}
	return nil, &NotLoadedError{edge: "appointment"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Patient) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // patient_ID
		&sql.NullString{}, // name
		&sql.NullString{}, // cardID
		&sql.NullString{}, // tel
		&sql.NullInt64{},  // age
		&sql.NullTime{},   // birthday
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Patient) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // disease_id
		&sql.NullInt64{}, // gender_id
		&sql.NullInt64{}, // medicalcare_id
		&sql.NullInt64{}, // nurse_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Patient fields.
func (pa *Patient) assignValues(values ...interface{}) error {
	if m, n := len(values), len(patient.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pa.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field patient_ID", values[0])
	} else if value.Valid {
		pa.PatientID = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[1])
	} else if value.Valid {
		pa.Name = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field cardID", values[2])
	} else if value.Valid {
		pa.CardID = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field tel", values[3])
	} else if value.Valid {
		pa.Tel = value.String
	}
	if value, ok := values[4].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field age", values[4])
	} else if value.Valid {
		pa.Age = int(value.Int64)
	}
	if value, ok := values[5].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field birthday", values[5])
	} else if value.Valid {
		pa.Birthday = value.Time
	}
	values = values[6:]
	if len(values) == len(patient.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field disease_id", value)
		} else if value.Valid {
			pa.disease_id = new(int)
			*pa.disease_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field gender_id", value)
		} else if value.Valid {
			pa.gender_id = new(int)
			*pa.gender_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field medicalcare_id", value)
		} else if value.Valid {
			pa.medicalcare_id = new(int)
			*pa.medicalcare_id = int(value.Int64)
		}
		if value, ok := values[3].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field nurse_id", value)
		} else if value.Valid {
			pa.nurse_id = new(int)
			*pa.nurse_id = int(value.Int64)
		}
	}
	return nil
}

// QueryGender queries the gender edge of the Patient.
func (pa *Patient) QueryGender() *GenderQuery {
	return (&PatientClient{config: pa.config}).QueryGender(pa)
}

// QueryMedicalcare queries the medicalcare edge of the Patient.
func (pa *Patient) QueryMedicalcare() *MedicalCareQuery {
	return (&PatientClient{config: pa.config}).QueryMedicalcare(pa)
}

// QueryNurse queries the nurse edge of the Patient.
func (pa *Patient) QueryNurse() *NurseQuery {
	return (&PatientClient{config: pa.config}).QueryNurse(pa)
}

// QueryDisease queries the disease edge of the Patient.
func (pa *Patient) QueryDisease() *DiseaseQuery {
	return (&PatientClient{config: pa.config}).QueryDisease(pa)
}

// QueryMedicalfiles queries the medicalfiles edge of the Patient.
func (pa *Patient) QueryMedicalfiles() *MedicalfileQuery {
	return (&PatientClient{config: pa.config}).QueryMedicalfiles(pa)
}

// QueryQueue queries the queue edge of the Patient.
func (pa *Patient) QueryQueue() *QueueQuery {
	return (&PatientClient{config: pa.config}).QueryQueue(pa)
}

// QueryAppointment queries the appointment edge of the Patient.
func (pa *Patient) QueryAppointment() *AppointmentQuery {
	return (&PatientClient{config: pa.config}).QueryAppointment(pa)
}

// Update returns a builder for updating this Patient.
// Note that, you need to call Patient.Unwrap() before calling this method, if this Patient
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Patient) Update() *PatientUpdateOne {
	return (&PatientClient{config: pa.config}).UpdateOne(pa)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pa *Patient) Unwrap() *Patient {
	tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Patient is not a transactional entity")
	}
	pa.config.driver = tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Patient) String() string {
	var builder strings.Builder
	builder.WriteString("Patient(")
	builder.WriteString(fmt.Sprintf("id=%v", pa.ID))
	builder.WriteString(", patient_ID=")
	builder.WriteString(pa.PatientID)
	builder.WriteString(", name=")
	builder.WriteString(pa.Name)
	builder.WriteString(", cardID=")
	builder.WriteString(pa.CardID)
	builder.WriteString(", tel=")
	builder.WriteString(pa.Tel)
	builder.WriteString(", age=")
	builder.WriteString(fmt.Sprintf("%v", pa.Age))
	builder.WriteString(", birthday=")
	builder.WriteString(pa.Birthday.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Patients is a parsable slice of Patient.
type Patients []*Patient

func (pa Patients) config(cfg config) {
	for _i := range pa {
		pa[_i].config = cfg
	}
}
