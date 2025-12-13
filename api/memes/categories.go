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
	case http.MethodGet:
		getAllCategories(w, r, memeService)
	case http.MethodPost:
		createCategory(w, r, memeService)
	default:
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func getAllCategories(w http.ResponseWriter, r *http.Request, memeService *service.MemeService) {
	categories, err := memeService.GetAllCategoriesWithMemes()
	if err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, categories)
}

func createCategory(w http.ResponseWriter, r *http.Request, memeService *service.MemeService) {
	if !checkAuth(r) {
		writeError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req dto.CreateMemeCategoryRequest
	if !validateRequest(w, r, &req) {
		return
	}

	category, err := memeService.CreateCategory(req)
	if err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, category)
}
