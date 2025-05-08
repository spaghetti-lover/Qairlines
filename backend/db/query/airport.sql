-- name: CreateAirport :one
INSERT INTO airport (
  airport_code,
  city,
  name
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetAirport :one
SELECT * FROM airport
WHERE airport_code = $1 LIMIT 1;

-- name: ListAirports :many
SELECT * FROM airport
ORDER BY airport_code
LIMIT $1
OFFSET $2;

-- name: DeleteAirport :exec
DELETE FROM airport
WHERE airport_code = $1;