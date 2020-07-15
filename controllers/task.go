package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/OlfaBensoussia/back-etcd/db"
	"github.com/OlfaBensoussia/back-etcd/models"
	"github.com/gorilla/mux"
)

//get a single Task by ID from etcd
func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key := params["id"]
	task, _ := db.GetTaskDb(key)
	json.NewEncoder(w).Encode(task)
	// return task
}

//create new Task in etcd
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	task, _ = db.CreateTaskDb(task)
	json.NewEncoder(w).Encode(task)
	// return task
}

/*
//get all tasks from etcd
// func GetAllTaskS() ([]Task, error) {
// 	Tasks, err := GetAllTasksDb()
// 	return Tasks, err
// }

//update a Task in etcd
func UpdateTask(Task Task) (Task, error) {
	res, err := UpdateTaskDb(Task)
	return res, err
}

//delete a Task in etcd
func DeleteTask(Task Task) (Task, error) {
	res, err := DeleteTaskDb(Task)
	return res, err
}
*/
