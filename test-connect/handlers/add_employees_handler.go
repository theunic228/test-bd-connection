package handlers

import (
	"fmt"
	"net/http"
	"test-connect/adds"

	_ "github.com/lib/pq"
)

func AddEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	First_Name := r.FormValue("First_Name")
	Last_Name := r.FormValue("Last_Name")
	Patronymic := r.FormValue("Patronymic")
	Email := r.FormValue("Email")
	Password := r.FormValue("Password")
	Department_Id := r.FormValue("Department_Id")
	Position_id := r.FormValue("Position_Id")

	if First_Name == "" || Last_Name == "" || Patronymic == "" || Email == "" || Password == "" || Department_Id == "" || Position_id == "" {
		http.Error(w, "data are required", http.StatusBadRequest)
		return
	}

	err := adds.AddEmployees(First_Name, Last_Name, Patronymic, Email, Password, Department_Id, Position_id)
	if err != nil {

		http.Error(w, "Error adding Employees", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	http.Redirect(w, r, "/employees", http.StatusSeeOther)
}
