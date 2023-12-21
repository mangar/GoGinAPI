package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Log struct {
	Id string `json:"id"`
	Mensagem string `json:"mensagem"`
}


func GetAllLogs(db *sql.DB) ([]ContratoEquipamento, error) {
    var contratos []ContratoEquipamento

    rows, err := db.Query("SELECT id, uuid, mensagem FROM Notificacoes ORDER BY id")
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