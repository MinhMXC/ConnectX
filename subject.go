package main

import (
	"database/sql"
	"fmt"
)

type Subject struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	LevelID int    `json:"level_id"`
}

type SubjectCreate struct {
	Name    string `json:"name"`
	LevelID int    `json:"level_id"`
}

var getAllSubjects = getAllItemsFactory[Subject]("subject", func(item *Subject, rows *sql.Rows) error {
	return rows.Scan(&item.ID, &item.Name, &item.LevelID)
})

var getSubjectByID = getItemByIDFactory[Subject]("subject", func(item *Subject, row *sql.Row) error {
	return row.Scan(&item.ID, &item.Name, &item.LevelID)
})

var createSubject = createItemFactory[Subject, SubjectCreate](
	"subject",
	func(item *SubjectCreate) string {
		return fmt.Sprintf("INSERT INTO subject (name, level_id) VALUES ('%s', %d)", item.Name, item.LevelID)
	},
	func(item *Subject, row *sql.Row) error {
		return row.Scan(&item.ID, &item.Name, &item.LevelID)
	},
)

var deleteSubjectByID = deleteItemByIDFactory("subject")
