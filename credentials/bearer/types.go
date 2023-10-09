package bearer

import "context"

type Service interface {
	Get(ctx context.Context) (string, error)
}
