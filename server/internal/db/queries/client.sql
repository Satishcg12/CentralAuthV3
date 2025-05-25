-- name: CreateClient :one
INSERT INTO clients (
    name,
    description,
    client_id,
    client_secret,
    redirect_uris,
    website_url,
    is_active,
    is_confidential,
    created_by,
    created_at,
    updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
) RETURNING *;

-- name: GetAllClients :many
SELECT 
    id,
    name,
    description,
    client_id,
    redirect_uris,
    website_url,
    is_active,
    is_confidential,
    created_at,
    updated_at
FROM clients
WHERE is_active = true
ORDER BY created_at DESC;

-- name: GetClientById :one
SELECT 
    id,
    name,
    description,
    client_id,
    redirect_uris,
    website_url,
    is_active,
    is_confidential,
    created_at,
    updated_at
FROM clients
WHERE id = $1 AND is_active = true
LIMIT 1;

-- name: GetClientByClientId :one
SELECT *
FROM clients
WHERE client_id = $1 AND is_active = true
LIMIT 1;

-- name: UpdateClient :one
UPDATE clients
SET 
    name = $2,
    description = $3,
    redirect_uris = $4,
    website_url = $5,
    is_confidential = $6,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND is_active = true
RETURNING 
    id,
    name,
    description,
    client_id,
    redirect_uris,
    website_url,
    is_active,
    is_confidential,
    created_at,
    updated_at;

-- name: DeleteClient :exec
UPDATE clients
SET 
    is_active = false,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: RegenerateClientSecret :one
UPDATE clients
SET 
    client_secret = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND is_active = true
RETURNING *;

-- name: RegenerateClientSecretByClientId :one
UPDATE clients
SET 
    client_secret = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE client_id = $1 AND is_active = true
RETURNING *;

-- name: CountClients :one
SELECT COUNT(*) FROM clients WHERE is_active = true;