package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type ContratoEquipamento struct {
	Id string `json:"id"`
	UUID string `json:"uuid"`
	Mensagem string `json:"mensagem"`
}


func QueryNotificacoes(db *sql.DB) {
    rows, err := db.Query("SELECT id, uuid, mensagem FROM Notificacoes")
    if err != nil {
		fmt.Println(err)
    }
    defer rows.Close()

    for rows.Next() {
        // Supondo que a tabela tenha colunas como id, name, etc.
        var id int
		var uuid string
        var mensagem string
        err = rows.Scan(&id, &uuid, &mensagem)
        if err != nil {
            fmt.Println(err)
        }
        fmt.Println(id, uuid, mensagem)
    }

    // Verifique se houve erros durante a iteração
    if err = rows.Err(); err != nil {
        fmt.Println(err)
    }
}