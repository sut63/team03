// Code generated by entc, DO NOT EDIT.

package degree

const (
	// Label holds the string label denoting the degree type in the database.
	Label = "degree"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDegreeName holds the string denoting the degree_name field in the database.
	FieldDegreeName = "degree_name"

	// EdgeDentists holds the string denoting the dentists edge name in mutations.
	EdgeDentists = "dentists"

	// Table holds the table name of the degree in the database.
	Table = "degrees"
	// DentistsTable is the table the holds the dentists relation/edge.
	DentistsTable = "dentists"
	// DentistsInverseTable is the table name for the Dentist entity.
	// It exists in this package in order to avoid circular dependency with the "dentist" package.
	DentistsInverseTable = "dentists"
	// DentistsColumn is the table column denoting the dentists relation/edge.
	DentistsColumn = "degree_id"
)

// Columns holds all SQL columns for degree fields.
var Columns = []string{
	FieldID,
	FieldDegreeName,
}

var (
	// DegreeNameValidator is a validator for the "degree_name" field. It is called by the builders before save.
	DegreeNameValidator func(string) error
)