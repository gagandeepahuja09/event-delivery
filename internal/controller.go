package internal

import (
	"encoding/json"
	"net/http"
)

func proxy(w http.ResponseWriter, r *http.Request) {
	var pr proxyRequest
	json.NewDecoder(r.Body).Decode(&pr)
	handleProxyRequest(r.Context(), pr)
}
