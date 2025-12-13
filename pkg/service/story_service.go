package service

import (
	"database/sql"
	"fmt"
	"time"

	"resume-backend/pkg/database"
	"resume-backend/dto"
	"resume-backend/pkg/models"

	"github.com/google/uuid"
)

type StoryService struct {
	db *database.DB
}

func NewStoryService(db *database.DB) *StoryService {
	return &StoryService{db: db}
}

func (s *StoryService) Create(req dto.CreateStoryRequest) (*dto.StoryResponse, error) {
	id := uuid.New().String()
	now := time.Now()

	_, err := s.db.GetConn().Exec(`
		INSERT INTO stories (id, media, mimetype, title, description, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, id, req.Media, req.Mimetype, req.Title, req.Description, now, now)

	if err != nil {
		return nil, fmt.Errorf("failed to create story: %w", err)
	}

	return s.GetByID(id)
}

func (s *StoryService) GetByID(id string) (*dto.StoryResponse, error) {
	var story models.Story
	err := s.db.GetConn().QueryRow(`
		SELECT id, media, mimetype, title, description, created_at, updated_at
		FROM stories WHERE id = $1
	`, id).Scan(&story.ID, &story.Media, &story.Mimetype, &story.Title, &story.Description,
		&story.CreatedAt, &story.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("story not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get story: %w", err)
	}

	return s.toResponse(&story), nil
}

func (s *StoryService) GetAll() ([]dto.StoryResponse, error) {
	rows, err := s.db.GetConn().Query(`
		SELECT id, media, mimetype, title, description, created_at, updated_at
		FROM stories ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to get stories: %w", err)
	}
	defer rows.Close()

	var stories []dto.StoryResponse
	for rows.Next() {
		var story models.Story
		if err := rows.Scan(&story.ID, &story.Media, &story.Mimetype, &story.Title,
			&story.Description, &story.CreatedAt, &story.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan story: %w", err)
		}
		stories = append(stories, *s.toResponse(&story))
	}

	return stories, nil
}

func (s *StoryService) Update(id string, req dto.UpdateStoryRequest) (*dto.StoryResponse, error) {
	if _, err := s.GetByID(id); err != nil {
		return nil, err
	}

	updates := []string{}
	args := []interface{}{}
	paramNum := 1

	if req.Media != nil {
		updates = append(updates, fmt.Sprintf("media = $%d", paramNum))
		args = append(args, *req.Media)
		paramNum++
	}
	if req.Mimetype != nil {
		updates = append(updates, fmt.Sprintf("mimetype = $%d", paramNum))
		args = append(args, *req.Mimetype)
		paramNum++
	}
	if req.Title != nil {
		updates = append(updates, fmt.Sprintf("title = $%d", paramNum))
		args = append(args, *req.Title)
		paramNum++
	}
	if req.Description != nil {
		updates = append(updates, fmt.Sprintf("description = $%d", paramNum))
		args = append(args, *req.Description)
		paramNum++
	}

	if len(updates) > 0 {
		updates = append(updates, fmt.Sprintf("updated_at = $%d", paramNum))
		args = append(args, time.Now())
		paramNum++
		args = append(args, id)

		query := fmt.Sprintf("UPDATE stories SET %s WHERE id = $%d", joinUpdates(updates), paramNum)
		_, err := s.db.GetConn().Exec(query, args...)
		if err != nil {
			return nil, fmt.Errorf("failed to update story: %w", err)
		}
	}

	return s.GetByID(id)
}

func (s *StoryService) Delete(id string) error {
	_, err := s.db.GetConn().Exec("DELETE FROM stories WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete story: %w", err)
	}
	return nil
}

func (s *StoryService) toResponse(story *models.Story) *dto.StoryResponse {
	return &dto.StoryResponse{
		ID:          story.ID,
		Media:       story.Media,
		Mimetype:    story.Mimetype,
		Title:       story.Title,
		Description: story.Description,
		CreatedAt:   models.FormatTime(story.CreatedAt),
		UpdatedAt:   models.FormatTime(story.UpdatedAt),
	}
}

