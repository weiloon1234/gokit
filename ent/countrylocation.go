// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/weiloon1234/gokit/ent/country"
	"github.com/weiloon1234/gokit/ent/countrylocation"
)

// CountryLocation is the model entity for the CountryLocation schema.
type CountryLocation struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Country of the location
	CountryID *uint64 `json:"country_id,omitempty"`
	// If present, it is area; this is state id
	ParentID *uint64 `json:"parent_id,omitempty"`
	// Sorting order
	Sorting uint64 `json:"sorting,omitempty"`
	// Location name in English
	NameEn string `json:"name_en,omitempty"`
	// Location name in Chinese
	NameZh string `json:"name_zh,omitempty"`
	// Record creation timestamp
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Record update timestamp
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Record deletion timestamp
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CountryLocationQuery when eager-loading is set.
	Edges        CountryLocationEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CountryLocationEdges holds the relations/edges for other nodes in the graph.
type CountryLocationEdges struct {
	// Country of the location
	Country *Country `json:"country,omitempty"`
	// Parent location (e.g., state of an area)
	Parent *CountryLocation `json:"parent,omitempty"`
	// Child locations of this parent
	ChildLocations []*CountryLocation `json:"child_locations,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// CountryOrErr returns the Country value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CountryLocationEdges) CountryOrErr() (*Country, error) {
	if e.Country != nil {
		return e.Country, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: country.Label}
	}
	return nil, &NotLoadedError{edge: "country"}
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CountryLocationEdges) ParentOrErr() (*CountryLocation, error) {
	if e.Parent != nil {
		return e.Parent, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: countrylocation.Label}
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildLocationsOrErr returns the ChildLocations value or an error if the edge
// was not loaded in eager-loading.
func (e CountryLocationEdges) ChildLocationsOrErr() ([]*CountryLocation, error) {
	if e.loadedTypes[2] {
		return e.ChildLocations, nil
	}
	return nil, &NotLoadedError{edge: "child_locations"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CountryLocation) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case countrylocation.FieldID, countrylocation.FieldCountryID, countrylocation.FieldParentID, countrylocation.FieldSorting:
			values[i] = new(sql.NullInt64)
		case countrylocation.FieldNameEn, countrylocation.FieldNameZh:
			values[i] = new(sql.NullString)
		case countrylocation.FieldCreatedAt, countrylocation.FieldUpdatedAt, countrylocation.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CountryLocation fields.
func (cl *CountryLocation) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case countrylocation.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cl.ID = uint64(value.Int64)
		case countrylocation.FieldCountryID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field country_id", values[i])
			} else if value.Valid {
				cl.CountryID = new(uint64)
				*cl.CountryID = uint64(value.Int64)
			}
		case countrylocation.FieldParentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				cl.ParentID = new(uint64)
				*cl.ParentID = uint64(value.Int64)
			}
		case countrylocation.FieldSorting:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sorting", values[i])
			} else if value.Valid {
				cl.Sorting = uint64(value.Int64)
			}
		case countrylocation.FieldNameEn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_en", values[i])
			} else if value.Valid {
				cl.NameEn = value.String
			}
		case countrylocation.FieldNameZh:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_zh", values[i])
			} else if value.Valid {
				cl.NameZh = value.String
			}
		case countrylocation.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				cl.CreatedAt = value.Time
			}
		case countrylocation.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				cl.UpdatedAt = value.Time
			}
		case countrylocation.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				cl.DeletedAt = new(time.Time)
				*cl.DeletedAt = value.Time
			}
		default:
			cl.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the CountryLocation.
// This includes values selected through modifiers, order, etc.
func (cl *CountryLocation) Value(name string) (ent.Value, error) {
	return cl.selectValues.Get(name)
}

// QueryCountry queries the "country" edge of the CountryLocation entity.
func (cl *CountryLocation) QueryCountry() *CountryQuery {
	return NewCountryLocationClient(cl.config).QueryCountry(cl)
}

// QueryParent queries the "parent" edge of the CountryLocation entity.
func (cl *CountryLocation) QueryParent() *CountryLocationQuery {
	return NewCountryLocationClient(cl.config).QueryParent(cl)
}

// QueryChildLocations queries the "child_locations" edge of the CountryLocation entity.
func (cl *CountryLocation) QueryChildLocations() *CountryLocationQuery {
	return NewCountryLocationClient(cl.config).QueryChildLocations(cl)
}

// Update returns a builder for updating this CountryLocation.
// Note that you need to call CountryLocation.Unwrap() before calling this method if this CountryLocation
// was returned from a transaction, and the transaction was committed or rolled back.
func (cl *CountryLocation) Update() *CountryLocationUpdateOne {
	return NewCountryLocationClient(cl.config).UpdateOne(cl)
}

// Unwrap unwraps the CountryLocation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cl *CountryLocation) Unwrap() *CountryLocation {
	_tx, ok := cl.config.driver.(*txDriver)
	if !ok {
		panic("ent: CountryLocation is not a transactional entity")
	}
	cl.config.driver = _tx.drv
	return cl
}

// String implements the fmt.Stringer.
func (cl *CountryLocation) String() string {
	var builder strings.Builder
	builder.WriteString("CountryLocation(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cl.ID))
	if v := cl.CountryID; v != nil {
		builder.WriteString("country_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := cl.ParentID; v != nil {
		builder.WriteString("parent_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("sorting=")
	builder.WriteString(fmt.Sprintf("%v", cl.Sorting))
	builder.WriteString(", ")
	builder.WriteString("name_en=")
	builder.WriteString(cl.NameEn)
	builder.WriteString(", ")
	builder.WriteString("name_zh=")
	builder.WriteString(cl.NameZh)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(cl.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(cl.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := cl.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// CountryLocations is a parsable slice of CountryLocation.
type CountryLocations []*CountryLocation