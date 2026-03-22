-- +goose Up
CREATE TABLE employee (
    id uuid PRIMARY KEY,
    Salary INT NOT NULL,
    role TEXT NOT NULL,
    userId UUID REFERENCES users (id) ON DELETE CASCADE,
    companyEmail TEXT UNIQUE not Null,
    createdAt TIME NOT NULL,
    updatedAt TIME NOT NULL
);
-- +goose Down
Drop table employee;