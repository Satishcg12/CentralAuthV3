package client

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

// ==========
// Client DTOs
// ==========

// === Create Client Dto ===
type CreateClientRequest struct {
	Name           string   `json:"name" validate:"required,min=3,max=100"`
	Description    string   `json:"description" validate:"max=500"`
	RedirectURIs   []string `json:"redirect_uris" validate:"required,min=1,dive,url"`
	WebsiteURL     string   `json:"website_url" validate:"omitempty,url,max=255"`
	IsConfidential bool     `json:"is_confidential"`
}

type CreateClientResponse struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	ClientID       string    `json:"client_id"`
	ClientSecret   string    `json:"client_secret"`
	RedirectURIs   []string  `json:"redirect_uris"`
	WebsiteURL     string    `json:"website_url"`
	IsActive       bool      `json:"is_active"`
	IsConfidential bool      `json:"is_confidential"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// === Get Client Dto ===
type ClientResponse struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	ClientID       string    `json:"client_id"`
	RedirectURIs   []string  `json:"redirect_uris"`
	WebsiteURL     string    `json:"website_url"`
	IsActive       bool      `json:"is_active"`
	IsConfidential bool      `json:"is_confidential"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type ClientDetailResponse struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	ClientID       string    `json:"client_id"`
	ClientSecret   string    `json:"client_secret"`
	RedirectURIs   []string  `json:"redirect_uris"`
	WebsiteURL     string    `json:"website_url"`
	IsActive       bool      `json:"is_active"`
	IsConfidential bool      `json:"is_confidential"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// === Update Client Dto ===
type UpdateClientRequest struct {
	Name           string   `json:"name" validate:"required,min=3,max=100"`
	Description    string   `json:"description" validate:"max=500"`
	RedirectURIs   []string `json:"redirect_uris" validate:"required,min=1,dive,url"`
	WebsiteURL     string   `json:"website_url" validate:"omitempty,url,max=255"`
	IsConfidential bool     `json:"is_confidential"`
}

// === List Clients Dto ===
type ListClientsResponse struct {
	Clients []ClientResponse `json:"clients"`
	Total   int64            `json:"total"`
}

// === Regenerate Secret Dto ===
type RegenerateSecretResponse struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	ClientID       string    `json:"client_id"`
	ClientSecret   string    `json:"client_secret"`
	RedirectURIs   []string  `json:"redirect_uris"`
	WebsiteURL     string    `json:"website_url"`
	IsActive       bool      `json:"is_active"`
	IsConfidential bool      `json:"is_confidential"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// Helper function to convert pq.StringArray to []string
func StringArrayToSlice(arr pq.StringArray) []string {
	if arr == nil {
		return []string{}
	}
	return []string(arr)
}

// Helper function to convert []string to pq.StringArray
func SliceToStringArray(slice []string) pq.StringArray {
	if slice == nil {
		return pq.StringArray{}
	}
	return pq.StringArray(slice)
}
