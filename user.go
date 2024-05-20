package main

import (
	"database/sql"
	"fmt"
	"time"
)

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	IsParent  bool   `json:"is_parent"`
	Gender    bool   `json:"gender"`
	CreatedAt int    `json:"created_at"`
}

type UserCreate struct {
	Name     string `json:"name"`
	IsParent bool   `json:"is_parent"`
	Gender   bool   `json:"gender"`
}

var getAllUsers = getAllItemsFactory[User]("user", func(item *User, rows *sql.Rows) error {
	return rows.Scan(&item.ID, &item.Name, &item.IsParent, &item.Gender, &item.CreatedAt)
})

var getUserByID = getItemByIDFactory[User]("user", func(item *User, row *sql.Row) error {
	return row.Scan(&item.ID, &item.Name, &item.IsParent, &item.Gender, item.CreatedAt)
})

var createUser = createItemFactory[User, UserCreate](
	"user",
	func(item *UserCreate) string {
		return fmt.Sprintf("INSERT INTO user (name, is_parent, gender, created_at) VALUES ('%s', %t, %t, %d)",
			item.Name, item.IsParent, item.Gender, time.Now().Unix())
	},
	func(item *User, row *sql.Row) error {
		return row.Scan(&item.ID, &item.Name, &item.IsParent, &item.Gender, &item.CreatedAt)
	},
)

var deleteUserByID = deleteItemByIDFactory("user")
