-- name: CreateSession :one
INSERT INTO sessions (
    user_id,
    ip_address,
    user_agent
) VALUES (
    $1, $2, $3
) RETURNING id;

-- name: GetSession :one
SELECT * FROM sessions WHERE id = $1 AND expires_at > CURRENT_TIMESTAMP LIMIT 1;

-- name: ListSessionsByUser :many
SELECT * FROM sessions WHERE user_id = $1 ORDER BY created_at DESC;

-- name: DeleteSession :exec
UPDATE sessions
SET expires_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: CreateAccessToken :one
INSERT INTO access_tokens (
    token,
    client_id,
    user_id,
    scope,
    expires_at
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING id;

-- name: GetAccessToken :one
SELECT * FROM access_tokens WHERE token = $1 AND expires_at > CURRENT_TIMESTAMP LIMIT 1;

-- name: InvalidateAccessToken :exec
UPDATE access_tokens
SET expires_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (
    token,
    client_id,
    user_id,
    scope,
    expires_at
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING id;

-- name: GetRefreshToken :one
SELECT * FROM refresh_tokens WHERE token = $1 AND expires_at > CURRENT_TIMESTAMP LIMIT 1;

-- name: InvalidateRefreshToken :exec
UPDATE refresh_tokens
SET expires_at = CURRENT_TIMESTAMP
WHERE id = $1;
