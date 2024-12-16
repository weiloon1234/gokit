// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/weiloon1234/gokit/ent/countrylocation"
	"github.com/weiloon1234/gokit/ent/predicate"
)

// CountryLocationUpdate is the builder for updating CountryLocation entities.
type CountryLocationUpdate struct {
	config
	hooks    []Hook
	mutation *CountryLocationMutation
}

// Where appends a list predicates to the CountryLocationUpdate builder.
func (clu *CountryLocationUpdate) Where(ps ...predicate.CountryLocation) *CountryLocationUpdate {
	clu.mutation.Where(ps...)
	return clu
}

// SetCountryID sets the "country_id" field.
func (clu *CountryLocationUpdate) SetCountryID(u uint64) *CountryLocationUpdate {
	clu.mutation.ResetCountryID()
	clu.mutation.SetCountryID(u)
	return clu
}

// SetNillableCountryID sets the "country_id" field if the given value is not nil.
func (clu *CountryLocationUpdate) SetNillableCountryID(u *uint64) *CountryLocationUpdate {
	if u != nil {
		clu.SetCountryID(*u)
	}
	return clu
}

// AddCountryID adds u to the "country_id" field.
func (clu *CountryLocationUpdate) AddCountryID(u int64) *CountryLocationUpdate {
	clu.mutation.AddCountryID(u)
	return clu
}

// ClearCountryID clears the value of the "country_id" field.
func (clu *CountryLocationUpdate) ClearCountryID() *CountryLocationUpdate {
	clu.mutation.ClearCountryID()
	return clu
}

