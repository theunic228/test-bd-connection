package handlers

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq" // импортируем pq для работы с PostgreSQL
)

func AddAuthorizationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println(username, password)

	if username == "admin" && password == "123" {
		fmt.Println(username, password)
		http.Redirect(w, r, "/main", http.StatusSeeOther)
		return
	}
}
