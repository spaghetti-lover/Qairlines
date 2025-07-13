-- name: CreateFlight :one
INSERT INTO flights (
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
  )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;
-- name: GetFlight :one
SELECT *
FROM flights
WHERE flight_id = $1
LIMIT 1;
-- name: GetAllFlights :many
SELECT flight_id,
  flight_number,
  aircraft_type,
  departure_city,
  arrival_city,
  departure_time,
  arrival_time,
  base_price,
  status
FROM Flights;
-- name: ListFlights :many
SELECT flight_id,
  flight_number,
  airline,
  departure_city,
  arrival_city,
  departure_time,
  arrival_time,
  departure_airport,
  arrival_airport,
  aircraft_type,
  base_price
FROM flights
ORDER BY flight_id
LIMIT $1 OFFSET $2;
-- name: GetFlightsByStatus :one
SELECT status
FROM flights
WHERE flight_id = $1
LIMIT 1;
-- name: DeleteFlight :one
DELETE FROM flights
WHERE flight_id = $1
RETURNING flight_id;
-- name: UpdateFlightTimes :one
UPDATE Flights
SET departure_time = $2,
  arrival_time = $3
WHERE flight_id = $1
RETURNING flight_id,
  departure_time,
  arrival_time;
-- name: SearchFlights :many
SELECT flight_id,
  flight_number,
  airline,
  departure_city,
  arrival_city,
  departure_time,
  arrival_time,
  departure_airport,
  arrival_airport,
  aircraft_type,
  base_price
FROM Flights
WHERE departure_city = $1
  AND arrival_city = $2
  AND DATE(departure_time) = $3;