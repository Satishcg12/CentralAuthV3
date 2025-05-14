-- name: RegisterUser :one
INSERT INTO users (
    email,
    password_hash,
    full_name,
    date_of_birth,
    email_verified,
    role
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: GetUserByIdentifier :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users SET
    full_name = COALESCE(sqlc.narg(full_name), full_name),
    date_of_birth = COALESCE(sqlc.narg(date_of_birth), date_of_birth),
    email_verified = COALESCE(sqlc.narg(email_verified), email_verified),
    updated_at = CURRENT_TIMESTAMP
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY created_at DESC
LIMIT $1
OFFSET $2;
