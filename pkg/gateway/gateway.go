package gateway

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/golang/glog"
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
	go func() {

	}()
	go func() {
		defer utilruntime.HandleCrash()
		for {
			server.InitGatewayServer(gw.Server)
			if err := gw.Server.ListenAndServe(); err != nil {
				glog.Errorf("Unable to listen for server (%v); will try again.", err)
			}
			time.Sleep(15 * time.Second)
		}
	}()
	glog.Fatalf("Server quit.....")
}

func AddDefaultHandlers(router *httprouter.Router, authenticator authn.Authenticator) {
	chain := alice.New(authenticator.Authenticator, middlewareTwo)
	router.GET("/", wrapHandler(chain.Then(http.HandlerFunc(restapis.RestAPIHandler))))
	router.GET("/hello", wrapHandler(chain.Then(http.HandlerFunc(HelloHandler))))
	router.GET("/metrics", wrapHandler(chain.Then(http.HandlerFunc(HelloHandler))))
	router.GET("/healthz", wrapHandler(chain.Then(http.HandlerFunc(HelloHandler))))

	router.GET("/restapis", wrapHandler(chain.Then(http.HandlerFunc(restapis.RestAPIHandler))))
	router.PUT("/restapis", wrapHandler(chain.Then(http.HandlerFunc(restapis.RestAPIHandler))))
	router.POST("/restapis", wrapHandler(chain.Then(http.HandlerFunc(restapis.RestAPIHandler))))
	router.DELETE("/restapis", wrapHandler(chain.Then(http.HandlerFunc(restapis.RestAPIHandler))))
}

func middlewareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareTwo")
		if r.URL.Path != "/" {
			return
		}
		next.ServeHTTP(w, r)
		log.Println("Executing middlewareTwo again")
	})
}

func wrapHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request,
		ps httprouter.Params) {
		h.ServeHTTP(w, r)
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Thyra....!\n"))
}
