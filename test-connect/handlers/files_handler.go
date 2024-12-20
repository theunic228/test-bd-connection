package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"test-connect/gets"
)

func FilesHandler(w http.ResponseWriter, r *http.Request) {
	files, err := gets.GetFiles()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error getting files", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/files_page.html")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, files)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}
