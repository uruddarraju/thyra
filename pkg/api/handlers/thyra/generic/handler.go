package restapis

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/uruddarraju/thyra/pkg/api/server"
)

// Check for the media type
// Use this to check for the media type:
// mediaType := req.Header.Get("Content-Type")
// mediaType, options, err := mime.ParseMediaType(mediaType)
// support only json and yaml for now
// http://www.alexedwards.net/blog/golang-response-snippets
// decode to object based on the content type
/*
func readBody(req *http.Request) ([]byte, error) {
	defer req.Body.Close()
	return ioutil.ReadAll(req.Body)
}
*/
// decode bytes after this
// get object name from the route path
// Create, Get, List or Delete based on the operation and return appropriate http codes

func List(w http.ResponseWriter, r *http.Request) {
	gatewayServer := server.CurrentGatewayServer()

	log.Infof("gateway server: %s", gatewayServer)
}

func Get(w http.ResponseWriter, r *http.Request) {
	gatewayServer := server.CurrentGatewayServer()
	test := r.Context().Value("test")
	log.Infof("%s/%s", gatewayServer, test)
}

func Post(w http.ResponseWriter, r *http.Request) {
	gatewayServer := server.CurrentGatewayServer()
	log.Infof("%s", gatewayServer)
}

func Put(w http.ResponseWriter, r *http.Request) {
	gatewayServer := server.CurrentGatewayServer()
	log.Infof("%s", gatewayServer)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	gatewayServer := server.CurrentGatewayServer()
	log.Infof("%s", gatewayServer)
}
