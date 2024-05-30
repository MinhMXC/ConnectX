package main

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	IsParent  bool   `json:"is_parent"`
	CreatedAt int    `json:"created_at"`
}

type UserCreate struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IsParent bool   `json:"is_parent"`
}

type UserUpdate struct {
	Username string `json:"username"`
	IsParent bool   `json:"is_parent"`
}

func scanUserRows(item *User, rows *sql.Rows) error {
	var temp string
	return rows.Scan(&item.ID, &item.Username, &temp, &item.IsParent, &item.CreatedAt)
}

func scanUserRow(item *User, row *sql.Row) error {
	var temp string
	return row.Scan(&item.ID, &item.Username, &temp, &item.IsParent, &item.CreatedAt)
}

func generateRandomString(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}

// Should be accessed only when the JWT expire
// Return a new JWT and a new refreshToken
func refreshHandler(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	oldRefreshToken := r.Header.Get("Refresh-Token")
	if oldRefreshToken == "" {
		return http.StatusUnauthorized, errors.New("Refresh-Token header missing")
	}

	row := context.db.QueryRow("SELECT refresh.expiry, user.username FROM refresh INNER JOIN user ON refresh.user_id = user.id WHERE refresh_token=?", oldRefreshToken)
	var expiry int64
	var username string
	err := row.Scan(&expiry, &username)
	if err != nil && err == sql.ErrNoRows {
		return http.StatusUnauthorized, errors.New("Refresh-Token not found")
	} else if err != nil {
		return http.StatusInternalServerError, err
	}

	if time.Now().Unix() > expiry {
		return http.StatusUnauthorized, errors.New("Refresh-Token expired")
	}

	newRefreshToken, err := generateRandomString(32)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Replacing old token
	_, err = context.db.Exec(fmt.Sprintf("UPDATE refresh SET refresh_token='%s', expiry=%d WHERE refresh_token='%s'",
		newRefreshToken, time.Now().Unix()+86400*7, oldRefreshToken))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	w.Header().Set("Refresh-Token", newRefreshToken)

	// Creating a new JWT
	tokenString, err := createToken(username)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	iat := time.Now().Unix()
	exp := iat + 86400*3
	_, err = fmt.Fprint(w, fmt.Sprintf(`{
		"token": "%s",
		"iss": "connectx",
		"sub": "%s",
		"iat": "%d",
		"exp": "%d"
	}`, tokenString, username, iat, exp))

	return http.StatusOK, nil
}

func loginHandler(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	w.Header().Set("Content-Type", "application/json")

	credential, status, err := verifyJson[Login](w, r)
	if status != 200 {
		return status, err
	}

	sum := sha256.Sum256([]byte(credential.Password))
	// Checking if the username really exist with the corresponding password
	row := context.db.QueryRow(fmt.Sprintf("SELECT id, username, is_parent, created_at FROM user WHERE username='%s' AND password='%x'", credential.Username, sum))
	user := new(User)
	err = row.Scan(&user.ID, &user.Username, &user.IsParent, &user.CreatedAt)
	if err != nil {
		return http.StatusUnauthorized, err
	}

	tokenString, err := createToken(credential.Username)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	refreshToken, err := generateRandomString(32)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	w.Header().Set("Refresh-Token", refreshToken)

	_, err = context.db.Exec(fmt.Sprintf("INSERT INTO refresh (refresh_token, expiry, user_id) VALUES ('%s', %d, %d) ON DUPLICATE KEY UPDATE refresh_token='%s', expiry=%d",
		refreshToken, time.Now().Unix()+86400*7, user.ID, refreshToken, time.Now().Unix()+86400*7))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	iat := time.Now().Unix()
	exp := iat + 86400*3
	_, err = fmt.Fprint(w, fmt.Sprintf(`{
		"token": "%s",
		"iss": "connectx",
		"sub": "%s",
		"iat": "%d",
		"exp": "%d"
	}`, tokenString, credential.Username, iat, exp))

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func verifyHandler(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return http.StatusUnauthorized, errors.New("No access token")
	}

	err := verifyToken(token)
	if err != nil {
		return http.StatusUnauthorized, err
	}

	return http.StatusOK, nil
}

var getAllUsers = getAllItemsFactory[User]("user", scanUserRows)

var getUserByID = getItemByIDFactory[User]("user", scanUserRow)

var createUser = createItemFactory[User, UserCreate](
	"user",
	func(item *UserCreate) string {
		sum := sha256.Sum256([]byte(item.Password))
		return fmt.Sprintf("INSERT INTO user (username, password, is_parent, created_at) VALUES ('%s', '%s', %t, %d)",
			item.Username, fmt.Sprintf("%x", sum), item.IsParent, time.Now().Unix())
	},
	scanUserRow,
)

var updateUser = updateItemFactory[User, UserUpdate](
	"user",
	func(item *UserUpdate) string {
		return fmt.Sprintf("UPDATE user SET username = '%s', is_parent = %t",
			item.Username, item.IsParent)
	},
	scanUserRow,
)

var deleteUserByID = deleteItemByIDFactory("user")
