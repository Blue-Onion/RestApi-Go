-- name: CreateUser :one
INSERT INTO users (
    name,
    email,
    password
)
VALUES ($1, $2, $3)
RETURNING
    id; 

-- name: GetUser :one
SELECT
    id,
    name,
    email,
    createdAt,
    updatedAt
FROM users
WHERE id = $1;


-- name: GetUserByEmail :one
SELECT
    id,
    name,
    email,
    password,
    FROM users
WHERE email = $1;

-- name: UpdateUser :one
UPDATE users
SET
    name = COALESCE(sqlc.narg(name), name),
    email = COALESCE(sqlc.narg(email), email),
WHERE id = sqlc.arg(id)
RETURNING id;


-- name: UpdateUserPassword :one
UPDATE users
SET
    password= $2
WHERE id = $1
RETURNING id, name, email;


-- name: DeleteUser :one
DELETE FROM users
WHERE id = $1
RETURNING id;
