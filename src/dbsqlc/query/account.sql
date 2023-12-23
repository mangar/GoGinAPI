-- name: CreateAccount :execresult
INSERT INTO Accounts ( name, email ) VALUES (?, ?);

-- name: GetAccountByEmail :one
SELECT * FROM Accounts
WHERE email = ?
LIMIT 1;

-- name: GetAccount :one
SELECT * FROM Accounts
WHERE id = ?
LIMIT 1;

-- name: GetAccounts :many
SELECT * FROM Accounts;
