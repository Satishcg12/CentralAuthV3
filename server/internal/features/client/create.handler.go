package client

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"

	"github.com/Satishcg12/CentralAuthV3/server/internal/db/sqlc"
	"github.com/Satishcg12/CentralAuthV3/server/internal/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

// CreateClient handles client creation
func (h *ClientHandler) CreateClient(c echo.Context) error {
	// Parse the request body
	req := new(CreateClientRequest)
	if err := c.Bind(req); err != nil {
		return utils.RespondWithError(
			c,
			utils.StatusCodeBadRequest,
			"Invalid request data",
			utils.ErrorCodeInvalidRequest,
			"Could not parse request body",
			err,
		)
	}

	// Validate the request body
	if err := c.Validate(req); err != nil {
		return err
	}

	// Generate client ID and secret
	clientID, err := generateClientID()
	if err != nil {
		return utils.RespondWithInternalError(c, "Failed to generate client ID", err)
	}

	clientSecret, err := generateClientSecret()
	if err != nil {
		return utils.RespondWithInternalError(c, "Failed to generate client secret", err)
	}

	// TODO: Get user ID from JWT token when authentication is implemented

	// Create the client
	client, err := h.store.CreateClient(c.Request().Context(), sqlc.CreateClientParams{
		Name:           req.Name,
		Description:    sql.NullString{String: req.Description, Valid: req.Description != ""},
		ClientID:       clientID,
		ClientSecret:   clientSecret,
		RedirectUris:   pq.StringArray(req.RedirectURIs),
		WebsiteUrl:     sql.NullString{String: req.WebsiteURL, Valid: req.WebsiteURL != ""},
		IsActive:       sql.NullBool{Bool: true, Valid: true},
		IsConfidential: sql.NullBool{Bool: req.IsConfidential, Valid: true},
		CreatedBy:      uuid.NullUUID{}, // Empty for now
	})
	if err != nil {
		return utils.RespondWithError(
			c,
			utils.StatusCodeInternalError,
			"Internal server error",
			utils.ErrorCodeDatabaseError,
			"Could not create client",
			err,
		)
	}

	// Create the response
	res := CreateClientResponse{
		ID:             client.ID,
		Name:           client.Name,
		Description:    client.Description.String,
		ClientID:       client.ClientID,
		ClientSecret:   client.ClientSecret,
		RedirectURIs:   StringArrayToSlice(client.RedirectUris),
		WebsiteURL:     client.WebsiteUrl.String,
		IsActive:       client.IsActive.Bool,
		IsConfidential: client.IsConfidential.Bool,
		CreatedAt:      client.CreatedAt.Time,
		UpdatedAt:      client.UpdatedAt.Time,
	}

	// Send the response
	return utils.RespondWithSuccess(
		c,
		utils.StatusCodeCreated,
		"Client created successfully",
		res,
	)
}

// generateClientID generates a unique client ID
func generateClientID() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return fmt.Sprintf("client_%s", hex.EncodeToString(bytes)[:16]), nil
}

// generateClientSecret generates a secure client secret
func generateClientSecret() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
