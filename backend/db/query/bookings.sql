-- name: CreateBooking :one
INSERT INTO bookings (
  user_email,
  trip_type,
  departure_flight_id,
  return_flight_id,
  status
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetBooking :one
SELECT * FROM bookings
WHERE booking_id = $1 LIMIT 1;

-- name: ListBookings :many
SELECT * FROM bookings
ORDER BY booking_id
LIMIT $1
OFFSET $2;

-- name: GetBookingHistoryByUID :many
SELECT booking_id
FROM Bookings
JOIN Users u ON Bookings.user_email = u.email
WHERE u.user_id = $1;


-- name: DeleteBookings :exec
DELETE FROM bookings
WHERE booking_id = $1;

-- name: RemoveUserFromBookings :exec
UPDATE bookings
SET user_email = NULL,
    updated_at = NOW()
WHERE user_email = $1;
