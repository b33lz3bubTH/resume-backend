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

	_, _, _, storyService, _ := hutils.GetServices(db)

	switch r.Method {
	case http.MethodGet:
		getAllStories(w, r, storyService)
	case http.MethodPost:
		createStory(w, r, storyService)
	default:
		hutils.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func getAllStories(w http.ResponseWriter, r *http.Request, storyService *service.StoryService) {
	stories, err := storyService.GetAll()
	if err != nil {
		hutils.HandleError(w, err)
		return
	}
	hutils.WriteJSON(w, http.StatusOK, stories)
}

func createStory(w http.ResponseWriter, r *http.Request, storyService *service.StoryService) {
	if !hutils.CheckAuth(r) {
		hutils.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req dto.CreateStoryRequest
	if !hutils.ValidateRequest(w, r, &req) {
		return
	}

	story, err := storyService.Create(req)
	if err != nil {
		hutils.HandleError(w, err)
		return
	}

	hutils.WriteJSON(w, http.StatusCreated, story)
}
