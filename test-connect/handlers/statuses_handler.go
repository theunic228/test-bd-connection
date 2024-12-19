package handlers

import (
	"html/template"
	"net/http"
	"test-connect/gets"
)

func StatusesHandler(w http.ResponseWriter, r *http.Request) {
	// Исправляем вызов функции на существующий GetStatuses
	statuses, err := gets.GetStatuses()
	if err != nil {
		http.Error(w, "Error getting departments", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/statuses_page.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, statuses)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}
