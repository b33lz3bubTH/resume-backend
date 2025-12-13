package dto

type CreateMemeCategoryRequest struct {
	Name string `json:"name" validate:"required,min=1,max=100"`
}

type MemeCategoryResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

type CreateMemeRequest struct {
	CategoryID string `json:"category_id" validate:"required"`
	Type       string `json:"type" validate:"required,oneof=img yt mp4 webm"`
	Src        string `json:"src" validate:"required,min=1"`
}

type UpdateMemeRequest struct {
	CategoryID *string `json:"category_id,omitempty"`
	Type       *string `json:"type,omitempty" validate:"omitempty,oneof=img yt mp4 webm"`
	Src        *string `json:"src,omitempty" validate:"omitempty,min=1"`
}

type MemeResponse struct {
	ID         string `json:"id"`
	CategoryID string `json:"category_id"`
	Type       string `json:"type"`
	Src        string `json:"src"`
	CreatedAt  string `json:"created_at"`
}

type MemeCategoryWithMemesResponse struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Memes     []MemeResponse `json:"memes"`
	CreatedAt string         `json:"created_at"`
}


