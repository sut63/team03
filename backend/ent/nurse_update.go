// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team03/app/ent/appointment"
	"github.com/team03/app/ent/dentalexpense"
	"github.com/team03/app/ent/dentist"
	"github.com/team03/app/ent/medicalfile"
	"github.com/team03/app/ent/nurse"
	"github.com/team03/app/ent/patient"
	"github.com/team03/app/ent/predicate"
	"github.com/team03/app/ent/queue"
)

// NurseUpdate is the builder for updating Nurse entities.
type NurseUpdate struct {
	config
	hooks      []Hook
	mutation   *NurseMutation
	predicates []predicate.Nurse
}

// Where adds a new predicate for the builder.
func (nu *NurseUpdate) Where(ps ...predicate.Nurse) *NurseUpdate {
	nu.predicates = append(nu.predicates, ps...)
	return nu
}

// SetNurseName sets the nurse_name field.
func (nu *NurseUpdate) SetNurseName(s string) *NurseUpdate {
	nu.mutation.SetNurseName(s)
	return nu
}

// SetNurseAge sets the nurse_age field.
func (nu *NurseUpdate) SetNurseAge(i int) *NurseUpdate {
	nu.mutation.ResetNurseAge()
	nu.mutation.SetNurseAge(i)
	return nu
}

// AddNurseAge adds i to nurse_age.
func (nu *NurseUpdate) AddNurseAge(i int) *NurseUpdate {
	nu.mutation.AddNurseAge(i)
	return nu
}

// SetEmail sets the email field.
func (nu *NurseUpdate) SetEmail(s string) *NurseUpdate {
	nu.mutation.SetEmail(s)
	return nu
}

// SetPassword sets the password field.
func (nu *NurseUpdate) SetPassword(s string) *NurseUpdate {
	nu.mutation.SetPassword(s)
	return nu
}

// AddQueueIDs adds the queue edge to Queue by ids.
func (nu *NurseUpdate) AddQueueIDs(ids ...int) *NurseUpdate {
	nu.mutation.AddQueueIDs(ids...)
	return nu
}

// AddQueue adds the queue edges to Queue.
func (nu *NurseUpdate) AddQueue(q ...*Queue) *NurseUpdate {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return nu.AddQueueIDs(ids...)
}

// AddMedicalfileIDs adds the medicalfiles edge to Medicalfile by ids.
func (nu *NurseUpdate) AddMedicalfileIDs(ids ...int) *NurseUpdate {
	nu.mutation.AddMedicalfileIDs(ids...)
	return nu
}

// AddMedicalfiles adds the medicalfiles edges to Medicalfile.
func (nu *NurseUpdate) AddMedicalfiles(m ...*Medicalfile) *NurseUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return nu.AddMedicalfileIDs(ids...)
}

// AddDentalexpenseIDs adds the dentalexpenses edge to Dentalexpense by ids.
func (nu *NurseUpdate) AddDentalexpenseIDs(ids ...int) *NurseUpdate {
	nu.mutation.AddDentalexpenseIDs(ids...)
	return nu
}

// AddDentalexpenses adds the dentalexpenses edges to Dentalexpense.
func (nu *NurseUpdate) AddDentalexpenses(d ...*Dentalexpense) *NurseUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return nu.AddDentalexpenseIDs(ids...)
}

// AddPatientIDs adds the patients edge to Patient by ids.
func (nu *NurseUpdate) AddPatientIDs(ids ...int) *NurseUpdate {
	nu.mutation.AddPatientIDs(ids...)
	return nu
}

// AddPatients adds the patients edges to Patient.
func (nu *NurseUpdate) AddPatients(p ...*Patient) *NurseUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return nu.AddPatientIDs(ids...)
}

// AddDentistIDs adds the dentists edge to Dentist by ids.
func (nu *NurseUpdate) AddDentistIDs(ids ...int) *NurseUpdate {
	nu.mutation.AddDentistIDs(ids...)
	return nu
}

// AddDentists adds the dentists edges to Dentist.
func (nu *NurseUpdate) AddDentists(d ...*Dentist) *NurseUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return nu.AddDentistIDs(ids...)
}

