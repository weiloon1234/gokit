package utils

import (
	"entgo.io/ent"
)

//**********************************************************
// Sample for chain all the func here
/**
func (User) Fields() []ent.Field {
	baseFields := schema.User{}.Fields()

	// Insert password3 after password2
	modifiedFields := utils.InsertFieldsAfter(baseFields, "password2",
		field.String("password3").Optional().Comment("Third password for the user"),
	)

	// Insert username2 after username
	modifiedFields = utils.InsertFields(modifiedFields, map[string][]ent.Field{
		"username": {field.String("username2").Optional().Comment("Second username for the user")},
	})

	// Modify first_login field to have a default value of true
	modifiedFields = utils.ModifyFields(modifiedFields,
		func(f field.Field) bool {
			return f.Descriptor().Name == "first_login"
		},
		func(result []field.Field, f field.Field) []field.Field {
			result = append(result, f.UpdateDefault(true)) // Change the default value
			return result
		},
	)

	return modifiedFields
}
**/
//**********************************************************

// InsertFieldsAfter inserts one or more new fields after a specified field name.
/**
 * Usage:
    // Insert multiple fields after "password2"
	return gokit.utils.InsertFieldsAfter(
		schema.User{}.Fields(),
		"password2",
		field.String("password3").Optional().Comment("Third password for the user"),
		field.String("password4").Optional().Comment("Fourth password for the user"),
		field.String("password5").Optional().Comment("Fifth password for the user"),
	)
*/
func InsertFieldsAfter(fields []ent.Field, afterFieldName string, newFields ...ent.Field) []ent.Field {
	var result []ent.Field

	for _, f := range fields {
		result = append(result, f) // Add the current field
		if f.Descriptor().Name == afterFieldName {
			result = append(result, newFields...) // Add new fields immediately after
		}
	}

	return result
}

// InsertFields inserts multiple fields dynamically at specified positions.
// `insertions` maps a field name to the fields to be inserted after it.
/**
 * Usage:
	// Define the base fields from gokit
	baseFields := schema.User{}.Fields()

	insertions := map[string][]ent.Field{
		"password2": {field.String("password3").Optional().Comment("Third password for the user")},
		"username":  {field.String("username2").Optional().Comment("Second username for the user")},
	}

	// Use the helper to insert fields dynamically
	return gokit.utils.InsertFields(baseFields, insertions)
**/

func InsertFields(fields []ent.Field, insertions map[string][]ent.Field) []ent.Field {
	var result []ent.Field

	for _, f := range fields {
		result = append(result, f) // Add the current field
		if newFields, ok := insertions[f.Descriptor().Name]; ok {
			result = append(result, newFields...) // Add the new fields at the specified position
		}
	}

	return result
}

// ModifyFields allows modification of fields based on a condition.
// `condition` determines whether a field should be modified.
// `modify` specifies how the field list should be changed.
/**
 * Usage:
	// Define the base fields from gokit
	baseFields := schema.User{}.Fields()

	// Use ModifyFields to insert password3 and username2
	return utils.ModifyFields(baseFields,
		// Condition to find where to modify
		func(f field.Field) bool {
			return f.Descriptor().Name == "password2" || f.Descriptor().Name == "username"
		},
		// Modify function to insert new fields
		func(result []field.Field, f field.Field) []field.Field {
			result = append(result, f) // Keep the original field
			if f.Descriptor().Name == "password2" {
				// Add password3 after password2
				result = append(result, field.String("password3").Optional().Comment("Third password for the user"))
			}
			if f.Descriptor().Name == "username" {
				// Add username2 after username
				result = append(result, field.String("username2").Optional().Comment("Second username for the user"))
			}
			return result
		},
	)
**/
func ModifyFields(fields []ent.Field, condition func(f ent.Field) bool, modify func(result []ent.Field, f ent.Field) []ent.Field) []ent.Field {
	var result []ent.Field

	for _, f := range fields {
		if condition(f) {
			result = modify(result, f) // Apply the modification
		} else {
			result = append(result, f) // Otherwise, keep the field
		}
	}

	return result
}
