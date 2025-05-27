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

-- name: GetSeatByTicketID :one
SELECT s.seat_id, s.seat_code, s.class, s.is_available
FROM Seats s
JOIN Tickets t ON s.seat_id = t.seat_id
WHERE t.ticket_id = $1;


-- name: UpdateSeatAvailability :exec
UPDATE Seats
SET is_available = $2
WHERE seat_id = (SELECT seat_id FROM Tickets WHERE ticket_id = $1);