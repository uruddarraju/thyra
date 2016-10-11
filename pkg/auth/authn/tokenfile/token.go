package tokenfile

import (
	"log"
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

func (ta *TokenAuthenticator) Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing authenticator")
		next.ServeHTTP(w, r)
		log.Println("Executing authenticator again")
	})
}
