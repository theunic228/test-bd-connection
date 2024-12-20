package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"test-connect/gets"
)

func ManagmentsHandler(w http.ResponseWriter, r *http.Request) {
	managments, err := gets.GetManagments()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error getting managments", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/managments_page.html")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, managments)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}
