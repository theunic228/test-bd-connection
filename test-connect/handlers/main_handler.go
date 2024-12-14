package handlers

import (
	"html/template"
	"net/http"
)

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("index").Parse(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Главная страница</title>
		</head>
		<body>
			<h1>Добро пожаловать!</h1>
			<p>Выберите таблицу, нажав на кнопку ниже:</p>
			<form action="http://localhost:8080/products" method="get" style="display:inline;">
				<button type="submit">products</button>
			</form>
			<form action="http://localhost:8080/test" method="get" style="display:inline;">
				<button type="submit">test</button>
			</form>
		</body>
		</html>
	`)
	if err != nil {
		http.Error(w, "Ошибка при создании шаблона", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Ошибка при обработке шаблона", http.StatusInternalServerError)
	}
}
