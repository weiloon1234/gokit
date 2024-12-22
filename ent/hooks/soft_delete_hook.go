package hooks

import (
	"context"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// SoftDeleteHook applies a soft delete filter to all mutations dynamically.
func SoftDeleteHook() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, mutation ent.Mutation) (ent.Value, error) {
			// Apply global filters to exclude soft-deleted records.
			if q, ok := mutation.(interface{ WhereP(...func(*sql.Selector)) }); ok {
				q.WhereP(func(s *sql.Selector) {
					s.Where(sql.IsNull(s.C("deleted_at"))) // Exclude soft-deleted rows
				})
			}
			return next.Mutate(ctx, mutation)
		})
	}
}
