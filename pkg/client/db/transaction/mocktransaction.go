package transaction

import (
	"context"

	"github.com/Paul1k96/microservices_course_platform_common/pkg/client/db"
)

// NopTxManager is a transaction manager that does nothing.
// It is used in tests where transactions are not needed.
type NopTxManager struct{}

// NewNopTxManager creates a new NopTxManager.
func NewNopTxManager() *NopTxManager {
	return &NopTxManager{}
}

// ReadCommitted executes the user-specified handler in a read committed transaction
func (NopTxManager) ReadCommitted(ctx context.Context, fn db.Handler) error {
	return fn(ctx)
}
