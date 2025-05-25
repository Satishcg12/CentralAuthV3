package auth

// ==========
// Auth DTOs
// ==========

// === Register Dto ===
type RegisterRequest struct {
	Email       string `json:"email" validate:"required,email,max=255"`
	Password    string `json:"password" validate:"required,min=8,max=72"`
	FullName    string `json:"full_name" validate:"required,min=2,max=255"`
	DateOfBirth string `json:"date_of_birth" validate:"required,datetime=2006-01-02"`
}

type RegisterResponse struct {
	UserID string `json:"user_id"`
}

// === Login Dto ===
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=8,max=72"`
}

type LoginResponse struct {
	// set session token in cookie and return user ID
	AccessToken string `json:"access_token"`
}

// === Logout Dto ===
type LogoutRequest struct {
	// No additional fields needed as session token comes from cookie
}

type LogoutResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// === Logout All Sessions Dto ===
type LogoutAllRequest struct {
	// No additional fields needed as user ID comes from JWT token
}

type LogoutAllResponse struct {
	Success       bool   `json:"success"`
	Message       string `json:"message"`
	SessionsEnded int    `json:"sessions_ended"`
}

// === Refresh Token Dto ===
type RefreshTokenRequest struct {
	// No additional fields needed as session token comes from cookie
}

type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expire_at"`
}
