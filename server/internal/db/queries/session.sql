-- name: CreateSession :one
INSERT INTO sessions (
    user_id,
    session_token,
    user_agent,
    ip_address,
    expires_at
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetSessionByToken :one
SELECT * FROM sessions 
WHERE session_token = $1 
AND is_active = TRUE 
AND expires_at > CURRENT_TIMESTAMP;

-- name: DeactivateSession :exec
UPDATE sessions 
SET is_active = FALSE 
WHERE session_token = $1;

-- name: DeactivateAllUserSessions :exec
UPDATE sessions 
SET is_active = FALSE 
WHERE user_id = $1;

-- name: CountActiveUserSessions :one
SELECT COUNT(*) FROM sessions 
WHERE user_id = $1 
AND is_active = TRUE 
AND expires_at > CURRENT_TIMESTAMP;
