package main

import (
	"log"
	"main/db"
	"main/models"
	"main/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db.DbConnection()

	db.DB.AutoMigrate(&models.Task{})
	db.DB.AutoMigrate(&models.User{})

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", routes.HomeHandler)
	// users routes
	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/users/{id}", routes.DeleteUsersHandler).Methods("DELETE")

	//tasks routes
	router.HandleFunc("/tasksById/{id}", routes.GetTasksByIdHandler).Methods("GET")
	router.HandleFunc("/tasks", routes.CreateTaskHandler).Methods("POST")
	router.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	router.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")

	log.Fatal(
		http.ListenAndServe(":3000", router),
	)
}