// AddAppointmentIDs adds the appointment edge to Appointment by ids.
func (nu *NurseUpdate) AddAppointmentIDs(ids ...int) *NurseUpdate {
	nu.mutation.AddAppointmentIDs(ids...)
	return nu
}

// AddAppointment adds the appointment edges to Appointment.
func (nu *NurseUpdate) AddAppointment(a ...*Appointment) *NurseUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return nu.AddAppointmentIDs(ids...)
}

// Mutation returns the NurseMutation object of the builder.
func (nu *NurseUpdate) Mutation() *NurseMutation {
	return nu.mutation
}

// RemoveQueueIDs removes the queue edge to Queue by ids.
func (nu *NurseUpdate) RemoveQueueIDs(ids ...int) *NurseUpdate {
	nu.mutation.RemoveQueueIDs(ids...)
	return nu
}

// RemoveQueue removes queue edges to Queue.
func (nu *NurseUpdate) RemoveQueue(q ...*Queue) *NurseUpdate {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return nu.RemoveQueueIDs(ids...)
}

// RemoveMedicalfileIDs removes the medicalfiles edge to Medicalfile by ids.
func (nu *NurseUpdate) RemoveMedicalfileIDs(ids ...int) *NurseUpdate {
	nu.mutation.RemoveMedicalfileIDs(ids...)
	return nu
}

// RemoveMedicalfiles removes medicalfiles edges to Medicalfile.
func (nu *NurseUpdate) RemoveMedicalfiles(m ...*Medicalfile) *NurseUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return nu.RemoveMedicalfileIDs(ids...)
}

// RemoveDentalexpenseIDs removes the dentalexpenses edge to Dentalexpense by ids.
func (nu *NurseUpdate) RemoveDentalexpenseIDs(ids ...int) *NurseUpdate {
	nu.mutation.RemoveDentalexpenseIDs(ids...)
	return nu
}

// RemoveDentalexpenses removes dentalexpenses edges to Dentalexpense.
func (nu *NurseUpdate) RemoveDentalexpenses(d ...*Dentalexpense) *NurseUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return nu.RemoveDentalexpenseIDs(ids...)
}

// RemovePatientIDs removes the patients edge to Patient by ids.
func (nu *NurseUpdate) RemovePatientIDs(ids ...int) *NurseUpdate {
	nu.mutation.RemovePatientIDs(ids...)
	return nu
}

// RemovePatients removes patients edges to Patient.
func (nu *NurseUpdate) RemovePatients(p ...*Patient) *NurseUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return nu.RemovePatientIDs(ids...)
}

// RemoveDentistIDs removes the dentists edge to Dentist by ids.
func (nu *NurseUpdate) RemoveDentistIDs(ids ...int) *NurseUpdate {
	nu.mutation.RemoveDentistIDs(ids...)
	return nu
}

// RemoveDentists removes dentists edges to Dentist.
func (nu *NurseUpdate) RemoveDentists(d ...*Dentist) *NurseUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return nu.RemoveDentistIDs(ids...)
}

// RemoveAppointmentIDs removes the appointment edge to Appointment by ids.
func (nu *NurseUpdate) RemoveAppointmentIDs(ids ...int) *NurseUpdate {
	nu.mutation.RemoveAppointmentIDs(ids...)
	return nu
}

