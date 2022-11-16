CREATE TABLE user_profiles (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    login varchar(255) NOT NUll,
    password varchar(255) NOT NUll,
    avatar varchar(255),
    name varchar(255),
    lastname varchar(255),
    personal_id varchar(255)
);
