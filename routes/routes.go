package routes

import (
	"log"
	"net/http"
	"github.com/OlfaBensoussia/mock-back-end/models"

	"github.com/gorilla/mux"
)

func handleRequests() {
	task := models.Task{}
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", task.homePage())
	myRouter.HandleFunc("/tasks", task.getAllTasks())
	myRouter.HandleFunc("/task/{id}", task.getTask()).Methods("GET")
	myRouter.HandleFunc("/tasks", task.createTask()).Methods("POST")
	myRouter.HandleFunc("/task/{id}", task.updateTask()).Methods("PUT")
	myRouter.HandleFunc("/task/{id}", task.deleteTask()).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
