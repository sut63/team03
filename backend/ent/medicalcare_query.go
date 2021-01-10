// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team03/app/ent/medicalcare"
	"github.com/team03/app/ent/medicalfile"
	"github.com/team03/app/ent/patient"
	"github.com/team03/app/ent/predicate"
)

// MedicalCareQuery is the builder for querying MedicalCare entities.
type MedicalCareQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.MedicalCare
	// eager-loading edges.
	withPatients     *PatientQuery
	withMedicalfiles *MedicalfileQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (mcq *MedicalCareQuery) Where(ps ...predicate.MedicalCare) *MedicalCareQuery {
	mcq.predicates = append(mcq.predicates, ps...)
	return mcq
}

// Limit adds a limit step to the query.
func (mcq *MedicalCareQuery) Limit(limit int) *MedicalCareQuery {
	mcq.limit = &limit
	return mcq
}

// Offset adds an offset step to the query.
func (mcq *MedicalCareQuery) Offset(offset int) *MedicalCareQuery {
	mcq.offset = &offset
	return mcq
}

// Order adds an order step to the query.
func (mcq *MedicalCareQuery) Order(o ...OrderFunc) *MedicalCareQuery {
	mcq.order = append(mcq.order, o...)
	return mcq
}

// QueryPatients chains the current query on the patients edge.
func (mcq *MedicalCareQuery) QueryPatients() *PatientQuery {
	query := &PatientQuery{config: mcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(medicalcare.Table, medicalcare.FieldID, mcq.sqlQuery()),
			sqlgraph.To(patient.Table, patient.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, medicalcare.PatientsTable, medicalcare.PatientsColumn),
		)
		fromU = sqlgraph.SetNeighbors(mcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMedicalfiles chains the current query on the medicalfiles edge.
func (mcq *MedicalCareQuery) QueryMedicalfiles() *MedicalfileQuery {
	query := &MedicalfileQuery{config: mcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(medicalcare.Table, medicalcare.FieldID, mcq.sqlQuery()),
			sqlgraph.To(medicalfile.Table, medicalfile.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, medicalcare.MedicalfilesTable, medicalcare.MedicalfilesColumn),
		)
		fromU = sqlgraph.SetNeighbors(mcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MedicalCare entity in the query. Returns *NotFoundError when no medicalcare was found.
func (mcq *MedicalCareQuery) First(ctx context.Context) (*MedicalCare, error) {
	mcs, err := mcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(mcs) == 0 {
		return nil, &NotFoundError{medicalcare.Label}
	}
	return mcs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mcq *MedicalCareQuery) FirstX(ctx context.Context) *MedicalCare {
	mc, err := mcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return mc
}

// FirstID returns the first MedicalCare id in the query. Returns *NotFoundError when no id was found.
func (mcq *MedicalCareQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{medicalcare.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (mcq *MedicalCareQuery) FirstXID(ctx context.Context) int {
	id, err := mcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only MedicalCare entity in the query, returns an error if not exactly one entity was returned.
func (mcq *MedicalCareQuery) Only(ctx context.Context) (*MedicalCare, error) {
	mcs, err := mcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(mcs) {
	case 1:
		return mcs[0], nil
	case 0:
		return nil, &NotFoundError{medicalcare.Label}
	default:
		return nil, &NotSingularError{medicalcare.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mcq *MedicalCareQuery) OnlyX(ctx context.Context) *MedicalCare {
	mc, err := mcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return mc
}

// OnlyID returns the only MedicalCare id in the query, returns an error if not exactly one id was returned.
func (mcq *MedicalCareQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{medicalcare.Label}
	default:
		err = &NotSingularError{medicalcare.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mcq *MedicalCareQuery) OnlyIDX(ctx context.Context) int {
	id, err := mcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MedicalCares.
func (mcq *MedicalCareQuery) All(ctx context.Context) ([]*MedicalCare, error) {
	if err := mcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mcq *MedicalCareQuery) AllX(ctx context.Context) []*MedicalCare {
	mcs, err := mcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return mcs
}

// IDs executes the query and returns a list of MedicalCare ids.
func (mcq *MedicalCareQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := mcq.Select(medicalcare.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mcq *MedicalCareQuery) IDsX(ctx context.Context) []int {
	ids, err := mcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mcq *MedicalCareQuery) Count(ctx context.Context) (int, error) {
	if err := mcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mcq *MedicalCareQuery) CountX(ctx context.Context) int {
	count, err := mcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mcq *MedicalCareQuery) Exist(ctx context.Context) (bool, error) {
	if err := mcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mcq *MedicalCareQuery) ExistX(ctx context.Context) bool {
	exist, err := mcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mcq *MedicalCareQuery) Clone() *MedicalCareQuery {
	return &MedicalCareQuery{
		config:     mcq.config,
		limit:      mcq.limit,
		offset:     mcq.offset,
		order:      append([]OrderFunc{}, mcq.order...),
		unique:     append([]string{}, mcq.unique...),
		predicates: append([]predicate.MedicalCare{}, mcq.predicates...),
		// clone intermediate query.
		sql:  mcq.sql.Clone(),
		path: mcq.path,
	}
}

//  WithPatients tells the query-builder to eager-loads the nodes that are connected to
// the "patients" edge. The optional arguments used to configure the query builder of the edge.
func (mcq *MedicalCareQuery) WithPatients(opts ...func(*PatientQuery)) *MedicalCareQuery {
	query := &PatientQuery{config: mcq.config}
	for _, opt := range opts {
		opt(query)
	}
	mcq.withPatients = query
	return mcq
}

//  WithMedicalfiles tells the query-builder to eager-loads the nodes that are connected to
// the "medicalfiles" edge. The optional arguments used to configure the query builder of the edge.
func (mcq *MedicalCareQuery) WithMedicalfiles(opts ...func(*MedicalfileQuery)) *MedicalCareQuery {
	query := &MedicalfileQuery{config: mcq.config}
	for _, opt := range opts {
		opt(query)
	}
	mcq.withMedicalfiles = query
	return mcq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MedicalCare.Query().
//		GroupBy(medicalcare.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (mcq *MedicalCareQuery) GroupBy(field string, fields ...string) *MedicalCareGroupBy {
	group := &MedicalCareGroupBy{config: mcq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mcq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.MedicalCare.Query().
//		Select(medicalcare.FieldName).
//		Scan(ctx, &v)
//
func (mcq *MedicalCareQuery) Select(field string, fields ...string) *MedicalCareSelect {
	selector := &MedicalCareSelect{config: mcq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mcq.sqlQuery(), nil
	}
	return selector
}

func (mcq *MedicalCareQuery) prepareQuery(ctx context.Context) error {
	if mcq.path != nil {
		prev, err := mcq.path(ctx)
		if err != nil {
			return err
		}
		mcq.sql = prev
	}
	return nil
}

func (mcq *MedicalCareQuery) sqlAll(ctx context.Context) ([]*MedicalCare, error) {
	var (
		nodes       = []*MedicalCare{}
		_spec       = mcq.querySpec()
		loadedTypes = [2]bool{
			mcq.withPatients != nil,
			mcq.withMedicalfiles != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &MedicalCare{config: mcq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, mcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := mcq.withPatients; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*MedicalCare)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Patient(func(s *sql.Selector) {
			s.Where(sql.InValues(medicalcare.PatientsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.medicalcare_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "medicalcare_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "medicalcare_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Patients = append(node.Edges.Patients, n)
		}
	}

	if query := mcq.withMedicalfiles; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*MedicalCare)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Medicalfile(func(s *sql.Selector) {
			s.Where(sql.InValues(medicalcare.MedicalfilesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.medicalcare_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "medicalcare_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "medicalcare_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Medicalfiles = append(node.Edges.Medicalfiles, n)
		}
	}

	return nodes, nil
}

func (mcq *MedicalCareQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mcq.querySpec()
	return sqlgraph.CountNodes(ctx, mcq.driver, _spec)
}

func (mcq *MedicalCareQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := mcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (mcq *MedicalCareQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   medicalcare.Table,
			Columns: medicalcare.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: medicalcare.FieldID,
			},
		},
		From:   mcq.sql,
		Unique: true,
	}
	if ps := mcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mcq *MedicalCareQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(mcq.driver.Dialect())
	t1 := builder.Table(medicalcare.Table)
	selector := builder.Select(t1.Columns(medicalcare.Columns...)...).From(t1)
	if mcq.sql != nil {
		selector = mcq.sql
		selector.Select(selector.Columns(medicalcare.Columns...)...)
	}
	for _, p := range mcq.predicates {
		p(selector)
	}
	for _, p := range mcq.order {
		p(selector)
	}
	if offset := mcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MedicalCareGroupBy is the builder for group-by MedicalCare entities.
type MedicalCareGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mcgb *MedicalCareGroupBy) Aggregate(fns ...AggregateFunc) *MedicalCareGroupBy {
	mcgb.fns = append(mcgb.fns, fns...)
	return mcgb
}

// Scan applies the group-by query and scan the result into the given value.
func (mcgb *MedicalCareGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := mcgb.path(ctx)
	if err != nil {
		return err
	}
	mcgb.sql = query
	return mcgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mcgb *MedicalCareGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := mcgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (mcgb *MedicalCareGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(mcgb.fields) > 1 {
		return nil, errors.New("ent: MedicalCareGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := mcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mcgb *MedicalCareGroupBy) StringsX(ctx context.Context) []string {
	v, err := mcgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (mcgb *MedicalCareGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mcgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medicalcare.Label}
	default:
		err = fmt.Errorf("ent: MedicalCareGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mcgb *MedicalCareGroupBy) StringX(ctx context.Context) string {
	v, err := mcgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (mcgb *MedicalCareGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(mcgb.fields) > 1 {
		return nil, errors.New("ent: MedicalCareGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := mcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mcgb *MedicalCareGroupBy) IntsX(ctx context.Context) []int {
	v, err := mcgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (mcgb *MedicalCareGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mcgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medicalcare.Label}
	default:
		err = fmt.Errorf("ent: MedicalCareGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mcgb *MedicalCareGroupBy) IntX(ctx context.Context) int {
	v, err := mcgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (mcgb *MedicalCareGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(mcgb.fields) > 1 {
		return nil, errors.New("ent: MedicalCareGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := mcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mcgb *MedicalCareGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := mcgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (mcgb *MedicalCareGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mcgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medicalcare.Label}
	default:
		err = fmt.Errorf("ent: MedicalCareGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mcgb *MedicalCareGroupBy) Float64X(ctx context.Context) float64 {
	v, err := mcgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (mcgb *MedicalCareGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(mcgb.fields) > 1 {
		return nil, errors.New("ent: MedicalCareGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := mcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mcgb *MedicalCareGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := mcgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (mcgb *MedicalCareGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mcgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medicalcare.Label}
	default:
		err = fmt.Errorf("ent: MedicalCareGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mcgb *MedicalCareGroupBy) BoolX(ctx context.Context) bool {
	v, err := mcgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mcgb *MedicalCareGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := mcgb.sqlQuery().Query()
	if err := mcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mcgb *MedicalCareGroupBy) sqlQuery() *sql.Selector {
	selector := mcgb.sql
	columns := make([]string, 0, len(mcgb.fields)+len(mcgb.fns))
	columns = append(columns, mcgb.fields...)
	for _, fn := range mcgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(mcgb.fields...)
}

// MedicalCareSelect is the builder for select fields of MedicalCare entities.
type MedicalCareSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (mcs *MedicalCareSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := mcs.path(ctx)
	if err != nil {
		return err
	}
	mcs.sql = query
	return mcs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mcs *MedicalCareSelect) ScanX(ctx context.Context, v interface{}) {
	if err := mcs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (mcs *MedicalCareSelect) Strings(ctx context.Context) ([]string, error) {
	if len(mcs.fields) > 1 {
		return nil, errors.New("ent: MedicalCareSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := mcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mcs *MedicalCareSelect) StringsX(ctx context.Context) []string {
	v, err := mcs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (mcs *MedicalCareSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mcs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medicalcare.Label}
	default:
		err = fmt.Errorf("ent: MedicalCareSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mcs *MedicalCareSelect) StringX(ctx context.Context) string {
	v, err := mcs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (mcs *MedicalCareSelect) Ints(ctx context.Context) ([]int, error) {
	if len(mcs.fields) > 1 {
		return nil, errors.New("ent: MedicalCareSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := mcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mcs *MedicalCareSelect) IntsX(ctx context.Context) []int {
	v, err := mcs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (mcs *MedicalCareSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mcs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medicalcare.Label}
	default:
		err = fmt.Errorf("ent: MedicalCareSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mcs *MedicalCareSelect) IntX(ctx context.Context) int {
	v, err := mcs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (mcs *MedicalCareSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(mcs.fields) > 1 {
		return nil, errors.New("ent: MedicalCareSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := mcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mcs *MedicalCareSelect) Float64sX(ctx context.Context) []float64 {
	v, err := mcs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (mcs *MedicalCareSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mcs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medicalcare.Label}
	default:
		err = fmt.Errorf("ent: MedicalCareSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mcs *MedicalCareSelect) Float64X(ctx context.Context) float64 {
	v, err := mcs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (mcs *MedicalCareSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(mcs.fields) > 1 {
		return nil, errors.New("ent: MedicalCareSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := mcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mcs *MedicalCareSelect) BoolsX(ctx context.Context) []bool {
	v, err := mcs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (mcs *MedicalCareSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mcs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medicalcare.Label}
	default:
		err = fmt.Errorf("ent: MedicalCareSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mcs *MedicalCareSelect) BoolX(ctx context.Context) bool {
	v, err := mcs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mcs *MedicalCareSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := mcs.sqlQuery().Query()
	if err := mcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mcs *MedicalCareSelect) sqlQuery() sql.Querier {
	selector := mcs.sql
	selector.Select(selector.Columns(mcs.fields...)...)
	return selector
}
