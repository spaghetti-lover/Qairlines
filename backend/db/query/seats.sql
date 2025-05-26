-- name: CreateSeat :one
INSERT INTO "seats" (
  flight_id,
  seat_code,
  is_available,
  class
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetSeat :one
SELECT * FROM "seats"
WHERE seat_id = $1 LIMIT 1;

-- name: GetAllSeats :many
SELECT * FROM "seats";

-- name: ListSeatsWithFlightId :many
SELECT * FROM "seats"
WHERE flight_id = $1;

-- name: UpdateSeat :exec
UPDATE "seats"
SET seat_code = $2,
    is_available = $3,
    class = $4,
    updated_at = NOW()
WHERE seat_id = $1;

-- name: CountOccupiedSeats :one
SELECT COUNT(*) FROM "seats"
WHERE flight_id = $1 AND is_available = false;

-- name: CheckSeatAvailability :one
SELECT is_available FROM "seats"
WHERE seat_code = $1 and flight_id = $2;

-- name: MarkSeatUnavailable :exec
UPDATE "seats"
SET is_available = false
WHERE seat_code = $1 AND flight_id = $2;