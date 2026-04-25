-- name: CreateUser :one
INSERT INTO Users (Id, Created_At, Updated_At, Name)
VALUES ($1, $2, $3, $4)
RETURNING *;
