package handlers

import (
	"html/template"
	"net/http"
	"test-connect/gets"
)

func DepartmentsHandler(w http.ResponseWriter, r *http.Request) {
	departments, err := gets.GetDepartments()
	if err != nil {
		http.Error(w, "Error getting departments", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/departments_page.html")

	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, departments)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}
