// Code generated by ent, DO NOT EDIT.

package country

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/weiloon1234/gokit/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Country {
	return predicate.Country(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Country {
	return predicate.Country(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.Country {
	return predicate.Country(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.Country {
	return predicate.Country(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Country {
	return predicate.Country(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Country {
	return predicate.Country(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Country {
	return predicate.Country(sql.FieldLTE(FieldID, id))
}

// Iso2 applies equality check predicate on the "iso2" field. It's identical to Iso2EQ.
func Iso2(v string) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldIso2, v))
}

// Iso3 applies equality check predicate on the "iso3" field. It's identical to Iso3EQ.
func Iso3(v string) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldIso3, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldName, v))
}

// OfficialName applies equality check predicate on the "official_name" field. It's identical to OfficialNameEQ.
func OfficialName(v string) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldOfficialName, v))
}

// NumericCode applies equality check predicate on the "numeric_code" field. It's identical to NumericCodeEQ.
func NumericCode(v string) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldNumericCode, v))
}

// PhoneCode applies equality check predicate on the "phone_code" field. It's identical to PhoneCodeEQ.
func PhoneCode(v string) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldPhoneCode, v))
}

// Capital applies equality check predicate on the "capital" field. It's identical to CapitalEQ.
func Capital(v string) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldCapital, v))
}

// CurrencyName applies equality check predicate on the "currency_name" field. It's identical to CurrencyNameEQ.
func CurrencyName(v string) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldCurrencyName, v))
}

// CurrencyCode applies equality check predicate on the "currency_code" field. It's identical to CurrencyCodeEQ.
func CurrencyCode(v string) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldCurrencyCode, v))
}

// CurrencySymbol applies equality check predicate on the "currency_symbol" field. It's identical to CurrencySymbolEQ.
func CurrencySymbol(v string) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldCurrencySymbol, v))
}

// ConversionRate applies equality check predicate on the "conversion_rate" field. It's identical to ConversionRateEQ.
func ConversionRate(v float64) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldConversionRate, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v uint8) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldStatus, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldUpdatedAt, v))
}

// Iso2EQ applies the EQ predicate on the "iso2" field.
func Iso2EQ(v string) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldIso2, v))
}

// Iso2NEQ applies the NEQ predicate on the "iso2" field.
func Iso2NEQ(v string) predicate.Country {
	return predicate.Country(sql.FieldNEQ(FieldIso2, v))
}

// Iso2In applies the In predicate on the "iso2" field.
func Iso2In(vs ...string) predicate.Country {
	return predicate.Country(sql.FieldIn(FieldIso2, vs...))
}

// Iso2NotIn applies the NotIn predicate on the "iso2" field.
func Iso2NotIn(vs ...string) predicate.Country {
	return predicate.Country(sql.FieldNotIn(FieldIso2, vs...))
}

// Iso2GT applies the GT predicate on the "iso2" field.
func Iso2GT(v string) predicate.Country {
	return predicate.Country(sql.FieldGT(FieldIso2, v))
}

// Iso2GTE applies the GTE predicate on the "iso2" field.
func Iso2GTE(v string) predicate.Country {
	return predicate.Country(sql.FieldGTE(FieldIso2, v))
}

// Iso2LT applies the LT predicate on the "iso2" field.
func Iso2LT(v string) predicate.Country {
	return predicate.Country(sql.FieldLT(FieldIso2, v))
}

// Iso2LTE applies the LTE predicate on the "iso2" field.
func Iso2LTE(v string) predicate.Country {
	return predicate.Country(sql.FieldLTE(FieldIso2, v))
}

// Iso2Contains applies the Contains predicate on the "iso2" field.
func Iso2Contains(v string) predicate.Country {
	return predicate.Country(sql.FieldContains(FieldIso2, v))
}

// Iso2HasPrefix applies the HasPrefix predicate on the "iso2" field.
func Iso2HasPrefix(v string) predicate.Country {
	return predicate.Country(sql.FieldHasPrefix(FieldIso2, v))
}

// Iso2HasSuffix applies the HasSuffix predicate on the "iso2" field.
func Iso2HasSuffix(v string) predicate.Country {
	return predicate.Country(sql.FieldHasSuffix(FieldIso2, v))
}

