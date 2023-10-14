// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: product.sql

package db

import (
	"context"
	"database/sql"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO products (
  product_name,
  thumb,
  price
) VALUES (
  $1, $2, $3
) RETURNING id, product_name, thumb, price
`

type CreateProductParams struct {
	ProductName string  `json:"product_name"`
	Thumb       string  `json:"thumb"`
	Price       float64 `json:"price"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, createProduct, arg.ProductName, arg.Thumb, arg.Price)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.ProductName,
		&i.Thumb,
		&i.Price,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM products WHERE id = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteProduct, id)
	return err
}

const findProduct = `-- name: FindProduct :many
SELECT products.id, products.product_name, products.thumb, products.price
FROM products
INNER JOIN descriptions_product
ON products.id = descriptions_product.product_id
WHERE (
  product_name ILIKE '%' || $1 || '%'
  OR descriptions_product.gender ILIKE '%' || $1 || '%'
  OR descriptions_product.material ILIKE '%' || $1 || '%'
)
LIMIT $2
OFFSET $3
`

type FindProductParams struct {
	Column1 sql.NullString `json:"column_1"`
	Limit   int32          `json:"limit"`
	Offset  int32          `json:"offset"`
}

func (q *Queries) FindProduct(ctx context.Context, arg FindProductParams) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, findProduct, arg.Column1, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.ProductName,
			&i.Thumb,
			&i.Price,
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

const getProduct = `-- name: GetProduct :one
SELECT id, product_name, thumb, price FROM products
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProduct(ctx context.Context, id int64) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProduct, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.ProductName,
		&i.Thumb,
		&i.Price,
	)
	return i, err
}

const listProducts = `-- name: ListProducts :many
SELECT id, product_name, thumb, price FROM products
ORDER BY product_name
LIMIT $1
OFFSET $2
`

type ListProductsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListProducts(ctx context.Context, arg ListProductsParams) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, listProducts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.ProductName,
			&i.Thumb,
			&i.Price,
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

const updateProduct = `-- name: UpdateProduct :one
UPDATE products
SET product_name = $2, thumb = $3, price = $4
WHERE id = $1
RETURNING id, product_name, thumb, price
`

type UpdateProductParams struct {
	ID          int64   `json:"id"`
	ProductName string  `json:"product_name"`
	Thumb       string  `json:"thumb"`
	Price       float64 `json:"price"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, updateProduct,
		arg.ID,
		arg.ProductName,
		arg.Thumb,
		arg.Price,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.ProductName,
		&i.Thumb,
		&i.Price,
	)
	return i, err
}
