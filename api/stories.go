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

	_, _, _, storyService, _ := getServices(db)

	switch r.Method {
	case http.MethodGet:
		getAllStories(w, r, storyService)
	case http.MethodPost:
		createStory(w, r, storyService)
	default:
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func getAllStories(w http.ResponseWriter, r *http.Request, storyService *service.StoryService) {
	stories, err := storyService.GetAll()
	if err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, stories)
}

func createStory(w http.ResponseWriter, r *http.Request, storyService *service.StoryService) {
	if !checkAuth(r) {
		writeError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req dto.CreateStoryRequest
	if !validateRequest(w, r, &req) {
		return
	}

	story, err := storyService.Create(req)
	if err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, story)
}
