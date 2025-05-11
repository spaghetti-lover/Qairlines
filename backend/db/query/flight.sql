-- name: CreateFlight :one
INSERT INTO flight (
  flight_number,
  registration_number,
  estimated_departure_time,
  actual_departure_time,
  estimated_arrival_time,
  actual_arrival_time,
  departure_airport_id,
  destination_airport_id,
  flight_price,
  status
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) RETURNING *;


-- name: GetFlight :one
SELECT * FROM flight
WHERE flight_id = $1 LIMIT 1;

-- name: ListFlights :many
SELECT * FROM flight
ORDER BY flight_number
LIMIT $1
OFFSET $2;


-- name: DeleteFlight :exec
DELETE FROM flight
WHERE flight_number = $1;