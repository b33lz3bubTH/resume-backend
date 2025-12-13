package models

import "time"

type Bootcamp struct {
	ID              string          `json:"id"`
	Title           string          `json:"title"`
	Subtitle        string          `json:"subtitle"`
	Description     string          `json:"description"`
	LongDescription string          `json:"long_description"`
	TechStack       []string        `json:"tech_stack"`
	Duration        string          `json:"duration"`
	Level           string          `json:"level"`
	Price           string          `json:"price"`
	Highlights      []string        `json:"highlights"`
	Modules         []BootcampModule `json:"modules"`
	ProjectFeatures []string        `json:"project_features"`
	TargetAudience  []string        `json:"target_audience"`
	Images          []string        `json:"images,omitempty"`
	Videos          []string        `json:"videos,omitempty"`
	GithubRepo      *string         `json:"github_repo,omitempty"`
	DemoURL         *string         `json:"demo_url,omitempty"`
	Status          string          `json:"status"`
	EnrolledCount   int             `json:"enrolled_count"`
	Rating          *float64        `json:"rating,omitempty"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

type BootcampModule struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Duration    string   `json:"duration"`
	Topics      []string `json:"topics"`
}

type JournalEntry struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	Summary    string    `json:"summary"`
	PublishedOn string   `json:"published_on"`
	Category   *string   `json:"category,omitempty"`
	Tags       []string  `json:"tags,omitempty"`
	Author     *string   `json:"author,omitempty"`
	ReadTime   *string   `json:"read_time,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type MemeCategory struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Meme struct {
	ID         string    `json:"id"`
	CategoryID string    `json:"category_id"`
	Type       string    `json:"type"`
	Src        string    `json:"src"`
	CreatedAt  time.Time `json:"created_at"`
}

type Story struct {
	ID          string    `json:"id"`
	Media       string    `json:"media"`
	Mimetype    string    `json:"mimetype"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Contact struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Message   string    `json:"message"`
	Subject   *string   `json:"subject,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

