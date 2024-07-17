package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
)

type TuitionCenter struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Picture     string `json:"picture"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	AddressLink string `json:"address_link"`
	Description string `json:"description"`
	Website     string `json:"website"`
	CreatedAt   int    `json:"created_at"`
}

type TuitionCenterCreate struct {
	Name        string `json:"name"`
	Picture     string `json:"picture"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	AddressLink string `json:"address_link"`
	Description string `json:"description"`
	Website     string `json:"website"`
}

func scanTuitionCenterRows(item *TuitionCenter, rows *sql.Rows) error {
	var temp string
	return rows.Scan(&item.ID, &item.Name, &item.Phone, &item.Picture, &item.Address, &item.AddressLink, &item.Description, &item.Website, &temp, &item.Email, &temp, &temp, &item.CreatedAt)
}

func scanTuitionCenterRow(item *TuitionCenter, row *sql.Row) error {
	var temp string
	return row.Scan(&item.ID, &item.Name, &item.Phone, &item.Picture, &item.Address, &item.AddressLink, &item.Description, &item.Website, &temp, &item.Email, &temp, &temp, &item.CreatedAt)
}

var getAllTuitionCenters = getAllItemsFactory[TuitionCenter](
	"tuition_center INNER JOIN base_user ON base_user.id = tuition_center.user_id",
	noFilter,
	scanTuitionCenterRows,
)

var getTuitionCenterByID = getItemByIDFactory[TuitionCenter]("tuition_center INNER JOIN base_user ON base_user.id = tuition_center.user_id", scanTuitionCenterRow)

func tuitionCenterSetup(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	data, status, err := verifyJson[TuitionCenterCreate](w, r)
	if status != 200 {
		return status, err
	}

	user, err := authorize(w, r, false)
	if err != nil {
		return http.StatusUnauthorized, err
	}

	res := context.db.QueryRow(fmt.Sprintf("SELECT user_type FROM base_user WHERE id = %d", user.ID))
	var userType int
	err = res.Scan(&userType)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if userType >= 0 {
		return http.StatusBadRequest, errors.New("You have already went through setup")
	}

	_, err = context.db.Exec(fmt.Sprintf("INSERT INTO tuition_center (user_id, name, picture, phone, address, address_link, description, website) VALUES (%d, '%s', '%s', '%s', '%s', '%s', '%s', '%s')",
		user.ID, data.Name, data.Picture, data.Phone, data.Address, data.AddressLink, data.Description, data.Website))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	_, err = context.db.Exec(fmt.Sprintf("UPDATE base_user SET user_type = 2 WHERE id = %d", user.ID))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

var updateTuitionCenter = updateItemFactory[TuitionCenter, TuitionCenterCreate](
	"tuition_center INNER JOIN base_user ON base_user.id = tuition_center.user_id",
	func(item *TuitionCenterCreate) string {
		return fmt.Sprintf("UPDATE tuition_center INNER JOIN base_user ON base_user.id = tuition_center.user_id SET name = '%s', picture = '%s', phone = '%s', address = '%s', address_link = '%s', description = '%s', website = '%s'",
			item.Name, item.Picture, item.Phone, item.Address, item.AddressLink, item.Description, item.Website)
	},
	scanTuitionCenterRow,
)
