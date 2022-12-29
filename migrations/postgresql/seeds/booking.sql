INSERT INTO booking (
    flight_id,
    seats_id,
    user_profiles_id,
    status
) SELECT 
    flight_id, 
    id AS seats_id,
    (SELECT id FROM user_profiles ORDER BY random() LIMIT 1) AS user_profiles_id,
    booking_status
    FROM seats 
    WHERE booking_status = 'booked';

INSERT INTO booking (
    flight_id,
    seats_id,
    user_profiles_id,
    status
) SELECT 
    flight_id, 
    id AS seats_id,
    (SELECT id FROM user_profiles ORDER BY random() LIMIT 1) AS user_profiles_id,
    booking_status
    FROM seats 
    WHERE booking_status = 'booking';
