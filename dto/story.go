package dto

type CreateStoryRequest struct {
	Media       string `json:"media" validate:"required,min=1"`
	Mimetype    string `json:"mimetype" validate:"required,min=1,max=100"`
	Title       string `json:"title" validate:"required,min=1,max=500"`
	Description string `json:"description" validate:"required,min=1,max=1000"`
}

type UpdateStoryRequest struct {
	Media       *string `json:"media,omitempty" validate:"omitempty,min=1"`
	Mimetype    *string `json:"mimetype,omitempty" validate:"omitempty,min=1,max=100"`
	Title       *string `json:"title,omitempty" validate:"omitempty,min=1,max=500"`
	Description *string `json:"description,omitempty" validate:"omitempty,min=1,max=1000"`
}

type StoryResponse struct {
	ID          string `json:"id"`
	Media       string `json:"media"`
	Mimetype    string `json:"mimetype"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}


