package main

import (
	"log"
	"net/http"

	"goproject/db"
	"goproject/user"

	"github.com/gorilla/mux"
)

func main() {
	db.InitGorm()
	initMigration()
	initializeRoutes()
}

func initializeRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/users", user.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", user.GetUser).Methods("GET")
	r.HandleFunc("/users", user.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", user.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", user.DeleteUser).Methods("DELETE")

	log.Println("Running server on port 9000")
	log.Fatal(http.ListenAndServe(":9000", r))
}

func initMigration() {
	user.InitMigration()
}
