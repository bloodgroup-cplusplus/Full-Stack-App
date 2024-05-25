package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// connect to database
	db, err := sql.Open("postgres", os.Getenv(("DATABASE_URL")))
	if err != nil {
		log.Fatal(err)

	}
	defer db.Close()

	// create table if not exists

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY,name TEXT,email TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	// create router
	router := mux.NewRouter()
	router.HandleFunc("/api/go/users", getUsers(db)).Methods("GET")
	router.HandleFunc("/api/go/users", createUser(db)).Methods("POST")
	router.HandleFunc("/api/go/users/{id}", getUser(db)).Methods("GET")
	router.HandleFunc("/api/go/users/{id}", getUser(db)).Methods("PUT")
	router.HandleFunc("/api/go/users/{id}", getUser(db)).Methods("DELETE")

}
