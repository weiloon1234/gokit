package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/weiloon1234/gokit-base-entity/ent/mixin"
)

// Bank holds the schema definition for the Bank entity.
type Bank struct {
	ent.Schema
}

// Fields of the Bank.
func (Bank) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").
			Unique().
			Immutable().
			Comment("Primary key for the bank"),
		field.String("name_en").
			NotEmpty().
			Comment("Bank name in English"),
		field.String("name_zh").
			NotEmpty().
			Comment("Bank name in Chinese"),
		field.Uint64("country_id").
			Nillable().
			Optional().
			Comment("Country of the bank"),
		field.Time("created_at").
			Default(time.Now).
			Comment("Record creation timestamp"),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("Record update timestamp"),
		field.Time("deleted_at").
			Nillable().
			Optional().
			Comment("Record deleted timestamp"),
	}
}

// Mixins adds the SoftDeleteMixin.
func (Bank) Mixins() []ent.Mixin {
	return []ent.Mixin{
		mixin.SoftDeleteMixin{},
	}
}

// Edges of the Bank.
func (Bank) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("country", Country.Type).
			Ref("banks").
			Field("country_id").
			Unique().
			Comment("Country to which the bank belongs"),

		edge.To("users", User.Type).
			Comment("User bind with this bank"),
	}
}

// Indexes of the Bank.
func (Bank) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("country_id"),
	}
}
