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

	_, _, _, _, contactService := hutils.GetServices(db)

	switch r.Method {
	case http.MethodGet:
		getContactByID(w, r, id, contactService)
	case http.MethodDelete:
		deleteContact(w, r, id, contactService)
	default:
		hutils.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func getContactByID(w http.ResponseWriter, r *http.Request, id string, contactService *service.ContactService) {
	if !hutils.CheckAuth(r) {
		hutils.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	contact, err := contactService.GetByID(id)
	if err != nil {
		hutils.HandleError(w, err)
		return
	}
	hutils.WriteJSON(w, http.StatusOK, contact)
}

func deleteContact(w http.ResponseWriter, r *http.Request, id string, contactService *service.ContactService) {
	if !hutils.CheckAuth(r) {
		hutils.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if err := contactService.Delete(id); err != nil {
		hutils.HandleError(w, err)
		return
	}

	hutils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Contact deleted successfully"})
}
