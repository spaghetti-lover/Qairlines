-- name: CreateFlight :one
INSERT INTO flights (
  flight_id,
  flight_number,
  aircraft_type,
  departure_city,
  arrival_city,
  departure_airport,
  arrival_airport,
  departure_time,
  arrival_time,
  base_price,
  status
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
) RETURNING *;


-- name: GetFlight :one
SELECT * FROM flights
WHERE flight_id = $1 LIMIT 1;

-- name: ListFlights :many
SELECT * FROM flights
ORDER BY flight_id
LIMIT $1
OFFSET $2;

-- name: GetFlightsByStatus :one
SELECT status FROM flights
WHERE flight_id = $1 LIMIT 1;

-- name: DeleteFlight :exec
DELETE FROM flights
WHERE flight_id = $1;