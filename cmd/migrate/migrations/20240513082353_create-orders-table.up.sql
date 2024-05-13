CREATE TYPE status_enum AS ENUM ('pending','completed','cancelled');
CREATE TABLE IF NOT EXISTS ORDERS (
    id SERIAL PRIMARY KEY,
    userId  INTEGER NOT NULL,
    total DECIMAL(10,2) NOT NULL,
    status  status_enum not null DEFAULT 'pending',
    address TEXT not null,
    createdAt TIMESTAMP not null DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP not null DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_userid FOREIGN KEY(userId) REFERENCES USERS(id)
);