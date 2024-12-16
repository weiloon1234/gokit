// ent/schema/country.go
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Country holds the schema definition for the Country entity.
type Country struct {
	ent.Schema
}

// Fields of the Country.
func (Country) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").
			Unique().
			Immutable(),
		field.String("iso2").
			NotEmpty().
			Comment("ISO 3166-1 alpha-2 country code"),
		field.String("iso3").
			NotEmpty().
			Comment("ISO 3166-1 alpha-3 country code"),
		field.String("name").
			NotEmpty().
			Comment("Country name"),
		field.String("official_name").
			Nillable().
			Optional().
			Comment("Official country name"),
		field.String("numeric_code").
			Nillable().
			Optional().
			Comment("ISO 3166-1 numeric code"),
		field.String("phone_code").
			Nillable().
			Optional().
			Comment("Country calling code, e.g., +1 for US"),
		field.String("capital").
			Nillable().
			Optional().
			Comment("Capital city of the country"),
		field.String("currency_name").
			Nillable().
			Optional().
			Comment("Currency name"),
		field.String("currency_code").
			Nillable().
			Optional().
			Comment("ISO 4217 currency code"),
		field.String("currency_symbol").
			Nillable().
			Optional().
			Comment("Currency symbol"),
		field.Float("conversion_rate").
			Default(0.00000).
			Comment("Conversion rate"),
		field.Uint8("status").
			Default(0).
			Comment("Status"),
		field.Time("created_at").
			Default(time.Now).
			Comment("Record creation timestamp"),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("Record update timestamp"),
	}
}

// Edges of the Country.
func (Country) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("locations", CountryLocation.Type).
			Comment("Locations belonging to the country"),
		edge.To("banks", Bank.Type).
			Comment("Banks belonging to the country"),
		edge.To("users", User.Type).
			Comment("Users belonging to the country"),
		edge.To("contact_users", User.Type).
			Comment("Contact country Users belonging to the country"),
	}
}

// Indexes of the Country.
func (Country) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("iso2").Unique(),
		index.Fields("iso3").Unique(),
		index.Fields("status"),
	}
}
