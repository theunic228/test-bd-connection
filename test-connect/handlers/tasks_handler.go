package handlers

import (
	"html/template"
	"net/http"
	"test-connect/gets"
)

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	// Исправляем вызов функции на существующий GetTasks
	tasks, err := gets.GetTasks()
	if err != nil {
		http.Error(w, "Error getting tasks", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.New("tasks").Parse(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>tasks</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					background-color: #f9f9f9;
					color: #333;
					margin: 0;
					padding: 20px;
					display: flex;
					flex-direction: column;
					align-items: center;
				}
				h1 {
					color: #388E3C; /* Зеленый цвет для заголовка */
					margin-bottom: 20px;
					font-size: 2.5rem;
				}
				h2 {
					color: #388E3C; /* Зеленый цвет для заголовка "Add New Department" */
					font-size: 2rem;
					margin-top: 20px;
				}
				table {
					border-collapse: collapse;
					width: 80%;
					margin-bottom: 20px;
					background-color: white;
					box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
				}
				th, td {
					border: 1px solid #ddd;
					padding: 12px 15px;
					text-align: left;
					font-size: 1rem;
				}
				th {
					background-color: #388E3C; /* Зеленый фон для заголовков таблицы */
					color: white;
					font-size: 1.1rem;
				}
				tr:nth-child(even) {
					background-color: #f2f2f2;
				}
				tr:hover {
					background-color: #ddd;
				}
				button {
					background-color: #388E3C; /* Зеленая кнопка */
					color: white;
					border: none;
					padding: 12px 25px;
					font-size: 16px;
					border-radius: 5px;
					cursor: pointer;
					transition: background-color 0.3s ease, transform 0.2s ease;
				}
				button:hover {
					background-color: #2C6E2F;
					transform: translateY(-2px);
				}
				button:active {
					transform: translateY(1px);
				}
				form {
					margin-top: 30px;
					width: 80%;
					background-color: #fff;
					padding: 25px;
					border-radius: 8px;
					box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
				}
				form label {
					display: block;
					margin-bottom: 10px;
					font-weight: bold;
					font-size: 1rem;
					color: #388E3C; /* Зеленый цвет для меток */
				}
				form input[type="text"], form input[type="submit"] {
					width: 100%;
					padding: 12px;
					margin-bottom: 20px;
					border: 1px solid #ccc;
					border-radius: 8px;
					font-size: 1rem;
					box-sizing: border-box;
				}
				form input[type="text"]:focus {
					border-color: #388E3C; /* Зеленый цвет рамки при фокусе */
					outline: none;
					box-shadow: 0 0 5px rgba(56, 142, 60, 0.4); /* Легкое зеленое свечение */
				}
				form input[type="submit"] {
					background-color: #388E3C; /* Зеленая кнопка в форме */
					color: white;
					border: none;
					cursor: pointer;
					padding: 12px 20px;
					font-size: 1rem;
					border-radius: 8px;
					transition: background-color 0.3s ease;
				}
				form input[type="submit"]:hover {
					background-color: #2C6E2F;
				}
			</style>
		</head>
		<body>
			<h1>Tasks List</h1>
			<table border="1">
				<tr>
					<th>Task_Id</th>
					<th>Title</th>
					<th>Description</th>
					<th>Status</th>
					<th>Created_By</th>
					<th>Created_At</th>
					<th>Updated_At</th>
					<th>Due_Date</th>
				</tr>
				{{range .}}
				<tr>
					<td>{{.Task_Id}}</td>
					<td>{{.Title}}</td>
					<td>{{.Description}}</td>
					<td>{{.Status}}</td>
					<td>{{.Created_By}}</td>
					<td>{{.Created_At}}</td>
					<td>{{.Updated_At}}</td>
					<td>{{.Due_Date}}</td>
				</tr>
				{{end}}
			</table>
			<br>
			<h2>Add New Tasks</h2>
			<form action="/add_tasks" method="POST">
				<label for="Title">Title:</label>
				<input type="text" id="Title" name="Title" required><br>
				<label for="Description">Description:</label>
				<input type="text" id="Description" name="Description" required><br>
				<label for="Status">Status:</label>
				<input type="text" id="Status" name="Status" required><br>
				<label for="Created_By">Created By:</label>
				<input type="text" id="Created_By" name="Created_By" required><br>
				<label for="Due_Date">Due Date:</label>
				<input type="text" id="Due_Date" name="Due_Date" required><br>
				<input type="submit" value="Add tasks">
			</form>
			<form action="http://localhost:8080/" method="get" style="display:inline;">
				<button type="submit">main</button>
			</form>
		</body>
		</html>
	`)

	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, tasks)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}
