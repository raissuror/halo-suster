CREATE TABLE IF NOT EXISTS patients(
    id serial PRIMARY KEY,
    identity_number bigint NOT NULL,
    phone_number VARCHAR (300),
    name VARCHAR (30) NOT NULL,
    birthdate TIMESTAMP NOT NULL,
    gender VARCHAR (6) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);