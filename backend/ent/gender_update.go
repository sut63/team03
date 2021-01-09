// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/tanapon395/playlist-video/ent/dentist"
	"github.com/tanapon395/playlist-video/ent/gender"
	"github.com/tanapon395/playlist-video/ent/patient"
	"github.com/tanapon395/playlist-video/ent/predicate"
)

// GenderUpdate is the builder for updating Gender entities.
type GenderUpdate struct {
	config
	hooks      []Hook
	mutation   *GenderMutation
	predicates []predicate.Gender
}

// Where adds a new predicate for the builder.
func (gu *GenderUpdate) Where(ps ...predicate.Gender) *GenderUpdate {
	gu.predicates = append(gu.predicates, ps...)
	return gu
}

// SetName sets the name field.
func (gu *GenderUpdate) SetName(s string) *GenderUpdate {
	gu.mutation.SetName(s)
	return gu
}

// AddPatientIDs adds the patients edge to Patient by ids.
func (gu *GenderUpdate) AddPatientIDs(ids ...int) *GenderUpdate {
	gu.mutation.AddPatientIDs(ids...)
	return gu
}

// AddPatients adds the patients edges to Patient.
func (gu *GenderUpdate) AddPatients(p ...*Patient) *GenderUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return gu.AddPatientIDs(ids...)
}

// AddDentistIDs adds the dentists edge to Dentist by ids.
func (gu *GenderUpdate) AddDentistIDs(ids ...int) *GenderUpdate {
	gu.mutation.AddDentistIDs(ids...)
	return gu
}

// AddDentists adds the dentists edges to Dentist.
func (gu *GenderUpdate) AddDentists(d ...*Dentist) *GenderUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return gu.AddDentistIDs(ids...)
}

// Mutation returns the GenderMutation object of the builder.
func (gu *GenderUpdate) Mutation() *GenderMutation {
	return gu.mutation
}

// RemovePatientIDs removes the patients edge to Patient by ids.
func (gu *GenderUpdate) RemovePatientIDs(ids ...int) *GenderUpdate {
	gu.mutation.RemovePatientIDs(ids...)
	return gu
}

// RemovePatients removes patients edges to Patient.
func (gu *GenderUpdate) RemovePatients(p ...*Patient) *GenderUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return gu.RemovePatientIDs(ids...)
}

// RemoveDentistIDs removes the dentists edge to Dentist by ids.
func (gu *GenderUpdate) RemoveDentistIDs(ids ...int) *GenderUpdate {
	gu.mutation.RemoveDentistIDs(ids...)
	return gu
}

// RemoveDentists removes dentists edges to Dentist.
func (gu *GenderUpdate) RemoveDentists(d ...*Dentist) *GenderUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return gu.RemoveDentistIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (gu *GenderUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := gu.mutation.Name(); ok {
		if err := gender.NameValidator(v); err != nil {
			return 0, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(gu.hooks) == 0 {
		affected, err = gu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GenderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gu.mutation = mutation
			affected, err = gu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gu.hooks) - 1; i >= 0; i-- {
			mut = gu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GenderUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GenderUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GenderUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gu *GenderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   gender.Table,
			Columns: gender.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: gender.FieldID,
			},
		},
	}
	if ps := gu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gender.FieldName,
		})
	}
	if nodes := gu.mutation.RemovedPatientsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.PatientsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := gu.mutation.RemovedDentistsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.DentistsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gender.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// GenderUpdateOne is the builder for updating a single Gender entity.
type GenderUpdateOne struct {
	config
	hooks    []Hook
	mutation *GenderMutation
}

// SetName sets the name field.
func (guo *GenderUpdateOne) SetName(s string) *GenderUpdateOne {
	guo.mutation.SetName(s)
	return guo
}

// AddPatientIDs adds the patients edge to Patient by ids.
func (guo *GenderUpdateOne) AddPatientIDs(ids ...int) *GenderUpdateOne {
	guo.mutation.AddPatientIDs(ids...)
	return guo
}

// AddPatients adds the patients edges to Patient.
func (guo *GenderUpdateOne) AddPatients(p ...*Patient) *GenderUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return guo.AddPatientIDs(ids...)
}

// AddDentistIDs adds the dentists edge to Dentist by ids.
func (guo *GenderUpdateOne) AddDentistIDs(ids ...int) *GenderUpdateOne {
	guo.mutation.AddDentistIDs(ids...)
	return guo
}

// AddDentists adds the dentists edges to Dentist.
func (guo *GenderUpdateOne) AddDentists(d ...*Dentist) *GenderUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return guo.AddDentistIDs(ids...)
}

// Mutation returns the GenderMutation object of the builder.
func (guo *GenderUpdateOne) Mutation() *GenderMutation {
	return guo.mutation
}

// RemovePatientIDs removes the patients edge to Patient by ids.
func (guo *GenderUpdateOne) RemovePatientIDs(ids ...int) *GenderUpdateOne {
	guo.mutation.RemovePatientIDs(ids...)
	return guo
}

// RemovePatients removes patients edges to Patient.
func (guo *GenderUpdateOne) RemovePatients(p ...*Patient) *GenderUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return guo.RemovePatientIDs(ids...)
}

// RemoveDentistIDs removes the dentists edge to Dentist by ids.
func (guo *GenderUpdateOne) RemoveDentistIDs(ids ...int) *GenderUpdateOne {
	guo.mutation.RemoveDentistIDs(ids...)
	return guo
}

// RemoveDentists removes dentists edges to Dentist.
func (guo *GenderUpdateOne) RemoveDentists(d ...*Dentist) *GenderUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return guo.RemoveDentistIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (guo *GenderUpdateOne) Save(ctx context.Context) (*Gender, error) {
	if v, ok := guo.mutation.Name(); ok {
		if err := gender.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}

	var (
		err  error
		node *Gender
	)
	if len(guo.hooks) == 0 {
		node, err = guo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GenderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			guo.mutation = mutation
			node, err = guo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(guo.hooks) - 1; i >= 0; i-- {
			mut = guo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, guo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GenderUpdateOne) SaveX(ctx context.Context) *Gender {
	ge, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return ge
}

// Exec executes the query on the entity.
func (guo *GenderUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GenderUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (guo *GenderUpdateOne) sqlSave(ctx context.Context) (ge *Gender, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   gender.Table,
			Columns: gender.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: gender.FieldID,
			},
		},
	}
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Gender.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := guo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gender.FieldName,
		})
	}
	if nodes := guo.mutation.RemovedPatientsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.PatientsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := guo.mutation.RemovedDentistsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.DentistsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	ge = &Gender{config: guo.config}
	_spec.Assign = ge.assignValues
	_spec.ScanValues = ge.scanValues()
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gender.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return ge, nil
}
