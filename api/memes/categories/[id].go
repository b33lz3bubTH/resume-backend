package handler

import (
	"net/http"

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
		getCategoryByID(w, r, id, memeService)
	case http.MethodDelete:
		deleteCategory(w, r, id, memeService)
	default:
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func getCategoryByID(w http.ResponseWriter, r *http.Request, id string, memeService *service.MemeService) {
	category, err := memeService.GetCategoryWithMemes(id)
	if err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, category)
}

func deleteCategory(w http.ResponseWriter, r *http.Request, id string, memeService *service.MemeService) {
	if !checkAuth(r) {
		writeError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if err := memeService.DeleteCategory(id); err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "Category deleted successfully"})
}
