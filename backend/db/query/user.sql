-- name: CreateUser :one
INSERT INTO "user" (
  username,
  password,
  role
) VALUES (
  $1, $2, $3
) RETURNING *;


-- name: GetUser :one
SELECT * FROM "user"
WHERE user_id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM "user"
ORDER BY user_id
LIMIT $1
OFFSET $2;


-- name: DeleteUser :exec
DELETE FROM "user"
WHERE user_id = $1;