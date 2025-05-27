  -- booking_id VARCHAR(30) PRIMARY KEY,
  -- user_email VARCHAR(255) REFERENCES Users(email) ON DELETE SET NULL,
  -- trip_type trip_type NOT NULL,
  -- departure_flight_id VARCHAR(20) REFERENCES Flights(flight_id),
  -- return_flight_id VARCHAR(20) REFERENCES Flights(flight_id),
  -- status booking_status NOT NULL DEFAULT 'pending',
  -- created_at timestamptz NOT NULL DEFAULT (now()),
  -- updated_at timestamptz NOT NULL DEFAULT (now())

-- name: CreateBooking :one
INSERT INTO bookings (
  booking_id,
  user_email,
  trip_type,
  departure_flight_id,
  return_flight_id,
  status
) VALUES (
  $1, $2, $3, $4, $5, $6
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
