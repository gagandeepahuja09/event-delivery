package internal

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func proxy(w http.ResponseWriter, r *http.Request) {
	var pr proxyRequest
	if err := json.NewDecoder(r.Body).Decode(&pr); err != nil {
		log.Error("REQUEST_BODY_DECODE_ERROR", err)
	}
	err := handleProxyRequest(r.Context(), pr)
	Respond(r, w, map[string]interface{}{"successfully_queued": true}, err)
}

func Respond(r *http.Request, w http.ResponseWriter, payload interface{}, err error) {
	w.Header().Set("Content-Type", "application/json")

	var res []byte
	if err == nil {
		w.WriteHeader(200)
		response := payload
		res, _ = json.Marshal(response)
	} else {
		w.WriteHeader(err.Code())
		res, _ = json.Marshal(errorResp{Error: err.Error()})
	}

	_, berr := w.Write(res)
	if berr != nil {
		log.Error("RESPONSE_WRITE_FAILED")
	}
}
