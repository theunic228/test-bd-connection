package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"test-connect/gets"

	_ "github.com/lib/pq"
)

func AddAuthorizationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	emp, err := gets.GetEmployees()
	if err != nil {
		fmt.Println(err)
	}

	// taskDetails, err := gets.GetTaskDetails()
	// if err != nil {
	// 	fmt.Println(err)
	// 	http.Error(w, "Error getting task details", http.StatusInternalServerError)
	// 	return
	// }

	for _, e := range emp {
		if username == "admin" && password == "123" {
			tmpl, err := template.ParseFiles("templates/main_page.html")
			if err != nil {
				fmt.Println(err)
				http.Error(w, "Ошибка при создании шаблона", http.StatusInternalServerError)
				return
			}

			if err := tmpl.Execute(w, nil); err != nil {
				fmt.Println(err)
				http.Error(w, "Ошибка при обработке шаблона", http.StatusInternalServerError)
			}

			//http.Redirect(w, r, "/main", http.StatusSeeOther)
			return
		} else if username == e.Email && password == e.Password {
			fmt.Println("вы вошли под "+username, password)

			// data := struct {
			// 	Employees   gets.Employees
			// 	TaskDetails []gets.TaskDetails
			// }{
			// 	Employees:   e,
			// 	TaskDetails: taskDetails,
			// }

			tmpl, err := template.ParseFiles("templates/employees_main_page.html")
			if err != nil {
				fmt.Println(err)
				http.Error(w, "Ошибка при создании шаблона", http.StatusInternalServerError)
				return
			}
			fmt.Println(e)
			if err := tmpl.Execute(w, e); err != nil {
				fmt.Println(err)
				http.Error(w, "Ошибка при обработке шаблона", http.StatusInternalServerError)
			}

			//http.Redirect(w, r, "/employees_main", http.StatusSeeOther)
			return
		}
	}

}
