// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: province.sql

package db

import (
	"context"
)

const createProvince = `-- name: CreateProvince :one
INSERT INTO provinces (
  name
) VALUES (
  $1
) RETURNING id, name
`

func (q *Queries) CreateProvince(ctx context.Context, name string) (Province, error) {
	row := q.db.QueryRowContext(ctx, createProvince, name)
	var i Province
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getProvince = `-- name: GetProvince :one
SELECT id, name FROM provinces
WHERE name = $1 LIMIT 1
`

func (q *Queries) GetProvince(ctx context.Context, name string) (Province, error) {
	row := q.db.QueryRowContext(ctx, getProvince, name)
	var i Province
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getProvinceByID = `-- name: GetProvinceByID :one
SELECT id, name FROM provinces
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProvinceByID(ctx context.Context, id int64) (Province, error) {
	row := q.db.QueryRowContext(ctx, getProvinceByID, id)
	var i Province
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const listProvinces = `-- name: ListProvinces :many
SELECT name FROM provinces
ORDER BY name
`

func (q *Queries) ListProvinces(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, listProvinces)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
