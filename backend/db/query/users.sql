-- name: CreateUser :one
INSERT INTO users (
  email,
  hashed_password,
  first_name,
  last_name,
  role
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetUser :one
SELECT *
FROM users
WHERE user_id = $1;

-- name: GetAllUser :many
SELECT * FROM users;

-- name: ListUsers :many
SELECT *
FROM users
ORDER BY user_id
LIMIT $1
OFFSET $2;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;


-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = $1;

-- name: UpdatePassword :exec
UPDATE users
SET hashed_password = $2
WHERE user_id = $1;

-- name: UpdateUser :exec
UPDATE users
SET first_name = $2,
    last_name = $3,
    updated_at = NOW()
WHERE user_id = $1;

-- name: DeactivateUser :exec
UPDATE Users
SET is_active = false,
    updated_at = now()
WHERE user_id = $1;