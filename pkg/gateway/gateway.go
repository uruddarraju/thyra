package gateway

import (
	"log"
	"net/http"
	"time"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/uruddarraju/thyra/pkg/api/handlers/restapis"
	"github.com/uruddarraju/thyra/pkg/auth/authn"
	"github.com/uruddarraju/thyra/pkg/auth/authn/union"
)

type Gateway struct {
	Address       string
	DefaultRouter *mux.Router
	Server        *http.Server
	Authenticator authn.Authenticator
}

type GatewayOpts struct{}

func NewDefaultGateway() *Gateway {
	defaultRouter := mux.NewRouter()
	authn, err := union.NewUnionAuthenticator("keystone-url", "token-file")
	if err != nil {
		glog.Errorf("Initializing Union Authentication Failed: %s", err.Error())
		return nil
	}
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

func AddDefaultHandlers(router *mux.Router, authenticator authn.Authenticator) {

	chain := alice.New(authenticator.Authenticator, middlewareTwo)
	router.Handle("/", chain.Then(http.HandlerFunc(restapis.RestAPIHandler)))
	router.Handle("/hello", chain.Then(http.HandlerFunc(HelloHandler)))
	router.Handle("/metrics", chain.Then(http.HandlerFunc(HelloHandler)))
	router.Handle("/healthz", chain.Then(http.HandlerFunc(HelloHandler)))
	router.Handle("/restapis", chain.Then(http.HandlerFunc(restapis.RestAPIHandler)))
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
