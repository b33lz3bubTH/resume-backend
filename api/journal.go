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

	_, journalService, _, _, _ := hutils.GetServices(db)

	switch r.Method {
	case http.MethodGet:
		getAllJournalEntries(w, r, journalService)
	case http.MethodPost:
		createJournalEntry(w, r, journalService)
	default:
		hutils.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func getAllJournalEntries(w http.ResponseWriter, r *http.Request, journalService *service.JournalService) {
	entries, err := journalService.GetAll()
	if err != nil {
		hutils.HandleError(w, err)
		return
	}
	hutils.WriteJSON(w, http.StatusOK, entries)
}

func createJournalEntry(w http.ResponseWriter, r *http.Request, journalService *service.JournalService) {
	if !hutils.CheckAuth(r) {
		hutils.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req dto.CreateJournalRequest
	if !hutils.ValidateRequest(w, r, &req) {
		return
	}

	entry, err := journalService.Create(req)
	if err != nil {
		hutils.HandleError(w, err)
		return
	}

	hutils.WriteJSON(w, http.StatusCreated, entry)
}
