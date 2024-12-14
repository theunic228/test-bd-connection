package main

import (
	"fmt"
	"log"
	"net/http"
	"test-connect/handlers"

	_ "github.com/lib/pq" // импортируем pq для работы с PostgreSQL
)

func main() {
	// Запуск HTTP-сервера
	http.HandleFunc("/", handlers.MainPageHandler)

	http.HandleFunc("/departments", handlers.DepartmentsHandler)
	http.HandleFunc("/employees", handlers.EmployeesHandler)
	http.HandleFunc("/files", handlers.FilesHandler)
	http.HandleFunc("/history_base", handlers.HistoryBaseHandler)
	http.HandleFunc("/statuses", handlers.StatusesHandler)
	http.HandleFunc("/task_assignees", handlers.TaskAssigneesHandler)
	http.HandleFunc("/task_comments", handlers.TaskCommentsHandler)
	http.HandleFunc("/task_history", handlers.TaskHistoryHandler)
	http.HandleFunc("/tasks", handlers.TasksHandler)

	//http.HandleFunc("/add_departments", handlers.AddDepartmentsHandler)

	// Начало работы сервера на порту 8080
	fmt.Println("Server started at http://localhost:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
