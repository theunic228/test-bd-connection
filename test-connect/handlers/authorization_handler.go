package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func AuthorizationHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/authorization_page.html")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		fmt.Println(err)
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}

}
