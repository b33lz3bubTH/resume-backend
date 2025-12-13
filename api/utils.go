package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"resume-backend/pkg/config"
	"resume-backend/pkg/database"
	"resume-backend/pkg/middleware"
	"resume-backend/pkg/service"
)

func setCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

func checkAuth(r *http.Request) bool {
	cfg := config.Load()
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return false
	}
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return false
	}
	return parts[1] == cfg.RootKey
}

func getDB() (*database.DB, error) {
	cfg := config.Load()
	return database.NewDB(cfg.DatabaseURL)
}

func getServices(db *database.DB) (*service.BootcampService, *service.JournalService, *service.MemeService, *service.StoryService, *service.ContactService) {
	return service.NewBootcampService(db),
		service.NewJournalService(db),
		service.NewMemeService(db),
		service.NewStoryService(db),
		service.NewContactService(db)
}

func handleCORS(w http.ResponseWriter, r *http.Request) bool {
	setCORS(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return true
	}
	return false
}

func validateRequest(w http.ResponseWriter, r *http.Request, req interface{}) bool {
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return false
	}

	if errors := middleware.ValidateStruct(req); len(errors) > 0 {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"error":  "Validation failed",
			"errors": errors,
		})
		return false
	}
	return true
}

func getIDFromPath(r *http.Request) string {
	path := strings.TrimPrefix(r.URL.Path, "/api/")
	parts := strings.Split(path, "/")
	if len(parts) > 0 {
		lastPart := parts[len(parts)-1]
		if lastPart != "" {
			return lastPart
		}
		if len(parts) > 1 {
			return parts[len(parts)-2]
		}
	}
	return ""
}

func handleError(w http.ResponseWriter, err error) {
	if strings.Contains(err.Error(), "not found") {
		writeError(w, http.StatusNotFound, err.Error())
	} else {
		writeError(w, http.StatusInternalServerError, err.Error())
	}
}

