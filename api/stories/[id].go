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

	_, _, _, storyService, _ := getServices(db)

	switch r.Method {
	case http.MethodGet:
		getStoryByID(w, r, id, storyService)
	case http.MethodPut:
		updateStory(w, r, id, storyService)
	case http.MethodDelete:
		deleteStory(w, r, id, storyService)
	default:
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func getStoryByID(w http.ResponseWriter, r *http.Request, id string, storyService *service.StoryService) {
	story, err := storyService.GetByID(id)
	if err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, story)
}

func updateStory(w http.ResponseWriter, r *http.Request, id string, storyService *service.StoryService) {
	if !checkAuth(r) {
		writeError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req dto.UpdateStoryRequest
	if !validateRequest(w, r, &req) {
		return
	}

	story, err := storyService.Update(id, req)
	if err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, story)
}

func deleteStory(w http.ResponseWriter, r *http.Request, id string, storyService *service.StoryService) {
	if !checkAuth(r) {
		writeError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if err := storyService.Delete(id); err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "Story deleted successfully"})
}
