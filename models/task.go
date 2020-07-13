package models

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
