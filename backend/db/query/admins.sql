-- name: CreateAdmin :one
INSERT INTO admins (
  user_id
) VALUES (
  $1
) RETURNING *;

-- name: GetAdmin :one
SELECT * FROM admins
WHERE user_id = $1 LIMIT 1;

-- name: GetAdminByEmail :one
SELECT a.* FROM admins
JOIN users u ON a.user_id = u.user_id
WHERE u.email = $1 LIMIT 1;

-- name: GetAllAdmin :many
SELECT * FROM admins;

-- name: ListAdmins :many
SELECT * FROM admins
ORDER BY user_id
LIMIT $1
OFFSET $2;

-- name: DeleteAdmin :exec
DELETE FROM admins
WHERE user_id = $1;

-- name: IsAdmin :one
SELECT EXISTS (
  SELECT 1
  FROM admins
  WHERE user_id = $1
) AS is_admin;