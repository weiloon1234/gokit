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
	"github.com/weiloon1234/gokit/ent/country"
	"github.com/weiloon1234/gokit/ent/countrylocation"
	"github.com/weiloon1234/gokit/ent/predicate"
)

// CountryUpdate is the builder for updating Country entities.
type CountryUpdate struct {
	config
	hooks    []Hook
	mutation *CountryMutation
}

// Where appends a list predicates to the CountryUpdate builder.
func (cu *CountryUpdate) Where(ps ...predicate.Country) *CountryUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetIso2 sets the "iso2" field.
func (cu *CountryUpdate) SetIso2(s string) *CountryUpdate {
	cu.mutation.SetIso2(s)
	return cu
}

// SetNillableIso2 sets the "iso2" field if the given value is not nil.
func (cu *CountryUpdate) SetNillableIso2(s *string) *CountryUpdate {
	if s != nil {
		cu.SetIso2(*s)
	}
	return cu
}

// SetIso3 sets the "iso3" field.
func (cu *CountryUpdate) SetIso3(s string) *CountryUpdate {
	cu.mutation.SetIso3(s)
	return cu
}

// SetNillableIso3 sets the "iso3" field if the given value is not nil.
func (cu *CountryUpdate) SetNillableIso3(s *string) *CountryUpdate {
	if s != nil {
		cu.SetIso3(*s)
	}
	return cu
}

// SetName sets the "name" field.
func (cu *CountryUpdate) SetName(s string) *CountryUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cu *CountryUpdate) SetNillableName(s *string) *CountryUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// SetOfficialName sets the "official_name" field.
func (cu *CountryUpdate) SetOfficialName(s string) *CountryUpdate {
	cu.mutation.SetOfficialName(s)
	return cu
}

// SetNillableOfficialName sets the "official_name" field if the given value is not nil.
func (cu *CountryUpdate) SetNillableOfficialName(s *string) *CountryUpdate {
	if s != nil {
		cu.SetOfficialName(*s)
	}
	return cu
}

// ClearOfficialName clears the value of the "official_name" field.
func (cu *CountryUpdate) ClearOfficialName() *CountryUpdate {
	cu.mutation.ClearOfficialName()
	return cu
}

// SetNumericCode sets the "numeric_code" field.
func (cu *CountryUpdate) SetNumericCode(s string) *CountryUpdate {
	cu.mutation.SetNumericCode(s)
	return cu
}

// SetNillableNumericCode sets the "numeric_code" field if the given value is not nil.
func (cu *CountryUpdate) SetNillableNumericCode(s *string) *CountryUpdate {
	if s != nil {
		cu.SetNumericCode(*s)
	}
	return cu
}

// ClearNumericCode clears the value of the "numeric_code" field.
func (cu *CountryUpdate) ClearNumericCode() *CountryUpdate {
	cu.mutation.ClearNumericCode()
	return cu
}

// SetPhoneCode sets the "phone_code" field.
func (cu *CountryUpdate) SetPhoneCode(s string) *CountryUpdate {
	cu.mutation.SetPhoneCode(s)
	return cu
}

// SetNillablePhoneCode sets the "phone_code" field if the given value is not nil.
func (cu *CountryUpdate) SetNillablePhoneCode(s *string) *CountryUpdate {
	if s != nil {
		cu.SetPhoneCode(*s)
	}
	return cu
}

// ClearPhoneCode clears the value of the "phone_code" field.
func (cu *CountryUpdate) ClearPhoneCode() *CountryUpdate {
	cu.mutation.ClearPhoneCode()
	return cu
}

// SetCapital sets the "capital" field.
func (cu *CountryUpdate) SetCapital(s string) *CountryUpdate {
	cu.mutation.SetCapital(s)
	return cu
}

// SetNillableCapital sets the "capital" field if the given value is not nil.
func (cu *CountryUpdate) SetNillableCapital(s *string) *CountryUpdate {
	if s != nil {
		cu.SetCapital(*s)
	}
	return cu
}

