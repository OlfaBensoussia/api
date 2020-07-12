package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


// Task - Our struct for all Tasks
type Task struct {
	Id    string `json:"Id"`
	Title string `json:"Title"`
	Desc  string `json:"desc"`
}

var Tasks []Task

func initMockTasks() {
	Tasks = []Task{
		Task{Id: "1", Title: "task 1", Desc: "Task Description"},
		Task{Id: "2", Title: "task 2", Desc: "Task Description"},
	}
}


func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page!")
	fmt.Println("Endpoint Hit: home Page")
}

func getAllTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Endpoint Hit: return all Tasks")
	json.NewEncoder(w).Encode(Tasks)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	for _, Task := range Tasks {
		if Task.Id == key {
			json.NewEncoder(w).Encode(Task)
		}
	}
}

func createTask(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request and unmarshal it into a new Task struct
	reqBody, _ := ioutil.ReadAll(r.Body)
	// _ = json.NewDecoder(r.Body).Decode(&task)
	var task Task
	json.Unmarshal(reqBody, &task)
	task.Id = strconv.Itoa(len(Tasks) + 1)
	// update our global Tasks array to include our new Task
	Tasks = append(Tasks, task)
	json.NewEncoder(w).Encode(task)

}

func updateTask(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request and unmarshal it into a new Task struct
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newTask Task
	json.Unmarshal(reqBody, &newTask)
	key := newTask.Id
	for index, Task := range Tasks {
		if Task.Id == key {
			Tasks[index] = newTask
			// Tasks[index].Title = newTask.Title
			// Tasks[index].Desc = newTask.Desc
		}
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for index, Task := range Tasks {
		if Task.Id == id {
			Tasks = append(Tasks[:index], Tasks[index+1:]...)
			break
		}
	}
}
