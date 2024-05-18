package main

import "database/sql"

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	IsParent  bool   `json:"is_parent"`
	Gender    string `json:"gender"`
	CreatedAt int    `json:"created_at"`
}

var getAllUsers = getAllItemsFactory[User]("user", func(item *User, rows *sql.Rows) error {
	return rows.Scan(&item.ID, &item.Name, &item.IsParent, &item.Gender, &item.CreatedAt)
})

var getUserByID = getItemByIDFactory[User]("user", func(item *User, row *sql.Row) error {
	return row.Scan(&item.ID, &item.Name, &item.IsParent, &item.Gender, item.CreatedAt)
})

var deleteUserByID = deleteItemByIDFactory("user")
