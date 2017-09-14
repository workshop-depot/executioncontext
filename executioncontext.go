package executioncontext

import (
	"context"
	"sync"
)

// WaitGroup interface for built-in WaitGroup
type WaitGroup interface {
	Add(delta int)
	Done()
	Wait()
}

type executionContext struct {
	ctx context.Context
	wg  *sync.WaitGroup
}

func (etx *executionContext) Context() context.Context { return etx.ctx }
func (etx *executionContext) WaitGroup() WaitGroup     { return etx.wg }

// embedding context.Context would make usage of Done ambiguous and error-prone

// Context combination of context.Context & WaitGroup, an execution context
type Context interface {
	Context() context.Context
	WaitGroup() WaitGroup
}

// ErrNilContext means we got a nil context while we shouldn't
var ErrNilContext = errorf("ERR_NIL_CONTEXT")

// New .
func New(ctx context.Context) (Context, error) {
	if ctx == nil {
		return nil, ErrNilContext
	}
	return &executionContext{
		ctx: ctx,
		wg:  &sync.WaitGroup{},
	}, nil
}
