// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package dbsqlc

import (
	"database/sql"
)

type Log struct {
	ID       int32
	Mensagem sql.NullString
}