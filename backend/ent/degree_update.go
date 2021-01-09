// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/tanapon395/playlist-video/ent/degree"
	"github.com/tanapon395/playlist-video/ent/dentist"
	"github.com/tanapon395/playlist-video/ent/predicate"
)

// DegreeUpdate is the builder for updating Degree entities.
type DegreeUpdate struct {
	config
	hooks      []Hook
	mutation   *DegreeMutation
	predicates []predicate.Degree
}

// Where adds a new predicate for the builder.
func (du *DegreeUpdate) Where(ps ...predicate.Degree) *DegreeUpdate {
	du.predicates = append(du.predicates, ps...)
	return du
}

// SetDegreeName sets the degree_name field.
func (du *DegreeUpdate) SetDegreeName(s string) *DegreeUpdate {
	du.mutation.SetDegreeName(s)
	return du
}

// AddDentistIDs adds the dentists edge to Dentist by ids.
func (du *DegreeUpdate) AddDentistIDs(ids ...int) *DegreeUpdate {
	du.mutation.AddDentistIDs(ids...)
	return du
}

// AddDentists adds the dentists edges to Dentist.
func (du *DegreeUpdate) AddDentists(d ...*Dentist) *DegreeUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.AddDentistIDs(ids...)
}

// Mutation returns the DegreeMutation object of the builder.
func (du *DegreeUpdate) Mutation() *DegreeMutation {
	return du.mutation
}

// RemoveDentistIDs removes the dentists edge to Dentist by ids.
func (du *DegreeUpdate) RemoveDentistIDs(ids ...int) *DegreeUpdate {
	du.mutation.RemoveDentistIDs(ids...)
	return du
}

// RemoveDentists removes dentists edges to Dentist.
func (du *DegreeUpdate) RemoveDentists(d ...*Dentist) *DegreeUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.RemoveDentistIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (du *DegreeUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := du.mutation.DegreeName(); ok {
		if err := degree.DegreeNameValidator(v); err != nil {
			return 0, &ValidationError{Name: "degree_name", err: fmt.Errorf("ent: validator failed for field \"degree_name\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DegreeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DegreeUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DegreeUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DegreeUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DegreeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   degree.Table,
			Columns: degree.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: degree.FieldID,
			},
		},
	}
	if ps := du.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.DegreeName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: degree.FieldDegreeName,
		})
	}
	if nodes := du.mutation.RemovedDentistsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.DentistsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{degree.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DegreeUpdateOne is the builder for updating a single Degree entity.
type DegreeUpdateOne struct {
	config
	hooks    []Hook
	mutation *DegreeMutation
}

// SetDegreeName sets the degree_name field.
func (duo *DegreeUpdateOne) SetDegreeName(s string) *DegreeUpdateOne {
	duo.mutation.SetDegreeName(s)
	return duo
}

// AddDentistIDs adds the dentists edge to Dentist by ids.
func (duo *DegreeUpdateOne) AddDentistIDs(ids ...int) *DegreeUpdateOne {
	duo.mutation.AddDentistIDs(ids...)
	return duo
}

// AddDentists adds the dentists edges to Dentist.
func (duo *DegreeUpdateOne) AddDentists(d ...*Dentist) *DegreeUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.AddDentistIDs(ids...)
}

// Mutation returns the DegreeMutation object of the builder.
func (duo *DegreeUpdateOne) Mutation() *DegreeMutation {
	return duo.mutation
}

// RemoveDentistIDs removes the dentists edge to Dentist by ids.
func (duo *DegreeUpdateOne) RemoveDentistIDs(ids ...int) *DegreeUpdateOne {
	duo.mutation.RemoveDentistIDs(ids...)
	return duo
}

// RemoveDentists removes dentists edges to Dentist.
func (duo *DegreeUpdateOne) RemoveDentists(d ...*Dentist) *DegreeUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.RemoveDentistIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (duo *DegreeUpdateOne) Save(ctx context.Context) (*Degree, error) {
	if v, ok := duo.mutation.DegreeName(); ok {
		if err := degree.DegreeNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "degree_name", err: fmt.Errorf("ent: validator failed for field \"degree_name\": %w", err)}
		}
	}

	var (
		err  error
		node *Degree
	)
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DegreeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DegreeUpdateOne) SaveX(ctx context.Context) *Degree {
	d, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return d
}

// Exec executes the query on the entity.
func (duo *DegreeUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DegreeUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DegreeUpdateOne) sqlSave(ctx context.Context) (d *Degree, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   degree.Table,
			Columns: degree.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: degree.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Degree.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := duo.mutation.DegreeName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: degree.FieldDegreeName,
		})
	}
	if nodes := duo.mutation.RemovedDentistsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.DentistsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	d = &Degree{config: duo.config}
	_spec.Assign = d.assignValues
	_spec.ScanValues = d.scanValues()
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{degree.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return d, nil
}