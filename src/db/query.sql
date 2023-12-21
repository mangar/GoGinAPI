-- name: GetAllNotificacoes :many
SELECT * FROM Notificacoes

-- name: GetLastNotificacao :one
SELECT * FROM Notificacoes
WHERE id = (SELECT MAX(id) FROM Notificacoes); 