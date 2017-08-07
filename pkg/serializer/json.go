package serializer

import (
	"encoding/json"
	"net/http"

	"github.com/uruddarraju/thyra/pkg/api/runtime"
	"github.com/uruddarraju/thyra/pkg/apis/thyra/v1"
)

func Decode(req *http.Request) (runtime.Object, error) {
	decoder := json.NewDecoder(req.Body)
	var meta v1.ObjectMeta
	err := decoder.Decode(meta)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
