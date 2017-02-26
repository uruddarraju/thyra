package union

import (
	"net/http"

	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/uruddarraju/thyra/pkg/auth/authn"
	"github.com/uruddarraju/thyra/pkg/auth/user"
	"github.com/uruddarraju/thyra/pkg/plugins/authentication/keystone"
	"github.com/uruddarraju/thyra/pkg/plugins/authentication/tokenfile"
)

type UnionAuthenticator struct {
	authenticators []authn.Authenticator
}

func NewUnionAuthenticator(keystoneURL string, tokenFile string) (*UnionAuthenticator, error) {
	authenticators := []authn.Authenticator{}
	if len(keystoneURL) > 0 {
		ka, err := keystone.NewKeystoneAuthenticator(keystoneURL)
		if err != nil {
			log.Errorf("Unable to initialize Keystone authentication: %v", err)
			return nil, fmt.Errorf("Unable to initialize Keystone authentication: %v", err)
		}
		authenticators = append(authenticators, ka)
	}
	if len(tokenFile) > 0 {
		ta := tokenfile.NewTokenAuthenticator(tokenFile)
		authenticators = append(authenticators, ta)
	}
	return &UnionAuthenticator{authenticators: authenticators}, nil
}

// TODO: Add authentication logic here to read of a token file and populate the user context
func (ua *UnionAuthenticator) AuthenticateRequest(req *http.Request) (userInfo user.UserInfo, authenticated bool, err error) {
	for _, authenticator := range ua.authenticators {
		userInfo, authenticated, err = authenticator.AuthenticateRequest(req)
		if err != nil || !authenticated {
			continue
		}
		return
	}
	return nil, false, fmt.Errorf("Union of authenticators could not authenticate the request")
}

func (ua *UnionAuthenticator) Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, authenticated, err := ua.AuthenticateRequest(r)
		if !authenticated || err != nil {
			log.Warningf("Attempted administrative access with invalid or missing key!")
			message := "Unauthorized"
			w.WriteHeader(401)
			fmt.Fprintf(w, message)
			return
		}

		next.ServeHTTP(w, r)
		log.Infof("Executing authenticator again")
	})
}
