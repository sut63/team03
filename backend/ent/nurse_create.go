// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team03/app/ent/dentalexpense"
	"github.com/team03/app/ent/dentist"
	"github.com/team03/app/ent/medicalfile"
	"github.com/team03/app/ent/nurse"
	"github.com/team03/app/ent/patient"
	"github.com/team03/app/ent/queue"
)

// NurseCreate is the builder for creating a Nurse entity.
type NurseCreate struct {
	config
	mutation *NurseMutation
	hooks    []Hook
}

// SetNurseName sets the nurse_name field.
func (nc *NurseCreate) SetNurseName(s string) *NurseCreate {
	nc.mutation.SetNurseName(s)
	return nc
}

// SetNurseAge sets the nurse_age field.
func (nc *NurseCreate) SetNurseAge(i int) *NurseCreate {
	nc.mutation.SetNurseAge(i)
	return nc
}

// SetNurseEmail sets the nurse_email field.
func (nc *NurseCreate) SetNurseEmail(s string) *NurseCreate {
	nc.mutation.SetNurseEmail(s)
	return nc
}

// SetNursePassword sets the nurse_password field.
func (nc *NurseCreate) SetNursePassword(s string) *NurseCreate {
	nc.mutation.SetNursePassword(s)
	return nc
}

// AddQueueIDs adds the queue edge to Queue by ids.
func (nc *NurseCreate) AddQueueIDs(ids ...int) *NurseCreate {
	nc.mutation.AddQueueIDs(ids...)
	return nc
}

// AddQueue adds the queue edges to Queue.
func (nc *NurseCreate) AddQueue(q ...*Queue) *NurseCreate {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return nc.AddQueueIDs(ids...)
}

// AddMedicalfileIDs adds the medicalfiles edge to Medicalfile by ids.
func (nc *NurseCreate) AddMedicalfileIDs(ids ...int) *NurseCreate {
	nc.mutation.AddMedicalfileIDs(ids...)
	return nc
}

// AddMedicalfiles adds the medicalfiles edges to Medicalfile.
func (nc *NurseCreate) AddMedicalfiles(m ...*Medicalfile) *NurseCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return nc.AddMedicalfileIDs(ids...)
}

// AddDentalexpenseIDs adds the dentalexpenses edge to DentalExpense by ids.
func (nc *NurseCreate) AddDentalexpenseIDs(ids ...int) *NurseCreate {
	nc.mutation.AddDentalexpenseIDs(ids...)
	return nc
}

// AddDentalexpenses adds the dentalexpenses edges to DentalExpense.
func (nc *NurseCreate) AddDentalexpenses(d ...*DentalExpense) *NurseCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return nc.AddDentalexpenseIDs(ids...)
}

// AddPatientIDs adds the patients edge to Patient by ids.
func (nc *NurseCreate) AddPatientIDs(ids ...int) *NurseCreate {
	nc.mutation.AddPatientIDs(ids...)
	return nc
}

// AddPatients adds the patients edges to Patient.
func (nc *NurseCreate) AddPatients(p ...*Patient) *NurseCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return nc.AddPatientIDs(ids...)
}

// AddDentistIDs adds the dentists edge to Dentist by ids.
func (nc *NurseCreate) AddDentistIDs(ids ...int) *NurseCreate {
	nc.mutation.AddDentistIDs(ids...)
	return nc
}

// AddDentists adds the dentists edges to Dentist.
func (nc *NurseCreate) AddDentists(d ...*Dentist) *NurseCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return nc.AddDentistIDs(ids...)
}

// Mutation returns the NurseMutation object of the builder.
func (nc *NurseCreate) Mutation() *NurseMutation {
	return nc.mutation
}

// Save creates the Nurse in the database.
func (nc *NurseCreate) Save(ctx context.Context) (*Nurse, error) {
	if _, ok := nc.mutation.NurseName(); !ok {
		return nil, &ValidationError{Name: "nurse_name", err: errors.New("ent: missing required field \"nurse_name\"")}
	}
	if v, ok := nc.mutation.NurseName(); ok {
		if err := nurse.NurseNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "nurse_name", err: fmt.Errorf("ent: validator failed for field \"nurse_name\": %w", err)}
		}
	}
	if _, ok := nc.mutation.NurseAge(); !ok {
		return nil, &ValidationError{Name: "nurse_age", err: errors.New("ent: missing required field \"nurse_age\"")}
	}
	if _, ok := nc.mutation.NurseEmail(); !ok {
		return nil, &ValidationError{Name: "nurse_email", err: errors.New("ent: missing required field \"nurse_email\"")}
	}
	if _, ok := nc.mutation.NursePassword(); !ok {
		return nil, &ValidationError{Name: "nurse_password", err: errors.New("ent: missing required field \"nurse_password\"")}
	}
	var (
		err  error
		node *Nurse
	)
	if len(nc.hooks) == 0 {
		node, err = nc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NurseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nc.mutation = mutation
			node, err = nc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(nc.hooks) - 1; i >= 0; i-- {
			mut = nc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NurseCreate) SaveX(ctx context.Context) *Nurse {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (nc *NurseCreate) sqlSave(ctx context.Context) (*Nurse, error) {
	n, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	n.ID = int(id)
	return n, nil
}

func (nc *NurseCreate) createSpec() (*Nurse, *sqlgraph.CreateSpec) {
	var (
		n     = &Nurse{config: nc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: nurse.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: nurse.FieldID,
			},
		}
	)
	if value, ok := nc.mutation.NurseName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nurse.FieldNurseName,
		})
		n.NurseName = value
	}
	if value, ok := nc.mutation.NurseAge(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: nurse.FieldNurseAge,
		})
		n.NurseAge = value
	}
	if value, ok := nc.mutation.NurseEmail(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nurse.FieldNurseEmail,
		})
		n.NurseEmail = value
	}
	if value, ok := nc.mutation.NursePassword(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nurse.FieldNursePassword,
		})
		n.NursePassword = value
	}
	if nodes := nc.mutation.QueueIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.MedicalfilesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.DentalexpensesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.PatientsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.DentistsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return n, _spec
}
