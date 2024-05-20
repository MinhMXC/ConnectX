package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func getAllItemsFactory[T any](tableName string, scanFn func(item *T, rows *sql.Rows) error) func(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	getAll := func(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
		rows, err := context.db.Query(fmt.Sprintf("SELECT * FROM %s", tableName))
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

func getItemByIDFactory[T any](tableName string, scanFn func(item *T, rows *sql.Row) error) func(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	getItemByID := func(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
		row := context.db.QueryRow(fmt.Sprintf("SELECT * FROM %s WHERE id = %s", tableName, r.PathValue("id")))

		item := new(T)
		err := scanFn(item, row)
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

func createItemFactory[T any, U any](tableName string, query func(item *U) string, scanFn func(item *T, rows *sql.Row) error) func(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	createItem := func(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
		ct := r.Header.Get("Content-Type")
		if ct != "" {
			mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
			if mediaType != "application/json" {
				msg := "Content-Type header is not application/json"
				return http.StatusUnsupportedMediaType, errors.New(msg)
			}
		}

		r.Body = http.MaxBytesReader(w, r.Body, 1048576)
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		item := new(U)
		err := decoder.Decode(item)
		if err != nil {
			var syntaxError *json.SyntaxError
			var unmarshalTypeError *json.UnmarshalTypeError

			switch {
			case errors.As(err, &syntaxError):
				msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
				return http.StatusBadRequest, errors.New(msg)
			case errors.Is(err, io.ErrUnexpectedEOF):
				msg := "Request body contains badly-formed JSON"
				return http.StatusBadRequest, errors.New(msg)
			case errors.As(err, &unmarshalTypeError):
				msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
				return http.StatusBadRequest, errors.New(msg)
			case strings.HasPrefix(err.Error(), "json: unknown field"):
				fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
				msg := fmt.Sprintf("Request body conatains unknown field %s", fieldName)
				return http.StatusBadRequest, errors.New(msg)
			case errors.Is(err, io.EOF):
				msg := "Request body must not be empty"
				return http.StatusBadRequest, errors.New(msg)
			case err.Error() == "http: request body too large":
				msg := "Request body must not be larger than 1MB"
				return http.StatusBadRequest, errors.New(msg)
			default:
				log.Print(err.Error())
				return http.StatusInternalServerError, nil
			}
		}

		err = decoder.Decode(&struct{}{})
		if !errors.Is(err, io.EOF) {
			msg := "Request body must only contain a single JSON object"
			return http.StatusBadRequest, errors.New(msg)
		}

		res, err := context.db.Exec(query(item))
		if err != nil {
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
