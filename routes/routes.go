package routes

import (
	"log"
	"net/http"

	"github.com/OlfaBensoussia/api/controllers"
	"github.com/gorilla/mux"
)

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", controllers.HomePage)
	myRouter.HandleFunc("/tasks", controllers.GetAllTasks).Methods("GET")
	myRouter.HandleFunc("/task/{id}", controllers.GetTask).Methods("GET")
	myRouter.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")
	myRouter.HandleFunc("/task/{id}", controllers.UpdateTask).Methods("PUT")
	myRouter.HandleFunc("/task/{id}", controllers.DeleteTask).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
