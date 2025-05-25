package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/Satishcg12/CentralAuthV3/server/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

var (
	// jwtConfig holds the JWT configuration
	jwtConfig config.JWTConfig
)

// InitJWT initializes the JWT configuration
func InitJWT(cfg config.JWTConfig) {
	jwtConfig = cfg
}

// AccessTokenClaims represents the claims for access tokens
type AccessTokenClaims struct {
	UserID        string `json:"user_id"`
	Email         string `json:"email,omitempty"`
	FullName      string `json:"full_name,omitempty"`
	EmailVerified bool   `json:"email_verified,omitempty"`
	jwt.RegisteredClaims
}

// CreateAccessToken generates a JWT access token for the authenticated user with roles and permissions
// It returns the token string, expiry time in seconds, and any error
func CreateAccessToken(claims AccessTokenClaims) (string, int, error) {
	// Set the token expiry time
	expiry := jwtConfig.ExpiryHours * 3600 // Convert hours to seconds

	// Set the expiration time in the claims
	expirationTime := time.Now().Add(time.Duration(expiry) * time.Second)
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	// Create the token using the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	log.Println(jwtConfig.Secret)
	// Sign the token with the secret key
	if jwtConfig.Secret == "" {
		return "", 0, fmt.Errorf("JWT secret is not configured")
	}

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(jwtConfig.Secret))
	if err != nil {
		return "", 0, fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, expiry, nil
}

// GetTokenFromRequest extracts the JWT token from the Authorization header
func GetUserIDFromAccessToken(tokenString string) (string, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &AccessTokenClaims{}, func(token *jwt.Token) (any, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the key for validation
		return []byte(jwtConfig.Secret), nil
	})

	if err != nil {
		return "", fmt.Errorf("invalid token: %w", err)
	}

	// Extract and validate claims
	if claims, ok := token.Claims.(*AccessTokenClaims); ok && token.Valid {
		return claims.UserID, nil
	}

	return "", fmt.Errorf("invalid token claims")
}

// ValidateToken validates and parses a JWT token
func ValidateToken(tokenString string) (*AccessTokenClaims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the key for validation
		return []byte(jwtConfig.Secret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	// Extract and validate claims
	if claims, ok := token.Claims.(*AccessTokenClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token claims")
}
