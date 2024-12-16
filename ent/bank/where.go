// Code generated by ent, DO NOT EDIT.

package bank

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/weiloon1234/gokit/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Bank {
	return predicate.Bank(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Bank {
	return predicate.Bank(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Bank {
	return predicate.Bank(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Bank {
	return predicate.Bank(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.Bank {
	return predicate.Bank(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.Bank {
	return predicate.Bank(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Bank {
	return predicate.Bank(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Bank {
	return predicate.Bank(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Bank {
	return predicate.Bank(sql.FieldLTE(FieldID, id))
}

// NameEn applies equality check predicate on the "name_en" field. It's identical to NameEnEQ.
func NameEn(v string) predicate.Bank {
	return predicate.Bank(sql.FieldEQ(FieldNameEn, v))
}

// NameZh applies equality check predicate on the "name_zh" field. It's identical to NameZhEQ.
func NameZh(v string) predicate.Bank {
	return predicate.Bank(sql.FieldEQ(FieldNameZh, v))
}

// CountryID applies equality check predicate on the "country_id" field. It's identical to CountryIDEQ.
func CountryID(v uint64) predicate.Bank {
	return predicate.Bank(sql.FieldEQ(FieldCountryID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Bank {
	return predicate.Bank(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Bank {
	return predicate.Bank(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Bank {
	return predicate.Bank(sql.FieldEQ(FieldDeletedAt, v))
}

// NameEnEQ applies the EQ predicate on the "name_en" field.
func NameEnEQ(v string) predicate.Bank {
	return predicate.Bank(sql.FieldEQ(FieldNameEn, v))
}

// NameEnNEQ applies the NEQ predicate on the "name_en" field.
func NameEnNEQ(v string) predicate.Bank {
	return predicate.Bank(sql.FieldNEQ(FieldNameEn, v))
}

// NameEnIn applies the In predicate on the "name_en" field.
func NameEnIn(vs ...string) predicate.Bank {
	return predicate.Bank(sql.FieldIn(FieldNameEn, vs...))
}

// NameEnNotIn applies the NotIn predicate on the "name_en" field.
func NameEnNotIn(vs ...string) predicate.Bank {
	return predicate.Bank(sql.FieldNotIn(FieldNameEn, vs...))
}

// NameEnGT applies the GT predicate on the "name_en" field.
func NameEnGT(v string) predicate.Bank {
	return predicate.Bank(sql.FieldGT(FieldNameEn, v))
}

// NameEnGTE applies the GTE predicate on the "name_en" field.
func NameEnGTE(v string) predicate.Bank {
	return predicate.Bank(sql.FieldGTE(FieldNameEn, v))
}

// NameEnLT applies the LT predicate on the "name_en" field.
func NameEnLT(v string) predicate.Bank {
	return predicate.Bank(sql.FieldLT(FieldNameEn, v))
}

// NameEnLTE applies the LTE predicate on the "name_en" field.
func NameEnLTE(v string) predicate.Bank {
	return predicate.Bank(sql.FieldLTE(FieldNameEn, v))
}

// NameEnContains applies the Contains predicate on the "name_en" field.
func NameEnContains(v string) predicate.Bank {
	return predicate.Bank(sql.FieldContains(FieldNameEn, v))
}

// NameEnHasPrefix applies the HasPrefix predicate on the "name_en" field.
func NameEnHasPrefix(v string) predicate.Bank {
	return predicate.Bank(sql.FieldHasPrefix(FieldNameEn, v))
}

// NameEnHasSuffix applies the HasSuffix predicate on the "name_en" field.
func NameEnHasSuffix(v string) predicate.Bank {
	return predicate.Bank(sql.FieldHasSuffix(FieldNameEn, v))
}

// NameEnEqualFold applies the EqualFold predicate on the "name_en" field.
func NameEnEqualFold(v string) predicate.Bank {
	return predicate.Bank(sql.FieldEqualFold(FieldNameEn, v))
}

// NameEnContainsFold applies the ContainsFold predicate on the "name_en" field.
func NameEnContainsFold(v string) predicate.Bank {
	return predicate.Bank(sql.FieldContainsFold(FieldNameEn, v))
}

// NameZhEQ applies the EQ predicate on the "name_zh" field.
func NameZhEQ(v string) predicate.Bank {
	return predicate.Bank(sql.FieldEQ(FieldNameZh, v))
}

// NameZhNEQ applies the NEQ predicate on the "name_zh" field.
func NameZhNEQ(v string) predicate.Bank {
	return predicate.Bank(sql.FieldNEQ(FieldNameZh, v))
}

// NameZhIn applies the In predicate on the "name_zh" field.
func NameZhIn(vs ...string) predicate.Bank {
	return predicate.Bank(sql.FieldIn(FieldNameZh, vs...))
}

// NameZhNotIn applies the NotIn predicate on the "name_zh" field.
func NameZhNotIn(vs ...string) predicate.Bank {
	return predicate.Bank(sql.FieldNotIn(FieldNameZh, vs...))
}

// NameZhGT applies the GT predicate on the "name_zh" field.
func NameZhGT(v string) predicate.Bank {
	return predicate.Bank(sql.FieldGT(FieldNameZh, v))
}

// NameZhGTE applies the GTE predicate on the "name_zh" field.
func NameZhGTE(v string) predicate.Bank {
	return predicate.Bank(sql.FieldGTE(FieldNameZh, v))
}

// NameZhLT applies the LT predicate on the "name_zh" field.
func NameZhLT(v string) predicate.Bank {
	return predicate.Bank(sql.FieldLT(FieldNameZh, v))
}

// NameZhLTE applies the LTE predicate on the "name_zh" field.
func NameZhLTE(v string) predicate.Bank {
	return predicate.Bank(sql.FieldLTE(FieldNameZh, v))
}

// NameZhContains applies the Contains predicate on the "name_zh" field.
func NameZhContains(v string) predicate.Bank {
	return predicate.Bank(sql.FieldContains(FieldNameZh, v))
}

// NameZhHasPrefix applies the HasPrefix predicate on the "name_zh" field.
func NameZhHasPrefix(v string) predicate.Bank {
	return predicate.Bank(sql.FieldHasPrefix(FieldNameZh, v))
}

// NameZhHasSuffix applies the HasSuffix predicate on the "name_zh" field.
func NameZhHasSuffix(v string) predicate.Bank {
	return predicate.Bank(sql.FieldHasSuffix(FieldNameZh, v))
}

// NameZhEqualFold applies the EqualFold predicate on the "name_zh" field.
func NameZhEqualFold(v string) predicate.Bank {
	return predicate.Bank(sql.FieldEqualFold(FieldNameZh, v))
}

// NameZhContainsFold applies the ContainsFold predicate on the "name_zh" field.
func NameZhContainsFold(v string) predicate.Bank {
	return predicate.Bank(sql.FieldContainsFold(FieldNameZh, v))
}

// CountryIDEQ applies the EQ predicate on the "country_id" field.
func CountryIDEQ(v uint64) predicate.Bank {
	return predicate.Bank(sql.FieldEQ(FieldCountryID, v))
}

// CountryIDNEQ applies the NEQ predicate on the "country_id" field.
func CountryIDNEQ(v uint64) predicate.Bank {
	return predicate.Bank(sql.FieldNEQ(FieldCountryID, v))
}

// CountryIDIn applies the In predicate on the "country_id" field.
func CountryIDIn(vs ...uint64) predicate.Bank {
	return predicate.Bank(sql.FieldIn(FieldCountryID, vs...))
}

// CountryIDNotIn applies the NotIn predicate on the "country_id" field.
func CountryIDNotIn(vs ...uint64) predicate.Bank {
	return predicate.Bank(sql.FieldNotIn(FieldCountryID, vs...))
}

// CountryIDIsNil applies the IsNil predicate on the "country_id" field.
func CountryIDIsNil() predicate.Bank {
	return predicate.Bank(sql.FieldIsNull(FieldCountryID))
}

// CountryIDNotNil applies the NotNil predicate on the "country_id" field.
func CountryIDNotNil() predicate.Bank {
	return predicate.Bank(sql.FieldNotNull(FieldCountryID))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Bank {
	return predicate.Bank(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Bank {
	return predicate.Bank(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Bank {
	return predicate.Bank(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Bank {
	return predicate.Bank(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Bank {
	return predicate.Bank(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Bank {
	return predicate.Bank(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Bank {
	return predicate.Bank(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Bank {
	return predicate.Bank(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Bank {
	return predicate.Bank(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Bank {
	return predicate.Bank(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Bank {
	return predicate.Bank(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Bank {
	return predicate.Bank(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Bank {
	return predicate.Bank(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Bank {
	return predicate.Bank(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Bank {
	return predicate.Bank(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Bank {
	return predicate.Bank(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Bank {
	return predicate.Bank(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Bank {
	return predicate.Bank(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Bank {
	return predicate.Bank(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Bank {
	return predicate.Bank(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Bank {
	return predicate.Bank(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Bank {
	return predicate.Bank(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Bank {
	return predicate.Bank(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Bank {
	return predicate.Bank(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Bank {
	return predicate.Bank(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Bank {
	return predicate.Bank(sql.FieldNotNull(FieldDeletedAt))
}

// HasCountry applies the HasEdge predicate on the "country" edge.
func HasCountry() predicate.Bank {
	return predicate.Bank(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CountryTable, CountryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCountryWith applies the HasEdge predicate on the "country" edge with a given conditions (other predicates).
func HasCountryWith(preds ...predicate.Country) predicate.Bank {
	return predicate.Bank(func(s *sql.Selector) {
		step := newCountryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Bank {
	return predicate.Bank(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Bank {
	return predicate.Bank(func(s *sql.Selector) {
		step := newUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Bank) predicate.Bank {
	return predicate.Bank(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Bank) predicate.Bank {
	return predicate.Bank(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Bank) predicate.Bank {
	return predicate.Bank(sql.NotPredicates(p))
}
