// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/weiloon1234/gokit/ent/country"
	"github.com/weiloon1234/gokit/ent/countrylocation"
	"github.com/weiloon1234/gokit/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	countryFields := schema.Country{}.Fields()
	_ = countryFields
	// countryDescIso2 is the schema descriptor for iso2 field.
	countryDescIso2 := countryFields[1].Descriptor()
	// country.Iso2Validator is a validator for the "iso2" field. It is called by the builders before save.
	country.Iso2Validator = countryDescIso2.Validators[0].(func(string) error)
	// countryDescIso3 is the schema descriptor for iso3 field.
	countryDescIso3 := countryFields[2].Descriptor()
	// country.Iso3Validator is a validator for the "iso3" field. It is called by the builders before save.
	country.Iso3Validator = countryDescIso3.Validators[0].(func(string) error)
	// countryDescName is the schema descriptor for name field.
	countryDescName := countryFields[3].Descriptor()
	// country.NameValidator is a validator for the "name" field. It is called by the builders before save.
	country.NameValidator = countryDescName.Validators[0].(func(string) error)
	// countryDescConversionRate is the schema descriptor for conversion_rate field.
	countryDescConversionRate := countryFields[11].Descriptor()
	// country.DefaultConversionRate holds the default value on creation for the conversion_rate field.
	country.DefaultConversionRate = countryDescConversionRate.Default.(float64)
	// countryDescStatus is the schema descriptor for status field.
	countryDescStatus := countryFields[12].Descriptor()
	// country.DefaultStatus holds the default value on creation for the status field.
	country.DefaultStatus = countryDescStatus.Default.(uint8)
	// countryDescCreatedAt is the schema descriptor for created_at field.
	countryDescCreatedAt := countryFields[13].Descriptor()
	// country.DefaultCreatedAt holds the default value on creation for the created_at field.
	country.DefaultCreatedAt = countryDescCreatedAt.Default.(func() time.Time)
	// countryDescUpdatedAt is the schema descriptor for updated_at field.
	countryDescUpdatedAt := countryFields[14].Descriptor()
	// country.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	country.DefaultUpdatedAt = countryDescUpdatedAt.Default.(func() time.Time)
	// country.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	country.UpdateDefaultUpdatedAt = countryDescUpdatedAt.UpdateDefault.(func() time.Time)
	countrylocationFields := schema.CountryLocation{}.Fields()
	_ = countrylocationFields
	// countrylocationDescSorting is the schema descriptor for sorting field.
	countrylocationDescSorting := countrylocationFields[3].Descriptor()
	// countrylocation.DefaultSorting holds the default value on creation for the sorting field.
	countrylocation.DefaultSorting = countrylocationDescSorting.Default.(uint64)
	// countrylocationDescNameEn is the schema descriptor for name_en field.
	countrylocationDescNameEn := countrylocationFields[4].Descriptor()
	// countrylocation.NameEnValidator is a validator for the "name_en" field. It is called by the builders before save.
	countrylocation.NameEnValidator = countrylocationDescNameEn.Validators[0].(func(string) error)
	// countrylocationDescNameZh is the schema descriptor for name_zh field.
	countrylocationDescNameZh := countrylocationFields[5].Descriptor()
	// countrylocation.NameZhValidator is a validator for the "name_zh" field. It is called by the builders before save.
	countrylocation.NameZhValidator = countrylocationDescNameZh.Validators[0].(func(string) error)
	// countrylocationDescCreatedAt is the schema descriptor for created_at field.
	countrylocationDescCreatedAt := countrylocationFields[6].Descriptor()
	// countrylocation.DefaultCreatedAt holds the default value on creation for the created_at field.
	countrylocation.DefaultCreatedAt = countrylocationDescCreatedAt.Default.(func() time.Time)
	// countrylocationDescUpdatedAt is the schema descriptor for updated_at field.
	countrylocationDescUpdatedAt := countrylocationFields[7].Descriptor()
	// countrylocation.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	countrylocation.DefaultUpdatedAt = countrylocationDescUpdatedAt.Default.(time.Time)
	// countrylocation.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	countrylocation.UpdateDefaultUpdatedAt = countrylocationDescUpdatedAt.UpdateDefault.(func() time.Time)
}
