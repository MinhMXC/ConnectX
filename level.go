package main

import (
	"database/sql"
)

type Level struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var getAllLevels = getAllItemsFactory[Level]("level", func(item *Level, rows *sql.Rows) error {
	return rows.Scan(&item.ID, &item.Name)
})

var getLevelByID = getItemByIDFactory[Level]("level", func(item *Level, row *sql.Row) error {
	return row.Scan(&item.ID, &item.Name)
})

var deleteLevelByID = deleteItemByIDFactory("level")
