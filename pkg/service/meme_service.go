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

type MemeService struct {
	db *database.DB
}

func NewMemeService(db *database.DB) *MemeService {
	return &MemeService{db: db}
}

func (s *MemeService) CreateCategory(req dto.CreateMemeCategoryRequest) (*dto.MemeCategoryResponse, error) {
	id := uuid.New().String()
	now := time.Now()

	_, err := s.db.GetConn().Exec(`
		INSERT INTO meme_categories (id, name, created_at) VALUES ($1, $2, $3)
	`, id, req.Name, now)

	if err != nil {
		return nil, fmt.Errorf("failed to create meme category: %w", err)
	}

	return &dto.MemeCategoryResponse{
		ID:        id,
		Name:      req.Name,
		CreatedAt: models.FormatTime(now),
	}, nil
}

func (s *MemeService) GetAllCategories() ([]dto.MemeCategoryResponse, error) {
	rows, err := s.db.GetConn().Query(`
		SELECT id, name, created_at FROM meme_categories ORDER BY name
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to get categories: %w", err)
	}
	defer rows.Close()

	var categories []dto.MemeCategoryResponse
	for rows.Next() {
		var cat models.MemeCategory
		if err := rows.Scan(&cat.ID, &cat.Name, &cat.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan category: %w", err)
		}
		categories = append(categories, dto.MemeCategoryResponse{
			ID:        cat.ID,
			Name:      cat.Name,
			CreatedAt: models.FormatTime(cat.CreatedAt),
		})
	}

	return categories, nil
}

func (s *MemeService) GetCategoryWithMemes(categoryID string) (*dto.MemeCategoryWithMemesResponse, error) {
	var cat models.MemeCategory
	err := s.db.GetConn().QueryRow(`
		SELECT id, name, created_at FROM meme_categories WHERE id = $1
	`, categoryID).Scan(&cat.ID, &cat.Name, &cat.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("category not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get category: %w", err)
	}

	rows, err := s.db.GetConn().Query(`
		SELECT id, category_id, type, src, created_at
		FROM memes WHERE category_id = $1 ORDER BY created_at DESC
	`, categoryID)
	if err != nil {
		return nil, fmt.Errorf("failed to get memes: %w", err)
	}
	defer rows.Close()

	var memes []dto.MemeResponse
	for rows.Next() {
		var meme models.Meme
		if err := rows.Scan(&meme.ID, &meme.CategoryID, &meme.Type, &meme.Src, &meme.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan meme: %w", err)
		}
		memes = append(memes, dto.MemeResponse{
			ID:         meme.ID,
			CategoryID: meme.CategoryID,
			Type:       meme.Type,
			Src:        meme.Src,
			CreatedAt:  models.FormatTime(meme.CreatedAt),
		})
	}

	return &dto.MemeCategoryWithMemesResponse{
		ID:        cat.ID,
		Name:      cat.Name,
		Memes:     memes,
		CreatedAt: models.FormatTime(cat.CreatedAt),
	}, nil
}

func (s *MemeService) GetAllCategoriesWithMemes() ([]dto.MemeCategoryWithMemesResponse, error) {
	categories, err := s.GetAllCategories()
	if err != nil {
		return nil, err
	}

	result := make([]dto.MemeCategoryWithMemesResponse, len(categories))
	for i, cat := range categories {
		withMemes, err := s.GetCategoryWithMemes(cat.ID)
		if err != nil {
			return nil, err
		}
		result[i] = *withMemes
	}

	return result, nil
}

func (s *MemeService) CreateMeme(req dto.CreateMemeRequest) (*dto.MemeResponse, error) {
	id := uuid.New().String()
	now := time.Now()

	_, err := s.db.GetConn().Exec(`
		INSERT INTO memes (id, category_id, type, src, created_at) VALUES ($1, $2, $3, $4, $5)
	`, id, req.CategoryID, req.Type, req.Src, now)

	if err != nil {
		return nil, fmt.Errorf("failed to create meme: %w", err)
	}

	return &dto.MemeResponse{
		ID:         id,
		CategoryID: req.CategoryID,
		Type:       req.Type,
		Src:        req.Src,
		CreatedAt:  models.FormatTime(now),
	}, nil
}

func (s *MemeService) GetMemeByID(id string) (*dto.MemeResponse, error) {
	var meme models.Meme
	err := s.db.GetConn().QueryRow(`
		SELECT id, category_id, type, src, created_at FROM memes WHERE id = $1
	`, id).Scan(&meme.ID, &meme.CategoryID, &meme.Type, &meme.Src, &meme.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("meme not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get meme: %w", err)
	}

	return &dto.MemeResponse{
		ID:         meme.ID,
		CategoryID: meme.CategoryID,
		Type:       meme.Type,
		Src:        meme.Src,
		CreatedAt:  models.FormatTime(meme.CreatedAt),
	}, nil
}

func (s *MemeService) UpdateMeme(id string, req dto.UpdateMemeRequest) (*dto.MemeResponse, error) {
	existing, err := s.GetMemeByID(id)
	if err != nil {
		return nil, err
	}

	updates := []string{}
	args := []interface{}{}
	paramNum := 1

	if req.CategoryID != nil {
		updates = append(updates, fmt.Sprintf("category_id = $%d", paramNum))
		args = append(args, *req.CategoryID)
		paramNum++
	}
	if req.Type != nil {
		updates = append(updates, fmt.Sprintf("type = $%d", paramNum))
		args = append(args, *req.Type)
		paramNum++
	}
	if req.Src != nil {
		updates = append(updates, fmt.Sprintf("src = $%d", paramNum))
		args = append(args, *req.Src)
		paramNum++
	}

	if len(updates) > 0 {
		args = append(args, id)
		query := fmt.Sprintf("UPDATE memes SET %s WHERE id = $%d", joinUpdates(updates), paramNum)
		_, err = s.db.GetConn().Exec(query, args...)
		if err != nil {
			return nil, fmt.Errorf("failed to update meme: %w", err)
		}
	} else {
		return existing, nil
	}

	return s.GetMemeByID(id)
}

func (s *MemeService) DeleteMeme(id string) error {
	_, err := s.db.GetConn().Exec("DELETE FROM memes WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete meme: %w", err)
	}
	return nil
}

func (s *MemeService) DeleteCategory(id string) error {
	_, err := s.db.GetConn().Exec("DELETE FROM meme_categories WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete category: %w", err)
	}
	return nil
}

