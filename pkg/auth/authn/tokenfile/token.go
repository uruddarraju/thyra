package tokenfile

import (
	"net/http"

	"github.com/uruddarraju/thyra/pkg/auth/user"
)

type TokenAuthenticator struct {
	Filename string
}

func NewTokenAuthenticator(filename string) *TokenAuthenticator {
	return &TokenAuthenticator{Filename: filename}
}

// TODO: Add authentication logic here to read of a token file and populate the user context
func (ta *TokenAuthenticator) AuthenticateRequest(req *http.Request) (user.UserInfo, bool, error) {
	return user.UserInfo{Username: "uruddarraju"}, true, nil
}
