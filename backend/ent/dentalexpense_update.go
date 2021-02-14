// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team03/app/ent/dentalexpense"
	"github.com/team03/app/ent/medicalfile"
	"github.com/team03/app/ent/nurse"
	"github.com/team03/app/ent/predicate"
	"github.com/team03/app/ent/pricetype"
)

// DentalexpenseUpdate is the builder for updating Dentalexpense entities.
type DentalexpenseUpdate struct {
	config
	hooks      []Hook
	mutation   *DentalexpenseMutation
	predicates []predicate.Dentalexpense
}

// Where adds a new predicate for the builder.
func (du *DentalexpenseUpdate) Where(ps ...predicate.Dentalexpense) *DentalexpenseUpdate {
	du.predicates = append(du.predicates, ps...)
	return du
}

// SetName sets the Name field.
func (du *DentalexpenseUpdate) SetName(s string) *DentalexpenseUpdate {
	du.mutation.SetName(s)
	return du
}

// SetPhone sets the Phone field.
func (du *DentalexpenseUpdate) SetPhone(s string) *DentalexpenseUpdate {
	du.mutation.SetPhone(s)
	return du
}

// SetAddedTime sets the AddedTime field.
func (du *DentalexpenseUpdate) SetAddedTime(t time.Time) *DentalexpenseUpdate {
	du.mutation.SetAddedTime(t)
	return du
}

// SetAmount sets the Amount field.
func (du *DentalexpenseUpdate) SetAmount(i int) *DentalexpenseUpdate {
	du.mutation.ResetAmount()
	du.mutation.SetAmount(i)
	return du
}

// AddAmount adds i to Amount.
func (du *DentalexpenseUpdate) AddAmount(i int) *DentalexpenseUpdate {
	du.mutation.AddAmount(i)
	return du
}

// SetTax sets the Tax field.
func (du *DentalexpenseUpdate) SetTax(s string) *DentalexpenseUpdate {
	du.mutation.SetTax(s)
	return du
}

// SetNurseID sets the nurse edge to Nurse by id.
func (du *DentalexpenseUpdate) SetNurseID(id int) *DentalexpenseUpdate {
	du.mutation.SetNurseID(id)
	return du
}

// SetNillableNurseID sets the nurse edge to Nurse by id if the given value is not nil.
func (du *DentalexpenseUpdate) SetNillableNurseID(id *int) *DentalexpenseUpdate {
	if id != nil {
		du = du.SetNurseID(*id)
	}
	return du
}

// SetNurse sets the nurse edge to Nurse.
func (du *DentalexpenseUpdate) SetNurse(n *Nurse) *DentalexpenseUpdate {
	return du.SetNurseID(n.ID)
}

// SetMedicalfileID sets the medicalfile edge to Medicalfile by id.
func (du *DentalexpenseUpdate) SetMedicalfileID(id int) *DentalexpenseUpdate {
	du.mutation.SetMedicalfileID(id)
	return du
}

// SetNillableMedicalfileID sets the medicalfile edge to Medicalfile by id if the given value is not nil.
func (du *DentalexpenseUpdate) SetNillableMedicalfileID(id *int) *DentalexpenseUpdate {
	if id != nil {
		du = du.SetMedicalfileID(*id)
	}
	return du
}

// SetMedicalfile sets the medicalfile edge to Medicalfile.
func (du *DentalexpenseUpdate) SetMedicalfile(m *Medicalfile) *DentalexpenseUpdate {
	return du.SetMedicalfileID(m.ID)
}

// SetPricetypeID sets the pricetype edge to Pricetype by id.
func (du *DentalexpenseUpdate) SetPricetypeID(id int) *DentalexpenseUpdate {
	du.mutation.SetPricetypeID(id)
	return du
}

// SetNillablePricetypeID sets the pricetype edge to Pricetype by id if the given value is not nil.
func (du *DentalexpenseUpdate) SetNillablePricetypeID(id *int) *DentalexpenseUpdate {
	if id != nil {
		du = du.SetPricetypeID(*id)
	}
	return du
}