// ClearCapital clears the value of the "capital" field.
func (cu *CountryUpdate) ClearCapital() *CountryUpdate {
	cu.mutation.ClearCapital()
	return cu
}

// SetCurrencyName sets the "currency_name" field.
func (cu *CountryUpdate) SetCurrencyName(s string) *CountryUpdate {
	cu.mutation.SetCurrencyName(s)
	return cu
}

// SetNillableCurrencyName sets the "currency_name" field if the given value is not nil.
func (cu *CountryUpdate) SetNillableCurrencyName(s *string) *CountryUpdate {
	if s != nil {
		cu.SetCurrencyName(*s)
	}
	return cu
}

// ClearCurrencyName clears the value of the "currency_name" field.
func (cu *CountryUpdate) ClearCurrencyName() *CountryUpdate {
	cu.mutation.ClearCurrencyName()
	return cu
}

// SetCurrencyCode sets the "currency_code" field.
func (cu *CountryUpdate) SetCurrencyCode(s string) *CountryUpdate {
	cu.mutation.SetCurrencyCode(s)
	return cu
}

// SetNillableCurrencyCode sets the "currency_code" field if the given value is not nil.
func (cu *CountryUpdate) SetNillableCurrencyCode(s *string) *CountryUpdate {
	if s != nil {
		cu.SetCurrencyCode(*s)
	}
	return cu
}

// ClearCurrencyCode clears the value of the "currency_code" field.
func (cu *CountryUpdate) ClearCurrencyCode() *CountryUpdate {
	cu.mutation.ClearCurrencyCode()
	return cu
}

// SetCurrencySymbol sets the "currency_symbol" field.
func (cu *CountryUpdate) SetCurrencySymbol(s string) *CountryUpdate {
	cu.mutation.SetCurrencySymbol(s)
	return cu
}

// SetNillableCurrencySymbol sets the "currency_symbol" field if the given value is not nil.
func (cu *CountryUpdate) SetNillableCurrencySymbol(s *string) *CountryUpdate {
	if s != nil {
		cu.SetCurrencySymbol(*s)
	}
	return cu
}

// ClearCurrencySymbol clears the value of the "currency_symbol" field.
func (cu *CountryUpdate) ClearCurrencySymbol() *CountryUpdate {
	cu.mutation.ClearCurrencySymbol()
	return cu
}

// SetConversionRate sets the "conversion_rate" field.
func (cu *CountryUpdate) SetConversionRate(f float64) *CountryUpdate {
	cu.mutation.ResetConversionRate()
	cu.mutation.SetConversionRate(f)
	return cu
}

// SetNillableConversionRate sets the "conversion_rate" field if the given value is not nil.
func (cu *CountryUpdate) SetNillableConversionRate(f *float64) *CountryUpdate {
	if f != nil {
		cu.SetConversionRate(*f)
	}
	return cu
}

// AddConversionRate adds f to the "conversion_rate" field.
func (cu *CountryUpdate) AddConversionRate(f float64) *CountryUpdate {
	cu.mutation.AddConversionRate(f)
	return cu
}

// SetStatus sets the "status" field.
func (cu *CountryUpdate) SetStatus(u uint8) *CountryUpdate {
	cu.mutation.ResetStatus()
	cu.mutation.SetStatus(u)
	return cu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cu *CountryUpdate) SetNillableStatus(u *uint8) *CountryUpdate {
	if u != nil {
		cu.SetStatus(*u)
	}
	return cu
}

// AddStatus adds u to the "status" field.
func (cu *CountryUpdate) AddStatus(u int8) *CountryUpdate {
	cu.mutation.AddStatus(u)
	return cu
}

