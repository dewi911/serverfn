package main

import (
	"fmt"
	"github.com/sourcegraph/conc/pool"
	"log"
	"net/http"

	"serverfn/handlers"
	"serverfn/models"
)

func main() {
	manager := models.NewTaskManager()
	taskQueue := make(chan *models.Task, 10)
	handler := handlers.NewTaskHandler(manager, taskQueue)

	p := pool.New().WithMaxGoroutines(10)

	http.HandleFunc("/task", handler.CreateTaskHandler(p))
	http.HandleFunc("/task/", handler.GetTaskHandler)

	fmt.Println("Server is listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

	p.Wait()
}
