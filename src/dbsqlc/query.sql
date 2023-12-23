-- name: GetAllLogs :many
SELECT * FROM Logs;

-- name: GetLastLog :one
SELECT * FROM Logs
WHERE id = (SELECT MAX(id) FROM Logs); 

-- name: InsertLog :exec
INSERT INTO Logs ( mensagem )
VALUES ($1);