package handler

import (
	"io"
	"net/http"

	hutils "resume-backend/pkg/handler"
	"resume-backend/pkg/service"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if hutils.HandleCORS(w, r) {
		return
	}

	if r.Method != http.MethodPost {
		hutils.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed. Only POST is supported")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		hutils.WriteError(w, http.StatusBadRequest, "Failed to read request body")
		return
	}
	defer r.Body.Close()

	openRouterService := service.NewOpenRouterService()
	response, statusCode, err := openRouterService.CreateChatCompletion(body, r.Header.Get("Referer"))
	if err != nil {
		hutils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	hutils.WriteJSON(w, statusCode, response)
}

