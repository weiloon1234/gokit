package hook

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/weiloon1234/gokit/ent"
)

// AddSoftDeleteFilter applies a query filter to exclude soft-deleted records.
func SoftDeleteHook(client *ent.Client) {
	client.Use(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, mutation ent.Mutation) (ent.Value, error) {
			// Apply global filters to exclude soft-deleted records.
			if q, ok := mutation.(interface{ WhereP(...func(*sql.Selector)) }); ok {
				q.WhereP(func(s *sql.Selector) {
					s.Where(sql.IsNull(s.C("deleted_at")))
				})
			}
			return next.Mutate(ctx, mutation)
		})
	})
}