// SetCreatedAt sets the "created_at" field.
func (cu *CountryUpdate) SetCreatedAt(t time.Time) *CountryUpdate {
	cu.mutation.SetCreatedAt(t)
	return cu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cu *CountryUpdate) SetNillableCreatedAt(t *time.Time) *CountryUpdate {
	if t != nil {
		cu.SetCreatedAt(*t)
	}
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CountryUpdate) SetUpdatedAt(t time.Time) *CountryUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// AddLocationIDs adds the "locations" edge to the CountryLocation entity by IDs.
func (cu *CountryUpdate) AddLocationIDs(ids ...uint64) *CountryUpdate {
	cu.mutation.AddLocationIDs(ids...)
	return cu
}

// AddLocations adds the "locations" edges to the CountryLocation entity.
func (cu *CountryUpdate) AddLocations(c ...*CountryLocation) *CountryUpdate {
	ids := make([]uint64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddLocationIDs(ids...)
}

// Mutation returns the CountryMutation object of the builder.
func (cu *CountryUpdate) Mutation() *CountryMutation {
	return cu.mutation
}

// ClearLocations clears all "locations" edges to the CountryLocation entity.
func (cu *CountryUpdate) ClearLocations() *CountryUpdate {
	cu.mutation.ClearLocations()
	return cu
}

// RemoveLocationIDs removes the "locations" edge to CountryLocation entities by IDs.
func (cu *CountryUpdate) RemoveLocationIDs(ids ...uint64) *CountryUpdate {
	cu.mutation.RemoveLocationIDs(ids...)
	return cu
}

// RemoveLocations removes "locations" edges to CountryLocation entities.
func (cu *CountryUpdate) RemoveLocations(c ...*CountryLocation) *CountryUpdate {
	ids := make([]uint64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveLocationIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CountryUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CountryUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CountryUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CountryUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CountryUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := country.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CountryUpdate) check() error {
	if v, ok := cu.mutation.Iso2(); ok {
		if err := country.Iso2Validator(v); err != nil {
			return &ValidationError{Name: "iso2", err: fmt.Errorf(`ent: validator failed for field "Country.iso2": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Iso3(); ok {
		if err := country.Iso3Validator(v); err != nil {
			return &ValidationError{Name: "iso3", err: fmt.Errorf(`ent: validator failed for field "Country.iso3": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Name(); ok {
		if err := country.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Country.name": %w`, err)}
		}
	}
	return nil
}

func (cu *CountryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(country.Table, country.Columns, sqlgraph.NewFieldSpec(country.FieldID, field.TypeUint64))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Iso2(); ok {
		_spec.SetField(country.FieldIso2, field.TypeString, value)
	}
	if value, ok := cu.mutation.Iso3(); ok {
		_spec.SetField(country.FieldIso3, field.TypeString, value)
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(country.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.OfficialName(); ok {
		_spec.SetField(country.FieldOfficialName, field.TypeString, value)
	}
	if cu.mutation.OfficialNameCleared() {
		_spec.ClearField(country.FieldOfficialName, field.TypeString)
	}
	if value, ok := cu.mutation.NumericCode(); ok {
		_spec.SetField(country.FieldNumericCode, field.TypeString, value)
	}
	if cu.mutation.NumericCodeCleared() {
		_spec.ClearField(country.FieldNumericCode, field.TypeString)
	}
	if value, ok := cu.mutation.PhoneCode(); ok {
		_spec.SetField(country.FieldPhoneCode, field.TypeString, value)
	}
	if cu.mutation.PhoneCodeCleared() {
		_spec.ClearField(country.FieldPhoneCode, field.TypeString)
	}
	if value, ok := cu.mutation.Capital(); ok {
		_spec.SetField(country.FieldCapital, field.TypeString, value)
	}
	if cu.mutation.CapitalCleared() {
		_spec.ClearField(country.FieldCapital, field.TypeString)
	}
	if value, ok := cu.mutation.CurrencyName(); ok {
		_spec.SetField(country.FieldCurrencyName, field.TypeString, value)
	}
	if cu.mutation.CurrencyNameCleared() {
		_spec.ClearField(country.FieldCurrencyName, field.TypeString)
	}
	if value, ok := cu.mutation.CurrencyCode(); ok {
		_spec.SetField(country.FieldCurrencyCode, field.TypeString, value)
	}
	if cu.mutation.CurrencyCodeCleared() {
		_spec.ClearField(country.FieldCurrencyCode, field.TypeString)
	}
	if value, ok := cu.mutation.CurrencySymbol(); ok {
		_spec.SetField(country.FieldCurrencySymbol, field.TypeString, value)
	}
	if cu.mutation.CurrencySymbolCleared() {
		_spec.ClearField(country.FieldCurrencySymbol, field.TypeString)
	}
	if value, ok := cu.mutation.ConversionRate(); ok {
		_spec.SetField(country.FieldConversionRate, field.TypeFloat64, value)
	}
	if value, ok := cu.mutation.AddedConversionRate(); ok {
		_spec.AddField(country.FieldConversionRate, field.TypeFloat64, value)
	}
	if value, ok := cu.mutation.Status(); ok {
		_spec.SetField(country.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := cu.mutation.AddedStatus(); ok {
		_spec.AddField(country.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := cu.mutation.CreatedAt(); ok {
		_spec.SetField(country.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(country.FieldUpdatedAt, field.TypeTime, value)
	}
	if cu.mutation.LocationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.LocationsTable,
			Columns: []string{country.LocationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(countrylocation.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedLocationsIDs(); len(nodes) > 0 && !cu.mutation.LocationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.LocationsTable,
			Columns: []string{country.LocationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(countrylocation.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.LocationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.LocationsTable,
			Columns: []string{country.LocationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(countrylocation.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{country.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CountryUpdateOne is the builder for updating a single Country entity.
type CountryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CountryMutation
}

// SetIso2 sets the "iso2" field.
func (cuo *CountryUpdateOne) SetIso2(s string) *CountryUpdateOne {
	cuo.mutation.SetIso2(s)
	return cuo
}

// SetNillableIso2 sets the "iso2" field if the given value is not nil.
func (cuo *CountryUpdateOne) SetNillableIso2(s *string) *CountryUpdateOne {
	if s != nil {
		cuo.SetIso2(*s)
	}
	return cuo
}

// SetIso3 sets the "iso3" field.
func (cuo *CountryUpdateOne) SetIso3(s string) *CountryUpdateOne {
	cuo.mutation.SetIso3(s)
	return cuo
}

// SetNillableIso3 sets the "iso3" field if the given value is not nil.
func (cuo *CountryUpdateOne) SetNillableIso3(s *string) *CountryUpdateOne {
	if s != nil {
		cuo.SetIso3(*s)
	}
	return cuo
}

// SetName sets the "name" field.
func (cuo *CountryUpdateOne) SetName(s string) *CountryUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cuo *CountryUpdateOne) SetNillableName(s *string) *CountryUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// SetOfficialName sets the "official_name" field.
func (cuo *CountryUpdateOne) SetOfficialName(s string) *CountryUpdateOne {
	cuo.mutation.SetOfficialName(s)
	return cuo
}

// SetNillableOfficialName sets the "official_name" field if the given value is not nil.
func (cuo *CountryUpdateOne) SetNillableOfficialName(s *string) *CountryUpdateOne {
	if s != nil {
		cuo.SetOfficialName(*s)
	}
	return cuo
}

// ClearOfficialName clears the value of the "official_name" field.
func (cuo *CountryUpdateOne) ClearOfficialName() *CountryUpdateOne {
	cuo.mutation.ClearOfficialName()
	return cuo
}

// SetNumericCode sets the "numeric_code" field.
func (cuo *CountryUpdateOne) SetNumericCode(s string) *CountryUpdateOne {
	cuo.mutation.SetNumericCode(s)
	return cuo
}

// SetNillableNumericCode sets the "numeric_code" field if the given value is not nil.
func (cuo *CountryUpdateOne) SetNillableNumericCode(s *string) *CountryUpdateOne {
	if s != nil {
		cuo.SetNumericCode(*s)
	}
	return cuo
}

// ClearNumericCode clears the value of the "numeric_code" field.
func (cuo *CountryUpdateOne) ClearNumericCode() *CountryUpdateOne {
	cuo.mutation.ClearNumericCode()
	return cuo
}

// SetPhoneCode sets the "phone_code" field.
func (cuo *CountryUpdateOne) SetPhoneCode(s string) *CountryUpdateOne {
	cuo.mutation.SetPhoneCode(s)
	return cuo
}

// SetNillablePhoneCode sets the "phone_code" field if the given value is not nil.
func (cuo *CountryUpdateOne) SetNillablePhoneCode(s *string) *CountryUpdateOne {
	if s != nil {
		cuo.SetPhoneCode(*s)
	}
	return cuo
}

// ClearPhoneCode clears the value of the "phone_code" field.
func (cuo *CountryUpdateOne) ClearPhoneCode() *CountryUpdateOne {
	cuo.mutation.ClearPhoneCode()
	return cuo
}

// SetCapital sets the "capital" field.
func (cuo *CountryUpdateOne) SetCapital(s string) *CountryUpdateOne {
	cuo.mutation.SetCapital(s)
	return cuo
}

// SetNillableCapital sets the "capital" field if the given value is not nil.
func (cuo *CountryUpdateOne) SetNillableCapital(s *string) *CountryUpdateOne {
	if s != nil {
		cuo.SetCapital(*s)
	}
	return cuo
}

// ClearCapital clears the value of the "capital" field.
func (cuo *CountryUpdateOne) ClearCapital() *CountryUpdateOne {
	cuo.mutation.ClearCapital()
	return cuo
}

// SetCurrencyName sets the "currency_name" field.
func (cuo *CountryUpdateOne) SetCurrencyName(s string) *CountryUpdateOne {
	cuo.mutation.SetCurrencyName(s)
	return cuo
}

// SetNillableCurrencyName sets the "currency_name" field if the given value is not nil.
func (cuo *CountryUpdateOne) SetNillableCurrencyName(s *string) *CountryUpdateOne {
	if s != nil {
		cuo.SetCurrencyName(*s)
	}
	return cuo
}

// ClearCurrencyName clears the value of the "currency_name" field.
func (cuo *CountryUpdateOne) ClearCurrencyName() *CountryUpdateOne {
	cuo.mutation.ClearCurrencyName()
	return cuo
}

// SetCurrencyCode sets the "currency_code" field.
func (cuo *CountryUpdateOne) SetCurrencyCode(s string) *CountryUpdateOne {
	cuo.mutation.SetCurrencyCode(s)
	return cuo
}

// SetNillableCurrencyCode sets the "currency_code" field if the given value is not nil.
func (cuo *CountryUpdateOne) SetNillableCurrencyCode(s *string) *CountryUpdateOne {
	if s != nil {
		cuo.SetCurrencyCode(*s)
	}
	return cuo
}

// ClearCurrencyCode clears the value of the "currency_code" field.
func (cuo *CountryUpdateOne) ClearCurrencyCode() *CountryUpdateOne {
	cuo.mutation.ClearCurrencyCode()
	return cuo
}

// SetCurrencySymbol sets the "currency_symbol" field.
func (cuo *CountryUpdateOne) SetCurrencySymbol(s string) *CountryUpdateOne {
	cuo.mutation.SetCurrencySymbol(s)
	return cuo
}

// SetNillableCurrencySymbol sets the "currency_symbol" field if the given value is not nil.
func (cuo *CountryUpdateOne) SetNillableCurrencySymbol(s *string) *CountryUpdateOne {
	if s != nil {
		cuo.SetCurrencySymbol(*s)
	}
	return cuo
}

// ClearCurrencySymbol clears the value of the "currency_symbol" field.
func (cuo *CountryUpdateOne) ClearCurrencySymbol() *CountryUpdateOne {
	cuo.mutation.ClearCurrencySymbol()
	return cuo
}

// SetConversionRate sets the "conversion_rate" field.
func (cuo *CountryUpdateOne) SetConversionRate(f float64) *CountryUpdateOne {
	cuo.mutation.ResetConversionRate()
	cuo.mutation.SetConversionRate(f)
	return cuo
}

// SetNillableConversionRate sets the "conversion_rate" field if the given value is not nil.
func (cuo *CountryUpdateOne) SetNillableConversionRate(f *float64) *CountryUpdateOne {
	if f != nil {
		cuo.SetConversionRate(*f)
	}
	return cuo
}

// AddConversionRate adds f to the "conversion_rate" field.
func (cuo *CountryUpdateOne) AddConversionRate(f float64) *CountryUpdateOne {
	cuo.mutation.AddConversionRate(f)
	return cuo
}

// SetStatus sets the "status" field.
func (cuo *CountryUpdateOne) SetStatus(u uint8) *CountryUpdateOne {
	cuo.mutation.ResetStatus()
	cuo.mutation.SetStatus(u)
	return cuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cuo *CountryUpdateOne) SetNillableStatus(u *uint8) *CountryUpdateOne {
	if u != nil {
		cuo.SetStatus(*u)
	}
	return cuo
}

// AddStatus adds u to the "status" field.
func (cuo *CountryUpdateOne) AddStatus(u int8) *CountryUpdateOne {
	cuo.mutation.AddStatus(u)
	return cuo
}

// SetCreatedAt sets the "created_at" field.
func (cuo *CountryUpdateOne) SetCreatedAt(t time.Time) *CountryUpdateOne {
	cuo.mutation.SetCreatedAt(t)
	return cuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cuo *CountryUpdateOne) SetNillableCreatedAt(t *time.Time) *CountryUpdateOne {
	if t != nil {
		cuo.SetCreatedAt(*t)
	}
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CountryUpdateOne) SetUpdatedAt(t time.Time) *CountryUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// AddLocationIDs adds the "locations" edge to the CountryLocation entity by IDs.
func (cuo *CountryUpdateOne) AddLocationIDs(ids ...uint64) *CountryUpdateOne {
	cuo.mutation.AddLocationIDs(ids...)
	return cuo
}

// AddLocations adds the "locations" edges to the CountryLocation entity.
func (cuo *CountryUpdateOne) AddLocations(c ...*CountryLocation) *CountryUpdateOne {
	ids := make([]uint64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddLocationIDs(ids...)
}

// Mutation returns the CountryMutation object of the builder.
func (cuo *CountryUpdateOne) Mutation() *CountryMutation {
	return cuo.mutation
}

// ClearLocations clears all "locations" edges to the CountryLocation entity.
func (cuo *CountryUpdateOne) ClearLocations() *CountryUpdateOne {
	cuo.mutation.ClearLocations()
	return cuo
}

// RemoveLocationIDs removes the "locations" edge to CountryLocation entities by IDs.
func (cuo *CountryUpdateOne) RemoveLocationIDs(ids ...uint64) *CountryUpdateOne {
	cuo.mutation.RemoveLocationIDs(ids...)
	return cuo
}

// RemoveLocations removes "locations" edges to CountryLocation entities.
func (cuo *CountryUpdateOne) RemoveLocations(c ...*CountryLocation) *CountryUpdateOne {
	ids := make([]uint64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveLocationIDs(ids...)
}

// Where appends a list predicates to the CountryUpdate builder.
func (cuo *CountryUpdateOne) Where(ps ...predicate.Country) *CountryUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CountryUpdateOne) Select(field string, fields ...string) *CountryUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Country entity.
func (cuo *CountryUpdateOne) Save(ctx context.Context) (*Country, error) {
	cuo.defaults()
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CountryUpdateOne) SaveX(ctx context.Context) *Country {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CountryUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CountryUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CountryUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := country.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CountryUpdateOne) check() error {
	if v, ok := cuo.mutation.Iso2(); ok {
		if err := country.Iso2Validator(v); err != nil {
			return &ValidationError{Name: "iso2", err: fmt.Errorf(`ent: validator failed for field "Country.iso2": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Iso3(); ok {
		if err := country.Iso3Validator(v); err != nil {
			return &ValidationError{Name: "iso3", err: fmt.Errorf(`ent: validator failed for field "Country.iso3": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Name(); ok {
		if err := country.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Country.name": %w`, err)}
		}
	}
	return nil
}

func (cuo *CountryUpdateOne) sqlSave(ctx context.Context) (_node *Country, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(country.Table, country.Columns, sqlgraph.NewFieldSpec(country.FieldID, field.TypeUint64))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Country.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, country.FieldID)
		for _, f := range fields {
			if !country.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != country.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Iso2(); ok {
		_spec.SetField(country.FieldIso2, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Iso3(); ok {
		_spec.SetField(country.FieldIso3, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(country.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.OfficialName(); ok {
		_spec.SetField(country.FieldOfficialName, field.TypeString, value)
	}
	if cuo.mutation.OfficialNameCleared() {
		_spec.ClearField(country.FieldOfficialName, field.TypeString)
	}
	if value, ok := cuo.mutation.NumericCode(); ok {
		_spec.SetField(country.FieldNumericCode, field.TypeString, value)
	}
	if cuo.mutation.NumericCodeCleared() {
		_spec.ClearField(country.FieldNumericCode, field.TypeString)
	}
	if value, ok := cuo.mutation.PhoneCode(); ok {
		_spec.SetField(country.FieldPhoneCode, field.TypeString, value)
	}
	if cuo.mutation.PhoneCodeCleared() {
		_spec.ClearField(country.FieldPhoneCode, field.TypeString)
	}
	if value, ok := cuo.mutation.Capital(); ok {
		_spec.SetField(country.FieldCapital, field.TypeString, value)
	}
	if cuo.mutation.CapitalCleared() {
		_spec.ClearField(country.FieldCapital, field.TypeString)
	}
	if value, ok := cuo.mutation.CurrencyName(); ok {
		_spec.SetField(country.FieldCurrencyName, field.TypeString, value)
	}
	if cuo.mutation.CurrencyNameCleared() {
		_spec.ClearField(country.FieldCurrencyName, field.TypeString)
	}
	if value, ok := cuo.mutation.CurrencyCode(); ok {
		_spec.SetField(country.FieldCurrencyCode, field.TypeString, value)
	}
	if cuo.mutation.CurrencyCodeCleared() {
		_spec.ClearField(country.FieldCurrencyCode, field.TypeString)
	}
	if value, ok := cuo.mutation.CurrencySymbol(); ok {
		_spec.SetField(country.FieldCurrencySymbol, field.TypeString, value)
	}
	if cuo.mutation.CurrencySymbolCleared() {
		_spec.ClearField(country.FieldCurrencySymbol, field.TypeString)
	}
	if value, ok := cuo.mutation.ConversionRate(); ok {
		_spec.SetField(country.FieldConversionRate, field.TypeFloat64, value)
	}
	if value, ok := cuo.mutation.AddedConversionRate(); ok {
		_spec.AddField(country.FieldConversionRate, field.TypeFloat64, value)
	}
	if value, ok := cuo.mutation.Status(); ok {
		_spec.SetField(country.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := cuo.mutation.AddedStatus(); ok {
		_spec.AddField(country.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := cuo.mutation.CreatedAt(); ok {
		_spec.SetField(country.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(country.FieldUpdatedAt, field.TypeTime, value)
	}
	if cuo.mutation.LocationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.LocationsTable,
			Columns: []string{country.LocationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(countrylocation.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedLocationsIDs(); len(nodes) > 0 && !cuo.mutation.LocationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.LocationsTable,
			Columns: []string{country.LocationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(countrylocation.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.LocationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.LocationsTable,
			Columns: []string{country.LocationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(countrylocation.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Country{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{country.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
