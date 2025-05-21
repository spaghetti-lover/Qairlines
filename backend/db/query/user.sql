-- name: CreateUser :one
INSERT INTO "user" (
  username,
  hashed_password,
  role,
  email
) VALUES (
  $1, $2, $3, $4
) RETURNING *;


-- name: GetUser :one
SELECT * FROM "user"
WHERE user_id = $1 LIMIT 1;

-- name: GetAllUser :many
SELECT * FROM "user";

-- name: ListUsers :many
SELECT * FROM "user"
ORDER BY user_id
LIMIT $1
OFFSET $2;


-- name: DeleteUser :exec
DELETE FROM "user"
WHERE user_id = $1;