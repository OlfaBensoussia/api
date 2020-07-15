package db

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/OlfaBensoussia/back-etcd/models"
	"go.etcd.io/etcd/clientv3"
)

//Db is an instance of the client etcd
var Db *clientv3.Client

// func InitEtcd() (*clientv3.Client, error) {
// 	cli, err := clientv3.New(clientv3.Config{
// 		Endpoints:   []string{"http://127.0.0.0.2379"},
// 		DialTimeout: 2 * time.Second,
// 	})
// 	if err != nil {
// 		err = errors.New("etcd instanciation error")
// 	}
// 	return cli, err
// }

func InitEtcd() *clientv3.Client {

	c := clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 2 * time.Second,
	}
	fmt.Println("Creating a new etcd client")

	cli, err := clientv3.New(c)
	if err != nil {
		fmt.Println("etcd instanciation error : %s", err)
	}
	return cli
}

func GetTaskDb(key string) (models.Task, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	res, err := Db.Get(ctx, "task."+key)
	task := models.Task{}
	if err != nil {
		fmt.Println("etcd get error")
	}
	// add check if key is not there
	json.Unmarshal(res.Kvs[0].Value, &task)
	cancel()
	return task, err
}

func CreateTaskDb(task models.Task) (models.Task, error) {
	var res, err = json.Marshal(task)
	if err != nil {
		return task, err
	}
	key := "tasks." + task.Id
	value := string(res)
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	_, err = Db.Put(ctx, key, value)
	if err != nil {
		fmt.Println("etcd put error")
	}
	cancel()
	return task, err
}
