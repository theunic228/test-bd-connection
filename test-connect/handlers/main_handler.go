package handlers

import (
	"html/template"
	"net/http"
)

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/main_page.html")
	if err != nil {
		http.Error(w, "Ошибка при создании шаблона", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Ошибка при обработке шаблона", http.StatusInternalServerError)
	}
}
