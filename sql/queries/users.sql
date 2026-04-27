-- name: CreateUser :one
INSERT INTO Users (Id, Created_At, Updated_At, Name, api_key)
VALUES ($1, $2, $3, $4,
    encode(sha256(random()::text::bytea), 'hex')
)
RETURNING *;

-- name: GetUserByApiKey :one
SELECT * FROM Users WHERE api_key = $1;


