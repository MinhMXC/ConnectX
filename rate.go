package main

import (
	"database/sql"
	"fmt"
)

type Rate struct {
	ID              int     `json:"id"`
	Amount          float32 `json:"amount"`
	IsOpen          bool    `json:"is_open"`
	SubjectID       int     `json:"subject_id"`
	TutorID         *int    `json:"tutor_id"`
	TuitionCenterID *int    `json:"tuition_center_id"`
}

type RateCreate struct {
	Amount          float32 `json:"amount"`
	IsOpen          bool    `json:"is_open"`
	SubjectID       int     `json:"subject_id"`
	TutorID         *int    `json:"tutor_id"`
	TuitionCenterID *int    `json:"tuition_center_id"`
}

var getAllRates = getAllItemsFactory[Rate]("rate", func(item *Rate, rows *sql.Rows) error {
	return rows.Scan(&item.ID, &item.Amount, &item.IsOpen, &item.SubjectID, &item.TutorID, &item.TuitionCenterID)
})

var getRateByID = getItemByIDFactory[Rate]("rate", func(item *Rate, row *sql.Row) error {
	return row.Scan(&item.ID, &item.Amount, &item.IsOpen, &item.SubjectID, &item.TutorID, &item.TuitionCenterID)
})

var createRate = createItemFactory[Rate, RateCreate](
	"rate",
	func(item *RateCreate) string {
		if item.TuitionCenterID == nil {
			return fmt.Sprintf("INSERT INTO rate (amount, is_open, subject_id, tutor_id, tuition_center_id) VALUES (%f, %t, %d, %d, null)",
				item.Amount, item.IsOpen, item.SubjectID, *item.TutorID)
		} else {
			return fmt.Sprintf("INSERT INTO rate (amount, is_open, subject_id, tutor_id, tuition_center_id) VALUES (%f, %t, %d, null, %d)",
				item.Amount, item.IsOpen, item.SubjectID, *item.TuitionCenterID)
		}
	},
	func(item *Rate, row *sql.Row) error {
		return row.Scan(&item.ID, &item.Amount, &item.IsOpen, &item.SubjectID, &item.TutorID, &item.TuitionCenterID)
	},
)

var deleteRateByID = deleteItemByIDFactory("rate")
