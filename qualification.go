package main

import "database/sql"

type Qualification struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Time        int    `json:"time"`
	LevelID     int    `json:"level_id"`
	TutorID     int    `json:"tutor_id"`
}

var getAllQualifications = getAllItemsFactory[Qualification]("qualification", func(item *Qualification, rows *sql.Rows) error {
	return rows.Scan(&item.ID, &item.Name, &item.Description, &item.Time, &item.LevelID, &item.TutorID)
})

var getQualificationByID = getItemByIDFactory[Qualification]("qualification", func(item *Qualification, row *sql.Row) error {
	return row.Scan(&item.ID, &item.Name, &item.Description, &item.Time, &item.LevelID, &item.TutorID)
})

var deleteQualificationByID = deleteItemByIDFactory("qualification")
