package main

import (
	"database/sql"
	"fmt"
	"net/http"
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

func scanRateRows(item *Rate, rows *sql.Rows) error {
	return rows.Scan(&item.ID, &item.Amount, &item.IsOpen, &item.SubjectID, &item.TutorID, &item.TuitionCenterID)
}

func scanRateRow(item *Rate, row *sql.Row) error {
	return row.Scan(&item.ID, &item.Amount, &item.IsOpen, &item.SubjectID, &item.TutorID, &item.TuitionCenterID)
}

var getAllRates = getAllItemsFactory[Rate](
	"rate",
	noFilter,
	scanRateRows,
)

var getRateByID = getItemByIDFactory[Rate]("rate", scanRateRow)

var getRateByTutorID = getAllItemsFactory[Rate](
	"rate",
	func(r *http.Request) string {
		return fmt.Sprintf("tutor_id = %s", r.PathValue("id"))
	},
	scanRateRows,
)

var getRateByTuitionCenterID = getAllItemsFactory[Rate](
	"rate",
	func(r *http.Request) string {
		return fmt.Sprintf("tuition_center_id = %s", r.PathValue("id"))
	},
	scanRateRows,
)

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
	scanRateRow,
)

// TODO: add middleware to ensure that only one of tutor_id and tuition_center_id can be null
var updateRate = updateItemFactory[Rate, RateCreate](
	"rate",
	func(item *RateCreate) string {
		if item.TuitionCenterID == nil {
			return fmt.Sprintf("UPDATE rate SET amount = %f, is_open = %t, subject_id = %d, tutor_id = %d, tuition_center_id = null",
				item.Amount, item.IsOpen, item.SubjectID, *item.TutorID)
		} else {
			return fmt.Sprintf("UPDATE rate SET amount = %f, is_open = %t, subject_id = %d, tutor_id = null, tuition_center_id = %d",
				item.Amount, item.IsOpen, item.SubjectID, *item.TuitionCenterID)
		}
	},
	scanRateRow,
)

var deleteRateByID = deleteItemByIDFactory("rate")
