package gateway

import (
	"context"
	"net/http"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"github.com/uruddarraju/thyra/pkg/api/handlers/restapis"
	"github.com/uruddarraju/thyra/pkg/api/server"
	"github.com/uruddarraju/thyra/pkg/auth/authn"
	"github.com/uruddarraju/thyra/pkg/auth/authn/tokenfile"
	utilruntime "k8s.io/kubernetes/pkg/util/runtime"
)

var gateway *gatewayImpl
var once sync.Once

type Gateway interface {
	Start()
}

type GatewayOpts struct {
	KeystoneAuthnConfigFile string
	BasicAuthnFile          string
}

type gatewayImpl struct {
	Address       string
	DefaultRouter *httprouter.Router
	Server        *http.Server
	Authenticator authn.Authenticator
}

func DefaultGateway() Gateway {
	once.Do(func() {
		defaultRouter := httprouter.New()
		authn := tokenfile.NewTokenAuthenticator("")
		AddDefaultHandlers(defaultRouter, authn)
		srv := &http.Server{
			Handler:      defaultRouter,
			Addr:         "127.0.0.1:8000",
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}

		gateway = &gatewayImpl{
			DefaultRouter: defaultRouter,
			Server:        srv,
		}
	})
	return gateway

}

func (gw *gatewayImpl) Start() {

	defer utilruntime.HandleCrash()
	for {
		server.InitGatewayServer(gw.Server)
		if err := gw.Server.ListenAndServe(); err != nil {
			log.Errorf("Unable to listen for server (%v); will try again.", err)
		}
		time.Sleep(15 * time.Second)
	}
	log.Fatalf("Server quit.....")
}

func AddDefaultHandlers(router *httprouter.Router, authenticator authn.Authenticator) {

	// TODO: Add Union of authenticators
	chain := alice.New(authenticator.Authenticator, middlewareTwo)
	router.GET("/", wrapHandler(chain.Then(http.HandlerFunc(restapis.Get))))
	router.GET("/hello", wrapHandler(chain.Then(http.HandlerFunc(HelloHandler))))
	router.GET("/metrics", wrapHandler(chain.Then(http.HandlerFunc(HelloHandler))))
	router.GET("/healthz", wrapHandler(chain.Then(http.HandlerFunc(HelloHandler))))

	router.GET("/restapis", wrapHandler(chain.Then(http.HandlerFunc(restapis.List))))
	router.POST("/restapis", wrapHandler(chain.Then(http.HandlerFunc(restapis.Post))))

}

func middlewareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("Executing middlewareTwo")
		if r.URL.Path != "/" {
			return
		}
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
