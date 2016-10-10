package gateway

import (
	"net/http"
	"time"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
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

	r.HandleFunc("/", HelloHandler)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
	glog.Fatalf("Server quit.....")
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Thyra....!\n"))
}
