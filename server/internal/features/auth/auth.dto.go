package auth

// ==========
// Auth DTOs
// ==========

// === Register Dto ===
type RegisterRequest struct {
	Email       string `json:"email" validate:"required,email,max=255"`
	Password    string `json:"password" validate:"required,min=8,max=72"`
	FullName    string `json:"full_name,omitempty" validate:"omitempty,min=2,max=255"`
	DateOfBirth string `json:"date_of_birth,omitempty" validate:"omitempty,datetime=2006-01-02"`
}

type RegisterResponse struct {
	UserID string `json:"user_id"`
}
