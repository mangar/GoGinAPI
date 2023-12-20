package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// initDB inicializa e retorna uma conexão com o banco de dados.
func InitDB() *sql.DB {
	var mysql_server string = "root:password@tcp(192.168.1.110:3306)/goginapi"
    db, err := sql.Open("mysql", mysql_server)
    if err != nil {
        log.Fatalf("Erro ao abrir conexão com banco de dados: %v", err)
    }

    // Verifica se a conexão com o banco está funcionando.
    if err := db.Ping(); err != nil {
        log.Fatalf("Erro ao conectar com o banco de dados: %v", err)
    }

    return db
}