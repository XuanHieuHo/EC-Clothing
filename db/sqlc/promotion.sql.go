// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: promotion.sql

package db

import (
	"context"
	"time"
)

const createPromotion = `-- name: CreatePromotion :one
INSERT INTO promotions (
  title,
  description,
  discount_percent,
  start_date,
  end_date
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING id, title, description, discount_percent, start_date, end_date
`

type CreatePromotionParams struct {
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	DiscountPercent float64   `json:"discount_percent"`
	StartDate       time.Time `json:"start_date"`
	EndDate         time.Time `json:"end_date"`
}

func (q *Queries) CreatePromotion(ctx context.Context, arg CreatePromotionParams) (Promotion, error) {
	row := q.db.QueryRowContext(ctx, createPromotion,
		arg.Title,
		arg.Description,
		arg.DiscountPercent,
		arg.StartDate,
		arg.EndDate,
	)
	var i Promotion
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.DiscountPercent,
		&i.StartDate,
		&i.EndDate,
	)
	return i, err
}

const deletePromotion = `-- name: DeletePromotion :exec
DELETE FROM promotions WHERE id = $1
`

func (q *Queries) DeletePromotion(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletePromotion, id)
	return err
}

const getPromotion = `-- name: GetPromotion :one
SELECT id, title, description, discount_percent, start_date, end_date FROM promotions
WHERE title = $1 LIMIT 1
`

func (q *Queries) GetPromotion(ctx context.Context, title string) (Promotion, error) {
	row := q.db.QueryRowContext(ctx, getPromotion, title)
	var i Promotion
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.DiscountPercent,
		&i.StartDate,
		&i.EndDate,
	)
	return i, err
}

const listPromotions = `-- name: ListPromotions :many
SELECT id, title, description, discount_percent, start_date, end_date FROM promotions
ORDER BY title
LIMIT $1
OFFSET $2
`

type ListPromotionsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListPromotions(ctx context.Context, arg ListPromotionsParams) ([]Promotion, error) {
	rows, err := q.db.QueryContext(ctx, listPromotions, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Promotion{}
	for rows.Next() {
		var i Promotion
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.DiscountPercent,
			&i.StartDate,
			&i.EndDate,
		); err != nil {
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

const updatePromotion = `-- name: UpdatePromotion :one
UPDATE promotions
SET description = $2, discount_percent = $3, end_date = $4
WHERE id = $1
RETURNING id, title, description, discount_percent, start_date, end_date
`

type UpdatePromotionParams struct {
	ID              int64     `json:"id"`
	Description     string    `json:"description"`
	DiscountPercent float64   `json:"discount_percent"`
	EndDate         time.Time `json:"end_date"`
}

func (q *Queries) UpdatePromotion(ctx context.Context, arg UpdatePromotionParams) (Promotion, error) {
	row := q.db.QueryRowContext(ctx, updatePromotion,
		arg.ID,
		arg.Description,
		arg.DiscountPercent,
		arg.EndDate,
	)
	var i Promotion
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.DiscountPercent,
		&i.StartDate,
		&i.EndDate,
	)
	return i, err
}
