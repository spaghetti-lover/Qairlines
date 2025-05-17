-- name: CreateAdmin :one
INSERT INTO admin (
  username,
  password
) VALUES (
  $1, $2
) RETURNING *;


-- name: GetAdmin :one
SELECT * FROM admin
WHERE admin_id = $1 LIMIT 1;

-- name: ListAdmins :many
SELECT * FROM admin
ORDER BY admin_id
LIMIT $1
OFFSET $2;


-- name: DeleteAdmin :exec
DELETE FROM admin
WHERE admin_id = $1;