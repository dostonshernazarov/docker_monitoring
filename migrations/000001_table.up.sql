CREATE TABLE container_status (
    id SERIAL PRIMARY KEY,
    ip_address TEXT NOT NULL,
    ping_time INT NOT NULL,
    last_success TIMESTAMP NOT NULL
);