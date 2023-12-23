-- name: GetAllLogs :many
SELECT * FROM Logs;

-- name: GetLastLog :one
SELECT * FROM Logs
WHERE id = (SELECT MAX(id) FROM Logs); 

-- name: InsertLog :exec
INSERT INTO Logs ( mensagem )
VALUES ($1);


-- name: InsertAccount :exec
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
