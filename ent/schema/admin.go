package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/weiloon1234/gokit-base-entity/ent/mixin"
)

// Admin holds the schema definition for the Admin entity.
type Admin struct {
	ent.Schema
}

// Fields of the Admin.
func (Admin) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").
			Unique().
			Immutable().
			Comment("Primary key for the admin"),
		field.String("username").
			Optional().
			Comment("Username of the admin"),
		field.String("name").
			Optional().
			Comment("Full name of the admin"),
		field.String("email").
			Optional().
			Comment("Email address of the admin"),
		field.Time("email_verified_at").
			Nillable().
			Optional().
			Comment("Timestamp when email was verified"),
		field.String("password").
			NotEmpty().
			Comment("Password for the admin"),
		field.String("password2").
			NotEmpty().
			Comment("Second password for the admin"),
		field.String("lang").
			Default("en").
			Comment("Preferred language of the admin"),
		field.Text("avatar").
			Nillable().
			Optional().
			Comment("Avatar of the admin"),
		field.Uint8("type").
			Default(0).
			Comment("Type of the admin"),
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
func (Admin) Mixins() []ent.Mixin {
	return []ent.Mixin{
		mixin.SoftDeleteMixin{},
	}
}

// Indexes of the Admin.
func (Admin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username"),
		index.Fields("email"),
	}
}
