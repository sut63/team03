// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/tanapon395/playlist-video/ent/dentist"
	"github.com/tanapon395/playlist-video/ent/nurse"
	"github.com/tanapon395/playlist-video/ent/patient"
	"github.com/tanapon395/playlist-video/ent/queue"
)

// QueueCreate is the builder for creating a Queue entity.
type QueueCreate struct {
	config
	mutation *QueueMutation
	hooks    []Hook
}

// SetDental sets the dental field.
func (qc *QueueCreate) SetDental(s string) *QueueCreate {
	qc.mutation.SetDental(s)
	return qc
}

// SetQueueTime sets the queue_time field.
func (qc *QueueCreate) SetQueueTime(t time.Time) *QueueCreate {
	qc.mutation.SetQueueTime(t)
	return qc
}

// SetDentistID sets the dentist edge to Dentist by id.
func (qc *QueueCreate) SetDentistID(id int) *QueueCreate {
	qc.mutation.SetDentistID(id)
	return qc
}

// SetNillableDentistID sets the dentist edge to Dentist by id if the given value is not nil.
func (qc *QueueCreate) SetNillableDentistID(id *int) *QueueCreate {
	if id != nil {
		qc = qc.SetDentistID(*id)
	}
	return qc
}

// SetDentist sets the dentist edge to Dentist.
func (qc *QueueCreate) SetDentist(d *Dentist) *QueueCreate {
	return qc.SetDentistID(d.ID)
}

// SetNurseID sets the nurse edge to Nurse by id.
func (qc *QueueCreate) SetNurseID(id int) *QueueCreate {
	qc.mutation.SetNurseID(id)
	return qc
}

// SetNillableNurseID sets the nurse edge to Nurse by id if the given value is not nil.
func (qc *QueueCreate) SetNillableNurseID(id *int) *QueueCreate {
	if id != nil {
		qc = qc.SetNurseID(*id)
	}
	return qc
}

// SetNurse sets the nurse edge to Nurse.
func (qc *QueueCreate) SetNurse(n *Nurse) *QueueCreate {
	return qc.SetNurseID(n.ID)
}

// SetPatientID sets the patient edge to Patient by id.
func (qc *QueueCreate) SetPatientID(id int) *QueueCreate {
	qc.mutation.SetPatientID(id)
	return qc
}

// SetNillablePatientID sets the patient edge to Patient by id if the given value is not nil.
func (qc *QueueCreate) SetNillablePatientID(id *int) *QueueCreate {
	if id != nil {
		qc = qc.SetPatientID(*id)
	}
	return qc
}

// SetPatient sets the patient edge to Patient.
func (qc *QueueCreate) SetPatient(p *Patient) *QueueCreate {
	return qc.SetPatientID(p.ID)
}

// Mutation returns the QueueMutation object of the builder.
func (qc *QueueCreate) Mutation() *QueueMutation {
	return qc.mutation
}

// Save creates the Queue in the database.
func (qc *QueueCreate) Save(ctx context.Context) (*Queue, error) {
	if _, ok := qc.mutation.Dental(); !ok {
		return nil, &ValidationError{Name: "dental", err: errors.New("ent: missing required field \"dental\"")}
	}
	if _, ok := qc.mutation.QueueTime(); !ok {
		return nil, &ValidationError{Name: "queue_time", err: errors.New("ent: missing required field \"queue_time\"")}
	}
	var (
		err  error
		node *Queue
	)
	if len(qc.hooks) == 0 {
		node, err = qc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QueueMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			qc.mutation = mutation
			node, err = qc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(qc.hooks) - 1; i >= 0; i-- {
			mut = qc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, qc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (qc *QueueCreate) SaveX(ctx context.Context) *Queue {
	v, err := qc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (qc *QueueCreate) sqlSave(ctx context.Context) (*Queue, error) {
	q, _spec := qc.createSpec()
	if err := sqlgraph.CreateNode(ctx, qc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	q.ID = int(id)
	return q, nil
}

func (qc *QueueCreate) createSpec() (*Queue, *sqlgraph.CreateSpec) {
	var (
		q     = &Queue{config: qc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: queue.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: queue.FieldID,
			},
		}
	)
	if value, ok := qc.mutation.Dental(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: queue.FieldDental,
		})
		q.Dental = value
	}
	if value, ok := qc.mutation.QueueTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: queue.FieldQueueTime,
		})
		q.QueueTime = value
	}
	if nodes := qc.mutation.DentistIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := qc.mutation.NurseIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := qc.mutation.PatientIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return q, _spec
}
