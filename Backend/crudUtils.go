package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

func noFilter(r *http.Request) string {
	return "id IS NOT NULL"
}

// A factory method for getting all rows in a table.
//
// tableName: name of the table in database.
//
// scanFn: scan the row into the provided item.
//
// Return an HTTP handler
func getAllItemsFactory[T any](
	tableName string,
	filter func(r *http.Request) string,
	scanFn func(item *T, rows *sql.Rows) error,
) func(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	getAll := func(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
		rows, err := context.db.Query(fmt.Sprintf("SELECT * FROM %s WHERE %s", tableName, filter(r)))
		if err != nil {
			return http.StatusInternalServerError, err
		}

		var items []*T
		for rows.Next() {
			item := new(T)
			err := scanFn(item, rows)
			if err != nil {
				return http.StatusInternalServerError, err
			}

			items = append(items, item)
		}

		err = json.NewEncoder(w).Encode(items)
		if err != nil {
			return http.StatusInternalServerError, err
		}

		return 200, nil
	}

	return getAll
}

// A factory method for a row by UserID in a table.
//
// tableName: name of the table in database.
//
// scanFn: scan the row into the provided item.
//
// Return an HTTP handler
func getItemByIDFactory[T any](tableName string, scanFn func(item *T, rows *sql.Row) error) func(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	getItemByID := func(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
		row := context.db.QueryRow(fmt.Sprintf("SELECT * FROM %s WHERE id = %s", tableName, r.PathValue("id")))

		item := new(T)
		err := scanFn(item, row)
		if errors.Is(err, sql.ErrNoRows) {
			return http.StatusNotFound, err
		}
		if err != nil {
			return http.StatusInternalServerError, err
		}

		err = json.NewEncoder(w).Encode(item)
		if err != nil {
			return http.StatusInternalServerError, err
		}

		return 200, nil
	}

	return getItemByID
}

// A factory method for creating a new row in a table
//
// tableName: name of the table in database.
//
// queryFn: the SQL command to put the item in the database
//
// scanFn: scan the row into the provided item (for returning to the requester)
//
// Return an HTTP handler
func createItemFactory[T any, U any](
	tableName string,
	queryFn func(item *U) string,
	scanFn func(item *T, rows *sql.Row) error,
) func(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	createItem := func(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
		item, httpCode, err := verifyJson[U](w, r)
		if err != nil {
			return httpCode, err
		}

		query := queryFn(item)
		res, err := context.db.Exec(query)
		if err != nil {
			match, _ := regexp.MatchString("Duplicate entry '.*' for key 'base_user\\.username'", err.Error())
			if match {
				return http.StatusBadRequest, errors.New("This username has already existed")
			}

			return http.StatusInternalServerError, err
		}

		newID, _ := res.LastInsertId()
		row := context.db.QueryRow(fmt.Sprintf("SELECT * FROM %s WHERE id = %d", tableName, newID))
		outputItem := new(T)
		err = scanFn(outputItem, row)
		if err != nil {
			return http.StatusInternalServerError, err
		}

		err = json.NewEncoder(w).Encode(outputItem)
		if err != nil {
			return http.StatusInternalServerError, err
		}

		return 201, nil
	}

	return createItem
}

// A factory method for updating a row in a table
//
// tableName: name of the table in database.
//
// queryFn: the SQL command to put the item in the database
//
// scanFn: scan the row into the provided item (for returning to the requester)
//
// Return an HTTP handler
func updateItemFactory[T any, U any](
	tableName string,
	queryFn func(item *U) string,
	scanFn func(item *T, rows *sql.Row) error,
) func(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	updateItem := func(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
		item, httpCode, err := verifyJson[U](w, r)
		if err != nil {
			return httpCode, err
		}

		row := context.db.QueryRow(fmt.Sprintf("SELECT * FROM %s WHERE id = %s", tableName, r.PathValue("id")))
		initialItem := new(T)
		err = scanFn(initialItem, row)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return http.StatusNotFound, err
			}
			return http.StatusInternalServerError, err
		}

		query := fmt.Sprintf("%s WHERE id = %s", queryFn(item), r.PathValue("id"))
		res, err := context.db.Exec(query)
		if err != nil {
			return http.StatusInternalServerError, err
		}

		newID, _ := res.RowsAffected()
		log.Println(newID)
		row = context.db.QueryRow(fmt.Sprintf("SELECT * FROM %s WHERE id = %s", tableName, r.PathValue("id")))

		// UPDATE where every value remains the same will return a newID of 0
		// as it is not considered a change
		if newID == 0 {
			row = context.db.QueryRow(fmt.Sprintf("SELECT * FROM %s WHERE id = %s", tableName, r.PathValue("id")))
		}

		outputItem := new(T)
		err = scanFn(outputItem, row)
		if err != nil {
			return http.StatusInternalServerError, err
		}

		err = json.NewEncoder(w).Encode(outputItem)
		if err != nil {
			return http.StatusInternalServerError, err
		}

		return 200, nil
	}

	return updateItem
}

// A factory method for deleting a row by UserID in a table.
//
// tableName: name of the table in database.
//
// Return an HTTP handler
func deleteItemByIDFactory(tableName string) func(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	deleteItemByID := func(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
		_, err := context.db.Exec(fmt.Sprintf("DELETE FROM %s WHERE id = %s", tableName, r.PathValue("id")))
		if errors.Is(err, sql.ErrNoRows) {
			return http.StatusNotFound, err
		}

		if err != nil {
			return http.StatusInternalServerError, err
		}

		return 200, nil
	}

	return deleteItemByID
}
