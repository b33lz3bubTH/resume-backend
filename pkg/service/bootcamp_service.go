package service

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"resume-backend/pkg/database"
	"resume-backend/dto"
	"resume-backend/pkg/models"

	"github.com/google/uuid"
)

type BootcampService struct {
	db *database.DB
}

func NewBootcampService(db *database.DB) *BootcampService {
	return &BootcampService{db: db}
}

func (s *BootcampService) Create(req dto.CreateBootcampRequest) (*dto.BootcampResponse, error) {
	id := uuid.New().String()
	now := time.Now()

	techStackJSON := models.StringSliceToJSON(req.TechStack)
	highlightsJSON := models.StringSliceToJSON(req.Highlights)
	projectFeaturesJSON := models.StringSliceToJSON(req.ProjectFeatures)
	targetAudienceJSON := models.StringSliceToJSON(req.TargetAudience)
	imagesJSON := models.StringSliceToJSON(req.Images)
	videosJSON := models.StringSliceToJSON(req.Videos)

	tx, err := s.db.GetConn().Begin()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
		INSERT INTO bootcamps (
			id, title, subtitle, description, long_description, tech_stack,
			duration, level, price, highlights, project_features, target_audience,
			images, videos, github_repo, demo_url, status, enrolled_count, rating,
			created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)
	`, id, req.Title, req.Subtitle, req.Description, req.LongDescription, techStackJSON,
		req.Duration, req.Level, req.Price, highlightsJSON, projectFeaturesJSON, targetAudienceJSON,
		imagesJSON, videosJSON, req.GithubRepo, req.DemoURL, req.Status, req.EnrolledCount, req.Rating,
		now, now)

	if err != nil {
		return nil, fmt.Errorf("failed to create bootcamp: %w", err)
	}

	for _, module := range req.Modules {
		moduleID := uuid.New().String()
		topicsJSON := models.StringSliceToJSON(module.Topics)
		_, err = tx.Exec(`
			INSERT INTO bootcamp_modules (id, bootcamp_id, title, description, duration, topics, created_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
		`, moduleID, id, module.Title, module.Description, module.Duration, topicsJSON, now)
		if err != nil {
			return nil, fmt.Errorf("failed to create module: %w", err)
		}
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return s.GetByID(id)
}

func (s *BootcampService) GetByID(id string) (*dto.BootcampResponse, error) {
	var bootcamp models.Bootcamp
	var techStackJSON, highlightsJSON, projectFeaturesJSON, targetAudienceJSON, imagesJSON, videosJSON string
	var githubRepo, demoURL sql.NullString
	var rating sql.NullFloat64

	err := s.db.GetConn().QueryRow(`
		SELECT id, title, subtitle, description, long_description, tech_stack,
			duration, level, price, highlights, project_features, target_audience,
			images, videos, github_repo, demo_url, status, enrolled_count, rating,
			created_at, updated_at
		FROM bootcamps WHERE id = $1
	`, id).Scan(
		&bootcamp.ID, &bootcamp.Title, &bootcamp.Subtitle, &bootcamp.Description,
		&bootcamp.LongDescription, &techStackJSON, &bootcamp.Duration, &bootcamp.Level,
		&bootcamp.Price, &highlightsJSON, &projectFeaturesJSON, &targetAudienceJSON,
		&imagesJSON, &videosJSON, &githubRepo, &demoURL, &bootcamp.Status,
		&bootcamp.EnrolledCount, &rating, &bootcamp.CreatedAt, &bootcamp.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("bootcamp not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get bootcamp: %w", err)
	}

	bootcamp.TechStack = models.JSONToStringSlice(techStackJSON)
	bootcamp.Highlights = models.JSONToStringSlice(highlightsJSON)
	bootcamp.ProjectFeatures = models.JSONToStringSlice(projectFeaturesJSON)
	bootcamp.TargetAudience = models.JSONToStringSlice(targetAudienceJSON)
	bootcamp.Images = models.JSONToStringSlice(imagesJSON)
	bootcamp.Videos = models.JSONToStringSlice(videosJSON)
	bootcamp.GithubRepo = models.NullStringToStringPtr(githubRepo)
	bootcamp.DemoURL = models.NullStringToStringPtr(demoURL)
	bootcamp.Rating = models.NullFloat64ToFloat64Ptr(rating)

	rows, err := s.db.GetConn().Query(`
		SELECT id, title, description, duration, topics
		FROM bootcamp_modules WHERE bootcamp_id = $1
	`, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get modules: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var module models.BootcampModule
		var topicsJSON string
		if err := rows.Scan(&module.ID, &module.Title, &module.Description, &module.Duration, &topicsJSON); err != nil {
			return nil, fmt.Errorf("failed to scan module: %w", err)
		}
		module.Topics = models.JSONToStringSlice(topicsJSON)
		bootcamp.Modules = append(bootcamp.Modules, module)
	}

	return s.toResponse(&bootcamp), nil
}

func (s *BootcampService) GetAll() ([]dto.BootcampResponse, error) {
	rows, err := s.db.GetConn().Query(`
		SELECT id, title, subtitle, description, long_description, tech_stack,
			duration, level, price, highlights, project_features, target_audience,
			images, videos, github_repo, demo_url, status, enrolled_count, rating,
			created_at, updated_at
		FROM bootcamps ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to get bootcamps: %w", err)
	}
	defer rows.Close()

	var bootcamps []dto.BootcampResponse
	for rows.Next() {
		var bootcamp models.Bootcamp
		var techStackJSON, highlightsJSON, projectFeaturesJSON, targetAudienceJSON, imagesJSON, videosJSON string
		var githubRepo, demoURL sql.NullString
		var rating sql.NullFloat64

		if err := rows.Scan(
			&bootcamp.ID, &bootcamp.Title, &bootcamp.Subtitle, &bootcamp.Description,
			&bootcamp.LongDescription, &techStackJSON, &bootcamp.Duration, &bootcamp.Level,
			&bootcamp.Price, &highlightsJSON, &projectFeaturesJSON, &targetAudienceJSON,
			&imagesJSON, &videosJSON, &githubRepo, &demoURL, &bootcamp.Status,
			&bootcamp.EnrolledCount, &rating, &bootcamp.CreatedAt, &bootcamp.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan bootcamp: %w", err)
		}

		bootcamp.TechStack = models.JSONToStringSlice(techStackJSON)
		bootcamp.Highlights = models.JSONToStringSlice(highlightsJSON)
		bootcamp.ProjectFeatures = models.JSONToStringSlice(projectFeaturesJSON)
		bootcamp.TargetAudience = models.JSONToStringSlice(targetAudienceJSON)
		bootcamp.Images = models.JSONToStringSlice(imagesJSON)
		bootcamp.Videos = models.JSONToStringSlice(videosJSON)
		bootcamp.GithubRepo = models.NullStringToStringPtr(githubRepo)
		bootcamp.DemoURL = models.NullStringToStringPtr(demoURL)
		bootcamp.Rating = models.NullFloat64ToFloat64Ptr(rating)

		moduleRows, err := s.db.GetConn().Query(`
			SELECT id, title, description, duration, topics
			FROM bootcamp_modules WHERE bootcamp_id = $1
		`, bootcamp.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get modules: %w", err)
		}

		for moduleRows.Next() {
			var module models.BootcampModule
			var topicsJSON string
			if err := moduleRows.Scan(&module.ID, &module.Title, &module.Description, &module.Duration, &topicsJSON); err != nil {
				moduleRows.Close()
				return nil, fmt.Errorf("failed to scan module: %w", err)
			}
			module.Topics = models.JSONToStringSlice(topicsJSON)
			bootcamp.Modules = append(bootcamp.Modules, module)
		}
		moduleRows.Close()

		bootcamps = append(bootcamps, *s.toResponse(&bootcamp))
	}

	return bootcamps, nil
}

