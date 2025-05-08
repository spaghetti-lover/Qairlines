-- name: CreateAirplaneModel :one
INSERT INTO airplane_model (
  name,
  manufacturer,
  total_seats
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetAirplaneModel :one
SELECT * FROM airplane_model
WHERE airplane_model_id = $1 LIMIT 1;

-- name: ListAirplaneModels :many
SELECT * FROM airplane_model
ORDER BY airplane_model_id
LIMIT $1
OFFSET $2;


-- name: DeleteAirplaneModel :exec
DELETE FROM airplane_model
WHERE airplane_model_id = $1;