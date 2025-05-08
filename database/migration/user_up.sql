start transaction;

CREATE TABLE IF NOT EXISTS users (
    id BIGINT PRIMARY KEY,
    name VARCHAR(255),
    age INTEGER,
    active BOOLEAN
);

commit transaction;