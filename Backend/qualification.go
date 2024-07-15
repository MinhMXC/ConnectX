package main

import (
	"database/sql"
	"fmt"
	"net/http"
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

func scanQualificationRows(item *Qualification, rows *sql.Rows) error {
	return rows.Scan(&item.ID, &item.Name, &item.Description, &item.Time, &item.LevelID, &item.TutorID)
}

func scanQualificationRow(item *Qualification, rows *sql.Row) error {
	return rows.Scan(&item.ID, &item.Name, &item.Description, &item.Time, &item.LevelID, &item.TutorID)
}

var getAllQualifications = getAllItemsFactory[Qualification](
	"qualification",
	noFilter,
	scanQualificationRows,
)

var getQualificationByTutorID = getAllItemsFactory[Qualification](
	"qualification",
	func(r *http.Request) string {
		return fmt.Sprintf("tutor_id = %s", r.PathValue("id"))
	},
	scanQualificationRows,
)

var getQualificationByID = getItemByIDFactory[Qualification]("qualification", scanQualificationRow)

var createQualification = createItemFactory[Qualification, QualificationCreate](
	"qualification",
	func(item *QualificationCreate) string {
		return fmt.Sprintf("INSERT INTO qualification (name, description, time, level_id, tutor_id) VALUES ('%s', '%s', %d, %d, %d)",
			item.Name, item.Description, item.Time, item.LevelID, item.TutorID)
	},
	scanQualificationRow,
)

var updateQualification = updateItemFactory[Qualification, QualificationCreate](
	"qualification",
	func(item *QualificationCreate) string {
		return fmt.Sprintf("UPDATE qualification SET name = '%s', description = '%s', time = %d, level_id = %d, tutor_id = %d",
			item.Name, item.Description, item.Time, item.LevelID, item.TutorID)
	},
	scanQualificationRow,
)

var deleteQualificationByID = deleteItemByIDFactory("qualification")
