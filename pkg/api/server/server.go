package server

import (
	"net/http"

	"github.com/uruddarraju/thyra/pkg/apis/thyra"
	"github.com/uruddarraju/thyra/pkg/storage"
	"github.com/uruddarraju/thyra/pkg/storage/local"
)

var gateway GatewayServer

func CurrentGatewayServer() GatewayServer {
	return gateway
}

type GatewayServer interface {
	AddAPIGroup(apiGroup string)
	DeleteAPIGroup(apiGroup string)
	AddResource(apiGroup, kind string)
	DeleteResource(apiGroup, kind string)
	AddMethod(apiGroup, kind string, method thyra.HttpMethod)
	DeleteMethod(apiGroup, kind string, method thyra.HttpMethod)
	Storage() storage.Storage
}

type defaultGatewayServer struct {
	server                  *http.Server
	apiGroupResourceMapping map[string]string
	storage                 storage.Storage
}

func InitGatewayServer(srv *http.Server) {
	gateway = &defaultGatewayServer{
		server:  srv,
		storage: local.NewLocalStorage(),
	}
}

func (gs *defaultGatewayServer) AddResource(apiGroup, kind string) {
	err := gs.storage.RegisterKind(nil, apiGroup, kind)
	if err != nil {
		return
	}
}

func (gs *defaultGatewayServer) AddMethod(apiGroup, kind string, method thyra.HttpMethod) {}

func (gs *defaultGatewayServer) DeleteResource(apiGroup, kind string) {}

func (gs *defaultGatewayServer) DeleteMethod(apiGroup, kind string, method thyra.HttpMethod) {}

func (gs *defaultGatewayServer) AddAPIGroup(apiGroup string) {
	err := gs.storage.RegisterGroup(nil, apiGroup)
	if err != nil {
		return
	}
}

func (gs *defaultGatewayServer) DeleteAPIGroup(apiGroup string) {}

func (gs *defaultGatewayServer) Storage() storage.Storage { return gs.storage }
