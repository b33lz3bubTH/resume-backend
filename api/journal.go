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

	_, journalService, _, _, _ := getServices(db)

	switch r.Method {
	case http.MethodGet:
		getAllJournalEntries(w, r, journalService)
	case http.MethodPost:
		createJournalEntry(w, r, journalService)
	default:
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func getAllJournalEntries(w http.ResponseWriter, r *http.Request, journalService *service.JournalService) {
	entries, err := journalService.GetAll()
	if err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, entries)
}

func createJournalEntry(w http.ResponseWriter, r *http.Request, journalService *service.JournalService) {
	if !checkAuth(r) {
		writeError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req dto.CreateJournalRequest
	if !validateRequest(w, r, &req) {
		return
	}

	entry, err := journalService.Create(req)
	if err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, entry)
}
