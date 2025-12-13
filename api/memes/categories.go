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

	db, err := hutils.GetDB()
	if err != nil {
		hutils.WriteError(w, http.StatusInternalServerError, "Database connection failed")
		return
	}
	defer db.Close()

	_, _, memeService, _, _ := hutils.GetServices(db)

	switch r.Method {
	case http.MethodGet:
		getAllCategories(w, r, memeService)
	case http.MethodPost:
		createCategory(w, r, memeService)
	default:
		hutils.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func getAllCategories(w http.ResponseWriter, r *http.Request, memeService *service.MemeService) {
	categories, err := memeService.GetAllCategoriesWithMemes()
	if err != nil {
		hutils.HandleError(w, err)
		return
	}
	hutils.WriteJSON(w, http.StatusOK, categories)
}

func createCategory(w http.ResponseWriter, r *http.Request, memeService *service.MemeService) {
	if !hutils.CheckAuth(r) {
		hutils.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req dto.CreateMemeCategoryRequest
	if !hutils.ValidateRequest(w, r, &req) {
		return
	}

	category, err := memeService.CreateCategory(req)
	if err != nil {
		hutils.HandleError(w, err)
		return
	}

	hutils.WriteJSON(w, http.StatusCreated, category)
}
