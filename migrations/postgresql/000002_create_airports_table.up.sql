CREATE TABLE airports (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    code varchar(255) NOT NUll,
    name varchar(255) NOT NUll,
    description text,
    country varchar(255) NOT NUll,
    city varchar(255) NOT NUll
);
