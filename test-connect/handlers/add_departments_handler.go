package handlers

import (
	"net/http"
	"test-connect/adds"

	_ "github.com/lib/pq"
)

func AddDepartmentsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	Department_Name := r.FormValue("Department_Name")
	Managment_Id := r.FormValue("Managment_Id")

	if Department_Name == "" || Managment_Id == "" {
		http.Error(w, "Name and Description are required", http.StatusBadRequest)
		return
	}

	err := adds.AddDepartments(Department_Name, Managment_Id)
	if err != nil {
		http.Error(w, "Error adding Department", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/departments", http.StatusSeeOther)
}
