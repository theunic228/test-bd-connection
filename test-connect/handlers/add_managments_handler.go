package handlers

import (
	"net/http"
	"test-connect/adds"

	_ "github.com/lib/pq"
)

func AddManagmentsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	Managment_Name := r.FormValue("Managment_Name")

	if Managment_Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	err := adds.AddManagments(Managment_Name)
	if err != nil {
		http.Error(w, "Error adding Managment", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/managments", http.StatusSeeOther)
}
