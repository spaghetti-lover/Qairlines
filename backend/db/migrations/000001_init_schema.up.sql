-- Define ENUM types
CREATE TYPE user_role AS ENUM ('customer', 'admin');
CREATE TYPE gender_type AS ENUM ('Male', 'Female', 'Other');
CREATE TYPE trip_type AS ENUM ('oneWay', 'roundTrip');
CREATE TYPE booking_status AS ENUM ('confirmed', 'cancelled', 'pending');
CREATE TYPE ticket_status AS ENUM ('booked', 'cancelled', 'used');
CREATE TYPE flight_status AS ENUM ('On Time', 'Delayed', 'Cancelled', 'Boarding', 'Takeoff', 'Landing', 'Landed');
CREATE TYPE flight_class AS ENUM ('economy', 'business', 'firstClass');

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
  gender gender_type NOT NULL DEFAULT 'Other',
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
  airline VARCHAR(100),
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
  departure_flight_id BIGINT REFERENCES Flights(flight_id) ON DELETE CASCADE,
  return_flight_id BIGINT REFERENCES Flights(flight_id) ON DELETE CASCADE CHECK (
    trip_type = 'oneWay' AND return_flight_id IS NULL OR
    trip_type = 'roundTrip' AND return_flight_id IS NOT NULL
  ),
  status booking_status NOT NULL DEFAULT 'pending',
  created_at timestamptz NOT NULL DEFAULT (now()),
  updated_at timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE Seats (
  seat_id BIGSERIAL PRIMARY KEY,
  flight_id BIGINT REFERENCES Flights(flight_id) ON DELETE CASCADE,
  seat_code VARCHAR(3) NOT NULL,
  is_available BOOLEAN NOT NULL DEFAULT TRUE,
  class flight_class NOT NULL DEFAULT 'economy',
  UNIQUE (flight_id, seat_code)
);

-- Create Tickets table
CREATE TABLE Tickets (
  ticket_id BIGSERIAL PRIMARY KEY,
  seat_id BIGINT REFERENCES Seats(seat_id) ON DELETE CASCADE,
  flight_class flight_class NOT NULL DEFAULT 'economy',
  price INT NOT NULL CHECK (price >= 0),
  status ticket_status NOT NULL DEFAULT 'booked',
  booking_id BIGINT REFERENCES Bookings(booking_id) ON DELETE CASCADE,
  flight_id BIGINT REFERENCES Flights(flight_id) ON DELETE CASCADE NOT NULL,
  created_at timestamptz NOT NULL DEFAULT (now()),
  updated_at timestamptz NOT NULL DEFAULT (now())
);

