package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/weiloon1234/gokit-base-entity/ent/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").
			Unique().
			Immutable().
			Comment("Primary key for the user"),
		field.String("username").
			Optional().
			Comment("Username of the user"),
		field.String("name").
			Optional().
			Comment("Full name of the user"),
		field.String("email").
			Optional().
			Comment("Email address of the user"),
		field.Time("email_verified_at").
			Nillable().
			Optional().
			Comment("Timestamp when email was verified"),
		field.String("password").
			NotEmpty().
			Comment("Password for the user"),
		field.String("password2").
			NotEmpty().
			Comment("Second password for the user"),
		field.Uint64("country_id").
			Nillable().
			Optional().
			Comment("Country ID of the user"),
		field.Uint64("contact_country_id").
			Nillable().
			Optional().
			Comment("Contact country ID of the user"),
		field.String("contact_number").
			Nillable().
			Optional().
			Comment("Contact number of the user"),
		field.String("full_contact_number").
			Nillable().
			Optional().
			Comment("Full contact number of the user"),
		field.Uint64("introducer_user_id").
			Nillable().
			Optional().
			Comment("ID of the introducer user"),
		field.String("lang").
			Default("en").
			Comment("Preferred language of the user"),
		field.Text("avatar").
			Nillable().
			Optional().
			Comment("Avatar of the user"),
		field.Float("credit_1").
			Default(0.0).
			Comment("Credit 1 balance"),
		field.Float("credit_2").
			Default(0.0).
			Comment("Credit 2 balance"),
		field.Float("credit_3").
			Default(0.0).
			Comment("Credit 3 balance"),
		field.Float("credit_4").
			Default(0.0).
			Comment("Credit 4 balance"),
		field.Float("credit_5").
			Default(0.0).
			Comment("Credit 5 balance"),
		field.Uint64("bank_id").
			Nillable().
			Optional().
			Comment("Bank ID of the user"),
		field.String("bank_account_name").
			Nillable().
			Optional().
			Comment("Bank account name of the user"),
		field.String("bank_account_number").
			Nillable().
			Optional().
			Comment("Bank account number of the user"),
		field.String("national_id").
			Nillable().
			Optional().
			Comment("National ID of the user"),
		field.Bool("first_login").
			Default(false).
			Comment("Whether this is the user's first login"),
		field.Time("ban_until").
			Nillable().
			Optional().
			Comment("Timestamp until the user is banned"),
		field.Time("new_login_at").
			Nillable().
			Optional().
			Comment("Timestamp of the user's latest login"),
		field.Time("last_login_at").
			Nillable().
			Optional().
			Comment("Timestamp of the user's last login"),
		field.Uint64("unilevel").
			Nillable().
			Optional().
			Comment("Unilevel of the user"),
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

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// Country the user belongs to
		edge.From("country", Country.Type).
			Ref("users").
			Field("country_id").
			Unique().
			Comment("Country of the user"),

		// Contact country
		edge.From("contact_country", Country.Type).
			Ref("contact_users").
			Field("contact_country_id").
			Unique().
			Comment("Contact country of the user"),

		// Introducer user
		edge.From("introducer", User.Type).
			Ref("introduced_users").
			Field("introducer_user_id").
			Unique().
			Comment("Introducer user"),

		// Bank associated with the user
		edge.From("bank", Bank.Type).
			Ref("users").
			Field("bank_id").
			Unique().
			Comment("Bank associated with the user"),

		edge.To("introduced_users", User.Type).
			Comment("Users introduced by this user"),
	}
}

// Mixins adds the SoftDeleteMixin.
func (User) Mixins() []ent.Mixin {
	return []ent.Mixin{
		mixin.SoftDeleteMixin{},
	}
}

// Indexes of the User.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username"),
		index.Fields("email"),
		index.Fields("country_id"),
		index.Fields("contact_country_id"),
		index.Fields("introducer_user_id"),
		index.Fields("bank_id"),
	}
}
