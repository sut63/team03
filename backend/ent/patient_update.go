// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team03/app/ent/appointment"
	"github.com/team03/app/ent/disease"
	"github.com/team03/app/ent/gender"
	"github.com/team03/app/ent/medicalcare"
	"github.com/team03/app/ent/medicalfile"
	"github.com/team03/app/ent/nurse"
	"github.com/team03/app/ent/patient"
	"github.com/team03/app/ent/predicate"
	"github.com/team03/app/ent/queue"
)

// PatientUpdate is the builder for updating Patient entities.
type PatientUpdate struct {
	config
	hooks      []Hook
	mutation   *PatientMutation
	predicates []predicate.Patient
}

// Where adds a new predicate for the builder.
func (pu *PatientUpdate) Where(ps ...predicate.Patient) *PatientUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetPatientID sets the patient_ID field.
func (pu *PatientUpdate) SetPatientID(s string) *PatientUpdate {
	pu.mutation.SetPatientID(s)
	return pu
}

// SetName sets the name field.
func (pu *PatientUpdate) SetName(s string) *PatientUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetCardID sets the cardID field.
func (pu *PatientUpdate) SetCardID(s string) *PatientUpdate {
	pu.mutation.SetCardID(s)
	return pu
}

// SetTel sets the tel field.
func (pu *PatientUpdate) SetTel(s string) *PatientUpdate {
	pu.mutation.SetTel(s)
	return pu
}

// SetAge sets the age field.
func (pu *PatientUpdate) SetAge(i int) *PatientUpdate {
	pu.mutation.ResetAge()
	pu.mutation.SetAge(i)
	return pu
}

// AddAge adds i to age.
func (pu *PatientUpdate) AddAge(i int) *PatientUpdate {
	pu.mutation.AddAge(i)
	return pu
}

// SetBirthday sets the birthday field.
func (pu *PatientUpdate) SetBirthday(t time.Time) *PatientUpdate {
	pu.mutation.SetBirthday(t)
	return pu
}

// SetGenderID sets the gender edge to Gender by id.
func (pu *PatientUpdate) SetGenderID(id int) *PatientUpdate {
	pu.mutation.SetGenderID(id)
	return pu
}

// SetNillableGenderID sets the gender edge to Gender by id if the given value is not nil.
func (pu *PatientUpdate) SetNillableGenderID(id *int) *PatientUpdate {
	if id != nil {
		pu = pu.SetGenderID(*id)
	}
	return pu
}

// SetGender sets the gender edge to Gender.
func (pu *PatientUpdate) SetGender(g *Gender) *PatientUpdate {
	return pu.SetGenderID(g.ID)
}

// SetMedicalcareID sets the medicalcare edge to MedicalCare by id.
func (pu *PatientUpdate) SetMedicalcareID(id int) *PatientUpdate {
	pu.mutation.SetMedicalcareID(id)
	return pu
}

// SetNillableMedicalcareID sets the medicalcare edge to MedicalCare by id if the given value is not nil.
func (pu *PatientUpdate) SetNillableMedicalcareID(id *int) *PatientUpdate {
	if id != nil {
		pu = pu.SetMedicalcareID(*id)
	}
	return pu
}

// SetMedicalcare sets the medicalcare edge to MedicalCare.
func (pu *PatientUpdate) SetMedicalcare(m *MedicalCare) *PatientUpdate {
	return pu.SetMedicalcareID(m.ID)
}

// SetNurseID sets the nurse edge to Nurse by id.
func (pu *PatientUpdate) SetNurseID(id int) *PatientUpdate {
	pu.mutation.SetNurseID(id)
	return pu
}

// SetNillableNurseID sets the nurse edge to Nurse by id if the given value is not nil.
func (pu *PatientUpdate) SetNillableNurseID(id *int) *PatientUpdate {
	if id != nil {
		pu = pu.SetNurseID(*id)
	}
	return pu
}

