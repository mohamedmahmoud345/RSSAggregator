-- +goose Up

CREATE TABLE Users (
    Id UUID PRIMARY KEY,
    Created_At TIMESTAMP NOT NULL,
    Updated_At TIMESTAMP NOT NULL,
    Name TEXT NOT NULL
);

-- +goose Down

DROP TABLE Users;

