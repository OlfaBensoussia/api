package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/OlfaBensoussia/api/models"

	"github.com/gorilla/mux"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: home Page !")
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	for _, Task := range models.Tasks {
		if Task.Id == key {
			json.NewEncoder(w).Encode(Task)
			break
		}
		return
	}
	json.NewEncoder(w).Encode(&models.Task{})
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request and unmarshal it into a new Task struct
	task := models.Task{}
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.Id = strconv.Itoa(len(models.Tasks) + 1)
	// update our global models.Tasks array to include our new Task
	models.Tasks = append(models.Tasks, task)
	json.NewEncoder(w).Encode(&task)

}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request and unmarshal it into a new Task struct
	reqBody, _ := ioutil.ReadAll(r.Body)
	newTask := models.Task{}
	json.Unmarshal(reqBody, &newTask)
	key := newTask.Id
	for index, Task := range models.Tasks {
		if Task.Id == key {
			models.Tasks[index] = newTask
		}
	}
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for index, Task := range models.Tasks {
		if Task.Id == id {
			models.Tasks = append(models.Tasks[:index], models.Tasks[index+1:]...)
			break
		}
	}
}
