CREATE TABLE IF NOT EXISTS Logs
(
    id int auto_increment not null,
    mensagem longtext null,
    primary key (id)
);

CREATE TABLE IF NOT EXISTS Accounts
(
    id int auto_increment not null,
    name varchar(255) null,
    email varchar(255) null,
    primary key (id)
);
