-- name: CreateUser :one
INSERT INTO users (
    name,
    email,
    password,
    createdAt,
    updatedAt
)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, name, email, createdAt, updatedAt;


-- name: GetUser :one
SELECT id, name, email, createdAt, updatedAt
FROM users
WHERE id = $1;


-- name: GetUserByEmail :one
SELECT id, name, email, password, createdAt, updatedAt
FROM users
WHERE email = $1;


-- name: UpdateUser :one
UPDATE users
SET
    name = $2,
    email = $3,
    password = $4,
    updatedAt = $5
WHERE id = $1
RETURNING id, name, email, createdAt, updatedAt;


-- name: UpdateUserProfile :one
UPDATE users
SET
    name = $2,
    email = $3,
    updatedAt = $4
WHERE id = $1
RETURNING id, name, email, createdAt, updatedAt;


-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
