package client

import (
	"github.com/Satishcg12/CentralAuthV3/server/internal/utils"
	"github.com/labstack/echo/v4"
)

// GetAllClients handles getting all clients
func (h *ClientHandler) GetAllClients(c echo.Context) error {
	// Get all clients
	clients, err := h.store.GetAllClients(c.Request().Context())
	if err != nil {
		return utils.RespondWithError(
			c,
			utils.StatusCodeInternalError,
			"Internal server error",
			utils.ErrorCodeDatabaseError,
			"Could not fetch clients",
			err,
		)
	}

	// Get total count
	total, err := h.store.CountClients(c.Request().Context())
	if err != nil {
		return utils.RespondWithError(
			c,
			utils.StatusCodeInternalError,
			"Internal server error",
			utils.ErrorCodeDatabaseError,
			"Could not count clients",
			err,
		)
	}

	// Convert to response format
	clientResponses := make([]ClientResponse, len(clients))
	for i, client := range clients {
		clientResponses[i] = ClientResponse{
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
	}

	// Create the response
	res := ListClientsResponse{
		Clients: clientResponses,
		Total:   total,
	}

	// Send the response
	return utils.RespondWithSuccess(
		c,
		utils.StatusCodeSuccess,
		"Clients retrieved successfully",
		res,
	)
}
