package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/tasks", getAllTasks)
	myRouter.HandleFunc("/task/{id}", getTask).Methods("GET")
	myRouter.HandleFunc("/tasks", createTask).Methods("POST")
	myRouter.HandleFunc("/task/{id}", updateTask).Methods("PUT")
	myRouter.HandleFunc("/task/{id}", deleteTask).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
