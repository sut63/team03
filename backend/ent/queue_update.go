// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team03/app/ent/dentist"
	"github.com/team03/app/ent/nurse"
	"github.com/team03/app/ent/patient"
	"github.com/team03/app/ent/predicate"
	"github.com/team03/app/ent/queue"
)

// QueueUpdate is the builder for updating Queue entities.
type QueueUpdate struct {
	config
	hooks      []Hook
	mutation   *QueueMutation
	predicates []predicate.Queue
}

// Where adds a new predicate for the builder.
func (qu *QueueUpdate) Where(ps ...predicate.Queue) *QueueUpdate {
	qu.predicates = append(qu.predicates, ps...)
	return qu
}

// SetQueueID sets the QueueID field.
func (qu *QueueUpdate) SetQueueID(s string) *QueueUpdate {
	qu.mutation.SetQueueID(s)
	return qu
}

// SetPhone sets the Phone field.
func (qu *QueueUpdate) SetPhone(s string) *QueueUpdate {
	qu.mutation.SetPhone(s)
	return qu
}

// SetDental sets the Dental field.
func (qu *QueueUpdate) SetDental(s string) *QueueUpdate {
	qu.mutation.SetDental(s)
	return qu
}

// SetQueueTime sets the QueueTime field.
func (qu *QueueUpdate) SetQueueTime(t time.Time) *QueueUpdate {
	qu.mutation.SetQueueTime(t)
	return qu
}

// SetDentistID sets the dentist edge to Dentist by id.
func (qu *QueueUpdate) SetDentistID(id int) *QueueUpdate {
	qu.mutation.SetDentistID(id)
	return qu
}

// SetNillableDentistID sets the dentist edge to Dentist by id if the given value is not nil.
func (qu *QueueUpdate) SetNillableDentistID(id *int) *QueueUpdate {
	if id != nil {
		qu = qu.SetDentistID(*id)
	}
	return qu
}

// SetDentist sets the dentist edge to Dentist.
func (qu *QueueUpdate) SetDentist(d *Dentist) *QueueUpdate {
	return qu.SetDentistID(d.ID)
}

// SetNurseID sets the nurse edge to Nurse by id.
func (qu *QueueUpdate) SetNurseID(id int) *QueueUpdate {
	qu.mutation.SetNurseID(id)
	return qu
}

// SetNillableNurseID sets the nurse edge to Nurse by id if the given value is not nil.
func (qu *QueueUpdate) SetNillableNurseID(id *int) *QueueUpdate {
	if id != nil {
		qu = qu.SetNurseID(*id)
	}
	return qu
}

// SetNurse sets the nurse edge to Nurse.
func (qu *QueueUpdate) SetNurse(n *Nurse) *QueueUpdate {
	return qu.SetNurseID(n.ID)
}

// SetPatientID sets the patient edge to Patient by id.
func (qu *QueueUpdate) SetPatientID(id int) *QueueUpdate {
	qu.mutation.SetPatientID(id)
	return qu
}

// SetNillablePatientID sets the patient edge to Patient by id if the given value is not nil.
func (qu *QueueUpdate) SetNillablePatientID(id *int) *QueueUpdate {
	if id != nil {
		qu = qu.SetPatientID(*id)
	}
	return qu
}

// SetPatient sets the patient edge to Patient.
func (qu *QueueUpdate) SetPatient(p *Patient) *QueueUpdate {
	return qu.SetPatientID(p.ID)
}

// Mutation returns the QueueMutation object of the builder.
func (qu *QueueUpdate) Mutation() *QueueMutation {
	return qu.mutation
}

// ClearDentist clears the dentist edge to Dentist.
func (qu *QueueUpdate) ClearDentist() *QueueUpdate {
	qu.mutation.ClearDentist()
	return qu
}

// ClearNurse clears the nurse edge to Nurse.
func (qu *QueueUpdate) ClearNurse() *QueueUpdate {
	qu.mutation.ClearNurse()
	return qu
}

