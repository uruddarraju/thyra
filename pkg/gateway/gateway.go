package gateway

import (
	"net/http"
	"net/url"
	"time"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/vulcand/oxy/forward"
)

type Gateway struct {
	Address string
}

func NewDefaultGateway() *Gateway {
	return &Gateway{}
}

type GatewayOpts struct{}

func (gw *Gateway) Start(opts *GatewayOpts) {
	r := mux.NewRouter()

	r.HandleFunc("/", InfoHandler)
	r.HandleFunc("/hello", HelloHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
	glog.Fatalf("Server quit.....")
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Thyra....!\n"))
}

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	fwd, _ := forward.New()
	url := &url.URL{
		Scheme: "http",
		Host:   "www.ebay.com",
	}
	r.URL = url
	glog.Infof("%s", url.String())
	glog.Infof("%s", r.Header)
	fwd.ServeHTTP(w, r)
}
