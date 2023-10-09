package bearer_test

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/authwisecom/api-client-go/credentials/bearer"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestLocalStore_Get(t *testing.T) {

	type s struct {
		arrange func(ctx context.Context) (context.Context, string)
		assert  func(got string, err error)
	}

	token := "test-token"

	cases := map[string]s{
		"file not exists": {
			arrange: func(ctx context.Context) (context.Context, string) {

				f, err := os.MkdirTemp("", "unit")
				if err != nil {
					panic(err)
				}
				return ctx, filepath.Join(f, "nothing.json")
			},
			assert: func(got string, err error) {
				assert.Empty(t, got)
				assert.Nil(t, err)
			},
		},
		"token expired": {
			arrange: func(ctx context.Context) (context.Context, string) {

				f, err := os.CreateTemp("", "creds*.json")
				if err != nil {
					panic(err)
				}
				data := &bearer.Data{
					AccessToken: token,
					ExpiresAt:   time.Now().Unix() + 59,
				}
				err = json.NewEncoder(f).Encode(data)
				if err != nil {
					panic(err)
				}
				return ctx, f.Name()
			},
			assert: func(got string, err error) {
				assert.Empty(t, got)
				assert.Nil(t, err)
			},
		},
		"no token": {
			arrange: func(ctx context.Context) (context.Context, string) {

				f, err := os.CreateTemp("", "creds*.json")
				if err != nil {
					panic(err)
				}
				data := &bearer.Data{}
				err = json.NewEncoder(f).Encode(data)
				if err != nil {
					panic(err)
				}
				return ctx, f.Name()
			},
			assert: func(got string, err error) {
				assert.Empty(t, got)
				assert.Nil(t, err)
			},
		},
		"token": {
			arrange: func(ctx context.Context) (context.Context, string) {

				f, err := os.CreateTemp("", "creds*.json")
				if err != nil {
					panic(err)
				}
				data := &bearer.Data{
					AccessToken: token,
					ExpiresAt:   time.Now().Unix() + 70,
				}
				err = json.NewEncoder(f).Encode(data)
				if err != nil {
					panic(err)
				}
				return ctx, f.Name()
			},
			assert: func(got string, err error) {
				assert.Equal(t, token, got)
				assert.Nil(t, err)
			},
		},
	}

	for k, v := range cases {
		t.Run(k, func(t *testing.T) {

			ctx, path := v.arrange(context.Background())
			unit := bearer.NewLocalStoreService(path)
			v.assert(unit.Get(ctx))

		})
	}
}
