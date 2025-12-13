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

func SetCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func WriteError(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, map[string]string{"error": message})
}

func CheckAuth(r *http.Request) bool {
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

func GetDB() (*database.DB, error) {
	cfg := config.Load()
	return database.NewDB(cfg.DatabaseURL)
}

func GetServices(db *database.DB) (*service.BootcampService, *service.JournalService, *service.MemeService, *service.StoryService, *service.ContactService) {
	return service.NewBootcampService(db),
		service.NewJournalService(db),
		service.NewMemeService(db),
		service.NewStoryService(db),
		service.NewContactService(db)
}

func HandleCORS(w http.ResponseWriter, r *http.Request) bool {
	SetCORS(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return true
	}
	return false
}

func ValidateRequest(w http.ResponseWriter, r *http.Request, req interface{}) bool {
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		WriteError(w, http.StatusBadRequest, "Invalid request body")
		return false
	}

	if errors := middleware.ValidateStruct(req); len(errors) > 0 {
		WriteJSON(w, http.StatusBadRequest, map[string]interface{}{
			"error":  "Validation failed",
			"errors": errors,
		})
		return false
	}
	return true
}

func GetIDFromPath(r *http.Request) string {
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

func HandleError(w http.ResponseWriter, err error) {
	if strings.Contains(err.Error(), "not found") {
		WriteError(w, http.StatusNotFound, err.Error())
	} else {
		WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

