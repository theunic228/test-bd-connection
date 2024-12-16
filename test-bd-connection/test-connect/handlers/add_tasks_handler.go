package handlers

import (
	"fmt"
	"net/http"
	"test-connect/adds"

	_ "github.com/lib/pq" // импортируем pq для работы с PostgreSQL
)

func AddTasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	Title := r.FormValue("Title")
	Description := r.FormValue("Description")
	Status := r.FormValue("Status")
	Created_By := r.FormValue("Created_By")
	Due_Date := r.FormValue("Due_Date")
	if Title == "" || Description == "" || Status == "" || Created_By == "" || Due_Date == "" {
		http.Error(w, "data are required", http.StatusBadRequest)
		return
	}

	err := adds.AddTasks(Title, Description, Status, Created_By, Due_Date)
	if err != nil {
		http.Error(w, "Error adding tasks", http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}
