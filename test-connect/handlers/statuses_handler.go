package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"test-connect/gets"
)

func StatusesHandler(w http.ResponseWriter, r *http.Request) {

	statuses, err := gets.GetStatuses()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error getting departments", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/statuses_page.html")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, statuses)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}
