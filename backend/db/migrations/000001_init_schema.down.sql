-- Xóa các ràng buộc đã thêm trong file up
ALTER TABLE Flights DROP CONSTRAINT IF EXISTS departure_before_arrival;
ALTER TABLE Seats DROP CONSTRAINT IF EXISTS unique_flight_seat;

-- Drop indexes for Tickets
DROP INDEX IF EXISTS idx_tickets_flight_id;
DROP INDEX IF EXISTS idx_tickets_booking_id;

-- Drop indexes for Bookings
DROP INDEX IF EXISTS idx_bookings_return_flight_id;
DROP INDEX IF EXISTS idx_bookings_departure_flight_id;
DROP INDEX IF EXISTS idx_bookings_user_id;

-- Drop indexes for Flights
DROP INDEX IF EXISTS idx_flights_status;
DROP INDEX IF EXISTS idx_flights_arrival_time;
DROP INDEX IF EXISTS idx_flights_departure_time;

-- Drop TicketOwnerSnapshot table
DROP TABLE IF EXISTS TicketOwnerSnapshot;

-- Drop Tickets table
DROP TABLE IF EXISTS Tickets;

-- Drop Seats table
DROP TABLE IF EXISTS Seats;

-- Drop Bookings table
DROP TABLE IF EXISTS Bookings;

-- Drop Flights table
DROP TABLE IF EXISTS Flights;

-- Drop News table
DROP TABLE IF EXISTS News;

-- Drop Admins table
DROP TABLE IF EXISTS Admins;

-- Drop Customers table
DROP TABLE IF EXISTS Customers;

-- Drop Users table
DROP TABLE IF EXISTS Users;

-- Drop ENUM types
DROP TYPE IF EXISTS user_role;
DROP TYPE IF EXISTS ticket_status;
DROP TYPE IF EXISTS booking_status;
DROP TYPE IF EXISTS trip_type;
DROP TYPE IF EXISTS gender_type;
DROP TYPE IF EXISTS flight_status;
DROP TYPE IF EXISTS flight_class;
