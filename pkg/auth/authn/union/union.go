package union

import (
	"log"
	"net/http"

	"github.com/uruddarraju/thyra/pkg/auth/user"
	"fmt"
	"github.com/golang/glog"
	"github.com/uruddarraju/thyra/pkg/auth/authn"
)

type UnionAuthenticator struct {
	authenticators []authn.Authenticator
}

func NewUnionAuthenticator(authenticators []authn.Authenticator) *UnionAuthenticator {
	return &UnionAuthenticator{authenticators: authenticators}
}

// TODO: Add authentication logic here to read of a token file and populate the user context
func (ua *UnionAuthenticator) AuthenticateRequest(req *http.Request) (userInfo user.UserInfo, authenticated bool, err error) {
	for _, authenticator := range ua.authenticators {
		userInfo, authenticated, err = authenticator.AuthenticateRequest(req)
		if err != nil || !authenticated {
			continue;
		}
		return
	}
	return nil, false, fmt.Errorf("Union of authenticators could not authenticate the request")
}

func (ua *UnionAuthenticator) Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, authenticated, err := ua.AuthenticateRequest(r)
		if !authenticated || err != nil {
			glog.Warningf("Attempted administrative access with invalid or missing key!")
			message := "Unauthorized"
			w.WriteHeader(401)
			fmt.Fprintf(w, message)
			return
		}

		next.ServeHTTP(w, r)
		log.Println("Executing authenticator again")
	})
}
