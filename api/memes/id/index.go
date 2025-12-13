package handler

import (
	"net/http"

	"resume-backend/dto"
	hutils "resume-backend/pkg/handler"
	"resume-backend/pkg/service"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if hutils.HandleCORS(w, r) {
		return
	}

	id := hutils.GetIDFromPath(r)
	if id == "" {
		hutils.WriteError(w, http.StatusBadRequest, "ID is required")
		return
	}

	db, err := hutils.GetDB()
	if err != nil {
		hutils.WriteError(w, http.StatusInternalServerError, "Database connection failed")
		return
	}
	defer db.Close()

	_, _, memeService, _, _ := hutils.GetServices(db)

	switch r.Method {
	case http.MethodGet:
		getMemeByID(w, r, id, memeService)
	case http.MethodPut:
		updateMeme(w, r, id, memeService)
	case http.MethodDelete:
		deleteMeme(w, r, id, memeService)
	default:
		hutils.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func getMemeByID(w http.ResponseWriter, r *http.Request, id string, memeService *service.MemeService) {
	meme, err := memeService.GetMemeByID(id)
	if err != nil {
		hutils.HandleError(w, err)
		return
	}
	hutils.WriteJSON(w, http.StatusOK, meme)
}

func updateMeme(w http.ResponseWriter, r *http.Request, id string, memeService *service.MemeService) {
	if !hutils.CheckAuth(r) {
		hutils.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req dto.UpdateMemeRequest
	if !hutils.ValidateRequest(w, r, &req) {
		return
	}

	meme, err := memeService.UpdateMeme(id, req)
	if err != nil {
		hutils.HandleError(w, err)
		return
	}

	hutils.WriteJSON(w, http.StatusOK, meme)
}

func deleteMeme(w http.ResponseWriter, r *http.Request, id string, memeService *service.MemeService) {
	if !hutils.CheckAuth(r) {
		hutils.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if err := memeService.DeleteMeme(id); err != nil {
		hutils.HandleError(w, err)
		return
	}

	hutils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Meme deleted successfully"})
}
