package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/OlfaBensoussia/back-etcd/controllers"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

//InitRoute is used to initialize all routes
func InitRoutes() {
	//init router
	r := mux.NewRouter().StrictSlash(true)
	//Route handlers for tasks
	r.HandleFunc("/", homeLink)
	// r.HandleFunc("/api/task", controllers.GetAllTasks).Methods("GET")
	r.HandleFunc("/task/{id}", controllers.GetTask).Methods("GET")
	r.HandleFunc("/task", controllers.CreateTask).Methods("POST")
	// r.HandleFunc("/task/{id}", controllers.UpdateTask).Methods("PUT")
	// r.HandleFunc("/task/{id}", controllers.Task.DeleteTask).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}
