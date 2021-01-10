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
	"github.com/team03/app/ent/dentist"
	"github.com/team03/app/ent/patient"
	"github.com/team03/app/ent/predicate"
	"github.com/team03/app/ent/room"
)

// AppointmentUpdate is the builder for updating Appointment entities.
type AppointmentUpdate struct {
	config
	hooks      []Hook
	mutation   *AppointmentMutation
	predicates []predicate.Appointment
}

// Where adds a new predicate for the builder.
func (au *AppointmentUpdate) Where(ps ...predicate.Appointment) *AppointmentUpdate {
	au.predicates = append(au.predicates, ps...)
	return au
}

// SetDetail sets the detail field.
func (au *AppointmentUpdate) SetDetail(s string) *AppointmentUpdate {
	au.mutation.SetDetail(s)
	return au
}

// SetDatetime sets the datetime field.
func (au *AppointmentUpdate) SetDatetime(t time.Time) *AppointmentUpdate {
	au.mutation.SetDatetime(t)
	return au
}

// SetPatientID sets the patient edge to Patient by id.
func (au *AppointmentUpdate) SetPatientID(id int) *AppointmentUpdate {
	au.mutation.SetPatientID(id)
	return au
}

// SetNillablePatientID sets the patient edge to Patient by id if the given value is not nil.
func (au *AppointmentUpdate) SetNillablePatientID(id *int) *AppointmentUpdate {
	if id != nil {
		au = au.SetPatientID(*id)
	}
	return au
}

// SetPatient sets the patient edge to Patient.
func (au *AppointmentUpdate) SetPatient(p *Patient) *AppointmentUpdate {
	return au.SetPatientID(p.ID)
}

// SetRoomID sets the room edge to Room by id.
func (au *AppointmentUpdate) SetRoomID(id int) *AppointmentUpdate {
	au.mutation.SetRoomID(id)
	return au
}

// SetNillableRoomID sets the room edge to Room by id if the given value is not nil.
func (au *AppointmentUpdate) SetNillableRoomID(id *int) *AppointmentUpdate {
	if id != nil {
		au = au.SetRoomID(*id)
	}
	return au
}

// SetRoom sets the room edge to Room.
func (au *AppointmentUpdate) SetRoom(r *Room) *AppointmentUpdate {
	return au.SetRoomID(r.ID)
}

// SetDentistID sets the dentist edge to Dentist by id.
func (au *AppointmentUpdate) SetDentistID(id int) *AppointmentUpdate {
	au.mutation.SetDentistID(id)
	return au
}

// SetNillableDentistID sets the dentist edge to Dentist by id if the given value is not nil.
func (au *AppointmentUpdate) SetNillableDentistID(id *int) *AppointmentUpdate {
	if id != nil {
		au = au.SetDentistID(*id)
	}
	return au
}

// SetDentist sets the dentist edge to Dentist.
func (au *AppointmentUpdate) SetDentist(d *Dentist) *AppointmentUpdate {
	return au.SetDentistID(d.ID)
}

// Mutation returns the AppointmentMutation object of the builder.
func (au *AppointmentUpdate) Mutation() *AppointmentMutation {
	return au.mutation
}

// ClearPatient clears the patient edge to Patient.
func (au *AppointmentUpdate) ClearPatient() *AppointmentUpdate {
	au.mutation.ClearPatient()
	return au
}

// ClearRoom clears the room edge to Room.
func (au *AppointmentUpdate) ClearRoom() *AppointmentUpdate {
	au.mutation.ClearRoom()
	return au
}

