CREATE TABLE IF NOT EXISTS records(
    id serial PRIMARY KEY,
    created_by_user_id INT NOT NULL,
    patient_id VARCHAR (300),
    symptoms VARCHAR (2000) NOT NULL,
    medications VARCHAR (2000) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_user_id
        FOREIGN KEY(created_by_user_id)
            REFERENCES users(id)
            ON DELETE CASCADE
);