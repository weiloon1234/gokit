package hooks

import (
	"context"
	"fmt"

	"entgo.io/ent"
	"github.com/weiloon1234/gokit/globalcontext"
	"github.com/weiloon1234/gokit/logger"
)

func AuditTrailHook(next ent.Mutator) ent.Mutator {
	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		// Retrieve the Gin context
		ginCtx := globalcontext.GetGinContext(ctx)
		if ginCtx != nil {
			requestID := ginCtx.GetString("requestID")
			logger.GetLogger().Info(fmt.Sprintf("Audit Trail - RequestID: %s, Entity: %s, Operation: %s", requestID, m.Type(), m.Op()))
		} else {
			logger.GetLogger().Info("Gin context not found in Ent hook")
		}

		// Proceed with the next mutator
		return next.Mutate(ctx, m)
	})
}
