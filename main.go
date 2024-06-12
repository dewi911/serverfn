package main

import (
	"fmt"
	"log"
	"net/http"

	"serverfn/handlers"
	"serverfn/models"
)

func main() {
	manager := models.NewTaskManager()
	handler := handlers.NewTaskHandler(manager)

	http.HandleFunc("/task", handler.CreateTaskHandler)
	http.HandleFunc("/task/", handler.GetTaskHandler)

	fmt.Println("Server is listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
