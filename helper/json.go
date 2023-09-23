package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSONEncode(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json := json.NewEncoder(w)
	err := json.Encode(v)
	if err != nil {
		log.Fatal(err)
	}
}

func BodyRequest(r *http.Request, v any) {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		log.Fatal(err)
	}
}
