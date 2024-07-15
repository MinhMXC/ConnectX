package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

type AppContext struct {
	db *sql.DB
}

type AppHandler struct {
	*AppContext
	H func(*AppContext, http.ResponseWriter, *http.Request) (int, error)
}

func handleOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Accept, Authorization, content-type, X-Requested-With")
	w.Header().Add("Access-Control-Max-Age", "86400")
	w.WriteHeader(http.StatusOK)
	return
}

func (handler AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Accept, Authorization, content-type, X-Requested-With")
	w.Header().Add("Access-Control-Expose-Headers", "Authorization, Refresh-Token")
	w.Header().Add("Access-Control-Max-Age", "86400")

	status, err := handler.H(handler.AppContext, w, r)

	if err != nil {
		log.Printf("HTTP %d: %q", status, err)

		var errorText string

		switch status {
		case http.StatusNotFound:
			errorText = http.StatusText(status)
		case http.StatusUnauthorized:
			errorText = err.Error()
		case http.StatusUnprocessableEntity:
			errorText = err.Error()
		case http.StatusBadRequest:
			errorText = err.Error()
		case http.StatusInternalServerError:
			errorText = http.StatusText(status)
		default:
			http.Error(w, http.StatusText(status), status)
		}

		w.WriteHeader(status)
		errJson := json.NewEncoder(w).Encode(map[string]string{"error": errorText})
		if errJson != nil {
			http.Error(w, "", http.StatusInternalServerError)
		}
	}
}

func main() {
	args := os.Args[1:]
	if len(args) > 0 && args[0] == "setup" {
		setup()
		return
	}

	mux := http.NewServeMux()

	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/orbital")
	if err != nil {
		panic(err.Error())
	}

	context := &AppContext{db: db}
	defer db.Close()

	mux.HandleFunc("OPTIONS /", handleOptions)

	mux.HandleFunc("GET /level", AppHandler{context, getAllLevels}.ServeHTTP)
	mux.HandleFunc("GET /level/{id}", AppHandler{context, getLevelByID}.ServeHTTP)
	mux.HandleFunc("POST /level", AppHandler{context, createLevel}.ServeHTTP)
	mux.HandleFunc("PATCH /level/{id}", AppHandler{context, updateLevel}.ServeHTTP)
	mux.HandleFunc("DELETE /level/{id}", AppHandler{context, deleteLevelByID}.ServeHTTP)

	mux.HandleFunc("GET /subject", AppHandler{context, getAllSubjects}.ServeHTTP)
	mux.HandleFunc("GET /subject/{id}", AppHandler{context, getSubjectByID}.ServeHTTP)
	mux.HandleFunc("POST /subject", AppHandler{context, createSubject}.ServeHTTP)
	mux.HandleFunc("PATCH /subject/{id}", AppHandler{context, updateSubject}.ServeHTTP)
	mux.HandleFunc("DELETE /subject/{id}", AppHandler{context, deleteSubjectByID}.ServeHTTP)

	mux.HandleFunc("POST /user", AppHandler{context, signupHandler}.ServeHTTP)
	mux.HandleFunc("POST /user/login", AppHandler{context, loginHandler}.ServeHTTP)
	mux.HandleFunc("GET /user/logout", AppHandler{context, logoutHandler}.ServeHTTP)
	mux.HandleFunc("GET /user/verify", AppHandler{context, verifyHandler}.ServeHTTP)
	mux.HandleFunc("GET /user/refresh", AppHandler{context, refreshHandler}.ServeHTTP)
	mux.HandleFunc("PATCH /user/change_password", AppHandler{context, changePasswordHandler}.ServeHTTP)
	mux.HandleFunc("DELETE /user", AppHandler{context, deleteBaseUserHandler}.ServeHTTP)

	mux.HandleFunc("GET /user", AppHandler{context, getAllUsers}.ServeHTTP)
	mux.HandleFunc("GET /user/{id}", AppHandler{context, getUserByID}.ServeHTTP)
	mux.HandleFunc("POST /user/setup", AppHandler{context, userSetup}.ServeHTTP)
	mux.HandleFunc("PATCH /user/{id}", AppHandler{context, updateUser}.ServeHTTP)

	mux.HandleFunc("GET /tutor", AppHandler{context, getAllTutors}.ServeHTTP)
	mux.HandleFunc("GET /tutor/{id}", AppHandler{context, getTutorByID}.ServeHTTP)
	mux.HandleFunc("POST /tutor/setup", AppHandler{context, tutorSetup}.ServeHTTP)
	mux.HandleFunc("PATCH /tutor/{id}", AppHandler{context, updateTutor}.ServeHTTP)

	mux.HandleFunc("GET /tuition_center", AppHandler{context, getAllTuitionCenters}.ServeHTTP)
	mux.HandleFunc("GET /tuition_center/{id}", AppHandler{context, getTuitionCenterByID}.ServeHTTP)
	mux.HandleFunc("POST /tuition_center/setup", AppHandler{context, tuitionCenterSetup}.ServeHTTP)
	mux.HandleFunc("PATCH /tuition_center/{id}", AppHandler{context, updateTuitionCenter}.ServeHTTP)

	mux.HandleFunc("GET /rate", AppHandler{context, getAllRates}.ServeHTTP)
	mux.HandleFunc("GET /rate/{id}", AppHandler{context, getRateByID}.ServeHTTP)
	mux.HandleFunc("POST /rate", AppHandler{context, createRate}.ServeHTTP)
	mux.HandleFunc("PATCH /rate/{id}", AppHandler{context, updateRate}.ServeHTTP)
	mux.HandleFunc("DELETE /rate/{id}", AppHandler{context, deleteRateByID}.ServeHTTP)

	mux.HandleFunc("GET /request", AppHandler{context, getAllRequests}.ServeHTTP)
	mux.HandleFunc("GET /request/{id}", AppHandler{context, getRequestByID}.ServeHTTP)
	mux.HandleFunc("POST /request", AppHandler{context, createRequest}.ServeHTTP)
	mux.HandleFunc("PATCH /request/{id}", AppHandler{context, updateRequest}.ServeHTTP)
	mux.HandleFunc("DELETE /request/{id}", AppHandler{context, deleteRequestByID}.ServeHTTP)

	mux.HandleFunc("GET /qualification", AppHandler{context, getAllQualifications}.ServeHTTP)
	mux.HandleFunc("GET /qualification/tutor/{id}", AppHandler{context, getQualificationByTutorID}.ServeHTTP)
	mux.HandleFunc("GET /qualification/{id}", AppHandler{context, getQualificationByID}.ServeHTTP)
	mux.HandleFunc("POST /qualification", AppHandler{context, createQualification}.ServeHTTP)
	mux.HandleFunc("PATCH /qualification/{id}", AppHandler{context, updateQualification}.ServeHTTP)
	mux.HandleFunc("DELETE /qualification/{id}", AppHandler{context, deleteQualificationByID}.ServeHTTP)

	fmt.Println("Server Started. Listening on port 8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err.Error())
	}
}
