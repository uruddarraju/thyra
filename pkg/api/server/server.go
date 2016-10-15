package server

import (
	"net/http"

	"github.com/uruddarraju/thyra/pkg/storage"
)

var gateway GatewayServer

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
	storage                 *storage.Storage
}

func InitGatewayServer(srv *http.Server) GatewayServer {
	gateway = &DefaultGatewayServer{
		server:  srv,
		storage: storage.NewDefaultStorage(),
	}
	return gateway
}

func (*DefaultGatewayServer) AddResource() {}

func (*DefaultGatewayServer) AddMethod() {}

func (*DefaultGatewayServer) DeleteResource() {}

func (*DefaultGatewayServer) DeleteMethod() {}

func (*DefaultGatewayServer) AddAPIGroup() {}

func (*DefaultGatewayServer) DeleteAPIGroup() {}