-- Create TicketOwnerSnapshot table
CREATE TABLE TicketOwnerSnapshot (
  ticket_id BIGSERIAL PRIMARY KEY REFERENCES Tickets(ticket_id) ON DELETE CASCADE,
  first_name VARCHAR(100),
  last_name VARCHAR(100),
  phone_number VARCHAR(20),
  gender gender_type NOT NULL DEFAULT 'Other',
  date_of_birth DATE,
  passport_number VARCHAR(50),
  identification_number VARCHAR(50),
  address TEXT
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

-- Thêm ràng buộc UNIQUE trong bảng Seats
ALTER TABLE Seats ADD CONSTRAINT unique_flight_seat UNIQUE (flight_id, seat_code);

-- Đảm bảo ngày đi trước ngày đến
ALTER TABLE Flights ADD CONSTRAINT departure_before_arrival CHECK (departure_time < arrival_time);

INSERT INTO Users (email, hashed_password, first_name, last_name, role, is_active)
VALUES
('john.doe@example.com', 'hashed_password_1', 'John', 'Doe', 'customer', TRUE),
('jane.smith@example.com', 'hashed_password_2', 'Jane', 'Smith', 'admin', TRUE),
('alice.wonderland@example.com', 'hashed_password_3', 'Alice', 'Wonderland', 'customer', TRUE),
('phungducanh2511@gmail.com', '123456', 'Phung', 'Duc Anh', 'customer', TRUE);

INSERT INTO Customers (user_id, phone_number, gender, date_of_birth, passport_number, identification_number, address, loyalty_points)
VALUES
(1, '1234567890', 'Male', '1990-01-01', 'A12345678', 'ID123456', '123 Main St, Cityville', 100),
(3, '0987654321', 'Female', '1995-05-15', 'B98765432', 'ID987654', '456 Elm St, Townsville', 200);

INSERT INTO Admins (user_id)
VALUES
(4);

INSERT INTO Flights (flight_number, airline, aircraft_type, departure_city, arrival_city, departure_airport, arrival_airport, departure_time, arrival_time, base_price, total_seats_row, total_seats_column, status)
VALUES
('VN101', 'Vietnam Airlines', 'Airbus A320', 'Hanoi', 'Ho Chi Minh City', 'Noi Bai', 'Tan Son Nhat', '2025-05-28 08:00:00', '2025-05-28 10:00:00', 1000000, 44, 6, 'On Time'),
('VN102', 'Vietnam Airlines', 'Boeing 737', 'Da Nang', 'Hanoi', 'Da Nang Airport', 'Noi Bai', '2025-05-28 14:00:00', '2025-05-28 16:00:00', 800000, 44, 6, 'Delayed'),
('VN123', 'Vietnam Airlines', 'Boeing 787', 'Hanoi', 'Ho Chi Minh City', 'Noi Bai', 'Tan Son Nhat', '2025-06-01 08:00:00', '2025-06-01 10:00:00', 1000000, 44, 6, 'On Time'),
('VN456', 'Vietnam Airlines', 'Airbus A320', 'Da Nang', 'Hanoi', 'Da Nang Airport', 'Noi Bai', '2025-06-02 14:00:00', '2025-06-02 16:00:00', 800000, 44, 6, 'Delayed'),
('VN789', 'Vietnam Airlines', 'Airbus A321', 'Ho Chi Minh City', 'Hanoi', 'Tan Son Nhat', 'Noi Bai', '2025-05-28 17:00:00', '2025-05-28 19:00:00', 1200000, 44, 6, 'On Time'),
('VN234', 'Vietnam Airlines', 'Boeing 737', 'Hanoi', 'Da Nang', 'Noi Bai', 'Da Nang Airport', '2025-06-03 09:00:00', '2025-06-03 10:30:00', 700000, 44, 6, 'On Time'),
('VN567', 'Vietnam Airlines', 'Airbus A330', 'Ho Chi Minh City', 'Da Nang', 'Tan Son Nhat', 'Da Nang Airport', '2025-06-04 15:00:00', '2025-06-04 16:30:00', 900000, 44, 6, 'On Time'),
('VN890', 'Vietnam Airlines', 'Boeing 777', 'Da Nang', 'Ho Chi Minh City', 'Da Nang Airport', 'Tan Son Nhat', '2025-06-05 18:00:00', '2025-06-05 19:30:00', 850000, 44, 6, 'Delayed'),
('VN345', 'Vietnam Airlines', 'Airbus A350', 'Hanoi', 'Phu Quoc', 'Noi Bai', 'Phu Quoc Airport', '2025-06-06 07:00:00', '2025-06-06 09:30:00', 1500000, 44, 6, 'On Time'),
('VN678', 'Vietnam Airlines', 'Boeing 737', 'Phu Quoc', 'Ho Chi Minh City', 'Phu Quoc Airport', 'Tan Son Nhat', '2025-06-07 11:00:00', '2025-06-07 12:30:00', 600000, 44, 6, 'On Time'),
('VN901', 'Vietnam Airlines', 'Airbus A320', 'Nha Trang', 'Hanoi', 'Cam Ranh Airport', 'Noi Bai', '2025-06-08 13:00:00', '2025-06-08 15:30:00', 1100000, 44, 6, 'On Time'),
('VN432', 'Vietnam Airlines', 'Boeing 787', 'Hanoi', 'Nha Trang', 'Noi Bai', 'Cam Ranh Airport', '2025-06-09 16:00:00', '2025-06-09 18:30:00', 1200000, 44, 6, 'On Time'),
('VN654', 'Vietnam Airlines', 'Airbus A321', 'Can Tho', 'Hanoi', 'Can Tho Airport', 'Noi Bai', '2025-06-10 06:00:00', '2025-06-10 08:30:00', 1300000, 44, 6, 'On Time'),
('VN876', 'Vietnam Airlines', 'Boeing 737', 'Hanoi', 'Can Tho', 'Noi Bai', 'Can Tho Airport', '2025-06-11 09:00:00', '2025-06-11 11:30:00', 1250000, 44, 6, 'On Time'),
('VN543', 'Vietnam Airlines', 'Airbus A330', 'Hue', 'Ho Chi Minh City', 'Phu Bai Airport', 'Tan Son Nhat', '2025-06-12 14:00:00', '2025-06-12 16:00:00', 950000, 44, 6, 'On Time'),
('VN765', 'Vietnam Airlines', 'Boeing 777', 'Ho Chi Minh City', 'Hue', 'Tan Son Nhat', 'Phu Bai Airport', '2025-06-13 17:00:00', '2025-06-13 19:00:00', 1000000, 44, 6, 'On Time'),
('VN987', 'Vietnam Airlines', 'Airbus A350', 'Hai Phong', 'Hanoi', 'Cat Bi Airport', 'Noi Bai', '2025-06-14 08:00:00', '2025-06-14 08:45:00', 500000, 44, 6, 'On Time'),
('VN210', 'Vietnam Airlines', 'Airbus A320', 'Hanoi', 'Hai Phong', 'Noi Bai', 'Cat Bi Airport', '2025-06-15 10:00:00', '2025-06-15 10:45:00', 500000, 44, 6, 'On Time'),
('VN211', 'Vietnam Airlines', 'Airbus A320', 'Ho Chi Minh City', 'Hanoi', 'Tan Son Nhat', 'Noi Bai', '2025-06-06 17:00:00', '2025-06-07 19:00:00', 1200000, 44, 6, 'On Time'),
('VN212', 'Vietnam Airlines', 'Boeing 787', 'SGN', 'HAN', 'Tan Son Nhat', 'Noi Bai', '2025-06-07 17:00:00', '2025-06-08 19:30:00', 1500000, 44, 6, 'On Time'),
('VN213', 'Vietnam Airlines', 'Boeing 787', 'HAN', 'SGN', 'Noi Bai', 'Tan Son Nhat', '2025-06-07 17:00:00', '2025-06-08 19:30:00', 1500000, 44, 6, 'On Time'),
('VN214', 'Vietnam Airlines', 'Airbus A320', 'HAN', 'SGN', 'Noi Bai', 'Tan Son Nhat', '2025-06-06 17:00:00', '2025-06-07 19:00:00', 1200000, 44, 6, 'On Time');

INSERT INTO Bookings (user_email, trip_type, departure_flight_id, return_flight_id, status)
VALUES
('john.doe@example.com', 'oneWay', 1, NULL, 'confirmed'),
('alice.wonderland@example.com', 'roundTrip', 3, 4, 'confirmed'),
('alice.wonderland@example.com', 'roundTrip', 2, 1, 'pending'),
('jane.smith@example.com', 'oneWay', 3, NULL, 'confirmed'),
('phungducanh2511@gmail.com', 'roundTrip', 4, 5, 'pending'),
('john.doe@example.com', 'roundTrip', 6, 7, 'confirmed'),
('alice.wonderland@example.com', 'oneWay', 8, NULL, 'cancelled'),
('jane.smith@example.com', 'roundTrip', 9, 10, 'pending'),
('phungducanh2511@gmail.com', 'oneWay', 11, NULL, 'confirmed'),
('john.doe@example.com', 'roundTrip', 12, 13, 'pending'),
('alice.wonderland@example.com', 'oneWay', 14, NULL, 'confirmed');

INSERT INTO Seats (flight_id, seat_code, is_available, class)
VALUES
(1, '1A', TRUE, 'economy'),
(1, '1B', FALSE, 'economy'),
(1, '1C', TRUE, 'economy'),
(1, '1D', TRUE, 'economy'),
(1, '1E', TRUE, 'economy'),
(1, '1F', TRUE, 'economy'),
(2, '2A', TRUE, 'business'),
(2, '2B', TRUE, 'business'),
(2, '2C', FALSE, 'business'),
(2, '2D', TRUE, 'business'),
(2, '2E', TRUE, 'business'),
(2, '2F', TRUE, 'business'),
(3, '3A', TRUE, 'economy'),
(3, '3B', TRUE, 'economy'),
(3, '3C', TRUE, 'economy'),
(3, '3D', FALSE, 'economy'),
(3, '3E', TRUE, 'economy'),
(3, '3F', TRUE, 'economy'),
(4, '4A', TRUE, 'economy'),
(4, '4B', TRUE, 'economy'),
(4, '4C', TRUE, 'economy'),
(4, '4D', TRUE, 'economy'),
(4, '4E', FALSE, 'economy'),
(4, '4F', TRUE, 'economy'),
(5, '5A', TRUE, 'business'),
(5, '5B', TRUE, 'business'),
(5, '5C', TRUE, 'business'),
(5, '5D', TRUE, 'business'),
(5, '5E', TRUE, 'business'),
(5, '5F', FALSE, 'business'),
(6, '6A', TRUE, 'economy'),
(6, '6B', TRUE, 'economy'),
(6, '6C', TRUE, 'economy'),
(6, '6D', TRUE, 'economy'),
(6, '6E', TRUE, 'economy'),
(6, '6F', TRUE, 'economy'),
(7, '7A', TRUE, 'economy'),
(7, '7B', TRUE, 'economy'),
(7, '7C', TRUE, 'economy'),
(7, '7D', TRUE, 'economy'),
(7, '7E', TRUE, 'economy'),
(7, '7F', TRUE, 'economy'),
(8, '8A', TRUE, 'economy'),
(8, '8B', TRUE, 'economy'),
(8, '8C', TRUE, 'economy'),
(8, '8D', TRUE, 'economy'),
(8, '8E', TRUE, 'economy'),
(8, '8F', TRUE, 'economy');

INSERT INTO Tickets (seat_id, flight_class, price, status, booking_id, flight_id)
VALUES
(1, 'economy', 1000000, 'booked', 1, 1),
(2, 'economy', 1000000, 'cancelled', 1, 1),
(3, 'business', 2000000, 'booked', 2, 2),
(4, 'economy', 1200000, 'booked', 3, 3),
(5, 'economy', 1200000, 'cancelled', 3, 3),
(6, 'economy', 700000, 'booked', 4, 4),
(7, 'business', 900000, 'booked', 4, 5),
(8, 'economy', 850000, 'booked', 5, 6),
(9, 'economy', 600000, 'booked', 5, 7),
(10, 'economy', 1100000, 'booked', 6, 8),
(11, 'economy', 1200000, 'booked', 7, 9),
(12, 'economy', 1300000, 'booked', 7, 10),
(13, 'economy', 1250000, 'booked', 8, 11),
(14, 'economy', 950000, 'booked', 9, 12),
(15, 'economy', 1000000, 'booked', 9, 13),
(16, 'economy', 500000, 'booked', 10, 14),
(17, 'economy', 500000, 'booked', 10, 15);

INSERT INTO TicketOwnerSnapshot (ticket_id, first_name, last_name, phone_number, gender)
VALUES
(1, 'John', 'Doe', '1234567890', 'Male'),
(3, 'Alice', 'Wonderland', '0987654321', 'Female'),
(4, 'Jane', 'Smith', NULL, 'Female'),
(6, 'Phung', 'Duc Anh', NULL, 'Male'),
(8, 'John', 'Doe', NULL, 'Male'),
(10, 'Alice', 'Wonderland', NULL, 'Female'),
(13, 'Phung', 'Duc Anh', NULL, 'Male'),
(16, 'Alice', 'Wonderland', NULL, 'Female');

INSERT INTO News (title, description, content, image, author_id)
VALUES
('Flight Delays Expected', 'Delays due to weather conditions', 'Full content about delays...', 'https://picsum.photos/536/354', 2),
('New Routes Announced', 'Vietnam Airlines announces new routes', 'Full content about new routes...', 'https://picsum.photos/536/354', 2),
('Special Discounts for Summer', 'Enjoy up to 50% off on selected flights', 'Full content about discounts...', 'https://picsum.photos/536/354', 2),
('New Aircraft Added to Fleet', 'Vietnam Airlines adds new aircraft to its fleet', 'Full content about new aircraft...', 'https://picsum.photos/536/354', 2),
('Safety Measures Enhanced', 'Enhanced safety measures for passengers', 'Full content about safety measures...', 'https://picsum.photos/536/354', 3),
('Loyalty Program Updates', 'Exciting updates to our loyalty program', 'Full content about loyalty program...', 'https://picsum.photos/536/354', 3),
('Holiday Travel Tips', 'Tips for a smooth holiday travel experience', 'Full content about travel tips...', 'https://picsum.photos/536/354', 4),
('New International Destinations', 'Explore new international destinations with us', 'Full content about destinations...', 'https://picsum.photos/536/354', 4);