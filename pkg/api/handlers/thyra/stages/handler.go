package stages

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/uruddarraju/thyra/pkg/api/server"
	"github.com/uruddarraju/thyra/pkg/storage"
)

func List(w http.ResponseWriter, r *http.Request) {
	//gatewayServer := server.CurrentGatewayServer()
	ctx := r.Context()
	name := ctx.Value("name")
	log.Infof("gateway server: %s", name)
}

func Get(w http.ResponseWriter, r *http.Request) {
	gatewayServer := server.CurrentGatewayServer()

	ctx := r.Context()
	name := ctx.Value("name")
	opts := storage.ListOptions{
		APIGroup: "thyra",
		Type:     "stages",
		Name:     name.(string),
	}
	test, err := gatewayServer.Storage().List(ctx, opts)
	log.Infof("in get %s == %s", test, err)
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
