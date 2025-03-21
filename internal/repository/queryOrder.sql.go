// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: queryOrder.sql

package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const OrderByID = `-- name: OrderByID :one
select id, user_id, order_date, total_amount from shop.orders o
where o.id = $1
`

func (q *Queries) OrderByID(ctx context.Context, id int) (*ShopOrder, error) {
	row := q.db.QueryRow(ctx, OrderByID, id)
	var i ShopOrder
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.OrderDate,
		&i.TotalAmount,
	)
	return &i, err
}

const OrderCreateByUser = `-- name: OrderCreateByUser :one
insert into shop.Orders (user_id)
values ((select id from shop.Users where name = $1))
returning id
`

func (q *Queries) OrderCreateByUser(ctx context.Context, name string) (int, error) {
	row := q.db.QueryRow(ctx, OrderCreateByUser, name)
	var id int
	err := row.Scan(&id)
	return id, err
}

const OrderDelete = `-- name: OrderDelete :exec
delete from shop.orders 
where id=$1
`

func (q *Queries) OrderDelete(ctx context.Context, id int) error {
	_, err := q.db.Exec(ctx, OrderDelete, id)
	return err
}

const OrderUpdateTotal = `-- name: OrderUpdateTotal :exec
update shop.orders ord
set total_amount=(
	select sum(p.price * op.quantity) from 
	shop.orderproducts op
	inner join shop.products p on op.product_id=p.id
	where op.order_id = $1
	group by op.order_id)
where ord.id = $1
`

func (q *Queries) OrderUpdateTotal(ctx context.Context, orderID int32) error {
	_, err := q.db.Exec(ctx, OrderUpdateTotal, orderID)
	return err
}

const OrderUser = `-- name: OrderUser :one
select  o.id, u.name, o.order_date, o.total_amount from shop.orders o 
inner join shop.users u on o.user_id = u.id
where o.id = $1
`

type OrderUserRow struct {
	ID          int                `db:"id" json:"id"`
	Name        string             `db:"name" json:"name"`
	OrderDate   pgtype.Timestamptz `db:"order_date" json:"order_date"`
	TotalAmount pgtype.Numeric     `db:"total_amount" json:"total_amount"`
}

func (q *Queries) OrderUser(ctx context.Context, id int) (*OrderUserRow, error) {
	row := q.db.QueryRow(ctx, OrderUser, id)
	var i OrderUserRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.OrderDate,
		&i.TotalAmount,
	)
	return &i, err
}

const Orders = `-- name: Orders :many
select  o.id, u.name, o.order_date, o.total_amount from shop.orders o 
inner join shop.users u on o.user_id = u.id
order by o.id, u.name
`

type OrdersRow struct {
	ID          int                `db:"id" json:"id"`
	Name        string             `db:"name" json:"name"`
	OrderDate   pgtype.Timestamptz `db:"order_date" json:"order_date"`
	TotalAmount pgtype.Numeric     `db:"total_amount" json:"total_amount"`
}

func (q *Queries) Orders(ctx context.Context) ([]*OrdersRow, error) {
	rows, err := q.db.Query(ctx, Orders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*OrdersRow{}
	for rows.Next() {
		var i OrdersRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.OrderDate,
			&i.TotalAmount,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