// SetNurse sets the nurse edge to Nurse.
func (pu *PatientUpdate) SetNurse(n *Nurse) *PatientUpdate {
	return pu.SetNurseID(n.ID)
}

// SetDiseaseID sets the disease edge to Disease by id.
func (pu *PatientUpdate) SetDiseaseID(id int) *PatientUpdate {
	pu.mutation.SetDiseaseID(id)
	return pu
}

// SetNillableDiseaseID sets the disease edge to Disease by id if the given value is not nil.
func (pu *PatientUpdate) SetNillableDiseaseID(id *int) *PatientUpdate {
	if id != nil {
		pu = pu.SetDiseaseID(*id)
	}
	return pu
}

// SetDisease sets the disease edge to Disease.
func (pu *PatientUpdate) SetDisease(d *Disease) *PatientUpdate {
	return pu.SetDiseaseID(d.ID)
}

// AddMedicalfileIDs adds the medicalfiles edge to Medicalfile by ids.
func (pu *PatientUpdate) AddMedicalfileIDs(ids ...int) *PatientUpdate {
	pu.mutation.AddMedicalfileIDs(ids...)
	return pu
}

// AddMedicalfiles adds the medicalfiles edges to Medicalfile.
func (pu *PatientUpdate) AddMedicalfiles(m ...*Medicalfile) *PatientUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return pu.AddMedicalfileIDs(ids...)
}

// AddQueueIDs adds the queue edge to Queue by ids.
func (pu *PatientUpdate) AddQueueIDs(ids ...int) *PatientUpdate {
	pu.mutation.AddQueueIDs(ids...)
	return pu
}

// AddQueue adds the queue edges to Queue.
func (pu *PatientUpdate) AddQueue(q ...*Queue) *PatientUpdate {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return pu.AddQueueIDs(ids...)
}

// AddAppointmentIDs adds the appointment edge to Appointment by ids.
func (pu *PatientUpdate) AddAppointmentIDs(ids ...int) *PatientUpdate {
	pu.mutation.AddAppointmentIDs(ids...)
	return pu
}

// AddAppointment adds the appointment edges to Appointment.
func (pu *PatientUpdate) AddAppointment(a ...*Appointment) *PatientUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pu.AddAppointmentIDs(ids...)
}

// Mutation returns the PatientMutation object of the builder.
func (pu *PatientUpdate) Mutation() *PatientMutation {
	return pu.mutation
}

// ClearGender clears the gender edge to Gender.
func (pu *PatientUpdate) ClearGender() *PatientUpdate {
	pu.mutation.ClearGender()
	return pu
}

// ClearMedicalcare clears the medicalcare edge to MedicalCare.
func (pu *PatientUpdate) ClearMedicalcare() *PatientUpdate {
	pu.mutation.ClearMedicalcare()
	return pu
}

// ClearNurse clears the nurse edge to Nurse.
func (pu *PatientUpdate) ClearNurse() *PatientUpdate {
	pu.mutation.ClearNurse()
	return pu
}

// ClearDisease clears the disease edge to Disease.
func (pu *PatientUpdate) ClearDisease() *PatientUpdate {
	pu.mutation.ClearDisease()
	return pu
}

// RemoveMedicalfileIDs removes the medicalfiles edge to Medicalfile by ids.
func (pu *PatientUpdate) RemoveMedicalfileIDs(ids ...int) *PatientUpdate {
	pu.mutation.RemoveMedicalfileIDs(ids...)
	return pu
}

// RemoveMedicalfiles removes medicalfiles edges to Medicalfile.
func (pu *PatientUpdate) RemoveMedicalfiles(m ...*Medicalfile) *PatientUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return pu.RemoveMedicalfileIDs(ids...)
}

// RemoveQueueIDs removes the queue edge to Queue by ids.
func (pu *PatientUpdate) RemoveQueueIDs(ids ...int) *PatientUpdate {
	pu.mutation.RemoveQueueIDs(ids...)
	return pu
}

