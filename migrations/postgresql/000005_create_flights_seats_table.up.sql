CREATE TABLE flights_seats (
    flights_id uuid REFERENCES flights,
    seats_id uuid REFERENCES seats,
    PRIMARY KEY (flights_id, seats_id)
);
