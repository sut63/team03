// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/tanapon395/playlist-video/ent/dentalexpense"
	"github.com/tanapon395/playlist-video/ent/dentist"
	"github.com/tanapon395/playlist-video/ent/medicalcare"
	"github.com/tanapon395/playlist-video/ent/medicalfile"
	"github.com/tanapon395/playlist-video/ent/nurse"
	"github.com/tanapon395/playlist-video/ent/patient"
)

// MedicalfileCreate is the builder for creating a Medicalfile entity.
type MedicalfileCreate struct {
	config
	mutation *MedicalfileMutation
	hooks    []Hook
}

// SetDetail sets the detail field.
func (mc *MedicalfileCreate) SetDetail(s string) *MedicalfileCreate {
	mc.mutation.SetDetail(s)
	return mc
}

// SetAddedTime sets the added_time field.
func (mc *MedicalfileCreate) SetAddedTime(t time.Time) *MedicalfileCreate {
	mc.mutation.SetAddedTime(t)
	return mc
}

// SetDentistID sets the dentist edge to Dentist by id.
func (mc *MedicalfileCreate) SetDentistID(id int) *MedicalfileCreate {
	mc.mutation.SetDentistID(id)
	return mc
}

// SetNillableDentistID sets the dentist edge to Dentist by id if the given value is not nil.
func (mc *MedicalfileCreate) SetNillableDentistID(id *int) *MedicalfileCreate {
	if id != nil {
		mc = mc.SetDentistID(*id)
	}
	return mc
}

// SetDentist sets the dentist edge to Dentist.
func (mc *MedicalfileCreate) SetDentist(d *Dentist) *MedicalfileCreate {
	return mc.SetDentistID(d.ID)
}

// SetPatientID sets the patient edge to Patient by id.
func (mc *MedicalfileCreate) SetPatientID(id int) *MedicalfileCreate {
	mc.mutation.SetPatientID(id)
	return mc
}

// SetNillablePatientID sets the patient edge to Patient by id if the given value is not nil.
func (mc *MedicalfileCreate) SetNillablePatientID(id *int) *MedicalfileCreate {
	if id != nil {
		mc = mc.SetPatientID(*id)
	}
	return mc
}

// SetPatient sets the patient edge to Patient.
func (mc *MedicalfileCreate) SetPatient(p *Patient) *MedicalfileCreate {
	return mc.SetPatientID(p.ID)
}

// SetNurseID sets the nurse edge to Nurse by id.
func (mc *MedicalfileCreate) SetNurseID(id int) *MedicalfileCreate {
	mc.mutation.SetNurseID(id)
	return mc
}

// SetNillableNurseID sets the nurse edge to Nurse by id if the given value is not nil.
func (mc *MedicalfileCreate) SetNillableNurseID(id *int) *MedicalfileCreate {
	if id != nil {
		mc = mc.SetNurseID(*id)
	}
	return mc
}

// SetNurse sets the nurse edge to Nurse.
func (mc *MedicalfileCreate) SetNurse(n *Nurse) *MedicalfileCreate {
	return mc.SetNurseID(n.ID)
}

// SetMedicalcareID sets the medicalcare edge to MedicalCare by id.
func (mc *MedicalfileCreate) SetMedicalcareID(id int) *MedicalfileCreate {
	mc.mutation.SetMedicalcareID(id)
	return mc
}

// SetNillableMedicalcareID sets the medicalcare edge to MedicalCare by id if the given value is not nil.
func (mc *MedicalfileCreate) SetNillableMedicalcareID(id *int) *MedicalfileCreate {
	if id != nil {
		mc = mc.SetMedicalcareID(*id)
	}
	return mc
}

// SetMedicalcare sets the medicalcare edge to MedicalCare.
func (mc *MedicalfileCreate) SetMedicalcare(m *MedicalCare) *MedicalfileCreate {
	return mc.SetMedicalcareID(m.ID)
}

// AddDentalexpenseIDs adds the dentalexpenses edge to DentalExpense by ids.
func (mc *MedicalfileCreate) AddDentalexpenseIDs(ids ...int) *MedicalfileCreate {
	mc.mutation.AddDentalexpenseIDs(ids...)
	return mc
}

// AddDentalexpenses adds the dentalexpenses edges to DentalExpense.
func (mc *MedicalfileCreate) AddDentalexpenses(d ...*DentalExpense) *MedicalfileCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return mc.AddDentalexpenseIDs(ids...)
}

// Mutation returns the MedicalfileMutation object of the builder.
func (mc *MedicalfileCreate) Mutation() *MedicalfileMutation {
	return mc.mutation
}

// Save creates the Medicalfile in the database.
func (mc *MedicalfileCreate) Save(ctx context.Context) (*Medicalfile, error) {
	if _, ok := mc.mutation.Detail(); !ok {
		return nil, &ValidationError{Name: "detail", err: errors.New("ent: missing required field \"detail\"")}
	}
	if v, ok := mc.mutation.Detail(); ok {
		if err := medicalfile.DetailValidator(v); err != nil {
			return nil, &ValidationError{Name: "detail", err: fmt.Errorf("ent: validator failed for field \"detail\": %w", err)}
		}
	}
	if _, ok := mc.mutation.AddedTime(); !ok {
		return nil, &ValidationError{Name: "added_time", err: errors.New("ent: missing required field \"added_time\"")}
	}
	var (
		err  error
		node *Medicalfile
	)
	if len(mc.hooks) == 0 {
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MedicalfileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mc.mutation = mutation
			node, err = mc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MedicalfileCreate) SaveX(ctx context.Context) *Medicalfile {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mc *MedicalfileCreate) sqlSave(ctx context.Context) (*Medicalfile, error) {
	m, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	m.ID = int(id)
	return m, nil
}

func (mc *MedicalfileCreate) createSpec() (*Medicalfile, *sqlgraph.CreateSpec) {
	var (
		m     = &Medicalfile{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: medicalfile.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: medicalfile.FieldID,
			},
		}
	)
	if value, ok := mc.mutation.Detail(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: medicalfile.FieldDetail,
		})
		m.Detail = value
	}
	if value, ok := mc.mutation.AddedTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: medicalfile.FieldAddedTime,
		})
		m.AddedTime = value
	}
	if nodes := mc.mutation.DentistIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   medicalfile.DentistTable,
			Columns: []string{medicalfile.DentistColumn},
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
	if nodes := mc.mutation.PatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   medicalfile.PatientTable,
			Columns: []string{medicalfile.PatientColumn},
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
	if nodes := mc.mutation.NurseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   medicalfile.NurseTable,
			Columns: []string{medicalfile.NurseColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.MedicalcareIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   medicalfile.MedicalcareTable,
			Columns: []string{medicalfile.MedicalcareColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.DentalexpensesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   medicalfile.DentalexpensesTable,
			Columns: []string{medicalfile.DentalexpensesColumn},
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
	return m, _spec
}
