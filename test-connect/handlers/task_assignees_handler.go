package handlers

import (
	"html/template"
	"net/http"
	"test-connect/gets"
)

func TaskAssigneesHandler(w http.ResponseWriter, r *http.Request) {
	// Исправляем вызов функции на существующий GetTaskAssignees
	task_assignees, err := gets.GetTaskAssignees()
	if err != nil {
		http.Error(w, "Error getting task_assignees", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/task_assignees_page.html")

	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, task_assignees)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}
