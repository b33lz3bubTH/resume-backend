package handler

import (
	"net/http"

	"resume-backend/dto"
	hutils "resume-backend/pkg/handler"
	"resume-backend/pkg/service"
)

// Handler is the unified resource handler that dispatches requests based on resource type
func Handler(w http.ResponseWriter, r *http.Request) {
	if hutils.HandleCORS(w, r) {
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
		handleBootcampCollection(w, r, bootcampService)
	case ResourceTypeJournal:
		handleJournalCollection(w, r, journalService)
	case ResourceTypeMeme:
		handleMemeCollection(w, r, memeService)
	case ResourceTypeCategory:
		handleCategoryCollection(w, r, memeService)
	default:
		hutils.WriteError(w, http.StatusBadRequest, "unsupported resource type")
	}
}

// handleBootcampCollection handles bootcamp collection operations
func handleBootcampCollection(w http.ResponseWriter, r *http.Request, service *service.BootcampService) {
	switch r.Method {
	case http.MethodGet:
		getAllBootcamps(w, r, service)
	case http.MethodPost:
		createBootcamp(w, r, service)
	default:
		hutils.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed. Supported methods: GET, POST")
	}
}

// handleJournalCollection handles journal collection operations
func handleJournalCollection(w http.ResponseWriter, r *http.Request, service *service.JournalService) {
	switch r.Method {
	case http.MethodGet:
		getAllJournalEntries(w, r, service)
	case http.MethodPost:
		createJournalEntry(w, r, service)
	default:
		hutils.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed. Supported methods: GET, POST")
	}
}

// handleMemeCollection handles meme collection operations
func handleMemeCollection(w http.ResponseWriter, r *http.Request, service *service.MemeService) {
	switch r.Method {
	case http.MethodPost:
		createMeme(w, r, service)
	default:
		hutils.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed. Supported methods: POST")
	}
}

// handleCategoryCollection handles category collection operations
func handleCategoryCollection(w http.ResponseWriter, r *http.Request, service *service.MemeService) {
	switch r.Method {
	case http.MethodGet:
		getAllCategories(w, r, service)
	case http.MethodPost:
		createCategory(w, r, service)
	default:
		hutils.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed. Supported methods: GET, POST")
	}
}

// Bootcamp handlers
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

// Journal handlers
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

// Meme handlers
func createMeme(w http.ResponseWriter, r *http.Request, memeService *service.MemeService) {
	if !hutils.CheckAuth(r) {
		hutils.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req dto.CreateMemeRequest
	if !hutils.ValidateRequest(w, r, &req) {
		return
	}

	meme, err := memeService.CreateMeme(req)
	if err != nil {
		hutils.HandleError(w, err)
		return
	}

	hutils.WriteJSON(w, http.StatusCreated, meme)
}

// Category handlers
func getAllCategories(w http.ResponseWriter, r *http.Request, memeService *service.MemeService) {
	categories, err := memeService.GetAllCategoriesWithMemes()
	if err != nil {
		hutils.HandleError(w, err)
		return
	}
	hutils.WriteJSON(w, http.StatusOK, categories)
}

func createCategory(w http.ResponseWriter, r *http.Request, memeService *service.MemeService) {
	if !hutils.CheckAuth(r) {
		hutils.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req dto.CreateMemeCategoryRequest
	if !hutils.ValidateRequest(w, r, &req) {
		return
	}

	category, err := memeService.CreateCategory(req)
	if err != nil {
		hutils.HandleError(w, err)
		return
	}

	hutils.WriteJSON(w, http.StatusCreated, category)
}

