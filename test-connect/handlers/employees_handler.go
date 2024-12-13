package handlers

import (
	"html/template"
	"net/http"
	"test-connect/gets"
)

func EmployeesHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем список сотрудников из пакета gets
	employees, err := gets.GetEmployees()
	if err != nil {
		http.Error(w, "Error getting employees", http.StatusInternalServerError)
		return
	}

	// Получаем список отделов
	departments, err := gets.GetDepartments()
	if err != nil {
		http.Error(w, "Error getting departments", http.StatusInternalServerError)
		return
	}

	//Создаем структуру для передачи в шаблон
	data := struct {
		Employees   []gets.Employees
		Departments []gets.Departments
	}{
		Employees:   employees,
		Departments: departments,
	}

	tmpl, err := template.New("employees").Parse(`
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
			<h1>Employees List</h1>
			<table border="1">
				<tr>
					<th>Employee_id</th>
					<th>First_Name</th>
					<th>Last_Name</th>
					<th>Email</th>
					<th>Password</th>
					<th>Department_Id</th>
					<th>Hired_Date</th>
				</tr>
				{{range .Employees}}
				<tr>
					<td>{{.Employee_Id}}</td>
					<td>{{.First_Name}}</td>
					<td>{{.Last_Name}}</td>
					<td>{{.Email}}</td>
					<td>{{.Password}}</td>
					<td>{{.Department_Id}}</td>
					<td>{{.Hired_Date}}</td>
				</tr>
				{{end}}
			</table>
			<br>
			<h2>Add New Employees</h2>
			<form action="/add_employees" method="POST">
				<label for="First_Name">First Name:</label>
				<input type="text" id="First_Name" name="First_Name" required><br>
				<label for="Last_Name">Last Name:</label>
				<input type="text" id="Last_Name" name="Last_Name" required><br>
				<label for="Email">Email:</label>
				<input type="text" id="Email" name="Email" required><br>
				<label for="Password">Password:</label>
				<input type="text" id="Password" name="Password" required><br>

				<label for="Department_Id">Department:</label>
				<select id="Department_Id" name="Department_Id" required>
					<option value=""></option>
					{{range .Departments}}
						<option value="{{.Department_Id}}">{{.Name}}</option>
					{{end}}
				</select><br>


				<label for="Hired_Date">Hired Date:</label>
				<input type="text" id="Hired_Date" name="Hired_Date" required><br>
				<input type="submit" value="Add employees">
			</form>
			<form action="http://localhost:8080/" method="get" style="display:inline;">
				<button type="submit">Main</button>
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
