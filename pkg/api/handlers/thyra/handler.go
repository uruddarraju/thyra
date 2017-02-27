package thyra

import (
	"encoding/json"
	"net/http"
)

type APIResourceList struct {
	Group     string
	Resources []string
	Versions  []string
}

var thyraGroup APIResourceList = APIResourceList{
	Group:     "thyra",
	Resources: []string{"restapis", "resources", "methods", "stages", "deployments", "integrations", "authorizers"},
	Versions:  []string{"", "v1"},
}

func Get(w http.ResponseWriter, r *http.Request) {
	js, err := json.MarshalIndent(thyraGroup, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
