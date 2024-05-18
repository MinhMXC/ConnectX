package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func getAllItemsFactory[T any](tableName string, fn func(item *T, rows *sql.Rows) error) func(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	getAll := func(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
		rows, err := context.db.Query(fmt.Sprintf("SELECT * FROM %s", tableName))
		if err != nil {
			return http.StatusInternalServerError, err
		}

		var items []*T
		for rows.Next() {
			item := new(T)
			err := fn(item, rows)
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

func getItemByIDFactory[T any](tableName string, fn func(item *T, rows *sql.Row) error) func(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	getItemByID := func(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
		row := context.db.QueryRow(fmt.Sprintf("SELECT * FROM %s WHERE id = %s", tableName, r.PathValue("id")))

		item := new(T)
		err := fn(item, row)
		if errors.Is(err, sql.ErrNoRows) {
			return http.StatusNotFound, err
		}

		err = json.NewEncoder(w).Encode(item)
		if err != nil {
			return http.StatusInternalServerError, err
		}

		return 200, nil
	}

	return getItemByID
}

func deleteItemByIDFactory(tableName string) func(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	deleteItemByID := func(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
		_, err := context.db.Exec(fmt.Sprintf("DELETE FROM %s WHERE id = %s", tableName, r.PathValue("id")))
		if errors.Is(err, sql.ErrNoRows) {
			return http.StatusNotFound, err
		}

		if err != nil {
			fmt.Sprintf(err.Error())
			return http.StatusInternalServerError, err
		}

		return 200, nil
	}

	return deleteItemByID
}
