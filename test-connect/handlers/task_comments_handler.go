package handlers

import (
	"html/template"
	"net/http"
	"test-connect/gets"
)

func TaskCommentsHandler(w http.ResponseWriter, r *http.Request) {
	// Исправляем вызов функции на существующий GetDepartments
	task_comments, err := gets.GetTaskComments()
	if err != nil {
		http.Error(w, "Error getting task_comments", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/task_comments_page.html")

	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, task_comments)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}
