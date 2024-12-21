// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/weiloon1234/gokit/ent/bank"
	"github.com/weiloon1234/gokit/ent/country"
	"github.com/weiloon1234/gokit/ent/predicate"
	"github.com/weiloon1234/gokit/ent/user"
)

// BankQuery is the builder for querying Bank entities.
type BankQuery struct {
	config
	ctx         *QueryContext
	order       []bank.OrderOption
	inters      []Interceptor
	predicates  []predicate.Bank
	withCountry *CountryQuery
	withUsers   *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BankQuery builder.
func (bq *BankQuery) Where(ps ...predicate.Bank) *BankQuery {
	bq.predicates = append(bq.predicates, ps...)
	return bq
}

// Limit the number of records to be returned by this query.
func (bq *BankQuery) Limit(limit int) *BankQuery {
	bq.ctx.Limit = &limit
	return bq
}

// Offset to start from.
func (bq *BankQuery) Offset(offset int) *BankQuery {
	bq.ctx.Offset = &offset
	return bq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bq *BankQuery) Unique(unique bool) *BankQuery {
	bq.ctx.Unique = &unique
	return bq
}

// Order specifies how the records should be ordered.
func (bq *BankQuery) Order(o ...bank.OrderOption) *BankQuery {
	bq.order = append(bq.order, o...)
	return bq
}

// QueryCountry chains the current query on the "country" edge.
func (bq *BankQuery) QueryCountry() *CountryQuery {
	query := (&CountryClient{config: bq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bank.Table, bank.FieldID, selector),
			sqlgraph.To(country.Table, country.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, bank.CountryTable, bank.CountryColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUsers chains the current query on the "users" edge.
func (bq *BankQuery) QueryUsers() *UserQuery {
	query := (&UserClient{config: bq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bank.Table, bank.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, bank.UsersTable, bank.UsersColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Bank entity from the query.
// Returns a *NotFoundError when no Bank was found.
func (bq *BankQuery) First(ctx context.Context) (*Bank, error) {
	nodes, err := bq.Limit(1).All(setContextOp(ctx, bq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{bank.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bq *BankQuery) FirstX(ctx context.Context) *Bank {
	node, err := bq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Bank ID from the query.
// Returns a *NotFoundError when no Bank ID was found.
func (bq *BankQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = bq.Limit(1).IDs(setContextOp(ctx, bq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{bank.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bq *BankQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := bq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Bank entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Bank entity is found.
// Returns a *NotFoundError when no Bank entities are found.
func (bq *BankQuery) Only(ctx context.Context) (*Bank, error) {
	nodes, err := bq.Limit(2).All(setContextOp(ctx, bq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{bank.Label}
	default:
		return nil, &NotSingularError{bank.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bq *BankQuery) OnlyX(ctx context.Context) *Bank {
	node, err := bq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Bank ID in the query.
// Returns a *NotSingularError when more than one Bank ID is found.
// Returns a *NotFoundError when no entities are found.
func (bq *BankQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = bq.Limit(2).IDs(setContextOp(ctx, bq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{bank.Label}
	default:
		err = &NotSingularError{bank.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bq *BankQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := bq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Banks.
func (bq *BankQuery) All(ctx context.Context) ([]*Bank, error) {
	ctx = setContextOp(ctx, bq.ctx, ent.OpQueryAll)
	if err := bq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Bank, *BankQuery]()
	return withInterceptors[[]*Bank](ctx, bq, qr, bq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bq *BankQuery) AllX(ctx context.Context) []*Bank {
	nodes, err := bq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Bank IDs.
func (bq *BankQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if bq.ctx.Unique == nil && bq.path != nil {
		bq.Unique(true)
	}
	ctx = setContextOp(ctx, bq.ctx, ent.OpQueryIDs)
	if err = bq.Select(bank.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bq *BankQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := bq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bq *BankQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bq.ctx, ent.OpQueryCount)
	if err := bq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bq, querierCount[*BankQuery](), bq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bq *BankQuery) CountX(ctx context.Context) int {
	count, err := bq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bq *BankQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bq.ctx, ent.OpQueryExist)
	switch _, err := bq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bq *BankQuery) ExistX(ctx context.Context) bool {
	exist, err := bq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BankQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bq *BankQuery) Clone() *BankQuery {
	if bq == nil {
		return nil
	}
	return &BankQuery{
		config:      bq.config,
		ctx:         bq.ctx.Clone(),
		order:       append([]bank.OrderOption{}, bq.order...),
		inters:      append([]Interceptor{}, bq.inters...),
		predicates:  append([]predicate.Bank{}, bq.predicates...),
		withCountry: bq.withCountry.Clone(),
		withUsers:   bq.withUsers.Clone(),
		// clone intermediate query.
		sql:  bq.sql.Clone(),
		path: bq.path,
	}
}

// WithCountry tells the query-builder to eager-load the nodes that are connected to
// the "country" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BankQuery) WithCountry(opts ...func(*CountryQuery)) *BankQuery {
	query := (&CountryClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bq.withCountry = query
	return bq
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BankQuery) WithUsers(opts ...func(*UserQuery)) *BankQuery {
	query := (&UserClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bq.withUsers = query
	return bq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		NameEn string `json:"name_en,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Bank.Query().
//		GroupBy(bank.FieldNameEn).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bq *BankQuery) GroupBy(field string, fields ...string) *BankGroupBy {
	bq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BankGroupBy{build: bq}
	grbuild.flds = &bq.ctx.Fields
	grbuild.label = bank.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		NameEn string `json:"name_en,omitempty"`
//	}
//
//	client.Bank.Query().
//		Select(bank.FieldNameEn).
//		Scan(ctx, &v)
func (bq *BankQuery) Select(fields ...string) *BankSelect {
	bq.ctx.Fields = append(bq.ctx.Fields, fields...)
	sbuild := &BankSelect{BankQuery: bq}
	sbuild.label = bank.Label
	sbuild.flds, sbuild.scan = &bq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BankSelect configured with the given aggregations.
func (bq *BankQuery) Aggregate(fns ...AggregateFunc) *BankSelect {
	return bq.Select().Aggregate(fns...)
}

func (bq *BankQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bq); err != nil {
				return err
			}
		}
	}
	for _, f := range bq.ctx.Fields {
		if !bank.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bq.path != nil {
		prev, err := bq.path(ctx)
		if err != nil {
			return err
		}
		bq.sql = prev
	}
	return nil
}

func (bq *BankQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Bank, error) {
	var (
		nodes       = []*Bank{}
		_spec       = bq.querySpec()
		loadedTypes = [2]bool{
			bq.withCountry != nil,
			bq.withUsers != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Bank).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Bank{config: bq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := bq.withCountry; query != nil {
		if err := bq.loadCountry(ctx, query, nodes, nil,
			func(n *Bank, e *Country) { n.Edges.Country = e }); err != nil {
			return nil, err
		}
	}
	if query := bq.withUsers; query != nil {
		if err := bq.loadUsers(ctx, query, nodes,
			func(n *Bank) { n.Edges.Users = []*User{} },
			func(n *Bank, e *User) { n.Edges.Users = append(n.Edges.Users, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (bq *BankQuery) loadCountry(ctx context.Context, query *CountryQuery, nodes []*Bank, init func(*Bank), assign func(*Bank, *Country)) error {
	ids := make([]uint64, 0, len(nodes))
	nodeids := make(map[uint64][]*Bank)
	for i := range nodes {
		if nodes[i].CountryID == nil {
			continue
		}
		fk := *nodes[i].CountryID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(country.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "country_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (bq *BankQuery) loadUsers(ctx context.Context, query *UserQuery, nodes []*Bank, init func(*Bank), assign func(*Bank, *User)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uint64]*Bank)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(user.FieldBankID)
	}
	query.Where(predicate.User(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(bank.UsersColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.BankID
		if fk == nil {
			return fmt.Errorf(`foreign-key "bank_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "bank_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (bq *BankQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bq.querySpec()
	_spec.Node.Columns = bq.ctx.Fields
	if len(bq.ctx.Fields) > 0 {
		_spec.Unique = bq.ctx.Unique != nil && *bq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bq.driver, _spec)
}

func (bq *BankQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(bank.Table, bank.Columns, sqlgraph.NewFieldSpec(bank.FieldID, field.TypeUint64))
	_spec.From = bq.sql
	if unique := bq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bq.path != nil {
		_spec.Unique = true
	}
	if fields := bq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bank.FieldID)
		for i := range fields {
			if fields[i] != bank.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if bq.withCountry != nil {
			_spec.Node.AddColumnOnce(bank.FieldCountryID)
		}
	}
	if ps := bq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bq *BankQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bq.driver.Dialect())
	t1 := builder.Table(bank.Table)
	columns := bq.ctx.Fields
	if len(columns) == 0 {
		columns = bank.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bq.sql != nil {
		selector = bq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bq.ctx.Unique != nil && *bq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range bq.predicates {
		p(selector)
	}
	for _, p := range bq.order {
		p(selector)
	}
	if offset := bq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BankGroupBy is the group-by builder for Bank entities.
type BankGroupBy struct {
	selector
	build *BankQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BankGroupBy) Aggregate(fns ...AggregateFunc) *BankGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the selector query and scans the result into the given value.
func (bgb *BankGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bgb.build.ctx, ent.OpQueryGroupBy)
	if err := bgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BankQuery, *BankGroupBy](ctx, bgb.build, bgb, bgb.build.inters, v)
}

func (bgb *BankGroupBy) sqlScan(ctx context.Context, root *BankQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bgb.fns))
	for _, fn := range bgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bgb.flds)+len(bgb.fns))
		for _, f := range *bgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BankSelect is the builder for selecting fields of Bank entities.
type BankSelect struct {
	*BankQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bs *BankSelect) Aggregate(fns ...AggregateFunc) *BankSelect {
	bs.fns = append(bs.fns, fns...)
	return bs
}

// Scan applies the selector query and scans the result into the given value.
func (bs *BankSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bs.ctx, ent.OpQuerySelect)
	if err := bs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BankQuery, *BankSelect](ctx, bs.BankQuery, bs, bs.inters, v)
}

func (bs *BankSelect) sqlScan(ctx context.Context, root *BankQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bs.fns))
	for _, fn := range bs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
