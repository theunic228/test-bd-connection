package handlers

import (
	"html/template"
	"net/http"
	"test-connect/gets"
)

func EmployeesHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем список сотрудников из пакета gets
	employees, err := gets.GetEmployees()
	if err != nil {
		http.Error(w, "Error getting employees", http.StatusInternalServerError)
		return
	}

	// Получаем список отделов
	departments, err := gets.GetDepartments()
	if err != nil {
		http.Error(w, "Error getting departments", http.StatusInternalServerError)
		return
	}

	//Создаем структуру для передачи в шаблон
	data := struct {
		Employees   []gets.Employees
		Departments []gets.Departments
	}{
		Employees:   employees,
		Departments: departments,
	}

	tmpl, err := template.ParseFiles("templates/employees_page.html")

	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	// Отправляем данные в шаблон
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}
