package bearer_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/authwisecom/api-client-go/credentials/bearer"
	"sync"
	"testing"
)

type singleInvokeService struct {
	called        bool
	passedContext context.Context
	resultToken   string
	resultError   error
	lock          sync.Mutex
}

func (s *singleInvokeService) Get(ctx context.Context) (string, error) {

	s.lock.Lock()
	defer s.lock.Unlock()

	if s.called {
		panic("already called")
	}
	s.called = true
	s.passedContext = ctx
	return s.resultToken, s.resultError

}

func TestMainService_Get(t *testing.T) {

	type s struct {
		arrange func(ctx context.Context) (context.Context, []bearer.Service)
		assert  func(ctx context.Context, got string, err error)
	}

	var delegate1 *singleInvokeService
	var delegate2 *singleInvokeService

	token := "test-token"
	testErr := errors.New("test-error")

	cases := map[string]s{
		"no delegates": {
			arrange: func(ctx context.Context) (context.Context, []bearer.Service) {
				return ctx, nil
			},
			assert: func(ctx context.Context, got string, err error) {
				assert.Empty(t, got)
				assert.EqualError(t, err, "bearer token could not be retrieved")
			},
		},
		"delegates no return": {
			arrange: func(ctx context.Context) (context.Context, []bearer.Service) {
				delegate1 = &singleInvokeService{
					resultToken: "",
					resultError: nil,
				}
				delegate2 = &singleInvokeService{
					resultToken: "",
					resultError: nil,
				}
				return ctx, []bearer.Service{delegate1, delegate2}
			},
			assert: func(ctx context.Context, got string, err error) {
				assert.Empty(t, got)
				assert.EqualError(t, err, "bearer token could not be retrieved")
				assert.Same(t, ctx, delegate1.passedContext)
				assert.Same(t, ctx, delegate2.passedContext)
				assert.True(t, delegate1.called)
				assert.True(t, delegate2.called)
			},
		},
		"delegates first return": {
			arrange: func(ctx context.Context) (context.Context, []bearer.Service) {
				delegate1 = &singleInvokeService{
					resultToken: token,
					resultError: nil,
				}
				delegate2 = &singleInvokeService{
					resultToken: "",
					resultError: nil,
				}
				return ctx, []bearer.Service{delegate1, delegate2}
			},
			assert: func(ctx context.Context, got string, err error) {
				assert.Equal(t, token, got)
				assert.Nil(t, err)
				assert.Same(t, ctx, delegate1.passedContext)
				assert.Nil(t, delegate2.passedContext)
				assert.True(t, delegate1.called)
				assert.False(t, delegate2.called)
			},
		},
		"delegates first error": {
			arrange: func(ctx context.Context) (context.Context, []bearer.Service) {
				delegate1 = &singleInvokeService{
					resultToken: token,
					resultError: testErr,
				}
				delegate2 = &singleInvokeService{
					resultToken: "",
					resultError: nil,
				}
				return ctx, []bearer.Service{delegate1, delegate2}
			},
			assert: func(ctx context.Context, got string, err error) {
				assert.Empty(t, got)
				assert.Same(t, testErr, err)
				assert.Same(t, ctx, delegate1.passedContext)
				assert.Nil(t, delegate2.passedContext)
				assert.True(t, delegate1.called)
				assert.False(t, delegate2.called)
			},
		},
		"delegates second": {
			arrange: func(ctx context.Context) (context.Context, []bearer.Service) {
				delegate1 = &singleInvokeService{
					resultToken: "",
					resultError: nil,
				}
				delegate2 = &singleInvokeService{
					resultToken: token,
					resultError: nil,
				}
				return ctx, []bearer.Service{delegate1, delegate2}
			},
			assert: func(ctx context.Context, got string, err error) {
				assert.Equal(t, token, got)
				assert.Nil(t, err)
				assert.Same(t, ctx, delegate1.passedContext)
				assert.Same(t, ctx, delegate2.passedContext)
				assert.True(t, delegate1.called)
				assert.True(t, delegate2.called)
			},
		},
	}

	for k, v := range cases {
		t.Run(k, func(t *testing.T) {

			ctx, delegates := v.arrange(context.Background())
			unit := bearer.NewService(delegates)
			got, err := unit.Get(ctx)
			v.assert(ctx, got, err)
		})
	}
}
