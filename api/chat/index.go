package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"resume-backend/dto"
	hutils "resume-backend/pkg/handler"
	"resume-backend/pkg/config"
	"resume-backend/pkg/service"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if hutils.HandleCORS(w, r) {
		return
	}

	if r.Method != http.MethodPost {
		hutils.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed. Only POST is supported")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		hutils.WriteError(w, http.StatusBadRequest, "Failed to read request body")
		return
	}
	defer r.Body.Close()

	var req dto.ChatRequest
	if err := json.Unmarshal(body, &req); err != nil {
		hutils.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Message == "" || req.SessionID == "" {
		hutils.WriteError(w, http.StatusBadRequest, "message and session_id are required")
		return
	}

	db, err := hutils.GetDB()
	if err != nil {
		hutils.WriteError(w, http.StatusInternalServerError, "Database connection failed")
		return
	}
	defer db.Close()

	chatService := service.NewChatService(db)
	openRouterService := service.NewOpenRouterService()

	session, err := chatService.GetOrCreateSession(req.SessionID)
	if err != nil {
		hutils.WriteError(w, http.StatusInternalServerError, "Failed to get or create session")
		return
	}

	lastMessages, err := chatService.GetLastMessages(session.ID, 5)
	if err != nil {
		hutils.WriteError(w, http.StatusInternalServerError, "Failed to get last messages")
		return
	}

	_, err = chatService.SaveMessage(session.ID, "user", req.Message, nil)
	if err != nil {
		hutils.WriteError(w, http.StatusInternalServerError, "Failed to save user message")
		return
	}

	cfg := config.Load()
	model := cfg.OpenRouterModel
	response, statusCode, err := openRouterService.CreateChatCompletionWithContext(
		lastMessages,
		req.Message,
		model,
		r.Header.Get("Referer"),
	)
	if err != nil {
		hutils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if statusCode != http.StatusOK {
		hutils.WriteJSON(w, statusCode, response)
		return
	}

	choices, ok := response["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		hutils.WriteError(w, http.StatusInternalServerError, "Invalid response from OpenRouter")
		return
	}

	choice, ok := choices[0].(map[string]interface{})
	if !ok {
		hutils.WriteError(w, http.StatusInternalServerError, "Invalid response format")
		return
	}

	message, ok := choice["message"].(map[string]interface{})
	if !ok {
		hutils.WriteError(w, http.StatusInternalServerError, "Invalid message format")
		return
	}

	content, ok := message["content"].(string)
	if !ok {
		hutils.WriteError(w, http.StatusInternalServerError, "Invalid content format")
		return
	}

	reasoningDetails, _ := message["reasoning_details"]

	assistantMessageID, err := chatService.SaveMessage(session.ID, "assistant", content, reasoningDetails)
	if err != nil {
		hutils.WriteError(w, http.StatusInternalServerError, "Failed to save assistant message")
		return
	}

	chatResponse := dto.ChatResponse{
		Answer:    content,
		SessionID: session.ID,
		MessageID: assistantMessageID,
	}

	hutils.WriteJSON(w, http.StatusOK, chatResponse)
}

