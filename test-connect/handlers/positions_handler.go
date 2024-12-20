package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"test-connect/gets"
)

func PositionsHandler(w http.ResponseWriter, r *http.Request) {
	positions, err := gets.GetPositions()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error getting positions", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/positions_page.html")

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, positions)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}
