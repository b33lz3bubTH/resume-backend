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

	bootcampService, _, _, _, _ := getServices(db)

	switch r.Method {
	case http.MethodGet:
		getAllBootcamps(w, r, bootcampService)
	case http.MethodPost:
		createBootcamp(w, r, bootcampService)
	default:
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func getAllBootcamps(w http.ResponseWriter, r *http.Request, bootcampService *service.BootcampService) {
	bootcamps, err := bootcampService.GetAll()
	if err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, bootcamps)
}

func createBootcamp(w http.ResponseWriter, r *http.Request, bootcampService *service.BootcampService) {
	if !checkAuth(r) {
		writeError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req dto.CreateBootcampRequest
	if !validateRequest(w, r, &req) {
		return
	}

	bootcamp, err := bootcampService.Create(req)
	if err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, bootcamp)
}
