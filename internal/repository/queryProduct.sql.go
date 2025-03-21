// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: queryProduct.sql

package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const ProductCreate = `-- name: ProductCreate :one
insert into shop.Products(name, price)
values ($1, $2) returning id
`

type ProductCreateParams struct {
	Name  string         `db:"name" json:"name"`
	Price pgtype.Numeric `db:"price" json:"price"`
}

func (q *Queries) ProductCreate(ctx context.Context, arg ProductCreateParams) (int, error) {
	row := q.db.QueryRow(ctx, ProductCreate, arg.Name, arg.Price)
	var id int
	err := row.Scan(&id)
	return id, err
}

const ProductDelete = `-- name: ProductDelete :exec
delete from shop.products 
where name = $1
`

func (q *Queries) ProductDelete(ctx context.Context, name string) error {
	_, err := q.db.Exec(ctx, ProductDelete, name)
	return err
}

const ProductGetById = `-- name: ProductGetById :one
select id, name, price from shop.products p 
where p.id = $1
`

func (q *Queries) ProductGetById(ctx context.Context, id int) (*ShopProduct, error) {
	row := q.db.QueryRow(ctx, ProductGetById, id)
	var i ShopProduct
	err := row.Scan(&i.ID, &i.Name, &i.Price)
	return &i, err
}

const ProductGetByName = `-- name: ProductGetByName :one
select id, name, price from shop.products p 
where p.name = $1
`

func (q *Queries) ProductGetByName(ctx context.Context, name string) (*ShopProduct, error) {
	row := q.db.QueryRow(ctx, ProductGetByName, name)
	var i ShopProduct
	err := row.Scan(&i.ID, &i.Name, &i.Price)
	return &i, err
}

const ProductGetRangePrice = `-- name: ProductGetRangePrice :many
select id, name, price from shop.products p 
where price between $1 and $2
order by price
`

type ProductGetRangePriceParams struct {
	Price   pgtype.Numeric `db:"price" json:"price"`
	Price_2 pgtype.Numeric `db:"price_2" json:"price_2"`
}

func (q *Queries) ProductGetRangePrice(ctx context.Context, arg ProductGetRangePriceParams) ([]*ShopProduct, error) {
	rows, err := q.db.Query(ctx, ProductGetRangePrice, arg.Price, arg.Price_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*ShopProduct{}
	for rows.Next() {
		var i ShopProduct
		if err := rows.Scan(&i.ID, &i.Name, &i.Price); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const ProductID = `-- name: ProductID :one
select id from shop.products p
where p.name = $1
`

func (q *Queries) ProductID(ctx context.Context, name string) (int, error) {
	row := q.db.QueryRow(ctx, ProductID, name)
	var id int
	err := row.Scan(&id)
	return id, err
}

const ProductUpdate = `-- name: ProductUpdate :exec
update shop.products
set price = $1
where name = $2
`

type ProductUpdateParams struct {
	Price pgtype.Numeric `db:"price" json:"price"`
	Name  string         `db:"name" json:"name"`
}

func (q *Queries) ProductUpdate(ctx context.Context, arg ProductUpdateParams) error {
	_, err := q.db.Exec(ctx, ProductUpdate, arg.Price, arg.Name)
	return err
}

const Products = `-- name: Products :many
select id, name, price from shop.products p
`

func (q *Queries) Products(ctx context.Context) ([]*ShopProduct, error) {
	rows, err := q.db.Query(ctx, Products)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*ShopProduct{}
	for rows.Next() {
		var i ShopProduct
		if err := rows.Scan(&i.ID, &i.Name, &i.Price); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
