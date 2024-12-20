package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"test-connect/gets"
)

func TaskAssigneesHandler(w http.ResponseWriter, r *http.Request) {

	task_assignees, err := gets.GetTaskAssignees()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error getting task_assignees", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/task_assignees_page.html")

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, task_assignees)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}
