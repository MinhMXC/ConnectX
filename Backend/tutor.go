package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type Tutor struct {
	UserID      int    `json:"user_id"`
	Username    string `json:"username"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Gender      bool   `json:"gender"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
}

type TutorCreate struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Gender      bool   `json:"gender"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
}

func scanTutorRows(item *Tutor, rows *sql.Rows) error {
	var temp string
	return rows.Scan(&item.UserID, &item.Name, &item.Age, &item.Gender, &item.Phone, &item.Description, &temp, &item.Username, &temp, &temp, &item.CreatedAt)
}

func scanTutorRow(item *Tutor, row *sql.Row) error {
	var temp string
	return row.Scan(&item.UserID, &item.Name, &item.Age, &item.Gender, &item.Phone, &item.Description, &temp, &item.Username, &temp, &temp, &item.CreatedAt)
}

var getAllTutors = getAllItemsFactory[Tutor]("tutor INNER JOIN base_user ON base_user.id = tutor.user_id", scanTutorRows)

var getTutorByID = getItemByIDFactory[Tutor]("tutor INNER JOIN base_user ON base_user.id = tutor.user_id", scanTutorRow)

func tutorSetup(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	data, status, err := verifyJson[TutorCreate](w, r)
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

	_, err = context.db.Exec(fmt.Sprintf("INSERT INTO tutor (user_id, name, age, gender, phone, description) VALUES (%d, '%s', %d, %t, '%s', '%s')",
		user.ID, data.Name, data.Age, data.Gender, data.Phone, data.Description))
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return http.StatusBadRequest, errors.New("You have already went through setup")
		}
		return http.StatusInternalServerError, err
	}

	_, err = context.db.Exec(fmt.Sprintf("UPDATE base_user SET user_type = 2 WHERE id = %d", user.ID))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

var updateTutor = updateItemFactory[Tutor, TutorCreate](
	"tutor INNER JOIN base_user ON base_user.id = tutor.user_id",
	func(item *TutorCreate) string {
		return fmt.Sprintf("UPDATE tutor INNER JOIN base_user ON base_user.id = tutor.user_id SET name = '%s', age = %d, gender = %t, phone = '%s', description = '%s'",
			item.Name, item.Age, item.Gender, item.Phone, item.Description)
	},
	scanTutorRow,
)
