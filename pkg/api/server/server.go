package server

import (
	"net/http"

	"github.com/uruddarraju/thyra/pkg/storage"
)

var gateway GatewayServer

func CurrentGatewayServer() GatewayServer {
	return gateway
}

type GatewayServer interface {
	AddResource()
	DeleteResource()
	AddMethod()
	DeleteMethod()

	AddAPIGroup()
	DeleteAPIGroup()
}

type DefaultGatewayServer struct {
	server                  *http.Server
	apiGroupResourceMapping map[string]string
	storage                 storage.Storage
}

func InitGatewayServer(srv *http.Server) {
	gateway = &DefaultGatewayServer{
		server:  srv,
		storage: storage.NewDefaultStorage(),
	}
}

func (*DefaultGatewayServer) AddResource() {}

func (*DefaultGatewayServer) AddMethod() {}

func (*DefaultGatewayServer) DeleteResource() {}

func (*DefaultGatewayServer) DeleteMethod() {}

func (*DefaultGatewayServer) AddAPIGroup() {}

func (*DefaultGatewayServer) DeleteAPIGroup() {}