func (s *BootcampService) Update(id string, req dto.UpdateBootcampRequest) (*dto.BootcampResponse, error) {
	if _, err := s.GetByID(id); err != nil {
		return nil, err
	}

	tx, err := s.db.GetConn().Begin()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	updates := []string{}
	args := []interface{}{}
	paramNum := 1

	if req.Title != nil {
		updates = append(updates, fmt.Sprintf("title = $%d", paramNum))
		args = append(args, *req.Title)
		paramNum++
	}
	if req.Subtitle != nil {
		updates = append(updates, fmt.Sprintf("subtitle = $%d", paramNum))
		args = append(args, *req.Subtitle)
		paramNum++
	}
	if req.Description != nil {
		updates = append(updates, fmt.Sprintf("description = $%d", paramNum))
		args = append(args, *req.Description)
		paramNum++
	}
	if req.LongDescription != nil {
		updates = append(updates, fmt.Sprintf("long_description = $%d", paramNum))
		args = append(args, *req.LongDescription)
		paramNum++
	}
	if req.TechStack != nil {
		updates = append(updates, fmt.Sprintf("tech_stack = $%d", paramNum))
		args = append(args, models.StringSliceToJSON(req.TechStack))
		paramNum++
	}
	if req.Duration != nil {
		updates = append(updates, fmt.Sprintf("duration = $%d", paramNum))
		args = append(args, *req.Duration)
		paramNum++
	}
	if req.Level != nil {
		updates = append(updates, fmt.Sprintf("level = $%d", paramNum))
		args = append(args, *req.Level)
		paramNum++
	}
	if req.Price != nil {
		updates = append(updates, fmt.Sprintf("price = $%d", paramNum))
		args = append(args, *req.Price)
		paramNum++
	}
	if req.Highlights != nil {
		updates = append(updates, fmt.Sprintf("highlights = $%d", paramNum))
		args = append(args, models.StringSliceToJSON(req.Highlights))
		paramNum++
	}
	if req.ProjectFeatures != nil {
		updates = append(updates, fmt.Sprintf("project_features = $%d", paramNum))
		args = append(args, models.StringSliceToJSON(req.ProjectFeatures))
		paramNum++
	}
	if req.TargetAudience != nil {
		updates = append(updates, fmt.Sprintf("target_audience = $%d", paramNum))
		args = append(args, models.StringSliceToJSON(req.TargetAudience))
		paramNum++
	}
	if req.Images != nil {
		updates = append(updates, fmt.Sprintf("images = $%d", paramNum))
		args = append(args, models.StringSliceToJSON(req.Images))
		paramNum++
	}
	if req.Videos != nil {
		updates = append(updates, fmt.Sprintf("videos = $%d", paramNum))
		args = append(args, models.StringSliceToJSON(req.Videos))
		paramNum++
	}
	if req.GithubRepo != nil {
		updates = append(updates, fmt.Sprintf("github_repo = $%d", paramNum))
		args = append(args, *req.GithubRepo)
		paramNum++
	}
	if req.DemoURL != nil {
		updates = append(updates, fmt.Sprintf("demo_url = $%d", paramNum))
		args = append(args, *req.DemoURL)
		paramNum++
	}
	if req.Status != nil {
		updates = append(updates, fmt.Sprintf("status = $%d", paramNum))
		args = append(args, *req.Status)
		paramNum++
	}
	if req.EnrolledCount != nil {
		updates = append(updates, fmt.Sprintf("enrolled_count = $%d", paramNum))
		args = append(args, *req.EnrolledCount)
		paramNum++
	}
	if req.Rating != nil {
		updates = append(updates, fmt.Sprintf("rating = $%d", paramNum))
		args = append(args, *req.Rating)
		paramNum++
	}

	if len(updates) > 0 {
		updates = append(updates, fmt.Sprintf("updated_at = $%d", paramNum))
		args = append(args, time.Now())
		paramNum++
		args = append(args, id)

		_, err = tx.Exec(fmt.Sprintf("UPDATE bootcamps SET %s WHERE id = $%d", joinUpdates(updates), paramNum), args...)
		if err != nil {
			return nil, fmt.Errorf("failed to update bootcamp: %w", err)
		}
	}

	if req.Modules != nil {
		_, err = tx.Exec("DELETE FROM bootcamp_modules WHERE bootcamp_id = $1", id)
		if err != nil {
			return nil, fmt.Errorf("failed to delete modules: %w", err)
		}

		for _, module := range req.Modules {
			moduleID := uuid.New().String()
			topicsJSON := models.StringSliceToJSON(module.Topics)
			_, err = tx.Exec(`
				INSERT INTO bootcamp_modules (id, bootcamp_id, title, description, duration, topics, created_at)
				VALUES ($1, $2, $3, $4, $5, $6, $7)
			`, moduleID, id, module.Title, module.Description, module.Duration, topicsJSON, time.Now())
			if err != nil {
				return nil, fmt.Errorf("failed to create module: %w", err)
			}
		}
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return s.GetByID(id)
}

func (s *BootcampService) Delete(id string) error {
	_, err := s.db.GetConn().Exec("DELETE FROM bootcamps WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete bootcamp: %w", err)
	}
	return nil
}

