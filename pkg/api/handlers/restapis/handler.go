package restapis

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/uruddarraju/thyra/pkg/api/server"
)

func List(w http.ResponseWriter, r *http.Request) {
	gatewayServer := server.CurrentGatewayServer()
	log.Infof("%s", gatewayServer)
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
