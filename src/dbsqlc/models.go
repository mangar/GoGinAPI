// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package dbsqlc

import (
	"database/sql"
)

type Account struct {
	ID    int32
	Name  sql.NullString
	Email sql.NullString
}

type Log struct {
	ID       int32
	Mensagem sql.NullString
}

type Notificaco struct {
	ID        int32
	Timestamp sql.NullTime
	Uuid      sql.NullString
	Mensagem  sql.NullString
	IsOk      sql.NullString
}