// ClearPatient clears the patient edge to Patient.
func (qu *QueueUpdate) ClearPatient() *QueueUpdate {
	qu.mutation.ClearPatient()
	return qu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (qu *QueueUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := qu.mutation.QueueID(); ok {
		if err := queue.QueueIDValidator(v); err != nil {
			return 0, &ValidationError{Name: "QueueID", err: fmt.Errorf("ent: validator failed for field \"QueueID\": %w", err)}
		}
	}
	if v, ok := qu.mutation.Phone(); ok {
		if err := queue.PhoneValidator(v); err != nil {
			return 0, &ValidationError{Name: "Phone", err: fmt.Errorf("ent: validator failed for field \"Phone\": %w", err)}
		}
	}
	if v, ok := qu.mutation.Dental(); ok {
		if err := queue.DentalValidator(v); err != nil {
			return 0, &ValidationError{Name: "Dental", err: fmt.Errorf("ent: validator failed for field \"Dental\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(qu.hooks) == 0 {
		affected, err = qu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QueueMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			qu.mutation = mutation
			affected, err = qu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(qu.hooks) - 1; i >= 0; i-- {
			mut = qu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, qu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (qu *QueueUpdate) SaveX(ctx context.Context) int {
	affected, err := qu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (qu *QueueUpdate) Exec(ctx context.Context) error {
	_, err := qu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qu *QueueUpdate) ExecX(ctx context.Context) {
	if err := qu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (qu *QueueUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   queue.Table,
			Columns: queue.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: queue.FieldID,
			},
		},
	}
	if ps := qu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := qu.mutation.QueueID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: queue.FieldQueueID,
		})
	}
	if value, ok := qu.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: queue.FieldPhone,
		})
	}
	if value, ok := qu.mutation.Dental(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: queue.FieldDental,
		})
	}
	if value, ok := qu.mutation.QueueTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: queue.FieldQueueTime,
		})
	}
	if qu.mutation.DentistCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   queue.DentistTable,
			Columns: []string{queue.DentistColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dentist.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qu.mutation.DentistIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   queue.DentistTable,
			Columns: []string{queue.DentistColumn},
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
	if qu.mutation.NurseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   queue.NurseTable,
			Columns: []string{queue.NurseColumn},
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
	if nodes := qu.mutation.NurseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   queue.NurseTable,
			Columns: []string{queue.NurseColumn},
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
	if qu.mutation.PatientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   queue.PatientTable,
			Columns: []string{queue.PatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qu.mutation.PatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   queue.PatientTable,
			Columns: []string{queue.PatientColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, qu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{queue.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// QueueUpdateOne is the builder for updating a single Queue entity.
type QueueUpdateOne struct {
	config
	hooks    []Hook
	mutation *QueueMutation
}

// SetQueueID sets the QueueID field.
func (quo *QueueUpdateOne) SetQueueID(s string) *QueueUpdateOne {
	quo.mutation.SetQueueID(s)
	return quo
}

// SetPhone sets the Phone field.
func (quo *QueueUpdateOne) SetPhone(s string) *QueueUpdateOne {
	quo.mutation.SetPhone(s)
	return quo
}

// SetDental sets the Dental field.
func (quo *QueueUpdateOne) SetDental(s string) *QueueUpdateOne {
	quo.mutation.SetDental(s)
	return quo
}

// SetQueueTime sets the QueueTime field.
func (quo *QueueUpdateOne) SetQueueTime(t time.Time) *QueueUpdateOne {
	quo.mutation.SetQueueTime(t)
	return quo
}

// SetDentistID sets the dentist edge to Dentist by id.
func (quo *QueueUpdateOne) SetDentistID(id int) *QueueUpdateOne {
	quo.mutation.SetDentistID(id)
	return quo
}

// SetNillableDentistID sets the dentist edge to Dentist by id if the given value is not nil.
func (quo *QueueUpdateOne) SetNillableDentistID(id *int) *QueueUpdateOne {
	if id != nil {
		quo = quo.SetDentistID(*id)
	}
	return quo
}

// SetDentist sets the dentist edge to Dentist.
func (quo *QueueUpdateOne) SetDentist(d *Dentist) *QueueUpdateOne {
	return quo.SetDentistID(d.ID)
}

// SetNurseID sets the nurse edge to Nurse by id.
func (quo *QueueUpdateOne) SetNurseID(id int) *QueueUpdateOne {
	quo.mutation.SetNurseID(id)
	return quo
}

// SetNillableNurseID sets the nurse edge to Nurse by id if the given value is not nil.
func (quo *QueueUpdateOne) SetNillableNurseID(id *int) *QueueUpdateOne {
	if id != nil {
		quo = quo.SetNurseID(*id)
	}
	return quo
}

// SetNurse sets the nurse edge to Nurse.
func (quo *QueueUpdateOne) SetNurse(n *Nurse) *QueueUpdateOne {
	return quo.SetNurseID(n.ID)
}

// SetPatientID sets the patient edge to Patient by id.
func (quo *QueueUpdateOne) SetPatientID(id int) *QueueUpdateOne {
	quo.mutation.SetPatientID(id)
	return quo
}

// SetNillablePatientID sets the patient edge to Patient by id if the given value is not nil.
func (quo *QueueUpdateOne) SetNillablePatientID(id *int) *QueueUpdateOne {
	if id != nil {
		quo = quo.SetPatientID(*id)
	}
	return quo
}

// SetPatient sets the patient edge to Patient.
func (quo *QueueUpdateOne) SetPatient(p *Patient) *QueueUpdateOne {
	return quo.SetPatientID(p.ID)
}

// Mutation returns the QueueMutation object of the builder.
func (quo *QueueUpdateOne) Mutation() *QueueMutation {
	return quo.mutation
}

// ClearDentist clears the dentist edge to Dentist.
func (quo *QueueUpdateOne) ClearDentist() *QueueUpdateOne {
	quo.mutation.ClearDentist()
	return quo
}

// ClearNurse clears the nurse edge to Nurse.
func (quo *QueueUpdateOne) ClearNurse() *QueueUpdateOne {
	quo.mutation.ClearNurse()
	return quo
}

// ClearPatient clears the patient edge to Patient.
func (quo *QueueUpdateOne) ClearPatient() *QueueUpdateOne {
	quo.mutation.ClearPatient()
	return quo
}

// Save executes the query and returns the updated entity.
func (quo *QueueUpdateOne) Save(ctx context.Context) (*Queue, error) {
	if v, ok := quo.mutation.QueueID(); ok {
		if err := queue.QueueIDValidator(v); err != nil {
			return nil, &ValidationError{Name: "QueueID", err: fmt.Errorf("ent: validator failed for field \"QueueID\": %w", err)}
		}
	}
	if v, ok := quo.mutation.Phone(); ok {
		if err := queue.PhoneValidator(v); err != nil {
			return nil, &ValidationError{Name: "Phone", err: fmt.Errorf("ent: validator failed for field \"Phone\": %w", err)}
		}
	}
	if v, ok := quo.mutation.Dental(); ok {
		if err := queue.DentalValidator(v); err != nil {
			return nil, &ValidationError{Name: "Dental", err: fmt.Errorf("ent: validator failed for field \"Dental\": %w", err)}
		}
	}

	var (
		err  error
		node *Queue
	)
	if len(quo.hooks) == 0 {
		node, err = quo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QueueMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			quo.mutation = mutation
			node, err = quo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(quo.hooks) - 1; i >= 0; i-- {
			mut = quo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, quo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (quo *QueueUpdateOne) SaveX(ctx context.Context) *Queue {
	q, err := quo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return q
}

// Exec executes the query on the entity.
func (quo *QueueUpdateOne) Exec(ctx context.Context) error {
	_, err := quo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (quo *QueueUpdateOne) ExecX(ctx context.Context) {
	if err := quo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (quo *QueueUpdateOne) sqlSave(ctx context.Context) (q *Queue, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   queue.Table,
			Columns: queue.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: queue.FieldID,
			},
		},
	}
	id, ok := quo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Queue.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := quo.mutation.QueueID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: queue.FieldQueueID,
		})
	}
	if value, ok := quo.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: queue.FieldPhone,
		})
	}
	if value, ok := quo.mutation.Dental(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: queue.FieldDental,
		})
	}
	if value, ok := quo.mutation.QueueTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: queue.FieldQueueTime,
		})
	}
	if quo.mutation.DentistCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   queue.DentistTable,
			Columns: []string{queue.DentistColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dentist.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := quo.mutation.DentistIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   queue.DentistTable,
			Columns: []string{queue.DentistColumn},
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
	if quo.mutation.NurseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   queue.NurseTable,
			Columns: []string{queue.NurseColumn},
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
	if nodes := quo.mutation.NurseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   queue.NurseTable,
			Columns: []string{queue.NurseColumn},
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
	if quo.mutation.PatientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   queue.PatientTable,
			Columns: []string{queue.PatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := quo.mutation.PatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   queue.PatientTable,
			Columns: []string{queue.PatientColumn},
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
	q = &Queue{config: quo.config}
	_spec.Assign = q.assignValues
	_spec.ScanValues = q.scanValues()
	if err = sqlgraph.UpdateNode(ctx, quo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{queue.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return q, nil
}
