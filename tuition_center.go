package main

import "database/sql"

type TuitionCenter struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	AddressLink string `json:"address_link"`
	Description string `json:"description"`
	Website     string `json:"website"`
	IsOpen      bool   `json:"is_open"`
	CreatedAt   string `json:"created_at"`
}

var getAllTuitionCenters = getAllItemsFactory[TuitionCenter]("tuition_center", func(item *TuitionCenter, rows *sql.Rows) error {
	return rows.Scan(&item.ID, &item.Name, &item.Phone, &item.Address, &item.AddressLink,
		&item.Description, &item.Website, &item.IsOpen, &item.CreatedAt)
})

var getTuitionCenterByID = getItemByIDFactory[TuitionCenter]("tuition_center", func(item *TuitionCenter, row *sql.Row) error {
	return row.Scan(&item.ID, &item.Name, &item.Phone, &item.Address, &item.AddressLink,
		&item.Description, &item.Website, &item.IsOpen, &item.CreatedAt)
})

var deleteTuitionCenterByID = deleteItemByIDFactory("tuition_center")
