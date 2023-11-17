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

-- name: UpdateUser :one
UPDATE "user"
SET
  hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
  password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at),
  full_name = COALESCE(sqlc.narg(full_name), full_name),
  email = COALESCE(sqlc.narg(email), email),
  is_email_verified = COALESCE(sqlc.narg(is_email_verified), is_email_verified)
WHERE
  id = sqlc.arg(id)
RETURNING *;
