package main

import "database/sql"

type Subject struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Level_ID int    `json:"level_id"`
}

var getAllSubjects = getAllItemsFactory[Subject]("subject", func(item *Subject, rows *sql.Rows) error {
	return rows.Scan(&item.ID, &item.Name, &item.Level_ID)
})

var getSubjectByID = getItemByIDFactory[Subject]("subject", func(item *Subject, row *sql.Row) error {
	return row.Scan(&item.ID, &item.Name, &item.Level_ID)
})

var deleteSubjectByID = deleteItemByIDFactory("subject")
