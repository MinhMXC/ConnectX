package main

import "database/sql"

type Rate struct {
	ID              int     `json:"id"`
	Amount          float32 `json:"amount"`
	SubjectID       int     `json:"subject_id"`
	TutorID         *int    `json:"tutor_id"`
	TuitionCenterID *int    `json:"tuition_center_id"`
}

var getAllRates = getAllItemsFactory[Rate]("rate", func(item *Rate, rows *sql.Rows) error {
	return rows.Scan(&item.ID, &item.Amount, &item.SubjectID, &item.TutorID, &item.TuitionCenterID)
})

var getRateByID = getItemByIDFactory[Rate]("rate", func(item *Rate, row *sql.Row) error {
	return row.Scan(&item.ID, &item.Amount, &item.SubjectID, &item.TutorID, &item.TuitionCenterID)
})

var deleteRateByID = deleteItemByIDFactory("rate")
