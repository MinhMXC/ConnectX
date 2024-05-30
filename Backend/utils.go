package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/golang-jwt/jwt/v5"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.MapClaims
}

var secretKey = []byte("secret")

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":      "connectx",
		"username": username,
		"exp":      time.Now().Add(time.Minute * 5).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyToken(tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if _, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return nil
	}

	return fmt.Errorf("Invalid token")
}

func verifyJson[T any](w http.ResponseWriter, r *http.Request) (*T, int, error) {
	ct := r.Header.Get("Content-Type")
	if ct != "" {
		mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
		if mediaType != "application/json" {
			msg := "Content-Type header is not application/json"
			return nil, http.StatusUnsupportedMediaType, errors.New(msg)
		}
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	item := new(T)
	err := decoder.Decode(item)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			return nil, http.StatusBadRequest, errors.New(msg)
		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := "Request body contains badly-formed JSON"
			return nil, http.StatusBadRequest, errors.New(msg)
		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			return nil, http.StatusBadRequest, errors.New(msg)
		case strings.HasPrefix(err.Error(), "json: unknown field"):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body conatains unknown field %s", fieldName)
			return nil, http.StatusBadRequest, errors.New(msg)
		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			return nil, http.StatusBadRequest, errors.New(msg)
		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			return nil, http.StatusBadRequest, errors.New(msg)
		default:
			log.Print(err.Error())
			return nil, http.StatusInternalServerError, nil
		}
	}

	err = decoder.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
		msg := "Request body must only contain a single JSON object"
		return nil, http.StatusBadRequest, errors.New(msg)
	}

	return item, http.StatusOK, nil
}
