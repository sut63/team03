// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team03/app/ent/dentalexpense"
	"github.com/team03/app/ent/pricetype"
)

// PriceTypeCreate is the builder for creating a PriceType entity.
type PriceTypeCreate struct {
	config
	mutation *PriceTypeMutation
	hooks    []Hook
}

// SetName sets the name field.
func (ptc *PriceTypeCreate) SetName(s string) *PriceTypeCreate {
	ptc.mutation.SetName(s)
	return ptc
}

// AddDentalexpenseIDs adds the dentalexpenses edge to DentalExpense by ids.
func (ptc *PriceTypeCreate) AddDentalexpenseIDs(ids ...int) *PriceTypeCreate {
	ptc.mutation.AddDentalexpenseIDs(ids...)
	return ptc
}

// AddDentalexpenses adds the dentalexpenses edges to DentalExpense.
func (ptc *PriceTypeCreate) AddDentalexpenses(d ...*DentalExpense) *PriceTypeCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return ptc.AddDentalexpenseIDs(ids...)
}

// Mutation returns the PriceTypeMutation object of the builder.
func (ptc *PriceTypeCreate) Mutation() *PriceTypeMutation {
	return ptc.mutation
}

// Save creates the PriceType in the database.
func (ptc *PriceTypeCreate) Save(ctx context.Context) (*PriceType, error) {
	if _, ok := ptc.mutation.Name(); !ok {
		return nil, &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	var (
		err  error
		node *PriceType
	)
	if len(ptc.hooks) == 0 {
		node, err = ptc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PriceTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ptc.mutation = mutation
			node, err = ptc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ptc.hooks) - 1; i >= 0; i-- {
			mut = ptc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ptc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ptc *PriceTypeCreate) SaveX(ctx context.Context) *PriceType {
	v, err := ptc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ptc *PriceTypeCreate) sqlSave(ctx context.Context) (*PriceType, error) {
	pt, _spec := ptc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ptc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	pt.ID = int(id)
	return pt, nil
}

func (ptc *PriceTypeCreate) createSpec() (*PriceType, *sqlgraph.CreateSpec) {
	var (
		pt    = &PriceType{config: ptc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: pricetype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pricetype.FieldID,
			},
		}
	)
	if value, ok := ptc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pricetype.FieldName,
		})
		pt.Name = value
	}
	if nodes := ptc.mutation.DentalexpensesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pricetype.DentalexpensesTable,
			Columns: []string{pricetype.DentalexpensesColumn},
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
	return pt, _spec
}
