// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: role.sql

package db

import (
	"context"
)

const getRole = `-- name: GetRole :one
SELECT id, name FROM roles
WHERE name = $1 LIMIT 1
`

func (q *Queries) GetRole(ctx context.Context, name string) (Role, error) {
	row := q.db.QueryRowContext(ctx, getRole, name)
	var i Role
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}
