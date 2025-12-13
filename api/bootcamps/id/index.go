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

	bootcampService, _, _, _, _ := hutils.GetServices(db)

	switch r.Method {
	case http.MethodGet:
		getBootcampByID(w, r, id, bootcampService)
	case http.MethodPut:
		updateBootcamp(w, r, id, bootcampService)
	case http.MethodDelete:
		deleteBootcamp(w, r, id, bootcampService)
	default:
		hutils.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func getBootcampByID(w http.ResponseWriter, r *http.Request, id string, bootcampService *service.BootcampService) {
	bootcamp, err := bootcampService.GetByID(id)
	if err != nil {
		hutils.HandleError(w, err)
		return
	}
	hutils.WriteJSON(w, http.StatusOK, bootcamp)
}

func updateBootcamp(w http.ResponseWriter, r *http.Request, id string, bootcampService *service.BootcampService) {
	if !hutils.CheckAuth(r) {
		hutils.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req dto.UpdateBootcampRequest
	if !hutils.ValidateRequest(w, r, &req) {
		return
	}

	bootcamp, err := bootcampService.Update(id, req)
	if err != nil {
		hutils.HandleError(w, err)
		return
	}

	hutils.WriteJSON(w, http.StatusOK, bootcamp)
}

func deleteBootcamp(w http.ResponseWriter, r *http.Request, id string, bootcampService *service.BootcampService) {
	if !hutils.CheckAuth(r) {
		hutils.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if err := bootcampService.Delete(id); err != nil {
		hutils.HandleError(w, err)
		return
	}

	hutils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Bootcamp deleted successfully"})
}
