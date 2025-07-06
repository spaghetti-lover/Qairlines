-- name: CreateTicketOwnerSnapshot :one
INSERT INTO TicketOwnerSnapshots (
  ticket_id, first_name, last_name, phone_number, gender, date_of_birth,
  passport_number, identification_number, address
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: GetTicketOwnerSnapshot :one
SELECT * FROM TicketOwnerSnapshots
WHERE ticket_id = $1
LIMIT 1;

-- name: GetAllTicketOwnerSnapshots :many
SELECT * FROM TicketOwnerSnapshots;

-- name: ListTicketOwnerSnapshots :many
SELECT * FROM TicketOwnerSnapshots
ORDER BY ticket_id
LIMIT $1
OFFSET $2;