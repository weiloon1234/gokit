// Code generated by ent, DO NOT EDIT.

package countrylocation

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the countrylocation type in the database.
	Label = "country_location"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCountryID holds the string denoting the country_id field in the database.
	FieldCountryID = "country_id"
	// FieldParentID holds the string denoting the parent_id field in the database.
	FieldParentID = "parent_id"
	// FieldSorting holds the string denoting the sorting field in the database.
	FieldSorting = "sorting"
	// FieldNameEn holds the string denoting the name_en field in the database.
	FieldNameEn = "name_en"
	// FieldNameZh holds the string denoting the name_zh field in the database.
	FieldNameZh = "name_zh"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeCountry holds the string denoting the country edge name in mutations.
	EdgeCountry = "country"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildLocations holds the string denoting the child_locations edge name in mutations.
	EdgeChildLocations = "child_locations"
	// Table holds the table name of the countrylocation in the database.
	Table = "country_locations"
	// CountryTable is the table that holds the country relation/edge.
	CountryTable = "country_locations"
	// CountryInverseTable is the table name for the Country entity.
	// It exists in this package in order to avoid circular dependency with the "country" package.
	CountryInverseTable = "countries"
	// CountryColumn is the table column denoting the country relation/edge.
	CountryColumn = "country_id"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "country_locations"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "parent_id"
	// ChildLocationsTable is the table that holds the child_locations relation/edge.
	ChildLocationsTable = "country_locations"
	// ChildLocationsColumn is the table column denoting the child_locations relation/edge.
	ChildLocationsColumn = "parent_id"
)

// Columns holds all SQL columns for countrylocation fields.
var Columns = []string{
	FieldID,
	FieldCountryID,
	FieldParentID,
	FieldSorting,
	FieldNameEn,
	FieldNameZh,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultSorting holds the default value on creation for the "sorting" field.
	DefaultSorting uint64
	// NameEnValidator is a validator for the "name_en" field. It is called by the builders before save.
	NameEnValidator func(string) error
	// NameZhValidator is a validator for the "name_zh" field. It is called by the builders before save.
	NameZhValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the CountryLocation queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCountryID orders the results by the country_id field.
func ByCountryID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCountryID, opts...).ToFunc()
}

// ByParentID orders the results by the parent_id field.
func ByParentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParentID, opts...).ToFunc()
}

// BySorting orders the results by the sorting field.
func BySorting(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSorting, opts...).ToFunc()
}

// ByNameEn orders the results by the name_en field.
func ByNameEn(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNameEn, opts...).ToFunc()
}

// ByNameZh orders the results by the name_zh field.
func ByNameZh(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNameZh, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByCountryField orders the results by country field.
func ByCountryField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCountryStep(), sql.OrderByField(field, opts...))
	}
}

// ByParentField orders the results by parent field.
func ByParentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newParentStep(), sql.OrderByField(field, opts...))
	}
}

// ByChildLocationsCount orders the results by child_locations count.
func ByChildLocationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newChildLocationsStep(), opts...)
	}
}

// ByChildLocations orders the results by child_locations terms.
func ByChildLocations(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChildLocationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCountryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CountryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CountryTable, CountryColumn),
	)
}
func newParentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
	)
}
func newChildLocationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ChildLocationsTable, ChildLocationsColumn),
	)
}