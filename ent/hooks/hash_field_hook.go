package hooks

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"entgo.io/ent"
)

// HashFieldHook creates a hook that hashes the specified field dynamically.
func HashFieldHook(fieldName string) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, mutation ent.Mutation) (ent.Value, error) {
			// Retrieve the field value to check if it is being modified
			fieldValue, exists := mutation.Field(fieldName)
			if exists {
				// Ensure the value is a string before hashing
				password, ok := fieldValue.(string)
				if !ok {
					return nil, fmt.Errorf("field %s value is not a string", fieldName)
				}

				// Hash the password
				hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
				if err != nil {
					return nil, fmt.Errorf("failed to hash %s: %w", fieldName, err)
				}

				// Update the mutation with the hashed value
				mutation.SetField(fieldName, string(hashedPassword))
			}

			// Proceed to the next mutator in the chain
			return next.Mutate(ctx, mutation)
		})
	}
}
