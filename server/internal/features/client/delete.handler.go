package client

import (
	"database/sql"

	"github.com/Satishcg12/CentralAuthV3/server/internal/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *ClientHandler) DeleteClient(c echo.Context) error {
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

	// Check if client exists before deletion
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

	// Delete the client
	err = h.store.DeleteClient(c.Request().Context(), clientID)
	if err != nil {
		return utils.RespondWithInternalError(c, "Failed to delete client", err)
	}

	return utils.RespondWithSuccess(
		c,
		utils.StatusCodeSuccess,
		"Client deleted successfully",
		nil,
	)
}
