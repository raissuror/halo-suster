CREATE TABLE IF NOT EXISTS users(
    id serial PRIMARY KEY,
    name VARCHAR (50) NOT NULL,
    password VARCHAR (300),
    identify_card_scan_img VARCHAR (300) NOT NULL,
    role VARCHAR (16) NOT NULL,
    status VARCHAR (16) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);