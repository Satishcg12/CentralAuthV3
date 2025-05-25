package client

import (
	"database/sql"

	"github.com/Satishcg12/CentralAuthV3/server/internal/db/sqlc"
	"github.com/Satishcg12/CentralAuthV3/server/internal/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *ClientHandler) RegenerateClientSecret(c echo.Context) error {
	// Parse client ID from URL parameter
	clientIDStr := c.Param("id")
	clientID, err := uuid.Parse(clientIDStr)
	if err != nil {
		return utils.RespondWithError(
			c,
			utils.StatusCodeBadRequest,
			"Invalid client ID",
			utils.ErrorCodeInvalidRequest,
			"Client ID must be a valid UUID",
			err,
		)
	}

	// Check if client exists
	_, err = h.store.GetClientById(c.Request().Context(), clientID)
	if err != nil {
		if err == sql.ErrNoRows {
			return utils.RespondWithError(
				c,
				utils.StatusCodeNotFound,
				"Client not found",
				utils.ErrorCodeResourceNotFound,
				"The specified client does not exist",
				nil,
			)
		}
		return utils.RespondWithInternalError(c, "Failed to check client existence", err)
	}

	// Generate new client secret (reusing function from create.handler.go)
	newSecret, err := generateClientSecret()
	if err != nil {
		return utils.RespondWithInternalError(c, "Failed to generate client secret", err)
	}

	// Update client secret in database
	updatedClient, err := h.store.RegenerateClientSecret(c.Request().Context(), sqlc.RegenerateClientSecretParams{
		ID:           clientID,
		ClientSecret: newSecret,
	})
	if err != nil {
		return utils.RespondWithInternalError(c, "Failed to update client secret", err)
	}

	// Convert to response DTO
	response := RegenerateSecretResponse{
		ID:             updatedClient.ID,
		ClientID:       updatedClient.ClientID,
		ClientSecret:   updatedClient.ClientSecret,
		Name:           updatedClient.Name,
		Description:    updatedClient.Description.String,
		RedirectURIs:   StringArrayToSlice(updatedClient.RedirectUris),
		WebsiteURL:     updatedClient.WebsiteUrl.String,
		IsActive:       updatedClient.IsActive.Bool,
		IsConfidential: updatedClient.IsConfidential.Bool,
		CreatedAt:      updatedClient.CreatedAt.Time,
		UpdatedAt:      updatedClient.UpdatedAt.Time,
	}

	return utils.RespondWithSuccess(
		c,
		utils.StatusCodeSuccess,
		"Client secret regenerated successfully",
		map[string]interface{}{
			"client": response,
		},
	)
}

func (h *ClientHandler) RegenerateClientSecretByClientID(c echo.Context) error {
	// Get client_id from URL parameter
	clientIDParam := c.Param("client_id")
	if clientIDParam == "" {
		return utils.RespondWithError(
			c,
			utils.StatusCodeBadRequest,
			"Missing client_id parameter",
			utils.ErrorCodeInvalidRequest,
			"Client ID parameter is required",
			nil,
		)
	}

	// Check if client exists
	_, err := h.store.GetClientByClientId(c.Request().Context(), clientIDParam)
	if err != nil {
		if err == sql.ErrNoRows {
			return utils.RespondWithError(
				c,
				utils.StatusCodeNotFound,
				"Client not found",
				utils.ErrorCodeResourceNotFound,
				"The specified client does not exist",
				nil,
			)
		}
		return utils.RespondWithInternalError(c, "Failed to check client existence", err)
	}

	// Generate new client secret
	newSecret, err := generateClientSecret()
	if err != nil {
		return utils.RespondWithInternalError(c, "Failed to generate client secret", err)
	}

	// Update client secret in database using client_id
	updatedClient, err := h.store.RegenerateClientSecretByClientId(c.Request().Context(), sqlc.RegenerateClientSecretByClientIdParams{
		ClientID:     clientIDParam,
		ClientSecret: newSecret,
	})
	if err != nil {
		return utils.RespondWithInternalError(c, "Failed to update client secret", err)
	}

	// Convert to response DTO
	response := RegenerateSecretResponse{
		ID:             updatedClient.ID,
		ClientID:       updatedClient.ClientID,
		ClientSecret:   updatedClient.ClientSecret,
		Name:           updatedClient.Name,
		Description:    updatedClient.Description.String,
		RedirectURIs:   StringArrayToSlice(updatedClient.RedirectUris),
		WebsiteURL:     updatedClient.WebsiteUrl.String,
		IsActive:       updatedClient.IsActive.Bool,
		IsConfidential: updatedClient.IsConfidential.Bool,
		CreatedAt:      updatedClient.CreatedAt.Time,
		UpdatedAt:      updatedClient.UpdatedAt.Time,
	}

	return utils.RespondWithSuccess(
		c,
		utils.StatusCodeSuccess,
		"Client secret regenerated successfully",
		map[string]interface{}{
			"client": response,
		},
	)
}
