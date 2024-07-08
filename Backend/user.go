package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Picture   string `json:"picture"`
	IsParent  bool   `json:"is_parent"`
	Gender    bool   `json:"gender"`
	CreatedAt int64  `json:"created_at"`
}

type UserCreate struct {
	Username string `json:"username"`
	Picture  string `json:"picture"`
	IsParent bool   `json:"is_parent"`
	Gender   bool   `json:"gender"`
}

func scanUserRows(item *User, rows *sql.Rows) error {
	var temp string
	return rows.Scan(&item.ID, &item.Username, &item.Picture, &item.IsParent, &item.Gender, &temp, &item.Email, &temp, &temp, &item.CreatedAt)
}

func scanUserRow(item *User, row *sql.Row) error {
	var temp string
	return row.Scan(&item.ID, &item.Username, &item.Picture, &item.IsParent, &item.Gender, &temp, &item.Email, &temp, &temp, &item.CreatedAt)
}

func userSetup(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	data, status, err := verifyJson[UserCreate](w, r)
	if status != 200 {
		return status, err
	}

	user, err := authorize(w, r, false)
	if err != nil {
		return http.StatusUnauthorized, err
	}

	res := context.db.QueryRow(fmt.Sprintf("SELECT user_type FROM base_user WHERE id = %d", user.ID))
	var userType int
	err = res.Scan(&userType)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if userType >= 0 {
		return http.StatusBadRequest, errors.New("You have already went through setup")
	}

	_, err = context.db.Exec(fmt.Sprintf("INSERT INTO user (user_id, username, picture, is_parent, gender) VALUES (%d, '%s', '%s', %t, %t)",
		user.ID, data.Username, data.Picture, data.IsParent, data.Gender))
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return http.StatusBadRequest, errors.New("You have already went through setup")
		}
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

var updateUser = updateItemFactory[User, UserCreate](
	"user INNER JOIN base_user ON base_user.id = user.user_id",
	func(item *UserCreate) string {
		return fmt.Sprintf("UPDATE user INNER JOIN base_user ON base_user.id = user.user_id SET is_parent = %t, gender = %t",
			item.IsParent, item.Gender)
	},
	scanUserRow,
)
