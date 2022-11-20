CREATE TABLE flights (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    code varchar(255) NOT NUll,
    from_airport_id uuid REFERENCES airports,
    from_timestamp timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
    to_airport_id uuid  REFERENCES airports,
    to_timestamp timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
    status varchar(255) NOT NULL
);
