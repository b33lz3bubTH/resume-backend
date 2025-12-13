package dto

type CreateBootcampRequest struct {
	Title           string            `json:"title" validate:"required,min=1,max=200"`
	Subtitle        string            `json:"subtitle" validate:"required,min=1,max=300"`
	Description     string            `json:"description" validate:"required,min=1"`
	LongDescription string            `json:"long_description" validate:"required,min=1"`
	TechStack       []string          `json:"tech_stack" validate:"required,min=1"`
	Duration        string            `json:"duration" validate:"required,min=1,max=50"`
	Level           string            `json:"level" validate:"required,min=1,max=100"`
	Price           string            `json:"price" validate:"required,min=1,max=50"`
	Highlights      []string          `json:"highlights" validate:"required,min=1"`
	Modules         []BootcampModuleDTO `json:"modules" validate:"required,min=1"`
	ProjectFeatures []string          `json:"project_features" validate:"required,min=1"`
	TargetAudience  []string          `json:"target_audience" validate:"required,min=1"`
	Images          []string          `json:"images,omitempty"`
	Videos          []string          `json:"videos,omitempty"`
	GithubRepo      *string           `json:"github_repo,omitempty"`
	DemoURL         *string           `json:"demo_url,omitempty"`
	Status          string            `json:"status" validate:"required,oneof=active upcoming completed"`
	EnrolledCount   int               `json:"enrolled_count" validate:"gte=0"`
	Rating          *float64          `json:"rating,omitempty" validate:"omitempty,gte=0,lte=5"`
}

type BootcampModuleDTO struct {
	Title       string   `json:"title" validate:"required,min=1,max=200"`
	Description string   `json:"description" validate:"required,min=1"`
	Duration    string   `json:"duration" validate:"required,min=1,max=50"`
	Topics      []string `json:"topics" validate:"required,min=1"`
}

type UpdateBootcampRequest struct {
	Title           *string           `json:"title,omitempty" validate:"omitempty,min=1,max=200"`
	Subtitle        *string           `json:"subtitle,omitempty" validate:"omitempty,min=1,max=300"`
	Description     *string           `json:"description,omitempty" validate:"omitempty,min=1"`
	LongDescription *string           `json:"long_description,omitempty" validate:"omitempty,min=1"`
	TechStack       []string          `json:"tech_stack,omitempty" validate:"omitempty,min=1"`
	Duration        *string           `json:"duration,omitempty" validate:"omitempty,min=1,max=50"`
	Level           *string           `json:"level,omitempty" validate:"omitempty,min=1,max=100"`
	Price           *string           `json:"price,omitempty" validate:"omitempty,min=1,max=50"`
	Highlights      []string          `json:"highlights,omitempty" validate:"omitempty,min=1"`
	Modules         []BootcampModuleDTO `json:"modules,omitempty" validate:"omitempty,min=1"`
	ProjectFeatures []string          `json:"project_features,omitempty" validate:"omitempty,min=1"`
	TargetAudience  []string          `json:"target_audience,omitempty" validate:"omitempty,min=1"`
	Images          []string          `json:"images,omitempty"`
	Videos          []string          `json:"videos,omitempty"`
	GithubRepo      *string           `json:"github_repo,omitempty"`
	DemoURL         *string           `json:"demo_url,omitempty"`
	Status          *string           `json:"status,omitempty" validate:"omitempty,oneof=active upcoming completed"`
	EnrolledCount   *int              `json:"enrolled_count,omitempty" validate:"omitempty,gte=0"`
	Rating          *float64          `json:"rating,omitempty" validate:"omitempty,gte=0,lte=5"`
}

type BootcampResponse struct {
	ID              string            `json:"id"`
	Title           string            `json:"title"`
	Subtitle        string            `json:"subtitle"`
	Description     string            `json:"description"`
	LongDescription string            `json:"long_description"`
	TechStack       []string          `json:"tech_stack"`
	Duration        string            `json:"duration"`
	Level           string            `json:"level"`
	Price           string            `json:"price"`
	Highlights      []string          `json:"highlights"`
	Modules         []BootcampModuleDTO `json:"modules"`
	ProjectFeatures []string          `json:"project_features"`
	TargetAudience  []string          `json:"target_audience"`
	Images          []string          `json:"images,omitempty"`
	Videos          []string          `json:"videos,omitempty"`
	GithubRepo      *string           `json:"github_repo,omitempty"`
	DemoURL         *string           `json:"demo_url,omitempty"`
	Status          string            `json:"status"`
	EnrolledCount   int               `json:"enrolled_count"`
	Rating          *float64          `json:"rating,omitempty"`
	CreatedAt       string            `json:"created_at"`
	UpdatedAt       string            `json:"updated_at"`
}


