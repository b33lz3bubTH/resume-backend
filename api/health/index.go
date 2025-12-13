package handler

import (
	"net/http"

	hutils "resume-backend/pkg/handler"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	hutils.SetCORS(w)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
