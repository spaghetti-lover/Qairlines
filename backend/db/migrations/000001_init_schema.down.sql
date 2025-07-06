-- Drop triggers
DROP TRIGGER IF EXISTS set_updated_at ON Users;
DROP TRIGGER IF EXISTS set_updated_at ON News;
DROP TRIGGER IF EXISTS set_updated_at ON Bookings;
DROP TRIGGER IF EXISTS set_updated_at ON Tickets;
-- Drop trigger functions
DROP FUNCTION IF EXISTS update_user_updated_at_column;
DROP FUNCTION IF EXISTS update_news_updated_at_column;
DROP FUNCTION IF EXISTS update_booking_updated_at_column;
DROP FUNCTION IF EXISTS update_ticket_updated_at_column;
-- Drop tables (in reverse dependency order)
DROP TABLE IF EXISTS TicketOwnerSnapshots;
DROP TABLE IF EXISTS Tickets;
DROP TABLE IF EXISTS Seats;
DROP TABLE IF EXISTS Bookings;
DROP TABLE IF EXISTS Flights;
DROP TABLE IF EXISTS Admins;
DROP TABLE IF EXISTS Customers;
DROP TABLE IF EXISTS News;
DROP TABLE IF EXISTS Users;
-- Drop ENUM types
DROP TYPE IF EXISTS user_role;
DROP TYPE IF EXISTS gender_type;
DROP TYPE IF EXISTS trip_type;
DROP TYPE IF EXISTS booking_status;
DROP TYPE IF EXISTS ticket_status;
DROP TYPE IF EXISTS flight_status;
DROP TYPE IF EXISTS flight_class;