// ClearDentist clears the dentist edge to Dentist.
func (au *AppointmentUpdate) ClearDentist() *AppointmentUpdate {
	au.mutation.ClearDentist()
	return au
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (au *AppointmentUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := au.mutation.Detail(); ok {
		if err := appointment.DetailValidator(v); err != nil {
			return 0, &ValidationError{Name: "detail", err: fmt.Errorf("ent: validator failed for field \"detail\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppointmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AppointmentUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AppointmentUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AppointmentUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AppointmentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appointment.Table,
			Columns: appointment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: appointment.FieldID,
			},
		},
	}
	if ps := au.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Detail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appointment.FieldDetail,
		})
	}
	if value, ok := au.mutation.Datetime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: appointment.FieldDatetime,
		})
	}
	if au.mutation.PatientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appointment.PatientTable,
			Columns: []string{appointment.PatientColumn},
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
	if nodes := au.mutation.PatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appointment.PatientTable,
			Columns: []string{appointment.PatientColumn},
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
	if au.mutation.RoomCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appointment.RoomTable,
			Columns: []string{appointment.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appointment.RoomTable,
			Columns: []string{appointment.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.DentistCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appointment.DentistTable,
			Columns: []string{appointment.DentistColumn},
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
	if nodes := au.mutation.DentistIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appointment.DentistTable,
			Columns: []string{appointment.DentistColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appointment.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AppointmentUpdateOne is the builder for updating a single Appointment entity.
type AppointmentUpdateOne struct {
	config
	hooks    []Hook
	mutation *AppointmentMutation
}

// SetDetail sets the detail field.
func (auo *AppointmentUpdateOne) SetDetail(s string) *AppointmentUpdateOne {
	auo.mutation.SetDetail(s)
	return auo
}

// SetDatetime sets the datetime field.
func (auo *AppointmentUpdateOne) SetDatetime(t time.Time) *AppointmentUpdateOne {
	auo.mutation.SetDatetime(t)
	return auo
}

// SetPatientID sets the patient edge to Patient by id.
func (auo *AppointmentUpdateOne) SetPatientID(id int) *AppointmentUpdateOne {
	auo.mutation.SetPatientID(id)
	return auo
}

// SetNillablePatientID sets the patient edge to Patient by id if the given value is not nil.
func (auo *AppointmentUpdateOne) SetNillablePatientID(id *int) *AppointmentUpdateOne {
	if id != nil {
		auo = auo.SetPatientID(*id)
	}
	return auo
}

// SetPatient sets the patient edge to Patient.
func (auo *AppointmentUpdateOne) SetPatient(p *Patient) *AppointmentUpdateOne {
	return auo.SetPatientID(p.ID)
}

// SetRoomID sets the room edge to Room by id.
func (auo *AppointmentUpdateOne) SetRoomID(id int) *AppointmentUpdateOne {
	auo.mutation.SetRoomID(id)
	return auo
}

// SetNillableRoomID sets the room edge to Room by id if the given value is not nil.
func (auo *AppointmentUpdateOne) SetNillableRoomID(id *int) *AppointmentUpdateOne {
	if id != nil {
		auo = auo.SetRoomID(*id)
	}
	return auo
}

// SetRoom sets the room edge to Room.
func (auo *AppointmentUpdateOne) SetRoom(r *Room) *AppointmentUpdateOne {
	return auo.SetRoomID(r.ID)
}

// SetDentistID sets the dentist edge to Dentist by id.
func (auo *AppointmentUpdateOne) SetDentistID(id int) *AppointmentUpdateOne {
	auo.mutation.SetDentistID(id)
	return auo
}

// SetNillableDentistID sets the dentist edge to Dentist by id if the given value is not nil.
func (auo *AppointmentUpdateOne) SetNillableDentistID(id *int) *AppointmentUpdateOne {
	if id != nil {
		auo = auo.SetDentistID(*id)
	}
	return auo
}

// SetDentist sets the dentist edge to Dentist.
func (auo *AppointmentUpdateOne) SetDentist(d *Dentist) *AppointmentUpdateOne {
	return auo.SetDentistID(d.ID)
}

// Mutation returns the AppointmentMutation object of the builder.
func (auo *AppointmentUpdateOne) Mutation() *AppointmentMutation {
	return auo.mutation
}

// ClearPatient clears the patient edge to Patient.
func (auo *AppointmentUpdateOne) ClearPatient() *AppointmentUpdateOne {
	auo.mutation.ClearPatient()
	return auo
}

// ClearRoom clears the room edge to Room.
func (auo *AppointmentUpdateOne) ClearRoom() *AppointmentUpdateOne {
	auo.mutation.ClearRoom()
	return auo
}

// ClearDentist clears the dentist edge to Dentist.
func (auo *AppointmentUpdateOne) ClearDentist() *AppointmentUpdateOne {
	auo.mutation.ClearDentist()
	return auo
}

// Save executes the query and returns the updated entity.
func (auo *AppointmentUpdateOne) Save(ctx context.Context) (*Appointment, error) {
	if v, ok := auo.mutation.Detail(); ok {
		if err := appointment.DetailValidator(v); err != nil {
			return nil, &ValidationError{Name: "detail", err: fmt.Errorf("ent: validator failed for field \"detail\": %w", err)}
		}
	}

	var (
		err  error
		node *Appointment
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppointmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AppointmentUpdateOne) SaveX(ctx context.Context) *Appointment {
	a, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return a
}

// Exec executes the query on the entity.
func (auo *AppointmentUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AppointmentUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AppointmentUpdateOne) sqlSave(ctx context.Context) (a *Appointment, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appointment.Table,
			Columns: appointment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: appointment.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Appointment.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := auo.mutation.Detail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appointment.FieldDetail,
		})
	}
	if value, ok := auo.mutation.Datetime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: appointment.FieldDatetime,
		})
	}
	if auo.mutation.PatientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appointment.PatientTable,
			Columns: []string{appointment.PatientColumn},
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
	if nodes := auo.mutation.PatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appointment.PatientTable,
			Columns: []string{appointment.PatientColumn},
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
	if auo.mutation.RoomCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appointment.RoomTable,
			Columns: []string{appointment.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appointment.RoomTable,
			Columns: []string{appointment.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.DentistCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appointment.DentistTable,
			Columns: []string{appointment.DentistColumn},
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
	if nodes := auo.mutation.DentistIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appointment.DentistTable,
			Columns: []string{appointment.DentistColumn},
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
	a = &Appointment{config: auo.config}
	_spec.Assign = a.assignValues
	_spec.ScanValues = a.scanValues()
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appointment.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return a, nil
}
