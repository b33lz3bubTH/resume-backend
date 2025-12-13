package handler

import (
	"net/http"

	"resume-backend/dto"
	"resume-backend/pkg/service"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if handleCORS(w, r) {
		return
	}

	db, err := getDB()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Database connection failed")
		return
	}
	defer db.Close()

	_, _, memeService, _, _ := getServices(db)

	switch r.Method {
	case http.MethodPost:
		createMeme(w, r, memeService)
	default:
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func createMeme(w http.ResponseWriter, r *http.Request, memeService *service.MemeService) {
	if !checkAuth(r) {
		writeError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req dto.CreateMemeRequest
	if !validateRequest(w, r, &req) {
		return
	}

	meme, err := memeService.CreateMeme(req)
	if err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, meme)
}
