package handlers

import (
	"net/http"
	"test-connect/adds"

	_ "github.com/lib/pq"
)

func AddStatusesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	Status_Name := r.FormValue("Status_Name")

	if Status_Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	err := adds.AddStatuses(Status_Name)
	if err != nil {
		http.Error(w, "Error adding Status", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/statuses", http.StatusSeeOther)
}
