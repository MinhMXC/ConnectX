package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	IsParent  bool   `json:"is_parent"`
	Gender    bool   `json:"gender"`
	CreatedAt int64  `json:"created_at"`
}

type UserUpdate struct {
	IsParent bool `json:"is_parent"`
	Gender   bool `json:"gender"`
}

func scanUserRows(item *User, rows *sql.Rows) error {
	var temp string
	return rows.Scan(&item.ID, &item.IsParent, &item.Gender, &temp, &item.Username, &temp, &temp, &item.CreatedAt)
}

func scanUserRow(item *User, row *sql.Row) error {
	var temp string
	return row.Scan(&item.ID, &item.IsParent, &item.Gender, &temp, &item.Username, &temp, &temp, &item.CreatedAt)
}

func userSetup(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	data, status, err := verifyJson[UserUpdate](w, r)
	if status != 200 {
		return status, err
	}

	user, err := authorize(w, r, false)
	if err != nil {
		return http.StatusUnauthorized, err
	}

	_, err = context.db.Exec(fmt.Sprintf("INSERT INTO user (user_id, is_parent, gender) VALUES (%d, %t, %t)",
		user.ID, data.IsParent, data.Gender))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	_, err = context.db.Exec(fmt.Sprintf("UPDATE base_user SET user_type = 1 WHERE id=%d", user.ID))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

var getAllUsers = getAllItemsFactory[User]("user INNER JOIN base_user ON base_user.id = user.user_id", scanUserRows)

var getUserByID = getItemByIDFactory[User]("user INNER JOIN base_user ON base_user.id = user.user_id", scanUserRow)

var updateUser = updateItemFactory[User, UserUpdate](
	"user INNER JOIN base_user ON base_user.id = user.user_id",
	func(item *UserUpdate) string {
		return fmt.Sprintf("UPDATE user INNER JOIN base_user ON base_user.id = user.user_id SET is_parent = %t, gender = %t",
			item.IsParent, item.Gender)
	},
	scanUserRow,
)
