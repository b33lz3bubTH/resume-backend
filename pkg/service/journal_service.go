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

type JournalService struct {
	db *database.DB
}

func NewJournalService(db *database.DB) *JournalService {
	return &JournalService{db: db}
}

func (s *JournalService) Create(req dto.CreateJournalRequest) (*dto.JournalResponse, error) {
	id := uuid.New().String()
	now := time.Now()

	tagsJSON := models.StringSliceToJSON(req.Tags)

	_, err := s.db.GetConn().Exec(`
		INSERT INTO journal_entries (
			id, title, body, summary, published_on, category, tags, author, read_time,
			created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`, id, req.Title, req.Body, req.Summary, req.PublishedOn, req.Category, tagsJSON,
		req.Author, req.ReadTime, now, now)

	if err != nil {
		return nil, fmt.Errorf("failed to create journal entry: %w", err)
	}

	return s.GetByID(id)
}

func (s *JournalService) GetByID(id string) (*dto.JournalResponse, error) {
	var entry models.JournalEntry
	var tagsJSON string
	var category, author, readTime sql.NullString

	err := s.db.GetConn().QueryRow(`
		SELECT id, title, body, summary, published_on, category, tags, author, read_time,
			created_at, updated_at
		FROM journal_entries WHERE id = $1
	`, id).Scan(
		&entry.ID, &entry.Title, &entry.Body, &entry.Summary, &entry.PublishedOn,
		&category, &tagsJSON, &author, &readTime, &entry.CreatedAt, &entry.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("journal entry not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get journal entry: %w", err)
	}

	entry.Category = models.NullStringToStringPtr(category)
	entry.Tags = models.JSONToStringSlice(tagsJSON)
	entry.Author = models.NullStringToStringPtr(author)
	entry.ReadTime = models.NullStringToStringPtr(readTime)

	return s.toResponse(&entry), nil
}

func (s *JournalService) GetAll() ([]dto.JournalResponse, error) {
	rows, err := s.db.GetConn().Query(`
		SELECT id, title, body, summary, published_on, category, tags, author, read_time,
			created_at, updated_at
		FROM journal_entries ORDER BY published_on DESC, created_at DESC
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to get journal entries: %w", err)
	}
	defer rows.Close()

	var entries []dto.JournalResponse
	for rows.Next() {
		var entry models.JournalEntry
		var tagsJSON string
		var category, author, readTime sql.NullString

		if err := rows.Scan(
			&entry.ID, &entry.Title, &entry.Body, &entry.Summary, &entry.PublishedOn,
			&category, &tagsJSON, &author, &readTime, &entry.CreatedAt, &entry.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan journal entry: %w", err)
		}

		entry.Category = models.NullStringToStringPtr(category)
		entry.Tags = models.JSONToStringSlice(tagsJSON)
		entry.Author = models.NullStringToStringPtr(author)
		entry.ReadTime = models.NullStringToStringPtr(readTime)

		entries = append(entries, *s.toResponse(&entry))
	}

	return entries, nil
}

func (s *JournalService) Update(id string, req dto.UpdateJournalRequest) (*dto.JournalResponse, error) {
	if _, err := s.GetByID(id); err != nil {
		return nil, err
	}

	updates := []string{}
	args := []interface{}{}
	paramNum := 1

	if req.Title != nil {
		updates = append(updates, fmt.Sprintf("title = $%d", paramNum))
		args = append(args, *req.Title)
		paramNum++
	}
	if req.Body != nil {
		updates = append(updates, fmt.Sprintf("body = $%d", paramNum))
		args = append(args, *req.Body)
		paramNum++
	}
	if req.Summary != nil {
		updates = append(updates, fmt.Sprintf("summary = $%d", paramNum))
		args = append(args, *req.Summary)
		paramNum++
	}
	if req.PublishedOn != nil {
		updates = append(updates, fmt.Sprintf("published_on = $%d", paramNum))
		args = append(args, *req.PublishedOn)
		paramNum++
	}
	if req.Category != nil {
		updates = append(updates, fmt.Sprintf("category = $%d", paramNum))
		args = append(args, *req.Category)
		paramNum++
	}
	if req.Tags != nil {
		updates = append(updates, fmt.Sprintf("tags = $%d", paramNum))
		args = append(args, models.StringSliceToJSON(req.Tags))
		paramNum++
	}
	if req.Author != nil {
		updates = append(updates, fmt.Sprintf("author = $%d", paramNum))
		args = append(args, *req.Author)
		paramNum++
	}
	if req.ReadTime != nil {
		updates = append(updates, fmt.Sprintf("read_time = $%d", paramNum))
		args = append(args, *req.ReadTime)
		paramNum++
	}

	if len(updates) > 0 {
		updates = append(updates, fmt.Sprintf("updated_at = $%d", paramNum))
		args = append(args, time.Now())
		paramNum++
		args = append(args, id)

		query := fmt.Sprintf("UPDATE journal_entries SET %s WHERE id = $%d", joinUpdates(updates), paramNum)
		_, err := s.db.GetConn().Exec(query, args...)
		if err != nil {
			return nil, fmt.Errorf("failed to update journal entry: %w", err)
		}
	}

	return s.GetByID(id)
}

func (s *JournalService) Delete(id string) error {
	_, err := s.db.GetConn().Exec("DELETE FROM journal_entries WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete journal entry: %w", err)
	}
	return nil
}

func (s *JournalService) toResponse(e *models.JournalEntry) *dto.JournalResponse {
	return &dto.JournalResponse{
		ID:          e.ID,
		Title:       e.Title,
		Body:        e.Body,
		Summary:     e.Summary,
		PublishedOn: e.PublishedOn,
		Category:    e.Category,
		Tags:        e.Tags,
		Author:      e.Author,
		ReadTime:    e.ReadTime,
		CreatedAt:   models.FormatTime(e.CreatedAt),
		UpdatedAt:   models.FormatTime(e.UpdatedAt),
	}
}

