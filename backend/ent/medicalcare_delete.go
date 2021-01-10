// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team03/app/ent/medicalcare"
	"github.com/team03/app/ent/predicate"
)

// MedicalCareDelete is the builder for deleting a MedicalCare entity.
type MedicalCareDelete struct {
	config
	hooks      []Hook
	mutation   *MedicalCareMutation
	predicates []predicate.MedicalCare
}

// Where adds a new predicate to the delete builder.
func (mcd *MedicalCareDelete) Where(ps ...predicate.MedicalCare) *MedicalCareDelete {
	mcd.predicates = append(mcd.predicates, ps...)
	return mcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mcd *MedicalCareDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mcd.hooks) == 0 {
		affected, err = mcd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MedicalCareMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mcd.mutation = mutation
			affected, err = mcd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mcd.hooks) - 1; i >= 0; i-- {
			mut = mcd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mcd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcd *MedicalCareDelete) ExecX(ctx context.Context) int {
	n, err := mcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mcd *MedicalCareDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: medicalcare.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: medicalcare.FieldID,
			},
		},
	}
	if ps := mcd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, mcd.driver, _spec)
}

// MedicalCareDeleteOne is the builder for deleting a single MedicalCare entity.
type MedicalCareDeleteOne struct {
	mcd *MedicalCareDelete
}

// Exec executes the deletion query.
func (mcdo *MedicalCareDeleteOne) Exec(ctx context.Context) error {
	n, err := mcdo.mcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{medicalcare.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mcdo *MedicalCareDeleteOne) ExecX(ctx context.Context) {
	mcdo.mcd.ExecX(ctx)
}
