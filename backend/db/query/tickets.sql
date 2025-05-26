-- name: CreateTicket :one
INSERT INTO tickets (
  flight_class,
  price,
  status,
  booking_id,
  flight_id
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetTicket :one
SELECT * FROM tickets
WHERE ticket_id = $1 LIMIT 1;

-- name: ListTickets :many
SELECT * FROM tickets
ORDER BY ticket_id
LIMIT $1
OFFSET $2;

-- name: GetTicketByFlightId :many
SELECT * FROM tickets
WHERE flight_id = $1
ORDER BY ticket_id;

-- name: DeleteTicket :exec
DELETE FROM tickets
WHERE ticket_id = $1;

-- name: UpdateTicket :exec
UPDATE tickets
SET
  flight_class = $2,
  price = $3,
  status = $4,
  booking_id = $5,
  flight_id = $6,
  updated_at = NOW()
WHERE ticket_id = $1;

-- name: GetTicketsByFlightID :many
SELECT
    t.ticket_id,
    t.status,
    t.flight_class,
    t.price,
    t.booking_id,
    t.flight_id,
    t.created_at,
    t.updated_at,
    tos.first_name,
    tos.last_name,
    tos.phone_number,
    tos.gender,
    s.seat_code,
    s.class AS seat_class
FROM Tickets t
LEFT JOIN TicketOwnerSnapshot tos ON t.ticket_id = tos.ticket_id
LEFT JOIN Seats s ON t.flight_id = s.flight_id
WHERE t.flight_id = $1;

-- name: UpdateTicketStatus :one
UPDATE Tickets
SET status = $2, updated_at = NOW()
WHERE ticket_id = $1
RETURNING *;

-- name: GetTicketByID :one
SELECT
    t.ticket_id,
    t.seat_id,
    t.status,
    t.flight_class,
    t.price,
    t.booking_id,
    t.flight_id,
    t.created_at,
    t.updated_at,
    s.seat_code,
    s.class AS seat_class,
    tos.first_name,
    tos.last_name,
    tos.phone_number,
    tos.gender
FROM Tickets t
JOIN Seats s ON t.seat_id = s.seat_id
JOIN TicketOwnerSnapshot tos ON t.ticket_id = tos.ticket_id
WHERE t.ticket_id = $1
LIMIT 1;