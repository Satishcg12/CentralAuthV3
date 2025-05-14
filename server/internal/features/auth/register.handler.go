package auth

import (
	"database/sql"
	"time"

	"github.com/Satishcg12/CentralAuthV3/server/internal/db/sqlc"
	"github.com/Satishcg12/CentralAuthV3/server/internal/utils"
	"github.com/labstack/echo/v4"
)

func (h *AuthHandler) Register(c echo.Context) error {
	// Parse the request body
	req := new(RegisterRequest)
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

	// Check if the email already exists
	_, err := h.store.GetUserByEmail(c.Request().Context(), req.Email)
	if err == nil {
		return utils.RespondWithError(
			c,
			utils.StatusCodeConflict,
			"User already exists",
			utils.ErrorCodeDuplicateEntry,
			"User with this email already exists",
			map[string]any{
				"email": "Email already exists",
			},
		)
	} else if err != sql.ErrNoRows {
		return utils.RespondWithError(
			c,
			utils.StatusCodeInternalError,
			"Internal server error",
			utils.ErrorCodeDatabaseError,
			"Could not check if user exists",
			err,
		)
	}

	// Hash the password
	hashedPassword, err := utils.Hash(req.Password)
	if err != nil {
		return utils.RespondWithError(
			c,
			utils.StatusCodeInternalError,
			"Internal server error",
			utils.ErrorCodeInternalError,
			"Could not hash password",
			err,
		)
	}

	// Parse date of birth if provided
	var dob sql.NullTime
	if req.DateOfBirth != "" {
		parsedDOB, err := time.Parse("2006-01-02", req.DateOfBirth)
		if err != nil {
			return utils.RespondWithError(
				c,
				utils.StatusCodeBadRequest,
				"Invalid date format",
				utils.ErrorCodeInvalidRequest,
				"Date of birth must be in YYYY-MM-DD format",
				err,
			)
		}
		dob = sql.NullTime{
			Time:  parsedDOB,
			Valid: true,
		}
	}

	// Create the user
	user, err := h.store.CreateUser(c.Request().Context(), sqlc.CreateUserParams{
		Email:         req.Email,
		PasswordHash:  hashedPassword,
		FullName:      req.FullName,
		DateOfBirth:   dob,
		Active:        sql.NullBool{Bool: true, Valid: true},
		EmailVerified: sql.NullBool{Bool: false, Valid: true},
		Role:          sql.NullString{String: "user", Valid: true},
	})
	if err != nil {
		return utils.RespondWithError(
			c,
			utils.StatusCodeInternalError,
			"Internal server error",
			utils.ErrorCodeDatabaseError,
			"Could not create user",
			err,
		)
	}

	// Create the response
	res := RegisterResponse{
		UserID: user.ID.String(),
	}

	// Send the response
	return utils.RespondWithSuccess(
		c,
		utils.StatusCodeCreated,
		"User created successfully",
		res,
	)
}
