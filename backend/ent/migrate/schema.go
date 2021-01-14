// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// AppointmentsColumns holds the columns for the "appointments" table.
	AppointmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "detail", Type: field.TypeString},
		{Name: "datetime", Type: field.TypeTime},
		{Name: "dentist_id", Type: field.TypeInt, Nullable: true},
		{Name: "patient_id", Type: field.TypeInt, Nullable: true},
		{Name: "room_id", Type: field.TypeInt, Nullable: true},
	}
	// AppointmentsTable holds the schema information for the "appointments" table.
	AppointmentsTable = &schema.Table{
		Name:       "appointments",
		Columns:    AppointmentsColumns,
		PrimaryKey: []*schema.Column{AppointmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "appointments_dentists_appointment",
				Columns: []*schema.Column{AppointmentsColumns[3]},

				RefColumns: []*schema.Column{DentistsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "appointments_patients_appointment",
				Columns: []*schema.Column{AppointmentsColumns[4]},

				RefColumns: []*schema.Column{PatientsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "appointments_rooms_appointment",
				Columns: []*schema.Column{AppointmentsColumns[5]},

				RefColumns: []*schema.Column{RoomsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DegreesColumns holds the columns for the "degrees" table.
	DegreesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// DegreesTable holds the schema information for the "degrees" table.
	DegreesTable = &schema.Table{
		Name:        "degrees",
		Columns:     DegreesColumns,
		PrimaryKey:  []*schema.Column{DegreesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// DentalExpensesColumns holds the columns for the "dental_expenses" table.
	DentalExpensesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "added_time", Type: field.TypeTime},
		{Name: "medicalfile_id", Type: field.TypeInt, Nullable: true},
		{Name: "nurse_id", Type: field.TypeInt, Nullable: true},
		{Name: "pricetype_id", Type: field.TypeInt, Nullable: true},
	}
	// DentalExpensesTable holds the schema information for the "dental_expenses" table.
	DentalExpensesTable = &schema.Table{
		Name:       "dental_expenses",
		Columns:    DentalExpensesColumns,
		PrimaryKey: []*schema.Column{DentalExpensesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "dental_expenses_medicalfiles_dentalexpenses",
				Columns: []*schema.Column{DentalExpensesColumns[2]},

				RefColumns: []*schema.Column{MedicalfilesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "dental_expenses_nurses_dentalexpenses",
				Columns: []*schema.Column{DentalExpensesColumns[3]},

				RefColumns: []*schema.Column{NursesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "dental_expenses_price_types_dentalexpenses",
				Columns: []*schema.Column{DentalExpensesColumns[4]},

				RefColumns: []*schema.Column{PriceTypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DentistsColumns holds the columns for the "dentists" table.
	DentistsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "age", Type: field.TypeInt},
		{Name: "cardid", Type: field.TypeString},
		{Name: "birthday", Type: field.TypeTime},
		{Name: "experience", Type: field.TypeString},
		{Name: "tel", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "degree_id", Type: field.TypeInt, Nullable: true},
		{Name: "expert_id", Type: field.TypeInt, Nullable: true},
		{Name: "gender_id", Type: field.TypeInt, Nullable: true},
		{Name: "nurse_id", Type: field.TypeInt, Nullable: true},
	}
	// DentistsTable holds the schema information for the "dentists" table.
	DentistsTable = &schema.Table{
		Name:       "dentists",
		Columns:    DentistsColumns,
		PrimaryKey: []*schema.Column{DentistsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "dentists_degrees_dentists",
				Columns: []*schema.Column{DentistsColumns[9]},

				RefColumns: []*schema.Column{DegreesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "dentists_experts_dentists",
				Columns: []*schema.Column{DentistsColumns[10]},

				RefColumns: []*schema.Column{ExpertsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "dentists_genders_dentists",
				Columns: []*schema.Column{DentistsColumns[11]},

				RefColumns: []*schema.Column{GendersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "dentists_nurses_dentists",
				Columns: []*schema.Column{DentistsColumns[12]},

				RefColumns: []*schema.Column{NursesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DiseasesColumns holds the columns for the "diseases" table.
	DiseasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// DiseasesTable holds the schema information for the "diseases" table.
	DiseasesTable = &schema.Table{
		Name:        "diseases",
		Columns:     DiseasesColumns,
		PrimaryKey:  []*schema.Column{DiseasesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ExpertsColumns holds the columns for the "experts" table.
	ExpertsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// ExpertsTable holds the schema information for the "experts" table.
	ExpertsTable = &schema.Table{
		Name:        "experts",
		Columns:     ExpertsColumns,
		PrimaryKey:  []*schema.Column{ExpertsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// GendersColumns holds the columns for the "genders" table.
	GendersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// GendersTable holds the schema information for the "genders" table.
	GendersTable = &schema.Table{
		Name:        "genders",
		Columns:     GendersColumns,
		PrimaryKey:  []*schema.Column{GendersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// MedicalCaresColumns holds the columns for the "medical_cares" table.
	MedicalCaresColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// MedicalCaresTable holds the schema information for the "medical_cares" table.
	MedicalCaresTable = &schema.Table{
		Name:        "medical_cares",
		Columns:     MedicalCaresColumns,
		PrimaryKey:  []*schema.Column{MedicalCaresColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// MedicalfilesColumns holds the columns for the "medicalfiles" table.
	MedicalfilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "detail", Type: field.TypeString},
		{Name: "added_time", Type: field.TypeTime},
		{Name: "dentist_id", Type: field.TypeInt, Nullable: true},
		{Name: "nurse_id", Type: field.TypeInt, Nullable: true},
		{Name: "patient_id", Type: field.TypeInt, Nullable: true},
	}
	// MedicalfilesTable holds the schema information for the "medicalfiles" table.
	MedicalfilesTable = &schema.Table{
		Name:       "medicalfiles",
		Columns:    MedicalfilesColumns,
		PrimaryKey: []*schema.Column{MedicalfilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "medicalfiles_dentists_medicalfiles",
				Columns: []*schema.Column{MedicalfilesColumns[3]},

				RefColumns: []*schema.Column{DentistsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "medicalfiles_nurses_medicalfiles",
				Columns: []*schema.Column{MedicalfilesColumns[4]},

				RefColumns: []*schema.Column{NursesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "medicalfiles_patients_medicalfiles",
				Columns: []*schema.Column{MedicalfilesColumns[5]},

				RefColumns: []*schema.Column{PatientsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// NursesColumns holds the columns for the "nurses" table.
	NursesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "nurse_name", Type: field.TypeString},
		{Name: "nurse_age", Type: field.TypeInt},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
	}
	// NursesTable holds the schema information for the "nurses" table.
	NursesTable = &schema.Table{
		Name:        "nurses",
		Columns:     NursesColumns,
		PrimaryKey:  []*schema.Column{NursesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PatientsColumns holds the columns for the "patients" table.
	PatientsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "patient_id", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "card_id", Type: field.TypeString},
		{Name: "tel", Type: field.TypeString},
		{Name: "age", Type: field.TypeInt},
		{Name: "birthday", Type: field.TypeTime},
		{Name: "disease_id", Type: field.TypeInt, Nullable: true},
		{Name: "gender_id", Type: field.TypeInt, Nullable: true},
		{Name: "medicalcare_id", Type: field.TypeInt, Nullable: true},
		{Name: "nurse_id", Type: field.TypeInt, Nullable: true},
	}
	// PatientsTable holds the schema information for the "patients" table.
	PatientsTable = &schema.Table{
		Name:       "patients",
		Columns:    PatientsColumns,
		PrimaryKey: []*schema.Column{PatientsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "patients_diseases_patients",
				Columns: []*schema.Column{PatientsColumns[7]},

				RefColumns: []*schema.Column{DiseasesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "patients_genders_patients",
				Columns: []*schema.Column{PatientsColumns[8]},

				RefColumns: []*schema.Column{GendersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "patients_medical_cares_patients",
				Columns: []*schema.Column{PatientsColumns[9]},

				RefColumns: []*schema.Column{MedicalCaresColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "patients_nurses_patients",
				Columns: []*schema.Column{PatientsColumns[10]},

				RefColumns: []*schema.Column{NursesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PriceTypesColumns holds the columns for the "price_types" table.
	PriceTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// PriceTypesTable holds the schema information for the "price_types" table.
	PriceTypesTable = &schema.Table{
		Name:        "price_types",
		Columns:     PriceTypesColumns,
		PrimaryKey:  []*schema.Column{PriceTypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// QueuesColumns holds the columns for the "queues" table.
	QueuesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "dental", Type: field.TypeString},
		{Name: "queue_time", Type: field.TypeTime},
		{Name: "dentist_id", Type: field.TypeInt, Nullable: true},
		{Name: "nurse_id", Type: field.TypeInt, Nullable: true},
		{Name: "patient_id", Type: field.TypeInt, Nullable: true},
	}
	// QueuesTable holds the schema information for the "queues" table.
	QueuesTable = &schema.Table{
		Name:       "queues",
		Columns:    QueuesColumns,
		PrimaryKey: []*schema.Column{QueuesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "queues_dentists_queue",
				Columns: []*schema.Column{QueuesColumns[3]},

				RefColumns: []*schema.Column{DentistsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "queues_nurses_queue",
				Columns: []*schema.Column{QueuesColumns[4]},

				RefColumns: []*schema.Column{NursesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "queues_patients_queue",
				Columns: []*schema.Column{QueuesColumns[5]},

				RefColumns: []*schema.Column{PatientsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RoomsColumns holds the columns for the "rooms" table.
	RoomsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// RoomsTable holds the schema information for the "rooms" table.
	RoomsTable = &schema.Table{
		Name:        "rooms",
		Columns:     RoomsColumns,
		PrimaryKey:  []*schema.Column{RoomsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AppointmentsTable,
		DegreesTable,
		DentalExpensesTable,
		DentistsTable,
		DiseasesTable,
		ExpertsTable,
		GendersTable,
		MedicalCaresTable,
		MedicalfilesTable,
		NursesTable,
		PatientsTable,
		PriceTypesTable,
		QueuesTable,
		RoomsTable,
	}
)

func init() {
	AppointmentsTable.ForeignKeys[0].RefTable = DentistsTable
	AppointmentsTable.ForeignKeys[1].RefTable = PatientsTable
	AppointmentsTable.ForeignKeys[2].RefTable = RoomsTable
	DentalExpensesTable.ForeignKeys[0].RefTable = MedicalfilesTable
	DentalExpensesTable.ForeignKeys[1].RefTable = NursesTable
	DentalExpensesTable.ForeignKeys[2].RefTable = PriceTypesTable
	DentistsTable.ForeignKeys[0].RefTable = DegreesTable
	DentistsTable.ForeignKeys[1].RefTable = ExpertsTable
	DentistsTable.ForeignKeys[2].RefTable = GendersTable
	DentistsTable.ForeignKeys[3].RefTable = NursesTable
	MedicalfilesTable.ForeignKeys[0].RefTable = DentistsTable
	MedicalfilesTable.ForeignKeys[1].RefTable = NursesTable
	MedicalfilesTable.ForeignKeys[2].RefTable = PatientsTable
	PatientsTable.ForeignKeys[0].RefTable = DiseasesTable
	PatientsTable.ForeignKeys[1].RefTable = GendersTable
	PatientsTable.ForeignKeys[2].RefTable = MedicalCaresTable
	PatientsTable.ForeignKeys[3].RefTable = NursesTable
	QueuesTable.ForeignKeys[0].RefTable = DentistsTable
	QueuesTable.ForeignKeys[1].RefTable = NursesTable
	QueuesTable.ForeignKeys[2].RefTable = PatientsTable
}
