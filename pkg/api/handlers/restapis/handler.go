package restapis

import (
	"net/http"

	"github.com/golang/glog"
	"github.com/uruddarraju/thyra/pkg/api/server"
)

func List(w http.ResponseWriter, r *http.Request) {
	gatewayServer := server.CurrentGatewayServer()
	glog.Infof("%s", gatewayServer)
}

func Get(w http.ResponseWriter, r *http.Request) {
	gatewayServer := server.CurrentGatewayServer()
	test := r.Context().Value("test")
	glog.Infof("%s/%s", gatewayServer, test)
}

func Post(w http.ResponseWriter, r *http.Request) {
	gatewayServer := server.CurrentGatewayServer()
	glog.Infof("%s", gatewayServer)
}

func Put(w http.ResponseWriter, r *http.Request) {
	gatewayServer := server.CurrentGatewayServer()
	glog.Infof("%s", gatewayServer)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	gatewayServer := server.CurrentGatewayServer()
	glog.Infof("%s", gatewayServer)
}
