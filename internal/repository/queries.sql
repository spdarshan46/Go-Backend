-- name: CreateUser :one
INSERT INTO users (name, dob)
VALUES ($1, $2)
RETURNING id, name, dob, created_at, updated_at;

-- name: GetUserByID :one
SELECT id, name, dob, created_at, updated_at
FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
SET name = COALESCE($2, name),
    dob = COALESCE($3, dob),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, name, dob, created_at, updated_at;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT id, name, dob, created_at, updated_at
FROM users
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: CountUsers :one
SELECT COUNT(*) FROM users;

-- name: CheckUserExists :one
SELECT EXISTS(SELECT 1 FROM users WHERE id = $1);