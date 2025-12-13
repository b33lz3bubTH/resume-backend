package handler

import (
	"net/http"
	"strconv"

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

	_, _, _, _, contactService := getServices(db)

	switch r.Method {
	case http.MethodGet:
		getAllContacts(w, r, contactService)
	case http.MethodPost:
		createContact(w, r, contactService)
	default:
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func getAllContacts(w http.ResponseWriter, r *http.Request, contactService *service.ContactService) {
	if !checkAuth(r) {
		writeError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	page := 1
	pageSize := 20

	if pageStr := r.URL.Query().Get("page"); pageStr != "" {
		if parsedPage, err := strconv.Atoi(pageStr); err == nil && parsedPage > 0 {
			page = parsedPage
		}
	}

	if pageSizeStr := r.URL.Query().Get("page_size"); pageSizeStr != "" {
		if parsedPageSize, err := strconv.Atoi(pageSizeStr); err == nil && parsedPageSize > 0 && parsedPageSize <= 100 {
			pageSize = parsedPageSize
		}
	}

	contacts, total, err := contactService.GetAll(page, pageSize)
	if err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"contacts":    contacts,
		"total":       total,
		"page":        page,
		"page_size":   pageSize,
		"total_pages": (total + pageSize - 1) / pageSize,
	})
}

func createContact(w http.ResponseWriter, r *http.Request, contactService *service.ContactService) {
	var req dto.CreateContactRequest
	if !validateRequest(w, r, &req) {
		return
	}

	contact, err := contactService.Create(req)
	if err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, contact)
}
