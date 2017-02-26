package gateway

import (
	"context"
	"net/http"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/uruddarraju/thyra/pkg/api/handlers/restapis"
	"github.com/uruddarraju/thyra/pkg/api/server"
	"github.com/uruddarraju/thyra/pkg/api/types"
	"github.com/uruddarraju/thyra/pkg/auth/authn"
	"github.com/uruddarraju/thyra/pkg/plugins/authentication/token/tokenfile"
	utilruntime "k8s.io/kubernetes/pkg/util/runtime"
)

var gateway *defaultGateway
var once sync.Once

type Interface interface {
	Start()
}

type Opts struct {
	KeystoneAuthnConfigFile string
	BasicAuthnFile          string
}

type defaultGateway struct {
	Address       string
	DefaultRouter *httprouter.Router
	Server        *http.Server
	Authenticator authn.Authenticator
}

func DefaultGateway() Interface {
	once.Do(func() {
		defaultRouter := httprouter.New()
		srv := &http.Server{
			Handler:      defaultRouter,
			Addr:         "127.0.0.1:8000",
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}

		gateway = &defaultGateway{
			DefaultRouter: defaultRouter,
			Server:        srv,
		}
	})
	return gateway
}

func (gw *defaultGateway) Start() {

	defer utilruntime.HandleCrash()
	for {
		authn := tokenfile.NewTokenAuthenticator("")
		server.InitGatewayServer(gw.Server)
		gw.AddDefaultHandlers(gw.DefaultRouter, authn)

		chain := alice.New(authn.Authenticator, middlewareTwo)
		gwServer := server.CurrentGatewayServer()
		gwServer.AddAPIGroup("thyra")
		gwServer.AddResource("thyra", "restapis")
		gwServer.AddMethod("thyra", "restapis", api.HTTPGet)
		gw.DefaultRouter.GET("/thyra/restapis/:name", wrapHandler(chain.Then(http.HandlerFunc(restapis.Get))))
		gw.DefaultRouter.GET("/thyra/restapis", wrapHandler(chain.Then(http.HandlerFunc(restapis.List))))

		if err := gw.Server.ListenAndServe(); err != nil {
			log.Errorf("Unable to listen for server (%v); will try again.", err)
		}

		time.Sleep(15 * time.Second)
	}
	log.Fatalf("Server quit.....")
}

func (gw *defaultGateway) AddDefaultHandlers(router *httprouter.Router, authenticator authn.Authenticator) {

	// TODO: Add Union of authenticators
	chain := alice.New(authenticator.Authenticator, middlewareTwo)
	router.GET("/", wrapHandler(chain.Then(http.HandlerFunc(restapis.Get))))
	router.GET("/hello", wrapHandler(chain.Then(http.HandlerFunc(HelloHandler))))
	router.GET("/healthz", wrapHandler(chain.Then(http.HandlerFunc(HealthzHandler))))

	router.GET("/metrics", wrapHandler(chain.Then(http.HandlerFunc(defaultMetricsHandler))))

}

var defaultMetricsHandler = prometheus.Handler().ServeHTTP

func middlewareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("Executing middlewareTwo")
		if r.URL.Path == "/" {
			return
		}
		log.Infof("Executing middlewareTwo .")
		next.ServeHTTP(w, r)
		log.Infof("Executing middlewareTwo again")
	})
}

// We need to do this as the handler function for httprouter is different from that of a regular http.HandleFunc
func wrapHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request,
		ps httprouter.Params) {
		// Setting the path parameters in the context
		ctx := r.Context()
		for _, param := range ps {
			ctx = context.WithValue(ctx, param.Key, param.Value)
		}
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Thyra....!\n"))
}

func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
