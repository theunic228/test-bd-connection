package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"test-connect/gets"
)

func TaskCommentsHandler(w http.ResponseWriter, r *http.Request) {

	task_comments, err := gets.GetTaskComments()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error getting task_comments", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/task_comments_page.html")

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, task_comments)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}
