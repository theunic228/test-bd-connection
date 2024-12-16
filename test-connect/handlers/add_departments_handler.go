package handlers

import (
	"net/http"
	"test-connect/adds"

	_ "github.com/lib/pq" // импортируем pq для работы с PostgreSQL
)

func AddDepartmentsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	Name := r.FormValue("Name")
	Description := r.FormValue("Description")
	if Name == "" || Description == "" {
		http.Error(w, "Name and Description are required", http.StatusBadRequest)
		return
	}

	err := adds.AddDepartments(Name, Description)
	if err != nil {
		http.Error(w, "Error adding Department", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/departments", http.StatusSeeOther)
}
