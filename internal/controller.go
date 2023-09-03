package internal

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func proxy(w http.ResponseWriter, r *http.Request) {
	var pr proxyRequest
	if err := json.NewDecoder(r.Body).Decode(&pr); err != nil {
		log.Errorf("REQUEST_BODY_DECODE_ERROR", err)
	}
	prJson, err := json.Marshal(pr)
	if err != nil {
		log.Errorf("JSON_MARSHAL_ERROR", err)
	}
	handleProxyRequest(r.Context(), string(prJson))
}
