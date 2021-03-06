// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team03/app/ent/degree"
	"github.com/team03/app/ent/dentist"
)

// DegreeCreate is the builder for creating a Degree entity.
type DegreeCreate struct {
	config
	mutation *DegreeMutation
	hooks    []Hook
}

// SetName sets the name field.
func (dc *DegreeCreate) SetName(s string) *DegreeCreate {
	dc.mutation.SetName(s)
	return dc
}

// AddDentistIDs adds the dentists edge to Dentist by ids.
func (dc *DegreeCreate) AddDentistIDs(ids ...int) *DegreeCreate {
	dc.mutation.AddDentistIDs(ids...)
	return dc
}

// AddDentists adds the dentists edges to Dentist.
func (dc *DegreeCreate) AddDentists(d ...*Dentist) *DegreeCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dc.AddDentistIDs(ids...)
}

// Mutation returns the DegreeMutation object of the builder.
func (dc *DegreeCreate) Mutation() *DegreeMutation {
	return dc.mutation
}

// Save creates the Degree in the database.
func (dc *DegreeCreate) Save(ctx context.Context) (*Degree, error) {
	if _, ok := dc.mutation.Name(); !ok {
		return nil, &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := dc.mutation.Name(); ok {
		if err := degree.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	var (
		err  error
		node *Degree
	)
	if len(dc.hooks) == 0 {
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DegreeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dc.mutation = mutation
			node, err = dc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DegreeCreate) SaveX(ctx context.Context) *Degree {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dc *DegreeCreate) sqlSave(ctx context.Context) (*Degree, error) {
	d, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	d.ID = int(id)
	return d, nil
}

func (dc *DegreeCreate) createSpec() (*Degree, *sqlgraph.CreateSpec) {
	var (
		d     = &Degree{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: degree.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: degree.FieldID,
			},
		}
	)
	if value, ok := dc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: degree.FieldName,
		})
		d.Name = value
	}
	if nodes := dc.mutation.DentistsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   degree.DentistsTable,
			Columns: []string{degree.DentistsColumn},
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
	return d, _spec
}
