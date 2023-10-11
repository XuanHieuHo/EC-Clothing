// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: items_order.sql

package db

import (
	"context"
)

const createItemsOrder = `-- name: CreateItemsOrder :one
INSERT INTO items_order (
  booking_id,
  product_id,
  quantity,
  price,
  size
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING id, booking_id, product_id, quantity, price, size
`

type CreateItemsOrderParams struct {
	BookingID string  `json:"booking_id"`
	ProductID int64   `json:"product_id"`
	Quantity  int32   `json:"quantity"`
	Price     float64 `json:"price"`
	Size      string  `json:"size"`
}

func (q *Queries) CreateItemsOrder(ctx context.Context, arg CreateItemsOrderParams) (ItemsOrder, error) {
	row := q.db.QueryRowContext(ctx, createItemsOrder,
		arg.BookingID,
		arg.ProductID,
		arg.Quantity,
		arg.Price,
		arg.Size,
	)
	var i ItemsOrder
	err := row.Scan(
		&i.ID,
		&i.BookingID,
		&i.ProductID,
		&i.Quantity,
		&i.Price,
		&i.Size,
	)
	return i, err
}

const deleteItemsOrder = `-- name: DeleteItemsOrder :exec
DELETE FROM items_order WHERE id = $1
`

func (q *Queries) DeleteItemsOrder(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteItemsOrder, id)
	return err
}

const deleteItemsOrderByBookingID = `-- name: DeleteItemsOrderByBookingID :exec
DELETE FROM items_order WHERE booking_id = $1
`

func (q *Queries) DeleteItemsOrderByBookingID(ctx context.Context, bookingID string) error {
	_, err := q.db.ExecContext(ctx, deleteItemsOrderByBookingID, bookingID)
	return err
}

const getItemsOrder = `-- name: GetItemsOrder :one
SELECT id, booking_id, product_id, quantity, price, size FROM items_order
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetItemsOrder(ctx context.Context, id int64) (ItemsOrder, error) {
	row := q.db.QueryRowContext(ctx, getItemsOrder, id)
	var i ItemsOrder
	err := row.Scan(
		&i.ID,
		&i.BookingID,
		&i.ProductID,
		&i.Quantity,
		&i.Price,
		&i.Size,
	)
	return i, err
}

const listItemsOrderByBookingID = `-- name: ListItemsOrderByBookingID :many
SELECT id, booking_id, product_id, quantity, price, size FROM items_order
WHERE booking_id = $1
ORDER BY id
`

func (q *Queries) ListItemsOrderByBookingID(ctx context.Context, bookingID string) ([]ItemsOrder, error) {
	rows, err := q.db.QueryContext(ctx, listItemsOrderByBookingID, bookingID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ItemsOrder{}
	for rows.Next() {
		var i ItemsOrder
		if err := rows.Scan(
			&i.ID,
			&i.BookingID,
			&i.ProductID,
			&i.Quantity,
			&i.Price,
			&i.Size,
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
