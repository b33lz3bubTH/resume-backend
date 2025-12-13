package handler

import (
	"net/http"

	"resume-backend/dto"
	hutils "resume-backend/pkg/handler"
	"resume-backend/pkg/service"
)

// ResourceType represents the type of resource being accessed
type ResourceType string

const (
	ResourceTypeBootcamp ResourceType = "bootcamps"
	ResourceTypeJournal  ResourceType = "journal"
	ResourceTypeMeme     ResourceType = "memes"
	ResourceTypeCategory ResourceType = "categories"
)

// isValidResourceType checks if the resource type is valid
func isValidResourceType(rt ResourceType) bool {
	return rt == ResourceTypeBootcamp || rt == ResourceTypeJournal ||
		rt == ResourceTypeMeme || rt == ResourceTypeCategory
}

// Handler handles individual resource operations by ID
func Handler(w http.ResponseWriter, r *http.Request) {
	if hutils.HandleCORS(w, r) {
		return
	}

	// Get ID from query parameter (set by Vercel rewrite) or path
	id := r.URL.Query().Get("id")
	if id == "" {
		id = hutils.GetIDFromPath(r)
	}
	if id == "" {
		hutils.WriteError(w, http.StatusBadRequest, "ID is required")
		return
	}

	// Get resource type from query parameter
	resourceType := ResourceType(r.URL.Query().Get("resource"))
	if resourceType == "" {
		hutils.WriteError(w, http.StatusBadRequest, "resource parameter is required. Valid values: bootcamps, journal, memes, categories")
		return
	}

	// Validate resource type
	if !isValidResourceType(resourceType) {
		hutils.WriteError(w, http.StatusBadRequest, "invalid resource type. Valid values: bootcamps, journal, memes, categories")
		return
	}

	// Initialize database connection
	db, err := hutils.GetDB()
	if err != nil {
		hutils.WriteError(w, http.StatusInternalServerError, "Database connection failed")
		return
	}
	defer db.Close()

	// Get services
	bootcampService, journalService, memeService, _, _ := hutils.GetServices(db)

	// Route to appropriate handler based on resource type and HTTP method
	switch resourceType {
	case ResourceTypeBootcamp:
		handleBootcampItem(w, r, id, bootcampService)
	case ResourceTypeJournal:
		handleJournalItem(w, r, id, journalService)
	case ResourceTypeMeme:
		handleMemeItem(w, r, id, memeService)
	case ResourceTypeCategory:
		handleCategoryItem(w, r, id, memeService)
	default:
		hutils.WriteError(w, http.StatusBadRequest, "unsupported resource type")
	}
}

// handleBootcampItem handles individual bootcamp operations
func handleBootcampItem(w http.ResponseWriter, r *http.Request, id string, service *service.BootcampService) {
	switch r.Method {
	case http.MethodGet:
		getBootcampByID(w, r, id, service)
	case http.MethodPut:
		updateBootcamp(w, r, id, service)
	case http.MethodDelete:
		deleteBootcamp(w, r, id, service)
	default:
		hutils.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed. Supported methods: GET, PUT, DELETE")
	}
}

// handleJournalItem handles individual journal entry operations
func handleJournalItem(w http.ResponseWriter, r *http.Request, id string, service *service.JournalService) {
	switch r.Method {
	case http.MethodGet:
		getJournalEntryByID(w, r, id, service)
	case http.MethodPut:
		updateJournalEntry(w, r, id, service)
	case http.MethodDelete:
		deleteJournalEntry(w, r, id, service)
	default:
		hutils.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed. Supported methods: GET, PUT, DELETE")
	}
}

// handleMemeItem handles individual meme operations
func handleMemeItem(w http.ResponseWriter, r *http.Request, id string, service *service.MemeService) {
	switch r.Method {
	case http.MethodGet:
		getMemeByID(w, r, id, service)
	case http.MethodPut:
		updateMeme(w, r, id, service)
	case http.MethodDelete:
		deleteMeme(w, r, id, service)
	default:
		hutils.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed. Supported methods: GET, PUT, DELETE")
	}
}

// handleCategoryItem handles individual category operations
func handleCategoryItem(w http.ResponseWriter, r *http.Request, id string, service *service.MemeService) {
	switch r.Method {
	case http.MethodGet:
		getCategoryByID(w, r, id, service)
	case http.MethodDelete:
		deleteCategory(w, r, id, service)
	default:
		hutils.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed. Supported methods: GET, DELETE")
	}
}

// Bootcamp item handlers
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

// Journal item handlers
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

// Meme item handlers
func getMemeByID(w http.ResponseWriter, r *http.Request, id string, memeService *service.MemeService) {
	meme, err := memeService.GetMemeByID(id)
	if err != nil {
		hutils.HandleError(w, err)
		return
	}
	hutils.WriteJSON(w, http.StatusOK, meme)
}

func updateMeme(w http.ResponseWriter, r *http.Request, id string, memeService *service.MemeService) {
	if !hutils.CheckAuth(r) {
		hutils.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req dto.UpdateMemeRequest
	if !hutils.ValidateRequest(w, r, &req) {
		return
	}

	meme, err := memeService.UpdateMeme(id, req)
	if err != nil {
		hutils.HandleError(w, err)
		return
	}

	hutils.WriteJSON(w, http.StatusOK, meme)
}

func deleteMeme(w http.ResponseWriter, r *http.Request, id string, memeService *service.MemeService) {
	if !hutils.CheckAuth(r) {
		hutils.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if err := memeService.DeleteMeme(id); err != nil {
		hutils.HandleError(w, err)
		return
	}

	hutils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Meme deleted successfully"})
}

// Category item handlers
func getCategoryByID(w http.ResponseWriter, r *http.Request, id string, memeService *service.MemeService) {
	category, err := memeService.GetCategoryWithMemes(id)
	if err != nil {
		hutils.HandleError(w, err)
		return
	}
	hutils.WriteJSON(w, http.StatusOK, category)
}

func deleteCategory(w http.ResponseWriter, r *http.Request, id string, memeService *service.MemeService) {
	if !hutils.CheckAuth(r) {
		hutils.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if err := memeService.DeleteCategory(id); err != nil {
		hutils.HandleError(w, err)
		return
	}

	hutils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Category deleted successfully"})
}

