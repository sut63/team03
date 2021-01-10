// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team03/app/ent/dentist"
	"github.com/team03/app/ent/gender"
	"github.com/team03/app/ent/patient"
)

// GenderCreate is the builder for creating a Gender entity.
type GenderCreate struct {
	config
	mutation *GenderMutation
	hooks    []Hook
}

// SetName sets the name field.
func (gc *GenderCreate) SetName(s string) *GenderCreate {
	gc.mutation.SetName(s)
	return gc
}

// AddPatientIDs adds the patients edge to Patient by ids.
func (gc *GenderCreate) AddPatientIDs(ids ...int) *GenderCreate {
	gc.mutation.AddPatientIDs(ids...)
	return gc
}

// AddPatients adds the patients edges to Patient.
func (gc *GenderCreate) AddPatients(p ...*Patient) *GenderCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return gc.AddPatientIDs(ids...)
}

// AddDentistIDs adds the dentists edge to Dentist by ids.
func (gc *GenderCreate) AddDentistIDs(ids ...int) *GenderCreate {
	gc.mutation.AddDentistIDs(ids...)
	return gc
}

// AddDentists adds the dentists edges to Dentist.
func (gc *GenderCreate) AddDentists(d ...*Dentist) *GenderCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return gc.AddDentistIDs(ids...)
}

// Mutation returns the GenderMutation object of the builder.
func (gc *GenderCreate) Mutation() *GenderMutation {
	return gc.mutation
}

// Save creates the Gender in the database.
func (gc *GenderCreate) Save(ctx context.Context) (*Gender, error) {
	if _, ok := gc.mutation.Name(); !ok {
		return nil, &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := gc.mutation.Name(); ok {
		if err := gender.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	var (
		err  error
		node *Gender
	)
	if len(gc.hooks) == 0 {
		node, err = gc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GenderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gc.mutation = mutation
			node, err = gc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gc.hooks) - 1; i >= 0; i-- {
			mut = gc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GenderCreate) SaveX(ctx context.Context) *Gender {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gc *GenderCreate) sqlSave(ctx context.Context) (*Gender, error) {
	ge, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	ge.ID = int(id)
	return ge, nil
}

func (gc *GenderCreate) createSpec() (*Gender, *sqlgraph.CreateSpec) {
	var (
		ge    = &Gender{config: gc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: gender.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: gender.FieldID,
			},
		}
	)
	if value, ok := gc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gender.FieldName,
		})
		ge.Name = value
	}
	if nodes := gc.mutation.PatientsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gender.PatientsTable,
			Columns: []string{gender.PatientsColumn},
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
	if nodes := gc.mutation.DentistsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gender.DentistsTable,
			Columns: []string{gender.DentistsColumn},
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
	return ge, _spec
}
