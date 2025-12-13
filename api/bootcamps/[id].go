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

	bootcampService, _, _, _, _ := getServices(db)

	switch r.Method {
	case http.MethodGet:
		getBootcampByID(w, r, id, bootcampService)
	case http.MethodPut:
		updateBootcamp(w, r, id, bootcampService)
	case http.MethodDelete:
		deleteBootcamp(w, r, id, bootcampService)
	default:
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func getBootcampByID(w http.ResponseWriter, r *http.Request, id string, bootcampService *service.BootcampService) {
	bootcamp, err := bootcampService.GetByID(id)
	if err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, bootcamp)
}

func updateBootcamp(w http.ResponseWriter, r *http.Request, id string, bootcampService *service.BootcampService) {
	if !checkAuth(r) {
		writeError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req dto.UpdateBootcampRequest
	if !validateRequest(w, r, &req) {
		return
	}

	bootcamp, err := bootcampService.Update(id, req)
	if err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, bootcamp)
}

func deleteBootcamp(w http.ResponseWriter, r *http.Request, id string, bootcampService *service.BootcampService) {
	if !checkAuth(r) {
		writeError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if err := bootcampService.Delete(id); err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "Bootcamp deleted successfully"})
}
