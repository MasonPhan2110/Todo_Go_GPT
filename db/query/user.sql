-- name: CreateUser :one
INSERT INTO "user" (
    username,
    hashed_password,
    full_name,
    email
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetUser :one
SELECT * FROM "user" 
WHERE username = $1 LIMIT 1;

-- name: GetUserForUpdate :one
SELECT * FROM "user" 
WHERE username = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: UpdateUserHashedPassword :one
UPDATE "user"
SET hashed_password = $2, update_at = (now())
WHERE id = $1
RETURNING *; 