// SetParentID sets the "parent_id" field.
func (clu *CountryLocationUpdate) SetParentID(u uint64) *CountryLocationUpdate {
	clu.mutation.ResetParentID()
	clu.mutation.SetParentID(u)
	return clu
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (clu *CountryLocationUpdate) SetNillableParentID(u *uint64) *CountryLocationUpdate {
	if u != nil {
		clu.SetParentID(*u)
	}
	return clu
}

// AddParentID adds u to the "parent_id" field.
func (clu *CountryLocationUpdate) AddParentID(u int64) *CountryLocationUpdate {
	clu.mutation.AddParentID(u)
	return clu
}

// ClearParentID clears the value of the "parent_id" field.
func (clu *CountryLocationUpdate) ClearParentID() *CountryLocationUpdate {
	clu.mutation.ClearParentID()
	return clu
}

// SetSorting sets the "sorting" field.
func (clu *CountryLocationUpdate) SetSorting(u uint64) *CountryLocationUpdate {
	clu.mutation.ResetSorting()
	clu.mutation.SetSorting(u)
	return clu
}

// SetNillableSorting sets the "sorting" field if the given value is not nil.
func (clu *CountryLocationUpdate) SetNillableSorting(u *uint64) *CountryLocationUpdate {
	if u != nil {
		clu.SetSorting(*u)
	}
	return clu
}

// AddSorting adds u to the "sorting" field.
func (clu *CountryLocationUpdate) AddSorting(u int64) *CountryLocationUpdate {
	clu.mutation.AddSorting(u)
	return clu
}

// SetNameEn sets the "name_en" field.
func (clu *CountryLocationUpdate) SetNameEn(s string) *CountryLocationUpdate {
	clu.mutation.SetNameEn(s)
	return clu
}

// SetNillableNameEn sets the "name_en" field if the given value is not nil.
func (clu *CountryLocationUpdate) SetNillableNameEn(s *string) *CountryLocationUpdate {
	if s != nil {
		clu.SetNameEn(*s)
	}
	return clu
}

// SetNameZh sets the "name_zh" field.
func (clu *CountryLocationUpdate) SetNameZh(s string) *CountryLocationUpdate {
	clu.mutation.SetNameZh(s)
	return clu
}

// SetNillableNameZh sets the "name_zh" field if the given value is not nil.
func (clu *CountryLocationUpdate) SetNillableNameZh(s *string) *CountryLocationUpdate {
	if s != nil {
		clu.SetNameZh(*s)
	}
	return clu
}

// SetCreatedAt sets the "created_at" field.
func (clu *CountryLocationUpdate) SetCreatedAt(t time.Time) *CountryLocationUpdate {
	clu.mutation.SetCreatedAt(t)
	return clu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (clu *CountryLocationUpdate) SetNillableCreatedAt(t *time.Time) *CountryLocationUpdate {
	if t != nil {
		clu.SetCreatedAt(*t)
	}
	return clu
}

// SetUpdatedAt sets the "updated_at" field.
func (clu *CountryLocationUpdate) SetUpdatedAt(t time.Time) *CountryLocationUpdate {
	clu.mutation.SetUpdatedAt(t)
	return clu
}

// SetDeletedAt sets the "deleted_at" field.
func (clu *CountryLocationUpdate) SetDeletedAt(t time.Time) *CountryLocationUpdate {
	clu.mutation.SetDeletedAt(t)
	return clu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (clu *CountryLocationUpdate) SetNillableDeletedAt(t *time.Time) *CountryLocationUpdate {
	if t != nil {
		clu.SetDeletedAt(*t)
	}
	return clu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (clu *CountryLocationUpdate) ClearDeletedAt() *CountryLocationUpdate {
	clu.mutation.ClearDeletedAt()
	return clu
}

// Mutation returns the CountryLocationMutation object of the builder.
func (clu *CountryLocationUpdate) Mutation() *CountryLocationMutation {
	return clu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (clu *CountryLocationUpdate) Save(ctx context.Context) (int, error) {
	clu.defaults()
	return withHooks(ctx, clu.sqlSave, clu.mutation, clu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (clu *CountryLocationUpdate) SaveX(ctx context.Context) int {
	affected, err := clu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (clu *CountryLocationUpdate) Exec(ctx context.Context) error {
	_, err := clu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (clu *CountryLocationUpdate) ExecX(ctx context.Context) {
	if err := clu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (clu *CountryLocationUpdate) defaults() {
	if _, ok := clu.mutation.UpdatedAt(); !ok {
		v := countrylocation.UpdateDefaultUpdatedAt()
		clu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (clu *CountryLocationUpdate) check() error {
	if v, ok := clu.mutation.NameEn(); ok {
		if err := countrylocation.NameEnValidator(v); err != nil {
			return &ValidationError{Name: "name_en", err: fmt.Errorf(`ent: validator failed for field "CountryLocation.name_en": %w`, err)}
		}
	}
	if v, ok := clu.mutation.NameZh(); ok {
		if err := countrylocation.NameZhValidator(v); err != nil {
			return &ValidationError{Name: "name_zh", err: fmt.Errorf(`ent: validator failed for field "CountryLocation.name_zh": %w`, err)}
		}
	}
	return nil
}

func (clu *CountryLocationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := clu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(countrylocation.Table, countrylocation.Columns, sqlgraph.NewFieldSpec(countrylocation.FieldID, field.TypeUint64))
	if ps := clu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := clu.mutation.CountryID(); ok {
		_spec.SetField(countrylocation.FieldCountryID, field.TypeUint64, value)
	}
	if value, ok := clu.mutation.AddedCountryID(); ok {
		_spec.AddField(countrylocation.FieldCountryID, field.TypeUint64, value)
	}
	if clu.mutation.CountryIDCleared() {
		_spec.ClearField(countrylocation.FieldCountryID, field.TypeUint64)
	}
	if value, ok := clu.mutation.ParentID(); ok {
		_spec.SetField(countrylocation.FieldParentID, field.TypeUint64, value)
	}
	if value, ok := clu.mutation.AddedParentID(); ok {
		_spec.AddField(countrylocation.FieldParentID, field.TypeUint64, value)
	}
	if clu.mutation.ParentIDCleared() {
		_spec.ClearField(countrylocation.FieldParentID, field.TypeUint64)
	}
	if value, ok := clu.mutation.Sorting(); ok {
		_spec.SetField(countrylocation.FieldSorting, field.TypeUint64, value)
	}
	if value, ok := clu.mutation.AddedSorting(); ok {
		_spec.AddField(countrylocation.FieldSorting, field.TypeUint64, value)
	}
	if value, ok := clu.mutation.NameEn(); ok {
		_spec.SetField(countrylocation.FieldNameEn, field.TypeString, value)
	}
	if value, ok := clu.mutation.NameZh(); ok {
		_spec.SetField(countrylocation.FieldNameZh, field.TypeString, value)
	}
	if value, ok := clu.mutation.CreatedAt(); ok {
		_spec.SetField(countrylocation.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := clu.mutation.UpdatedAt(); ok {
		_spec.SetField(countrylocation.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := clu.mutation.DeletedAt(); ok {
		_spec.SetField(countrylocation.FieldDeletedAt, field.TypeTime, value)
	}
	if clu.mutation.DeletedAtCleared() {
		_spec.ClearField(countrylocation.FieldDeletedAt, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, clu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{countrylocation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	clu.mutation.done = true
	return n, nil
}

// CountryLocationUpdateOne is the builder for updating a single CountryLocation entity.
type CountryLocationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CountryLocationMutation
}

// SetCountryID sets the "country_id" field.
func (cluo *CountryLocationUpdateOne) SetCountryID(u uint64) *CountryLocationUpdateOne {
	cluo.mutation.ResetCountryID()
	cluo.mutation.SetCountryID(u)
	return cluo
}

// SetNillableCountryID sets the "country_id" field if the given value is not nil.
func (cluo *CountryLocationUpdateOne) SetNillableCountryID(u *uint64) *CountryLocationUpdateOne {
	if u != nil {
		cluo.SetCountryID(*u)
	}
	return cluo
}

// AddCountryID adds u to the "country_id" field.
func (cluo *CountryLocationUpdateOne) AddCountryID(u int64) *CountryLocationUpdateOne {
	cluo.mutation.AddCountryID(u)
	return cluo
}

// ClearCountryID clears the value of the "country_id" field.
func (cluo *CountryLocationUpdateOne) ClearCountryID() *CountryLocationUpdateOne {
	cluo.mutation.ClearCountryID()
	return cluo
}

// SetParentID sets the "parent_id" field.
func (cluo *CountryLocationUpdateOne) SetParentID(u uint64) *CountryLocationUpdateOne {
	cluo.mutation.ResetParentID()
	cluo.mutation.SetParentID(u)
	return cluo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (cluo *CountryLocationUpdateOne) SetNillableParentID(u *uint64) *CountryLocationUpdateOne {
	if u != nil {
		cluo.SetParentID(*u)
	}
	return cluo
}

// AddParentID adds u to the "parent_id" field.
func (cluo *CountryLocationUpdateOne) AddParentID(u int64) *CountryLocationUpdateOne {
	cluo.mutation.AddParentID(u)
	return cluo
}

// ClearParentID clears the value of the "parent_id" field.
func (cluo *CountryLocationUpdateOne) ClearParentID() *CountryLocationUpdateOne {
	cluo.mutation.ClearParentID()
	return cluo
}

// SetSorting sets the "sorting" field.
func (cluo *CountryLocationUpdateOne) SetSorting(u uint64) *CountryLocationUpdateOne {
	cluo.mutation.ResetSorting()
	cluo.mutation.SetSorting(u)
	return cluo
}

// SetNillableSorting sets the "sorting" field if the given value is not nil.
func (cluo *CountryLocationUpdateOne) SetNillableSorting(u *uint64) *CountryLocationUpdateOne {
	if u != nil {
		cluo.SetSorting(*u)
	}
	return cluo
}

// AddSorting adds u to the "sorting" field.
func (cluo *CountryLocationUpdateOne) AddSorting(u int64) *CountryLocationUpdateOne {
	cluo.mutation.AddSorting(u)
	return cluo
}

// SetNameEn sets the "name_en" field.
func (cluo *CountryLocationUpdateOne) SetNameEn(s string) *CountryLocationUpdateOne {
	cluo.mutation.SetNameEn(s)
	return cluo
}

// SetNillableNameEn sets the "name_en" field if the given value is not nil.
func (cluo *CountryLocationUpdateOne) SetNillableNameEn(s *string) *CountryLocationUpdateOne {
	if s != nil {
		cluo.SetNameEn(*s)
	}
	return cluo
}

// SetNameZh sets the "name_zh" field.
func (cluo *CountryLocationUpdateOne) SetNameZh(s string) *CountryLocationUpdateOne {
	cluo.mutation.SetNameZh(s)
	return cluo
}

// SetNillableNameZh sets the "name_zh" field if the given value is not nil.
func (cluo *CountryLocationUpdateOne) SetNillableNameZh(s *string) *CountryLocationUpdateOne {
	if s != nil {
		cluo.SetNameZh(*s)
	}
	return cluo
}

// SetCreatedAt sets the "created_at" field.
func (cluo *CountryLocationUpdateOne) SetCreatedAt(t time.Time) *CountryLocationUpdateOne {
	cluo.mutation.SetCreatedAt(t)
	return cluo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cluo *CountryLocationUpdateOne) SetNillableCreatedAt(t *time.Time) *CountryLocationUpdateOne {
	if t != nil {
		cluo.SetCreatedAt(*t)
	}
	return cluo
}

// SetUpdatedAt sets the "updated_at" field.
func (cluo *CountryLocationUpdateOne) SetUpdatedAt(t time.Time) *CountryLocationUpdateOne {
	cluo.mutation.SetUpdatedAt(t)
	return cluo
}

// SetDeletedAt sets the "deleted_at" field.
func (cluo *CountryLocationUpdateOne) SetDeletedAt(t time.Time) *CountryLocationUpdateOne {
	cluo.mutation.SetDeletedAt(t)
	return cluo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cluo *CountryLocationUpdateOne) SetNillableDeletedAt(t *time.Time) *CountryLocationUpdateOne {
	if t != nil {
		cluo.SetDeletedAt(*t)
	}
	return cluo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (cluo *CountryLocationUpdateOne) ClearDeletedAt() *CountryLocationUpdateOne {
	cluo.mutation.ClearDeletedAt()
	return cluo
}

// Mutation returns the CountryLocationMutation object of the builder.
func (cluo *CountryLocationUpdateOne) Mutation() *CountryLocationMutation {
	return cluo.mutation
}

// Where appends a list predicates to the CountryLocationUpdate builder.
func (cluo *CountryLocationUpdateOne) Where(ps ...predicate.CountryLocation) *CountryLocationUpdateOne {
	cluo.mutation.Where(ps...)
	return cluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cluo *CountryLocationUpdateOne) Select(field string, fields ...string) *CountryLocationUpdateOne {
	cluo.fields = append([]string{field}, fields...)
	return cluo
}

// Save executes the query and returns the updated CountryLocation entity.
func (cluo *CountryLocationUpdateOne) Save(ctx context.Context) (*CountryLocation, error) {
	cluo.defaults()
	return withHooks(ctx, cluo.sqlSave, cluo.mutation, cluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cluo *CountryLocationUpdateOne) SaveX(ctx context.Context) *CountryLocation {
	node, err := cluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cluo *CountryLocationUpdateOne) Exec(ctx context.Context) error {
	_, err := cluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cluo *CountryLocationUpdateOne) ExecX(ctx context.Context) {
	if err := cluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cluo *CountryLocationUpdateOne) defaults() {
	if _, ok := cluo.mutation.UpdatedAt(); !ok {
		v := countrylocation.UpdateDefaultUpdatedAt()
		cluo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cluo *CountryLocationUpdateOne) check() error {
	if v, ok := cluo.mutation.NameEn(); ok {
		if err := countrylocation.NameEnValidator(v); err != nil {
			return &ValidationError{Name: "name_en", err: fmt.Errorf(`ent: validator failed for field "CountryLocation.name_en": %w`, err)}
		}
	}
	if v, ok := cluo.mutation.NameZh(); ok {
		if err := countrylocation.NameZhValidator(v); err != nil {
			return &ValidationError{Name: "name_zh", err: fmt.Errorf(`ent: validator failed for field "CountryLocation.name_zh": %w`, err)}
		}
	}
	return nil
}

func (cluo *CountryLocationUpdateOne) sqlSave(ctx context.Context) (_node *CountryLocation, err error) {
	if err := cluo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(countrylocation.Table, countrylocation.Columns, sqlgraph.NewFieldSpec(countrylocation.FieldID, field.TypeUint64))
	id, ok := cluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CountryLocation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, countrylocation.FieldID)
		for _, f := range fields {
			if !countrylocation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != countrylocation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cluo.mutation.CountryID(); ok {
		_spec.SetField(countrylocation.FieldCountryID, field.TypeUint64, value)
	}
	if value, ok := cluo.mutation.AddedCountryID(); ok {
		_spec.AddField(countrylocation.FieldCountryID, field.TypeUint64, value)
	}
	if cluo.mutation.CountryIDCleared() {
		_spec.ClearField(countrylocation.FieldCountryID, field.TypeUint64)
	}
	if value, ok := cluo.mutation.ParentID(); ok {
		_spec.SetField(countrylocation.FieldParentID, field.TypeUint64, value)
	}
	if value, ok := cluo.mutation.AddedParentID(); ok {
		_spec.AddField(countrylocation.FieldParentID, field.TypeUint64, value)
	}
	if cluo.mutation.ParentIDCleared() {
		_spec.ClearField(countrylocation.FieldParentID, field.TypeUint64)
	}
	if value, ok := cluo.mutation.Sorting(); ok {
		_spec.SetField(countrylocation.FieldSorting, field.TypeUint64, value)
	}
	if value, ok := cluo.mutation.AddedSorting(); ok {
		_spec.AddField(countrylocation.FieldSorting, field.TypeUint64, value)
	}
	if value, ok := cluo.mutation.NameEn(); ok {
		_spec.SetField(countrylocation.FieldNameEn, field.TypeString, value)
	}
	if value, ok := cluo.mutation.NameZh(); ok {
		_spec.SetField(countrylocation.FieldNameZh, field.TypeString, value)
	}
	if value, ok := cluo.mutation.CreatedAt(); ok {
		_spec.SetField(countrylocation.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := cluo.mutation.UpdatedAt(); ok {
		_spec.SetField(countrylocation.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cluo.mutation.DeletedAt(); ok {
		_spec.SetField(countrylocation.FieldDeletedAt, field.TypeTime, value)
	}
	if cluo.mutation.DeletedAtCleared() {
		_spec.ClearField(countrylocation.FieldDeletedAt, field.TypeTime)
	}
	_node = &CountryLocation{config: cluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{countrylocation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cluo.mutation.done = true
	return _node, nil
}