func (s *BootcampService) toResponse(b *models.Bootcamp) *dto.BootcampResponse {
	modules := make([]dto.BootcampModuleDTO, len(b.Modules))
	for i, m := range b.Modules {
		modules[i] = dto.BootcampModuleDTO{
			Title:       m.Title,
			Description: m.Description,
			Duration:    m.Duration,
			Topics:      m.Topics,
		}
	}

	return &dto.BootcampResponse{
		ID:              b.ID,
		Title:           b.Title,
		Subtitle:        b.Subtitle,
		Description:     b.Description,
		LongDescription: b.LongDescription,
		TechStack:       b.TechStack,
		Duration:        b.Duration,
		Level:           b.Level,
		Price:           b.Price,
		Highlights:      b.Highlights,
		Modules:         modules,
		ProjectFeatures: b.ProjectFeatures,
		TargetAudience:  b.TargetAudience,
		Images:          b.Images,
		Videos:          b.Videos,
		GithubRepo:      b.GithubRepo,
		DemoURL:         b.DemoURL,
		Status:          b.Status,
		EnrolledCount:   b.EnrolledCount,
		Rating:          b.Rating,
		CreatedAt:       models.FormatTime(b.CreatedAt),
		UpdatedAt:       models.FormatTime(b.UpdatedAt),
	}
}

func joinUpdates(updates []string) string {
	return strings.Join(updates, ", ")
}

