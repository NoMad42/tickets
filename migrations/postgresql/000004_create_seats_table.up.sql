CREATE TYPE booking_status AS ENUM ('free', 'booking', 'booked');
CREATE TYPE seet_position AS ENUM ('middle', 'aisle', 'window');
CREATE TYPE seat_type AS ENUM ('economy', 'buisness');

CREATE TABLE seats (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    code varchar(255) NOT NUll,
    postition seet_position DEFAULT 'middle',
    price numeric NOT NULL,
    booking_status booking_status DEFAULT 'free',
    seat_type seat_type DEFAULT 'economy'
);
