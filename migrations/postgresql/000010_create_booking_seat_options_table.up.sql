CREATE TABLE booking_seat_options (
    booking_id uuid REFERENCES booking,
    seat_options_id uuid REFERENCES seat_options,
    PRIMARY KEY (booking_id, seat_options_id)
);
