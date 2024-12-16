// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CountriesColumns holds the columns for the "countries" table.
	CountriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "iso2", Type: field.TypeString},
		{Name: "iso3", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "official_name", Type: field.TypeString, Nullable: true},
		{Name: "numeric_code", Type: field.TypeString, Nullable: true},
		{Name: "phone_code", Type: field.TypeString, Nullable: true},
		{Name: "capital", Type: field.TypeString, Nullable: true},
		{Name: "currency_name", Type: field.TypeString, Nullable: true},
		{Name: "currency_code", Type: field.TypeString, Nullable: true},
		{Name: "currency_symbol", Type: field.TypeString, Nullable: true},
		{Name: "conversion_rate", Type: field.TypeFloat64, Default: 0},
		{Name: "status", Type: field.TypeUint8, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// CountriesTable holds the schema information for the "countries" table.
	CountriesTable = &schema.Table{
		Name:       "countries",
		Columns:    CountriesColumns,
		PrimaryKey: []*schema.Column{CountriesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "country_iso2",
				Unique:  true,
				Columns: []*schema.Column{CountriesColumns[1]},
			},
			{
				Name:    "country_iso3",
				Unique:  true,
				Columns: []*schema.Column{CountriesColumns[2]},
			},
			{
				Name:    "country_status",
				Unique:  false,
				Columns: []*schema.Column{CountriesColumns[12]},
			},
		},
	}
	// CountryLocationsColumns holds the columns for the "country_locations" table.
	CountryLocationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "country_id", Type: field.TypeUint64, Nullable: true},
		{Name: "parent_id", Type: field.TypeUint64, Nullable: true},
		{Name: "sorting", Type: field.TypeUint64, Default: 0},
		{Name: "name_en", Type: field.TypeString},
		{Name: "name_zh", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// CountryLocationsTable holds the schema information for the "country_locations" table.
	CountryLocationsTable = &schema.Table{
		Name:       "country_locations",
		Columns:    CountryLocationsColumns,
		PrimaryKey: []*schema.Column{CountryLocationsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "countrylocation_country_id",
				Unique:  false,
				Columns: []*schema.Column{CountryLocationsColumns[1]},
			},
			{
				Name:    "countrylocation_parent_id",
				Unique:  false,
				Columns: []*schema.Column{CountryLocationsColumns[2]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CountriesTable,
		CountryLocationsTable,
	}
)

func init() {
}