// Iso2EqualFold applies the EqualFold predicate on the "iso2" field.
func Iso2EqualFold(v string) predicate.Country {
	return predicate.Country(sql.FieldEqualFold(FieldIso2, v))
}

// Iso2ContainsFold applies the ContainsFold predicate on the "iso2" field.
func Iso2ContainsFold(v string) predicate.Country {
	return predicate.Country(sql.FieldContainsFold(FieldIso2, v))
}

// Iso3EQ applies the EQ predicate on the "iso3" field.
func Iso3EQ(v string) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldIso3, v))
}

// Iso3NEQ applies the NEQ predicate on the "iso3" field.
func Iso3NEQ(v string) predicate.Country {
	return predicate.Country(sql.FieldNEQ(FieldIso3, v))
}

// Iso3In applies the In predicate on the "iso3" field.
func Iso3In(vs ...string) predicate.Country {
	return predicate.Country(sql.FieldIn(FieldIso3, vs...))
}

// Iso3NotIn applies the NotIn predicate on the "iso3" field.
func Iso3NotIn(vs ...string) predicate.Country {
	return predicate.Country(sql.FieldNotIn(FieldIso3, vs...))
}

// Iso3GT applies the GT predicate on the "iso3" field.
func Iso3GT(v string) predicate.Country {
	return predicate.Country(sql.FieldGT(FieldIso3, v))
}

// Iso3GTE applies the GTE predicate on the "iso3" field.
func Iso3GTE(v string) predicate.Country {
	return predicate.Country(sql.FieldGTE(FieldIso3, v))
}

// Iso3LT applies the LT predicate on the "iso3" field.
func Iso3LT(v string) predicate.Country {
	return predicate.Country(sql.FieldLT(FieldIso3, v))
}

// Iso3LTE applies the LTE predicate on the "iso3" field.
func Iso3LTE(v string) predicate.Country {
	return predicate.Country(sql.FieldLTE(FieldIso3, v))
}

// Iso3Contains applies the Contains predicate on the "iso3" field.
func Iso3Contains(v string) predicate.Country {
	return predicate.Country(sql.FieldContains(FieldIso3, v))
}

// Iso3HasPrefix applies the HasPrefix predicate on the "iso3" field.
func Iso3HasPrefix(v string) predicate.Country {
	return predicate.Country(sql.FieldHasPrefix(FieldIso3, v))
}

// Iso3HasSuffix applies the HasSuffix predicate on the "iso3" field.
func Iso3HasSuffix(v string) predicate.Country {
	return predicate.Country(sql.FieldHasSuffix(FieldIso3, v))
}

// Iso3EqualFold applies the EqualFold predicate on the "iso3" field.
func Iso3EqualFold(v string) predicate.Country {
	return predicate.Country(sql.FieldEqualFold(FieldIso3, v))
}

