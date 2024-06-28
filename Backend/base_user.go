package main

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type BaseUser struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	UserType  int    `json:"user_type"`
	CreatedAt int    `json:"created_at"`
}

type BaseUserCreate struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	UserType        int    `json:"user_type"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type BaseUserUpdate struct {
	Password string `json:"password"`
}

func generateRandomString(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}

func signupHandler(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	credential, status, err := verifyJson[BaseUserCreate](w, r)
	if status != 200 {
		return status, err
	}

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(credential.Email) {
		return http.StatusBadRequest, errors.New("Invalid email address")
	}

	// Password checking
	// Password must be at least 8 character & contains uppercase, lowercase and number
	if credential.Password != credential.ConfirmPassword {
		return http.StatusBadRequest, errors.New("Password and Confirm Password does not match")
	}

	if len(credential.Password) < 8 {
		return http.StatusBadRequest, errors.New("Password must be at least 8 characters")
	}
	var uppercaseCount int
	var lowercaseCount int
	var numberCount int
	for _, char := range credential.Password {
		if 65 <= char && char <= 90 {
			uppercaseCount++
		} else if 97 <= char && char <= 122 {
			lowercaseCount++
		} else if 48 <= char && char <= 57 {
			numberCount++
		}
	}
	if uppercaseCount == 0 {
		return http.StatusBadRequest, errors.New("Password must contains at least one uppercase letter")
	}
	if lowercaseCount == 0 {
		return http.StatusBadRequest, errors.New("Password must contains at least one lowercase letter")
	}
	if numberCount == 0 {
		return http.StatusBadRequest, errors.New("Password must contains at least a number")
	}

	sum := sha256.Sum256([]byte(credential.Password))

	_, err = context.db.Exec(fmt.Sprintf("INSERT INTO base_user (email, password, created_at) VALUES ('%s', '%s', %d)",
		credential.Email, fmt.Sprintf("%x", sum), time.Now().Unix()))
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return http.StatusBadRequest, errors.New("This email has already been registered")
		}

		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func loginHandler(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	w.Header().Set("Content-Type", "application/json")

	credential, status, err := verifyJson[Login](w, r)
	if status != 200 {
		return status, err
	}

	sum := sha256.Sum256([]byte(credential.Password))

	// Checking if the email really exist with the corresponding password
	row := context.db.QueryRow(fmt.Sprintf("SELECT id, email, user_type, created_at FROM base_user WHERE email='%s' AND password='%x'", credential.Email, sum))
	baseUser := new(BaseUser)
	err = row.Scan(&baseUser.ID, &baseUser.Email, &baseUser.UserType, &baseUser.CreatedAt)
	if err != nil {
		return http.StatusUnauthorized, errors.New("Invalid credentials")
	}

	err = attachToken(baseUser.ID, credential.Email, w)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	refreshToken, err := generateRandomString(32)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	w.Header().Set("Refresh-Token", refreshToken)

	_, err = context.db.Exec(fmt.Sprintf("INSERT INTO refresh (refresh_token, expiry, base_user_id) VALUES ('%s', %d, %d) ON DUPLICATE KEY UPDATE refresh_token='%s', expiry=%d",
		refreshToken, time.Now().Unix()+86400*7, baseUser.ID, refreshToken, time.Now().Unix()+86400*7))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	err = json.NewEncoder(w).Encode(baseUser)

	return http.StatusOK, nil
}

func changePasswordHandler(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	claims, err := authorize(w, r, false)
	if err != nil {
		return http.StatusUnauthorized, err
	}

	data, status, err := verifyJson[BaseUserUpdate](w, r)
	if status != 200 {
		return status, err
	}

	// Change the password hash
	sum := sha256.Sum256([]byte(data.Password))
	_, err = context.db.Exec(fmt.Sprintf("UPDATE base_user SET password = '%s' WHERE id = %d", fmt.Sprintf("%x", sum), claims.ID))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Logout the user
	_, err = context.db.Exec(
		fmt.Sprintf("UPDATE refresh SET expiry = 0 WHERE base_user_id = %d", claims.ID))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

// BaseUser should still be able to be authorized using the JWT before logging out
// Thus, frontend will have to remove the authorization key from memory
func logoutHandler(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	claims, err := authorize(w, r, false)
	if err != nil {
		return http.StatusUnauthorized, err
	}

	_, err = context.db.Exec(
		fmt.Sprintf("UPDATE refresh SET expiry = 0 WHERE base_user_id = %d", claims.ID))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func deleteBaseUserHandler(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	claims, err := authorize(w, r, false)
	if err != nil {
		return http.StatusUnauthorized, err
	}

	_, err = context.db.Exec(
		fmt.Sprintf("DELETE FROM base_user WHERE email = '%s'", claims.Username))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func verifyHandler(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	_, err := authorize(w, r, true)
	if err != nil {
		return http.StatusUnauthorized, err
	}

	return http.StatusOK, nil
}

// Should be accessed only when the JWT expire
// Return a new JWT and a new refreshToken
func refreshHandler(context *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	oldRefreshToken := r.Header.Get("Refresh-Token")
	if oldRefreshToken == "" {
		return http.StatusUnauthorized, errors.New("Refresh-Token header missing")
	}

	row := context.db.QueryRow("SELECT refresh.expiry, base_user.email, base_user.id FROM refresh INNER JOIN base_user ON refresh.base_user_id = base_user.id WHERE refresh_token=?", oldRefreshToken)
	var expiry int64
	var email string
	var id int
	err := row.Scan(&expiry, &email, &id)
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
	err = attachToken(id, email, w)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
