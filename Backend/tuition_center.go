package main

import (
	"database/sql"
	"fmt"
	"time"
)

type TuitionCenter struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	AddressLink string `json:"address_link"`
	Description string `json:"description"`
	Website     string `json:"website"`
	CreatedAt   string `json:"created_at"`
}

type TuitionCenterCreate struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	AddressLink string `json:"address_link"`
	Description string `json:"description"`
	Website     string `json:"website"`
}

var getAllTuitionCenters = getAllItemsFactory[TuitionCenter]("tuition_center", func(item *TuitionCenter, rows *sql.Rows) error {
	return rows.Scan(&item.ID, &item.Name, &item.Phone, &item.Address, &item.AddressLink,
		&item.Description, &item.Website, &item.CreatedAt)
})

var getTuitionCenterByID = getItemByIDFactory[TuitionCenter]("tuition_center", func(item *TuitionCenter, row *sql.Row) error {
	return row.Scan(&item.ID, &item.Name, &item.Phone, &item.Address, &item.AddressLink,
		&item.Description, &item.Website, &item.CreatedAt)
})

var createTuitionCenter = createItemFactory[TuitionCenter, TuitionCenterCreate](
	"tuition_center",
	func(item *TuitionCenterCreate) string {
		return fmt.Sprintf("INSERT INTO tuition_center (name, phone, address, address_link, description, website, created_at) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', %d)",
			item.Name, item.Phone, item.Address, item.AddressLink, item.Description, item.Website, time.Now().Unix())
	},
	func(item *TuitionCenter, row *sql.Row) error {
		return row.Scan(&item.ID, &item.Name, &item.Phone, &item.Address, &item.AddressLink, &item.Description, &item.Website, &item.CreatedAt)
	},
)

var updateTuitionCenter = updateItemFactory[TuitionCenter, TuitionCenterCreate](
	"tuition_center",
	func(item *TuitionCenterCreate) string {
		return fmt.Sprintf("UPDATE tuition_center SET name = '%s', phone = '%s', address = '%s', address_link = '%s', description = '%s', website = '%s'",
			item.Name, item.Phone, item.Address, item.AddressLink, item.Description, item.Website)
	},
	func(item *TuitionCenter, row *sql.Row) error {
		return row.Scan(&item.ID, &item.Name, &item.Phone, &item.Address, &item.AddressLink, &item.Description, &item.Website, &item.CreatedAt)
	},
)

var deleteTuitionCenterByID = deleteItemByIDFactory("tuition_center")
