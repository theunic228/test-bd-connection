package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/main_page.html")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Ошибка при создании шаблона", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		fmt.Println(err)
		http.Error(w, "Ошибка при обработке шаблона", http.StatusInternalServerError)
	}
}
