package client

import (
	"database/sql"

	"github.com/Satishcg12/CentralAuthV3/server/internal/db/sqlc"
	"github.com/Satishcg12/CentralAuthV3/server/internal/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

// UpdateClient handles updating a client
func (h *ClientHandler) UpdateClient(c echo.Context) error {
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

	// Parse the request body
	req := new(UpdateClientRequest)
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

	// Update the client
	client, err := h.store.UpdateClient(c.Request().Context(), sqlc.UpdateClientParams{
		ID:             clientID,
		Name:           req.Name,
		Description:    sql.NullString{String: req.Description, Valid: req.Description != ""},
		RedirectUris:   pq.StringArray(req.RedirectURIs),
		WebsiteUrl:     sql.NullString{String: req.WebsiteURL, Valid: req.WebsiteURL != ""},
		IsConfidential: sql.NullBool{Bool: req.IsConfidential, Valid: true},
	})
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
			"Could not update client",
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
		"Client updated successfully",
		res,
	)
}
