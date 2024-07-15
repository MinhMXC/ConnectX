package main

import (
	"database/sql"
	"fmt"
)

type Level struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type LevelCreate struct {
	Name string `json:"name"`
}

var getAllLevels = getAllItemsFactory[Level](
	"level",
	noFilter,
	func(item *Level, rows *sql.Rows) error {
		return rows.Scan(&item.ID, &item.Name)
	},
)

var createLevel = createItemFactory[Level, LevelCreate](
	"level",
	func(item *LevelCreate) string {
		return fmt.Sprintf("INSERT INTO level (name) VALUES ('%s')", item.Name)
	},
	func(item *Level, row *sql.Row) error {
		return row.Scan(&item.ID, &item.Name)
	},
)

var updateLevel = updateItemFactory[Level, LevelCreate](
	"level",
	func(item *LevelCreate) string {
		return fmt.Sprintf("UPDATE level SET name = '%s'", item.Name)
	},
	func(item *Level, row *sql.Row) error {
		return row.Scan(&item.ID, &item.Name)
	},
)

var getLevelByID = getItemByIDFactory[Level]("level", func(item *Level, row *sql.Row) error {
	return row.Scan(&item.ID, &item.Name)
})

var deleteLevelByID = deleteItemByIDFactory("level")
