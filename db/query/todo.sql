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

-- name: UpdateTask :one
UPDATE todo
SET
  "name" = COALESCE(sqlc.narg(name), "name"),
  description = COALESCE(sqlc.narg(description), description),
  status = COALESCE(sqlc.narg(status), status),
  deadline = COALESCE(sqlc.narg(deadline), deadline),
  update_at = (now())
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteTask :exec
DELETE FROM todo
WHERE id = $1;
