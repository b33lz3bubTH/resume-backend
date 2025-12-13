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

type ContactService struct {
	db *database.DB
}

func NewContactService(db *database.DB) *ContactService {
	return &ContactService{db: db}
}

func (s *ContactService) Create(req dto.CreateContactRequest) (*dto.ContactResponse, error) {
	id := uuid.New().String()
	now := time.Now()

	_, err := s.db.GetConn().Exec(`
		INSERT INTO contacts (id, name, email, message, subject, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, id, req.Name, req.Email, req.Message, req.Subject, now)

	if err != nil {
		return nil, fmt.Errorf("failed to create contact: %w", err)
	}

	return s.GetByID(id)
}

func (s *ContactService) GetByID(id string) (*dto.ContactResponse, error) {
	var contact models.Contact
	var subject sql.NullString

	err := s.db.GetConn().QueryRow(`
		SELECT id, name, email, message, subject, created_at
		FROM contacts WHERE id = $1
	`, id).Scan(&contact.ID, &contact.Name, &contact.Email, &contact.Message,
		&subject, &contact.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("contact not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get contact: %w", err)
	}

	contact.Subject = models.NullStringToStringPtr(subject)

	return s.toResponse(&contact), nil
}

func (s *ContactService) GetAll(page, pageSize int) ([]dto.ContactResponse, int, error) {
	offset := (page - 1) * pageSize

	var total int
	err := s.db.GetConn().QueryRow("SELECT COUNT(*) FROM contacts").Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count contacts: %w", err)
	}

	rows, err := s.db.GetConn().Query(`
		SELECT id, name, email, message, subject, created_at
		FROM contacts ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`, pageSize, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get contacts: %w", err)
	}
	defer rows.Close()

	var contacts []dto.ContactResponse
	for rows.Next() {
		var contact models.Contact
		var subject sql.NullString
		if err := rows.Scan(&contact.ID, &contact.Name, &contact.Email, &contact.Message,
			&subject, &contact.CreatedAt); err != nil {
			return nil, 0, fmt.Errorf("failed to scan contact: %w", err)
		}
		contact.Subject = models.NullStringToStringPtr(subject)
		contacts = append(contacts, *s.toResponse(&contact))
	}

	return contacts, total, nil
}

func (s *ContactService) Delete(id string) error {
	_, err := s.db.GetConn().Exec("DELETE FROM contacts WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete contact: %w", err)
	}
	return nil
}

func (s *ContactService) toResponse(c *models.Contact) *dto.ContactResponse {
	return &dto.ContactResponse{
		ID:        c.ID,
		Name:      c.Name,
		Email:     c.Email,
		Message:   c.Message,
		Subject:   c.Subject,
		CreatedAt: models.FormatTime(c.CreatedAt),
	}
}

