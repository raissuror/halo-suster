CREATE TABLE IF NOT EXISTS users(
    id serial PRIMARY KEY,
    nip VARCHAR (50) NOT NULL,
    name VARCHAR (50) NOT NULL,
    password VARCHAR (300) NOT NULL,
    identify_card_scan_img VARCHAR (300),
    role VARCHAR (16),
    status VARCHAR (16),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);