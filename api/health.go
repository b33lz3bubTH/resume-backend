package handler

import (
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	setCORS(w)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
