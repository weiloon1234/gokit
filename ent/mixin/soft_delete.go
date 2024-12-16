package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// SoftDeleteMixin adds a `deleted_at` field to implement soft deletes.
type SoftDeleteMixin struct {
	mixin.Schema
}

// Fields adds the `deleted_at` field to the schema.
func (SoftDeleteMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("deleted_at").
			Nillable().
			Optional().
			Comment("Record deletion timestamp for soft deletes"),
	}
}
