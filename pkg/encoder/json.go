package encoder

import (
	"context"
	"encoding/json"
	"net/http"
)

func EncodeJSONResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
