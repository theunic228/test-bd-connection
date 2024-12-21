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
	Email := r.FormValue("Email")
	Password := r.FormValue("Password")
	Department_Id := r.FormValue("Department_Id")
	Hired_Date := r.FormValue("Hired_Date")
	if First_Name == "" || Last_Name == "" || Email == "" || Password == "" || Department_Id == "" || Hired_Date == "" {
		http.Error(w, "data are required", http.StatusBadRequest)
		return
	}

	err := adds.AddEmployees(First_Name, Last_Name, Email, Password, Department_Id, Hired_Date)
	if err != nil {
		http.Error(w, "Error adding Employees", http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	http.Redirect(w, r, "/employees", http.StatusSeeOther)
}
