CREATE TABLE seat_options (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name varchar(255) NOT NUll,
    description text,
    price numeric
);
