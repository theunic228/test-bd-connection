package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"test-connect/gets"
)

func EmployeesHandler(w http.ResponseWriter, r *http.Request) {
	employees, err := gets.GetEmployees()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error getting employees", http.StatusInternalServerError)
		return
	}

	departments, err := gets.GetDepartments()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error getting departments", http.StatusInternalServerError)
		return
	}

	positions, err := gets.GetPositions()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error getting positions", http.StatusInternalServerError)
		return
	}

	//Создаем структуру для передачи в шаблон
	data := struct {
		Employees   []gets.Employees
		Departments []gets.Departments
		Positions   []gets.Positions
	}{
		Employees:   employees,
		Departments: departments,
		Positions:   positions,
	}

	tmpl, err := template.ParseFiles("templates/employees_page.html")

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	// Отправляем данные в шаблон
	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}
