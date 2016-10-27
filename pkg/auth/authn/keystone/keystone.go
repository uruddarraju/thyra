package keystone

import (
	"errors"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
)

type KeystoneAuthenticator struct {
	authURL string
}

// NewKeystoneAuthenticator returns a password authenticator that validates credentials using openstack keystone
func NewKeystoneAuthenticator(authURL string) (*KeystoneAuthenticator, error) {
	if !strings.HasPrefix(authURL, "https") {
		return nil, errors.New("Auth URL should be secure and start with https")
	}
	if authURL == "" {
		return nil, errors.New("Auth URL is empty")
	}

	return &KeystoneAuthenticator{authURL: authURL}, nil
}

func (keystoneAuthenticator *KeystoneAuthenticator) AuthenticatePassword(username string, password string) (user.Info, bool, error) {
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: keystoneAuthenticator.authURL,
		Username:         username,
		Password:         password,
	}

	_, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		log.Info("Failed: Starting openstack authenticate client")
		return nil, false, errors.New("Failed to authenticate")
	}

	return &user.DefaultInfo{Name: username}, true, nil
}
