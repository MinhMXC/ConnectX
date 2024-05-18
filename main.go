package main

import (
	"database/sql"
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

func (handler AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	status, err := handler.H(handler.AppContext, w, r)
	if err != nil {
		log.Printf("HTTP %d: %q", status, err)
		switch status {
		case http.StatusNotFound:
			http.NotFound(w, r)
		case http.StatusUnauthorized:
			http.Error(w, http.StatusText(status), status)
		case http.StatusInternalServerError:
			http.Error(w, http.StatusText(status), status)
		default:
			http.Error(w, http.StatusText(status), status)
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

	mux.HandleFunc("GET /level", AppHandler{context, getAllLevels}.ServeHTTP)
	mux.HandleFunc("GET /level/{id}", AppHandler{context, getLevelByID}.ServeHTTP)
	mux.HandleFunc("DELETE /level/{id}", AppHandler{context, deleteLevelByID}.ServeHTTP)

	mux.HandleFunc("GET /subject", AppHandler{context, getAllSubjects}.ServeHTTP)
	mux.HandleFunc("GET /subject/{id}", AppHandler{context, getSubjectByID}.ServeHTTP)
	mux.HandleFunc("DELETE /subject/{id}", AppHandler{context, deleteSubjectByID}.ServeHTTP)

	mux.HandleFunc("GET /user", AppHandler{context, getAllUsers}.ServeHTTP)
	mux.HandleFunc("GET /user/{id}", AppHandler{context, getUserByID}.ServeHTTP)
	mux.HandleFunc("DELETE /user/{id}", AppHandler{context, deleteUserByID}.ServeHTTP)

	mux.HandleFunc("GET /tutor", AppHandler{context, getAllTutors}.ServeHTTP)
	mux.HandleFunc("GET /tutor/{id}", AppHandler{context, getTutorByID}.ServeHTTP)
	mux.HandleFunc("DELETE /tutor/{id}", AppHandler{context, deleteTutorByID}.ServeHTTP)

	mux.HandleFunc("GET /tuition_center", AppHandler{context, getAllTuitionCenters}.ServeHTTP)
	mux.HandleFunc("GET /tuition_center/{id}", AppHandler{context, getTuitionCenterByID}.ServeHTTP)
	mux.HandleFunc("DELETE /tuition_center/{id}", AppHandler{context, deleteTuitionCenterByID}.ServeHTTP)

	mux.HandleFunc("GET /rate", AppHandler{context, getAllRates}.ServeHTTP)
	mux.HandleFunc("GET /rate/{id}", AppHandler{context, getRateByID}.ServeHTTP)
	mux.HandleFunc("DELETE /rate/{id}", AppHandler{context, deleteRateByID}.ServeHTTP)

	mux.HandleFunc("GET /request", AppHandler{context, getAllRequests}.ServeHTTP)
	mux.HandleFunc("GET /request/{id}", AppHandler{context, getRequestByID}.ServeHTTP)
	mux.HandleFunc("DELETE /request/{id}", AppHandler{context, deleteRequestByID}.ServeHTTP)

	mux.HandleFunc("GET /qualification", AppHandler{context, getAllQualifications}.ServeHTTP)
	mux.HandleFunc("GET /qualification/{id}", AppHandler{context, getQualificationByID}.ServeHTTP)
	mux.HandleFunc("DELETE /qualification/{id}", AppHandler{context, deleteQualificationByID}.ServeHTTP)

	fmt.Println("Server Started. Listening on port 8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err.Error())
	}
}
