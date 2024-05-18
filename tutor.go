package main

import "database/sql"

type Tutor struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
	IsOpen      bool   `json:"is_open"`
	Gender      bool   `json:"gender"`
	CreatedAt   int    `json:"created_at"`
}

var getAllTutors = getAllItemsFactory[Tutor]("tutor", func(item *Tutor, rows *sql.Rows) error {
	return rows.Scan(&item.ID, &item.Name, &item.Age, &item.Phone, &item.Description,
		&item.IsOpen, &item.Gender, &item.CreatedAt)
})

var getTutorByID = getItemByIDFactory[Tutor]("tutor", func(item *Tutor, row *sql.Row) error {
	return row.Scan(&item.ID, &item.Name, &item.Age, &item.Phone, &item.Description,
		&item.IsOpen, &item.Gender, &item.CreatedAt)
})

var deleteTutorByID = deleteItemByIDFactory("tutor")
