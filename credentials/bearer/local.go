package bearer

import (
	context "context"
	"encoding/json"
	homedir "github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
	"time"
)

type Data struct {
	Issuer      string `json:"issuer,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
	ExpiresAt   int64  `json:"expires_at,omitempty"`
}

type localStore struct {
	localStorePath string
}

func (l *localStore) Get(ctx context.Context) (string, error) {

	f, err := os.Open(l.localStorePath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		} else {
			return "", err
		}
	}
	defer f.Close()

	data := &Data{}

	err = json.NewDecoder(f).Decode(data)

	if err != nil {
		return "", err
	}

	if data.AccessToken == "" {
		return "", nil
	}

	// One minute grace period
	if data.ExpiresAt < (time.Now().Unix() + 60) {
		return "", nil
	}

	return data.AccessToken, nil

}

func NewDefaultLocalStoreService() Service {
	return NewLocalStoreService(DefaultLocalStorePath())
}

func NewLocalStoreService(localStorePath string) Service {
	return &localStore{localStorePath: localStorePath}
}

func DefaultLocalStorePath() string {

	path, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(path, ".awctl.d", "credentials", "credentials.json")
}
