package main

import (
	"database/sql"
	"fmt"
)

type Level struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type LevelCreate struct {
	Name string `json:"name"`
}

var getAllLevels = getAllItemsFactory[Level]("level", func(item *Level, rows *sql.Rows) error {
	return rows.Scan(&item.ID, &item.Name)
})

var createLevel = createItemFactory[Level, LevelCreate](
	"level",
	func(item *LevelCreate) string {
		return fmt.Sprintf("INSERT INTO level (name) VALUES ('%s')", item.Name)
	},
	func(item *Level, row *sql.Row) error {
		return row.Scan(&item.ID, &item.Name)
	},
)

//func createLevel(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
//	ct := r.Header.Get("Content-Type")
//	if ct != "" {
//		mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
//		if mediaType != "application/json" {
//			msg := "Content-Type header is not application/json"
//			return http.StatusUnsupportedMediaType, errors.New(msg)
//		}
//	}
//
//	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
//	decoder := json.NewDecoder(r.Body)
//	decoder.DisallowUnknownFields()
//
//	var level LevelCreate
//	err := decoder.Decode(&level)
//	if err != nil {
//		var syntaxError *json.SyntaxError
//		var unmarshalTypeError *json.UnmarshalTypeError
//
//		switch {
//		case errors.As(err, &syntaxError):
//			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
//			return http.StatusBadRequest, errors.New(msg)
//		case errors.Is(err, io.ErrUnexpectedEOF):
//			msg := "Request body contains badly-formed JSON"
//			return http.StatusBadRequest, errors.New(msg)
//		case errors.As(err, &unmarshalTypeError):
//			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
//			return http.StatusBadRequest, errors.New(msg)
//		case strings.HasPrefix(err.Error(), "json: unknown field"):
//			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
//			msg := fmt.Sprintf("Request body conatains unknown field %s", fieldName)
//			return http.StatusBadRequest, errors.New(msg)
//		case errors.Is(err, io.EOF):
//			msg := "Request body must not be empty"
//			return http.StatusBadRequest, errors.New(msg)
//		case err.Error() == "http: request body too large":
//			msg := "Request body must not be larger than 1MB"
//			return http.StatusBadRequest, errors.New(msg)
//		default:
//			log.Print(err.Error())
//			return http.StatusInternalServerError, nil
//		}
//	}
//
//	err = decoder.Decode(&struct{}{})
//	if !errors.Is(err, io.EOF) {
//		msg := "Request body must only contain a single JSON object"
//		return http.StatusBadRequest, errors.New(msg)
//	}
//
//	res, err := context.db.Exec(fmt.Sprintf("INSERT INTO level (name) VALUES ('%s')", level.Name))
//	if err != nil {
//		return http.StatusInternalServerError, err
//	}
//
//	newID, _ := res.LastInsertId()
//	row := context.db.QueryRow(fmt.Sprintf("SELECT * FROM level WHERE id = %d", newID))
//
//	outputLevel := new(Level)
//	err = row.Scan(&outputLevel.ID, &outputLevel.Name)
//	if err != nil {
//		return http.StatusInternalServerError, err
//	}
//
//	err = json.NewEncoder(w).Encode(outputLevel)
//	if err != nil {
//		return http.StatusInternalServerError, err
//	}
//
//	return 201, nil
//}

var getLevelByID = getItemByIDFactory[Level]("level", func(item *Level, row *sql.Row) error {
	return row.Scan(&item.ID, &item.Name)
})

var deleteLevelByID = deleteItemByIDFactory("level")
