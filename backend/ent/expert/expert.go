// Code generated by entc, DO NOT EDIT.

package expert

const (
	// Label holds the string label denoting the expert type in the database.
	Label = "expert"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"

	// EdgeDentists holds the string denoting the dentists edge name in mutations.
	EdgeDentists = "dentists"

	// Table holds the table name of the expert in the database.
	Table = "experts"
	// DentistsTable is the table the holds the dentists relation/edge.
	DentistsTable = "dentists"
	// DentistsInverseTable is the table name for the Dentist entity.
	// It exists in this package in order to avoid circular dependency with the "dentist" package.
	DentistsInverseTable = "dentists"
	// DentistsColumn is the table column denoting the dentists relation/edge.
	DentistsColumn = "expert_id"
)

// Columns holds all SQL columns for expert fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)
