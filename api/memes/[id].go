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

	id := getIDFromPath(r)
	if id == "" {
		writeError(w, http.StatusBadRequest, "ID is required")
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
	case http.MethodGet:
		getMemeByID(w, r, id, memeService)
	case http.MethodPut:
		updateMeme(w, r, id, memeService)
	case http.MethodDelete:
		deleteMeme(w, r, id, memeService)
	default:
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func getMemeByID(w http.ResponseWriter, r *http.Request, id string, memeService *service.MemeService) {
	meme, err := memeService.GetMemeByID(id)
	if err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, meme)
}

func updateMeme(w http.ResponseWriter, r *http.Request, id string, memeService *service.MemeService) {
	if !checkAuth(r) {
		writeError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req dto.UpdateMemeRequest
	if !validateRequest(w, r, &req) {
		return
	}

	meme, err := memeService.UpdateMeme(id, req)
	if err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, meme)
}

func deleteMeme(w http.ResponseWriter, r *http.Request, id string, memeService *service.MemeService) {
	if !checkAuth(r) {
		writeError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if err := memeService.DeleteMeme(id); err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "Meme deleted successfully"})
}
