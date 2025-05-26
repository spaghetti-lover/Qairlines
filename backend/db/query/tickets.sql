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