-- name: CreateTask :one
INSERT INTO "todo" (
    "user_id",
    "name",
    "description",
    "status",
    "deadline"
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetTask :one
SELECT * FROM todo
WHERE id = $1 LIMIT 1;

-- name: GetTaskForUpdate :one
SELECT * FROM todo
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListTasks :many
SELECT * FROM todo
WHERE "user_id" = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateStatus :one
UPDATE todo
SET status = $2
WHERE id = $1
RETURNING *;

-- name: UpdateDeadline :one
UPDATE todo
SET deadline = $2
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM todo
WHERE id = $1;
