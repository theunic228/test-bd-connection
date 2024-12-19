package handlers

import (
	"html/template"
	"net/http"
	"test-connect/gets"
)

func TaskHistoryHandler(w http.ResponseWriter, r *http.Request) {
	// Исправляем вызов функции на существующий GetTaskHistory
	task_history, err := gets.GetTaskHistory()
	if err != nil {
		http.Error(w, "Error getting task_history", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/task_history_page.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, task_history)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}
