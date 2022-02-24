package models

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) CheckUserWithNumber(employee string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select * from member where employee_id = $1 
	`
	row := m.DB.QueryRowContext(ctx, query, employee)

	var user User

	err := row.Scan(
		&user.ID,
		&user.EmployeeID,
		&user.Username,
	)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return nil, err
	}

	return &user, nil

}

// Get returns one movie and error, if any
func (m *DBModel) Get(id int) (*Menu, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select * from menu where id = $1 
	`

	row := m.DB.QueryRowContext(ctx, query, id)

	var menu Menu

	err := row.Scan(
		&menu.ID,
		&menu.Name,
		&menu.Type,
		&menu.Memo,
		&menu.FileString,
		&menu.CreatedAt,
		&menu.UpdatedAt,
		&menu.Opened,
	)
	if err != nil {
		return nil, err
	}

	return &menu, nil
}

func (m *DBModel) Create(newMenu Menu) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into menu (name, type, memo, image, created_at, updated_at, opened) 
		values ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := m.DB.ExecContext(ctx, stmt,
		newMenu.Name,
		newMenu.Type,
		newMenu.Memo,
		newMenu.FileString,
		newMenu.CreatedAt,
		newMenu.UpdatedAt,
		newMenu.Opened,
	)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (m *DBModel) AllMenu() ([]*Menu, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select * from menu
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var menus []*Menu
	for rows.Next() {
		var menu Menu
		err := rows.Scan(
			&menu.ID,
			&menu.Name,
			&menu.Type,
			&menu.Memo,
			&menu.FileString,
			&menu.CreatedAt,
			&menu.UpdatedAt,
			&menu.Opened,
		)
		if err != nil {
			return nil, err
		}
		menus = append(menus, &menu)
	}

	return menus, nil
}

func (m *DBModel) UpdateOpen(id int, name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var time = time.Now().Format("2006-01-02")
	stmt := `update menu set opened = 'true', updated_at = $1 where id = $2 and name = $3
	`
	_, err := m.DB.ExecContext(ctx, stmt, time, id, name)
	if err != nil {
		return err
	}

	return nil
}

func (m *DBModel) OpenedMenu() ([]*Menu, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var time = time.Now().Format("2006-01-02")
	log.Println(time)
	query := `select * from menu where opened = 'true' and updated_at = $1
	`
	rows, err := m.DB.QueryContext(ctx, query, time)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var menus []*Menu
	for rows.Next() {
		var menu Menu
		err := rows.Scan(
			&menu.ID,
			&menu.Name,
			&menu.Type,
			&menu.Memo,
			&menu.FileString,
			&menu.CreatedAt,
			&menu.UpdatedAt,
			&menu.Opened,
		)
		if err != nil {
			return nil, err
		}
		menus = append(menus, &menu)
	}

	return menus, nil
}

func (m *DBModel) AddOrder(order Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into orders (menu_id, name, type, item, sugar, ice, price, user_memo, updated_at, user_name, count)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`
	_, err := m.DB.ExecContext(ctx, stmt,
		order.MenuID,
		order.Name,
		order.Type,
		order.Item,
		order.Sugar,
		order.Ice,
		order.Price,
		order.UserMemo,
		order.UpdatedAt,
		order.User,
		order.Count,
	)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (m *DBModel) AllOrder() ([]*Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var time = time.Now().Format("2006-01-02")
	query := `select o.id, o.menu_id, o.name, o.type, o.item, o.sugar, o.ice, o.price, o.user_memo, o.updated_at, o.user_name, o.count from orders o
		INNER JOIN menu m
		ON m.name = o.name and m.opened = 'true'
		where o.updated_at = $1
	`

	rows, err := m.DB.QueryContext(ctx, query, time)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*Order
	for rows.Next() {
		var order Order
		err := rows.Scan(
			&order.ID,
			&order.MenuID,
			&order.Name,
			&order.Type,
			&order.Item,
			&order.Sugar,
			&order.Ice,
			&order.Price,
			&order.UserMemo,
			&order.UpdatedAt,
			&order.User,
			&order.Count,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}

	return orders, nil
}

func (m *DBModel) UpdateOrder(order Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `update orders set item = $3, sugar = $4, ice = $5, price = $6, user_memo = $7, updated_at = $8, count = $9 where id = $1 and menu_id = $2
	`
	_, err := m.DB.ExecContext(ctx, stmt,
		order.ID,
		order.MenuID,
		order.Item,
		order.Sugar,
		order.Ice,
		order.Price,
		order.UserMemo,
		order.UpdatedAt,
		order.Count,
	)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (m *DBModel) UpdateMenu(menu Menu) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `update menu set name = $2, type = $3, memo = $4, image = $5, updated_at = $6, opened=false where id = $1
	`
	_, err := m.DB.ExecContext(ctx, stmt,
		menu.ID,
		menu.Name,
		menu.Type,
		menu.Memo,
		menu.FileString,
		menu.UpdatedAt,
	)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (m *DBModel) DeleteOpenMenu(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `update menu set opened = false where id = $1
	`
	_, err := m.DB.ExecContext(ctx, stmt, id)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (m *DBModel) DeleteOrder(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `delete from orders where id = $1
	`
	_, err := m.DB.ExecContext(ctx, stmt, id)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
