-- +goose Up
CREATE TABLE USERS (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL, 
    password TEXT NOT NULL,
    email TEXT UNIQUE not Null, 
    createdAt TIMESTAMP NOT NULL, 
    updatedAt TIMESTAMP NOT NULL
);
-- +goose Down
Drop table USERS;
