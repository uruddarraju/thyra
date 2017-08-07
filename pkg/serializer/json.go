package serializer

import (
	"encoding/json"
	"net/http"

	"github.com/uruddarraju/thyra/pkg/api/runtime"
)

func Decode(req *http.Request, into runtime.Object) (runtime.Object, error) {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(into)
	if err != nil {
		return nil, err
	}

	return into, nil
}