-- name: CreateTicket :one
INSERT INTO tickets (
    seat_id,
    flight_class,
    price,
    status,
    booking_id,
    flight_id
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetTicketByID :one
SELECT
    t.ticket_id,
    t.status,
    t.flight_class,
    t.price,
    t.booking_id,
    t.flight_id,
    t.created_at,
    t.updated_at,
    s.seat_code,
    s.seat_id,
    s.is_available,
    s.flight_id,
    s.class AS seat_class,
    o.first_name AS owner_first_name,
    o.last_name AS owner_last_name,
    o.gender AS owner_gender,
    o.phone_number AS owner_phone_number,
    o.date_of_birth AS owner_date_of_birth,
    o.passport_number AS owner_passport_number,
    o.identification_number AS owner_identification_number,
    o.address AS owner_address
FROM Tickets t
LEFT JOIN Seats s ON t.seat_id = s.seat_id
LEFT JOIN TicketOwnerSnapshot o ON t.ticket_id = o.ticket_id
WHERE t.ticket_id = $1;

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
    t.seat_id,
    t.flight_class,
    t.price,
    t.status,
    t.booking_id,
    t.flight_id,
    t.created_at,
    t.updated_at,
    s.seat_code,
    s.is_available,
    s.class AS seat_class,
    o.first_name AS owner_first_name,
    o.last_name AS owner_last_name,
    o.phone_number AS owner_phone_number,
    o.gender AS owner_gender,
    o.date_of_birth AS owner_date_of_birth,
    o.passport_number AS owner_passport_number,
    o.identification_number AS owner_identification_number,
    o.address AS owner_address
FROM Tickets t
LEFT JOIN Seats s ON t.seat_id = s.seat_id
LEFT JOIN TicketOwnerSnapshot o ON t.ticket_id = o.ticket_id
WHERE t.flight_id = $1;

-- name: UpdateTicketStatus :one
UPDATE Tickets
SET status = $2, updated_at = NOW()
WHERE ticket_id = $1
RETURNING *;

-- name: CancelTicket :one
UPDATE Tickets
SET status = 'Cancelled', updated_at = NOW()
WHERE Tickets.ticket_id = $1 AND status = 'Active'
RETURNING ticket_id, status, flight_class, price, booking_id, flight_id, updated_at,
          (SELECT seat_code FROM Seats WHERE seat_id = Tickets.seat_id) AS seat_code,
          (SELECT first_name FROM TicketOwnerSnapshot WHERE ticket_id = Tickets.ticket_id) AS owner_first_name,
          (SELECT last_name FROM TicketOwnerSnapshot WHERE ticket_id = Tickets.ticket_id) AS owner_last_name,
          (SELECT phone_number FROM TicketOwnerSnapshot WHERE ticket_id = Tickets.ticket_id) AS owner_phone_number;

-- name: UpdateSeat :one
UPDATE Seats
SET
    seat_code = $2
WHERE seat_id = (SELECT seat_id FROM Tickets WHERE ticket_id = $1 AND seat_id IS NOT NULL)
RETURNING *;

-- name: GetTicketsByBookingIDAndType :many
SELECT
    ticket_id,
    seat_id,
    flight_class,
    price,
    status,
    booking_id,
    flight_id,
    created_at,
    updated_at
FROM
    Tickets
WHERE
    Tickets.booking_id = $1
    AND (
        ($2 = 'departure' AND flight_id = (SELECT departure_flight_id FROM Bookings WHERE Bookings.booking_id = $1))
        OR ($2 = 'return' AND flight_id = (SELECT return_flight_id FROM Bookings WHERE Bookings.booking_id = $1))
    );