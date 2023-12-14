package until

import (
	"encoding/json"
	"net/http"
)

func decodeJSON(r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}
