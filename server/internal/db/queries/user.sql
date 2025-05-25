-- name: GetUserByEmail :one
SELECT * 
FROM users
WHERE email = $1 
LIMIT 1;

-- name: GetUserByID :one
SELECT * 
FROM users
WHERE id = $1 
LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
    email,
    password_hash,
    full_name,
    date_of_birth,
    email_verified,
    active,
    created_at,
    updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
) RETURNING *;