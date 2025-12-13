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

	bootcampService, _, _, _, _ := hutils.GetServices(db)

	switch r.Method {
	case http.MethodGet:
		getAllBootcamps(w, r, bootcampService)
	case http.MethodPost:
		createBootcamp(w, r, bootcampService)
	default:
		hutils.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func getAllBootcamps(w http.ResponseWriter, r *http.Request, bootcampService *service.BootcampService) {
	bootcamps, err := bootcampService.GetAll()
	if err != nil {
		hutils.HandleError(w, err)
		return
	}
	hutils.WriteJSON(w, http.StatusOK, bootcamps)
}

func createBootcamp(w http.ResponseWriter, r *http.Request, bootcampService *service.BootcampService) {
	if !hutils.CheckAuth(r) {
		hutils.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req dto.CreateBootcampRequest
	if !hutils.ValidateRequest(w, r, &req) {
		return
	}

	bootcamp, err := bootcampService.Create(req)
	if err != nil {
		hutils.HandleError(w, err)
		return
	}

	hutils.WriteJSON(w, http.StatusCreated, bootcamp)
}
