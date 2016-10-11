package authn

import (
	"fmt"
	"net/http"

	"github.com/golang/glog"
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
		glog.Infof("Handler being called %s %s", user, authenticated)
		handler(w, r)
	}
}
