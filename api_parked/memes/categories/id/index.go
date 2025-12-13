package handler

import (
	"net/http"

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
		getCategoryByID(w, r, id, memeService)
	case http.MethodDelete:
		deleteCategory(w, r, id, memeService)
	default:
		hutils.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func getCategoryByID(w http.ResponseWriter, r *http.Request, id string, memeService *service.MemeService) {
	category, err := memeService.GetCategoryWithMemes(id)
	if err != nil {
		hutils.HandleError(w, err)
		return
	}
	hutils.WriteJSON(w, http.StatusOK, category)
}

func deleteCategory(w http.ResponseWriter, r *http.Request, id string, memeService *service.MemeService) {
	if !hutils.CheckAuth(r) {
		hutils.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if err := memeService.DeleteCategory(id); err != nil {
		hutils.HandleError(w, err)
		return
	}

	hutils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Category deleted successfully"})
}
