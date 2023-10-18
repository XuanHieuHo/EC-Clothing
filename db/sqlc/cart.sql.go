// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: cart.sql

package db

import (
	"context"
)

const createCart = `-- name: CreateCart :one
INSERT INTO carts (
  username,
  product_id,
  quantity,
  price,
  size
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING id, username, product_id, quantity, price, size
`

type CreateCartParams struct {
	Username  string  `json:"username"`
	ProductID int64   `json:"product_id"`
	Quantity  int32   `json:"quantity"`
	Price     float64 `json:"price"`
	Size      string  `json:"size"`
}

func (q *Queries) CreateCart(ctx context.Context, arg CreateCartParams) (Cart, error) {
	row := q.db.QueryRowContext(ctx, createCart,
		arg.Username,
		arg.ProductID,
		arg.Quantity,
		arg.Price,
		arg.Size,
	)
	var i Cart
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.ProductID,
		&i.Quantity,
		&i.Price,
		&i.Size,
	)
	return i, err
}

const deleteCart = `-- name: DeleteCart :exec
DELETE FROM carts WHERE id = $1
`

func (q *Queries) DeleteCart(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCart, id)
	return err
}

const deleteCartOfUser = `-- name: DeleteCartOfUser :exec
DELETE FROM carts WHERE username = $1
`

func (q *Queries) DeleteCartOfUser(ctx context.Context, username string) error {
	_, err := q.db.ExecContext(ctx, deleteCartOfUser, username)
	return err
}

const getCart = `-- name: GetCart :one
SELECT id, username, product_id, quantity, price, size FROM carts
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCart(ctx context.Context, id int64) (Cart, error) {
	row := q.db.QueryRowContext(ctx, getCart, id)
	var i Cart
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.ProductID,
		&i.Quantity,
		&i.Price,
		&i.Size,
	)
	return i, err
}

const getCartDetails = `-- name: GetCartDetails :one
SELECT id, username, product_id, quantity, price, size FROM carts
WHERE username = $1 AND product_id = $2 AND size = $3
LIMIT 1
`

type GetCartDetailsParams struct {
	Username  string `json:"username"`
	ProductID int64  `json:"product_id"`
	Size      string `json:"size"`
}

func (q *Queries) GetCartDetails(ctx context.Context, arg GetCartDetailsParams) (Cart, error) {
	row := q.db.QueryRowContext(ctx, getCartDetails, arg.Username, arg.ProductID, arg.Size)
	var i Cart
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.ProductID,
		&i.Quantity,
		&i.Price,
		&i.Size,
	)
	return i, err
}

const listCartOfUser = `-- name: ListCartOfUser :many
SELECT id, username, product_id, quantity, price, size FROM carts
WHERE username = $1
LIMIT $2
OFFSET $3
`

type ListCartOfUserParams struct {
	Username string `json:"username"`
	Limit    int32  `json:"limit"`
	Offset   int32  `json:"offset"`
}

func (q *Queries) ListCartOfUser(ctx context.Context, arg ListCartOfUserParams) ([]Cart, error) {
	rows, err := q.db.QueryContext(ctx, listCartOfUser, arg.Username, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Cart{}
	for rows.Next() {
		var i Cart
		if err := rows.Scan(
			&i.ID,
			&i.Username,
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

const updateCart = `-- name: UpdateCart :one
UPDATE carts
SET quantity = $2, size = $3, price = $4
WHERE id = $1
RETURNING id, username, product_id, quantity, price, size
`

type UpdateCartParams struct {
	ID       int64   `json:"id"`
	Quantity int32   `json:"quantity"`
	Size     string  `json:"size"`
	Price    float64 `json:"price"`
}

func (q *Queries) UpdateCart(ctx context.Context, arg UpdateCartParams) (Cart, error) {
	row := q.db.QueryRowContext(ctx, updateCart,
		arg.ID,
		arg.Quantity,
		arg.Size,
		arg.Price,
	)
	var i Cart
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.ProductID,
		&i.Quantity,
		&i.Price,
		&i.Size,
	)
	return i, err
}
