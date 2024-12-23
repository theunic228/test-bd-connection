package handlers

import (
	"net/http"
	"test-connect/adds"

	_ "github.com/lib/pq"
)

func AddPositionsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	Position_Name := r.FormValue("Position_Name")

	if Position_Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	err := adds.AddPositions(Position_Name)
	if err != nil {
		http.Error(w, "Error adding Position", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/positions", http.StatusSeeOther)
}
