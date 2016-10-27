package authn

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/uruddarraju/thyra/pkg/auth/authn"
)

func Authenticate(authenticator authn.Authenticator, handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user, authenticated, err := authenticator.AuthenticateRequest(r)
		if !authenticated || err != nil {
			responseMessage := "Unauthorized"
			w.WriteHeader(401)
			fmt.Fprintf(w, responseMessage)
			return
		}
		log.Infof("Handler being called %s %s", user, authenticated)
		handler(w, r)
	}
}