// Iso3ContainsFold applies the ContainsFold predicate on the "iso3" field.
func Iso3ContainsFold(v string) predicate.Country {
	return predicate.Country(sql.FieldContainsFold(FieldIso3, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Country {
	return predicate.Country(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Country {
	return predicate.Country(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Country {
	return predicate.Country(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Country {
	return predicate.Country(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Country {
	return predicate.Country(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Country {
	return predicate.Country(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Country {
	return predicate.Country(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Country {
	return predicate.Country(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Country {
	return predicate.Country(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Country {
	return predicate.Country(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Country {
	return predicate.Country(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Country {
	return predicate.Country(sql.FieldContainsFold(FieldName, v))
}

// OfficialNameEQ applies the EQ predicate on the "official_name" field.
func OfficialNameEQ(v string) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldOfficialName, v))
}

// OfficialNameNEQ applies the NEQ predicate on the "official_name" field.
func OfficialNameNEQ(v string) predicate.Country {
	return predicate.Country(sql.FieldNEQ(FieldOfficialName, v))
}

// OfficialNameIn applies the In predicate on the "official_name" field.
func OfficialNameIn(vs ...string) predicate.Country {
	return predicate.Country(sql.FieldIn(FieldOfficialName, vs...))
}

// OfficialNameNotIn applies the NotIn predicate on the "official_name" field.
func OfficialNameNotIn(vs ...string) predicate.Country {
	return predicate.Country(sql.FieldNotIn(FieldOfficialName, vs...))
}

// OfficialNameGT applies the GT predicate on the "official_name" field.
func OfficialNameGT(v string) predicate.Country {
	return predicate.Country(sql.FieldGT(FieldOfficialName, v))
}

// OfficialNameGTE applies the GTE predicate on the "official_name" field.
func OfficialNameGTE(v string) predicate.Country {
	return predicate.Country(sql.FieldGTE(FieldOfficialName, v))
}

// OfficialNameLT applies the LT predicate on the "official_name" field.
func OfficialNameLT(v string) predicate.Country {
	return predicate.Country(sql.FieldLT(FieldOfficialName, v))
}

// OfficialNameLTE applies the LTE predicate on the "official_name" field.
func OfficialNameLTE(v string) predicate.Country {
	return predicate.Country(sql.FieldLTE(FieldOfficialName, v))
}

// OfficialNameContains applies the Contains predicate on the "official_name" field.
func OfficialNameContains(v string) predicate.Country {
	return predicate.Country(sql.FieldContains(FieldOfficialName, v))
}

// OfficialNameHasPrefix applies the HasPrefix predicate on the "official_name" field.
func OfficialNameHasPrefix(v string) predicate.Country {
	return predicate.Country(sql.FieldHasPrefix(FieldOfficialName, v))
}

// OfficialNameHasSuffix applies the HasSuffix predicate on the "official_name" field.
func OfficialNameHasSuffix(v string) predicate.Country {
	return predicate.Country(sql.FieldHasSuffix(FieldOfficialName, v))
}

// OfficialNameIsNil applies the IsNil predicate on the "official_name" field.
func OfficialNameIsNil() predicate.Country {
	return predicate.Country(sql.FieldIsNull(FieldOfficialName))
}

// OfficialNameNotNil applies the NotNil predicate on the "official_name" field.
func OfficialNameNotNil() predicate.Country {
	return predicate.Country(sql.FieldNotNull(FieldOfficialName))
}

// OfficialNameEqualFold applies the EqualFold predicate on the "official_name" field.
func OfficialNameEqualFold(v string) predicate.Country {
	return predicate.Country(sql.FieldEqualFold(FieldOfficialName, v))
}

// OfficialNameContainsFold applies the ContainsFold predicate on the "official_name" field.
func OfficialNameContainsFold(v string) predicate.Country {
	return predicate.Country(sql.FieldContainsFold(FieldOfficialName, v))
}

// NumericCodeEQ applies the EQ predicate on the "numeric_code" field.
func NumericCodeEQ(v string) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldNumericCode, v))
}

// NumericCodeNEQ applies the NEQ predicate on the "numeric_code" field.
func NumericCodeNEQ(v string) predicate.Country {
	return predicate.Country(sql.FieldNEQ(FieldNumericCode, v))
}

// NumericCodeIn applies the In predicate on the "numeric_code" field.
func NumericCodeIn(vs ...string) predicate.Country {
	return predicate.Country(sql.FieldIn(FieldNumericCode, vs...))
}

// NumericCodeNotIn applies the NotIn predicate on the "numeric_code" field.
func NumericCodeNotIn(vs ...string) predicate.Country {
	return predicate.Country(sql.FieldNotIn(FieldNumericCode, vs...))
}

// NumericCodeGT applies the GT predicate on the "numeric_code" field.
func NumericCodeGT(v string) predicate.Country {
	return predicate.Country(sql.FieldGT(FieldNumericCode, v))
}

// NumericCodeGTE applies the GTE predicate on the "numeric_code" field.
func NumericCodeGTE(v string) predicate.Country {
	return predicate.Country(sql.FieldGTE(FieldNumericCode, v))
}

// NumericCodeLT applies the LT predicate on the "numeric_code" field.
func NumericCodeLT(v string) predicate.Country {
	return predicate.Country(sql.FieldLT(FieldNumericCode, v))
}

// NumericCodeLTE applies the LTE predicate on the "numeric_code" field.
func NumericCodeLTE(v string) predicate.Country {
	return predicate.Country(sql.FieldLTE(FieldNumericCode, v))
}

// NumericCodeContains applies the Contains predicate on the "numeric_code" field.
func NumericCodeContains(v string) predicate.Country {
	return predicate.Country(sql.FieldContains(FieldNumericCode, v))
}

// NumericCodeHasPrefix applies the HasPrefix predicate on the "numeric_code" field.
func NumericCodeHasPrefix(v string) predicate.Country {
	return predicate.Country(sql.FieldHasPrefix(FieldNumericCode, v))
}

// NumericCodeHasSuffix applies the HasSuffix predicate on the "numeric_code" field.
func NumericCodeHasSuffix(v string) predicate.Country {
	return predicate.Country(sql.FieldHasSuffix(FieldNumericCode, v))
}

// NumericCodeIsNil applies the IsNil predicate on the "numeric_code" field.
func NumericCodeIsNil() predicate.Country {
	return predicate.Country(sql.FieldIsNull(FieldNumericCode))
}

// NumericCodeNotNil applies the NotNil predicate on the "numeric_code" field.
func NumericCodeNotNil() predicate.Country {
	return predicate.Country(sql.FieldNotNull(FieldNumericCode))
}

// NumericCodeEqualFold applies the EqualFold predicate on the "numeric_code" field.
func NumericCodeEqualFold(v string) predicate.Country {
	return predicate.Country(sql.FieldEqualFold(FieldNumericCode, v))
}

// NumericCodeContainsFold applies the ContainsFold predicate on the "numeric_code" field.
func NumericCodeContainsFold(v string) predicate.Country {
	return predicate.Country(sql.FieldContainsFold(FieldNumericCode, v))
}

// PhoneCodeEQ applies the EQ predicate on the "phone_code" field.
func PhoneCodeEQ(v string) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldPhoneCode, v))
}

// PhoneCodeNEQ applies the NEQ predicate on the "phone_code" field.
func PhoneCodeNEQ(v string) predicate.Country {
	return predicate.Country(sql.FieldNEQ(FieldPhoneCode, v))
}

// PhoneCodeIn applies the In predicate on the "phone_code" field.
func PhoneCodeIn(vs ...string) predicate.Country {
	return predicate.Country(sql.FieldIn(FieldPhoneCode, vs...))
}

// PhoneCodeNotIn applies the NotIn predicate on the "phone_code" field.
func PhoneCodeNotIn(vs ...string) predicate.Country {
	return predicate.Country(sql.FieldNotIn(FieldPhoneCode, vs...))
}

// PhoneCodeGT applies the GT predicate on the "phone_code" field.
func PhoneCodeGT(v string) predicate.Country {
	return predicate.Country(sql.FieldGT(FieldPhoneCode, v))
}

// PhoneCodeGTE applies the GTE predicate on the "phone_code" field.
func PhoneCodeGTE(v string) predicate.Country {
	return predicate.Country(sql.FieldGTE(FieldPhoneCode, v))
}

// PhoneCodeLT applies the LT predicate on the "phone_code" field.
func PhoneCodeLT(v string) predicate.Country {
	return predicate.Country(sql.FieldLT(FieldPhoneCode, v))
}

// PhoneCodeLTE applies the LTE predicate on the "phone_code" field.
func PhoneCodeLTE(v string) predicate.Country {
	return predicate.Country(sql.FieldLTE(FieldPhoneCode, v))
}

// PhoneCodeContains applies the Contains predicate on the "phone_code" field.
func PhoneCodeContains(v string) predicate.Country {
	return predicate.Country(sql.FieldContains(FieldPhoneCode, v))
}

// PhoneCodeHasPrefix applies the HasPrefix predicate on the "phone_code" field.
func PhoneCodeHasPrefix(v string) predicate.Country {
	return predicate.Country(sql.FieldHasPrefix(FieldPhoneCode, v))
}

// PhoneCodeHasSuffix applies the HasSuffix predicate on the "phone_code" field.
func PhoneCodeHasSuffix(v string) predicate.Country {
	return predicate.Country(sql.FieldHasSuffix(FieldPhoneCode, v))
}

// PhoneCodeIsNil applies the IsNil predicate on the "phone_code" field.
func PhoneCodeIsNil() predicate.Country {
	return predicate.Country(sql.FieldIsNull(FieldPhoneCode))
}

// PhoneCodeNotNil applies the NotNil predicate on the "phone_code" field.
func PhoneCodeNotNil() predicate.Country {
	return predicate.Country(sql.FieldNotNull(FieldPhoneCode))
}

// PhoneCodeEqualFold applies the EqualFold predicate on the "phone_code" field.
func PhoneCodeEqualFold(v string) predicate.Country {
	return predicate.Country(sql.FieldEqualFold(FieldPhoneCode, v))
}

// PhoneCodeContainsFold applies the ContainsFold predicate on the "phone_code" field.
func PhoneCodeContainsFold(v string) predicate.Country {
	return predicate.Country(sql.FieldContainsFold(FieldPhoneCode, v))
}

// CapitalEQ applies the EQ predicate on the "capital" field.
func CapitalEQ(v string) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldCapital, v))
}

// CapitalNEQ applies the NEQ predicate on the "capital" field.
func CapitalNEQ(v string) predicate.Country {
	return predicate.Country(sql.FieldNEQ(FieldCapital, v))
}

// CapitalIn applies the In predicate on the "capital" field.
func CapitalIn(vs ...string) predicate.Country {
	return predicate.Country(sql.FieldIn(FieldCapital, vs...))
}

// CapitalNotIn applies the NotIn predicate on the "capital" field.
func CapitalNotIn(vs ...string) predicate.Country {
	return predicate.Country(sql.FieldNotIn(FieldCapital, vs...))
}

// CapitalGT applies the GT predicate on the "capital" field.
func CapitalGT(v string) predicate.Country {
	return predicate.Country(sql.FieldGT(FieldCapital, v))
}

// CapitalGTE applies the GTE predicate on the "capital" field.
func CapitalGTE(v string) predicate.Country {
	return predicate.Country(sql.FieldGTE(FieldCapital, v))
}

// CapitalLT applies the LT predicate on the "capital" field.
func CapitalLT(v string) predicate.Country {
	return predicate.Country(sql.FieldLT(FieldCapital, v))
}

// CapitalLTE applies the LTE predicate on the "capital" field.
func CapitalLTE(v string) predicate.Country {
	return predicate.Country(sql.FieldLTE(FieldCapital, v))
}

// CapitalContains applies the Contains predicate on the "capital" field.
func CapitalContains(v string) predicate.Country {
	return predicate.Country(sql.FieldContains(FieldCapital, v))
}

// CapitalHasPrefix applies the HasPrefix predicate on the "capital" field.
func CapitalHasPrefix(v string) predicate.Country {
	return predicate.Country(sql.FieldHasPrefix(FieldCapital, v))
}

// CapitalHasSuffix applies the HasSuffix predicate on the "capital" field.
func CapitalHasSuffix(v string) predicate.Country {
	return predicate.Country(sql.FieldHasSuffix(FieldCapital, v))
}

// CapitalIsNil applies the IsNil predicate on the "capital" field.
func CapitalIsNil() predicate.Country {
	return predicate.Country(sql.FieldIsNull(FieldCapital))
}

// CapitalNotNil applies the NotNil predicate on the "capital" field.
func CapitalNotNil() predicate.Country {
	return predicate.Country(sql.FieldNotNull(FieldCapital))
}

// CapitalEqualFold applies the EqualFold predicate on the "capital" field.
func CapitalEqualFold(v string) predicate.Country {
	return predicate.Country(sql.FieldEqualFold(FieldCapital, v))
}

// CapitalContainsFold applies the ContainsFold predicate on the "capital" field.
func CapitalContainsFold(v string) predicate.Country {
	return predicate.Country(sql.FieldContainsFold(FieldCapital, v))
}

// CurrencyNameEQ applies the EQ predicate on the "currency_name" field.
func CurrencyNameEQ(v string) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldCurrencyName, v))
}

// CurrencyNameNEQ applies the NEQ predicate on the "currency_name" field.
func CurrencyNameNEQ(v string) predicate.Country {
	return predicate.Country(sql.FieldNEQ(FieldCurrencyName, v))
}

// CurrencyNameIn applies the In predicate on the "currency_name" field.
func CurrencyNameIn(vs ...string) predicate.Country {
	return predicate.Country(sql.FieldIn(FieldCurrencyName, vs...))
}

// CurrencyNameNotIn applies the NotIn predicate on the "currency_name" field.
func CurrencyNameNotIn(vs ...string) predicate.Country {
	return predicate.Country(sql.FieldNotIn(FieldCurrencyName, vs...))
}

// CurrencyNameGT applies the GT predicate on the "currency_name" field.
func CurrencyNameGT(v string) predicate.Country {
	return predicate.Country(sql.FieldGT(FieldCurrencyName, v))
}

// CurrencyNameGTE applies the GTE predicate on the "currency_name" field.
func CurrencyNameGTE(v string) predicate.Country {
	return predicate.Country(sql.FieldGTE(FieldCurrencyName, v))
}

// CurrencyNameLT applies the LT predicate on the "currency_name" field.
func CurrencyNameLT(v string) predicate.Country {
	return predicate.Country(sql.FieldLT(FieldCurrencyName, v))
}

// CurrencyNameLTE applies the LTE predicate on the "currency_name" field.
func CurrencyNameLTE(v string) predicate.Country {
	return predicate.Country(sql.FieldLTE(FieldCurrencyName, v))
}

// CurrencyNameContains applies the Contains predicate on the "currency_name" field.
func CurrencyNameContains(v string) predicate.Country {
	return predicate.Country(sql.FieldContains(FieldCurrencyName, v))
}

// CurrencyNameHasPrefix applies the HasPrefix predicate on the "currency_name" field.
func CurrencyNameHasPrefix(v string) predicate.Country {
	return predicate.Country(sql.FieldHasPrefix(FieldCurrencyName, v))
}

// CurrencyNameHasSuffix applies the HasSuffix predicate on the "currency_name" field.
func CurrencyNameHasSuffix(v string) predicate.Country {
	return predicate.Country(sql.FieldHasSuffix(FieldCurrencyName, v))
}

// CurrencyNameIsNil applies the IsNil predicate on the "currency_name" field.
func CurrencyNameIsNil() predicate.Country {
	return predicate.Country(sql.FieldIsNull(FieldCurrencyName))
}

// CurrencyNameNotNil applies the NotNil predicate on the "currency_name" field.
func CurrencyNameNotNil() predicate.Country {
	return predicate.Country(sql.FieldNotNull(FieldCurrencyName))
}

// CurrencyNameEqualFold applies the EqualFold predicate on the "currency_name" field.
func CurrencyNameEqualFold(v string) predicate.Country {
	return predicate.Country(sql.FieldEqualFold(FieldCurrencyName, v))
}

// CurrencyNameContainsFold applies the ContainsFold predicate on the "currency_name" field.
func CurrencyNameContainsFold(v string) predicate.Country {
	return predicate.Country(sql.FieldContainsFold(FieldCurrencyName, v))
}

// CurrencyCodeEQ applies the EQ predicate on the "currency_code" field.
func CurrencyCodeEQ(v string) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldCurrencyCode, v))
}

// CurrencyCodeNEQ applies the NEQ predicate on the "currency_code" field.
func CurrencyCodeNEQ(v string) predicate.Country {
	return predicate.Country(sql.FieldNEQ(FieldCurrencyCode, v))
}

// CurrencyCodeIn applies the In predicate on the "currency_code" field.
func CurrencyCodeIn(vs ...string) predicate.Country {
	return predicate.Country(sql.FieldIn(FieldCurrencyCode, vs...))
}

// CurrencyCodeNotIn applies the NotIn predicate on the "currency_code" field.
func CurrencyCodeNotIn(vs ...string) predicate.Country {
	return predicate.Country(sql.FieldNotIn(FieldCurrencyCode, vs...))
}

// CurrencyCodeGT applies the GT predicate on the "currency_code" field.
func CurrencyCodeGT(v string) predicate.Country {
	return predicate.Country(sql.FieldGT(FieldCurrencyCode, v))
}

// CurrencyCodeGTE applies the GTE predicate on the "currency_code" field.
func CurrencyCodeGTE(v string) predicate.Country {
	return predicate.Country(sql.FieldGTE(FieldCurrencyCode, v))
}

// CurrencyCodeLT applies the LT predicate on the "currency_code" field.
func CurrencyCodeLT(v string) predicate.Country {
	return predicate.Country(sql.FieldLT(FieldCurrencyCode, v))
}

// CurrencyCodeLTE applies the LTE predicate on the "currency_code" field.
func CurrencyCodeLTE(v string) predicate.Country {
	return predicate.Country(sql.FieldLTE(FieldCurrencyCode, v))
}

// CurrencyCodeContains applies the Contains predicate on the "currency_code" field.
func CurrencyCodeContains(v string) predicate.Country {
	return predicate.Country(sql.FieldContains(FieldCurrencyCode, v))
}

// CurrencyCodeHasPrefix applies the HasPrefix predicate on the "currency_code" field.
func CurrencyCodeHasPrefix(v string) predicate.Country {
	return predicate.Country(sql.FieldHasPrefix(FieldCurrencyCode, v))
}

// CurrencyCodeHasSuffix applies the HasSuffix predicate on the "currency_code" field.
func CurrencyCodeHasSuffix(v string) predicate.Country {
	return predicate.Country(sql.FieldHasSuffix(FieldCurrencyCode, v))
}

// CurrencyCodeIsNil applies the IsNil predicate on the "currency_code" field.
func CurrencyCodeIsNil() predicate.Country {
	return predicate.Country(sql.FieldIsNull(FieldCurrencyCode))
}

// CurrencyCodeNotNil applies the NotNil predicate on the "currency_code" field.
func CurrencyCodeNotNil() predicate.Country {
	return predicate.Country(sql.FieldNotNull(FieldCurrencyCode))
}

// CurrencyCodeEqualFold applies the EqualFold predicate on the "currency_code" field.
func CurrencyCodeEqualFold(v string) predicate.Country {
	return predicate.Country(sql.FieldEqualFold(FieldCurrencyCode, v))
}

// CurrencyCodeContainsFold applies the ContainsFold predicate on the "currency_code" field.
func CurrencyCodeContainsFold(v string) predicate.Country {
	return predicate.Country(sql.FieldContainsFold(FieldCurrencyCode, v))
}

// CurrencySymbolEQ applies the EQ predicate on the "currency_symbol" field.
func CurrencySymbolEQ(v string) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldCurrencySymbol, v))
}

// CurrencySymbolNEQ applies the NEQ predicate on the "currency_symbol" field.
func CurrencySymbolNEQ(v string) predicate.Country {
	return predicate.Country(sql.FieldNEQ(FieldCurrencySymbol, v))
}

// CurrencySymbolIn applies the In predicate on the "currency_symbol" field.
func CurrencySymbolIn(vs ...string) predicate.Country {
	return predicate.Country(sql.FieldIn(FieldCurrencySymbol, vs...))
}

// CurrencySymbolNotIn applies the NotIn predicate on the "currency_symbol" field.
func CurrencySymbolNotIn(vs ...string) predicate.Country {
	return predicate.Country(sql.FieldNotIn(FieldCurrencySymbol, vs...))
}

// CurrencySymbolGT applies the GT predicate on the "currency_symbol" field.
func CurrencySymbolGT(v string) predicate.Country {
	return predicate.Country(sql.FieldGT(FieldCurrencySymbol, v))
}

// CurrencySymbolGTE applies the GTE predicate on the "currency_symbol" field.
func CurrencySymbolGTE(v string) predicate.Country {
	return predicate.Country(sql.FieldGTE(FieldCurrencySymbol, v))
}

// CurrencySymbolLT applies the LT predicate on the "currency_symbol" field.
func CurrencySymbolLT(v string) predicate.Country {
	return predicate.Country(sql.FieldLT(FieldCurrencySymbol, v))
}

// CurrencySymbolLTE applies the LTE predicate on the "currency_symbol" field.
func CurrencySymbolLTE(v string) predicate.Country {
	return predicate.Country(sql.FieldLTE(FieldCurrencySymbol, v))
}

// CurrencySymbolContains applies the Contains predicate on the "currency_symbol" field.
func CurrencySymbolContains(v string) predicate.Country {
	return predicate.Country(sql.FieldContains(FieldCurrencySymbol, v))
}

// CurrencySymbolHasPrefix applies the HasPrefix predicate on the "currency_symbol" field.
func CurrencySymbolHasPrefix(v string) predicate.Country {
	return predicate.Country(sql.FieldHasPrefix(FieldCurrencySymbol, v))
}

// CurrencySymbolHasSuffix applies the HasSuffix predicate on the "currency_symbol" field.
func CurrencySymbolHasSuffix(v string) predicate.Country {
	return predicate.Country(sql.FieldHasSuffix(FieldCurrencySymbol, v))
}

// CurrencySymbolIsNil applies the IsNil predicate on the "currency_symbol" field.
func CurrencySymbolIsNil() predicate.Country {
	return predicate.Country(sql.FieldIsNull(FieldCurrencySymbol))
}

// CurrencySymbolNotNil applies the NotNil predicate on the "currency_symbol" field.
func CurrencySymbolNotNil() predicate.Country {
	return predicate.Country(sql.FieldNotNull(FieldCurrencySymbol))
}

// CurrencySymbolEqualFold applies the EqualFold predicate on the "currency_symbol" field.
func CurrencySymbolEqualFold(v string) predicate.Country {
	return predicate.Country(sql.FieldEqualFold(FieldCurrencySymbol, v))
}

// CurrencySymbolContainsFold applies the ContainsFold predicate on the "currency_symbol" field.
func CurrencySymbolContainsFold(v string) predicate.Country {
	return predicate.Country(sql.FieldContainsFold(FieldCurrencySymbol, v))
}

// ConversionRateEQ applies the EQ predicate on the "conversion_rate" field.
func ConversionRateEQ(v float64) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldConversionRate, v))
}

// ConversionRateNEQ applies the NEQ predicate on the "conversion_rate" field.
func ConversionRateNEQ(v float64) predicate.Country {
	return predicate.Country(sql.FieldNEQ(FieldConversionRate, v))
}

// ConversionRateIn applies the In predicate on the "conversion_rate" field.
func ConversionRateIn(vs ...float64) predicate.Country {
	return predicate.Country(sql.FieldIn(FieldConversionRate, vs...))
}

// ConversionRateNotIn applies the NotIn predicate on the "conversion_rate" field.
func ConversionRateNotIn(vs ...float64) predicate.Country {
	return predicate.Country(sql.FieldNotIn(FieldConversionRate, vs...))
}

// ConversionRateGT applies the GT predicate on the "conversion_rate" field.
func ConversionRateGT(v float64) predicate.Country {
	return predicate.Country(sql.FieldGT(FieldConversionRate, v))
}

// ConversionRateGTE applies the GTE predicate on the "conversion_rate" field.
func ConversionRateGTE(v float64) predicate.Country {
	return predicate.Country(sql.FieldGTE(FieldConversionRate, v))
}

// ConversionRateLT applies the LT predicate on the "conversion_rate" field.
func ConversionRateLT(v float64) predicate.Country {
	return predicate.Country(sql.FieldLT(FieldConversionRate, v))
}

// ConversionRateLTE applies the LTE predicate on the "conversion_rate" field.
func ConversionRateLTE(v float64) predicate.Country {
	return predicate.Country(sql.FieldLTE(FieldConversionRate, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v uint8) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v uint8) predicate.Country {
	return predicate.Country(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...uint8) predicate.Country {
	return predicate.Country(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...uint8) predicate.Country {
	return predicate.Country(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v uint8) predicate.Country {
	return predicate.Country(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v uint8) predicate.Country {
	return predicate.Country(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v uint8) predicate.Country {
	return predicate.Country(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v uint8) predicate.Country {
	return predicate.Country(sql.FieldLTE(FieldStatus, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Country {
	return predicate.Country(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Country {
	return predicate.Country(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Country {
	return predicate.Country(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Country {
	return predicate.Country(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Country {
	return predicate.Country(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Country {
	return predicate.Country(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Country {
	return predicate.Country(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Country {
	return predicate.Country(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Country {
	return predicate.Country(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Country {
	return predicate.Country(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Country {
	return predicate.Country(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Country {
	return predicate.Country(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Country {
	return predicate.Country(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Country {
	return predicate.Country(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasLocations applies the HasEdge predicate on the "locations" edge.
func HasLocations() predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LocationsTable, LocationsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLocationsWith applies the HasEdge predicate on the "locations" edge with a given conditions (other predicates).
func HasLocationsWith(preds ...predicate.CountryLocation) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		step := newLocationsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Country) predicate.Country {
	return predicate.Country(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Country) predicate.Country {
	return predicate.Country(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Country) predicate.Country {
	return predicate.Country(sql.NotPredicates(p))
}
