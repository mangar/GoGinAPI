CREATE PROCEDURE InsertRecords()
BEGIN
    DECLARE i INT DEFAULT 1;
    WHILE i <= 100 DO
        INSERT INTO Notificacoes (timestamp, uuid, mensagem, is_ok)
        VALUES (NOW(), UUID(), CONCAT('Valor ', i), TRUE);
        SET i = i + 1;
    END WHILE;
END;

CALL InsertRecords();

DROP PROCEDURE InsertRecords;