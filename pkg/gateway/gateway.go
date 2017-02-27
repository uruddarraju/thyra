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
	"github.com/uruddarraju/thyra/pkg/api/handlers/thyra"
	"github.com/uruddarraju/thyra/pkg/api/handlers/thyra/authorizers"
	"github.com/uruddarraju/thyra/pkg/api/handlers/thyra/deployments"
	"github.com/uruddarraju/thyra/pkg/api/handlers/thyra/integrations"
	"github.com/uruddarraju/thyra/pkg/api/handlers/thyra/methods"
	"github.com/uruddarraju/thyra/pkg/api/handlers/thyra/resources"
	"github.com/uruddarraju/thyra/pkg/api/handlers/thyra/restapis"
	"github.com/uruddarraju/thyra/pkg/api/handlers/thyra/stages"
	"github.com/uruddarraju/thyra/pkg/api/server"
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
		gw.installThyraAPIs(gw.DefaultRouter, authn)

		if err := gw.Server.ListenAndServe(); err != nil {
			log.Errorf("Unable to listen for server (%v); will try again.", err)
		}
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

func (gw *defaultGateway) installThyraAPIs(router *httprouter.Router, authenticator authn.Authenticator) {

	chain := alice.New(authenticator.Authenticator, middlewareTwo)
	gwServer := server.CurrentGatewayServer()

	gwServer.AddAPIGroup("thyra")
	gw.DefaultRouter.GET("/thyra", wrapHandler(chain.Then(http.HandlerFunc(thyra.Get))))

	gwServer.AddResource("thyra", "restapis")
	gw.DefaultRouter.PUT("/thyra/restapis/:name", wrapHandler(chain.Then(http.HandlerFunc(restapis.Put))))
	gw.DefaultRouter.DELETE("/thyra/restapis/:name", wrapHandler(chain.Then(http.HandlerFunc(restapis.Delete))))
	gw.DefaultRouter.GET("/thyra/restapis/:name", wrapHandler(chain.Then(http.HandlerFunc(restapis.Get))))
	gw.DefaultRouter.GET("/thyra/restapis", wrapHandler(chain.Then(http.HandlerFunc(restapis.List))))
	gw.DefaultRouter.POST("/thyra/restapis", wrapHandler(chain.Then(http.HandlerFunc(restapis.Post))))

	gwServer.AddResource("thyra", "resources")
	gw.DefaultRouter.PUT("/thyra/resources/:name", wrapHandler(chain.Then(http.HandlerFunc(resources.Put))))
	gw.DefaultRouter.DELETE("/thyra/resources/:name", wrapHandler(chain.Then(http.HandlerFunc(resources.Delete))))
	gw.DefaultRouter.GET("/thyra/resources/:name", wrapHandler(chain.Then(http.HandlerFunc(resources.Get))))
	gw.DefaultRouter.GET("/thyra/resources", wrapHandler(chain.Then(http.HandlerFunc(resources.List))))
	gw.DefaultRouter.POST("/thyra/resources", wrapHandler(chain.Then(http.HandlerFunc(resources.Post))))

	gwServer.AddResource("thyra", "methods")
	gw.DefaultRouter.PUT("/thyra/methods/:name", wrapHandler(chain.Then(http.HandlerFunc(methods.Put))))
	gw.DefaultRouter.DELETE("/thyra/methods/:name", wrapHandler(chain.Then(http.HandlerFunc(methods.Delete))))
	gw.DefaultRouter.GET("/thyra/methods/:name", wrapHandler(chain.Then(http.HandlerFunc(methods.Get))))
	gw.DefaultRouter.GET("/thyra/methods", wrapHandler(chain.Then(http.HandlerFunc(methods.List))))
	gw.DefaultRouter.POST("/thyra/methods", wrapHandler(chain.Then(http.HandlerFunc(methods.Post))))

	gwServer.AddResource("thyra", "authorizers")
	gw.DefaultRouter.PUT("/thyra/authorizers/:name", wrapHandler(chain.Then(http.HandlerFunc(authorizers.Put))))
	gw.DefaultRouter.DELETE("/thyra/authorizers/:name", wrapHandler(chain.Then(http.HandlerFunc(authorizers.Delete))))
	gw.DefaultRouter.GET("/thyra/authorizers/:name", wrapHandler(chain.Then(http.HandlerFunc(authorizers.Get))))
	gw.DefaultRouter.GET("/thyra/authorizers", wrapHandler(chain.Then(http.HandlerFunc(authorizers.List))))
	gw.DefaultRouter.POST("/thyra/authorizers", wrapHandler(chain.Then(http.HandlerFunc(authorizers.Post))))

	gwServer.AddResource("thyra", "stages")
	gw.DefaultRouter.PUT("/thyra/stages/:name", wrapHandler(chain.Then(http.HandlerFunc(stages.Put))))
	gw.DefaultRouter.DELETE("/thyra/stages/:name", wrapHandler(chain.Then(http.HandlerFunc(stages.Delete))))
	gw.DefaultRouter.GET("/thyra/stages/:name", wrapHandler(chain.Then(http.HandlerFunc(stages.Get))))
	gw.DefaultRouter.GET("/thyra/stages", wrapHandler(chain.Then(http.HandlerFunc(stages.List))))
	gw.DefaultRouter.POST("/thyra/stages", wrapHandler(chain.Then(http.HandlerFunc(stages.Post))))

	gwServer.AddResource("thyra", "deployments")
	gw.DefaultRouter.PUT("/thyra/deployments/:name", wrapHandler(chain.Then(http.HandlerFunc(deployments.Put))))
	gw.DefaultRouter.DELETE("/thyra/deployments/:name", wrapHandler(chain.Then(http.HandlerFunc(deployments.Delete))))
	gw.DefaultRouter.GET("/thyra/deployments/:name", wrapHandler(chain.Then(http.HandlerFunc(deployments.Get))))
	gw.DefaultRouter.GET("/thyra/deployments", wrapHandler(chain.Then(http.HandlerFunc(deployments.List))))
	gw.DefaultRouter.POST("/thyra/deployments", wrapHandler(chain.Then(http.HandlerFunc(deployments.Post))))

	gwServer.AddResource("thyra", "integrations")
	gw.DefaultRouter.PUT("/thyra/integrations/:name", wrapHandler(chain.Then(http.HandlerFunc(integrations.Put))))
	gw.DefaultRouter.DELETE("/thyra/integrations/:name", wrapHandler(chain.Then(http.HandlerFunc(integrations.Delete))))
	gw.DefaultRouter.GET("/thyra/integrations/:name", wrapHandler(chain.Then(http.HandlerFunc(integrations.Get))))
	gw.DefaultRouter.GET("/thyra/integrations", wrapHandler(chain.Then(http.HandlerFunc(integrations.List))))
	gw.DefaultRouter.POST("/thyra/integrations", wrapHandler(chain.Then(http.HandlerFunc(integrations.Post))))

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
