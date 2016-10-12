package gateway

import (
	"log"
	"net/http"
	"time"

	"github.com/golang/glog"
	"github.com/justinas/alice"
	"github.com/uruddarraju/thyra/pkg/api/handlers/restapis"
	"github.com/uruddarraju/thyra/pkg/auth/authn"
	"github.com/uruddarraju/thyra/pkg/auth/authn/tokenfile"
	"github.com/julienschmidt/httprouter"
)

type Gateway struct {
	Address       string
	DefaultRouter *httprouter.Router
	Server        *http.Server
	Authenticator authn.Authenticator
}

type GatewayOpts struct{
	KeystoneAuthnConfigFile string
	BasicAuthnFile 		string
}

func NewDefaultGateway() *Gateway {
	defaultRouter := httprouter.New()
	authn := tokenfile.NewTokenAuthenticator("")
	AddDefaultHandlers(defaultRouter, authn)
	srv := &http.Server{
		Handler:      defaultRouter,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return &Gateway{
		DefaultRouter: defaultRouter,
		Server:        srv,
	}
}

func (gw *Gateway) Start() {
	gw.Server.ListenAndServe()
	glog.Fatalf("Server quit.....")
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Thyra....!\n"))
}

func AddDefaultHandlers(router *httprouter.Router, authenticator authn.Authenticator, ) {
	chain := alice.New(authenticator.Authenticator, middlewareTwo)
	router.GET("/", wrapHandler(chain.Then(http.HandlerFunc(restapis.RestAPIHandler))))
	router.GET("/hello", wrapHandler(chain.Then(http.HandlerFunc(HelloHandler))))
	router.GET("/metrics", wrapHandler(chain.Then(http.HandlerFunc(HelloHandler))))
	router.GET("/healthz", wrapHandler(chain.Then(http.HandlerFunc(HelloHandler))))
	router.GET("/restapis", wrapHandler(chain.Then(http.HandlerFunc(restapis.RestAPIHandler))))
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

// Essentially we are just taking the params and shoving them into the context and
// returning a proper httprouter.Handle.
func wrapHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request,
	ps httprouter.Params) {
		h.ServeHTTP(w, r)
	}
}
