package dto

type CreateJournalRequest struct {
	Title       string   `json:"title" validate:"required,min=1,max=500"`
	Body        string   `json:"body" validate:"required,min=1"`
	Summary     string   `json:"summary" validate:"required,min=1,max=1000"`
	PublishedOn string   `json:"published_on" validate:"required"`
	Category    *string  `json:"category,omitempty" validate:"omitempty,max=100"`
	Tags        []string `json:"tags,omitempty"`
	Author      *string  `json:"author,omitempty" validate:"omitempty,max=200"`
	ReadTime    *string  `json:"read_time,omitempty" validate:"omitempty,max=50"`
}

type UpdateJournalRequest struct {
	Title       *string  `json:"title,omitempty" validate:"omitempty,min=1,max=500"`
	Body        *string  `json:"body,omitempty" validate:"omitempty,min=1"`
	Summary     *string  `json:"summary,omitempty" validate:"omitempty,min=1,max=1000"`
	PublishedOn *string  `json:"published_on,omitempty"`
	Category    *string  `json:"category,omitempty" validate:"omitempty,max=100"`
	Tags        []string `json:"tags,omitempty"`
	Author      *string  `json:"author,omitempty" validate:"omitempty,max=200"`
	ReadTime    *string  `json:"read_time,omitempty" validate:"omitempty,max=50"`
}

type JournalResponse struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Body        string   `json:"body"`
	Summary     string   `json:"summary"`
	PublishedOn string   `json:"published_on"`
	Category    *string  `json:"category,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Author      *string  `json:"author,omitempty"`
	ReadTime    *string  `json:"read_time,omitempty"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}


