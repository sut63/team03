// Code generated by entc, DO NOT EDIT.

package nurse

const (
	// Label holds the string label denoting the nurse type in the database.
	Label = "nurse"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNurseName holds the string denoting the nurse_name field in the database.
	FieldNurseName = "nurse_name"
	// FieldNurseAge holds the string denoting the nurse_age field in the database.
	FieldNurseAge = "nurse_age"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"

	// EdgeQueue holds the string denoting the queue edge name in mutations.
	EdgeQueue = "queue"
	// EdgeMedicalfiles holds the string denoting the medicalfiles edge name in mutations.
	EdgeMedicalfiles = "medicalfiles"
	// EdgeDentalexpenses holds the string denoting the dentalexpenses edge name in mutations.
	EdgeDentalexpenses = "dentalexpenses"
	// EdgePatients holds the string denoting the patients edge name in mutations.
	EdgePatients = "patients"
	// EdgeDentists holds the string denoting the dentists edge name in mutations.
	EdgeDentists = "dentists"

	// Table holds the table name of the nurse in the database.
	Table = "nurses"
	// QueueTable is the table the holds the queue relation/edge.
	QueueTable = "queues"
	// QueueInverseTable is the table name for the Queue entity.
	// It exists in this package in order to avoid circular dependency with the "queue" package.
	QueueInverseTable = "queues"
	// QueueColumn is the table column denoting the queue relation/edge.
	QueueColumn = "nurse_id"
	// MedicalfilesTable is the table the holds the medicalfiles relation/edge.
	MedicalfilesTable = "medicalfiles"
	// MedicalfilesInverseTable is the table name for the Medicalfile entity.
	// It exists in this package in order to avoid circular dependency with the "medicalfile" package.
	MedicalfilesInverseTable = "medicalfiles"
	// MedicalfilesColumn is the table column denoting the medicalfiles relation/edge.
	MedicalfilesColumn = "nurse_id"
	// DentalexpensesTable is the table the holds the dentalexpenses relation/edge.
	DentalexpensesTable = "dental_expenses"
	// DentalexpensesInverseTable is the table name for the DentalExpense entity.
	// It exists in this package in order to avoid circular dependency with the "dentalexpense" package.
	DentalexpensesInverseTable = "dental_expenses"
	// DentalexpensesColumn is the table column denoting the dentalexpenses relation/edge.
	DentalexpensesColumn = "nurse_id"
	// PatientsTable is the table the holds the patients relation/edge.
	PatientsTable = "patients"
	// PatientsInverseTable is the table name for the Patient entity.
	// It exists in this package in order to avoid circular dependency with the "patient" package.
	PatientsInverseTable = "patients"
	// PatientsColumn is the table column denoting the patients relation/edge.
	PatientsColumn = "nurse_id"
	// DentistsTable is the table the holds the dentists relation/edge.
	DentistsTable = "dentists"
	// DentistsInverseTable is the table name for the Dentist entity.
	// It exists in this package in order to avoid circular dependency with the "dentist" package.
	DentistsInverseTable = "dentists"
	// DentistsColumn is the table column denoting the dentists relation/edge.
	DentistsColumn = "nurse_id"
)

// Columns holds all SQL columns for nurse fields.
var Columns = []string{
	FieldID,
	FieldNurseName,
	FieldNurseAge,
	FieldEmail,
	FieldPassword,
}

var (
	// NurseNameValidator is a validator for the "nurse_name" field. It is called by the builders before save.
	NurseNameValidator func(string) error
)
