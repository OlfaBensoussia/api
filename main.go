package main

import (
	"github.com/OlfaBensoussia/mock-back-end/routes"
	"github.com/OlfaBensoussia/mock-back-end/models"
)

func main() {
	routes.handleRequests()
	models.initMockTasks()
}

