// Code generated by entc, DO NOT EDIT.

package pricetype

const (
	// Label holds the string label denoting the pricetype type in the database.
	Label = "price_type"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"

	// EdgeDentalexpenses holds the string denoting the dentalexpenses edge name in mutations.
	EdgeDentalexpenses = "dentalexpenses"

	// Table holds the table name of the pricetype in the database.
	Table = "price_types"
	// DentalexpensesTable is the table the holds the dentalexpenses relation/edge.
	DentalexpensesTable = "dental_expenses"
	// DentalexpensesInverseTable is the table name for the DentalExpense entity.
	// It exists in this package in order to avoid circular dependency with the "dentalexpense" package.
	DentalexpensesInverseTable = "dental_expenses"
	// DentalexpensesColumn is the table column denoting the dentalexpenses relation/edge.
	DentalexpensesColumn = "pricetype_id"
)

// Columns holds all SQL columns for pricetype fields.
var Columns = []string{
	FieldID,
	FieldName,
}
