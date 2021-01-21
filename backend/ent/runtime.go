// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/team03/app/ent/appointment"
	"github.com/team03/app/ent/degree"
	"github.com/team03/app/ent/dentalexpense"
	"github.com/team03/app/ent/dentist"
	"github.com/team03/app/ent/disease"
	"github.com/team03/app/ent/expert"
	"github.com/team03/app/ent/gender"
	"github.com/team03/app/ent/medicalcare"
	"github.com/team03/app/ent/medicalfile"
	"github.com/team03/app/ent/nurse"
	"github.com/team03/app/ent/patient"
	"github.com/team03/app/ent/room"
	"github.com/team03/app/ent/schema"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	appointmentFields := schema.Appointment{}.Fields()
	_ = appointmentFields
	// appointmentDescAppointID is the schema descriptor for AppointID field.
	appointmentDescAppointID := appointmentFields[0].Descriptor()
	// appointment.AppointIDValidator is a validator for the "AppointID" field. It is called by the builders before save.
	appointment.AppointIDValidator = appointmentDescAppointID.Validators[0].(func(string) error)
	// appointmentDescDetail is the schema descriptor for Detail field.
	appointmentDescDetail := appointmentFields[1].Descriptor()
	// appointment.DetailValidator is a validator for the "Detail" field. It is called by the builders before save.
	appointment.DetailValidator = appointmentDescDetail.Validators[0].(func(string) error)
	// appointmentDescRemark is the schema descriptor for Remark field.
	appointmentDescRemark := appointmentFields[3].Descriptor()
	// appointment.RemarkValidator is a validator for the "Remark" field. It is called by the builders before save.
	appointment.RemarkValidator = appointmentDescRemark.Validators[0].(func(string) error)
	degreeFields := schema.Degree{}.Fields()
	_ = degreeFields
	// degreeDescName is the schema descriptor for name field.
	degreeDescName := degreeFields[0].Descriptor()
	// degree.NameValidator is a validator for the "name" field. It is called by the builders before save.
	degree.NameValidator = degreeDescName.Validators[0].(func(string) error)
	dentalexpenseFields := schema.DentalExpense{}.Fields()
	_ = dentalexpenseFields
	// dentalexpenseDescTax is the schema descriptor for tax field.
	dentalexpenseDescTax := dentalexpenseFields[0].Descriptor()
	// dentalexpense.TaxValidator is a validator for the "tax" field. It is called by the builders before save.
	dentalexpense.TaxValidator = dentalexpenseDescTax.Validators[0].(func(string) error)
	// dentalexpenseDescName is the schema descriptor for name field.
	dentalexpenseDescName := dentalexpenseFields[1].Descriptor()
	// dentalexpense.NameValidator is a validator for the "name" field. It is called by the builders before save.
	dentalexpense.NameValidator = dentalexpenseDescName.Validators[0].(func(string) error)
	// dentalexpenseDescRates is the schema descriptor for rates field.
	dentalexpenseDescRates := dentalexpenseFields[2].Descriptor()
	// dentalexpense.RatesValidator is a validator for the "rates" field. It is called by the builders before save.
	dentalexpense.RatesValidator = dentalexpenseDescRates.Validators[0].(func(int) error)
	// dentalexpenseDescPhone is the schema descriptor for phone field.
	dentalexpenseDescPhone := dentalexpenseFields[3].Descriptor()
	// dentalexpense.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	dentalexpense.PhoneValidator = dentalexpenseDescPhone.Validators[0].(func(string) error)
	dentistFields := schema.Dentist{}.Fields()
	_ = dentistFields
	// dentistDescName is the schema descriptor for name field.
	dentistDescName := dentistFields[0].Descriptor()
	// dentist.NameValidator is a validator for the "name" field. It is called by the builders before save.
	dentist.NameValidator = dentistDescName.Validators[0].(func(string) error)
	// dentistDescAge is the schema descriptor for age field.
	dentistDescAge := dentistFields[1].Descriptor()
	// dentist.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	dentist.AgeValidator = dentistDescAge.Validators[0].(func(int) error)
	// dentistDescCardid is the schema descriptor for cardid field.
	dentistDescCardid := dentistFields[2].Descriptor()
	// dentist.CardidValidator is a validator for the "cardid" field. It is called by the builders before save.
	dentist.CardidValidator = dentistDescCardid.Validators[0].(func(string) error)
	// dentistDescExperience is the schema descriptor for experience field.
	dentistDescExperience := dentistFields[4].Descriptor()
	// dentist.ExperienceValidator is a validator for the "experience" field. It is called by the builders before save.
	dentist.ExperienceValidator = dentistDescExperience.Validators[0].(func(string) error)
	// dentistDescTel is the schema descriptor for tel field.
	dentistDescTel := dentistFields[5].Descriptor()
	// dentist.TelValidator is a validator for the "tel" field. It is called by the builders before save.
	dentist.TelValidator = dentistDescTel.Validators[0].(func(string) error)
	// dentistDescEmail is the schema descriptor for email field.
	dentistDescEmail := dentistFields[6].Descriptor()
	// dentist.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	dentist.EmailValidator = dentistDescEmail.Validators[0].(func(string) error)
	// dentistDescPassword is the schema descriptor for password field.
	dentistDescPassword := dentistFields[7].Descriptor()
	// dentist.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	dentist.PasswordValidator = dentistDescPassword.Validators[0].(func(string) error)
	diseaseFields := schema.Disease{}.Fields()
	_ = diseaseFields
	// diseaseDescName is the schema descriptor for name field.
	diseaseDescName := diseaseFields[0].Descriptor()
	// disease.NameValidator is a validator for the "name" field. It is called by the builders before save.
	disease.NameValidator = diseaseDescName.Validators[0].(func(string) error)
	expertFields := schema.Expert{}.Fields()
	_ = expertFields
	// expertDescName is the schema descriptor for name field.
	expertDescName := expertFields[0].Descriptor()
	// expert.NameValidator is a validator for the "name" field. It is called by the builders before save.
	expert.NameValidator = expertDescName.Validators[0].(func(string) error)
	genderFields := schema.Gender{}.Fields()
	_ = genderFields
	// genderDescName is the schema descriptor for name field.
	genderDescName := genderFields[0].Descriptor()
	// gender.NameValidator is a validator for the "name" field. It is called by the builders before save.
	gender.NameValidator = genderDescName.Validators[0].(func(string) error)
	medicalcareFields := schema.MedicalCare{}.Fields()
	_ = medicalcareFields
	// medicalcareDescName is the schema descriptor for name field.
	medicalcareDescName := medicalcareFields[0].Descriptor()
	// medicalcare.NameValidator is a validator for the "name" field. It is called by the builders before save.
	medicalcare.NameValidator = medicalcareDescName.Validators[0].(func(string) error)
	medicalfileFields := schema.Medicalfile{}.Fields()
	_ = medicalfileFields
	// medicalfileDescDrugAllergy is the schema descriptor for DrugAllergy field.
	medicalfileDescDrugAllergy := medicalfileFields[0].Descriptor()
	// medicalfile.DrugAllergyValidator is a validator for the "DrugAllergy" field. It is called by the builders before save.
	medicalfile.DrugAllergyValidator = func() func(string) error {
		validators := medicalfileDescDrugAllergy.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_DrugAllergy string) error {
			for _, fn := range fns {
				if err := fn(_DrugAllergy); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// medicalfileDescDetial is the schema descriptor for Detial field.
	medicalfileDescDetial := medicalfileFields[1].Descriptor()
	// medicalfile.DetialValidator is a validator for the "Detial" field. It is called by the builders before save.
	medicalfile.DetialValidator = func() func(string) error {
		validators := medicalfileDescDetial.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_Detial string) error {
			for _, fn := range fns {
				if err := fn(_Detial); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// medicalfileDescMedno is the schema descriptor for Medno field.
	medicalfileDescMedno := medicalfileFields[3].Descriptor()
	// medicalfile.MednoValidator is a validator for the "Medno" field. It is called by the builders before save.
	medicalfile.MednoValidator = medicalfileDescMedno.Validators[0].(func(string) error)
	nurseFields := schema.Nurse{}.Fields()
	_ = nurseFields
	// nurseDescNurseName is the schema descriptor for nurse_name field.
	nurseDescNurseName := nurseFields[0].Descriptor()
	// nurse.NurseNameValidator is a validator for the "nurse_name" field. It is called by the builders before save.
	nurse.NurseNameValidator = nurseDescNurseName.Validators[0].(func(string) error)
	patientFields := schema.Patient{}.Fields()
	_ = patientFields
	// patientDescPatientID is the schema descriptor for patient_ID field.
	patientDescPatientID := patientFields[0].Descriptor()
	// patient.PatientIDValidator is a validator for the "patient_ID" field. It is called by the builders before save.
	patient.PatientIDValidator = patientDescPatientID.Validators[0].(func(string) error)
	// patientDescName is the schema descriptor for name field.
	patientDescName := patientFields[1].Descriptor()
	// patient.NameValidator is a validator for the "name" field. It is called by the builders before save.
	patient.NameValidator = patientDescName.Validators[0].(func(string) error)
	// patientDescCardID is the schema descriptor for cardID field.
	patientDescCardID := patientFields[2].Descriptor()
	// patient.CardIDValidator is a validator for the "cardID" field. It is called by the builders before save.
	patient.CardIDValidator = patientDescCardID.Validators[0].(func(string) error)
	// patientDescTel is the schema descriptor for tel field.
	patientDescTel := patientFields[3].Descriptor()
	// patient.TelValidator is a validator for the "tel" field. It is called by the builders before save.
	patient.TelValidator = patientDescTel.Validators[0].(func(string) error)
	// patientDescAge is the schema descriptor for age field.
	patientDescAge := patientFields[4].Descriptor()
	// patient.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	patient.AgeValidator = patientDescAge.Validators[0].(func(int) error)
	roomFields := schema.Room{}.Fields()
	_ = roomFields
	// roomDescName is the schema descriptor for name field.
	roomDescName := roomFields[0].Descriptor()
	// room.NameValidator is a validator for the "name" field. It is called by the builders before save.
	room.NameValidator = roomDescName.Validators[0].(func(string) error)
}
