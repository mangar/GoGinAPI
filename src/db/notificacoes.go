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


func QueryNotificacoes(db *sql.DB) ([]ContratoEquipamento, error) {
    var contratos []ContratoEquipamento

    rows, err := db.Query("SELECT id, uuid, mensagem FROM Notificacoes")
    if err != nil {
		fmt.Println(err)
    }
    defer rows.Close()

    for rows.Next() {
        var contrato ContratoEquipamento        
        err = rows.Scan(&contrato.Id, &contrato.UUID, &contrato.Mensagem)
        if err != nil {
            fmt.Println(err)
        }

        fmt.Println(&contrato)
        contratos = append(contratos, contrato)
    }

    // Verifique se houve erros durante a iteração
    if err = rows.Err(); err != nil {
        fmt.Println(err)
    }
    return contratos, nil
}