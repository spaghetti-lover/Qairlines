-- name: CreateAirplane :one
INSERT INTO airplane (
  airplane_model_id,
  registration_number
) VALUES (
  $1, $2
) RETURNING *;


-- name: GetAirplane :one
SELECT * FROM airplane
WHERE registration_number = $1 LIMIT 1;

-- name: ListAirplanes :many
SELECT * FROM airplane
ORDER BY registration_number
LIMIT $1
OFFSET $2;


-- name: DeleteAirplane :exec
DELETE FROM airplane
WHERE registration_number = $1;