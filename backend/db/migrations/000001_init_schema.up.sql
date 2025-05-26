-- Define ENUM types
CREATE TYPE user_role AS ENUM ('customer', 'admin');
CREATE TYPE gender_type AS ENUM ('male', 'female', 'other');
CREATE TYPE trip_type AS ENUM ('oneWay', 'roundTrip');
CREATE TYPE booking_status AS ENUM ('confirmed', 'cancelled', 'pending');
CREATE TYPE ticket_status AS ENUM ('booked', 'cancelled', 'used');
CREATE TYPE flight_status AS ENUM ('On Time', 'Delayed', 'Cancelled', 'Boarding', 'Takeoff', 'Landing', 'Landed');
CREATE TYPE flight_class AS ENUM ('Economy', 'Business', 'First Class');

-- Create Users table
CREATE TABLE Users (
  user_id BIGSERIAL PRIMARY KEY,
  email VARCHAR(255) UNIQUE NOT NULL,
  hashed_password VARCHAR(255) NOT NULL,
  first_name VARCHAR(100),
  last_name VARCHAR(100),
  role user_role NOT NULL DEFAULT 'customer',
  is_active BOOLEAN NOT NULL DEFAULT TRUE,
  created_at timestamptz NOT NULL DEFAULT (now()),
  updated_at timestamptz NOT NULL DEFAULT (now())
);

-- Create Customers table
CREATE TABLE Customers (
  user_id BIGINT PRIMARY KEY REFERENCES Users(user_id) ON DELETE CASCADE,
  phone_number VARCHAR(20),
  gender gender_type NOT NULL DEFAULT 'other',
  date_of_birth DATE,
  passport_number VARCHAR(50),
  identification_number VARCHAR(50),
  address TEXT,
  loyalty_points INT DEFAULT 0 CHECK (loyalty_points >= 0),
  created_at timestamptz NOT NULL DEFAULT (now()),
  updated_at timestamptz NOT NULL DEFAULT (now())
);

-- Create Admins table
CREATE TABLE Admins (
  user_id BIGINT PRIMARY KEY REFERENCES Users(user_id) ON DELETE CASCADE
);

-- Create News table
CREATE TABLE News (
  id BIGSERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  content TEXT,
  image TEXT,
  author_id BIGINT REFERENCES Users(user_id) ON DELETE SET NULL,
  created_at timestamptz NOT NULL DEFAULT (now()),
  updated_at timestamptz NOT NULL DEFAULT (now())
);

-- Create Flights table
CREATE TABLE Flights (
  flight_id BIGSERIAL PRIMARY KEY,
  flight_number VARCHAR(20) UNIQUE NOT NULL,
  aircraft_type VARCHAR(100),
  departure_city VARCHAR(100),
  arrival_city VARCHAR(100),
  departure_airport VARCHAR(255),
  arrival_airport VARCHAR(255),
  departure_time TIMESTAMPTZ NOT NULL,
  arrival_time TIMESTAMPTZ NOT NULL,
  base_price INT NOT NULL CHECK (base_price >= 0),
  total_seats_row INT NOT NULL CHECK (total_seats_row > 0) DEFAULT 44,
  total_seats_column INT NOT NULL CHECK (total_seats_column > 0) DEFAULT 6,
  status flight_status NOT NULL DEFAULT 'On Time'
);

-- Create Bookings table
CREATE TABLE Bookings (
  booking_id BIGSERIAL PRIMARY KEY,
  user_email VARCHAR(255) REFERENCES Users(email) ON DELETE SET NULL,
  trip_type trip_type NOT NULL,
  departure_flight_id BIGINT NOT NULL REFERENCES Flights(flight_id),
  return_flight_id BIGINT REFERENCES Flights(flight_id) CHECK (
    trip_type = 'oneWay' AND return_flight_id IS NULL OR
    trip_type = 'roundTrip' AND return_flight_id IS NOT NULL
  ),
  status booking_status NOT NULL DEFAULT 'pending',
  created_at timestamptz NOT NULL DEFAULT (now()),
  updated_at timestamptz NOT NULL DEFAULT (now())
);
-- Create Tickets table
CREATE TABLE Tickets (
  ticket_id BIGSERIAL PRIMARY KEY,
  flight_class flight_class NOT NULL DEFAULT 'Economy',
  price INT NOT NULL CHECK (price >= 0),
  status ticket_status NOT NULL DEFAULT 'booked',
  booking_id BIGINT REFERENCES Bookings(booking_id) ON DELETE CASCADE,
  flight_id BIGINT REFERENCES Flights(flight_id),
  created_at timestamptz NOT NULL DEFAULT (now()),
  updated_at timestamptz NOT NULL DEFAULT (now())
);

-- Create TicketOwnerSnapshot table
CREATE TABLE TicketOwnerSnapshot (
  ticket_id BIGSERIAL PRIMARY KEY REFERENCES Tickets(ticket_id) ON DELETE CASCADE,
  first_name VARCHAR(100),
  last_name VARCHAR(100),
  phone_number VARCHAR(20),
  gender gender_type NOT NULL DEFAULT 'other'
);

CREATE TABLE Seats (
  seat_id BIGSERIAL PRIMARY KEY,
  flight_id BIGINT REFERENCES Flights(flight_id) ON DELETE CASCADE,
  seat_code VARCHAR(3) NOT NULL,
  is_available BOOLEAN NOT NULL DEFAULT TRUE,
  class flight_class NOT NULL DEFAULT 'Economy'
);

-- Create indexes for Flights
CREATE INDEX idx_flights_departure_time ON Flights (departure_time);
CREATE INDEX idx_flights_arrival_time ON Flights (arrival_time);
CREATE INDEX idx_flights_status ON Flights (status);

-- Create indexes for Bookings
CREATE INDEX idx_bookings_user_id ON Bookings (user_email);
CREATE INDEX idx_bookings_departure_flight_id ON Bookings (departure_flight_id);
CREATE INDEX idx_bookings_return_flight_id ON Bookings (return_flight_id);

-- Create indexes for Tickets
CREATE INDEX idx_tickets_booking_id ON Tickets (booking_id);
CREATE INDEX idx_tickets_flight_id ON Tickets (flight_id);