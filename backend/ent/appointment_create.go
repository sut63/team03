// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team03/app/ent/appointment"
	"github.com/team03/app/ent/dentist"
	"github.com/team03/app/ent/patient"
	"github.com/team03/app/ent/room"
)

// AppointmentCreate is the builder for creating a Appointment entity.
type AppointmentCreate struct {
	config
	mutation *AppointmentMutation
	hooks    []Hook
}

// SetDetail sets the detail field.
func (ac *AppointmentCreate) SetDetail(s string) *AppointmentCreate {
	ac.mutation.SetDetail(s)
	return ac
}

// SetDatetime sets the datetime field.
func (ac *AppointmentCreate) SetDatetime(t time.Time) *AppointmentCreate {
	ac.mutation.SetDatetime(t)
	return ac
}

// SetPatientID sets the patient edge to Patient by id.
func (ac *AppointmentCreate) SetPatientID(id int) *AppointmentCreate {
	ac.mutation.SetPatientID(id)
	return ac
}

// SetNillablePatientID sets the patient edge to Patient by id if the given value is not nil.
func (ac *AppointmentCreate) SetNillablePatientID(id *int) *AppointmentCreate {
	if id != nil {
		ac = ac.SetPatientID(*id)
	}
	return ac
}

// SetPatient sets the patient edge to Patient.
func (ac *AppointmentCreate) SetPatient(p *Patient) *AppointmentCreate {
	return ac.SetPatientID(p.ID)
}

// SetRoomID sets the room edge to Room by id.
func (ac *AppointmentCreate) SetRoomID(id int) *AppointmentCreate {
	ac.mutation.SetRoomID(id)
	return ac
}

// SetNillableRoomID sets the room edge to Room by id if the given value is not nil.
func (ac *AppointmentCreate) SetNillableRoomID(id *int) *AppointmentCreate {
	if id != nil {
		ac = ac.SetRoomID(*id)
	}
	return ac
}

// SetRoom sets the room edge to Room.
func (ac *AppointmentCreate) SetRoom(r *Room) *AppointmentCreate {
	return ac.SetRoomID(r.ID)
}

// SetDentistID sets the dentist edge to Dentist by id.
func (ac *AppointmentCreate) SetDentistID(id int) *AppointmentCreate {
	ac.mutation.SetDentistID(id)
	return ac
}

// SetNillableDentistID sets the dentist edge to Dentist by id if the given value is not nil.
func (ac *AppointmentCreate) SetNillableDentistID(id *int) *AppointmentCreate {
	if id != nil {
		ac = ac.SetDentistID(*id)
	}
	return ac
}

// SetDentist sets the dentist edge to Dentist.
func (ac *AppointmentCreate) SetDentist(d *Dentist) *AppointmentCreate {
	return ac.SetDentistID(d.ID)
}

// Mutation returns the AppointmentMutation object of the builder.
func (ac *AppointmentCreate) Mutation() *AppointmentMutation {
	return ac.mutation
}

// Save creates the Appointment in the database.
func (ac *AppointmentCreate) Save(ctx context.Context) (*Appointment, error) {
	if _, ok := ac.mutation.Detail(); !ok {
		return nil, &ValidationError{Name: "detail", err: errors.New("ent: missing required field \"detail\"")}
	}
	if v, ok := ac.mutation.Detail(); ok {
		if err := appointment.DetailValidator(v); err != nil {
			return nil, &ValidationError{Name: "detail", err: fmt.Errorf("ent: validator failed for field \"detail\": %w", err)}
		}
	}
	if _, ok := ac.mutation.Datetime(); !ok {
		return nil, &ValidationError{Name: "datetime", err: errors.New("ent: missing required field \"datetime\"")}
	}
	var (
		err  error
		node *Appointment
	)
	if len(ac.hooks) == 0 {
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppointmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ac.mutation = mutation
			node, err = ac.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AppointmentCreate) SaveX(ctx context.Context) *Appointment {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ac *AppointmentCreate) sqlSave(ctx context.Context) (*Appointment, error) {
	a, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	a.ID = int(id)
	return a, nil
}

func (ac *AppointmentCreate) createSpec() (*Appointment, *sqlgraph.CreateSpec) {
	var (
		a     = &Appointment{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: appointment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: appointment.FieldID,
			},
		}
	)
	if value, ok := ac.mutation.Detail(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appointment.FieldDetail,
		})
		a.Detail = value
	}
	if value, ok := ac.mutation.Datetime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: appointment.FieldDatetime,
		})
		a.Datetime = value
	}
	if nodes := ac.mutation.PatientIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.RoomIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.DentistIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return a, _spec
}
