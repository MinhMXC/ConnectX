package main

import (
	"database/sql"
	"fmt"
)

type Qualification struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Time        int    `json:"time"`
	LevelID     int    `json:"level_id"`
	TutorID     int    `json:"tutor_id"`
}

type QualificationCreate struct {
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

var createQualification = createItemFactory[Qualification, QualificationCreate](
	"qualification",
	func(item *QualificationCreate) string {
		return fmt.Sprintf("INSERT INTO qualification (name, description, time, level_id, tutor_id) VALUES ('%s', '%s', %d, %d, %d)",
			item.Name, item.Description, item.Time, item.LevelID, item.TutorID)
	},
	func(item *Qualification, row *sql.Row) error {
		return row.Scan(&item.ID, &item.Name, &item.Description, &item.Time, &item.LevelID, &item.TutorID)
	},
)

var updateQualification = updateItemFactory[Qualification, QualificationCreate](
	"qualification",
	func(item *QualificationCreate) string {
		return fmt.Sprintf("UPDATE qualification SET name = '%s', description = '%s', time = %d, level_id = %d, tutor_id = %d",
			item.Name, item.Description, item.Time, item.LevelID, item.TutorID)
	},
	func(item *Qualification, row *sql.Row) error {
		return row.Scan(&item.ID, &item.Name, &item.Description, &item.Time, &item.LevelID, &item.TutorID)
	},
)

var deleteQualificationByID = deleteItemByIDFactory("qualification")