// RemoveQueue removes queue edges to Queue.
func (pu *PatientUpdate) RemoveQueue(q ...*Queue) *PatientUpdate {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return pu.RemoveQueueIDs(ids...)
}

// RemoveAppointmentIDs removes the appointment edge to Appointment by ids.
func (pu *PatientUpdate) RemoveAppointmentIDs(ids ...int) *PatientUpdate {
	pu.mutation.RemoveAppointmentIDs(ids...)
	return pu
}

// RemoveAppointment removes appointment edges to Appointment.
func (pu *PatientUpdate) RemoveAppointment(a ...*Appointment) *PatientUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pu.RemoveAppointmentIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *PatientUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := pu.mutation.PatientID(); ok {
		if err := patient.PatientIDValidator(v); err != nil {
			return 0, &ValidationError{Name: "patient_ID", err: fmt.Errorf("ent: validator failed for field \"patient_ID\": %w", err)}
		}
	}
	if v, ok := pu.mutation.Name(); ok {
		if err := patient.NameValidator(v); err != nil {
			return 0, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := pu.mutation.CardID(); ok {
		if err := patient.CardIDValidator(v); err != nil {
			return 0, &ValidationError{Name: "cardID", err: fmt.Errorf("ent: validator failed for field \"cardID\": %w", err)}
		}
	}
	if v, ok := pu.mutation.Tel(); ok {
		if err := patient.TelValidator(v); err != nil {
			return 0, &ValidationError{Name: "tel", err: fmt.Errorf("ent: validator failed for field \"tel\": %w", err)}
		}
	}
	if v, ok := pu.mutation.Age(); ok {
		if err := patient.AgeValidator(v); err != nil {
			return 0, &ValidationError{Name: "age", err: fmt.Errorf("ent: validator failed for field \"age\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PatientMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PatientUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PatientUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PatientUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PatientUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   patient.Table,
			Columns: patient.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patient.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.PatientID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldPatientID,
		})
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldName,
		})
	}
	if value, ok := pu.mutation.CardID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldCardID,
		})
	}
	if value, ok := pu.mutation.Tel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldTel,
		})
	}
	if value, ok := pu.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: patient.FieldAge,
		})
	}
	if value, ok := pu.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: patient.FieldAge,
		})
	}
	if value, ok := pu.mutation.Birthday(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: patient.FieldBirthday,
		})
	}
	if pu.mutation.GenderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.GenderTable,
			Columns: []string{patient.GenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gender.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.GenderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.GenderTable,
			Columns: []string{patient.GenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gender.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.MedicalcareCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.MedicalcareTable,
			Columns: []string{patient.MedicalcareColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicalcare.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.MedicalcareIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.MedicalcareTable,
			Columns: []string{patient.MedicalcareColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicalcare.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.NurseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.NurseTable,
			Columns: []string{patient.NurseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nurse.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.NurseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.NurseTable,
			Columns: []string{patient.NurseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nurse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.DiseaseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.DiseaseTable,
			Columns: []string{patient.DiseaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: disease.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.DiseaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.DiseaseTable,
			Columns: []string{patient.DiseaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: disease.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := pu.mutation.RemovedMedicalfilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.MedicalfilesTable,
			Columns: []string{patient.MedicalfilesColumn},
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
	if nodes := pu.mutation.MedicalfilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.MedicalfilesTable,
			Columns: []string{patient.MedicalfilesColumn},
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
	if nodes := pu.mutation.RemovedQueueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.QueueTable,
			Columns: []string{patient.QueueColumn},
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
	if nodes := pu.mutation.QueueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.QueueTable,
			Columns: []string{patient.QueueColumn},
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
	if nodes := pu.mutation.RemovedAppointmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.AppointmentTable,
			Columns: []string{patient.AppointmentColumn},
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
	if nodes := pu.mutation.AppointmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.AppointmentTable,
			Columns: []string{patient.AppointmentColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{patient.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PatientUpdateOne is the builder for updating a single Patient entity.
type PatientUpdateOne struct {
	config
	hooks    []Hook
	mutation *PatientMutation
}

// SetPatientID sets the patient_ID field.
func (puo *PatientUpdateOne) SetPatientID(s string) *PatientUpdateOne {
	puo.mutation.SetPatientID(s)
	return puo
}

// SetName sets the name field.
func (puo *PatientUpdateOne) SetName(s string) *PatientUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetCardID sets the cardID field.
func (puo *PatientUpdateOne) SetCardID(s string) *PatientUpdateOne {
	puo.mutation.SetCardID(s)
	return puo
}

// SetTel sets the tel field.
func (puo *PatientUpdateOne) SetTel(s string) *PatientUpdateOne {
	puo.mutation.SetTel(s)
	return puo
}

// SetAge sets the age field.
func (puo *PatientUpdateOne) SetAge(i int) *PatientUpdateOne {
	puo.mutation.ResetAge()
	puo.mutation.SetAge(i)
	return puo
}

// AddAge adds i to age.
func (puo *PatientUpdateOne) AddAge(i int) *PatientUpdateOne {
	puo.mutation.AddAge(i)
	return puo
}

// SetBirthday sets the birthday field.
func (puo *PatientUpdateOne) SetBirthday(t time.Time) *PatientUpdateOne {
	puo.mutation.SetBirthday(t)
	return puo
}

// SetGenderID sets the gender edge to Gender by id.
func (puo *PatientUpdateOne) SetGenderID(id int) *PatientUpdateOne {
	puo.mutation.SetGenderID(id)
	return puo
}

// SetNillableGenderID sets the gender edge to Gender by id if the given value is not nil.
func (puo *PatientUpdateOne) SetNillableGenderID(id *int) *PatientUpdateOne {
	if id != nil {
		puo = puo.SetGenderID(*id)
	}
	return puo
}

// SetGender sets the gender edge to Gender.
func (puo *PatientUpdateOne) SetGender(g *Gender) *PatientUpdateOne {
	return puo.SetGenderID(g.ID)
}

// SetMedicalcareID sets the medicalcare edge to MedicalCare by id.
func (puo *PatientUpdateOne) SetMedicalcareID(id int) *PatientUpdateOne {
	puo.mutation.SetMedicalcareID(id)
	return puo
}

// SetNillableMedicalcareID sets the medicalcare edge to MedicalCare by id if the given value is not nil.
func (puo *PatientUpdateOne) SetNillableMedicalcareID(id *int) *PatientUpdateOne {
	if id != nil {
		puo = puo.SetMedicalcareID(*id)
	}
	return puo
}

// SetMedicalcare sets the medicalcare edge to MedicalCare.
func (puo *PatientUpdateOne) SetMedicalcare(m *MedicalCare) *PatientUpdateOne {
	return puo.SetMedicalcareID(m.ID)
}

// SetNurseID sets the nurse edge to Nurse by id.
func (puo *PatientUpdateOne) SetNurseID(id int) *PatientUpdateOne {
	puo.mutation.SetNurseID(id)
	return puo
}

// SetNillableNurseID sets the nurse edge to Nurse by id if the given value is not nil.
func (puo *PatientUpdateOne) SetNillableNurseID(id *int) *PatientUpdateOne {
	if id != nil {
		puo = puo.SetNurseID(*id)
	}
	return puo
}

// SetNurse sets the nurse edge to Nurse.
func (puo *PatientUpdateOne) SetNurse(n *Nurse) *PatientUpdateOne {
	return puo.SetNurseID(n.ID)
}

// SetDiseaseID sets the disease edge to Disease by id.
func (puo *PatientUpdateOne) SetDiseaseID(id int) *PatientUpdateOne {
	puo.mutation.SetDiseaseID(id)
	return puo
}

// SetNillableDiseaseID sets the disease edge to Disease by id if the given value is not nil.
func (puo *PatientUpdateOne) SetNillableDiseaseID(id *int) *PatientUpdateOne {
	if id != nil {
		puo = puo.SetDiseaseID(*id)
	}
	return puo
}

// SetDisease sets the disease edge to Disease.
func (puo *PatientUpdateOne) SetDisease(d *Disease) *PatientUpdateOne {
	return puo.SetDiseaseID(d.ID)
}

// AddMedicalfileIDs adds the medicalfiles edge to Medicalfile by ids.
func (puo *PatientUpdateOne) AddMedicalfileIDs(ids ...int) *PatientUpdateOne {
	puo.mutation.AddMedicalfileIDs(ids...)
	return puo
}

// AddMedicalfiles adds the medicalfiles edges to Medicalfile.
func (puo *PatientUpdateOne) AddMedicalfiles(m ...*Medicalfile) *PatientUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return puo.AddMedicalfileIDs(ids...)
}

// AddQueueIDs adds the queue edge to Queue by ids.
func (puo *PatientUpdateOne) AddQueueIDs(ids ...int) *PatientUpdateOne {
	puo.mutation.AddQueueIDs(ids...)
	return puo
}

// AddQueue adds the queue edges to Queue.
func (puo *PatientUpdateOne) AddQueue(q ...*Queue) *PatientUpdateOne {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return puo.AddQueueIDs(ids...)
}

// AddAppointmentIDs adds the appointment edge to Appointment by ids.
func (puo *PatientUpdateOne) AddAppointmentIDs(ids ...int) *PatientUpdateOne {
	puo.mutation.AddAppointmentIDs(ids...)
	return puo
}

// AddAppointment adds the appointment edges to Appointment.
func (puo *PatientUpdateOne) AddAppointment(a ...*Appointment) *PatientUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return puo.AddAppointmentIDs(ids...)
}

// Mutation returns the PatientMutation object of the builder.
func (puo *PatientUpdateOne) Mutation() *PatientMutation {
	return puo.mutation
}

// ClearGender clears the gender edge to Gender.
func (puo *PatientUpdateOne) ClearGender() *PatientUpdateOne {
	puo.mutation.ClearGender()
	return puo
}

// ClearMedicalcare clears the medicalcare edge to MedicalCare.
func (puo *PatientUpdateOne) ClearMedicalcare() *PatientUpdateOne {
	puo.mutation.ClearMedicalcare()
	return puo
}

// ClearNurse clears the nurse edge to Nurse.
func (puo *PatientUpdateOne) ClearNurse() *PatientUpdateOne {
	puo.mutation.ClearNurse()
	return puo
}

// ClearDisease clears the disease edge to Disease.
func (puo *PatientUpdateOne) ClearDisease() *PatientUpdateOne {
	puo.mutation.ClearDisease()
	return puo
}

// RemoveMedicalfileIDs removes the medicalfiles edge to Medicalfile by ids.
func (puo *PatientUpdateOne) RemoveMedicalfileIDs(ids ...int) *PatientUpdateOne {
	puo.mutation.RemoveMedicalfileIDs(ids...)
	return puo
}

// RemoveMedicalfiles removes medicalfiles edges to Medicalfile.
func (puo *PatientUpdateOne) RemoveMedicalfiles(m ...*Medicalfile) *PatientUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return puo.RemoveMedicalfileIDs(ids...)
}

// RemoveQueueIDs removes the queue edge to Queue by ids.
func (puo *PatientUpdateOne) RemoveQueueIDs(ids ...int) *PatientUpdateOne {
	puo.mutation.RemoveQueueIDs(ids...)
	return puo
}

// RemoveQueue removes queue edges to Queue.
func (puo *PatientUpdateOne) RemoveQueue(q ...*Queue) *PatientUpdateOne {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return puo.RemoveQueueIDs(ids...)
}

// RemoveAppointmentIDs removes the appointment edge to Appointment by ids.
func (puo *PatientUpdateOne) RemoveAppointmentIDs(ids ...int) *PatientUpdateOne {
	puo.mutation.RemoveAppointmentIDs(ids...)
	return puo
}

// RemoveAppointment removes appointment edges to Appointment.
func (puo *PatientUpdateOne) RemoveAppointment(a ...*Appointment) *PatientUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return puo.RemoveAppointmentIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (puo *PatientUpdateOne) Save(ctx context.Context) (*Patient, error) {
	if v, ok := puo.mutation.PatientID(); ok {
		if err := patient.PatientIDValidator(v); err != nil {
			return nil, &ValidationError{Name: "patient_ID", err: fmt.Errorf("ent: validator failed for field \"patient_ID\": %w", err)}
		}
	}
	if v, ok := puo.mutation.Name(); ok {
		if err := patient.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := puo.mutation.CardID(); ok {
		if err := patient.CardIDValidator(v); err != nil {
			return nil, &ValidationError{Name: "cardID", err: fmt.Errorf("ent: validator failed for field \"cardID\": %w", err)}
		}
	}
	if v, ok := puo.mutation.Tel(); ok {
		if err := patient.TelValidator(v); err != nil {
			return nil, &ValidationError{Name: "tel", err: fmt.Errorf("ent: validator failed for field \"tel\": %w", err)}
		}
	}
	if v, ok := puo.mutation.Age(); ok {
		if err := patient.AgeValidator(v); err != nil {
			return nil, &ValidationError{Name: "age", err: fmt.Errorf("ent: validator failed for field \"age\": %w", err)}
		}
	}

	var (
		err  error
		node *Patient
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PatientMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PatientUpdateOne) SaveX(ctx context.Context) *Patient {
	pa, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pa
}

// Exec executes the query on the entity.
func (puo *PatientUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PatientUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PatientUpdateOne) sqlSave(ctx context.Context) (pa *Patient, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   patient.Table,
			Columns: patient.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patient.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Patient.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.PatientID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldPatientID,
		})
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldName,
		})
	}
	if value, ok := puo.mutation.CardID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldCardID,
		})
	}
	if value, ok := puo.mutation.Tel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldTel,
		})
	}
	if value, ok := puo.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: patient.FieldAge,
		})
	}
	if value, ok := puo.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: patient.FieldAge,
		})
	}
	if value, ok := puo.mutation.Birthday(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: patient.FieldBirthday,
		})
	}
	if puo.mutation.GenderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.GenderTable,
			Columns: []string{patient.GenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gender.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.GenderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.GenderTable,
			Columns: []string{patient.GenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gender.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.MedicalcareCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.MedicalcareTable,
			Columns: []string{patient.MedicalcareColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicalcare.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.MedicalcareIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.MedicalcareTable,
			Columns: []string{patient.MedicalcareColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicalcare.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.NurseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.NurseTable,
			Columns: []string{patient.NurseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nurse.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.NurseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.NurseTable,
			Columns: []string{patient.NurseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nurse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.DiseaseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.DiseaseTable,
			Columns: []string{patient.DiseaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: disease.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.DiseaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.DiseaseTable,
			Columns: []string{patient.DiseaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: disease.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := puo.mutation.RemovedMedicalfilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.MedicalfilesTable,
			Columns: []string{patient.MedicalfilesColumn},
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
	if nodes := puo.mutation.MedicalfilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.MedicalfilesTable,
			Columns: []string{patient.MedicalfilesColumn},
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
	if nodes := puo.mutation.RemovedQueueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.QueueTable,
			Columns: []string{patient.QueueColumn},
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
	if nodes := puo.mutation.QueueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.QueueTable,
			Columns: []string{patient.QueueColumn},
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
	if nodes := puo.mutation.RemovedAppointmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.AppointmentTable,
			Columns: []string{patient.AppointmentColumn},
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
	if nodes := puo.mutation.AppointmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.AppointmentTable,
			Columns: []string{patient.AppointmentColumn},
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
	pa = &Patient{config: puo.config}
	_spec.Assign = pa.assignValues
	_spec.ScanValues = pa.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{patient.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pa, nil
}
