CREATE TABLE transactions (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    amount numeric NOT NULL,
    created_at  timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
    user_profiles_id uuid REFERENCES user_profiles
);