// SetPricetype sets the pricetype edge to Pricetype.
func (du *DentalexpenseUpdate) SetPricetype(p *Pricetype) *DentalexpenseUpdate {
	return du.SetPricetypeID(p.ID)
}

// Mutation returns the DentalexpenseMutation object of the builder.
func (du *DentalexpenseUpdate) Mutation() *DentalexpenseMutation {
	return du.mutation
}

// ClearNurse clears the nurse edge to Nurse.
func (du *DentalexpenseUpdate) ClearNurse() *DentalexpenseUpdate {
	du.mutation.ClearNurse()
	return du
}

// ClearMedicalfile clears the medicalfile edge to Medicalfile.
func (du *DentalexpenseUpdate) ClearMedicalfile() *DentalexpenseUpdate {
	du.mutation.ClearMedicalfile()
	return du
}

// ClearPricetype clears the pricetype edge to Pricetype.
func (du *DentalexpenseUpdate) ClearPricetype() *DentalexpenseUpdate {
	du.mutation.ClearPricetype()
	return du
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (du *DentalexpenseUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := du.mutation.Name(); ok {
		if err := dentalexpense.NameValidator(v); err != nil {
			return 0, &ValidationError{Name: "Name", err: fmt.Errorf("ent: validator failed for field \"Name\": %w", err)}
		}
	}
	if v, ok := du.mutation.Phone(); ok {
		if err := dentalexpense.PhoneValidator(v); err != nil {
			return 0, &ValidationError{Name: "Phone", err: fmt.Errorf("ent: validator failed for field \"Phone\": %w", err)}
		}
	}
	if v, ok := du.mutation.Amount(); ok {
		if err := dentalexpense.AmountValidator(v); err != nil {
			return 0, &ValidationError{Name: "Amount", err: fmt.Errorf("ent: validator failed for field \"Amount\": %w", err)}
		}
	}
	if v, ok := du.mutation.Tax(); ok {
		if err := dentalexpense.TaxValidator(v); err != nil {
			return 0, &ValidationError{Name: "Tax", err: fmt.Errorf("ent: validator failed for field \"Tax\": %w", err)}
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
			mutation, ok := m.(*DentalexpenseMutation)
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
func (du *DentalexpenseUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DentalexpenseUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DentalexpenseUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DentalexpenseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dentalexpense.Table,
			Columns: dentalexpense.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dentalexpense.FieldID,
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
	if value, ok := du.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dentalexpense.FieldName,
		})
	}
	if value, ok := du.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dentalexpense.FieldPhone,
		})
	}
	if value, ok := du.mutation.AddedTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dentalexpense.FieldAddedTime,
		})
	}
	if value, ok := du.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: dentalexpense.FieldAmount,
		})
	}
	if value, ok := du.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: dentalexpense.FieldAmount,
		})
	}
	if value, ok := du.mutation.Tax(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dentalexpense.FieldTax,
		})
	}
	if du.mutation.NurseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dentalexpense.NurseTable,
			Columns: []string{dentalexpense.NurseColumn},
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
	if nodes := du.mutation.NurseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dentalexpense.NurseTable,
			Columns: []string{dentalexpense.NurseColumn},
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
	if du.mutation.MedicalfileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dentalexpense.MedicalfileTable,
			Columns: []string{dentalexpense.MedicalfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicalfile.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.MedicalfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dentalexpense.MedicalfileTable,
			Columns: []string{dentalexpense.MedicalfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicalfile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.PricetypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dentalexpense.PricetypeTable,
			Columns: []string{dentalexpense.PricetypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pricetype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.PricetypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dentalexpense.PricetypeTable,
			Columns: []string{dentalexpense.PricetypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pricetype.FieldID,
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
			err = &NotFoundError{dentalexpense.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DentalexpenseUpdateOne is the builder for updating a single Dentalexpense entity.
type DentalexpenseUpdateOne struct {
	config
	hooks    []Hook
	mutation *DentalexpenseMutation
}

// SetName sets the Name field.
func (duo *DentalexpenseUpdateOne) SetName(s string) *DentalexpenseUpdateOne {
	duo.mutation.SetName(s)
	return duo
}

// SetPhone sets the Phone field.
func (duo *DentalexpenseUpdateOne) SetPhone(s string) *DentalexpenseUpdateOne {
	duo.mutation.SetPhone(s)
	return duo
}

// SetAddedTime sets the AddedTime field.
func (duo *DentalexpenseUpdateOne) SetAddedTime(t time.Time) *DentalexpenseUpdateOne {
	duo.mutation.SetAddedTime(t)
	return duo
}

// SetAmount sets the Amount field.
func (duo *DentalexpenseUpdateOne) SetAmount(i int) *DentalexpenseUpdateOne {
	duo.mutation.ResetAmount()
	duo.mutation.SetAmount(i)
	return duo
}

// AddAmount adds i to Amount.
func (duo *DentalexpenseUpdateOne) AddAmount(i int) *DentalexpenseUpdateOne {
	duo.mutation.AddAmount(i)
	return duo
}

// SetTax sets the Tax field.
func (duo *DentalexpenseUpdateOne) SetTax(s string) *DentalexpenseUpdateOne {
	duo.mutation.SetTax(s)
	return duo
}

// SetNurseID sets the nurse edge to Nurse by id.
func (duo *DentalexpenseUpdateOne) SetNurseID(id int) *DentalexpenseUpdateOne {
	duo.mutation.SetNurseID(id)
	return duo
}

// SetNillableNurseID sets the nurse edge to Nurse by id if the given value is not nil.
func (duo *DentalexpenseUpdateOne) SetNillableNurseID(id *int) *DentalexpenseUpdateOne {
	if id != nil {
		duo = duo.SetNurseID(*id)
	}
	return duo
}

// SetNurse sets the nurse edge to Nurse.
func (duo *DentalexpenseUpdateOne) SetNurse(n *Nurse) *DentalexpenseUpdateOne {
	return duo.SetNurseID(n.ID)
}

// SetMedicalfileID sets the medicalfile edge to Medicalfile by id.
func (duo *DentalexpenseUpdateOne) SetMedicalfileID(id int) *DentalexpenseUpdateOne {
	duo.mutation.SetMedicalfileID(id)
	return duo
}

// SetNillableMedicalfileID sets the medicalfile edge to Medicalfile by id if the given value is not nil.
func (duo *DentalexpenseUpdateOne) SetNillableMedicalfileID(id *int) *DentalexpenseUpdateOne {
	if id != nil {
		duo = duo.SetMedicalfileID(*id)
	}
	return duo
}

// SetMedicalfile sets the medicalfile edge to Medicalfile.
func (duo *DentalexpenseUpdateOne) SetMedicalfile(m *Medicalfile) *DentalexpenseUpdateOne {
	return duo.SetMedicalfileID(m.ID)
}

// SetPricetypeID sets the pricetype edge to Pricetype by id.
func (duo *DentalexpenseUpdateOne) SetPricetypeID(id int) *DentalexpenseUpdateOne {
	duo.mutation.SetPricetypeID(id)
	return duo
}

// SetNillablePricetypeID sets the pricetype edge to Pricetype by id if the given value is not nil.
func (duo *DentalexpenseUpdateOne) SetNillablePricetypeID(id *int) *DentalexpenseUpdateOne {
	if id != nil {
		duo = duo.SetPricetypeID(*id)
	}
	return duo
}

// SetPricetype sets the pricetype edge to Pricetype.
func (duo *DentalexpenseUpdateOne) SetPricetype(p *Pricetype) *DentalexpenseUpdateOne {
	return duo.SetPricetypeID(p.ID)
}

// Mutation returns the DentalexpenseMutation object of the builder.
func (duo *DentalexpenseUpdateOne) Mutation() *DentalexpenseMutation {
	return duo.mutation
}

// ClearNurse clears the nurse edge to Nurse.
func (duo *DentalexpenseUpdateOne) ClearNurse() *DentalexpenseUpdateOne {
	duo.mutation.ClearNurse()
	return duo
}

// ClearMedicalfile clears the medicalfile edge to Medicalfile.
func (duo *DentalexpenseUpdateOne) ClearMedicalfile() *DentalexpenseUpdateOne {
	duo.mutation.ClearMedicalfile()
	return duo
}

// ClearPricetype clears the pricetype edge to Pricetype.
func (duo *DentalexpenseUpdateOne) ClearPricetype() *DentalexpenseUpdateOne {
	duo.mutation.ClearPricetype()
	return duo
}

// Save executes the query and returns the updated entity.
func (duo *DentalexpenseUpdateOne) Save(ctx context.Context) (*Dentalexpense, error) {
	if v, ok := duo.mutation.Name(); ok {
		if err := dentalexpense.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "Name", err: fmt.Errorf("ent: validator failed for field \"Name\": %w", err)}
		}
	}
	if v, ok := duo.mutation.Phone(); ok {
		if err := dentalexpense.PhoneValidator(v); err != nil {
			return nil, &ValidationError{Name: "Phone", err: fmt.Errorf("ent: validator failed for field \"Phone\": %w", err)}
		}
	}
	if v, ok := duo.mutation.Amount(); ok {
		if err := dentalexpense.AmountValidator(v); err != nil {
			return nil, &ValidationError{Name: "Amount", err: fmt.Errorf("ent: validator failed for field \"Amount\": %w", err)}
		}
	}
	if v, ok := duo.mutation.Tax(); ok {
		if err := dentalexpense.TaxValidator(v); err != nil {
			return nil, &ValidationError{Name: "Tax", err: fmt.Errorf("ent: validator failed for field \"Tax\": %w", err)}
		}
	}

	var (
		err  error
		node *Dentalexpense
	)
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DentalexpenseMutation)
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
func (duo *DentalexpenseUpdateOne) SaveX(ctx context.Context) *Dentalexpense {
	d, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return d
}

// Exec executes the query on the entity.
func (duo *DentalexpenseUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DentalexpenseUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DentalexpenseUpdateOne) sqlSave(ctx context.Context) (d *Dentalexpense, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dentalexpense.Table,
			Columns: dentalexpense.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dentalexpense.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Dentalexpense.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := duo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dentalexpense.FieldName,
		})
	}
	if value, ok := duo.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dentalexpense.FieldPhone,
		})
	}
	if value, ok := duo.mutation.AddedTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dentalexpense.FieldAddedTime,
		})
	}
	if value, ok := duo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: dentalexpense.FieldAmount,
		})
	}
	if value, ok := duo.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: dentalexpense.FieldAmount,
		})
	}
	if value, ok := duo.mutation.Tax(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dentalexpense.FieldTax,
		})
	}
	if duo.mutation.NurseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dentalexpense.NurseTable,
			Columns: []string{dentalexpense.NurseColumn},
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
	if nodes := duo.mutation.NurseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dentalexpense.NurseTable,
			Columns: []string{dentalexpense.NurseColumn},
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
	if duo.mutation.MedicalfileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dentalexpense.MedicalfileTable,
			Columns: []string{dentalexpense.MedicalfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicalfile.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.MedicalfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dentalexpense.MedicalfileTable,
			Columns: []string{dentalexpense.MedicalfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicalfile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.PricetypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dentalexpense.PricetypeTable,
			Columns: []string{dentalexpense.PricetypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pricetype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.PricetypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dentalexpense.PricetypeTable,
			Columns: []string{dentalexpense.PricetypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pricetype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	d = &Dentalexpense{config: duo.config}
	_spec.Assign = d.assignValues
	_spec.ScanValues = d.scanValues()
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dentalexpense.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return d, nil
}
