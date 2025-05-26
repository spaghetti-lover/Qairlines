-- name: CreateTicketOwnerSnapshot :one
INSERT INTO ticketownersnapshot (
  ticket_id,
  first_name,
  last_name,
  phone_number,
  gender
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetTicketOwnerSnapshot :one
SELECT * FROM ticketownersnapshot
WHERE ticket_id = $1 LIMIT 1;

-- name: GetAllTicketOwnerSnapshots :many
SELECT * FROM ticketownersnapshot;

-- name: ListTicketOwnerSnapshots :many
SELECT * FROM ticketownersnapshot
ORDER BY ticket_id
LIMIT $1
OFFSET $2;