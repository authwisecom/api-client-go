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

	a := assert.New(t)

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
				a.Empty(got)
				a.EqualError(err, "bearer token could not be retrieved")
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
				a.Empty(got)
				a.EqualError(err, "bearer token could not be retrieved")
				a.Equal(ctx, delegate1.passedContext)
				a.Equal(ctx, delegate2.passedContext)
				a.True(delegate1.called)
				a.True(delegate2.called)
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
				a.Equal(token, got)
				a.Nil(err)
				a.Equal(ctx, delegate1.passedContext)
				a.Nil(delegate2.passedContext)
				a.True(delegate1.called)
				a.False(delegate2.called)
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
				a.Empty(got)
				a.Equal(testErr, err)
				a.Equal(ctx, delegate1.passedContext)
				a.Nil(delegate2.passedContext)
				a.True(delegate1.called)
				a.False(delegate2.called)
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
				a.Equal(token, got)
				a.Nil(err)
				a.Equal(ctx, delegate1.passedContext)
				a.Equal(ctx, delegate2.passedContext)
				a.True(delegate1.called)
				a.True(delegate2.called)
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
