package bearer

import (
	context "context"
	"errors"
)

type mainService struct {
	delegates []Service
}

func (m *mainService) Get(ctx context.Context) (string, error) {
	for _, s := range m.delegates {
		token, err := s.Get(ctx)
		if err != nil {
			return "", err
		}
		if token != "" {
			return token, nil
		}
	}
	return "", errors.New("bearer token could not be retrieved")
}

func NewService(delegates []Service) Service {
	return &mainService{
		delegates: delegates,
	}
}

func NewDefaultService() Service {
	return NewService([]Service{
		NewDefaultLocalStoreService(),
	})
}
