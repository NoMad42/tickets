CREATE TABLE seats_seat_options (
    seats_id uuid REFERENCES seats,
    seat_options_id uuid REFERENCES seat_options,
    PRIMARY KEY (seats_id, seat_options_id)
)