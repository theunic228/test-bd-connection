package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"test-connect/gets"
)

func HistoryBaseHandler(w http.ResponseWriter, r *http.Request) {

	history_base, err := gets.GetHistoryBase()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error getting history", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/history_base_page.html")

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, history_base)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}
