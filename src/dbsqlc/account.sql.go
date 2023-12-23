// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: account.sql

package dbsqlc

import (
	"context"
)

const createAccount = `-- name: CreateAccount :exec
INSERT INTO Accounts ( name, email )
VALUES ($1, $2)
`

func (q *Queries) CreateAccount(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, createAccount)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT id, name, email FROM Accounts
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount)
	var i Account
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return i, err
}

const getAccountByEmail = `-- name: GetAccountByEmail :one
SELECT id, name, email FROM Accounts
WHERE email = $1
LIMIT 1
`

func (q *Queries) GetAccountByEmail(ctx context.Context) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountByEmail)
	var i Account
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return i, err
}

const getAccounts = `-- name: GetAccounts :many
SELECT id, name, email FROM Accounts
`

func (q *Queries) GetAccounts(ctx context.Context) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, getAccounts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Account
	for rows.Next() {
		var i Account
		if err := rows.Scan(&i.ID, &i.Name, &i.Email); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}