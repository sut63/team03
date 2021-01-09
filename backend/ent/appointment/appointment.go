// Code generated by entc, DO NOT EDIT.

package appointment

const (
	// Label holds the string label denoting the appointment type in the database.
	Label = "appointment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDetail holds the string denoting the detail field in the database.
	FieldDetail = "detail"
	// FieldDatetime holds the string denoting the datetime field in the database.
	FieldDatetime = "datetime"

	// EdgePatient holds the string denoting the patient edge name in mutations.
	EdgePatient = "patient"
	// EdgeRoom holds the string denoting the room edge name in mutations.
	EdgeRoom = "room"
	// EdgeDentist holds the string denoting the dentist edge name in mutations.
	EdgeDentist = "dentist"

	// Table holds the table name of the appointment in the database.
	Table = "appointments"
	// PatientTable is the table the holds the patient relation/edge.
	PatientTable = "appointments"
	// PatientInverseTable is the table name for the Patient entity.
	// It exists in this package in order to avoid circular dependency with the "patient" package.
	PatientInverseTable = "patients"
	// PatientColumn is the table column denoting the patient relation/edge.
	PatientColumn = "patient_id"
	// RoomTable is the table the holds the room relation/edge.
	RoomTable = "appointments"
	// RoomInverseTable is the table name for the Room entity.
	// It exists in this package in order to avoid circular dependency with the "room" package.
	RoomInverseTable = "rooms"
	// RoomColumn is the table column denoting the room relation/edge.
	RoomColumn = "room_id"
	// DentistTable is the table the holds the dentist relation/edge.
	DentistTable = "appointments"
	// DentistInverseTable is the table name for the Dentist entity.
	// It exists in this package in order to avoid circular dependency with the "dentist" package.
	DentistInverseTable = "dentists"
	// DentistColumn is the table column denoting the dentist relation/edge.
	DentistColumn = "dentist_id"
)

// Columns holds all SQL columns for appointment fields.
var Columns = []string{
	FieldID,
	FieldDetail,
	FieldDatetime,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Appointment type.
var ForeignKeys = []string{
	"dentist_id",
	"patient_id",
	"room_id",
}

var (
	// DetailValidator is a validator for the "detail" field. It is called by the builders before save.
	DetailValidator func(string) error
)