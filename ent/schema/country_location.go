// ent/schema/countrylocation.go
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/weiloon1234/gokit/ent/mixin"
)

// CountryLocation holds the schema definition for the CountryLocation entity.
type CountryLocation struct {
	ent.Schema
}

// Fields of the CountryLocation.
func (CountryLocation) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").
			Unique().
			Immutable(),
		field.Uint64("country_id").
			Nillable().
			Optional().
			Comment("Country of the location"),
		field.Uint64("parent_id").
			Nillable().
			Optional().
			Comment("If present, it is area; this is state id"),
		field.Uint64("sorting").
			Default(0).
			Comment("Sorting order"),
		field.String("name_en").
			NotEmpty().
			Comment("Location name in English"),
		field.String("name_zh").
			NotEmpty().
			Comment("Location name in Chinese"),
		field.Time("created_at").
			Default(time.Now).
			Comment("Record creation timestamp"),
		field.Time("updated_at").
			Default(time.Now()).
			UpdateDefault(time.Now).
			Comment("Record update timestamp"),
		field.Time("deleted_at").
			Nillable().
			Optional().
			Comment("Record deletion timestamp"),
	}
}

// Mixins adds the SoftDeleteMixin.
func (CountryLocation) Mixins() []ent.Mixin {
	return []ent.Mixin{
		mixin.SoftDeleteMixin{},
	}
}

// Edges of the CountryLocation.
func (CountryLocation) Edges() []ent.Edge {
	return []ent.Edge{
		// Country to which this location belongs
		edge.From("country", Country.Type).
			Ref("locations").
			Field("country_id").
			Unique().
			Comment("Country of the location"),

		// Parent location of this location (e.g., state for a city)
		edge.From("parent", CountryLocation.Type).
			Ref("child_locations").
			Field("parent_id").
			Unique().
			Comment("Parent location (e.g., state of an area)"),

		// Child locations of this location (e.g., cities under a state)
		edge.To("child_locations", CountryLocation.Type).
			Comment("Child locations of this parent"),
	}
}

// Indexes of the CountryLocation.
func (CountryLocation) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("country_id"),
		index.Fields("parent_id"),
		index.Fields("sorting"),
	}
}
