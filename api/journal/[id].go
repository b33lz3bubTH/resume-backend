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

	_, journalService, _, _, _ := getServices(db)

	switch r.Method {
	case http.MethodGet:
		getJournalEntryByID(w, r, id, journalService)
	case http.MethodPut:
		updateJournalEntry(w, r, id, journalService)
	case http.MethodDelete:
		deleteJournalEntry(w, r, id, journalService)
	default:
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func getJournalEntryByID(w http.ResponseWriter, r *http.Request, id string, journalService *service.JournalService) {
	entry, err := journalService.GetByID(id)
	if err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, entry)
}

func updateJournalEntry(w http.ResponseWriter, r *http.Request, id string, journalService *service.JournalService) {
	if !checkAuth(r) {
		writeError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req dto.UpdateJournalRequest
	if !validateRequest(w, r, &req) {
		return
	}

	entry, err := journalService.Update(id, req)
	if err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, entry)
}

func deleteJournalEntry(w http.ResponseWriter, r *http.Request, id string, journalService *service.JournalService) {
	if !checkAuth(r) {
		writeError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if err := journalService.Delete(id); err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "Journal entry deleted successfully"})
}
