-- name: CreateAccount :execresult
INSERT INTO Accounts ( name, email )
VALUES ($1, $2);

-- name: GetAccountByEmail :one
SELECT * FROM Accounts
WHERE email = $1
LIMIT 1;

-- name: GetAccount :one
SELECT * FROM Accounts
WHERE id = $1
LIMIT 1;

-- name: GetAccounts :many
SELECT * FROM Accounts;
