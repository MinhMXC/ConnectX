package main

import (
	"database/sql"
	"fmt"
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

type TutorUpdate struct {
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

var updateTutor = updateItemFactory[Tutor, TutorUpdate](
	"tutor INNER JOIN base_user ON base_user.id = tutor.user_id",
	func(item *TutorUpdate) string {
		return fmt.Sprintf("UPDATE tutor INNER JOIN base_user ON base_user.id = tutor.user_id SET name = '%s', age = %d, gender = %t, phone = '%s', description = '%s'",
			item.Name, item.Age, item.Gender, item.Phone, item.Description)
	},
	scanTutorRow,
)
