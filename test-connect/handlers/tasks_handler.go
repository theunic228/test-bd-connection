package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"test-connect/gets"
)

func TasksHandler(w http.ResponseWriter, r *http.Request) {

	tasks, err := gets.GetTasks()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error getting tasks", http.StatusInternalServerError)
		return
	}

	statuses, err := gets.GetStatuses()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error getting tasks + statuses", http.StatusInternalServerError)
		return
	}

	employees, err := gets.GetEmployees()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error getting tasks + employees", http.StatusInternalServerError)
		return
	}

	//Создаем структуру для передачи в шаблон
	data := struct {
		Tasks     []gets.Tasks
		Statuses  []gets.Statuses
		Employees []gets.Employees
	}{
		Tasks:     tasks,
		Statuses:  statuses,
		Employees: employees,
	}

	tmpl, err := template.ParseFiles("templates/tasks_page.html")

	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	// Отправляем данные в шаблон
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}