// RemoveAppointment removes appointment edges to Appointment.
func (nu *NurseUpdate) RemoveAppointment(a ...*Appointment) *NurseUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return nu.RemoveAppointmentIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (nu *NurseUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := nu.mutation.NurseName(); ok {
		if err := nurse.NurseNameValidator(v); err != nil {
			return 0, &ValidationError{Name: "nurse_name", err: fmt.Errorf("ent: validator failed for field \"nurse_name\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(nu.hooks) == 0 {
		affected, err = nu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NurseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nu.mutation = mutation
			affected, err = nu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nu.hooks) - 1; i >= 0; i-- {
			mut = nu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NurseUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NurseUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NurseUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nu *NurseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   nurse.Table,
			Columns: nurse.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: nurse.FieldID,
			},
		},
	}
	if ps := nu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.NurseName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nurse.FieldNurseName,
		})
	}
	if value, ok := nu.mutation.NurseAge(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: nurse.FieldNurseAge,
		})
	}
	if value, ok := nu.mutation.AddedNurseAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: nurse.FieldNurseAge,
		})
	}
	if value, ok := nu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nurse.FieldEmail,
		})
	}
	if value, ok := nu.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nurse.FieldPassword,
		})
	}
	if nodes := nu.mutation.RemovedQueueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.QueueTable,
			Columns: []string{nurse.QueueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: queue.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.QueueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.QueueTable,
			Columns: []string{nurse.QueueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: queue.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := nu.mutation.RemovedMedicalfilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.MedicalfilesTable,
			Columns: []string{nurse.MedicalfilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicalfile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.MedicalfilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.MedicalfilesTable,
			Columns: []string{nurse.MedicalfilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicalfile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := nu.mutation.RemovedDentalexpensesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.DentalexpensesTable,
			Columns: []string{nurse.DentalexpensesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dentalexpense.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.DentalexpensesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.DentalexpensesTable,
			Columns: []string{nurse.DentalexpensesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dentalexpense.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := nu.mutation.RemovedPatientsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.PatientsTable,
			Columns: []string{nurse.PatientsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.PatientsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.PatientsTable,
			Columns: []string{nurse.PatientsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := nu.mutation.RemovedDentistsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.DentistsTable,
			Columns: []string{nurse.DentistsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dentist.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.DentistsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.DentistsTable,
			Columns: []string{nurse.DentistsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dentist.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := nu.mutation.RemovedAppointmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.AppointmentTable,
			Columns: []string{nurse.AppointmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: appointment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.AppointmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.AppointmentTable,
			Columns: []string{nurse.AppointmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: appointment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nurse.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// NurseUpdateOne is the builder for updating a single Nurse entity.
type NurseUpdateOne struct {
	config
	hooks    []Hook
	mutation *NurseMutation
}

// SetNurseName sets the nurse_name field.
func (nuo *NurseUpdateOne) SetNurseName(s string) *NurseUpdateOne {
	nuo.mutation.SetNurseName(s)
	return nuo
}

// SetNurseAge sets the nurse_age field.
func (nuo *NurseUpdateOne) SetNurseAge(i int) *NurseUpdateOne {
	nuo.mutation.ResetNurseAge()
	nuo.mutation.SetNurseAge(i)
	return nuo
}

// AddNurseAge adds i to nurse_age.
func (nuo *NurseUpdateOne) AddNurseAge(i int) *NurseUpdateOne {
	nuo.mutation.AddNurseAge(i)
	return nuo
}

// SetEmail sets the email field.
func (nuo *NurseUpdateOne) SetEmail(s string) *NurseUpdateOne {
	nuo.mutation.SetEmail(s)
	return nuo
}

// SetPassword sets the password field.
func (nuo *NurseUpdateOne) SetPassword(s string) *NurseUpdateOne {
	nuo.mutation.SetPassword(s)
	return nuo
}

// AddQueueIDs adds the queue edge to Queue by ids.
func (nuo *NurseUpdateOne) AddQueueIDs(ids ...int) *NurseUpdateOne {
	nuo.mutation.AddQueueIDs(ids...)
	return nuo
}

// AddQueue adds the queue edges to Queue.
func (nuo *NurseUpdateOne) AddQueue(q ...*Queue) *NurseUpdateOne {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return nuo.AddQueueIDs(ids...)
}

// AddMedicalfileIDs adds the medicalfiles edge to Medicalfile by ids.
func (nuo *NurseUpdateOne) AddMedicalfileIDs(ids ...int) *NurseUpdateOne {
	nuo.mutation.AddMedicalfileIDs(ids...)
	return nuo
}

// AddMedicalfiles adds the medicalfiles edges to Medicalfile.
func (nuo *NurseUpdateOne) AddMedicalfiles(m ...*Medicalfile) *NurseUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return nuo.AddMedicalfileIDs(ids...)
}

// AddDentalexpenseIDs adds the dentalexpenses edge to Dentalexpense by ids.
func (nuo *NurseUpdateOne) AddDentalexpenseIDs(ids ...int) *NurseUpdateOne {
	nuo.mutation.AddDentalexpenseIDs(ids...)
	return nuo
}

// AddDentalexpenses adds the dentalexpenses edges to Dentalexpense.
func (nuo *NurseUpdateOne) AddDentalexpenses(d ...*Dentalexpense) *NurseUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return nuo.AddDentalexpenseIDs(ids...)
}

// AddPatientIDs adds the patients edge to Patient by ids.
func (nuo *NurseUpdateOne) AddPatientIDs(ids ...int) *NurseUpdateOne {
	nuo.mutation.AddPatientIDs(ids...)
	return nuo
}

// AddPatients adds the patients edges to Patient.
func (nuo *NurseUpdateOne) AddPatients(p ...*Patient) *NurseUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return nuo.AddPatientIDs(ids...)
}

// AddDentistIDs adds the dentists edge to Dentist by ids.
func (nuo *NurseUpdateOne) AddDentistIDs(ids ...int) *NurseUpdateOne {
	nuo.mutation.AddDentistIDs(ids...)
	return nuo
}

// AddDentists adds the dentists edges to Dentist.
func (nuo *NurseUpdateOne) AddDentists(d ...*Dentist) *NurseUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return nuo.AddDentistIDs(ids...)
}

// AddAppointmentIDs adds the appointment edge to Appointment by ids.
func (nuo *NurseUpdateOne) AddAppointmentIDs(ids ...int) *NurseUpdateOne {
	nuo.mutation.AddAppointmentIDs(ids...)
	return nuo
}

// AddAppointment adds the appointment edges to Appointment.
func (nuo *NurseUpdateOne) AddAppointment(a ...*Appointment) *NurseUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return nuo.AddAppointmentIDs(ids...)
}

// Mutation returns the NurseMutation object of the builder.
func (nuo *NurseUpdateOne) Mutation() *NurseMutation {
	return nuo.mutation
}

// RemoveQueueIDs removes the queue edge to Queue by ids.
func (nuo *NurseUpdateOne) RemoveQueueIDs(ids ...int) *NurseUpdateOne {
	nuo.mutation.RemoveQueueIDs(ids...)
	return nuo
}

// RemoveQueue removes queue edges to Queue.
func (nuo *NurseUpdateOne) RemoveQueue(q ...*Queue) *NurseUpdateOne {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return nuo.RemoveQueueIDs(ids...)
}

// RemoveMedicalfileIDs removes the medicalfiles edge to Medicalfile by ids.
func (nuo *NurseUpdateOne) RemoveMedicalfileIDs(ids ...int) *NurseUpdateOne {
	nuo.mutation.RemoveMedicalfileIDs(ids...)
	return nuo
}

// RemoveMedicalfiles removes medicalfiles edges to Medicalfile.
func (nuo *NurseUpdateOne) RemoveMedicalfiles(m ...*Medicalfile) *NurseUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return nuo.RemoveMedicalfileIDs(ids...)
}

// RemoveDentalexpenseIDs removes the dentalexpenses edge to Dentalexpense by ids.
func (nuo *NurseUpdateOne) RemoveDentalexpenseIDs(ids ...int) *NurseUpdateOne {
	nuo.mutation.RemoveDentalexpenseIDs(ids...)
	return nuo
}

// RemoveDentalexpenses removes dentalexpenses edges to Dentalexpense.
func (nuo *NurseUpdateOne) RemoveDentalexpenses(d ...*Dentalexpense) *NurseUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return nuo.RemoveDentalexpenseIDs(ids...)
}

// RemovePatientIDs removes the patients edge to Patient by ids.
func (nuo *NurseUpdateOne) RemovePatientIDs(ids ...int) *NurseUpdateOne {
	nuo.mutation.RemovePatientIDs(ids...)
	return nuo
}

// RemovePatients removes patients edges to Patient.
func (nuo *NurseUpdateOne) RemovePatients(p ...*Patient) *NurseUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return nuo.RemovePatientIDs(ids...)
}

// RemoveDentistIDs removes the dentists edge to Dentist by ids.
func (nuo *NurseUpdateOne) RemoveDentistIDs(ids ...int) *NurseUpdateOne {
	nuo.mutation.RemoveDentistIDs(ids...)
	return nuo
}

// RemoveDentists removes dentists edges to Dentist.
func (nuo *NurseUpdateOne) RemoveDentists(d ...*Dentist) *NurseUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return nuo.RemoveDentistIDs(ids...)
}

// RemoveAppointmentIDs removes the appointment edge to Appointment by ids.
func (nuo *NurseUpdateOne) RemoveAppointmentIDs(ids ...int) *NurseUpdateOne {
	nuo.mutation.RemoveAppointmentIDs(ids...)
	return nuo
}

// RemoveAppointment removes appointment edges to Appointment.
func (nuo *NurseUpdateOne) RemoveAppointment(a ...*Appointment) *NurseUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return nuo.RemoveAppointmentIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (nuo *NurseUpdateOne) Save(ctx context.Context) (*Nurse, error) {
	if v, ok := nuo.mutation.NurseName(); ok {
		if err := nurse.NurseNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "nurse_name", err: fmt.Errorf("ent: validator failed for field \"nurse_name\": %w", err)}
		}
	}

	var (
		err  error
		node *Nurse
	)
	if len(nuo.hooks) == 0 {
		node, err = nuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NurseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nuo.mutation = mutation
			node, err = nuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(nuo.hooks) - 1; i >= 0; i-- {
			mut = nuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NurseUpdateOne) SaveX(ctx context.Context) *Nurse {
	n, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

// Exec executes the query on the entity.
func (nuo *NurseUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NurseUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nuo *NurseUpdateOne) sqlSave(ctx context.Context) (n *Nurse, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   nurse.Table,
			Columns: nurse.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: nurse.FieldID,
			},
		},
	}
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Nurse.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := nuo.mutation.NurseName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nurse.FieldNurseName,
		})
	}
	if value, ok := nuo.mutation.NurseAge(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: nurse.FieldNurseAge,
		})
	}
	if value, ok := nuo.mutation.AddedNurseAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: nurse.FieldNurseAge,
		})
	}
	if value, ok := nuo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nurse.FieldEmail,
		})
	}
	if value, ok := nuo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nurse.FieldPassword,
		})
	}
	if nodes := nuo.mutation.RemovedQueueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.QueueTable,
			Columns: []string{nurse.QueueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: queue.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.QueueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.QueueTable,
			Columns: []string{nurse.QueueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: queue.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := nuo.mutation.RemovedMedicalfilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.MedicalfilesTable,
			Columns: []string{nurse.MedicalfilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicalfile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.MedicalfilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.MedicalfilesTable,
			Columns: []string{nurse.MedicalfilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicalfile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := nuo.mutation.RemovedDentalexpensesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.DentalexpensesTable,
			Columns: []string{nurse.DentalexpensesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dentalexpense.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.DentalexpensesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.DentalexpensesTable,
			Columns: []string{nurse.DentalexpensesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dentalexpense.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := nuo.mutation.RemovedPatientsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.PatientsTable,
			Columns: []string{nurse.PatientsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.PatientsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.PatientsTable,
			Columns: []string{nurse.PatientsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := nuo.mutation.RemovedDentistsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.DentistsTable,
			Columns: []string{nurse.DentistsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dentist.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.DentistsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.DentistsTable,
			Columns: []string{nurse.DentistsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dentist.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := nuo.mutation.RemovedAppointmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.AppointmentTable,
			Columns: []string{nurse.AppointmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: appointment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.AppointmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.AppointmentTable,
			Columns: []string{nurse.AppointmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: appointment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	n = &Nurse{config: nuo.config}
	_spec.Assign = n.assignValues
	_spec.ScanValues = n.scanValues()
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nurse.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return n, nil
}
