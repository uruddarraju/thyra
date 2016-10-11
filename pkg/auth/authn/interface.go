package authn

import (
	"net/http"

	"github.com/uruddarraju/thyra/pkg/auth/user"
)

type Authenticator interface {
	AuthenticateRequest(req *http.Request) (user.UserInfo, bool, error)
}
