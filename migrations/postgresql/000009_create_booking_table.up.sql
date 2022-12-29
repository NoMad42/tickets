CREATE TABLE booking (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    flight_id uuid REFERENCES flights,
    seats_id uuid REFERENCES seats,
    transaction_id uuid REFERENCES transactions,
    user_profiles_id uuid REFERENCES user_profiles,
    status booking_status DEFAULT 'free',
    UNIQUE (user_profiles_id, seats_id)
);
