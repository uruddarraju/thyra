package restapis

import (
	"net/http"
	"net/url"

	"github.com/golang/glog"
	"github.com/vulcand/oxy/forward"
)

func RestAPIHandler(w http.ResponseWriter, r *http.Request) {
	fwd, _ := forward.New()
	url := &url.URL{
		Scheme: "http",
		Host:   "www.ebay.com",
	}
	r.URL = url
	r.RequestURI = url.Path
	glog.Infof("%s", url.String())
	glog.Infof("%s", r.Header)
	glog.Infof("%s", r.RequestURI)
	fwd.ServeHTTP(w, r)
}
