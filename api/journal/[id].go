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

	_, journalService, _, _, _ := hutils.GetServices(db)

	switch r.Method {
	case http.MethodGet:
		getJournalEntryByID(w, r, id, journalService)
	case http.MethodPut:
		updateJournalEntry(w, r, id, journalService)
	case http.MethodDelete:
		deleteJournalEntry(w, r, id, journalService)
	default:
		hutils.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func getJournalEntryByID(w http.ResponseWriter, r *http.Request, id string, journalService *service.JournalService) {
	entry, err := journalService.GetByID(id)
	if err != nil {
		hutils.HandleError(w, err)
		return
	}
	hutils.WriteJSON(w, http.StatusOK, entry)
}

func updateJournalEntry(w http.ResponseWriter, r *http.Request, id string, journalService *service.JournalService) {
	if !hutils.CheckAuth(r) {
		hutils.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req dto.UpdateJournalRequest
	if !hutils.ValidateRequest(w, r, &req) {
		return
	}

	entry, err := journalService.Update(id, req)
	if err != nil {
		hutils.HandleError(w, err)
		return
	}

	hutils.WriteJSON(w, http.StatusOK, entry)
}

func deleteJournalEntry(w http.ResponseWriter, r *http.Request, id string, journalService *service.JournalService) {
	if !hutils.CheckAuth(r) {
		hutils.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if err := journalService.Delete(id); err != nil {
		hutils.HandleError(w, err)
		return
	}

	hutils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Journal entry deleted successfully"})
}
