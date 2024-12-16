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

	statuses, err := gets.GetStatuses()
	if err != nil {
		http.Error(w, "Error getting tasks", http.StatusInternalServerError)
		return
	}

	employees, err := gets.GetEmployees()
	if err != nil {
		http.Error(w, "Error getting tasks", http.StatusInternalServerError)
		return
	}

	//Создаем структуру для передачи в шаблон
	data := struct {
		Tasks     []gets.Tasks
		Statuses  []gets.Statuses
		Employees []gets.Employees
	}{
		Tasks:     tasks,
		Statuses:  statuses,
		Employees: employees,
	}

	tmpl, err := template.New("tasks").Parse(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>departments</title>
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

				/* Стили для выпадающего списка */
				select {
					width: 100%;
					padding: 12px;
					border: 1px solid #ccc;
					border-radius: 8px;
					font-size: 1rem;
					background-color: white;
					box-sizing: border-box;
					transition: all 0.3s ease;
				}

				/* Стиль при фокусе на выпадающем списке */
				select:focus {
					border-color: #388E3C;
					outline: none;
					box-shadow: 0 0 5px rgba(56, 142, 60, 0.4); /* Легкое зеленое свечение */
				}

				/* Кастомизация стрелки */
				select::-ms-expand {
					display: none;
				}

				/* Добавление красивых стрелок для вебkit-браузеров */
				select::-webkit-select {
					-webkit-appearance: none;
					appearance: none;
					background-image: url('data:image/svg+xml,%3Csvg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 12 12"%3E%3Cpath fill="none" stroke="black" stroke-width="1.5" d="M2 4l4 4 4-4"%3E%3C/path%3E%3C/svg%3E');
					background-position: right 10px center;
					background-repeat: no-repeat;
					background-size: 12px 12px;
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
				{{range .Tasks}}
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
				<select id="Status" name="Status" required>
					<option value=""></option>
					{{range .Statuses}}
						<option value="{{.Status_Id}}">{{.Status_Name}}</option>
					{{end}}
				</select><br>
				
				<label for="Created_By">Created By:</label>
				<select id="Created_By" name="Created_By" required>
					<option value=""></option>
					{{range .Employees}}
						<option value="{{.Employee_Id}}">{{.Last_Name}}</option>
					{{end}}
				</select><br>

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

	// Отправляем данные в шаблон
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}
