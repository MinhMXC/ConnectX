package main

import (
	"database/sql"
	"fmt"
	"time"
)

type Request struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Rate        float32 `json:"rate"`
	CreatedAt   int     `json:"created_at"`
	UserID      int     `json:"user_id"`
	SubjectID   int     `json:"subject_id"`
	LevelID     int     `json:"level_id"`
}

type RequestCreate struct {
	Description string  `json:"description"`
	Rate        float32 `json:"rate"`
	UserID      int     `json:"user_id"`
	SubjectID   int     `json:"subject_id"`
	LevelID     int     `json:"level_id"`
}

var getAllRequests = getAllItemsFactory[Request]("request", func(item *Request, rows *sql.Rows) error {
	return rows.Scan(&item.ID, &item.Description, &item.Rate, &item.CreatedAt, &item.UserID, &item.SubjectID, &item.LevelID)
})

var getRequestByID = getItemByIDFactory[Request]("request", func(item *Request, row *sql.Row) error {
	return row.Scan(&item.ID, &item.Description, &item.Rate, &item.CreatedAt, &item.UserID, &item.SubjectID, &item.LevelID)
})

var createRequest = createItemFactory[Request, RequestCreate](
	"request",
	func(item *RequestCreate) string {
		return fmt.Sprintf("INSERT INTO request (description, rate, created_at, user_id, subject_id, level_id) VALUES ('%s', %f, %d, %d, %d, %d)",
			item.Description, item.Rate, time.Now().Unix(), item.UserID, item.SubjectID, item.LevelID)
	},
	func(item *Request, row *sql.Row) error {
		return row.Scan(&item.ID, &item.Description, &item.Rate, &item.CreatedAt, &item.UserID, &item.SubjectID, &item.LevelID)
	},
)

var deleteRequestByID = deleteItemByIDFactory("request")
