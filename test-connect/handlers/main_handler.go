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
			<title>main</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					background-color: #f9f9f9;
					color: #333;
					margin: 0;
					padding: 0;
					display: flex;
					flex-direction: column;
					align-items: center;
					justify-content: center;
					height: 100vh;
				}
				h1 {
					color: #4CAF50;
					margin-bottom: 20px;
				}
				p {
					margin-bottom: 30px;
					font-size: 18px;
				}
				.button-container {
					display: flex;
					flex-wrap: wrap;
					justify-content: center;
					gap: 10px;
				}
				button {
					background-color: #4CAF50;
					color: white;
					border: none;
					padding: 10px 20px;
					font-size: 16px;
					border-radius: 5px;
					cursor: pointer;
					transition: background-color 0.3s ease, transform 0.2s ease;
				}
				button:hover {
					background-color: #45a049;
					transform: translateY(-2px);
				}
				button:active {
					transform: translateY(1px);
				}
			</style>
		</head>
		<body>
			<h1>Добро пожаловать!</h1>
			<p>Выберите таблицу, нажав на кнопку ниже:</p>
			<form action="http://localhost:8080/departments" method="get" style="display:inline;">
				<button type="submit">departments</button>
			</form><br>
			<form action="http://localhost:8080/employees" method="get" style="display:inline;">
				<button type="submit">employees</button>
			</form><br>
			<form action="http://localhost:8080/files" method="get" style="display:inline;">
				<button type="submit">files</button>
			</form><br>
			<form action="http://localhost:8080/history_base" method="get" style="display:inline;">
				<button type="submit">history base</button>
			</form><br>
			<form action="http://localhost:8080/statuses" method="get" style="display:inline;">
				<button type="submit">statuses</button>
			</form><br>
			<form action="http://localhost:8080/task_assignees" method="get" style="display:inline;">
				<button type="submit">task assignees</button>
			</form><br>
			<form action="http://localhost:8080/task_comments" method="get" style="display:inline;">
				<button type="submit">task comments</button>
			</form><br>
			<form action="http://localhost:8080/task_history" method="get" style="display:inline;">
				<button type="submit">task history</button>
			</form><br>
			<form action="http://localhost:8080/tasks" method="get" style="display:inline;">
				<button type="submit">tasks</button>
			</form><br>
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
