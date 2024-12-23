package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"test-connect/gets"
)

func DepartmentsHandler(w http.ResponseWriter, r *http.Request) {
	departments, err := gets.GetDepartments()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error getting departments", http.StatusInternalServerError)
		return
	}

	managments, err := gets.GetManagments()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error getting managments", http.StatusInternalServerError)
		return
	}

	data := struct {
		Departments []gets.Departments
		Managments  []gets.Managments
	}{
		Departments: departments,
		Managments:  managments,
	}

	tmpl, err := template.ParseFiles("templates/departments_page.html")

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}
