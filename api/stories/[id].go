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

	_, _, _, storyService, _ := hutils.GetServices(db)

	switch r.Method {
	case http.MethodGet:
		getStoryByID(w, r, id, storyService)
	case http.MethodPut:
		updateStory(w, r, id, storyService)
	case http.MethodDelete:
		deleteStory(w, r, id, storyService)
	default:
		hutils.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func getStoryByID(w http.ResponseWriter, r *http.Request, id string, storyService *service.StoryService) {
	story, err := storyService.GetByID(id)
	if err != nil {
		hutils.HandleError(w, err)
		return
	}
	hutils.WriteJSON(w, http.StatusOK, story)
}

func updateStory(w http.ResponseWriter, r *http.Request, id string, storyService *service.StoryService) {
	if !hutils.CheckAuth(r) {
		hutils.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req dto.UpdateStoryRequest
	if !hutils.ValidateRequest(w, r, &req) {
		return
	}

	story, err := storyService.Update(id, req)
	if err != nil {
		hutils.HandleError(w, err)
		return
	}

	hutils.WriteJSON(w, http.StatusOK, story)
}

func deleteStory(w http.ResponseWriter, r *http.Request, id string, storyService *service.StoryService) {
	if !hutils.CheckAuth(r) {
		hutils.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if err := storyService.Delete(id); err != nil {
		hutils.HandleError(w, err)
		return
	}

	hutils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Story deleted successfully"})
}
