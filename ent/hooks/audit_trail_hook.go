package hooks

import (
	"context"

	"entgo.io/ent"
	"github.com/weiloon1234/gokit/logger"
	"github.com/weiloon1234/gokit/middleware"
)

func AuditTrailHook() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, mutation ent.Mutation) (ent.Value, error) {
			if middleware.IsInAdminRoute(ctx) {
				logger.GetLogger().Infof("Admin user %s is performing %s operation on %s", "admin", mutation.Op(), mutation.Op().String())
			}
			return next.Mutate(ctx, mutation)
		})
	}
}
