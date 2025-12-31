package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"resume-backend/pkg/database"
	"resume-backend/pkg/models"

	"github.com/google/uuid"
)

type ChatService struct {
	db *database.DB
}

func NewChatService(db *database.DB) *ChatService {
	return &ChatService{db: db}
}

func (s *ChatService) GetOrCreateSession(sessionID string) (*models.ChatSession, error) {
	var session models.ChatSession
	err := s.db.GetConn().QueryRow(`
		SELECT id, created_at, updated_at
		FROM chat_sessions WHERE id = $1
	`, sessionID).Scan(&session.ID, &session.CreatedAt, &session.UpdatedAt)

	if err == sql.ErrNoRows {
		now := time.Now()
		session = models.ChatSession{
			ID:        sessionID,
			CreatedAt: now,
			UpdatedAt: now,
		}
		_, err = s.db.GetConn().Exec(`
			INSERT INTO chat_sessions (id, created_at, updated_at)
			VALUES ($1, $2, $3)
		`, session.ID, session.CreatedAt, session.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to create session: %w", err)
		}
		return &session, nil
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get session: %w", err)
	}

	return &session, nil
}

func (s *ChatService) SaveMessage(sessionID, role, content string, reasoningDetails interface{}) (string, error) {
	id := uuid.New().String()
	now := time.Now()

	var reasoningDetailsJSON *string
	if reasoningDetails != nil {
		jsonData, err := json.Marshal(reasoningDetails)
		if err == nil {
			jsonStr := string(jsonData)
			reasoningDetailsJSON = &jsonStr
		}
	}

	_, err := s.db.GetConn().Exec(`
		INSERT INTO chat_messages (id, session_id, role, content, reasoning_details, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, id, sessionID, role, content, reasoningDetailsJSON, now)

	if err != nil {
		return "", fmt.Errorf("failed to save message: %w", err)
	}

	_, err = s.db.GetConn().Exec(`
		UPDATE chat_sessions SET updated_at = $1 WHERE id = $2
	`, now, sessionID)

	if err != nil {
		return "", fmt.Errorf("failed to update session: %w", err)
	}

	return id, nil
}

func (s *ChatService) GetLastMessages(sessionID string, limit int) ([]models.ChatMessage, error) {
	rows, err := s.db.GetConn().Query(`
		SELECT id, session_id, role, content, reasoning_details, created_at
		FROM chat_messages
		WHERE session_id = $1
		ORDER BY created_at DESC
		LIMIT $2
	`, sessionID, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get messages: %w", err)
	}
	defer rows.Close()

	var messages []models.ChatMessage
	for rows.Next() {
		var msg models.ChatMessage
		var reasoningDetailsJSON sql.NullString
		if err := rows.Scan(&msg.ID, &msg.SessionID, &msg.Role, &msg.Content, &reasoningDetailsJSON, &msg.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan message: %w", err)
		}
		if reasoningDetailsJSON.Valid {
			msg.ReasoningDetails = &reasoningDetailsJSON.String
		}
		messages = append(messages, msg)
	}

	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}

	return messages, nil
}

