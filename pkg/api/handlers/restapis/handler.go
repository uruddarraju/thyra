package restapis

import (
	"net/http"

	"github.com/golang/glog"
	"github.com/uruddarraju/thyra/pkg/api/server"
)

type RestAPI struct {
}

func RestAPIHandler(w http.ResponseWriter, r *http.Request) {
	gatewayServer := server.CurrentGatewayServer()

	glog.Infof("%s", gatewayServer)
}
