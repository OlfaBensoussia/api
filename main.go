package main

import (
	"fmt"

	"github.com/OlfaBensoussia/api/models"
	"github.com/OlfaBensoussia/api/routes"
)

func main() {
	models.InitMockTasks()
	routes.HandleRequests()
	fmt.Println("app is running ")
}
