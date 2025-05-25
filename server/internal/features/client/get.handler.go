package client

import (
	"database/sql"

	"github.com/Satishcg12/CentralAuthV3/server/internal/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// GetClientById handles getting a client by ID
func (h *ClientHandler) GetClientById(c echo.Context) error {
	// Get client ID from URL parameter
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

	// Get the client
	client, err := h.store.GetClientById(c.Request().Context(), clientID)
	if err != nil {
		if err == sql.ErrNoRows {
			return utils.RespondWithError(
				c,
				utils.StatusCodeNotFound,
				"Client not found",
				utils.ErrorCodeResourceNotFound,
				"Client with the specified ID does not exist",
				nil,
			)
		}
		return utils.RespondWithError(
			c,
			utils.StatusCodeInternalError,
			"Internal server error",
			utils.ErrorCodeDatabaseError,
			"Could not fetch client",
			err,
		)
	}

	// Create the response
	res := ClientResponse{
		ID:             client.ID,
		Name:           client.Name,
		Description:    client.Description.String,
		ClientID:       client.ClientID,
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
		utils.StatusCodeSuccess,
		"Client retrieved successfully",
		res,
	)
}
