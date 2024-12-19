package handlers

import (
	"html/template"
	"net/http"
	"test-connect/gets"
)

func HistoryBaseHandler(w http.ResponseWriter, r *http.Request) {
	// Исправляем вызов функции на существующий GetHistoryBase
	history_base, err := gets.GetHistoryBase()
	if err != nil {
		http.Error(w, "Error getting departments", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/history_base_page.html")

	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, history_base)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}
