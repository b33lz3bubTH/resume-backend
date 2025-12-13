package dto

type CreateContactRequest struct {
	Name    string  `json:"name" validate:"required,min=1,max=200"`
	Email   string  `json:"email" validate:"required,email"`
	Message string  `json:"message" validate:"required,min=1,max=5000"`
	Subject *string `json:"subject,omitempty" validate:"omitempty,max=200"`
}

type ContactResponse struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Message   string  `json:"message"`
	Subject   *string `json:"subject,omitempty"`
	CreatedAt string  `json:"created_at"`
}


