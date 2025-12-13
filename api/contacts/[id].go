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

	_, _, _, _, contactService := getServices(db)

	switch r.Method {
	case http.MethodGet:
		getContactByID(w, r, id, contactService)
	case http.MethodDelete:
		deleteContact(w, r, id, contactService)
	default:
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func getContactByID(w http.ResponseWriter, r *http.Request, id string, contactService *service.ContactService) {
	if !checkAuth(r) {
		writeError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	contact, err := contactService.GetByID(id)
	if err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, contact)
}

func deleteContact(w http.ResponseWriter, r *http.Request, id string, contactService *service.ContactService) {
	if !checkAuth(r) {
		writeError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if err := contactService.Delete(id); err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "Contact deleted successfully"})
}
