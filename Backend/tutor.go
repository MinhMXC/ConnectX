package main

import (
	"database/sql"
	"fmt"
	"time"
)

type Tutor struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
	Gender      bool   `json:"gender"`
	CreatedAt   int    `json:"created_at"`
}

type TutorCreate struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
	Gender      bool   `json:"gender"`
}

var getAllTutors = getAllItemsFactory[Tutor]("tutor", func(item *Tutor, rows *sql.Rows) error {
	return rows.Scan(&item.ID, &item.Name, &item.Age, &item.Phone, &item.Description,
		&item.Gender, &item.CreatedAt)
})

var getTutorByID = getItemByIDFactory[Tutor]("tutor", func(item *Tutor, row *sql.Row) error {
	return row.Scan(&item.ID, &item.Name, &item.Age, &item.Phone, &item.Description,
		&item.Gender, &item.CreatedAt)
})

var createTutor = createItemFactory[Tutor, TutorCreate](
	"tutor",
	func(item *TutorCreate) string {
		return fmt.Sprintf("INSERT INTO tutor (name, age, phone, description, gender, created_at) VALUES ('%s', %d, '%s', '%s', %t, %d)",
			item.Name, item.Age, item.Phone, item.Description, item.Gender, time.Now().Unix())
	},
	func(item *Tutor, row *sql.Row) error {
		return row.Scan(&item.ID, &item.Name, &item.Age, &item.Phone, &item.Description, &item.Gender, &item.CreatedAt)
	},
)

var updateTutor = updateItemFactory[Tutor, TutorCreate](
	"tutor",
	func(item *TutorCreate) string {
		return fmt.Sprintf("UPDATE tutor SET name = '%s', age = %d, phone = '%s', description = '%s', gender = %t",
			item.Name, item.Age, item.Phone, item.Description, item.Gender)
	},
	func(item *Tutor, row *sql.Row) error {
		return row.Scan(&item.ID, &item.Name, &item.Age, &item.Phone, &item.Description, &item.Gender, &item.CreatedAt)
	},
)

var deleteTutorByID = deleteItemByIDFactory("tutor")
