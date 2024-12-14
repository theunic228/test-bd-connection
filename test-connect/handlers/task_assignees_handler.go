package handlers

import (
	"html/template"
	"net/http"
	"test-connect/gets"
)

func TaskAssigneesHandler(w http.ResponseWriter, r *http.Request) {
	// Исправляем вызов функции на существующий GetTaskAssignees
	task_assignees, err := gets.GetTaskAssignees()
	if err != nil {
		http.Error(w, "Error getting task_assignees", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.New("task_assignees").Parse(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>task_assignees</title>
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
			<h1>Task assignees List</h1>
			<table border="1">
				<tr>
					<th>Task_Id</th>
					<th>Employee_Id</th>
				</tr>
				{{range .}}
				<tr>
					<td>{{.Task_Id}}</td>
					<td>{{.Employee_Id}}</td>
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

	err = tmpl.Execute(w, task_assignees)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}
