-- name: CreateBooking :one
INSERT INTO booking (
  booker_email,
  number_of_adults,
  number_of_children,
  flight_class,
  cancelled,
  flight_id
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetBooking :one
SELECT * FROM booking
WHERE booking_id = $1 LIMIT 1;

-- name: ListBookings :many
SELECT * FROM booking
ORDER BY booking_id
LIMIT $1
OFFSET $2;


-- name: DeleteBookings :exec
DELETE FROM booking
WHERE booking_id = $1;