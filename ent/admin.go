// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/weiloon1234/gokit/ent/admin"
)

// Admin is the model entity for the Admin schema.
type Admin struct {
	config `json:"-"`
	// ID of the ent.
	// Primary key for the admin
	ID uint64 `json:"id,omitempty"`
	// Username of the admin
	Username string `json:"username,omitempty"`
	// Full name of the admin
	Name string `json:"name,omitempty"`
	// Email address of the admin
	Email string `json:"email,omitempty"`
	// Timestamp when email was verified
	EmailVerifiedAt *time.Time `json:"email_verified_at,omitempty"`
	// Password for the admin
	Password string `json:"password,omitempty"`
	// Second password for the admin
	Password2 string `json:"password2,omitempty"`
	// Preferred language of the admin
	Lang string `json:"lang,omitempty"`
	// Avatar of the admin
	Avatar *string `json:"avatar,omitempty"`
	// Type of the admin
	Type uint8 `json:"type,omitempty"`
	// Record creation timestamp
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Record update timestamp
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Record deleted timestamp
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Admin) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case admin.FieldID, admin.FieldType:
			values[i] = new(sql.NullInt64)
		case admin.FieldUsername, admin.FieldName, admin.FieldEmail, admin.FieldPassword, admin.FieldPassword2, admin.FieldLang, admin.FieldAvatar:
			values[i] = new(sql.NullString)
		case admin.FieldEmailVerifiedAt, admin.FieldCreatedAt, admin.FieldUpdatedAt, admin.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Admin fields.
func (a *Admin) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case admin.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = uint64(value.Int64)
		case admin.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				a.Username = value.String
			}
		case admin.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case admin.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				a.Email = value.String
			}
		case admin.FieldEmailVerifiedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field email_verified_at", values[i])
			} else if value.Valid {
				a.EmailVerifiedAt = new(time.Time)
				*a.EmailVerifiedAt = value.Time
			}
		case admin.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				a.Password = value.String
			}
		case admin.FieldPassword2:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password2", values[i])
			} else if value.Valid {
				a.Password2 = value.String
			}
		case admin.FieldLang:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field lang", values[i])
			} else if value.Valid {
				a.Lang = value.String
			}
		case admin.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				a.Avatar = new(string)
				*a.Avatar = value.String
			}
		case admin.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				a.Type = uint8(value.Int64)
			}
		case admin.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case admin.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case admin.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				a.DeletedAt = new(time.Time)
				*a.DeletedAt = value.Time
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Admin.
// This includes values selected through modifiers, order, etc.
func (a *Admin) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// Update returns a builder for updating this Admin.
// Note that you need to call Admin.Unwrap() before calling this method if this Admin
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Admin) Update() *AdminUpdateOne {
	return NewAdminClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Admin entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Admin) Unwrap() *Admin {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Admin is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Admin) String() string {
	var builder strings.Builder
	builder.WriteString("Admin(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("username=")
	builder.WriteString(a.Username)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(a.Name)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(a.Email)
	builder.WriteString(", ")
	if v := a.EmailVerifiedAt; v != nil {
		builder.WriteString("email_verified_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(a.Password)
	builder.WriteString(", ")
	builder.WriteString("password2=")
	builder.WriteString(a.Password2)
	builder.WriteString(", ")
	builder.WriteString("lang=")
	builder.WriteString(a.Lang)
	builder.WriteString(", ")
	if v := a.Avatar; v != nil {
		builder.WriteString("avatar=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", a.Type))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := a.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Admins is a parsable slice of Admin.
type Admins []*Admin
