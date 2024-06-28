package main

import (
	"database/sql"
	"fmt"
)

type TuitionCenter struct {
	UserID      int    `json:"user_id"`
	Username    string `json:"user_name"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	AddressLink string `json:"address_link"`
	Description string `json:"description"`
	Website     string `json:"website"`
	CreatedAt   int    `json:"created_at"`
}

type TuitionCenterCreate struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	AddressLink string `json:"address_link"`
	Description string `json:"description"`
	Website     string `json:"website"`
}

func scanTuitionCenterRows(item *TuitionCenter, rows *sql.Rows) error {
	var temp string
	return rows.Scan(&item.UserID, &item.Name, &item.Phone, &item.Address, &item.AddressLink, &item.Description, &item.Website, &temp, &item.Username, &temp, &temp, &item.CreatedAt)
}

func scanTuitionCenterRow(item *TuitionCenter, row *sql.Row) error {
	var temp string
	return row.Scan(&item.UserID, &item.Name, &item.Phone, &item.Address, &item.AddressLink, &item.Description, &item.Website, &temp, &item.Username, &temp, &temp, &item.CreatedAt)
}

var getAllTuitionCenters = getAllItemsFactory[TuitionCenter]("tuition_center INNER JOIN base_user ON base_user.id = tuition_center.user_id", scanTuitionCenterRows)

var getTuitionCenterByID = getItemByIDFactory[TuitionCenter]("tuition_center INNER JOIN base_user ON base_user.id = tuition_center.user_id", scanTuitionCenterRow)

var updateTuitionCenter = updateItemFactory[TuitionCenter, TuitionCenterCreate](
	"tuition_center INNER JOIN base_user ON base_user.id = tuition_center.user_id",
	func(item *TuitionCenterCreate) string {
		return fmt.Sprintf("UPDATE tuition_center INNER JOIN base_user ON base_user.id = tuition_center.user_id SET name = '%s', phone = '%s', address = '%s', address_link = '%s', description = '%s', website = '%s'",
			item.Name, item.Phone, item.Address, item.AddressLink, item.Description, item.Website)
	},
	scanTuitionCenterRow,
)
