package handlers

import (
	"html/template"
	"net/http"
	"test-connect/gets"
)

func EmployeesHandler(w http.ResponseWriter, r *http.Request) {
	// Исправляем вызов функции на существующий GetEmployees
	employees, err := gets.GetEmployees()
	if err != nil {
		http.Error(w, "Error getting employees", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.New("employees").Parse(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>employees</title>
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
            color: #4CAF50;
            margin-bottom: 20px;
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
            padding: 8px;
            text-align: left;
        }
        th {
            background-color: #4CAF50;
            color: white;
        }
        tr:nth-child(even) {
            background-color: #f2f2f2;
        }
        tr:hover {
            background-color: #ddd;
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
        form {
            margin-top: 20px;
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
				{{range .}}
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

	err = tmpl.Execute(w, employees)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}
