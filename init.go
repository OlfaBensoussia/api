package main

import (
	"github.com/OlfaBensoussia/back-etcd/db"
	"github.com/OlfaBensoussia/back-etcd/routes"
)

func main() {
	db.InitEtcd()
	routes.InitRoutes()
}
