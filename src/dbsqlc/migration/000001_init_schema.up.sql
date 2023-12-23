
CREATE TABLE IF NOT EXISTS Notificacoes
(
    id int auto_increment not null,
    timestamp datetime null,
    uuid varchar(255) null,
    mensagem longtext null,
    is_ok varchar(100) null,
    primary key (id)
);
